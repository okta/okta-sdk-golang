# \ThreatInsightApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentConfiguration**](ThreatInsightApi.md#GetCurrentConfiguration) | **Get** /api/v1/threats/configuration | Retrieve the ThreatInsight Configuration
[**UpdateConfiguration**](ThreatInsightApi.md#UpdateConfiguration) | **Post** /api/v1/threats/configuration | Update the ThreatInsight Configuration



## GetCurrentConfiguration

> ThreatInsightConfiguration GetCurrentConfiguration(ctx).Execute()

Retrieve the ThreatInsight Configuration



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
    resp, r, err := apiClient.ThreatInsightApi.GetCurrentConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThreatInsightApi.GetCurrentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentConfiguration`: ThreatInsightConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ThreatInsightApi.GetCurrentConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentConfigurationRequest struct via the builder pattern


### Return type

[**ThreatInsightConfiguration**](ThreatInsightConfiguration.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ThreatInsightConfiguration UpdateConfiguration(ctx).ThreatInsightConfiguration(threatInsightConfiguration).Execute()

Update the ThreatInsight Configuration



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
    threatInsightConfiguration := *openapiclient.NewThreatInsightConfiguration() // ThreatInsightConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThreatInsightApi.UpdateConfiguration(context.Background()).ThreatInsightConfiguration(threatInsightConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThreatInsightApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: ThreatInsightConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ThreatInsightApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threatInsightConfiguration** | [**ThreatInsightConfiguration**](ThreatInsightConfiguration.md) |  | 

### Return type

[**ThreatInsightConfiguration**](ThreatInsightConfiguration.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

