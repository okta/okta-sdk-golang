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
	"net/http"
	"os"
	"testing"

	okta "github.com/okta/okta-sdk-golang/v6/okta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nonExistentApiTokenId is a well formed API token ID that no org has.
const nonExistentApiTokenId = "00T00000000000000000"

func Test_okta_ApiTokenAPIService(t *testing.T) {
	err := okta.ValidateTestEnvironment()
	if err != nil {
		t.Skip("Test environment not configured:", err)
	}

	configuration, err := okta.NewConfiguration()
	require.Nil(t, err)
	apiClient := okta.NewAPIClient(configuration)

	// API tokens can't be minted through the API - they're created by an admin in the
	// Admin Console. That makes every revoke here one way: there's no way to put the
	// token back, and no way to tell from the list which entry is the token this suite
	// is authenticating with. The read operations run against the org's real tokens;
	// anything that revokes or rewrites one is skipped or opt in.
	requireExistingApiToken := func(t *testing.T) okta.ApiToken {
		t.Helper()

		tokens, httpRes, err := apiClient.ApiTokenAPI.ListApiTokens(context.Background()).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)

		if len(tokens) == 0 {
			t.Skip("Org has no API tokens; tokens can't be created through the API")
		}
		require.NotNil(t, tokens[0].Id)

		return tokens[0]
	}

	// disposableApiTokenId returns the token that opted in to being mutated, if any.
	disposableApiTokenId := func(t *testing.T) string {
		t.Helper()

		apiTokenId := os.Getenv("OKTA_TEST_API_TOKEN_ID")
		if apiTokenId == "" {
			t.Skip("Set OKTA_TEST_API_TOKEN_ID to a disposable API token - NOT the one in OKTA_CLIENT_TOKEN - to run this test")
		}

		return apiTokenId
	}

	t.Run("Test ApiTokenAPIService ListApiTokens", func(t *testing.T) {
		resp, httpRes, err := apiClient.ApiTokenAPI.ListApiTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		assert.IsType(t, []okta.ApiToken{}, resp)

		// The token this suite authenticates with is itself in the list
		assert.Greater(t, len(resp), 0, "The org should have at least the token used by these tests")

		for _, token := range resp {
			assert.NotNil(t, token.Id)
			assert.NotEmpty(t, token.Name)
			assert.NotNil(t, token.UserId)
			assert.NotNil(t, token.Created)
		}
	})

	t.Run("Test ApiTokenAPIService GetApiToken", func(t *testing.T) {
		token := requireExistingApiToken(t)

		resp, httpRes, err := apiClient.ApiTokenAPI.GetApiToken(context.Background(), *token.Id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		require.NotNil(t, resp.Id)
		assert.Equal(t, *token.Id, *resp.Id)
		assert.Equal(t, token.Name, resp.Name)
		assert.NotNil(t, resp.UserId)
		assert.NotNil(t, resp.Created)
	})

	t.Run("Test ApiTokenAPIService UpsertApiToken", func(t *testing.T) {
		apiTokenId := disposableApiTokenId(t)

		// Deliberately the permissive condition: a restrictive zone on the wrong token
		// would lock the caller out of the org's API.
		network := okta.NewApiTokenNetwork()
		network.SetConnection("ANYWHERE")

		tokenUpdate := okta.NewApiTokenUpdate()
		tokenUpdate.SetNetwork(*network)

		resp, httpRes, err := apiClient.ApiTokenAPI.UpsertApiToken(context.Background(), apiTokenId).ApiTokenUpdate(*tokenUpdate).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		require.NotNil(t, resp.Id)
		assert.Equal(t, apiTokenId, *resp.Id)
		if resp.Network != nil && resp.Network.Connection != nil {
			assert.Equal(t, "ANYWHERE", *resp.Network.Connection)
		}
	})

	t.Run("Test ApiTokenAPIService RevokeApiToken", func(t *testing.T) {
		// Revoking is irreversible and a replacement token can only be minted by an
		// admin in the Admin Console, so this can't be made repeatable. There's also no
		// way to tell which listed token is the one in OKTA_CLIENT_TOKEN, so picking one
		// automatically risks cutting off the whole suite. The 404 path is covered below.
		t.Skip("Revoking an API token is irreversible and tokens can't be recreated through the API")
	})

	t.Run("Test ApiTokenAPIService RevokeCurrentApiToken", func(t *testing.T) {
		// This revokes the token in OKTA_CLIENT_TOKEN. Every later test in this run, and
		// every run after it, would fail with 401 until an admin mints a new token by
		// hand. Never run this against an org you care about - there is no API level
		// recovery path.
		t.Skip("Revoking the current API token would invalidate OKTA_CLIENT_TOKEN for this and all later runs")
	})

	t.Run("Test ApiTokenAPIService Error Handling - Get Nonexistent Api Token", func(t *testing.T) {
		_, httpRes, err := apiClient.ApiTokenAPI.GetApiToken(context.Background(), nonExistentApiTokenId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test ApiTokenAPIService Error Handling - Upsert Nonexistent Api Token", func(t *testing.T) {
		network := okta.NewApiTokenNetwork()
		network.SetConnection("ANYWHERE")

		tokenUpdate := okta.NewApiTokenUpdate()
		tokenUpdate.SetNetwork(*network)

		_, httpRes, err := apiClient.ApiTokenAPI.UpsertApiToken(context.Background(), nonExistentApiTokenId).ApiTokenUpdate(*tokenUpdate).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test ApiTokenAPIService Error Handling - Revoke Nonexistent Api Token", func(t *testing.T) {
		// Safe to exercise for real: an all zeros ID matches no token, so nothing is
		// revoked. This covers the revoke code path without the destruction.
		httpRes, err := apiClient.ApiTokenAPI.RevokeApiToken(context.Background(), nonExistentApiTokenId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})
}
