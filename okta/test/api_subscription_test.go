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

func Test_okta_SubscriptionAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionAPIService GetSubscriptionsNotificationTypeRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleRef openapiclient.ListSubscriptionsRoleRoleRefParameter
		var notificationType string

		resp, httpRes, err := apiClient.SubscriptionAPI.GetSubscriptionsNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService GetSubscriptionsNotificationTypeUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationType string
		var userId string

		resp, httpRes, err := apiClient.SubscriptionAPI.GetSubscriptionsNotificationTypeUser(context.Background(), notificationType, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService ListSubscriptionsRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleRef openapiclient.ListSubscriptionsRoleRoleRefParameter

		resp, httpRes, err := apiClient.SubscriptionAPI.ListSubscriptionsRole(context.Background(), roleRef).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService ListSubscriptionsUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.SubscriptionAPI.ListSubscriptionsUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService SubscribeByNotificationTypeRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleRef openapiclient.ListSubscriptionsRoleRoleRefParameter
		var notificationType string

		httpRes, err := apiClient.SubscriptionAPI.SubscribeByNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService SubscribeByNotificationTypeUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationType string
		var userId string

		httpRes, err := apiClient.SubscriptionAPI.SubscribeByNotificationTypeUser(context.Background(), notificationType, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService UnsubscribeByNotificationTypeRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleRef openapiclient.ListSubscriptionsRoleRoleRefParameter
		var notificationType string

		httpRes, err := apiClient.SubscriptionAPI.UnsubscribeByNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAPIService UnsubscribeByNotificationTypeUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationType string
		var userId string

		httpRes, err := apiClient.SubscriptionAPI.UnsubscribeByNotificationTypeUser(context.Background(), notificationType, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
