/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateIdentityProvider(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	idpName := testName("Test Identity Provider")
	testIdp, err := createIdentityProvider(idpName)
	require.NoError(t, err)

	resultIpd, response, err := client.IdentityProvider.CreateIdentityProvider(ctx, *testIdp)
	go cleanupTestIdentityProvider(ctx, client, resultIpd)

	require.NoError(t, err, "creating an identity provider hook should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/idps")
	assert.Equal(t, idpName, resultIpd.Name)

}

func cleanupTestIdentityProvider(ctx context.Context, client *okta.Client, idp *okta.IdentityProvider) {
	if idp == nil || idp.Name == "" {
		return
	}
	_, _ = client.IdentityProvider.DeleteIdentityProvider(ctx, idp.Id)
}

func createIdentityProvider(name string) (*okta.IdentityProvider, error) {
	jsonIDP := `
		{
			"type": "OIDC",
			"name": "` + name + `",
			"protocol": {
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
			},
			"policy": {
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
			}
		  }
	`

	var idp okta.IdentityProvider

	err := json.Unmarshal([]byte(jsonIDP), &idp)
	if err != nil {
		return nil, err
	}

	return &idp, nil
}
