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

func Test_okta_IdentityProviderSigningKeysAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IdentityProviderSigningKeysAPIService CloneIdentityProviderKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string
		var kid string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.CloneIdentityProviderKey(context.Background(), idpId, kid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService GenerateCsrForIdentityProvider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.GenerateCsrForIdentityProvider(context.Background(), idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService GenerateIdentityProviderSigningKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.GenerateIdentityProviderSigningKey(context.Background(), idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService GetCsrForIdentityProvider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string
		var idpCsrId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.GetCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService GetIdentityProviderSigningKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string
		var kid string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.GetIdentityProviderSigningKey(context.Background(), idpId, kid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService ListActiveIdentityProviderSigningKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.ListActiveIdentityProviderSigningKey(context.Background(), idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService ListCsrsForIdentityProvider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.ListCsrsForIdentityProvider(context.Background(), idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService ListIdentityProviderSigningKeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.ListIdentityProviderSigningKeys(context.Background(), idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService PublishCsrForIdentityProvider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string
		var idpCsrId string

		resp, httpRes, err := apiClient.IdentityProviderSigningKeysAPI.PublishCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProviderSigningKeysAPIService RevokeCsrForIdentityProvider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idpId string
		var idpCsrId string

		httpRes, err := apiClient.IdentityProviderSigningKeysAPI.RevokeCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
