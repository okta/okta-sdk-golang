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

func Test_okta_ThemesAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ThemesAPIService DeleteBrandThemeBackgroundImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		httpRes, err := apiClient.ThemesAPI.DeleteBrandThemeBackgroundImage(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService DeleteBrandThemeFavicon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		httpRes, err := apiClient.ThemesAPI.DeleteBrandThemeFavicon(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService DeleteBrandThemeLogo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		httpRes, err := apiClient.ThemesAPI.DeleteBrandThemeLogo(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService GetBrandTheme", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		resp, httpRes, err := apiClient.ThemesAPI.GetBrandTheme(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService ListBrandThemes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.ThemesAPI.ListBrandThemes(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService ReplaceBrandTheme", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		resp, httpRes, err := apiClient.ThemesAPI.ReplaceBrandTheme(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService UploadBrandThemeBackgroundImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		resp, httpRes, err := apiClient.ThemesAPI.UploadBrandThemeBackgroundImage(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService UploadBrandThemeFavicon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		resp, httpRes, err := apiClient.ThemesAPI.UploadBrandThemeFavicon(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemesAPIService UploadBrandThemeLogo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var themeId string

		resp, httpRes, err := apiClient.ThemesAPI.UploadBrandThemeLogo(context.Background(), brandId, themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
