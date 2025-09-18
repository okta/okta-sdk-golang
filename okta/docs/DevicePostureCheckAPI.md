# \DevicePostureCheckAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevicePostureCheck**](DevicePostureCheckAPI.md#CreateDevicePostureCheck) | **Post** /api/v1/device-posture-checks | Create a device posture check
[**DeleteDevicePostureCheck**](DevicePostureCheckAPI.md#DeleteDevicePostureCheck) | **Delete** /api/v1/device-posture-checks/{postureCheckId} | Delete a device posture check
[**GetDevicePostureCheck**](DevicePostureCheckAPI.md#GetDevicePostureCheck) | **Get** /api/v1/device-posture-checks/{postureCheckId} | Retrieve a device posture check
[**ListDefaultDevicePostureChecks**](DevicePostureCheckAPI.md#ListDefaultDevicePostureChecks) | **Get** /api/v1/device-posture-checks/default | List all default device posture checks
[**ListDevicePostureChecks**](DevicePostureCheckAPI.md#ListDevicePostureChecks) | **Get** /api/v1/device-posture-checks | List all device posture checks
[**ReplaceDevicePostureCheck**](DevicePostureCheckAPI.md#ReplaceDevicePostureCheck) | **Put** /api/v1/device-posture-checks/{postureCheckId} | Replace a device posture check



## CreateDevicePostureCheck

> DevicePostureCheck CreateDevicePostureCheck(ctx).DevicePostureCheck(devicePostureCheck).Execute()

Create a device posture check



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
	devicePostureCheck := *openapiclient.NewDevicePostureCheck() // DevicePostureCheck | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicePostureCheckAPI.CreateDevicePostureCheck(context.Background()).DevicePostureCheck(devicePostureCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicePostureCheckAPI.CreateDevicePostureCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDevicePostureCheck`: DevicePostureCheck
	fmt.Fprintf(os.Stdout, "Response from `DevicePostureCheckAPI.CreateDevicePostureCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDevicePostureCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devicePostureCheck** | [**DevicePostureCheck**](DevicePostureCheck.md) |  | 

### Return type

[**DevicePostureCheck**](DevicePostureCheck.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevicePostureCheck

> DeleteDevicePostureCheck(ctx, postureCheckId).Execute()

Delete a device posture check



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
	postureCheckId := "postureCheckId_example" // string | ID of the device posture check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicePostureCheckAPI.DeleteDevicePostureCheck(context.Background(), postureCheckId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicePostureCheckAPI.DeleteDevicePostureCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postureCheckId** | **string** | ID of the device posture check | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevicePostureCheckRequest struct via the builder pattern


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


## GetDevicePostureCheck

> DevicePostureCheck GetDevicePostureCheck(ctx, postureCheckId).Execute()

Retrieve a device posture check



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
	postureCheckId := "postureCheckId_example" // string | ID of the device posture check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicePostureCheckAPI.GetDevicePostureCheck(context.Background(), postureCheckId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicePostureCheckAPI.GetDevicePostureCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicePostureCheck`: DevicePostureCheck
	fmt.Fprintf(os.Stdout, "Response from `DevicePostureCheckAPI.GetDevicePostureCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postureCheckId** | **string** | ID of the device posture check | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicePostureCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DevicePostureCheck**](DevicePostureCheck.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDefaultDevicePostureChecks

> []DevicePostureCheck ListDefaultDevicePostureChecks(ctx).Execute()

List all default device posture checks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicePostureCheckAPI.ListDefaultDevicePostureChecks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicePostureCheckAPI.ListDefaultDevicePostureChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDefaultDevicePostureChecks`: []DevicePostureCheck
	fmt.Fprintf(os.Stdout, "Response from `DevicePostureCheckAPI.ListDefaultDevicePostureChecks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDefaultDevicePostureChecksRequest struct via the builder pattern


### Return type

[**[]DevicePostureCheck**](DevicePostureCheck.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicePostureChecks

> []DevicePostureCheck ListDevicePostureChecks(ctx).Execute()

List all device posture checks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicePostureCheckAPI.ListDevicePostureChecks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicePostureCheckAPI.ListDevicePostureChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevicePostureChecks`: []DevicePostureCheck
	fmt.Fprintf(os.Stdout, "Response from `DevicePostureCheckAPI.ListDevicePostureChecks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDevicePostureChecksRequest struct via the builder pattern


### Return type

[**[]DevicePostureCheck**](DevicePostureCheck.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDevicePostureCheck

> DevicePostureCheck ReplaceDevicePostureCheck(ctx, postureCheckId).DevicePostureCheck(devicePostureCheck).Execute()

Replace a device posture check



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
	postureCheckId := "postureCheckId_example" // string | ID of the device posture check
	devicePostureCheck := *openapiclient.NewDevicePostureCheck() // DevicePostureCheck | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicePostureCheckAPI.ReplaceDevicePostureCheck(context.Background(), postureCheckId).DevicePostureCheck(devicePostureCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicePostureCheckAPI.ReplaceDevicePostureCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceDevicePostureCheck`: DevicePostureCheck
	fmt.Fprintf(os.Stdout, "Response from `DevicePostureCheckAPI.ReplaceDevicePostureCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postureCheckId** | **string** | ID of the device posture check | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDevicePostureCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicePostureCheck** | [**DevicePostureCheck**](DevicePostureCheck.md) |  | 

### Return type

[**DevicePostureCheck**](DevicePostureCheck.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

