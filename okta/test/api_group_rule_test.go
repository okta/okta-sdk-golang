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

func Test_okta_GroupRuleAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupRuleAPIService ActivateGroupRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupRuleId string

		httpRes, err := apiClient.GroupRuleAPI.ActivateGroupRule(context.Background(), groupRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupRuleAPIService CreateGroupRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupRuleAPI.CreateGroupRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupRuleAPIService DeactivateGroupRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupRuleId string

		httpRes, err := apiClient.GroupRuleAPI.DeactivateGroupRule(context.Background(), groupRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupRuleAPIService DeleteGroupRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupRuleId string

		httpRes, err := apiClient.GroupRuleAPI.DeleteGroupRule(context.Background(), groupRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupRuleAPIService GetGroupRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupRuleId string

		resp, httpRes, err := apiClient.GroupRuleAPI.GetGroupRule(context.Background(), groupRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupRuleAPIService ListGroupRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupRuleAPI.ListGroupRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupRuleAPIService ReplaceGroupRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupRuleId string

		resp, httpRes, err := apiClient.GroupRuleAPI.ReplaceGroupRule(context.Background(), groupRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
