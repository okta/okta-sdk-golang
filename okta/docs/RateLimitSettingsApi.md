# \RateLimitSettingsApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRateLimitSettingsAdminNotifications**](RateLimitSettingsApi.md#GetRateLimitSettingsAdminNotifications) | **Get** /api/v1/rate-limit-settings/admin-notifications | Retrieve the Rate Limit Admin Notification Settings
[**GetRateLimitSettingsPerClient**](RateLimitSettingsApi.md#GetRateLimitSettingsPerClient) | **Get** /api/v1/rate-limit-settings/per-client | Retrieve the Per-Client Rate Limit Settings
[**ReplaceRateLimitSettingsAdminNotifications**](RateLimitSettingsApi.md#ReplaceRateLimitSettingsAdminNotifications) | **Put** /api/v1/rate-limit-settings/admin-notifications | Replace the Rate Limit Admin Notification Settings
[**ReplaceRateLimitSettingsPerClient**](RateLimitSettingsApi.md#ReplaceRateLimitSettingsPerClient) | **Put** /api/v1/rate-limit-settings/per-client | Replace the Per-Client Rate Limit Settings



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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsApi.GetRateLimitSettingsAdminNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsApi.GetRateLimitSettingsAdminNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimitSettingsAdminNotifications`: RateLimitAdminNotifications
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsApi.GetRateLimitSettingsAdminNotifications`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsApi.GetRateLimitSettingsPerClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsApi.GetRateLimitSettingsPerClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimitSettingsPerClient`: PerClientRateLimitSettings
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsApi.GetRateLimitSettingsPerClient`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    rateLimitAdminNotifications := *openapiclient.NewRateLimitAdminNotifications(false) // RateLimitAdminNotifications | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsApi.ReplaceRateLimitSettingsAdminNotifications(context.Background()).RateLimitAdminNotifications(rateLimitAdminNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsApi.ReplaceRateLimitSettingsAdminNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRateLimitSettingsAdminNotifications`: RateLimitAdminNotifications
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsApi.ReplaceRateLimitSettingsAdminNotifications`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    perClientRateLimitSettings := *openapiclient.NewPerClientRateLimitSettings(openapiclient.PerClientRateLimitMode("DISABLE")) // PerClientRateLimitSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitSettingsApi.ReplaceRateLimitSettingsPerClient(context.Background()).PerClientRateLimitSettings(perClientRateLimitSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitSettingsApi.ReplaceRateLimitSettingsPerClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRateLimitSettingsPerClient`: PerClientRateLimitSettings
    fmt.Fprintf(os.Stdout, "Response from `RateLimitSettingsApi.ReplaceRateLimitSettingsPerClient`: %v\n", resp)
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

