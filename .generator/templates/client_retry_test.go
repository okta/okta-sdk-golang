package okta

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func generatePrivateKeyPem() ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	return pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}), nil
}

func TestAPIClient_doWithRetries(t *testing.T) {
	t.Run("successful request without retries", func(t *testing.T) {
		attemptCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attemptCount++
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 3
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "SSWS"

		client := NewAPIClient(cfg)

		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		require.NoError(t, err)

		resp, err := client.doWithRetries(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, "success", string(body))
		require.Equal(t, 1, attemptCount)
	})

	t.Run("retry on 429 Too Many Requests", func(t *testing.T) {
		attemptCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attemptCount++
			if attemptCount == 1 {
				// First attempt returns 429
				w.Header().Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
				w.Header().Set("X-Rate-Limit-Reset", "1")
				w.Header().Set("X-Okta-Request-Id", "test-request-id")
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("rate limited"))
				return
			}

			// Second attempt succeeds
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success after retry"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 3
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "SSWS"

		client := NewAPIClient(cfg)

		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		require.NoError(t, err)

		resp, err := client.doWithRetries(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, 2, attemptCount)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, "success after retry", string(body))

		// Verify retry headers were added
		require.Equal(t, "test-request-id", req.Header.Get("X-Okta-Retry-For"))
		require.Equal(t, "1", req.Header.Get("X-Okta-Retry-Count"))
	})

	t.Run("no retry on permanent error", func(t *testing.T) {
		attemptCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attemptCount++
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 3
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "SSWS"

		client := NewAPIClient(cfg)

		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		require.NoError(t, err)

		resp, err := client.doWithRetries(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		// Should only attempt once since 400 is not a retryable error
		require.Equal(t, 1, attemptCount)
	})

	t.Run("request body is rewound on retry", func(t *testing.T) {
		attemptCount := 0
		requestBodies := make([][]byte, 0)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attemptCount++
			body, _ := io.ReadAll(r.Body)
			requestBodies = append(requestBodies, body)
			if attemptCount == 1 {
				// First attempt returns 429
				w.Header().Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
				w.Header().Set("X-Rate-Limit-Reset", "1")
				w.Header().Set("X-Okta-Request-Id", "test-request-id")
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("rate limited"))
				return
			}

			// Second attempt succeeds
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 3
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "SSWS"

		client := NewAPIClient(cfg)

		bodyContent := []byte("test body content")
		req, err := http.NewRequest("POST", server.URL+"/test", bytes.NewReader(bodyContent))
		require.NoError(t, err)

		resp, err := client.doWithRetries(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		// Verify body was sent correctly in both attempts
		require.Equal(t, 2, len(requestBodies))
		require.Equal(t, bodyContent, requestBodies[0])
		require.Equal(t, bodyContent, requestBodies[1])
		require.Equal(t, 2, attemptCount)
	})

	t.Run("token retries for PrivateKey mode", func(t *testing.T) {
		attemptCount := 0
		authRequestCount := 0
		mockTokenResponse := &RequestAccessToken{
			AccessToken: "test-access-token",
			ExpiresIn:   3600,
			TokenType:   "DPoP",
			Scope:       "test-scope",
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/oauth2/v1/token" && authRequestCount == 0 {
				authRequestCount++
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid_dpop_proof"))
				return
			}
			if r.URL.Path == "/oauth2/v1/token" && authRequestCount == 1 {
				authRequestCount++
				w.Header().Set("Dpop-Nonce", "test-nonce")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("use_dpop_nonce"))
				return
			}
			if r.URL.Path == "/oauth2/v1/token" && authRequestCount == 2 {
				w.Header().Set("Content-Type", "application/json")
				require.NoError(t, json.NewEncoder(w).Encode(mockTokenResponse))
				return
			}
			if r.URL.Path == "/oauth2/v1/token" && authRequestCount == 3 {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal_server_error"))
				return
			}

			attemptCount++
			if attemptCount == 1 {
				// First attempt returns 429
				w.Header().Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
				w.Header().Set("X-Rate-Limit-Reset", "1")
				w.Header().Set("X-Okta-Request-Id", "test-request-id")
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("rate limited"))
				return
			}

			// Second attempt succeeds
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 3
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "PrivateKey"
		cfg.Okta.Client.OrgUrl = server.URL
		cfg.Okta.Client.ClientId = "test-client-id"
		privateKeyPem, err := generatePrivateKeyPem()
		require.NoError(t, err)
		cfg.Okta.Client.PrivateKey = string(privateKeyPem)

		client := NewAPIClient(cfg)

		// Set some values in token cache
		client.tokenCache.Set(AccessTokenCacheKey, "cached-token", 5*time.Minute)
		client.tokenCache.Set(DpopAccessTokenNonce, "cached-nonce", 5*time.Minute)
		client.tokenCache.Set(DpopAccessTokenPrivateKey, "cached-key", 5*time.Minute)

		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		require.NoError(t, err)

		// The request will attempt retry, and cache should be cleared on retry
		_, err = client.doWithRetries(context.Background(), req)
		require.NoError(t, err)

		require.Equal(t, 2, attemptCount)

		refreshed, found := client.tokenCache.Get(AccessTokenCacheKey)
		require.True(t, found)
		require.Equal(t, fmt.Sprintf("%s %s", mockTokenResponse.TokenType, mockTokenResponse.AccessToken), refreshed)

		refreshedNonce, found := client.tokenCache.Get(DpopAccessTokenNonce)
		require.True(t, found)
		require.Equal(t, "test-nonce", refreshedNonce)

		refreshedPrivateKey, found := client.tokenCache.Get(DpopAccessTokenPrivateKey)
		require.True(t, found)

		_, ok := refreshedPrivateKey.(*rsa.PrivateKey)
		require.True(t, ok)
	})

	t.Run("respects max retries limit", func(t *testing.T) {
		attemptCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attemptCount++
			// Always return 429
			w.Header().Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
			w.Header().Set("X-Rate-Limit-Reset", "1")
			w.Header().Set("X-Okta-Request-Id", "test-request-id")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("rate limited"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 2
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "SSWS"

		client := NewAPIClient(cfg)

		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		require.NoError(t, err)

		_, err = client.doWithRetries(context.Background(), req)
		// Should eventually fail after max retries
		require.Error(t, err)
		// Should attempt initial request + max retries
		require.Equal(t, attemptCount, 3)
	})

	t.Run("handles request with nil body", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		}))
		defer server.Close()

		cfg := &Configuration{
			HTTPClient: &http.Client{Timeout: 1 * time.Minute},
			Host:       strings.TrimPrefix(server.URL, "http://"),
			Scheme:     "http",
		}
		cfg.Okta.Client.RateLimit.MaxRetries = 3
		cfg.Okta.Client.RateLimit.MaxBackoff = 1
		cfg.Okta.Client.AuthorizationMode = "SSWS"

		client := NewAPIClient(cfg)

		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		require.NoError(t, err)

		resp, err := client.doWithRetries(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
