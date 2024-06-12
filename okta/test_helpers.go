package okta

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const (
	charSetAlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charSetAlphaLower = "abcdefghijklmnopqrstuvwxyz"
	charSetNumeric    = "0123456789"
	charSetAlpha      = charSetAlphaLower + charSetAlphaUpper + charSetNumeric
	testPrefix        = "SDK_TEST_"
)

func randomEmail() string {
	return randomTestString() + "@example.com"
}

// randStringFromCharSet generates a random string of 15 lower case letters
func randomTestString() string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 15)
	for i := 0; i < 15; i++ {
		result[i] = charSetAlphaLower[rand.Intn(len(charSetAlphaLower))]
	}
	return testPrefix + string(result)
}

// testPassword generates a random string of at least 4 characters in length
func testPassword(length int) string {
	if length < 5 {
		length = 4
	}
	result := make([]byte, length)
	result[0] = charSetAlphaLower[rand.Intn(len(charSetAlphaLower))]
	result[1] = charSetAlphaUpper[rand.Intn(len(charSetAlphaUpper))]
	result[2] = charSetNumeric[rand.Intn(len(charSetNumeric))]
	for i := 3; i < length; i++ {
		result[i] = charSetAlpha[rand.Intn(len(charSetAlpha))]
	}
	return string(result)
}

func testName(name string) string {
	s := fmt.Sprintf("%s %s", randomTestString(), name)
	if len(s) > 50 {
		s = s[:50]
	}
	return s
}

func contain(searchPhrase string, searchList []string) (res bool) {
	for _, searchResult := range searchList {
		if searchPhrase == searchResult {
			res = true
			break
		}
	}
	return
}

type TestFactory struct{}

var testFactory TestFactory

func (t *TestFactory) NewValidTestUserProfile() UserProfile {
	email := randomEmail()
	firstName := "John"
	lastName := "Doe"
	return UserProfile{
		FirstName: NullableString{value: &firstName, isSet: true},
		LastName:  NullableString{value: &lastName, isSet: true},
		Email:     &email,
		Login:     &email,
	}
}

func (t *TestFactory) NewValidTestUserCredentialsWithPassword() *UserCredentials {
	pc := t.NewValidTestPasswordCredential()
	return &UserCredentials{Password: pc}
}

func (t *TestFactory) NewValidTestPasswordCredential() *PasswordCredential {
	p := testPassword(10)
	return &PasswordCredential{Value: &p}
}

func (t *TestFactory) NewValidTestRecoveryQuestionCredential() *RecoveryQuestionCredential {
	question := "How many roads must a man walk down?"
	answer := "forty two"
	return &RecoveryQuestionCredential{Question: &question, Answer: &answer}
}

func (t *TestFactory) NewValidTestIdentityProvider() *IdentityProvider {
	res := IdentityProvider{}
	res.SetType("OIDC")
	res.SetName(randomTestString())
	res.SetProtocol(*t.NewValidTestProtocol())
	res.SetPolicy(*t.NewValidTestIdentityProviderPolicy())
	return &res
}

func (t *TestFactory) NewValidTestProtocol() *Protocol {
	payload := `{
		"algorithms": {
			"request": {
				"signature": {
					"algorithm": "SHA-256",
					"scope": "REQUEST"
				}
			},
			"response": {
				"signature": {
					"algorithm": "SHA-256",
					"scope": "ANY"
				}
			}
		},
		"endpoints": {
			"acs": {
				"binding": "HTTP-POST",
				"type": "INSTANCE"
			},
			"authorization": {
				"binding": "HTTP-REDIRECT",
				"url": "https://idp.example.com/authorize"
			},
			"token": {
				"binding": "HTTP-POST",
				"url": "https://idp.example.com/token"
			},
			"userInfo": {
				"binding": "HTTP-REDIRECT",
				"url": "https://idp.example.com/userinfo"
			},
			"jwks": {
				"binding": "HTTP-REDIRECT",
				"url": "https://idp.example.com/keys"
			}
		},
		"scopes": [
			"openid",
			"profile",
			"email"
		],
		"type": "OIDC",
		"credentials": {
			"client": {
				"client_id": "your-client-id",
				"client_secret": "your-client-secret"
			}
		},
		"issuer": {
			"url": "https://idp.example.com"
		}
	}`
	var ptc Protocol
	err := json.Unmarshal([]byte(payload), &ptc)
	if err != nil {
		return nil
	}
	return &ptc
}

func (t *TestFactory) NewValidTestIdentityProviderPolicy() *IdentityProviderPolicy {
	payload := `{
		"accountLink": {
			"action": "AUTO",
			"filter": null
		},
		"provisioning": {
			"action": "AUTO",
			"conditions": {
				"deprovisioned": {
					"action": "NONE"
				},
				"suspended": {
					"action": "NONE"
				}
			},
			"groups": {
				"action": "NONE"
			}
		},
		"maxClockSkew": 120000,
		"subject": {
			"userNameTemplate": {
				"template": "idpuser.email"
			},
			"matchType": "USERNAME"
		}
	
	}`
	var policy IdentityProviderPolicy
	err := json.Unmarshal([]byte(payload), &policy)
	if err != nil {
		return nil
	}
	return &policy
}

func (t *TestFactory) NewValidTestJsonWebKey() *JsonWebKey {
	keys := []string{`@"MIIDnjCCAoagAwIBAgIGAVG3MN+PMA0GCSqGSIb3DQEBBQUAMIGPMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5p
	YTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEDAOBgNVBAMM
	B2V4YW1wbGUxHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMTUxMjE4MjIyMjMyWhcNMjUxMjE4MjIyMzMyWjCB
	jzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoMBE9r
	dGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRAwDgYDVQQDDAdleGFtcGxlMRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29t
	MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtcnyvuVCrsFEKCwHDenS3Ocjed8eWDv3zLtD2K/iZfE8BMj2wpTf
	n6Ry8zCYey3mWlKdxIybnV9amrujGRnE0ab6Q16v9D6RlFQLOG6dwqoRKuZy33Uyg8PGdEudZjGbWuKCqqXEp+UKALJHV+k4
	wWeVH8g5d1n3KyR2TVajVJpCrPhLFmq1Il4G/IUnPe4MvjXqB6CpKkog1+ThWsItPRJPAM+RweFHXq7KfChXsYE7Mmfuly8s
	DQlvBmQyxZnFHVuiPfCvGHJjpvHy11YlHdOjfgqHRvZbmo30+y0X/oY/yV4YEJ00LL6eJWU4wi7ViY3HP6/VCdRjHoRdr5L/
	DwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCzzhOFkvyYLNFj2WDcq1YqD4sBy1iCia9QpRH3rjQvMKDwQDYWbi6EdOX0TQ/I
	YR7UWGj+2pXd6v0t33lYtoKocp/4lUvT3tfBnWZ5KnObi+J2uY2teUqoYkASN7F+GRPVOuMVoVgm05ss8tuMb2dLc9vsx93s
	Dt+XlMTv/2qi5VPwaDtqduKkzwW9lUfn4xIMkTiVvCpe0X2HneD2Bpuao3/U8Rk0uiPfq6TooWaoW3kjsmErhEAs9bA7xuqo
	1KKY9CdHcFhkSsMhoeaZylZHtzbnoipUlQKSLMdJQiiYZQ0bYL83/Ta9fulr1EERICMFt3GUmtYaZZKHpWSfdJp9"`}
	return &JsonWebKey{X5c: keys}
}

func (t *TestFactory) NewValidTestCSRMetadata() *CsrMetadata {
	payload := `{
		"subject": {
		  "countryName"            : "US",
		  "stateOrProvinceName"    : "California",
		  "localityName"           : "San Francisco",
		  "organizationName"       : "Okta, Inc.",
		  "organizationalUnitName" : "Dev",
		  "commonName"             : "SP Issuer"
		},
		"subjectAltNames": 
		  {
			"dnsNames": [ "dev.example.com" ] 
		  }
	  }`
	var csrm CsrMetadata
	err := json.Unmarshal([]byte(payload), &csrm)
	if err != nil {
		return nil
	}
	return &csrm
}

func (t *TestFactory) NewValidAccessPolicy(name string) *AccessPolicy {
	policyRule := NewPolicyRuleConditions()
	res := AccessPolicy{}
	res.SetType("ACCESS_POLICY")
	res.SetDescription(randomTestString())
	res.SetPriority(int32(1))
	res.SetConditions(*policyRule)
	res.SetName(name)
	return &res
}

func (t *TestFactory) NewValidBasicAuthApplication(label string) *BasicAuthApplication {
	app := NewBasicApplicationSettingsApplication()
	app.SetAuthURL("https://example.com/auth.html")
	app.SetUrl("https://example.com/auth.html")
	setting := NewBasicApplicationSettings()
	setting.SetApp(*app)
	res := BasicAuthApplication{}
	res.SetSettings(*setting)
	res.SetName("template_basic_auth")
	res.SetSignOnMode("BASIC_AUTH")
	res.SetLabel(label)
	return &res
}

func (t *TestFactory) NewValidBookmarkApplication(label string) *BookmarkApplication {
	app := NewBookmarkApplicationSettingsApplication()
	app.SetRequestIntegration(false)
	app.SetUrl("https://example.com/bookmark.html")
	setting := NewBookmarkApplicationSettings()
	setting.SetApp(*app)
	res := BookmarkApplication{}
	res.SetSettings(*setting)
	res.SetName("bookmark")
	res.SetSignOnMode("BOOKMARK")
	res.SetLabel(label)
	return &res
}

func (t *TestFactory) NewValidOrg2OrgApplication(label string) *SamlApplication {
	app := NewSamlApplicationSettingsApplication()
	app.SetAcsUrl("https://example.okta.com/sso/saml2/exampleid")
	app.SetAudRestriction("https://www.okta.com/saml2/service-provider/examplei")
	app.SetBaseUrl("https://example.okta.com")
	setting := NewSamlApplicationSettings()
	setting.SetApp(*app)
	res := SamlApplication{}
	res.SetSettings(*setting)
	res.SetName("okta_org2org")
	res.SetSignOnMode("SAML_2_0")
	res.SetLabel(label)
	return &res
}

func (t *TestFactory) NewValidOIDCApplication(label string) *OpenIdConnectApplication {
	settingClient := NewOpenIdConnectApplicationSettingsClient()
	settingClient.SetClientUri("https://example.com/client")
	settingClient.SetLogoUri("https://example.com/assets/images/logo-new.png")
	settingClient.SetResponseTypes([]string{"token", "id_token", "code"})
	settingClient.SetRedirectUris([]string{"https://example.com/oauth2/callback", "myapp://callback"})
	settingClient.SetPostLogoutRedirectUris([]string{"https://example.com/postlogout", "myapp://postlogoutcallback"})
	settingClient.SetGrantTypes([]string{"implicit", "authorization_code"})
	settingClient.SetApplicationType("native")
	settingClient.SetTosUri("https://example.com/client/tos")
	settingClient.SetPolicyUri("https://example.com/client/policy")
	setting := NewOpenIdConnectApplicationSettings()
	setting.SetOauthClient(*settingClient)
	credClient := NewApplicationCredentialsOAuthClient()
	credClient.SetTokenEndpointAuthMethod("client_secret_post")
	credClient.SetClientId(randomTestString())
	credClient.SetAutoKeyRotation(true)
	credentials := NewOAuthApplicationCredentials()
	credentials.SetOauthClient(*credClient)
	res := OpenIdConnectApplication{}
	res.SetSettings(*setting)
	res.SetCredentials(*credentials)
	res.SetName("oidc_client")
	res.SetSignOnMode("OPENID_CONNECT")
	res.SetLabel(label)
	return &res
}
