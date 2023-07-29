# \PushProviderApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePushProvider**](PushProviderApi.md#CreatePushProvider) | **Post** /api/v1/push-providers | Create a Push Provider
[**DeletePushProvider**](PushProviderApi.md#DeletePushProvider) | **Delete** /api/v1/push-providers/{pushProviderId} | Delete a Push Provider
[**GetPushProvider**](PushProviderApi.md#GetPushProvider) | **Get** /api/v1/push-providers/{pushProviderId} | Retrieve a Push Provider
[**ListPushProviders**](PushProviderApi.md#ListPushProviders) | **Get** /api/v1/push-providers | List all Push Providers
[**ReplacePushProvider**](PushProviderApi.md#ReplacePushProvider) | **Put** /api/v1/push-providers/{pushProviderId} | Replace a Push Provider



## CreatePushProvider

> ListPushProviders200ResponseInner CreatePushProvider(ctx).PushProvider(pushProvider).Execute()

Create a Push Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pushProvider := openapiclient.listPushProviders_200_response_inner{APNSPushProvider: openapiclient.NewAPNSPushProvider()} // ListPushProviders200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushProviderApi.CreatePushProvider(context.Background()).PushProvider(pushProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushProviderApi.CreatePushProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushProvider`: ListPushProviders200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PushProviderApi.CreatePushProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushProvider** | [**ListPushProviders200ResponseInner**](ListPushProviders200ResponseInner.md) |  | 

### Return type

[**ListPushProviders200ResponseInner**](ListPushProviders200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePushProvider

> DeletePushProvider(ctx, pushProviderId).Execute()

Delete a Push Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pushProviderId := "pushProviderId_example" // string | Id of the push provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushProviderApi.DeletePushProvider(context.Background(), pushProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushProviderApi.DeletePushProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pushProviderId** | **string** | Id of the push provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePushProviderRequest struct via the builder pattern


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


## GetPushProvider

> ListPushProviders200ResponseInner GetPushProvider(ctx, pushProviderId).Execute()

Retrieve a Push Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pushProviderId := "pushProviderId_example" // string | Id of the push provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushProviderApi.GetPushProvider(context.Background(), pushProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushProviderApi.GetPushProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushProvider`: ListPushProviders200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PushProviderApi.GetPushProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pushProviderId** | **string** | Id of the push provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPushProviders200ResponseInner**](ListPushProviders200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPushProviders

> []ListPushProviders200ResponseInner ListPushProviders(ctx).Type_(type_).Execute()

List all Push Providers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := openapiclient.ProviderType("APNS") // ProviderType | Filters push providers by `providerType` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushProviderApi.ListPushProviders(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushProviderApi.ListPushProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPushProviders`: []ListPushProviders200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PushProviderApi.ListPushProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPushProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ProviderType**](ProviderType.md) | Filters push providers by &#x60;providerType&#x60; | 

### Return type

[**[]ListPushProviders200ResponseInner**](ListPushProviders200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePushProvider

> ListPushProviders200ResponseInner ReplacePushProvider(ctx, pushProviderId).PushProvider(pushProvider).Execute()

Replace a Push Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pushProviderId := "pushProviderId_example" // string | Id of the push provider
    pushProvider := openapiclient.listPushProviders_200_response_inner{APNSPushProvider: openapiclient.NewAPNSPushProvider()} // ListPushProviders200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushProviderApi.ReplacePushProvider(context.Background(), pushProviderId).PushProvider(pushProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushProviderApi.ReplacePushProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePushProvider`: ListPushProviders200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PushProviderApi.ReplacePushProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pushProviderId** | **string** | Id of the push provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePushProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushProvider** | [**ListPushProviders200ResponseInner**](ListPushProviders200ResponseInner.md) |  | 

### Return type

[**ListPushProviders200ResponseInner**](ListPushProviders200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

