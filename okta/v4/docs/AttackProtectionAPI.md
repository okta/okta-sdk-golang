# \AttackProtectionAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthenticatorSettings**](AttackProtectionAPI.md#GetAuthenticatorSettings) | **Get** /attack-protection/api/v1/authenticator-settings | Retrieve the Authenticator Settings
[**GetUserLockoutSettings**](AttackProtectionAPI.md#GetUserLockoutSettings) | **Get** /attack-protection/api/v1/user-lockout-settings | Retrieve the User Lockout Settings
[**ReplaceAuthenticatorSettings**](AttackProtectionAPI.md#ReplaceAuthenticatorSettings) | **Put** /attack-protection/api/v1/authenticator-settings | Replace the Authenticator Settings
[**ReplaceUserLockoutSettings**](AttackProtectionAPI.md#ReplaceUserLockoutSettings) | **Put** /attack-protection/api/v1/user-lockout-settings | Replace the User Lockout Settings



## GetAuthenticatorSettings

> []AttackProtectionAuthenticatorSettings GetAuthenticatorSettings(ctx).Execute()

Retrieve the Authenticator Settings



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
    resp, r, err := apiClient.AttackProtectionAPI.GetAuthenticatorSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttackProtectionAPI.GetAuthenticatorSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticatorSettings`: []AttackProtectionAuthenticatorSettings
    fmt.Fprintf(os.Stdout, "Response from `AttackProtectionAPI.GetAuthenticatorSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatorSettingsRequest struct via the builder pattern


### Return type

[**[]AttackProtectionAuthenticatorSettings**](AttackProtectionAuthenticatorSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ReplaceAuthenticatorSettings

> AttackProtectionAuthenticatorSettings ReplaceAuthenticatorSettings(ctx).AuthenticatorSettings(authenticatorSettings).Execute()

Replace the Authenticator Settings



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
    authenticatorSettings := *openapiclient.NewAttackProtectionAuthenticatorSettings() // AttackProtectionAuthenticatorSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttackProtectionAPI.ReplaceAuthenticatorSettings(context.Background()).AuthenticatorSettings(authenticatorSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttackProtectionAPI.ReplaceAuthenticatorSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthenticatorSettings`: AttackProtectionAuthenticatorSettings
    fmt.Fprintf(os.Stdout, "Response from `AttackProtectionAPI.ReplaceAuthenticatorSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthenticatorSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorSettings** | [**AttackProtectionAuthenticatorSettings**](AttackProtectionAuthenticatorSettings.md) |  | 

### Return type

[**AttackProtectionAuthenticatorSettings**](AttackProtectionAuthenticatorSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
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

