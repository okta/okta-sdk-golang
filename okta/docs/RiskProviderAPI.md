# \RiskProviderAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskProvider**](RiskProviderAPI.md#CreateRiskProvider) | **Post** /api/v1/risk/providers | Create a Risk Provider
[**DeleteRiskProvider**](RiskProviderAPI.md#DeleteRiskProvider) | **Delete** /api/v1/risk/providers/{riskProviderId} | Delete a Risk Provider
[**GetRiskProvider**](RiskProviderAPI.md#GetRiskProvider) | **Get** /api/v1/risk/providers/{riskProviderId} | Retrieve a Risk Provider
[**ListRiskProviders**](RiskProviderAPI.md#ListRiskProviders) | **Get** /api/v1/risk/providers | List all Risk Providers
[**ReplaceRiskProvider**](RiskProviderAPI.md#ReplaceRiskProvider) | **Put** /api/v1/risk/providers/{riskProviderId} | Replace a Risk Provider



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    instance := *openapiclient.NewRiskProvider("Action_example", "00cjkjjkkgjkdkjdkkljjsd", "00rp12r4skkjkjgsn", "Risk-Partner-X", *openapiclient.NewLinksSelf()) // RiskProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderAPI.CreateRiskProvider(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderAPI.CreateRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskProvider`: RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderAPI.CreateRiskProvider`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    riskProviderId := "00rp12r4skkjkjgsn" // string | `id` of the Risk Provider object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RiskProviderAPI.DeleteRiskProvider(context.Background(), riskProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderAPI.DeleteRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProviderId** | **string** | &#x60;id&#x60; of the Risk Provider object | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    riskProviderId := "00rp12r4skkjkjgsn" // string | `id` of the Risk Provider object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderAPI.GetRiskProvider(context.Background(), riskProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderAPI.GetRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskProvider`: RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderAPI.GetRiskProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProviderId** | **string** | &#x60;id&#x60; of the Risk Provider object | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderAPI.ListRiskProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderAPI.ListRiskProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRiskProviders`: []RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderAPI.ListRiskProviders`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    riskProviderId := "00rp12r4skkjkjgsn" // string | `id` of the Risk Provider object
    instance := *openapiclient.NewRiskProvider("Action_example", "00cjkjjkkgjkdkjdkkljjsd", "00rp12r4skkjkjgsn", "Risk-Partner-X", *openapiclient.NewLinksSelf()) // RiskProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskProviderAPI.ReplaceRiskProvider(context.Background(), riskProviderId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskProviderAPI.ReplaceRiskProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRiskProvider`: RiskProvider
    fmt.Fprintf(os.Stdout, "Response from `RiskProviderAPI.ReplaceRiskProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskProviderId** | **string** | &#x60;id&#x60; of the Risk Provider object | 

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

