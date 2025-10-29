# \DeviceAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateDevice**](DeviceAPI.md#ActivateDevice) | **Post** /api/v1/devices/{deviceId}/lifecycle/activate | Activate a device
[**DeactivateDevice**](DeviceAPI.md#DeactivateDevice) | **Post** /api/v1/devices/{deviceId}/lifecycle/deactivate | Deactivate a device
[**DeleteDevice**](DeviceAPI.md#DeleteDevice) | **Delete** /api/v1/devices/{deviceId} | Delete a device
[**GetDevice**](DeviceAPI.md#GetDevice) | **Get** /api/v1/devices/{deviceId} | Retrieve a device
[**ListDeviceUsers**](DeviceAPI.md#ListDeviceUsers) | **Get** /api/v1/devices/{deviceId}/users | List all users for a device
[**ListDevices**](DeviceAPI.md#ListDevices) | **Get** /api/v1/devices | List all devices
[**SuspendDevice**](DeviceAPI.md#SuspendDevice) | **Post** /api/v1/devices/{deviceId}/lifecycle/suspend | Suspend a Device
[**UnsuspendDevice**](DeviceAPI.md#UnsuspendDevice) | **Post** /api/v1/devices/{deviceId}/lifecycle/unsuspend | Unsuspend a Device



## ActivateDevice

> ActivateDevice(ctx, deviceId).Execute()

Activate a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceAPI.ActivateDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ActivateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateDevice

> DeactivateDevice(ctx, deviceId).Execute()

Deactivate a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceAPI.DeactivateDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeactivateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, deviceId).Execute()

Delete a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceAPI.DeleteDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> Device GetDevice(ctx, deviceId).Execute()

Retrieve a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.GetDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.GetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevice`: Device
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Device**](Device.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceUsers

> []DeviceUser ListDeviceUsers(ctx, deviceId).Execute()

List all users for a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ListDeviceUsers(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ListDeviceUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceUsers`: []DeviceUser
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ListDeviceUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeviceUser**](DeviceUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevices

> []DeviceList ListDevices(ctx).After(after).Limit(limit).Search(search).Expand(expand).Execute()

List all devices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	after := "200u3des4afA47rYJu1d7" // string |  (optional)
	limit := int32(20) // int32 | A limit on the number of objects to return (recommend `20`) (optional) (default to 200)
	search := "status%20eq%20%22ACTIVE%22" // string | A SCIM filter expression that filters the results. Searches include all device `profile` properties and the device `id`, `status`, and `lastUpdated` properties.  Searches for devices can be filtered by the contains (`co`) operator. You can only use `co` with these select device profile attributes: `profile.displayName`, `profile.serialNumber`, `profile.imei`, `profile.meid`, `profile.udid`, and `profile.sid`. See [Operators](https://developer.okta.com/docs/api/#operators). (optional)
	expand := "user" // string | Includes associated user details and management status for the device in the `_embedded` attribute (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.ListDevices(context.Background()).After(after).Limit(limit).Search(search).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.ListDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevices`: []DeviceList
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.ListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** |  | 
 **limit** | **int32** | A limit on the number of objects to return (recommend &#x60;20&#x60;) | [default to 200]
 **search** | **string** | A SCIM filter expression that filters the results. Searches include all device &#x60;profile&#x60; properties and the device &#x60;id&#x60;, &#x60;status&#x60;, and &#x60;lastUpdated&#x60; properties.  Searches for devices can be filtered by the contains (&#x60;co&#x60;) operator. You can only use &#x60;co&#x60; with these select device profile attributes: &#x60;profile.displayName&#x60;, &#x60;profile.serialNumber&#x60;, &#x60;profile.imei&#x60;, &#x60;profile.meid&#x60;, &#x60;profile.udid&#x60;, and &#x60;profile.sid&#x60;. See [Operators](https://developer.okta.com/docs/api/#operators). | 
 **expand** | **string** | Includes associated user details and management status for the device in the &#x60;_embedded&#x60; attribute | 

### Return type

[**[]DeviceList**](DeviceList.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendDevice

> SuspendDevice(ctx, deviceId).Execute()

Suspend a Device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceAPI.SuspendDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.SuspendDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsuspendDevice

> UnsuspendDevice(ctx, deviceId).Execute()

Unsuspend a Device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	deviceId := "guo4a5u7JHHhjXrMK0g4" // string | `id` of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceAPI.UnsuspendDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.UnsuspendDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | &#x60;id&#x60; of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsuspendDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

