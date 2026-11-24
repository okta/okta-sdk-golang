package okta

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/stretchr/testify/require"
)

// Okta rejects a client assertion whose exp is more than one hour in the
// future, with no allowance for clock skew. The assertion must leave a margin
// so a client clock that runs a little ahead of Okta's still works.
// See https://github.com/okta/okta-sdk-golang/issues/573.
func Test_Client_Assertion_Expiry_Leaves_Clock_Skew_Margin(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	keyPEM := string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}))
	signer, err := createKeySigner(keyPEM, "test-kid")
	require.NoError(t, err)

	before := time.Now()
	assertion, err := createClientAssertion("https://example.okta.com", "0oaclient", signer)
	require.NoError(t, err)

	parsed, err := jwt.ParseSigned(assertion, []jose.SignatureAlgorithm{jose.RS256})
	require.NoError(t, err)
	var claims ClientAssertionClaims
	require.NoError(t, parsed.Claims(&key.PublicKey, &claims))

	// Okta's view of "now" if the client clock is one minute ahead.
	oktaNow := before.Add(-time.Minute)
	require.LessOrEqual(t, claims.Expiry.Time().Sub(oktaNow), time.Hour,
		"exp must stay within Okta's one hour limit when the client clock is a minute ahead")
	require.Equal(t, clientAssertionLifetime, claims.Expiry.Time().Sub(claims.IssuedAt.Time()))
	require.True(t, claims.Expiry.Time().After(before), "assertion must not already be expired")

	require.Equal(t, "0oaclient", claims.Issuer)
	require.Equal(t, "0oaclient", claims.Subject)
	require.Equal(t, "https://example.okta.com/oauth2/v1/token", claims.Audience)
	require.NotEmpty(t, claims.ID)
	require.Equal(t, "test-kid", parsed.Headers[0].KeyID)
}
