# \UserCredAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](UserCredAPI.md#ChangePassword) | **Post** /api/v1/users/{userId}/credentials/change_password | Update password
[**ChangeRecoveryQuestion**](UserCredAPI.md#ChangeRecoveryQuestion) | **Post** /api/v1/users/{userId}/credentials/change_recovery_question | Update recovery question
[**ExpirePassword**](UserCredAPI.md#ExpirePassword) | **Post** /api/v1/users/{id}/lifecycle/expire_password | Expire the password
[**ExpirePasswordWithTempPassword**](UserCredAPI.md#ExpirePasswordWithTempPassword) | **Post** /api/v1/users/{id}/lifecycle/expire_password_with_temp_password | Expire the password with a temporary password
[**ForgotPassword**](UserCredAPI.md#ForgotPassword) | **Post** /api/v1/users/{userId}/credentials/forgot_password | Start forgot password flow
[**ForgotPasswordSetNewPassword**](UserCredAPI.md#ForgotPasswordSetNewPassword) | **Post** /api/v1/users/{userId}/credentials/forgot_password_recovery_question | Reset password with recovery question
[**ResetPassword**](UserCredAPI.md#ResetPassword) | **Post** /api/v1/users/{id}/lifecycle/reset_password | Reset a password



## ChangePassword

> UserCredentials ChangePassword(ctx, userId).ChangePasswordRequest(changePasswordRequest).Strict(strict).Execute()

Update password



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	changePasswordRequest := *openapiclient.NewChangePasswordRequest() // ChangePasswordRequest | 
	strict := true // bool | If true, validates against the password minimum age policy (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ChangePassword(context.Background(), userId).ChangePasswordRequest(changePasswordRequest).Strict(strict).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePassword`: UserCredentials
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ChangePassword`: %v\n", resp)
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
 **strict** | **bool** | If true, validates against the password minimum age policy | [default to false]

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

Update recovery question



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	userCredentials := *openapiclient.NewUserCredentials() // UserCredentials | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ChangeRecoveryQuestion(context.Background(), userId).UserCredentials(userCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ChangeRecoveryQuestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeRecoveryQuestion`: UserCredentials
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ChangeRecoveryQuestion`: %v\n", resp)
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


## ExpirePassword

> User ExpirePassword(ctx, id).Execute()

Expire the password



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
	id := "id_example" // string | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ExpirePassword(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ExpirePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpirePassword`: User
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ExpirePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

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


## ExpirePasswordWithTempPassword

> User ExpirePasswordWithTempPassword(ctx, id).RevokeSessions(revokeSessions).Execute()

Expire the password with a temporary password



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
	id := "id_example" // string | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	revokeSessions := true // bool | Revokes the user's existing sessions if `true` (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ExpirePasswordWithTempPassword(context.Background(), id).RevokeSessions(revokeSessions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ExpirePasswordWithTempPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpirePasswordWithTempPassword`: User
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ExpirePasswordWithTempPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePasswordWithTempPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokeSessions** | **bool** | Revokes the user&#39;s existing sessions if &#x60;true&#x60; | [default to false]

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


## ForgotPassword

> ForgotPasswordResponse ForgotPassword(ctx, userId).SendEmail(sendEmail).Execute()

Start forgot password flow



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	sendEmail := true // bool | Sends a forgot password email to the user if `true` (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ForgotPassword(context.Background(), userId).SendEmail(sendEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPassword`: ForgotPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ForgotPassword`: %v\n", resp)
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

 **sendEmail** | **bool** | Sends a forgot password email to the user if &#x60;true&#x60; | [default to true]

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

Reset password with recovery question



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	userCredentials := *openapiclient.NewUserCredentials() // UserCredentials | 
	sendEmail := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ForgotPasswordSetNewPassword(context.Background(), userId).UserCredentials(userCredentials).SendEmail(sendEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ForgotPasswordSetNewPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPasswordSetNewPassword`: UserCredentials
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ForgotPasswordSetNewPassword`: %v\n", resp)
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


## ResetPassword

> ResetPasswordToken ResetPassword(ctx, id).SendEmail(sendEmail).RevokeSessions(revokeSessions).Execute()

Reset a password



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
	id := "id_example" // string | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	sendEmail := true // bool | 
	revokeSessions := true // bool | Revokes all user sessions, except for the current session, if set to `true` (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCredAPI.ResetPassword(context.Background(), id).SendEmail(sendEmail).RevokeSessions(revokeSessions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCredAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetPassword`: ResetPasswordToken
	fmt.Fprintf(os.Stdout, "Response from `UserCredAPI.ResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** |  | 
 **revokeSessions** | **bool** | Revokes all user sessions, except for the current session, if set to &#x60;true&#x60; | [default to false]

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

