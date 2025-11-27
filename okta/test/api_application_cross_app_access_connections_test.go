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

func Test_okta_ApplicationCrossAppAccessConnectionsAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationCrossAppAccessConnectionsAPIService CreateCrossAppAccessConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string

		resp, httpRes, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.CreateCrossAppAccessConnection(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCrossAppAccessConnectionsAPIService DeleteCrossAppAccessConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var connectionId string

		httpRes, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.DeleteCrossAppAccessConnection(context.Background(), appId, connectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCrossAppAccessConnectionsAPIService GetAllCrossAppAccessConnections", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string

		resp, httpRes, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.GetAllCrossAppAccessConnections(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCrossAppAccessConnectionsAPIService GetCrossAppAccessConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var connectionId string

		resp, httpRes, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.GetCrossAppAccessConnection(context.Background(), appId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCrossAppAccessConnectionsAPIService UpdateCrossAppAccessConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId string
		var connectionId string

		resp, httpRes, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.UpdateCrossAppAccessConnection(context.Background(), appId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
