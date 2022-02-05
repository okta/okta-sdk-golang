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

package unit

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/v2/okta"
)

func TestSocialAuthTokenMarshaling(t *testing.T) {
	type errorTestCases struct {
		name              string
		_json             []byte
		expectedScopes    []string
		expectedExpiredAt string
	}

	for _, scenario := range []errorTestCases{
		{
			name: "social auth base case",
			_json: []byte(`
				{
					"id": "<unique token identifier>",
					"token": "JBTWGV22G4ZGKV3N",
					"tokenType" : "urn:ietf:params:oauth:token-type:access_token",
					"tokenAuthScheme": "Bearer",
					"expiresAt" : "2014-08-06T16:56:31.000Z",
					"scopes" : [
						"email",
						"openid",
						"profile"
					]
			  	}
			`),
			expectedScopes:    []string{"email", "openid", "profile"},
			expectedExpiredAt: "2014-08-06T16:56:31.000Z",
		},
		{
			name: "social auth token with blank exiresAt",
			_json: []byte(`
				{
					"id": "<unique token identifier>",
					"token": "JBTWGV22G4ZGKV3N",
					"tokenType" : "urn:ietf:params:oauth:token-type:access_token",
					"tokenAuthScheme": null,
					"expiresAt" : "",
					"scopes" : []
			  	}
			`),
			expectedScopes: []string{},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			var token okta.SocialAuthToken
			err := json.Unmarshal(scenario._json, &token)
			require.NoError(t, err)

			if scenario.expectedExpiredAt == "" {
				require.Nil(t, token.ExpiresAt)
			} else {
				expectedExpiredAt, err := time.Parse(time.RFC3339, scenario.expectedExpiredAt)
				require.NoError(t, err)
				require.Equal(t, &expectedExpiredAt, token.ExpiresAt)
			}

			require.Equal(t, scenario.expectedScopes, token.Scopes)
		})
	}
}
