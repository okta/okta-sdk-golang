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

func Test_okta_GroupAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupAPIService AddGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupAPI.AddGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService AssignUserToGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var userId string

		httpRes, err := apiClient.GroupAPI.AssignUserToGroup(context.Background(), groupId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService DeleteGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		httpRes, err := apiClient.GroupAPI.DeleteGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService GetGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.GroupAPI.GetGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService ListAssignedApplicationsForGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.GroupAPI.ListAssignedApplicationsForGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService ListGroupUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.GroupAPI.ListGroupUsers(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService ListGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupAPI.ListGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService ReplaceGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.GroupAPI.ReplaceGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupAPIService UnassignUserFromGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string
		var userId string

		httpRes, err := apiClient.GroupAPI.UnassignUserFromGroup(context.Background(), groupId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
