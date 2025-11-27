# \AuthenticatorAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthenticator**](AuthenticatorAPI.md#ActivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/activate | Activate an authenticator
[**ActivateAuthenticatorMethod**](AuthenticatorAPI.md#ActivateAuthenticatorMethod) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/activate | Activate an authenticator method
[**CreateAuthenticator**](AuthenticatorAPI.md#CreateAuthenticator) | **Post** /api/v1/authenticators | Create an authenticator
[**CreateCustomAAGUID**](AuthenticatorAPI.md#CreateCustomAAGUID) | **Post** /api/v1/authenticators/{authenticatorId}/aaguids | Create a custom AAGUID
[**DeactivateAuthenticator**](AuthenticatorAPI.md#DeactivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/deactivate | Deactivate an authenticator
[**DeactivateAuthenticatorMethod**](AuthenticatorAPI.md#DeactivateAuthenticatorMethod) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/deactivate | Deactivate an authenticator method
[**DeleteCustomAAGUID**](AuthenticatorAPI.md#DeleteCustomAAGUID) | **Delete** /api/v1/authenticators/{authenticatorId}/aaguids/{aaguid} | Delete a custom AAGUID
[**GetAuthenticator**](AuthenticatorAPI.md#GetAuthenticator) | **Get** /api/v1/authenticators/{authenticatorId} | Retrieve an authenticator
[**GetAuthenticatorMethod**](AuthenticatorAPI.md#GetAuthenticatorMethod) | **Get** /api/v1/authenticators/{authenticatorId}/methods/{methodType} | Retrieve an authenticator method
[**GetCustomAAGUID**](AuthenticatorAPI.md#GetCustomAAGUID) | **Get** /api/v1/authenticators/{authenticatorId}/aaguids/{aaguid} | Retrieve a custom AAGUID
[**GetWellKnownAppAuthenticatorConfiguration**](AuthenticatorAPI.md#GetWellKnownAppAuthenticatorConfiguration) | **Get** /.well-known/app-authenticator-configuration | Retrieve the well-known app authenticator configuration
[**ListAllCustomAAGUIDs**](AuthenticatorAPI.md#ListAllCustomAAGUIDs) | **Get** /api/v1/authenticators/{authenticatorId}/aaguids | List all custom AAGUIDs
[**ListAuthenticatorMethods**](AuthenticatorAPI.md#ListAuthenticatorMethods) | **Get** /api/v1/authenticators/{authenticatorId}/methods | List all methods of an authenticator
[**ListAuthenticators**](AuthenticatorAPI.md#ListAuthenticators) | **Get** /api/v1/authenticators | List all authenticators
[**ReplaceAuthenticator**](AuthenticatorAPI.md#ReplaceAuthenticator) | **Put** /api/v1/authenticators/{authenticatorId} | Replace an authenticator
[**ReplaceAuthenticatorMethod**](AuthenticatorAPI.md#ReplaceAuthenticatorMethod) | **Put** /api/v1/authenticators/{authenticatorId}/methods/{methodType} | Replace an authenticator method
[**ReplaceCustomAAGUID**](AuthenticatorAPI.md#ReplaceCustomAAGUID) | **Put** /api/v1/authenticators/{authenticatorId}/aaguids/{aaguid} | Replace a custom AAGUID
[**UpdateCustomAAGUID**](AuthenticatorAPI.md#UpdateCustomAAGUID) | **Patch** /api/v1/authenticators/{authenticatorId}/aaguids/{aaguid} | Update a custom AAGUID
[**VerifyRpIdDomain**](AuthenticatorAPI.md#VerifyRpIdDomain) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{webAuthnMethodType}/verify-rp-id-domain | Verify a Relying Party ID domain



## ActivateAuthenticator

> AuthenticatorBase ActivateAuthenticator(ctx, authenticatorId).Execute()

Activate an authenticator



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ActivateAuthenticator(context.Background(), authenticatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ActivateAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateAuthenticator`: AuthenticatorBase
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ActivateAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorBase**](AuthenticatorBase.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivateAuthenticatorMethod

> ListAuthenticatorMethods200ResponseInner ActivateAuthenticatorMethod(ctx, authenticatorId, methodType).Execute()

Activate an authenticator method



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	methodType := "methodType_example" // string | Type of authenticator method

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ActivateAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ActivateAuthenticatorMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ActivateAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**methodType** | **string** | Type of authenticator method | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAuthenticatorMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAuthenticatorMethods200ResponseInner**](ListAuthenticatorMethods200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthenticator

> AuthenticatorBase CreateAuthenticator(ctx).Authenticator(authenticator).Activate(activate).Execute()

Create an authenticator



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
	authenticator := *openapiclient.NewAuthenticatorBase() // AuthenticatorBase | 
	activate := true // bool | Whether to execute the activation lifecycle operation when Okta creates the authenticator (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.CreateAuthenticator(context.Background()).Authenticator(authenticator).Activate(activate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.CreateAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthenticator`: AuthenticatorBase
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.CreateAuthenticator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticator** | [**AuthenticatorBase**](AuthenticatorBase.md) |  | 
 **activate** | **bool** | Whether to execute the activation lifecycle operation when Okta creates the authenticator | [default to true]

### Return type

[**AuthenticatorBase**](AuthenticatorBase.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomAAGUID

> CustomAAGUIDResponseObject CreateCustomAAGUID(ctx, authenticatorId).CustomAAGUIDCreateRequestObject(customAAGUIDCreateRequestObject).Execute()

Create a custom AAGUID



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	customAAGUIDCreateRequestObject := *openapiclient.NewCustomAAGUIDCreateRequestObject() // CustomAAGUIDCreateRequestObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.CreateCustomAAGUID(context.Background(), authenticatorId).CustomAAGUIDCreateRequestObject(customAAGUIDCreateRequestObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.CreateCustomAAGUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomAAGUID`: CustomAAGUIDResponseObject
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.CreateCustomAAGUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomAAGUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAAGUIDCreateRequestObject** | [**CustomAAGUIDCreateRequestObject**](CustomAAGUIDCreateRequestObject.md) |  | 

### Return type

[**CustomAAGUIDResponseObject**](CustomAAGUIDResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAuthenticator

> AuthenticatorBase DeactivateAuthenticator(ctx, authenticatorId).Execute()

Deactivate an authenticator



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.DeactivateAuthenticator(context.Background(), authenticatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.DeactivateAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateAuthenticator`: AuthenticatorBase
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.DeactivateAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorBase**](AuthenticatorBase.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAuthenticatorMethod

> ListAuthenticatorMethods200ResponseInner DeactivateAuthenticatorMethod(ctx, authenticatorId, methodType).Execute()

Deactivate an authenticator method



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	methodType := "methodType_example" // string | Type of authenticator method

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.DeactivateAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.DeactivateAuthenticatorMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.DeactivateAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**methodType** | **string** | Type of authenticator method | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAuthenticatorMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAuthenticatorMethods200ResponseInner**](ListAuthenticatorMethods200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomAAGUID

> DeleteCustomAAGUID(ctx, authenticatorId, aaguid).Execute()

Delete a custom AAGUID



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	aaguid := "cb69481e-8ff7-4039-93ec-0a272911111" // string | Unique ID of a custom AAGUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticatorAPI.DeleteCustomAAGUID(context.Background(), authenticatorId, aaguid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.DeleteCustomAAGUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**aaguid** | **string** | Unique ID of a custom AAGUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomAAGUIDRequest struct via the builder pattern


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


## GetAuthenticator

> AuthenticatorBase GetAuthenticator(ctx, authenticatorId).Execute()

Retrieve an authenticator



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.GetAuthenticator(context.Background(), authenticatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.GetAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticator`: AuthenticatorBase
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.GetAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorBase**](AuthenticatorBase.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticatorMethod

> ListAuthenticatorMethods200ResponseInner GetAuthenticatorMethod(ctx, authenticatorId, methodType).Execute()

Retrieve an authenticator method



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	methodType := "methodType_example" // string | Type of authenticator method

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.GetAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.GetAuthenticatorMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.GetAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**methodType** | **string** | Type of authenticator method | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatorMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAuthenticatorMethods200ResponseInner**](ListAuthenticatorMethods200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomAAGUID

> CustomAAGUIDResponseObject GetCustomAAGUID(ctx, authenticatorId, aaguid).Execute()

Retrieve a custom AAGUID



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	aaguid := "cb69481e-8ff7-4039-93ec-0a272911111" // string | Unique ID of a custom AAGUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.GetCustomAAGUID(context.Background(), authenticatorId, aaguid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.GetCustomAAGUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomAAGUID`: CustomAAGUIDResponseObject
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.GetCustomAAGUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**aaguid** | **string** | Unique ID of a custom AAGUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAAGUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomAAGUIDResponseObject**](CustomAAGUIDResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellKnownAppAuthenticatorConfiguration

> []WellKnownAppAuthenticatorConfiguration GetWellKnownAppAuthenticatorConfiguration(ctx).OauthClientId(oauthClientId).Execute()

Retrieve the well-known app authenticator configuration



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
	oauthClientId := "oauthClientId_example" // string | Filters app authenticator configurations by `oauthClientId`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.GetWellKnownAppAuthenticatorConfiguration(context.Background()).OauthClientId(oauthClientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.GetWellKnownAppAuthenticatorConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWellKnownAppAuthenticatorConfiguration`: []WellKnownAppAuthenticatorConfiguration
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.GetWellKnownAppAuthenticatorConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWellKnownAppAuthenticatorConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthClientId** | **string** | Filters app authenticator configurations by &#x60;oauthClientId&#x60; | 

### Return type

[**[]WellKnownAppAuthenticatorConfiguration**](WellKnownAppAuthenticatorConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllCustomAAGUIDs

> []CustomAAGUIDResponseObject ListAllCustomAAGUIDs(ctx, authenticatorId).Execute()

List all custom AAGUIDs



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ListAllCustomAAGUIDs(context.Background(), authenticatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ListAllCustomAAGUIDs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllCustomAAGUIDs`: []CustomAAGUIDResponseObject
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ListAllCustomAAGUIDs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllCustomAAGUIDsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomAAGUIDResponseObject**](CustomAAGUIDResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthenticatorMethods

> []ListAuthenticatorMethods200ResponseInner ListAuthenticatorMethods(ctx, authenticatorId).Execute()

List all methods of an authenticator



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ListAuthenticatorMethods(context.Background(), authenticatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ListAuthenticatorMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthenticatorMethods`: []ListAuthenticatorMethods200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ListAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthenticatorMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListAuthenticatorMethods200ResponseInner**](ListAuthenticatorMethods200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthenticators

> []ListAuthenticators200ResponseInner ListAuthenticators(ctx).Execute()

List all authenticators



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
	resp, r, err := apiClient.AuthenticatorAPI.ListAuthenticators(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ListAuthenticators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthenticators`: []ListAuthenticators200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ListAuthenticators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthenticatorsRequest struct via the builder pattern


### Return type

[**[]ListAuthenticators200ResponseInner**](ListAuthenticators200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAuthenticator

> AuthenticatorBase ReplaceAuthenticator(ctx, authenticatorId).Authenticator(authenticator).Execute()

Replace an authenticator



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	authenticator := *openapiclient.NewAuthenticatorBase() // AuthenticatorBase | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ReplaceAuthenticator(context.Background(), authenticatorId).Authenticator(authenticator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ReplaceAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceAuthenticator`: AuthenticatorBase
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ReplaceAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticator** | [**AuthenticatorBase**](AuthenticatorBase.md) |  | 

### Return type

[**AuthenticatorBase**](AuthenticatorBase.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAuthenticatorMethod

> ListAuthenticatorMethods200ResponseInner ReplaceAuthenticatorMethod(ctx, authenticatorId, methodType).ListAuthenticatorMethods200ResponseInner(listAuthenticatorMethods200ResponseInner).Execute()

Replace an authenticator method



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	methodType := "methodType_example" // string | Type of authenticator method
	listAuthenticatorMethods200ResponseInner := openapiclient.listAuthenticatorMethods_200_response_inner{AuthenticatorMethodOtp: openapiclient.NewAuthenticatorMethodOtp()} // ListAuthenticatorMethods200ResponseInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ReplaceAuthenticatorMethod(context.Background(), authenticatorId, methodType).ListAuthenticatorMethods200ResponseInner(listAuthenticatorMethods200ResponseInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ReplaceAuthenticatorMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ReplaceAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**methodType** | **string** | Type of authenticator method | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthenticatorMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **listAuthenticatorMethods200ResponseInner** | [**ListAuthenticatorMethods200ResponseInner**](ListAuthenticatorMethods200ResponseInner.md) |  | 

### Return type

[**ListAuthenticatorMethods200ResponseInner**](ListAuthenticatorMethods200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCustomAAGUID

> CustomAAGUIDResponseObject ReplaceCustomAAGUID(ctx, authenticatorId, aaguid).CustomAAGUIDUpdateRequestObject(customAAGUIDUpdateRequestObject).Execute()

Replace a custom AAGUID



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	aaguid := "cb69481e-8ff7-4039-93ec-0a272911111" // string | Unique ID of a custom AAGUID
	customAAGUIDUpdateRequestObject := *openapiclient.NewCustomAAGUIDUpdateRequestObject() // CustomAAGUIDUpdateRequestObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.ReplaceCustomAAGUID(context.Background(), authenticatorId, aaguid).CustomAAGUIDUpdateRequestObject(customAAGUIDUpdateRequestObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.ReplaceCustomAAGUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCustomAAGUID`: CustomAAGUIDResponseObject
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.ReplaceCustomAAGUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**aaguid** | **string** | Unique ID of a custom AAGUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCustomAAGUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customAAGUIDUpdateRequestObject** | [**CustomAAGUIDUpdateRequestObject**](CustomAAGUIDUpdateRequestObject.md) |  | 

### Return type

[**CustomAAGUIDResponseObject**](CustomAAGUIDResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomAAGUID

> CustomAAGUIDResponseObject UpdateCustomAAGUID(ctx, authenticatorId, aaguid).CustomAAGUIDUpdateRequestObject(customAAGUIDUpdateRequestObject).Execute()

Update a custom AAGUID



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	aaguid := "cb69481e-8ff7-4039-93ec-0a272911111" // string | Unique ID of a custom AAGUID
	customAAGUIDUpdateRequestObject := *openapiclient.NewCustomAAGUIDUpdateRequestObject() // CustomAAGUIDUpdateRequestObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticatorAPI.UpdateCustomAAGUID(context.Background(), authenticatorId, aaguid).CustomAAGUIDUpdateRequestObject(customAAGUIDUpdateRequestObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.UpdateCustomAAGUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomAAGUID`: CustomAAGUIDResponseObject
	fmt.Fprintf(os.Stdout, "Response from `AuthenticatorAPI.UpdateCustomAAGUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**aaguid** | **string** | Unique ID of a custom AAGUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomAAGUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customAAGUIDUpdateRequestObject** | [**CustomAAGUIDUpdateRequestObject**](CustomAAGUIDUpdateRequestObject.md) |  | 

### Return type

[**CustomAAGUIDResponseObject**](CustomAAGUIDResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyRpIdDomain

> VerifyRpIdDomain(ctx, authenticatorId, webAuthnMethodType).Execute()

Verify a Relying Party ID domain



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
	authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the authenticator
	webAuthnMethodType := "webAuthnMethodType_example" // string | Type of authenticator method

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticatorAPI.VerifyRpIdDomain(context.Background(), authenticatorId, webAuthnMethodType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorAPI.VerifyRpIdDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the authenticator | 
**webAuthnMethodType** | **string** | Type of authenticator method | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRpIdDomainRequest struct via the builder pattern


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

