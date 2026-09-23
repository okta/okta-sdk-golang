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
	"encoding/json"
	"net/http"
	"os"
	"testing"

	okta "github.com/okta/okta-sdk-golang/v6/okta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nonExistentDeviceId is a well formed device ID that is not enrolled in any org.
const nonExistentDeviceId = "guo00000000000000000"

func Test_okta_DeviceAPIService(t *testing.T) {
	err := okta.ValidateTestEnvironment()
	if err != nil {
		t.Skip("Test environment not configured:", err)
	}

	configuration, err := okta.NewConfiguration()
	require.Nil(t, err)
	apiClient := okta.NewAPIClient(configuration)

	// Devices cannot be created through the management API - they're enrolled by a
	// real device going through Okta Verify. The read operations below therefore run
	// against whatever devices the org already has, and the lifecycle operations are
	// opt in (see OKTA_TEST_DEVICE_ID) so that a test run never mutates someone's
	// device by accident.
	requireExistingDevice := func(t *testing.T) okta.DeviceList {
		t.Helper()

		devices, httpRes, err := apiClient.DeviceAPI.ListDevices(context.Background()).Limit(1).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)

		if len(devices) == 0 {
			t.Skip("Org has no enrolled devices; devices can't be created through the API")
		}
		require.NotNil(t, devices[0].Id)

		if payload, marshalErr := json.MarshalIndent(devices[0], "", "  "); marshalErr == nil {
			t.Logf("Using device %s:\n%s", *devices[0].Id, payload)
		} else {
			t.Logf("Using device %s: %+v", *devices[0].Id, devices[0])
		}

		return devices[0]
	}

	// lifecycleDeviceId returns the device that opted in to lifecycle testing, if any.
	lifecycleDeviceId := func(t *testing.T) string {
		t.Helper()

		deviceId := os.Getenv("OKTA_TEST_DEVICE_ID")
		if deviceId == "" {
			t.Skip("Set OKTA_TEST_DEVICE_ID to a disposable enrolled device to run device lifecycle tests")
		}

		return deviceId
	}

	t.Run("Test DeviceAPIService ListDevices", func(t *testing.T) {
		resp, httpRes, err := apiClient.DeviceAPI.ListDevices(context.Background()).Limit(20).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		assert.IsType(t, []okta.DeviceList{}, resp)

		for _, device := range resp {
			assert.NotNil(t, device.Id)
			assert.NotNil(t, device.Status)
			if device.Profile != nil {
				assert.NotEmpty(t, device.Profile.DisplayName)
				assert.NotEmpty(t, device.Profile.Platform)
			}
		}
	})

	t.Run("Test DeviceAPIService GetDevice", func(t *testing.T) {
		device := requireExistingDevice(t)

		resp, httpRes, err := apiClient.DeviceAPI.GetDevice(context.Background(), *device.Id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		require.NotNil(t, resp.Id)
		assert.Equal(t, *device.Id, *resp.Id)
		assert.NotNil(t, resp.Status)
		assert.NotNil(t, resp.Profile)

		if resp.Profile != nil && device.Profile != nil {
			assert.Equal(t, device.Profile.DisplayName, resp.Profile.DisplayName)
			assert.Equal(t, device.Profile.Platform, resp.Profile.Platform)
		}
	})

	t.Run("Test DeviceAPIService ListDeviceUsers", func(t *testing.T) {
		device := requireExistingDevice(t)

		resp, httpRes, err := apiClient.DeviceAPI.ListDeviceUsers(context.Background(), *device.Id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		assert.IsType(t, []okta.DeviceUser{}, resp)

		for _, deviceUser := range resp {
			require.NotNil(t, deviceUser.User)
			assert.NotNil(t, deviceUser.User.Id)
		}
	})

	t.Run("Test DeviceAPIService SuspendDevice", func(t *testing.T) {
		deviceId := lifecycleDeviceId(t)

		httpRes, err := apiClient.DeviceAPI.SuspendDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)

		// Always put the device back the way it was found
		t.Cleanup(func() {
			if _, err := apiClient.DeviceAPI.UnsuspendDevice(context.Background(), deviceId).Execute(); err != nil {
				t.Logf("Failed to unsuspend device %s: %v", deviceId, err)
			}
		})

		resp, _, err := apiClient.DeviceAPI.GetDevice(context.Background(), deviceId).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp.Status)
		assert.Equal(t, "SUSPENDED", *resp.Status)
	})

	t.Run("Test DeviceAPIService UnsuspendDevice", func(t *testing.T) {
		deviceId := lifecycleDeviceId(t)

		// Suspend first so there's something to unsuspend
		_, err := apiClient.DeviceAPI.SuspendDevice(context.Background(), deviceId).Execute()
		require.Nil(t, err)

		httpRes, err := apiClient.DeviceAPI.UnsuspendDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)

		resp, _, err := apiClient.DeviceAPI.GetDevice(context.Background(), deviceId).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp.Status)
		assert.Equal(t, "ACTIVE", *resp.Status)
	})

	t.Run("Test DeviceAPIService ActivateDevice", func(t *testing.T) {
		// Activation only applies to a device that was deactivated, and deactivation
		// forces the end user to re-enroll, so this isn't safe to run against a real
		// org. The error path below covers the SDK call itself.
		t.Skip("Device activation requires a deactivated device and re-enrollment by the end user")
	})

	t.Run("Test DeviceAPIService DeactivateDevice", func(t *testing.T) {
		t.Skip("Device deactivation forces the end user to re-enroll the device")
	})

	t.Run("Test DeviceAPIService DeleteDevice", func(t *testing.T) {
		// A device must be deactivated before it can be deleted, and deletion is
		// irreversible - the device has to be enrolled again from scratch.
		t.Skip("Device deletion is irreversible and requires a deactivated device")
	})

	t.Run("Test DeviceAPIService Error Handling - Get Nonexistent Device", func(t *testing.T) {
		_, httpRes, err := apiClient.DeviceAPI.GetDevice(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test DeviceAPIService Error Handling - List Users For Nonexistent Device", func(t *testing.T) {
		_, httpRes, err := apiClient.DeviceAPI.ListDeviceUsers(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test DeviceAPIService Error Handling - Activate Nonexistent Device", func(t *testing.T) {
		httpRes, err := apiClient.DeviceAPI.ActivateDevice(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test DeviceAPIService Error Handling - Deactivate Nonexistent Device", func(t *testing.T) {
		httpRes, err := apiClient.DeviceAPI.DeactivateDevice(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test DeviceAPIService Error Handling - Suspend Nonexistent Device", func(t *testing.T) {
		httpRes, err := apiClient.DeviceAPI.SuspendDevice(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test DeviceAPIService Error Handling - Unsuspend Nonexistent Device", func(t *testing.T) {
		httpRes, err := apiClient.DeviceAPI.UnsuspendDevice(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})

	t.Run("Test DeviceAPIService Error Handling - Delete Nonexistent Device", func(t *testing.T) {
		httpRes, err := apiClient.DeviceAPI.DeleteDevice(context.Background(), nonExistentDeviceId).Execute()

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusNotFound, httpRes.StatusCode)
	})
}
