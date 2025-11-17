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

func Test_okta_CustomTemplatesAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomTemplatesAPIService CreateEmailCustomization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.CreateEmailCustomization(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService DeleteAllCustomizations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		httpRes, err := apiClient.CustomTemplatesAPI.DeleteAllCustomizations(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService DeleteEmailCustomization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string
		var customizationId string

		httpRes, err := apiClient.CustomTemplatesAPI.DeleteEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService GetCustomizationPreview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string
		var customizationId string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.GetCustomizationPreview(context.Background(), brandId, templateName, customizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService GetEmailCustomization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string
		var customizationId string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.GetEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService GetEmailDefaultContent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.GetEmailDefaultContent(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService GetEmailDefaultPreview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.GetEmailDefaultPreview(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService GetEmailSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.GetEmailSettings(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService GetEmailTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.GetEmailTemplate(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService ListEmailCustomizations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.ListEmailCustomizations(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService ListEmailTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.ListEmailTemplates(context.Background(), brandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService ReplaceEmailCustomization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string
		var customizationId string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.ReplaceEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService ReplaceEmailSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		resp, httpRes, err := apiClient.CustomTemplatesAPI.ReplaceEmailSettings(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomTemplatesAPIService SendTestEmail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var brandId string
		var templateName string

		httpRes, err := apiClient.CustomTemplatesAPI.SendTestEmail(context.Background(), brandId, templateName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
