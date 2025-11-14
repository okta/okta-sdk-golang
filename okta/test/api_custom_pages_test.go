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

func Test_okta_CustomPagesAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomPagesAPIService DeleteCustomizedErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		httpRes, err := apiClient.CustomPagesAPI.DeleteCustomizedErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService DeleteCustomizedSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		httpRes, err := apiClient.CustomPagesAPI.DeleteCustomizedSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService DeletePreviewErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		httpRes, err := apiClient.CustomPagesAPI.DeletePreviewErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService DeletePreviewSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		httpRes, err := apiClient.CustomPagesAPI.DeletePreviewSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetCustomizedErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetCustomizedErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetCustomizedSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetCustomizedSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetDefaultErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetDefaultErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetDefaultSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetDefaultSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetPreviewErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetPreviewErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetPreviewSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetPreviewSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService GetSignOutPageSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.GetSignOutPageSettings(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService ListAllSignInWidgetVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.ListAllSignInWidgetVersions(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService ReplaceCustomizedErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.ReplaceCustomizedErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService ReplaceCustomizedSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.ReplaceCustomizedSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService ReplacePreviewErrorPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.ReplacePreviewErrorPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService ReplacePreviewSignInPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.ReplacePreviewSignInPage(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomPagesAPIService ReplaceSignOutPageSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomPagesAPI.ReplaceSignOutPageSettings(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
