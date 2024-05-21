# \ApplicationOktaApplicationSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirstPartyAppSettings**](ApplicationOktaApplicationSettingsAPI.md#GetFirstPartyAppSettings) | **Get** /api/v1/first-party-app-settings/{appName} | Retrieve the Okta app settings
[**ReplaceFirstPartyAppSettings**](ApplicationOktaApplicationSettingsAPI.md#ReplaceFirstPartyAppSettings) | **Put** /api/v1/first-party-app-settings/{appName} | Replace the Okta app settings



## GetFirstPartyAppSettings

> AdminConsoleSettings GetFirstPartyAppSettings(ctx, appName).Execute()

Retrieve the Okta app settings



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
    appName := "admin-console" // string | `appName` of the application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationOktaApplicationSettingsAPI.GetFirstPartyAppSettings(context.Background(), appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationOktaApplicationSettingsAPI.GetFirstPartyAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirstPartyAppSettings`: AdminConsoleSettings
    fmt.Fprintf(os.Stdout, "Response from `ApplicationOktaApplicationSettingsAPI.GetFirstPartyAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | &#x60;appName&#x60; of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirstPartyAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminConsoleSettings**](AdminConsoleSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFirstPartyAppSettings

> AdminConsoleSettings ReplaceFirstPartyAppSettings(ctx, appName).AdminConsoleSettings(adminConsoleSettings).Execute()

Replace the Okta app settings



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
    appName := "admin-console" // string | `appName` of the application
    adminConsoleSettings := *openapiclient.NewAdminConsoleSettings() // AdminConsoleSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationOktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings(context.Background(), appName).AdminConsoleSettings(adminConsoleSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationOktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceFirstPartyAppSettings`: AdminConsoleSettings
    fmt.Fprintf(os.Stdout, "Response from `ApplicationOktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | &#x60;appName&#x60; of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFirstPartyAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminConsoleSettings** | [**AdminConsoleSettings**](AdminConsoleSettings.md) |  | 

### Return type

[**AdminConsoleSettings**](AdminConsoleSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

