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
	"time"

	"github.com/okta/okta-sdk-golang/okta.v2"
	"github.com/okta/okta-sdk-golang/tests.v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_can_create_user_type(t *testing.T) {
	client, _ := tests.NewClient()

	ut := okta.UserType{
		Description: "My Custom User Type",
		DisplayName: "Display Name",
		Name:        "userTypeName",
	}

	userType, response, err := client.UserType.CreateUserType(ut)
	require.NoError(t, err, "creating a user type should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/meta/types/user")

	assert_user_type_model(t, userType)

	_, err = client.UserType.DeleteUserType(userType.Id)
	require.NoError(t, err, "deleting a user type should not error")

}

func Test_can_list_user_types(t *testing.T) {
	client, _ := tests.NewClient()

	userTypes, response, err := client.UserType.ListUserTypes()
	require.NoError(t, err, "creating a user type should not error")
	tests.Assert_response(t, response, "GET", "/api/v1/meta/types/user")

	assert_user_type_model(t, userTypes[0])

}

func assert_user_type_model(t *testing.T, userType *okta.UserType) {
	assert.NotEmpty(t, userType.Id, "id should not be empty")
	assert.NotEmpty(t, userType.Created, "created should not be empty")
	assert.IsType(t, &time.Time{}, userType.Created, "created should not be of type `*time.Time`")
	assert.NotEmpty(t, userType.CreatedBy, "createdBy should not be empty")
	assert.NotEmpty(t, userType.Description, "description should not be empty")
	assert.NotEmpty(t, userType.DisplayName, "displayName should not be empty")
	assert.NotEmpty(t, userType.LastUpdated, "lastUpdated should not be empty")
	assert.IsType(t, &time.Time{}, userType.LastUpdated, "lastUpdated should not be of type `*time.Time`")
	assert.NotEmpty(t, userType.LastUpdatedBy, "lastUpdatedBy should not be empty")
	assert.NotEmpty(t, userType.Name, "name should not be empty")

}
