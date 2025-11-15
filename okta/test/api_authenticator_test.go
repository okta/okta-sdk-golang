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

func Test_okta_AuthenticatorAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthenticatorAPIService ActivateAuthenticator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ActivateAuthenticator(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ActivateAuthenticatorMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var methodType string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ActivateAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService CreateAuthenticator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AuthenticatorAPI.CreateAuthenticator(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService CreateCustomAAGUID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.CreateCustomAAGUID(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService DeactivateAuthenticator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.DeactivateAuthenticator(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService DeactivateAuthenticatorMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var methodType string

		resp, httpRes, err := apiClient.AuthenticatorAPI.DeactivateAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService DeleteCustomAAGUID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var aaguid string

		httpRes, err := apiClient.AuthenticatorAPI.DeleteCustomAAGUID(context.Background(), authenticatorId, aaguid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService GetAuthenticator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.GetAuthenticator(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService GetAuthenticatorMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var methodType string

		resp, httpRes, err := apiClient.AuthenticatorAPI.GetAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService GetCustomAAGUID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var aaguid string

		resp, httpRes, err := apiClient.AuthenticatorAPI.GetCustomAAGUID(context.Background(), authenticatorId, aaguid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService GetWellKnownAppAuthenticatorConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AuthenticatorAPI.GetWellKnownAppAuthenticatorConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ListAllCustomAAGUIDs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ListAllCustomAAGUIDs(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ListAuthenticatorMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ListAuthenticatorMethods(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ListAuthenticators", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AuthenticatorAPI.ListAuthenticators(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ReplaceAuthenticator", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ReplaceAuthenticator(context.Background(), authenticatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ReplaceAuthenticatorMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var methodType string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ReplaceAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService ReplaceCustomAAGUID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var aaguid string

		resp, httpRes, err := apiClient.AuthenticatorAPI.ReplaceCustomAAGUID(context.Background(), authenticatorId, aaguid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService UpdateCustomAAGUID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var aaguid string

		resp, httpRes, err := apiClient.AuthenticatorAPI.UpdateCustomAAGUID(context.Background(), authenticatorId, aaguid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticatorAPIService VerifyRpIdDomain", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authenticatorId string
		var webAuthnMethodType string

		httpRes, err := apiClient.AuthenticatorAPI.VerifyRpIdDomain(context.Background(), authenticatorId, webAuthnMethodType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
