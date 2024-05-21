# \DeviceAssuranceAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceAssurancePolicy**](DeviceAssuranceAPI.md#CreateDeviceAssurancePolicy) | **Post** /api/v1/device-assurances | Create a Device Assurance Policy
[**DeleteDeviceAssurancePolicy**](DeviceAssuranceAPI.md#DeleteDeviceAssurancePolicy) | **Delete** /api/v1/device-assurances/{deviceAssuranceId} | Delete a Device Assurance Policy
[**GetDeviceAssurancePolicy**](DeviceAssuranceAPI.md#GetDeviceAssurancePolicy) | **Get** /api/v1/device-assurances/{deviceAssuranceId} | Retrieve a Device Assurance Policy
[**ListDeviceAssurancePolicies**](DeviceAssuranceAPI.md#ListDeviceAssurancePolicies) | **Get** /api/v1/device-assurances | List all Device Assurance Policies
[**ReplaceDeviceAssurancePolicy**](DeviceAssuranceAPI.md#ReplaceDeviceAssurancePolicy) | **Put** /api/v1/device-assurances/{deviceAssuranceId} | Replace a Device Assurance Policy



## CreateDeviceAssurancePolicy

> ListDeviceAssurancePolicies200ResponseInner CreateDeviceAssurancePolicy(ctx).DeviceAssurance(deviceAssurance).Execute()

Create a Device Assurance Policy



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
    deviceAssurance := openapiclient.listDeviceAssurancePolicies_200_response_inner{DeviceAssuranceAndroidPlatform: openapiclient.NewDeviceAssuranceAndroidPlatform()} // ListDeviceAssurancePolicies200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceAPI.CreateDeviceAssurancePolicy(context.Background()).DeviceAssurance(deviceAssurance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceAPI.CreateDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceAssurancePolicy`: ListDeviceAssurancePolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceAPI.CreateDeviceAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceAssurance** | [**ListDeviceAssurancePolicies200ResponseInner**](ListDeviceAssurancePolicies200ResponseInner.md) |  | 

### Return type

[**ListDeviceAssurancePolicies200ResponseInner**](ListDeviceAssurancePolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceAssurancePolicy

> DeleteDeviceAssurancePolicy(ctx, deviceAssuranceId).Execute()

Delete a Device Assurance Policy



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
    deviceAssuranceId := "deviceAssuranceId_example" // string | Id of the Device Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeviceAssuranceAPI.DeleteDeviceAssurancePolicy(context.Background(), deviceAssuranceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceAPI.DeleteDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAssuranceId** | **string** | Id of the Device Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceAssurancePolicyRequest struct via the builder pattern


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


## GetDeviceAssurancePolicy

> ListDeviceAssurancePolicies200ResponseInner GetDeviceAssurancePolicy(ctx, deviceAssuranceId).Execute()

Retrieve a Device Assurance Policy



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
    deviceAssuranceId := "deviceAssuranceId_example" // string | Id of the Device Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceAPI.GetDeviceAssurancePolicy(context.Background(), deviceAssuranceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceAPI.GetDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceAssurancePolicy`: ListDeviceAssurancePolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceAPI.GetDeviceAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAssuranceId** | **string** | Id of the Device Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListDeviceAssurancePolicies200ResponseInner**](ListDeviceAssurancePolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceAssurancePolicies

> []ListDeviceAssurancePolicies200ResponseInner ListDeviceAssurancePolicies(ctx).Execute()

List all Device Assurance Policies



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
    resp, r, err := apiClient.DeviceAssuranceAPI.ListDeviceAssurancePolicies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceAPI.ListDeviceAssurancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceAssurancePolicies`: []ListDeviceAssurancePolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceAPI.ListDeviceAssurancePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceAssurancePoliciesRequest struct via the builder pattern


### Return type

[**[]ListDeviceAssurancePolicies200ResponseInner**](ListDeviceAssurancePolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDeviceAssurancePolicy

> ListDeviceAssurancePolicies200ResponseInner ReplaceDeviceAssurancePolicy(ctx, deviceAssuranceId).DeviceAssurance(deviceAssurance).Execute()

Replace a Device Assurance Policy



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
    deviceAssuranceId := "deviceAssuranceId_example" // string | Id of the Device Assurance Policy
    deviceAssurance := openapiclient.listDeviceAssurancePolicies_200_response_inner{DeviceAssuranceAndroidPlatform: openapiclient.NewDeviceAssuranceAndroidPlatform()} // ListDeviceAssurancePolicies200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAssuranceAPI.ReplaceDeviceAssurancePolicy(context.Background(), deviceAssuranceId).DeviceAssurance(deviceAssurance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssuranceAPI.ReplaceDeviceAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDeviceAssurancePolicy`: ListDeviceAssurancePolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssuranceAPI.ReplaceDeviceAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAssuranceId** | **string** | Id of the Device Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDeviceAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceAssurance** | [**ListDeviceAssurancePolicies200ResponseInner**](ListDeviceAssurancePolicies200ResponseInner.md) |  | 

### Return type

[**ListDeviceAssurancePolicies200ResponseInner**](ListDeviceAssurancePolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

