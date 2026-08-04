package okta

import (
	"crypto/rand"
	"crypto/rsa"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/stretchr/testify/require"
)

// newTestSigner builds an RSA key and matching jose.Signer for exercising
// getAccessTokenForPrivateKey end-to-end without network. Matches the
// 2048-bit size used by the existing PrivateKey-mode tests.
func newTestSigner(t *testing.T) jose.Signer {
	t.Helper()
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: priv},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	require.NoError(t, err)
	return signer
}

// TestGetAccessTokenForPrivateKey_TokenEndpointErrors regression-tests the bug
// where a 4xx/5xx /oauth2/v1/token response was silently swallowed
// (return nil, "", nil, err with err==nil), causing callers to surface the
// misleading "Empty access token" message instead of Okta's actual error.
func TestGetAccessTokenForPrivateKey_TokenEndpointErrors(t *testing.T) {
	signer := newTestSigner(t)

	cases := []struct {
		name        string
		status      int
		body        string
		contentType string
		wantSubs    []string
	}{
		{
			name:        "401 invalid_client",
			status:      http.StatusUnauthorized,
			contentType: "application/json",
			body:        `{"error":"invalid_client","error_description":"Client authentication failed"}`,
			wantSubs:    []string{"401", "invalid_client", "Client authentication failed"},
		},
		{
			name:        "403 invalid_scope",
			status:      http.StatusForbidden,
			contentType: "application/json",
			body:        `{"error":"invalid_scope","error_description":"One or more scopes are not configured for the authorization server"}`,
			wantSubs:    []string{"403", "invalid_scope", "scopes are not configured"},
		},
		{
			name:        "500 non-JSON body falls back to raw body",
			status:      http.StatusInternalServerError,
			contentType: "text/plain",
			body:        "internal server error",
			wantSubs:    []string{"500", "internal server error"},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/oauth2/v1/token", r.URL.Path)
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer server.Close()

			httpClient := &http.Client{Timeout: 5 * time.Second}
			token, _, _, err := getAccessTokenForPrivateKey(
				httpClient, server.URL, "assertion-stub", "test-ua",
				[]string{"okta.users.read"}, 0, 1,
				"test-client-id", signer,
			)

			require.Error(t, err, "expected non-nil error for status %d", tc.status)
			require.Nil(t, token, "token must be nil when token endpoint returns an error")
			for _, sub := range tc.wantSubs {
				require.Contains(t, err.Error(), sub,
					"error %q must contain %q", err.Error(), sub)
			}
		})
	}
}

// TestGetAccessTokenForPrivateKey_Success is a regression check: a 200 response
// with a well-formed token payload must still return a populated token and a
// nil error (the fix must not break the happy path).
func TestGetAccessTokenForPrivateKey_Success(t *testing.T) {
	signer := newTestSigner(t)

	const tokenBody = `{"token_type":"Bearer","expires_in":3600,"access_token":"test-access-token","scope":"okta.users.read"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/oauth2/v1/token", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(tokenBody))
	}))
	defer server.Close()

	httpClient := &http.Client{Timeout: 5 * time.Second}
	got, nonce, dpopKey, err := getAccessTokenForPrivateKey(
		httpClient, server.URL, "assertion-stub", "test-ua",
		[]string{"okta.users.read"}, 0, 1,
		"test-client-id", signer,
	)

	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "test-access-token", got.AccessToken)
	require.Equal(t, "Bearer", got.TokenType)
	require.Equal(t, 3600, got.ExpiresIn)
	require.Equal(t, "", nonce)
	require.Nil(t, dpopKey)
}

// TestGetAccessTokenForPrivateKey_DpopPassthrough verifies that an
// `invalid_dpop_proof` response still routes into the DPoP retry path. This
// guards against the fix accidentally turning the DPoP delegation into a
// surfaced error.
func TestGetAccessTokenForPrivateKey_DpopPassthrough(t *testing.T) {
	signer := newTestSigner(t)

	const dpopTokenBody = `{"token_type":"DPoP","expires_in":3600,"access_token":"dpop-access-token","scope":"okta.users.read"}`

	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/oauth2/v1/token", r.URL.Path)
		n := calls.Add(1)
		switch n {
		case 1:
			// Initial call without DPoP header — trigger the DPoP path.
			require.Empty(t, r.Header.Get("DPoP"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":"invalid_dpop_proof","error_description":"DPoP proof required"}`))
		case 2:
			// DPoP-bearing retry — must succeed.
			require.NotEmpty(t, r.Header.Get("DPoP"), "DPoP retry must include DPoP header")
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(dpopTokenBody))
		default:
			t.Fatalf("unexpected request #%d", n)
		}
	}))
	defer server.Close()

	httpClient := &http.Client{Timeout: 5 * time.Second}
	got, _, dpopKey, err := getAccessTokenForPrivateKey(
		httpClient, server.URL, "assertion-stub", "test-ua",
		[]string{"okta.users.read"}, 0, 1,
		"test-client-id", signer,
	)

	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "dpop-access-token", got.AccessToken)
	require.Equal(t, "DPoP", got.TokenType)
	require.NotNil(t, dpopKey, "DPoP path must return a private key for proof generation")
	require.Equal(t, int32(2), calls.Load(), "expected exactly one DPoP retry")
}

// TestGetAccessTokenForDpopPrivateKey_NonNonceErrorReturned regression-tests
// the second buggy site: when the DPoP-bearing request itself fails with a
// non-`use_dpop_nonce` error, the helper must surface a real error rather
// than returning (nil, "", nil, nil).
func TestGetAccessTokenForDpopPrivateKey_NonNonceErrorReturned(t *testing.T) {
	signer := newTestSigner(t)

	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/oauth2/v1/token", r.URL.Path)
		n := calls.Add(1)
		switch n {
		case 1:
			// First (non-DPoP) call returns invalid_dpop_proof, triggering DPoP.
			require.Empty(t, r.Header.Get("DPoP"))
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":"invalid_dpop_proof"}`))
		case 2:
			// DPoP retry fails with a real error (not use_dpop_nonce).
			require.NotEmpty(t, r.Header.Get("DPoP"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"invalid_client","error_description":"Client authentication failed"}`))
		default:
			t.Fatalf("unexpected request #%d", n)
		}
	}))
	defer server.Close()

	httpClient := &http.Client{Timeout: 5 * time.Second}
	token, _, _, err := getAccessTokenForPrivateKey(
		httpClient, server.URL, "assertion-stub", "test-ua",
		[]string{"okta.users.read"}, 0, 1,
		"test-client-id", signer,
	)

	require.Error(t, err, "DPoP retry that fails with non-nonce error must surface an error")
	require.Nil(t, token)
	require.Contains(t, err.Error(), "401")
	require.Contains(t, err.Error(), "invalid_client")
}
