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

func Test_okta_LogStreamAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogStreamAPIService ActivateLogStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logStreamId string

		resp, httpRes, err := apiClient.LogStreamAPI.ActivateLogStream(context.Background(), logStreamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogStreamAPIService CreateLogStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LogStreamAPI.CreateLogStream(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogStreamAPIService DeactivateLogStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logStreamId string

		resp, httpRes, err := apiClient.LogStreamAPI.DeactivateLogStream(context.Background(), logStreamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogStreamAPIService DeleteLogStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logStreamId string

		httpRes, err := apiClient.LogStreamAPI.DeleteLogStream(context.Background(), logStreamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogStreamAPIService GetLogStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logStreamId string

		resp, httpRes, err := apiClient.LogStreamAPI.GetLogStream(context.Background(), logStreamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogStreamAPIService ListLogStreams", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LogStreamAPI.ListLogStreams(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogStreamAPIService ReplaceLogStream", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logStreamId string

		resp, httpRes, err := apiClient.LogStreamAPI.ReplaceLogStream(context.Background(), logStreamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
