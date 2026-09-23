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
	"slices"
	"testing"

	okta "github.com/okta/okta-sdk-golang/v6/okta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nonExistentUserRiskUserId is a well formed user ID that no org has.
const nonExistentUserRiskUserId = "00u00000000000000000"

func Test_okta_UserRiskAPIService(t *testing.T) {
	err := okta.ValidateTestEnvironment()
	if err != nil {
		t.Skip("Test environment not configured:", err)
	}

	configuration, err := okta.NewConfiguration()
	require.Nil(t, err)
	apiClient := okta.NewAPIClient(configuration)
	testDataManager := okta.GetTestDataManager()

	defer testDataManager.CleanupAllTestUsers()

	// User risk is not generally available - the spec marks it isGenerallyAvailable:
	// false - so an org without the entitlement answers 403 rather than serving the
	// endpoint. Skip in that case instead of reporting an SDK failure, but say so loudly
	// enough that it can't be mistaken for a pass. Risk is only ever read from or
	// written to a user these tests created, never an existing one.
	skipIfUserRiskUnavailable := func(t *testing.T, httpRes *okta.APIResponse) {
		t.Helper()

		if httpRes != nil && httpRes.StatusCode == http.StatusForbidden {
			t.Skip("User risk is not enabled for this org (403). It needs the Identity Threat Protection entitlement and the okta.userRisk scopes")
		}
	}

	createTestUser := func(t *testing.T) string {
		t.Helper()

		createdUser, err := testDataManager.CreateTestUser()
		require.NoError(t, err, "Failed to create test user")
		require.NotNil(t, createdUser.Id)

		return *createdUser.Id
	}

	// upsertRisk sets a risk level on a user. The spec documents both 200 and 201 for
	// this upsert - 201 when the risk object is created, 200 when an existing one is
	// updated - so either is accepted.
	upsertRisk := func(t *testing.T, userId, riskLevel, riskReason string) *okta.UserRiskPutResponse {
		t.Helper()

		riskRequest := okta.NewUserRiskRequest(riskLevel)
		riskRequest.SetRiskReason(riskReason)

		resp, httpRes, err := apiClient.UserRiskAPI.UpsertUserRisk(context.Background(), userId).UserRiskRequest(*riskRequest).Execute()
		skipIfUserRiskUnavailable(t, httpRes)

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, httpRes)
		assert.True(t, slices.Contains([]int{http.StatusOK, http.StatusCreated}, httpRes.StatusCode),
			"Upsert should return 200 or 201, got %d", httpRes.StatusCode)

		return resp
	}

	t.Run("Test UserRiskAPIService GetUserRisk", func(t *testing.T) {
		userId := createTestUser(t)

		resp, httpRes, err := apiClient.UserRiskAPI.GetUserRisk(context.Background(), userId).Execute()
		skipIfUserRiskUnavailable(t, httpRes)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		// The response is a oneOf keyed on riskLevel: NONE resolves to UserRiskLevelNone,
		// LOW/MEDIUM/HIGH to UserRiskLevelExists. A brand new user carries no risk.
		actualInstance := resp.GetActualInstance()
		require.NotNil(t, actualInstance, "Response should carry a user risk instance")

		riskLevelNone, ok := actualInstance.(*okta.UserRiskLevelNone)
		require.True(t, ok, "A newly created user should have no risk, got %T", actualInstance)
		require.NotNil(t, riskLevelNone.RiskLevel)
		assert.Equal(t, "NONE", *riskLevelNone.RiskLevel)
	})

	// An admin override is an input to the org's risk engine, not the last word on it:
	// on an org with Identity Threat Protection the engine re-evaluates the user against
	// its own signals, so a read-back can legitimately report a different level and a
	// reason it generated itself (USER_LOGOUT, say) rather than the riskReason sent here.
	// Assert the level and reason on the upsert response, which the SDK does control, and
	// on the read-back only assert what's stable: the oneOf resolves to the risk-exists
	// variant and carries one of the documented levels.
	assertRiskExists := func(t *testing.T, userId string) {
		t.Helper()

		getResp, getHttpRes, err := apiClient.UserRiskAPI.GetUserRisk(context.Background(), userId).Execute()
		require.Nil(t, err)
		require.NotNil(t, getHttpRes)
		assert.Equal(t, http.StatusOK, getHttpRes.StatusCode)

		riskLevelExists, ok := getResp.GetActualInstance().(*okta.UserRiskLevelExists)
		require.True(t, ok, "A user carrying risk should resolve to UserRiskLevelExists, got %T", getResp.GetActualInstance())
		require.NotNil(t, riskLevelExists.RiskLevel)
		assert.Contains(t, []string{"HIGH", "MEDIUM", "LOW"}, *riskLevelExists.RiskLevel)
	}

	t.Run("Test UserRiskAPIService UpsertUserRisk", func(t *testing.T) {
		userId := createTestUser(t)

		const riskReason = "Set by SDK tests"
		resp := upsertRisk(t, userId, "HIGH", riskReason)

		require.NotNil(t, resp.RiskLevel)
		assert.Equal(t, "HIGH", *resp.RiskLevel)
		if resp.Reason != nil {
			assert.Equal(t, riskReason, *resp.Reason)
		}

		// The user now has a risk object, so the GET crosses to the UserRiskLevelExists
		// side of the oneOf instead of UserRiskLevelNone
		assertRiskExists(t, userId)
	})

	t.Run("Test UserRiskAPIService UpsertUserRisk Updates An Existing Risk", func(t *testing.T) {
		userId := createTestUser(t)

		// First call creates the risk object, the second updates it
		createResp := upsertRisk(t, userId, "HIGH", "Set by SDK tests")
		require.NotNil(t, createResp.RiskLevel)
		assert.Equal(t, "HIGH", *createResp.RiskLevel)

		const lowerReason = "Lowered by SDK tests"
		updateResp := upsertRisk(t, userId, "LOW", lowerReason)
		require.NotNil(t, updateResp.RiskLevel)
		assert.Equal(t, "LOW", *updateResp.RiskLevel)
		if updateResp.Reason != nil {
			assert.Equal(t, lowerReason, *updateResp.Reason)
		}

		assertRiskExists(t, userId)
	})

	// NONE is a read-only risk level: the GET oneOf maps it to UserRiskLevelNone, but
	// UserRiskRequest.riskLevel only permits HIGH/MEDIUM/LOW and there's no DELETE on
	// the endpoint, so a risk level can be raised or lowered but never cleared. Assert
	// the rejection so a future spec change that does allow clearing shows up here.
	t.Run("Test UserRiskAPIService Error Handling - Upsert NONE Is Rejected", func(t *testing.T) {
		userId := createTestUser(t)

		upsertRisk(t, userId, "MEDIUM", "Set by SDK tests")

		riskRequest := okta.NewUserRiskRequest("NONE")

		_, httpRes, err := apiClient.UserRiskAPI.UpsertUserRisk(context.Background(), userId).UserRiskRequest(*riskRequest).Execute()
		skipIfUserRiskUnavailable(t, httpRes)

		assert.NotNil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, http.StatusBadRequest, httpRes.StatusCode)

		// The rejected call didn't clear the risk set above
		assertRiskExists(t, userId)
	})

	t.Run("Test UserRiskAPIService Error Handling - Get Risk For Nonexistent User", func(t *testing.T) {
		_, httpRes, err := apiClient.UserRiskAPI.GetUserRisk(context.Background(), nonExistentUserRiskUserId).Execute()
		skipIfUserRiskUnavailable(t, httpRes)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test UserRiskAPIService Error Handling - Upsert Risk For Nonexistent User", func(t *testing.T) {
		riskRequest := okta.NewUserRiskRequest("HIGH")

		_, httpRes, err := apiClient.UserRiskAPI.UpsertUserRisk(context.Background(), nonExistentUserRiskUserId).UserRiskRequest(*riskRequest).Execute()
		skipIfUserRiskUnavailable(t, httpRes)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test UserRiskAPIService Error Handling - Upsert Invalid Risk Level", func(t *testing.T) {
		userId := createTestUser(t)

		riskRequest := okta.NewUserRiskRequest("NOT_A_RISK_LEVEL")

		_, httpRes, err := apiClient.UserRiskAPI.UpsertUserRisk(context.Background(), userId).UserRiskRequest(*riskRequest).Execute()
		skipIfUserRiskUnavailable(t, httpRes)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, httpRes.StatusCode)
	})
}
