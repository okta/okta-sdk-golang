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

func Test_okta_PolicyAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PolicyAPIService ActivatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		httpRes, err := apiClient.PolicyAPI.ActivatePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ActivatePolicyRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var ruleId string

		httpRes, err := apiClient.PolicyAPI.ActivatePolicyRule(context.Background(), policyId, ruleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ClonePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.ClonePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService CreatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PolicyAPI.CreatePolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService CreatePolicyRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.CreatePolicyRule(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService CreatePolicySimulation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PolicyAPI.CreatePolicySimulation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService DeactivatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		httpRes, err := apiClient.PolicyAPI.DeactivatePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService DeactivatePolicyRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var ruleId string

		httpRes, err := apiClient.PolicyAPI.DeactivatePolicyRule(context.Background(), policyId, ruleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService DeletePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		httpRes, err := apiClient.PolicyAPI.DeletePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService DeletePolicyResourceMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var mappingId string

		httpRes, err := apiClient.PolicyAPI.DeletePolicyResourceMapping(context.Background(), policyId, mappingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService DeletePolicyRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var ruleId string

		httpRes, err := apiClient.PolicyAPI.DeletePolicyRule(context.Background(), policyId, ruleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService GetPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.GetPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService GetPolicyMapping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var mappingId string

		resp, httpRes, err := apiClient.PolicyAPI.GetPolicyMapping(context.Background(), policyId, mappingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService GetPolicyRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var ruleId string

		resp, httpRes, err := apiClient.PolicyAPI.GetPolicyRule(context.Background(), policyId, ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ListPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PolicyAPI.ListPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ListPolicyApps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.ListPolicyApps(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ListPolicyMappings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.ListPolicyMappings(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ListPolicyRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.ListPolicyRules(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService MapResourceToPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.MapResourceToPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ReplacePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PolicyAPI.ReplacePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyAPIService ReplacePolicyRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId string
		var ruleId string

		resp, httpRes, err := apiClient.PolicyAPI.ReplacePolicyRule(context.Background(), policyId, ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
