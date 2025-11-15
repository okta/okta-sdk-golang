# \UserAuthenticatorEnrollmentsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticatorEnrollment**](UserAuthenticatorEnrollmentsAPI.md#CreateAuthenticatorEnrollment) | **Post** /api/v1/users/{userId}/authenticator-enrollments/phone | Create an auto-activated Phone authenticator enrollment
[**CreateTacAuthenticatorEnrollment**](UserAuthenticatorEnrollmentsAPI.md#CreateTacAuthenticatorEnrollment) | **Post** /api/v1/users/{userId}/authenticator-enrollments/tac | Create an auto-activated TAC authenticator enrollment
[**DeleteAuthenticatorEnrollment**](UserAuthenticatorEnrollmentsAPI.md#DeleteAuthenticatorEnrollment) | **Delete** /api/v1/users/{userId}/authenticator-enrollments/{enrollmentId} | Delete an authenticator enrollment
[**GetAuthenticatorEnrollment**](UserAuthenticatorEnrollmentsAPI.md#GetAuthenticatorEnrollment) | **Get** /api/v1/users/{userId}/authenticator-enrollments/{enrollmentId} | Retrieve an authenticator enrollment
[**ListAuthenticatorEnrollments**](UserAuthenticatorEnrollmentsAPI.md#ListAuthenticatorEnrollments) | **Get** /api/v1/users/{userId}/authenticator-enrollments | List all authenticator enrollments



## CreateAuthenticatorEnrollment

> AuthenticatorEnrollment CreateAuthenticatorEnrollment(ctx, userId).Authenticator(authenticator).Execute()

Create an auto-activated Phone authenticator enrollment



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
	authenticator := *openapiclient.NewAuthenticatorEnrollmentCreateRequest("AuthenticatorId_example", *openapiclient.NewAuthenticatorProfile("PhoneNumber_example")) // AuthenticatorEnrollmentCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAuthenticatorEnrollmentsAPI.CreateAuthenticatorEnrollment(context.Background(), userId).Authenticator(authenticator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticatorEnrollmentsAPI.CreateAuthenticatorEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthenticatorEnrollment`: AuthenticatorEnrollment
	fmt.Fprintf(os.Stdout, "Response from `UserAuthenticatorEnrollmentsAPI.CreateAuthenticatorEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticatorEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticator** | [**AuthenticatorEnrollmentCreateRequest**](AuthenticatorEnrollmentCreateRequest.md) |  | 

### Return type

[**AuthenticatorEnrollment**](AuthenticatorEnrollment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTacAuthenticatorEnrollment

> TacAuthenticatorEnrollment CreateTacAuthenticatorEnrollment(ctx, userId).Authenticator(authenticator).Execute()

Create an auto-activated TAC authenticator enrollment



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
	authenticator := *openapiclient.NewAuthenticatorEnrollmentCreateRequestTac("AuthenticatorId_example") // AuthenticatorEnrollmentCreateRequestTac | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAuthenticatorEnrollmentsAPI.CreateTacAuthenticatorEnrollment(context.Background(), userId).Authenticator(authenticator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticatorEnrollmentsAPI.CreateTacAuthenticatorEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTacAuthenticatorEnrollment`: TacAuthenticatorEnrollment
	fmt.Fprintf(os.Stdout, "Response from `UserAuthenticatorEnrollmentsAPI.CreateTacAuthenticatorEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTacAuthenticatorEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticator** | [**AuthenticatorEnrollmentCreateRequestTac**](AuthenticatorEnrollmentCreateRequestTac.md) |  | 

### Return type

[**TacAuthenticatorEnrollment**](TacAuthenticatorEnrollment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticatorEnrollment

> DeleteAuthenticatorEnrollment(ctx, userId, enrollmentId).Execute()

Delete an authenticator enrollment



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
	enrollmentId := "sms8lqwuzSpWT4kVs0g4" // string | Unique identifier of an enrollment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAuthenticatorEnrollmentsAPI.DeleteAuthenticatorEnrollment(context.Background(), userId, enrollmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticatorEnrollmentsAPI.DeleteAuthenticatorEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**enrollmentId** | **string** | Unique identifier of an enrollment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticatorEnrollmentRequest struct via the builder pattern


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


## GetAuthenticatorEnrollment

> AuthenticatorEnrollment GetAuthenticatorEnrollment(ctx, userId, enrollmentId).DiscloseIdentifiers(discloseIdentifiers).Execute()

Retrieve an authenticator enrollment



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
	enrollmentId := "sms8lqwuzSpWT4kVs0g4" // string | Unique identifier of an enrollment
	discloseIdentifiers := []string{"DiscloseIdentifiers_example"} // []string | Indicates whether or not the identifier of an authenticator enrollment is disclosed or anonymized. If it's included in the operation query, then the identifier of the authenticator enrollment (the actual phone number, for example) is included in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAuthenticatorEnrollmentsAPI.GetAuthenticatorEnrollment(context.Background(), userId, enrollmentId).DiscloseIdentifiers(discloseIdentifiers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticatorEnrollmentsAPI.GetAuthenticatorEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticatorEnrollment`: AuthenticatorEnrollment
	fmt.Fprintf(os.Stdout, "Response from `UserAuthenticatorEnrollmentsAPI.GetAuthenticatorEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**enrollmentId** | **string** | Unique identifier of an enrollment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatorEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **discloseIdentifiers** | **[]string** | Indicates whether or not the identifier of an authenticator enrollment is disclosed or anonymized. If it&#39;s included in the operation query, then the identifier of the authenticator enrollment (the actual phone number, for example) is included in the response. | 

### Return type

[**AuthenticatorEnrollment**](AuthenticatorEnrollment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthenticatorEnrollments

> AuthenticatorEnrollment ListAuthenticatorEnrollments(ctx, userId).DiscloseIdentifiers(discloseIdentifiers).Execute()

List all authenticator enrollments



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
	discloseIdentifiers := []string{"DiscloseIdentifiers_example"} // []string | Indicates whether or not the identifier of an authenticator enrollment is disclosed or anonymized. If it's included in the operation query, then the identifier of the authenticator enrollment (the actual phone number, for example) is included in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAuthenticatorEnrollmentsAPI.ListAuthenticatorEnrollments(context.Background(), userId).DiscloseIdentifiers(discloseIdentifiers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticatorEnrollmentsAPI.ListAuthenticatorEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthenticatorEnrollments`: AuthenticatorEnrollment
	fmt.Fprintf(os.Stdout, "Response from `UserAuthenticatorEnrollmentsAPI.ListAuthenticatorEnrollments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthenticatorEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **discloseIdentifiers** | **[]string** | Indicates whether or not the identifier of an authenticator enrollment is disclosed or anonymized. If it&#39;s included in the operation query, then the identifier of the authenticator enrollment (the actual phone number, for example) is included in the response. | 

### Return type

[**AuthenticatorEnrollment**](AuthenticatorEnrollment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

