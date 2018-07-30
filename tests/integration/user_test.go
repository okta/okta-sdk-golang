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

	"github.com/okta/okta-sdk-golang/tests"

	"github.com/okta/okta-sdk-golang/okta"
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

	user, err := client.User.CreateUser(*u)
	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// Get the user by ID → GET /api/v1/users/{{userId}}
	ubid, err := client.User.GetUser(user.Id)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if ubid.Id != user.Id {
		t.Errorf("%v", "Error getting User By Id")
	}

	// Get the user by login name → GET /api/v1/users/{{login}}
	ubln, err := client.User.GetUser(user.Profile.Login)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if ubln.Profile.Login != user.Profile.Login {
		t.Errorf("%v", "Error getting User By Login")
	}

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	err = client.User.DeactivateUser(user.Id)
	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// Delete the user → DELETE /api/v1/users/{{userId}}
	err = client.User.DeactivateOrDeleteUser(user.Id)
	if err != nil {
		t.Errorf("%v", err.Error())
	}

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, err = client.User.GetUser(user.Id)
	if err == nil {
		t.Errorf("%v", err.Error())
	}
}
