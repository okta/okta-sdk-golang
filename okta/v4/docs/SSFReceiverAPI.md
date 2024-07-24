# \SSFReceiverAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSecurityEventsProviderInstance**](SSFReceiverAPI.md#ActivateSecurityEventsProviderInstance) | **Post** /api/v1/security-events-providers/{securityEventProviderId}/lifecycle/activate | Activate a Security Events Provider
[**CreateSecurityEventsProviderInstance**](SSFReceiverAPI.md#CreateSecurityEventsProviderInstance) | **Post** /api/v1/security-events-providers | Create a Security Events Provider
[**DeactivateSecurityEventsProviderInstance**](SSFReceiverAPI.md#DeactivateSecurityEventsProviderInstance) | **Post** /api/v1/security-events-providers/{securityEventProviderId}/lifecycle/deactivate | Deactivate a Security Events Provider
[**DeleteSecurityEventsProviderInstance**](SSFReceiverAPI.md#DeleteSecurityEventsProviderInstance) | **Delete** /api/v1/security-events-providers/{securityEventProviderId} | Delete a Security Events Provider
[**GetSecurityEventsProviderInstance**](SSFReceiverAPI.md#GetSecurityEventsProviderInstance) | **Get** /api/v1/security-events-providers/{securityEventProviderId} | Retrieve the Security Events Provider
[**ListSecurityEventsProviderInstances**](SSFReceiverAPI.md#ListSecurityEventsProviderInstances) | **Get** /api/v1/security-events-providers | List all Security Events Providers
[**ReplaceSecurityEventsProviderInstance**](SSFReceiverAPI.md#ReplaceSecurityEventsProviderInstance) | **Put** /api/v1/security-events-providers/{securityEventProviderId} | Replace a Security Events Provider



## ActivateSecurityEventsProviderInstance

> SecurityEventsProviderResponse ActivateSecurityEventsProviderInstance(ctx, securityEventProviderId).Execute()

Activate a Security Events Provider



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
    securityEventProviderId := "sse1qg25RpusjUP6m0g5" // string | `id` of the Security Events Provider instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFReceiverAPI.ActivateSecurityEventsProviderInstance(context.Background(), securityEventProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.ActivateSecurityEventsProviderInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateSecurityEventsProviderInstance`: SecurityEventsProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SSFReceiverAPI.ActivateSecurityEventsProviderInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityEventProviderId** | **string** | &#x60;id&#x60; of the Security Events Provider instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSecurityEventsProviderInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityEventsProviderResponse**](SecurityEventsProviderResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityEventsProviderInstance

> SecurityEventsProviderResponse CreateSecurityEventsProviderInstance(ctx).Instance(instance).Execute()

Create a Security Events Provider



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
    instance := *openapiclient.NewSecurityEventsProviderRequest("Target SSF Provider", openapiclient.SecurityEventsProviderRequest_settings{SecurityEventsProviderSettingsNonSSFCompliant: openapiclient.NewSecurityEventsProviderSettingsNonSSFCompliant("example.okta.com", "https://example.okta.com/oauth2/v1/keys")}, "okta") // SecurityEventsProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFReceiverAPI.CreateSecurityEventsProviderInstance(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.CreateSecurityEventsProviderInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityEventsProviderInstance`: SecurityEventsProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SSFReceiverAPI.CreateSecurityEventsProviderInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityEventsProviderInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**SecurityEventsProviderRequest**](SecurityEventsProviderRequest.md) |  | 

### Return type

[**SecurityEventsProviderResponse**](SecurityEventsProviderResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSecurityEventsProviderInstance

> SecurityEventsProviderResponse DeactivateSecurityEventsProviderInstance(ctx, securityEventProviderId).Execute()

Deactivate a Security Events Provider



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
    securityEventProviderId := "sse1qg25RpusjUP6m0g5" // string | `id` of the Security Events Provider instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFReceiverAPI.DeactivateSecurityEventsProviderInstance(context.Background(), securityEventProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.DeactivateSecurityEventsProviderInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateSecurityEventsProviderInstance`: SecurityEventsProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SSFReceiverAPI.DeactivateSecurityEventsProviderInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityEventProviderId** | **string** | &#x60;id&#x60; of the Security Events Provider instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateSecurityEventsProviderInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityEventsProviderResponse**](SecurityEventsProviderResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityEventsProviderInstance

> DeleteSecurityEventsProviderInstance(ctx, securityEventProviderId).Execute()

Delete a Security Events Provider



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
    securityEventProviderId := "sse1qg25RpusjUP6m0g5" // string | `id` of the Security Events Provider instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSFReceiverAPI.DeleteSecurityEventsProviderInstance(context.Background(), securityEventProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.DeleteSecurityEventsProviderInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityEventProviderId** | **string** | &#x60;id&#x60; of the Security Events Provider instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityEventsProviderInstanceRequest struct via the builder pattern


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


## GetSecurityEventsProviderInstance

> SecurityEventsProviderResponse GetSecurityEventsProviderInstance(ctx, securityEventProviderId).Execute()

Retrieve the Security Events Provider



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
    securityEventProviderId := "sse1qg25RpusjUP6m0g5" // string | `id` of the Security Events Provider instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFReceiverAPI.GetSecurityEventsProviderInstance(context.Background(), securityEventProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.GetSecurityEventsProviderInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityEventsProviderInstance`: SecurityEventsProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SSFReceiverAPI.GetSecurityEventsProviderInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityEventProviderId** | **string** | &#x60;id&#x60; of the Security Events Provider instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityEventsProviderInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityEventsProviderResponse**](SecurityEventsProviderResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityEventsProviderInstances

> []SecurityEventsProviderResponse ListSecurityEventsProviderInstances(ctx).Execute()

List all Security Events Providers



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
    resp, r, err := apiClient.SSFReceiverAPI.ListSecurityEventsProviderInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.ListSecurityEventsProviderInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityEventsProviderInstances`: []SecurityEventsProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SSFReceiverAPI.ListSecurityEventsProviderInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityEventsProviderInstancesRequest struct via the builder pattern


### Return type

[**[]SecurityEventsProviderResponse**](SecurityEventsProviderResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSecurityEventsProviderInstance

> SecurityEventsProviderResponse ReplaceSecurityEventsProviderInstance(ctx, securityEventProviderId).Instance(instance).Execute()

Replace a Security Events Provider



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
    securityEventProviderId := "sse1qg25RpusjUP6m0g5" // string | `id` of the Security Events Provider instance
    instance := *openapiclient.NewSecurityEventsProviderRequest("Target SSF Provider", openapiclient.SecurityEventsProviderRequest_settings{SecurityEventsProviderSettingsNonSSFCompliant: openapiclient.NewSecurityEventsProviderSettingsNonSSFCompliant("example.okta.com", "https://example.okta.com/oauth2/v1/keys")}, "okta") // SecurityEventsProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFReceiverAPI.ReplaceSecurityEventsProviderInstance(context.Background(), securityEventProviderId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFReceiverAPI.ReplaceSecurityEventsProviderInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSecurityEventsProviderInstance`: SecurityEventsProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SSFReceiverAPI.ReplaceSecurityEventsProviderInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityEventProviderId** | **string** | &#x60;id&#x60; of the Security Events Provider instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSecurityEventsProviderInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**SecurityEventsProviderRequest**](SecurityEventsProviderRequest.md) |  | 

### Return type

[**SecurityEventsProviderResponse**](SecurityEventsProviderResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

