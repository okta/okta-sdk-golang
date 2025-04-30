# \UserAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](UserAPI.md#ActivateUser) | **Post** /api/v1/users/{userId}/lifecycle/activate | Activate a User
[**ChangePassword**](UserAPI.md#ChangePassword) | **Post** /api/v1/users/{userId}/credentials/change_password | Change Password
[**ChangeRecoveryQuestion**](UserAPI.md#ChangeRecoveryQuestion) | **Post** /api/v1/users/{userId}/credentials/change_recovery_question | Change Recovery Question
[**CreateUser**](UserAPI.md#CreateUser) | **Post** /api/v1/users | Create a User
[**DeactivateUser**](UserAPI.md#DeactivateUser) | **Post** /api/v1/users/{userId}/lifecycle/deactivate | Deactivate a User
[**DeleteLinkedObjectForUser**](UserAPI.md#DeleteLinkedObjectForUser) | **Delete** /api/v1/users/{userIdOrLogin}/linkedObjects/{relationshipName} | Delete a Linked Object
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /api/v1/users/{userId} | Delete a User
[**ExpirePassword**](UserAPI.md#ExpirePassword) | **Post** /api/v1/users/{userId}/lifecycle/expire_password | Expire Password
[**ExpirePasswordAndGetTemporaryPassword**](UserAPI.md#ExpirePasswordAndGetTemporaryPassword) | **Post** /api/v1/users/{userId}/lifecycle/expire_password_with_temp_password | Expire Password and Set Temporary Password
[**ForgotPassword**](UserAPI.md#ForgotPassword) | **Post** /api/v1/users/{userId}/credentials/forgot_password | Initiate Forgot Password
[**ForgotPasswordSetNewPassword**](UserAPI.md#ForgotPasswordSetNewPassword) | **Post** /api/v1/users/{userId}/credentials/forgot_password_recovery_question | Reset Password with Recovery Question
[**GenerateResetPasswordToken**](UserAPI.md#GenerateResetPasswordToken) | **Post** /api/v1/users/{userId}/lifecycle/reset_password | Generate a Reset Password Token
[**GetRefreshTokenForUserAndClient**](UserAPI.md#GetRefreshTokenForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/tokens/{tokenId} | Retrieve a Refresh Token for a Client
[**GetUser**](UserAPI.md#GetUser) | **Get** /api/v1/users/{userId} | Retrieve a User
[**GetUserGrant**](UserAPI.md#GetUserGrant) | **Get** /api/v1/users/{userId}/grants/{grantId} | Retrieve a User Grant
[**ListAppLinks**](UserAPI.md#ListAppLinks) | **Get** /api/v1/users/{userId}/appLinks | List all Assigned Application Links
[**ListGrantsForUserAndClient**](UserAPI.md#ListGrantsForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/grants | List all Grants for a Client
[**ListLinkedObjectsForUser**](UserAPI.md#ListLinkedObjectsForUser) | **Get** /api/v1/users/{userIdOrLogin}/linkedObjects/{relationshipName} | List the primary or all of the associated Linked Object values
[**ListRefreshTokensForUserAndClient**](UserAPI.md#ListRefreshTokensForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/tokens | List all Refresh Tokens for a Client
[**ListUserBlocks**](UserAPI.md#ListUserBlocks) | **Get** /api/v1/users/{userId}/blocks | List all User Blocks
[**ListUserClients**](UserAPI.md#ListUserClients) | **Get** /api/v1/users/{userId}/clients | List all Clients
[**ListUserGrants**](UserAPI.md#ListUserGrants) | **Get** /api/v1/users/{userId}/grants | List all User Grants
[**ListUserGroups**](UserAPI.md#ListUserGroups) | **Get** /api/v1/users/{userId}/groups | List all Groups
[**ListUserIdentityProviders**](UserAPI.md#ListUserIdentityProviders) | **Get** /api/v1/users/{userId}/idps | List all Identity Providers
[**ListUsers**](UserAPI.md#ListUsers) | **Get** /api/v1/users | List all Users
[**ReactivateUser**](UserAPI.md#ReactivateUser) | **Post** /api/v1/users/{userId}/lifecycle/reactivate | Reactivate a User
[**ReplaceLinkedObjectForUser**](UserAPI.md#ReplaceLinkedObjectForUser) | **Put** /api/v1/users/{userIdOrLogin}/linkedObjects/{primaryRelationshipName}/{primaryUserId} | Replace the Linked Object value for &#x60;primary&#x60;
[**ReplaceUser**](UserAPI.md#ReplaceUser) | **Put** /api/v1/users/{userId} | Replace a User
[**ResetFactors**](UserAPI.md#ResetFactors) | **Post** /api/v1/users/{userId}/lifecycle/reset_factors | Reset all Factors
[**RevokeGrantsForUserAndClient**](UserAPI.md#RevokeGrantsForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/grants | Revoke all Grants for a Client
[**RevokeTokenForUserAndClient**](UserAPI.md#RevokeTokenForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/tokens/{tokenId} | Revoke a Token for a Client
[**RevokeTokensForUserAndClient**](UserAPI.md#RevokeTokensForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/tokens | Revoke all Refresh Tokens for a Client
[**RevokeUserGrant**](UserAPI.md#RevokeUserGrant) | **Delete** /api/v1/users/{userId}/grants/{grantId} | Revoke a User Grant
[**RevokeUserGrants**](UserAPI.md#RevokeUserGrants) | **Delete** /api/v1/users/{userId}/grants | Revoke all User Grants
[**RevokeUserSessions**](UserAPI.md#RevokeUserSessions) | **Delete** /api/v1/users/{userId}/sessions | Revoke all User Sessions
[**SuspendUser**](UserAPI.md#SuspendUser) | **Post** /api/v1/users/{userId}/lifecycle/suspend | Suspend a User
[**UnlockUser**](UserAPI.md#UnlockUser) | **Post** /api/v1/users/{userId}/lifecycle/unlock | Unlock a User
[**UnsuspendUser**](UserAPI.md#UnsuspendUser) | **Post** /api/v1/users/{userId}/lifecycle/unsuspend | Unsuspend a User
[**UpdateUser**](UserAPI.md#UpdateUser) | **Post** /api/v1/users/{userId} | Update a User



## ActivateUser

> UserActivationToken ActivateUser(ctx, userId).SendEmail(sendEmail).Execute()

Activate a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    sendEmail := true // bool | Sends an activation email to the user if true (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ActivateUser(context.Background(), userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ActivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateUser`: UserActivationToken
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ActivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** | Sends an activation email to the user if true | [default to true]

### Return type

[**UserActivationToken**](UserActivationToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePassword

> UserCredentials ChangePassword(ctx, userId).ChangePasswordRequest(changePasswordRequest).Strict(strict).Execute()

Change Password



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
    userId := "userId_example" // string | ID of an existing Okta user
    changePasswordRequest := *openapiclient.NewChangePasswordRequest() // ChangePasswordRequest | 
    strict := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ChangePassword(context.Background(), userId).ChangePasswordRequest(changePasswordRequest).Strict(strict).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangePassword`: UserCredentials
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ChangePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changePasswordRequest** | [**ChangePasswordRequest**](ChangePasswordRequest.md) |  | 
 **strict** | **bool** |  | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRecoveryQuestion

> UserCredentials ChangeRecoveryQuestion(ctx, userId).UserCredentials(userCredentials).Execute()

Change Recovery Question



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
    userId := "userId_example" // string | ID of an existing Okta user
    userCredentials := *openapiclient.NewUserCredentials() // UserCredentials | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ChangeRecoveryQuestion(context.Background(), userId).UserCredentials(userCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangeRecoveryQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeRecoveryQuestion`: UserCredentials
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ChangeRecoveryQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRecoveryQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCredentials** | [**UserCredentials**](UserCredentials.md) |  | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).Body(body).Activate(activate).Provider(provider).NextLogin(nextLogin).Execute()

Create a User



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
    body := *openapiclient.NewCreateUserRequest(*openapiclient.NewUserProfile()) // CreateUserRequest | 
    activate := true // bool | Executes activation lifecycle operation when creating the user (optional) (default to true)
    provider := true // bool | Indicates whether to create a user with a specified authentication provider (optional) (default to false)
    nextLogin := "nextLogin_example" // string | With activate=true, set nextLogin to \"changePassword\" to have the password be EXPIRED, so user must change it the next time they log in. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.CreateUser(context.Background()).Body(body).Activate(activate).Provider(provider).NextLogin(nextLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateUserRequest**](CreateUserRequest.md) |  | 
 **activate** | **bool** | Executes activation lifecycle operation when creating the user | [default to true]
 **provider** | **bool** | Indicates whether to create a user with a specified authentication provider | [default to false]
 **nextLogin** | **string** | With activate&#x3D;true, set nextLogin to \&quot;changePassword\&quot; to have the password be EXPIRED, so user must change it the next time they log in. | 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateUser

> DeactivateUser(ctx, userId).SendEmail(sendEmail).Execute()

Deactivate a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    sendEmail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.DeactivateUser(context.Background(), userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeactivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** |  | [default to false]

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


## DeleteLinkedObjectForUser

> DeleteLinkedObjectForUser(ctx, userIdOrLogin, relationshipName).Execute()

Delete a Linked Object



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
    userIdOrLogin := "userIdOrLogin_example" // string | User ID or login value of the user assigned the `associated` relationship
    relationshipName := "relationshipName_example" // string | Name of the `primary` or `associated` relationship being queried

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.DeleteLinkedObjectForUser(context.Background(), userIdOrLogin, relationshipName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteLinkedObjectForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrLogin** | **string** | User ID or login value of the user assigned the &#x60;associated&#x60; relationship | 
**relationshipName** | **string** | Name of the &#x60;primary&#x60; or &#x60;associated&#x60; relationship being queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkedObjectForUserRequest struct via the builder pattern


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


## DeleteUser

> DeleteUser(ctx, userId).SendEmail(sendEmail).Execute()

Delete a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    sendEmail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.DeleteUser(context.Background(), userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** |  | [default to false]

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


## ExpirePassword

> User ExpirePassword(ctx, userId).Execute()

Expire Password



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ExpirePassword(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ExpirePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirePassword`: User
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ExpirePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirePasswordAndGetTemporaryPassword

> TempPassword ExpirePasswordAndGetTemporaryPassword(ctx, userId).RevokeSessions(revokeSessions).Execute()

Expire Password and Set Temporary Password



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
    userId := "userId_example" // string | ID of an existing Okta user
    revokeSessions := true // bool | When set to `true` (and the session is a user session), all user sessions are revoked except the current session. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ExpirePasswordAndGetTemporaryPassword(context.Background(), userId).RevokeSessions(revokeSessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ExpirePasswordAndGetTemporaryPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirePasswordAndGetTemporaryPassword`: TempPassword
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ExpirePasswordAndGetTemporaryPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePasswordAndGetTemporaryPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokeSessions** | **bool** | When set to &#x60;true&#x60; (and the session is a user session), all user sessions are revoked except the current session. | [default to false]

### Return type

[**TempPassword**](TempPassword.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPassword

> ForgotPasswordResponse ForgotPassword(ctx, userId).SendEmail(sendEmail).Execute()

Initiate Forgot Password



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
    userId := "userId_example" // string | ID of an existing Okta user
    sendEmail := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ForgotPassword(context.Background(), userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ForgotPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForgotPassword`: ForgotPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ForgotPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** |  | [default to true]

### Return type

[**ForgotPasswordResponse**](ForgotPasswordResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPasswordSetNewPassword

> UserCredentials ForgotPasswordSetNewPassword(ctx, userId).UserCredentials(userCredentials).SendEmail(sendEmail).Execute()

Reset Password with Recovery Question



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
    userId := "userId_example" // string | ID of an existing Okta user
    userCredentials := *openapiclient.NewUserCredentials() // UserCredentials | 
    sendEmail := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ForgotPasswordSetNewPassword(context.Background(), userId).UserCredentials(userCredentials).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ForgotPasswordSetNewPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForgotPasswordSetNewPassword`: UserCredentials
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ForgotPasswordSetNewPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordSetNewPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCredentials** | [**UserCredentials**](UserCredentials.md) |  | 
 **sendEmail** | **bool** |  | [default to true]

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateResetPasswordToken

> ResetPasswordToken GenerateResetPasswordToken(ctx, userId).SendEmail(sendEmail).RevokeSessions(revokeSessions).Execute()

Generate a Reset Password Token



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
    userId := "userId_example" // string | ID of an existing Okta user
    sendEmail := true // bool | 
    revokeSessions := true // bool | When set to `true` (and the session is a user session), all user sessions are revoked except the current session. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GenerateResetPasswordToken(context.Background(), userId).SendEmail(sendEmail).RevokeSessions(revokeSessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GenerateResetPasswordToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateResetPasswordToken`: ResetPasswordToken
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GenerateResetPasswordToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateResetPasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** |  | 
 **revokeSessions** | **bool** | When set to &#x60;true&#x60; (and the session is a user session), all user sessions are revoked except the current session. | [default to false]

### Return type

[**ResetPasswordToken**](ResetPasswordToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefreshTokenForUserAndClient

> OAuth2RefreshToken GetRefreshTokenForUserAndClient(ctx, userId, clientId, tokenId).Expand(expand).Limit(limit).After(after).Execute()

Retrieve a Refresh Token for a Client



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
    userId := "userId_example" // string | ID of an existing Okta user
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token
    expand := "expand_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)
    after := "after_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetRefreshTokenForUserAndClient(context.Background(), userId, clientId, tokenId).Expand(expand).Limit(limit).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetRefreshTokenForUserAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefreshTokenForUserAndClient`: OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetRefreshTokenForUserAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefreshTokenForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **after** | **string** |  | 

### Return type

[**OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserGetSingleton GetUser(ctx, userId).Expand(expand).Execute()

Retrieve a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    expand := "blocks" // string | An optional parameter to include metadata in the `_embedded` attribute. Valid value: `blocks` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetUser(context.Background(), userId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserGetSingleton
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | An optional parameter to include metadata in the &#x60;_embedded&#x60; attribute. Valid value: &#x60;blocks&#x60; | 

### Return type

[**UserGetSingleton**](UserGetSingleton.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGrant

> OAuth2ScopeConsentGrant GetUserGrant(ctx, userId, grantId).Expand(expand).Execute()

Retrieve a User Grant



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
    userId := "userId_example" // string | ID of an existing Okta user
    grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetUserGrant(context.Background(), userId, grantId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserGrant`: OAuth2ScopeConsentGrant
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppLinks

> []AppLink ListAppLinks(ctx, userId).Execute()

List all Assigned Application Links



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListAppLinks(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListAppLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppLinks`: []AppLink
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListAppLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppLink**](AppLink.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrantsForUserAndClient

> []OAuth2ScopeConsentGrant ListGrantsForUserAndClient(ctx, userId, clientId).Expand(expand).After(after).Limit(limit).Execute()

List all Grants for a Client



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
    userId := "userId_example" // string | ID of an existing Okta user
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    expand := "expand_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListGrantsForUserAndClient(context.Background(), userId, clientId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListGrantsForUserAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGrantsForUserAndClient`: []OAuth2ScopeConsentGrant
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListGrantsForUserAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLinkedObjectsForUser

> []map[string]interface{} ListLinkedObjectsForUser(ctx, userIdOrLogin, relationshipName).After(after).Limit(limit).Execute()

List the primary or all of the associated Linked Object values



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
    userIdOrLogin := "userIdOrLogin_example" // string | User ID or login value of the user assigned the `associated` relationship
    relationshipName := "relationshipName_example" // string | Name of the `primary` or `associated` relationship being queried
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListLinkedObjectsForUser(context.Background(), userIdOrLogin, relationshipName).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListLinkedObjectsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLinkedObjectsForUser`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListLinkedObjectsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrLogin** | **string** | User ID or login value of the user assigned the &#x60;associated&#x60; relationship | 
**relationshipName** | **string** | Name of the &#x60;primary&#x60; or &#x60;associated&#x60; relationship being queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLinkedObjectsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]

### Return type

**[]map[string]interface{}**

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRefreshTokensForUserAndClient

> []OAuth2RefreshToken ListRefreshTokensForUserAndClient(ctx, userId, clientId).Expand(expand).After(after).Limit(limit).Execute()

List all Refresh Tokens for a Client



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
    userId := "userId_example" // string | ID of an existing Okta user
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    expand := "expand_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListRefreshTokensForUserAndClient(context.Background(), userId, clientId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListRefreshTokensForUserAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRefreshTokensForUserAndClient`: []OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListRefreshTokensForUserAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRefreshTokensForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBlocks

> []UserBlock ListUserBlocks(ctx, userId).Execute()

List all User Blocks



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUserBlocks(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserBlocks`: []UserBlock
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserBlock**](UserBlock.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserClients

> []OAuth2Client ListUserClients(ctx, userId).Execute()

List all Clients



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUserClients(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserClients`: []OAuth2Client
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGrants

> []OAuth2ScopeConsentGrant ListUserGrants(ctx, userId).ScopeId(scopeId).Expand(expand).After(after).Limit(limit).Execute()

List all User Grants



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
    userId := "userId_example" // string | ID of an existing Okta user
    scopeId := "scopeId_example" // string |  (optional)
    expand := "expand_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUserGrants(context.Background(), userId).ScopeId(scopeId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGrants`: []OAuth2ScopeConsentGrant
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeId** | **string** |  | 
 **expand** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroups

> []Group ListUserGroups(ctx, userId).After(after).Limit(limit).Execute()

List all Groups



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
    userId := "userId_example" // string | ID of an existing Okta user
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUserGroups(context.Background(), userId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGroups`: []Group
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**[]Group**](Group.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserIdentityProviders

> []IdentityProvider ListUserIdentityProviders(ctx, userId).Execute()

List all Identity Providers



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUserIdentityProviders(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserIdentityProviders`: []IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserIdentityProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx).Q(q).After(after).Limit(limit).Filter(filter).Search(search).SortBy(sortBy).SortOrder(sortOrder).Execute()

List all Users



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
    q := "q_example" // string | Finds a user that matches firstName, lastName, and email properties (optional)
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | Specifies the number of results returned. Defaults to 10 if `q` is provided. (optional) (default to 200)
    filter := "filter_example" // string | Filters users with a supported expression for a subset of properties (optional)
    search := "search_example" // string | Searches for users with a supported filtering expression for most properties. Okta recommends using this parameter for search for best performance. (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string | Sorting is done in ASCII sort order (that is, by ASCII character value), but isn't case sensitive. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUsers(context.Background()).Q(q).After(after).Limit(limit).Filter(filter).Search(search).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Finds a user that matches firstName, lastName, and email properties | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | Specifies the number of results returned. Defaults to 10 if &#x60;q&#x60; is provided. | [default to 200]
 **filter** | **string** | Filters users with a supported expression for a subset of properties | 
 **search** | **string** | Searches for users with a supported filtering expression for most properties. Okta recommends using this parameter for search for best performance. | 
 **sortBy** | **string** |  | 
 **sortOrder** | **string** | Sorting is done in ASCII sort order (that is, by ASCII character value), but isn&#39;t case sensitive. | 

### Return type

[**[]User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateUser

> UserActivationToken ReactivateUser(ctx, userId).SendEmail(sendEmail).Execute()

Reactivate a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    sendEmail := true // bool | Sends an activation email to the user if true (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ReactivateUser(context.Background(), userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ReactivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactivateUser`: UserActivationToken
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ReactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** | Sends an activation email to the user if true | [default to false]

### Return type

[**UserActivationToken**](UserActivationToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLinkedObjectForUser

> ReplaceLinkedObjectForUser(ctx, userIdOrLogin, primaryRelationshipName, primaryUserId).Execute()

Replace the Linked Object value for `primary`



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
    userIdOrLogin := "userIdOrLogin_example" // string | User ID or login value of the user assigned the `associated` relationship
    primaryRelationshipName := "primaryRelationshipName_example" // string | Name of the `primary` relationship being assigned
    primaryUserId := "ctxeQ5JnAVdGFBB7Zr7W" // string | User ID to be assigned to the `primary` relationship for the `associated` user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.ReplaceLinkedObjectForUser(context.Background(), userIdOrLogin, primaryRelationshipName, primaryUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ReplaceLinkedObjectForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrLogin** | **string** | User ID or login value of the user assigned the &#x60;associated&#x60; relationship | 
**primaryRelationshipName** | **string** | Name of the &#x60;primary&#x60; relationship being assigned | 
**primaryUserId** | **string** | User ID to be assigned to the &#x60;primary&#x60; relationship for the &#x60;associated&#x60; user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLinkedObjectForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUser

> User ReplaceUser(ctx, userId).User(user).Strict(strict).Execute()

Replace a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    user := *openapiclient.NewUser() // User | 
    strict := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ReplaceUser(context.Background(), userId).User(user).Strict(strict).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ReplaceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ReplaceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 
 **strict** | **bool** |  | 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetFactors

> ResetFactors(ctx, userId).RemoveRecoveryEnrollment(removeRecoveryEnrollment).Execute()

Reset all Factors



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
    userId := "userId_example" // string | ID of an existing Okta user
    removeRecoveryEnrollment := true // bool | If `true`, removes the phone number as both a recovery method and a Factor. Supported Factors: `sms` and `call` (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.ResetFactors(context.Background(), userId).RemoveRecoveryEnrollment(removeRecoveryEnrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ResetFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeRecoveryEnrollment** | **bool** | If &#x60;true&#x60;, removes the phone number as both a recovery method and a Factor. Supported Factors: &#x60;sms&#x60; and &#x60;call&#x60; | [default to false]

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


## RevokeGrantsForUserAndClient

> RevokeGrantsForUserAndClient(ctx, userId, clientId).Execute()

Revoke all Grants for a Client



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
    userId := "userId_example" // string | ID of an existing Okta user
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.RevokeGrantsForUserAndClient(context.Background(), userId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RevokeGrantsForUserAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeGrantsForUserAndClientRequest struct via the builder pattern


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


## RevokeTokenForUserAndClient

> RevokeTokenForUserAndClient(ctx, userId, clientId, tokenId).Execute()

Revoke a Token for a Client



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
    userId := "userId_example" // string | ID of an existing Okta user
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.RevokeTokenForUserAndClient(context.Background(), userId, clientId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RevokeTokenForUserAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenForUserAndClientRequest struct via the builder pattern


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


## RevokeTokensForUserAndClient

> RevokeTokensForUserAndClient(ctx, userId, clientId).Execute()

Revoke all Refresh Tokens for a Client



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
    userId := "userId_example" // string | ID of an existing Okta user
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.RevokeTokensForUserAndClient(context.Background(), userId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RevokeTokensForUserAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokensForUserAndClientRequest struct via the builder pattern


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


## RevokeUserGrant

> RevokeUserGrant(ctx, userId, grantId).Execute()

Revoke a User Grant



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
    userId := "userId_example" // string | ID of an existing Okta user
    grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.RevokeUserGrant(context.Background(), userId, grantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RevokeUserGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserGrantRequest struct via the builder pattern


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


## RevokeUserGrants

> RevokeUserGrants(ctx, userId).Execute()

Revoke all User Grants



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.RevokeUserGrants(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RevokeUserGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserGrantsRequest struct via the builder pattern


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


## RevokeUserSessions

> RevokeUserSessions(ctx, userId).OauthTokens(oauthTokens).Execute()

Revoke all User Sessions



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
    userId := "userId_example" // string | ID of an existing Okta user
    oauthTokens := true // bool | Revoke issued OpenID Connect and OAuth refresh and access tokens (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.RevokeUserSessions(context.Background(), userId).OauthTokens(oauthTokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RevokeUserSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oauthTokens** | **bool** | Revoke issued OpenID Connect and OAuth refresh and access tokens | [default to false]

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


## SuspendUser

> SuspendUser(ctx, userId).Execute()

Suspend a User



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.SuspendUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SuspendUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendUserRequest struct via the builder pattern


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


## UnlockUser

> UnlockUser(ctx, userId).Execute()

Unlock a User



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.UnlockUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnlockUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockUserRequest struct via the builder pattern


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


## UnsuspendUser

> UnsuspendUser(ctx, userId).Execute()

Unsuspend a User



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.UnsuspendUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnsuspendUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsuspendUserRequest struct via the builder pattern


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


## UpdateUser

> User UpdateUser(ctx, userId).User(user).Strict(strict).Execute()

Update a User



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
    userId := "userId_example" // string | ID of an existing Okta user
    user := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 
    strict := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UpdateUser(context.Background(), userId).User(user).Strict(strict).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 
 **strict** | **bool** |  | 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

