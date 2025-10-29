# \UserLifecycleAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](UserLifecycleAPI.md#ActivateUser) | **Post** /api/v1/users/{id}/lifecycle/activate | Activate a user
[**DeactivateUser**](UserLifecycleAPI.md#DeactivateUser) | **Post** /api/v1/users/{id}/lifecycle/deactivate | Deactivate a user
[**ReactivateUser**](UserLifecycleAPI.md#ReactivateUser) | **Post** /api/v1/users/{id}/lifecycle/reactivate | Reactivate a user
[**ResetFactors**](UserLifecycleAPI.md#ResetFactors) | **Post** /api/v1/users/{id}/lifecycle/reset_factors | Reset the factors
[**SuspendUser**](UserLifecycleAPI.md#SuspendUser) | **Post** /api/v1/users/{id}/lifecycle/suspend | Suspend a user
[**UnlockUser**](UserLifecycleAPI.md#UnlockUser) | **Post** /api/v1/users/{id}/lifecycle/unlock | Unlock a user
[**UnsuspendUser**](UserLifecycleAPI.md#UnsuspendUser) | **Post** /api/v1/users/{id}/lifecycle/unsuspend | Unsuspend a user



## ActivateUser

> UserActivationToken ActivateUser(ctx, id).SendEmail(sendEmail).Execute()

Activate a user



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
	sendEmail := true // bool | Sends an activation email to the user if `true` (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLifecycleAPI.ActivateUser(context.Background(), id).SendEmail(sendEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.ActivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateUser`: UserActivationToken
	fmt.Fprintf(os.Stdout, "Response from `UserLifecycleAPI.ActivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** | Sends an activation email to the user if &#x60;true&#x60; | [default to true]

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


## DeactivateUser

> DeactivateUser(ctx, id).SendEmail(sendEmail).Prefer(prefer).Execute()

Deactivate a user



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
	sendEmail := true // bool | Sends a deactivation email to the admin if `true` (optional) (default to false)
	prefer := "prefer_example" // string | Request asynchronous processing (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserLifecycleAPI.DeactivateUser(context.Background(), id).SendEmail(sendEmail).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.DeactivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** | Sends a deactivation email to the admin if &#x60;true&#x60; | [default to false]
 **prefer** | **string** | Request asynchronous processing | 

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


## ReactivateUser

> UserActivationToken ReactivateUser(ctx, id).SendEmail(sendEmail).Execute()

Reactivate a user



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
	sendEmail := true // bool | Sends an activation email to the user if `true` (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLifecycleAPI.ReactivateUser(context.Background(), id).SendEmail(sendEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.ReactivateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactivateUser`: UserActivationToken
	fmt.Fprintf(os.Stdout, "Response from `UserLifecycleAPI.ReactivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** | Sends an activation email to the user if &#x60;true&#x60; | [default to false]

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


## ResetFactors

> ResetFactors(ctx, id).Execute()

Reset the factors



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
	r, err := apiClient.UserLifecycleAPI.ResetFactors(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.ResetFactors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetFactorsRequest struct via the builder pattern


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


## SuspendUser

> SuspendUser(ctx, id).Execute()

Suspend a user



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
	r, err := apiClient.UserLifecycleAPI.SuspendUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.SuspendUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

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

> UnlockUser(ctx, id).Execute()

Unlock a user



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
	r, err := apiClient.UserLifecycleAPI.UnlockUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.UnlockUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

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

> UnsuspendUser(ctx, id).Execute()

Unsuspend a user



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
	r, err := apiClient.UserLifecycleAPI.UnsuspendUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLifecycleAPI.UnsuspendUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

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

