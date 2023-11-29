# \AttackProtectionAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserLockoutSettings**](AttackProtectionAPI.md#GetUserLockoutSettings) | **Get** /attack-protection/api/v1/user-lockout-settings | Retrieve the User Lockout Settings
[**ReplaceUserLockoutSettings**](AttackProtectionAPI.md#ReplaceUserLockoutSettings) | **Put** /attack-protection/api/v1/user-lockout-settings | Replace the User Lockout Settings



## GetUserLockoutSettings

> []UserLockoutSettings GetUserLockoutSettings(ctx).Execute()

Retrieve the User Lockout Settings



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
    resp, r, err := apiClient.AttackProtectionAPI.GetUserLockoutSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttackProtectionAPI.GetUserLockoutSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserLockoutSettings`: []UserLockoutSettings
    fmt.Fprintf(os.Stdout, "Response from `AttackProtectionAPI.GetUserLockoutSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLockoutSettingsRequest struct via the builder pattern


### Return type

[**[]UserLockoutSettings**](UserLockoutSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUserLockoutSettings

> UserLockoutSettings ReplaceUserLockoutSettings(ctx).LockoutSettings(lockoutSettings).Execute()

Replace the User Lockout Settings



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
    lockoutSettings := *openapiclient.NewUserLockoutSettings() // UserLockoutSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttackProtectionAPI.ReplaceUserLockoutSettings(context.Background()).LockoutSettings(lockoutSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttackProtectionAPI.ReplaceUserLockoutSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceUserLockoutSettings`: UserLockoutSettings
    fmt.Fprintf(os.Stdout, "Response from `AttackProtectionAPI.ReplaceUserLockoutSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserLockoutSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lockoutSettings** | [**UserLockoutSettings**](UserLockoutSettings.md) |  | 

### Return type

[**UserLockoutSettings**](UserLockoutSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

