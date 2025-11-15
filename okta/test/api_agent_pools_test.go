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

func Test_okta_AgentPoolsAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AgentPoolsAPIService ActivateAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.ActivateAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService CreateAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.CreateAgentPoolsUpdate(context.Background(), poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService DeactivateAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.DeactivateAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService DeleteAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		httpRes, err := apiClient.AgentPoolsAPI.DeleteAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService GetAgentPoolsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.GetAgentPoolsUpdateInstance(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService GetAgentPoolsUpdateSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.GetAgentPoolsUpdateSettings(context.Background(), poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService ListAgentPools", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AgentPoolsAPI.ListAgentPools(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService ListAgentPoolsUpdates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.ListAgentPoolsUpdates(context.Background(), poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService PauseAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.PauseAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService ResumeAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.ResumeAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService RetryAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.RetryAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService StopAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.StopAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService UpdateAgentPoolsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string
		var updateId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.UpdateAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentPoolsAPIService UpdateAgentPoolsUpdateSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var poolId string

		resp, httpRes, err := apiClient.AgentPoolsAPI.UpdateAgentPoolsUpdateSettings(context.Background(), poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
