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
	"testing"

	"github.com/okta/okta-sdk-golang/okta/query"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_can_get_a_user(t *testing.T) {
	client := tests.NewClient()
	// Create user with credentials → POST /api/v1/users?activate=false
	p := new(okta.PasswordCredential).WithValue("Abcd1234")
	uc := new(okta.UserCredentials).WithPassword(p)
	profile := new(okta.UserProfile).
		WithFirstName("John").
		WithLastName("Get-User").
		WithEmail("john-get-user@example.com").
		WithLogin("john-get-user@example.com")
	u := new(okta.User).WithCredentials(uc).WithProfile(profile)
	qp := query.NewQueryParams(query.WithSendEmail(false))

	user, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Get the user by ID → GET /api/v1/users/{{userId}}
	ubid, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "Getting a user by id should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Id")

	// Get the user by login name → GET /api/v1/users/{{login}}
	ubln, err := client.User.GetUser(user.Profile.Login, nil)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubln.Id, "Could not find user by Login")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}
