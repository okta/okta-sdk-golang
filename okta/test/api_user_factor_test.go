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

func Test_okta_UserFactorAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserFactorAPIService ActivateFactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var factorId string

		resp, httpRes, err := apiClient.UserFactorAPI.ActivateFactor(context.Background(), userId, factorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService EnrollFactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserFactorAPI.EnrollFactor(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService GetFactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var factorId string

		resp, httpRes, err := apiClient.UserFactorAPI.GetFactor(context.Background(), userId, factorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService GetFactorTransactionStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var factorId string
		var transactionId string

		resp, httpRes, err := apiClient.UserFactorAPI.GetFactorTransactionStatus(context.Background(), userId, factorId, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService GetYubikeyOtpTokenById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tokenId string

		resp, httpRes, err := apiClient.UserFactorAPI.GetYubikeyOtpTokenById(context.Background(), tokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService ListFactors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserFactorAPI.ListFactors(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService ListSupportedFactors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserFactorAPI.ListSupportedFactors(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService ListSupportedSecurityQuestions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserFactorAPI.ListSupportedSecurityQuestions(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService ListYubikeyOtpTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserFactorAPI.ListYubikeyOtpTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService ResendEnrollFactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var factorId string

		resp, httpRes, err := apiClient.UserFactorAPI.ResendEnrollFactor(context.Background(), userId, factorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService UnenrollFactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var factorId string

		httpRes, err := apiClient.UserFactorAPI.UnenrollFactor(context.Background(), userId, factorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService UploadYubikeyOtpTokenSeed", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserFactorAPI.UploadYubikeyOtpTokenSeed(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFactorAPIService VerifyFactor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var factorId string

		resp, httpRes, err := apiClient.UserFactorAPI.VerifyFactor(context.Background(), userId, factorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
