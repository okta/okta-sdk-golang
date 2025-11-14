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

func Test_okta_GovernanceBundleAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GovernanceBundleAPIService CreateGovernanceBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GovernanceBundleAPI.CreateGovernanceBundle(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService DeleteGovernanceBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId string

		httpRes, err := apiClient.GovernanceBundleAPI.DeleteGovernanceBundle(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService GetGovernanceBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId string

		resp, httpRes, err := apiClient.GovernanceBundleAPI.GetGovernanceBundle(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService GetOptInStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GovernanceBundleAPI.GetOptInStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService ListBundleEntitlementValues", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId string
		var entitlementId string

		resp, httpRes, err := apiClient.GovernanceBundleAPI.ListBundleEntitlementValues(context.Background(), bundleId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService ListBundleEntitlements", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId string

		resp, httpRes, err := apiClient.GovernanceBundleAPI.ListBundleEntitlements(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService ListGovernanceBundles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GovernanceBundleAPI.ListGovernanceBundles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService OptIn", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GovernanceBundleAPI.OptIn(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService OptOut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GovernanceBundleAPI.OptOut(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceBundleAPIService ReplaceGovernanceBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId string

		resp, httpRes, err := apiClient.GovernanceBundleAPI.ReplaceGovernanceBundle(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
