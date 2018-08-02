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
	"fmt"
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
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Get-User"
	profile["email"] = "john-get-user@example.com"
	profile["login"] = "john-get-user@example.com"
	u := new(okta.User).WithCredentials(uc).WithProfile(&profile)
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Get the user by ID → GET /api/v1/users/{{userId}}
	ubid, _, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "Getting a user by id should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Id")

	// Get the user by login name → GET /api/v1/users/{{login}}
	ubln, _, err := client.User.GetUser(profile["login"].(string), nil)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubln.Id, "Could not find user by Login")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_activate_a_user(t *testing.T) {
	client := tests.NewClient()
	//Create user with credentials → POST /api/v1/users?activate=false
	p := new(okta.PasswordCredential).WithValue("Abcd1234")
	uc := new(okta.UserCredentials).WithPassword(p)
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Activate"
	profile["email"] = "john-activate@example.com"
	profile["login"] = "john-activate@example.com"
	u := new(okta.User).WithCredentials(uc).WithProfile(&profile)
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Activate the user → POST /api/v1/users/{{userId}}/lifecycle/activate?sendEmail=false
	token, _, err := client.User.ActivateUser(user.Id, query.NewQueryParams(query.WithSendEmail(false)))
	require.NoError(t, err, "Could not activate the user")
	assert.NotEmpty(t, token, "Token was not provided")
	assert.IsType(t, &okta.UserActivationToken{}, token, "Activation did not return correct type")

	// Verify that the user is in the list of ACTIVE users with query parameter → GET /api/v1/users?filter=status eq "ACTIVE"
	filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))
	users, _, err := client.User.ListUsers(filter)
	require.NoError(t, err, "Could not get active users")
	found := false
	for _, u := range users {
		fmt.Printf("%+v\n", u)
		if user.Id == u.Id {
			found = true
		}
	}
	assert.True(t, found, "The user was not found")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_update_user_profile(t *testing.T) {
	client := tests.NewClient()
	// Create user with credentials → POST /api/v1/users?activate=false
	p := new(okta.PasswordCredential).WithValue("Abcd1234")
	uc := new(okta.UserCredentials).WithPassword(p)
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Profile-Update"
	profile["email"] = "john-profile-update@example.com"
	profile["login"] = "john-profile-update@example.com"
	u := new(okta.User).WithCredentials(uc).WithProfile(&profile)
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Update the user's profile by adding a nickname → PUT /api/v1/users/{{userId}}
	newProfile := *user.Profile
	newProfile["nickName"] = "Batman"
	updatedUser := new(okta.User).WithProfile(&newProfile)
	_, _, err = client.User.UpdateUser(user.Id, *updatedUser, nil)
	require.NoError(t, err, "Could not update the user")

	// Verify that user profile is updated by calling get on the user → GET /api/v1/users/{{userId}}
	tmpUser, _, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "User was not available to get")
	tmpProfile := *tmpUser.Profile
	assert.Equal(t, "Batman", tmpProfile["nickName"])
	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}
