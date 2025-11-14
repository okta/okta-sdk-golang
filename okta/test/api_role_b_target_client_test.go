/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"context"
	"testing"

	openapiclient "github.com/okta/okta-sdk-golang/v6/okta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_okta_RoleBTargetClientAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleBTargetClientAPIService AssignAppTargetInstanceRoleForClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string
		var appName string
		var appId string

		httpRes, err := apiClient.RoleBTargetClientAPI.AssignAppTargetInstanceRoleForClient(context.Background(), clientId, roleAssignmentId, appName, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService AssignAppTargetRoleToClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string
		var appName string

		httpRes, err := apiClient.RoleBTargetClientAPI.AssignAppTargetRoleToClient(context.Background(), clientId, roleAssignmentId, appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService AssignGroupTargetRoleForClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string
		var groupId string

		httpRes, err := apiClient.RoleBTargetClientAPI.AssignGroupTargetRoleForClient(context.Background(), clientId, roleAssignmentId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService ListAppTargetRoleToClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string

		resp, httpRes, err := apiClient.RoleBTargetClientAPI.ListAppTargetRoleToClient(context.Background(), clientId, roleAssignmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService ListGroupTargetRoleForClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string

		resp, httpRes, err := apiClient.RoleBTargetClientAPI.ListGroupTargetRoleForClient(context.Background(), clientId, roleAssignmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService RemoveAppTargetInstanceRoleForClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string
		var appName string
		var appId string

		httpRes, err := apiClient.RoleBTargetClientAPI.RemoveAppTargetInstanceRoleForClient(context.Background(), clientId, roleAssignmentId, appName, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService RemoveAppTargetRoleFromClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string
		var appName string

		httpRes, err := apiClient.RoleBTargetClientAPI.RemoveAppTargetRoleFromClient(context.Background(), clientId, roleAssignmentId, appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetClientAPIService RemoveGroupTargetRoleFromClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var roleAssignmentId string
		var groupId string

		httpRes, err := apiClient.RoleBTargetClientAPI.RemoveGroupTargetRoleFromClient(context.Background(), clientId, roleAssignmentId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
