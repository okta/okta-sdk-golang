# \AuthenticatorAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthenticator**](AuthenticatorAPI.md#ActivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/activate | Activate an Authenticator
[**ActivateAuthenticatorMethod**](AuthenticatorAPI.md#ActivateAuthenticatorMethod) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/activate | Activate an Authenticator Method
[**CreateAuthenticator**](AuthenticatorAPI.md#CreateAuthenticator) | **Post** /api/v1/authenticators | Create an Authenticator
[**DeactivateAuthenticator**](AuthenticatorAPI.md#DeactivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/deactivate | Deactivate an Authenticator
[**DeactivateAuthenticatorMethod**](AuthenticatorAPI.md#DeactivateAuthenticatorMethod) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/deactivate | Deactivate an Authenticator Method
[**GetAuthenticator**](AuthenticatorAPI.md#GetAuthenticator) | **Get** /api/v1/authenticators/{authenticatorId} | Retrieve an Authenticator
[**GetAuthenticatorMethod**](AuthenticatorAPI.md#GetAuthenticatorMethod) | **Get** /api/v1/authenticators/{authenticatorId}/methods/{methodType} | Retrieve an Authenticator Method
[**GetWellKnownAppAuthenticatorConfiguration**](AuthenticatorAPI.md#GetWellKnownAppAuthenticatorConfiguration) | **Get** /.well-known/app-authenticator-configuration | Retrieve the Well-Known App Authenticator Configuration
[**ListAuthenticatorMethods**](AuthenticatorAPI.md#ListAuthenticatorMethods) | **Get** /api/v1/authenticators/{authenticatorId}/methods | List all Methods of an Authenticator
[**ListAuthenticators**](AuthenticatorAPI.md#ListAuthenticators) | **Get** /api/v1/authenticators | List all Authenticators
[**ReplaceAuthenticator**](AuthenticatorAPI.md#ReplaceAuthenticator) | **Put** /api/v1/authenticators/{authenticatorId} | Replace an Authenticator
[**ReplaceAuthenticatorMethod**](AuthenticatorAPI.md#ReplaceAuthenticatorMethod) | **Put** /api/v1/authenticators/{authenticatorId}/methods/{methodType} | Replace an Authenticator Method



## ActivateAuthenticator

> AuthenticatorBase ActivateAuthenticator(ctx, authenticatorId).Execute()

Activate an Authenticator



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator

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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 

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

Activate an Authenticator Method



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator
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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
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

Create an Authenticator



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


## DeactivateAuthenticator

> AuthenticatorBase DeactivateAuthenticator(ctx, authenticatorId).Execute()

Deactivate an Authenticator



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator

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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 

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

Deactivate an Authenticator Method



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator
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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
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


## GetAuthenticator

> AuthenticatorBase GetAuthenticator(ctx, authenticatorId).Execute()

Retrieve an Authenticator



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator

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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 

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

Retrieve an Authenticator Method



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator
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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
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


## GetWellKnownAppAuthenticatorConfiguration

> []WellKnownAppAuthenticatorConfiguration GetWellKnownAppAuthenticatorConfiguration(ctx).OauthClientId(oauthClientId).Execute()

Retrieve the Well-Known App Authenticator Configuration



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


## ListAuthenticatorMethods

> []ListAuthenticatorMethods200ResponseInner ListAuthenticatorMethods(ctx, authenticatorId).Execute()

List all Methods of an Authenticator



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator

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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 

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

List all Authenticators



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

Replace an Authenticator



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator
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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 

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

Replace an Authenticator Method



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
    authenticatorId := "aut1nd8PQhGcQtSxB0g4" // string | `id` of the Authenticator
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
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
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

