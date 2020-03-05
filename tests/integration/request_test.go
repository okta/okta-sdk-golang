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
	"io"
	"testing"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/okta/query"
	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_private_key_request_contains_bearer_token(t *testing.T) {
	var buff io.ReadWriter

	client, _ := tests.NewClient(okta.WithAuthorizationMode("PrivateKey"), okta.WithScopes(([]string{"okta.users.manage"})))

	request, _ := client.GetRequestExecutor().NewRequest("GET", "https://example.com/", buff)

	assert.Contains(t, request.Header.Get("Authorization"), "Bearer", "does not contain a bearer token for the request")

}

func Test_private_key_request_can_create_a_user(t *testing.T) {
	client, _ := tests.NewClient(okta.WithAuthorizationMode("PrivateKey"), okta.WithScopes(([]string{"okta.users.manage"})))

	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Private_Key"
	profile["email"] = "john-private-key@example.com"
	profile["login"] = "john-private-key@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}

	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")
	assert.NotEmpty(t, user.Id, "appears the user was not created")
	tempProfile := *user.Profile
	assert.Equal(t, "john-private-key@example.com", tempProfile["email"], "did not get the correct user")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}
