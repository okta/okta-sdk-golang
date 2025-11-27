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

func Test_okta_RoleBTargetBGroupAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleBTargetBGroupAPIService AssignAppInstanceTargetToAppAdminRoleForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string
		var appName string
		var appId string

		httpRes, err := apiClient.RoleBTargetBGroupAPI.AssignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService AssignAppTargetToAdminRoleForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string
		var appName string

		httpRes, err := apiClient.RoleBTargetBGroupAPI.AssignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService AssignGroupTargetToGroupAdminRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string
		var targetGroupId string

		httpRes, err := apiClient.RoleBTargetBGroupAPI.AssignGroupTargetToGroupAdminRole(context.Background(), groupId, roleAssignmentId, targetGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService ListApplicationTargetsForApplicationAdministratorRoleForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string

		resp, httpRes, err := apiClient.RoleBTargetBGroupAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup(context.Background(), groupId, roleAssignmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService ListGroupTargetsForGroupRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string

		resp, httpRes, err := apiClient.RoleBTargetBGroupAPI.ListGroupTargetsForGroupRole(context.Background(), groupId, roleAssignmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService UnassignAppInstanceTargetToAppAdminRoleForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string
		var appName string
		var appId string

		httpRes, err := apiClient.RoleBTargetBGroupAPI.UnassignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService UnassignAppTargetToAdminRoleForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string
		var appName string

		httpRes, err := apiClient.RoleBTargetBGroupAPI.UnassignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleBTargetBGroupAPIService UnassignGroupTargetFromGroupAdminRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var roleAssignmentId string
		var targetGroupId string

		httpRes, err := apiClient.RoleBTargetBGroupAPI.UnassignGroupTargetFromGroupAdminRole(context.Background(), groupId, roleAssignmentId, targetGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
