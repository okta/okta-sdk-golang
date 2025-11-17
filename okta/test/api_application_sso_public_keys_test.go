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

func Test_okta_ApplicationSSOPublicKeysAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationSSOPublicKeysAPIService ActivateOAuth2ClientJsonWebKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var keyId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientJsonWebKey(context.Background(), appId, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService ActivateOAuth2ClientSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var secretId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientSecret(context.Background(), appId, secretId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService AddJwk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.AddJwk(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService CreateOAuth2ClientSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.CreateOAuth2ClientSecret(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService DeactivateOAuth2ClientJsonWebKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var keyId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientJsonWebKey(context.Background(), appId, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService DeactivateOAuth2ClientSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var secretId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientSecret(context.Background(), appId, secretId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService DeleteOAuth2ClientSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var secretId string

		httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.DeleteOAuth2ClientSecret(context.Background(), appId, secretId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService Deletejwk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var keyId string

		httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.Deletejwk(context.Background(), appId, keyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService GetJwk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var keyId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.GetJwk(context.Background(), appId, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService GetOAuth2ClientSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var secretId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.GetOAuth2ClientSecret(context.Background(), appId, secretId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService ListJwk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.ListJwk(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationSSOPublicKeysAPIService ListOAuth2ClientSecrets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string

		resp, httpRes, err := apiClient.ApplicationSSOPublicKeysAPI.ListOAuth2ClientSecrets(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
