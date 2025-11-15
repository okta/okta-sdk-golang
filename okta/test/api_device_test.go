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

func Test_okta_DeviceAPIService(t *testing.T) {

	configuration, err := openapiclient.NewConfiguration()
	require.Nil(t, err)
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceAPIService ActivateDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceAPI.ActivateDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService DeactivateDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceAPI.DeactivateDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService DeleteDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceAPI.DeleteDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService GetDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceAPI.GetDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService ListDeviceUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceAPI.ListDeviceUsers(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService ListDevices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeviceAPI.ListDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService SuspendDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceAPI.SuspendDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService UnsuspendDevice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceAPI.UnsuspendDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
