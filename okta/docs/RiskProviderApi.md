# \RiskProviderApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskProvider**](RiskProviderApi.md#CreateRiskProvider) | **Post** /api/v1/risk/providers | Create a Risk Provider
[**DeleteRiskProvider**](RiskProviderApi.md#DeleteRiskProvider) | **Delete** /api/v1/risk/providers/{riskProviderId} | Delete a Risk Provider
[**GetRiskProvider**](RiskProviderApi.md#GetRiskProvider) | **Get** /api/v1/risk/providers/{riskProviderId} | Retrieve a Risk Provider
[**ListRiskProviders**](RiskProviderApi.md#ListRiskProviders) | **Get** /api/v1/risk/providers | List all Risk Providers
[**ReplaceRiskProvider**](RiskProviderApi.md#ReplaceRiskProvider) | **Put** /api/v1/risk/providers/{riskProviderId} | Replace a Risk Provider



## CreateRiskProvider

> RiskProvider CreateRiskProvider(ctx).Instance(instance).Execute()

Create a Risk Provider



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
    instance := *openapiclient.NewRiskProvider("ClientId_example", "Name_example") // RiskProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderApi.CreateRiskProvider(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderApi.CreateRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskProvider`: RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderApi.CreateRiskProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**RiskProvider**](RiskProvider.md) |  | 

### Return type

[**RiskProvider**](RiskProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskProvider

> DeleteRiskProvider(ctx, riskProviderId).Execute()

Delete a Risk Provider



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
    riskProviderId := "00rp12r4skkjkjgsn" // string | `id` of the risk provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderApi.DeleteRiskProvider(context.Background(), riskProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderApi.DeleteRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProviderId** | **string** | &#x60;id&#x60; of the risk provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskProviderRequest struct via the builder pattern


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


## GetRiskProvider

> RiskProvider GetRiskProvider(ctx, riskProviderId).Execute()

Retrieve a Risk Provider



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
    riskProviderId := "00rp12r4skkjkjgsn" // string | `id` of the risk provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderApi.GetRiskProvider(context.Background(), riskProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderApi.GetRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskProvider`: RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderApi.GetRiskProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProviderId** | **string** | &#x60;id&#x60; of the risk provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskProvider**](RiskProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRiskProviders

> []RiskProvider ListRiskProviders(ctx).Execute()

List all Risk Providers



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderApi.ListRiskProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderApi.ListRiskProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRiskProviders`: []RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderApi.ListRiskProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRiskProvidersRequest struct via the builder pattern


### Return type

[**[]RiskProvider**](RiskProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRiskProvider

> RiskProvider ReplaceRiskProvider(ctx, riskProviderId).Instance(instance).Execute()

Replace a Risk Provider



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
    riskProviderId := "00rp12r4skkjkjgsn" // string | `id` of the risk provider
    instance := *openapiclient.NewRiskProvider("ClientId_example", "Name_example") // RiskProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderApi.ReplaceRiskProvider(context.Background(), riskProviderId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderApi.ReplaceRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRiskProvider`: RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderApi.ReplaceRiskProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProviderId** | **string** | &#x60;id&#x60; of the risk provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRiskProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**RiskProvider**](RiskProvider.md) |  | 

### Return type

[**RiskProvider**](RiskProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

