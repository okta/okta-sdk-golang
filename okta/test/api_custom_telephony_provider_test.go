/*
Okta Admin Management API

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

func Test_okta_CustomTelephonyProviderAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomTelephonyProviderAPIService ActivateCustomTelephonyCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.ActivateCustomTelephonyCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService CreateCustomTelephonyProviderCredentials", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.CreateCustomTelephonyProviderCredentials(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService DeactivateCustomTelephonyCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.DeactivateCustomTelephonyCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService DeleteCustomTelephonyProviderCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		httpRes, err := apiClient.CustomTelephonyProviderAPI.DeleteCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService GetCustomTelephonyProviderCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.GetCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService ListAllCustomTelephonyProviderCredentials", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.ListAllCustomTelephonyProviderCredentials(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService SendTestCustomTelephonyProviderCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		httpRes, err := apiClient.CustomTelephonyProviderAPI.SendTestCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService SetAsPrimaryCustomTelephonyCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.SetAsPrimaryCustomTelephonyCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTelephonyProviderAPIService UpdateCustomTelephonyProviderCredential", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customTelephonyProviderId string

		resp, httpRes, err := apiClient.CustomTelephonyProviderAPI.UpdateCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
