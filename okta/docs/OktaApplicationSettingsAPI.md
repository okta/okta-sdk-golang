# \OktaApplicationSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirstPartyAppSettings**](OktaApplicationSettingsAPI.md#GetFirstPartyAppSettings) | **Get** /api/v1/first-party-app-settings/{appName} | Retrieve the Okta Application Settings
[**ReplaceFirstPartyAppSettings**](OktaApplicationSettingsAPI.md#ReplaceFirstPartyAppSettings) | **Put** /api/v1/first-party-app-settings/{appName} | Replace the Okta Application Settings



## GetFirstPartyAppSettings

> AdminConsoleSettings GetFirstPartyAppSettings(ctx, appName).Execute()

Retrieve the Okta Application Settings



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
    appName := "admin-console" // string | The key name for the Okta app.<br> Supported apps:   * Okta Admin Console (`admin-console`) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OktaApplicationSettingsAPI.GetFirstPartyAppSettings(context.Background(), appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OktaApplicationSettingsAPI.GetFirstPartyAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirstPartyAppSettings`: AdminConsoleSettings
    fmt.Fprintf(os.Stdout, "Response from `OktaApplicationSettingsAPI.GetFirstPartyAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | The key name for the Okta app.&lt;br&gt; Supported apps:   * Okta Admin Console (&#x60;admin-console&#x60;)  | 

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

Replace the Okta Application Settings



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
    appName := "admin-console" // string | The key name for the Okta app.<br> Supported apps:   * Okta Admin Console (`admin-console`) 
    adminConsoleSettings := *openapiclient.NewAdminConsoleSettings() // AdminConsoleSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings(context.Background(), appName).AdminConsoleSettings(adminConsoleSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceFirstPartyAppSettings`: AdminConsoleSettings
    fmt.Fprintf(os.Stdout, "Response from `OktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | The key name for the Okta app.&lt;br&gt; Supported apps:   * Okta Admin Console (&#x60;admin-console&#x60;)  | 

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

