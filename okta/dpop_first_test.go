package okta

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetAccessTokenForPrivateKeyDpopFirst verifies the DPoP-first handshake:
// a proof is attached to the first token request (so Okta never sees, and never
// logs, a proofless invalid_dpop_proof rejection), and the result adapts to the
// token_type Okta returns.
func TestGetAccessTokenForPrivateKeyDpopFirst(t *testing.T) {
	t.Run("dpop-required app: proof sent up front, nonce challenge, no proofless request", func(t *testing.T) {
		prooflessSeen := false
		tokenRequests := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/v1/token", r.URL.Path)
			if r.Header.Get("DPoP") == "" {
				prooflessSeen = true
			}
			switch tokenRequests {
			case 0:
				tokenRequests++
				w.Header().Set("Dpop-Nonce", "nonce-1")
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"error":"use_dpop_nonce"}`))
			default:
				tokenRequests++
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(&RequestAccessToken{AccessToken: "dpop-token", TokenType: "DPoP", ExpiresIn: 3600})
			}
		}))
		defer server.Close()

		token, nonce, key, err := getAccessTokenForPrivateKey(server.Client(), server.URL, "assertion", "ua", []string{"okta.users.read"}, 2, 1, "", nil)
		require.NoError(t, err)
		require.False(t, prooflessSeen, "no token request should be sent without a DPoP proof")
		require.Equal(t, 2, tokenRequests, "expected 2 token requests (nonce challenge + success), not the old 3")
		require.NotNil(t, token)
		require.Equal(t, "DPoP", token.TokenType)
		require.Equal(t, "dpop-token", token.AccessToken)
		require.Equal(t, "nonce-1", nonce)
		require.NotNil(t, key, "DPoP key must be returned for per-request proofs")
	})

	t.Run("non-dpop app: bearer token returned, dpop material dropped", func(t *testing.T) {
		tokenRequests := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/oauth2/v1/token", r.URL.Path)
			require.NotEmpty(t, r.Header.Get("DPoP"), "a proof is still attached up front; Okta ignores it for a non-DPoP app")
			tokenRequests++
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(&RequestAccessToken{AccessToken: "bearer-token", TokenType: "Bearer", ExpiresIn: 3600})
		}))
		defer server.Close()

		token, nonce, key, err := getAccessTokenForPrivateKey(server.Client(), server.URL, "assertion", "ua", []string{"okta.users.read"}, 2, 1, "", nil)
		require.NoError(t, err)
		require.Equal(t, 1, tokenRequests, "a non-DPoP app needs only one token request")
		require.NotNil(t, token)
		require.Equal(t, "Bearer", token.TokenType)
		require.Empty(t, nonce, "no nonce for a Bearer token")
		require.Nil(t, key, "DPoP key must be dropped so no per-request proof is attached")
	})

	t.Run("with a signer, a fresh client assertion is minted for the nonce retry (PrivateKey/JWK modes)", func(t *testing.T) {
		pem, err := generatePrivateKeyPem()
		require.NoError(t, err)
		signer, err := createKeySigner(string(pem), "kid-1")
		require.NoError(t, err)

		var assertions []string
		count := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.NoError(t, r.ParseForm())
			assertions = append(assertions, r.FormValue("client_assertion"))
			if count == 0 {
				count++
				w.Header().Set("Dpop-Nonce", "n1")
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"error":"use_dpop_nonce"}`))
				return
			}
			count++
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(&RequestAccessToken{AccessToken: "dpop-token", TokenType: "DPoP", ExpiresIn: 3600})
		}))
		defer server.Close()

		_, _, key, err := getAccessTokenForPrivateKey(server.Client(), server.URL, "ignored-when-signer-present", "ua", []string{"okta.users.read"}, 2, 1, "client-id", signer)
		require.NoError(t, err)
		require.NotNil(t, key)
		require.Len(t, assertions, 2)
		require.NotEmpty(t, assertions[0])
		require.NotEqual(t, assertions[0], assertions[1], "the nonce retry must use a freshly minted (single-use) client assertion")
	})

	t.Run("without a signer (JWT mode), the caller-supplied assertion is replayed on the nonce retry", func(t *testing.T) {
		// signer==nil is the JWTAuth path: the SDK has no way to mint a fresh
		// assertion, so it reuses the caller-supplied one across the use_dpop_nonce
		// retry. This locks in the documented limitation that a Require-DPoP org
		// rejects the replayed jti; callers needing DPoP should use PrivateKeyAuth/JWKAuth.
		var assertions []string
		count := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.NoError(t, r.ParseForm())
			assertions = append(assertions, r.FormValue("client_assertion"))
			if count == 0 {
				count++
				w.Header().Set("Dpop-Nonce", "n1")
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"error":"use_dpop_nonce"}`))
				return
			}
			count++
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(&RequestAccessToken{AccessToken: "dpop-token", TokenType: "DPoP", ExpiresIn: 3600})
		}))
		defer server.Close()

		_, _, key, err := getAccessTokenForPrivateKey(server.Client(), server.URL, "caller-assertion", "ua", []string{"okta.users.read"}, 2, 1, "", nil)
		require.NoError(t, err)
		require.NotNil(t, key)
		require.Len(t, assertions, 2)
		require.Equal(t, "caller-assertion", assertions[0])
		require.Equal(t, assertions[0], assertions[1], "with no signer the same caller-supplied assertion is replayed on the nonce retry")
	})

	t.Run("use_dpop_nonce without a Dpop-Nonce header errors instead of looping forever", func(t *testing.T) {
		hits := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hits++
			// Always demand a nonce but never supply one — must not loop.
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":"use_dpop_nonce"}`))
		}))
		defer server.Close()

		token, _, key, err := getAccessTokenForPrivateKey(server.Client(), server.URL, "assertion", "ua", []string{"okta.users.read"}, 2, 1, "", nil)
		require.Error(t, err)
		require.Nil(t, token)
		require.Nil(t, key)
		require.LessOrEqual(t, hits, 2, "must not retry the nonce challenge more than once")
	})

	t.Run("hard failure surfaces an error instead of looping", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"invalid_client"}`))
		}))
		defer server.Close()

		token, _, key, err := getAccessTokenForPrivateKey(server.Client(), server.URL, "assertion", "ua", []string{"okta.users.read"}, 2, 1, "", nil)
		require.Error(t, err)
		require.Nil(t, token)
		require.Nil(t, key)
	})
}
