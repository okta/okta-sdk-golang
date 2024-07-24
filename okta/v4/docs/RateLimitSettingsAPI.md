# \RateLimitSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRateLimitSettingsAdminNotifications**](RateLimitSettingsAPI.md#GetRateLimitSettingsAdminNotifications) | **Get** /api/v1/rate-limit-settings/admin-notifications | Retrieve the Rate Limit Admin Notification Settings
[**GetRateLimitSettingsPerClient**](RateLimitSettingsAPI.md#GetRateLimitSettingsPerClient) | **Get** /api/v1/rate-limit-settings/per-client | Retrieve the Per-Client Rate Limit Settings
[**GetRateLimitSettingsWarningThreshold**](RateLimitSettingsAPI.md#GetRateLimitSettingsWarningThreshold) | **Get** /api/v1/rate-limit-settings/warning-threshold | Retrieve the Rate Limit Warning Threshold Percentage
[**ReplaceRateLimitSettingsAdminNotifications**](RateLimitSettingsAPI.md#ReplaceRateLimitSettingsAdminNotifications) | **Put** /api/v1/rate-limit-settings/admin-notifications | Replace the Rate Limit Admin Notification Settings
[**ReplaceRateLimitSettingsPerClient**](RateLimitSettingsAPI.md#ReplaceRateLimitSettingsPerClient) | **Put** /api/v1/rate-limit-settings/per-client | Replace the Per-Client Rate Limit Settings
[**ReplaceRateLimitSettingsWarningThreshold**](RateLimitSettingsAPI.md#ReplaceRateLimitSettingsWarningThreshold) | **Put** /api/v1/rate-limit-settings/warning-threshold | Replace the Rate Limit Warning Threshold Percentage



## GetRateLimitSettingsAdminNotifications

> RateLimitAdminNotifications GetRateLimitSettingsAdminNotifications(ctx).Execute()

Retrieve the Rate Limit Admin Notification Settings



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
    resp, r, err := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsAdminNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsAPI.GetRateLimitSettingsAdminNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimitSettingsAdminNotifications`: RateLimitAdminNotifications
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsAPI.GetRateLimitSettingsAdminNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitSettingsAdminNotificationsRequest struct via the builder pattern


### Return type

[**RateLimitAdminNotifications**](RateLimitAdminNotifications.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimitSettingsPerClient

> PerClientRateLimitSettings GetRateLimitSettingsPerClient(ctx).Execute()

Retrieve the Per-Client Rate Limit Settings



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
    resp, r, err := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsPerClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsAPI.GetRateLimitSettingsPerClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimitSettingsPerClient`: PerClientRateLimitSettings
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsAPI.GetRateLimitSettingsPerClient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitSettingsPerClientRequest struct via the builder pattern


### Return type

[**PerClientRateLimitSettings**](PerClientRateLimitSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimitSettingsWarningThreshold

> RateLimitWarningThresholdResponse GetRateLimitSettingsWarningThreshold(ctx).Execute()

Retrieve the Rate Limit Warning Threshold Percentage



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
    resp, r, err := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsWarningThreshold(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsAPI.GetRateLimitSettingsWarningThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimitSettingsWarningThreshold`: RateLimitWarningThresholdResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsAPI.GetRateLimitSettingsWarningThreshold`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitSettingsWarningThresholdRequest struct via the builder pattern


### Return type

[**RateLimitWarningThresholdResponse**](RateLimitWarningThresholdResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRateLimitSettingsAdminNotifications

> RateLimitAdminNotifications ReplaceRateLimitSettingsAdminNotifications(ctx).RateLimitAdminNotifications(rateLimitAdminNotifications).Execute()

Replace the Rate Limit Admin Notification Settings



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
    rateLimitAdminNotifications := *openapiclient.NewRateLimitAdminNotifications(false) // RateLimitAdminNotifications | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsAdminNotifications(context.Background()).RateLimitAdminNotifications(rateLimitAdminNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsAPI.ReplaceRateLimitSettingsAdminNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRateLimitSettingsAdminNotifications`: RateLimitAdminNotifications
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsAPI.ReplaceRateLimitSettingsAdminNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRateLimitSettingsAdminNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rateLimitAdminNotifications** | [**RateLimitAdminNotifications**](RateLimitAdminNotifications.md) |  | 

### Return type

[**RateLimitAdminNotifications**](RateLimitAdminNotifications.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRateLimitSettingsPerClient

> PerClientRateLimitSettings ReplaceRateLimitSettingsPerClient(ctx).PerClientRateLimitSettings(perClientRateLimitSettings).Execute()

Replace the Per-Client Rate Limit Settings



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
    perClientRateLimitSettings := *openapiclient.NewPerClientRateLimitSettings("DefaultMode_example") // PerClientRateLimitSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsPerClient(context.Background()).PerClientRateLimitSettings(perClientRateLimitSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsAPI.ReplaceRateLimitSettingsPerClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRateLimitSettingsPerClient`: PerClientRateLimitSettings
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsAPI.ReplaceRateLimitSettingsPerClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRateLimitSettingsPerClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perClientRateLimitSettings** | [**PerClientRateLimitSettings**](PerClientRateLimitSettings.md) |  | 

### Return type

[**PerClientRateLimitSettings**](PerClientRateLimitSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRateLimitSettingsWarningThreshold

> RateLimitWarningThresholdResponse ReplaceRateLimitSettingsWarningThreshold(ctx).RateLimitWarningThreshold(rateLimitWarningThreshold).Execute()

Replace the Rate Limit Warning Threshold Percentage



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
    rateLimitWarningThreshold := *openapiclient.NewRateLimitWarningThresholdRequest(int32(123)) // RateLimitWarningThresholdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsWarningThreshold(context.Background()).RateLimitWarningThreshold(rateLimitWarningThreshold).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsAPI.ReplaceRateLimitSettingsWarningThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRateLimitSettingsWarningThreshold`: RateLimitWarningThresholdResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsAPI.ReplaceRateLimitSettingsWarningThreshold`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRateLimitSettingsWarningThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rateLimitWarningThreshold** | [**RateLimitWarningThresholdRequest**](RateLimitWarningThresholdRequest.md) |  | 

### Return type

[**RateLimitWarningThresholdResponse**](RateLimitWarningThresholdResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

