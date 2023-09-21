# \AuthenticatorApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthenticator**](AuthenticatorApi.md#ActivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/activate | Activate an Authenticator
[**ActivateAuthenticatorMethod**](AuthenticatorApi.md#ActivateAuthenticatorMethod) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/activate | Activate an Authenticator Method
[**CreateAuthenticator**](AuthenticatorApi.md#CreateAuthenticator) | **Post** /api/v1/authenticators | Create an Authenticator
[**DeactivateAuthenticator**](AuthenticatorApi.md#DeactivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/deactivate | Deactivate an Authenticator
[**DeactivateAuthenticatorMethod**](AuthenticatorApi.md#DeactivateAuthenticatorMethod) | **Post** /api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/deactivate | Deactivate an Authenticator Method
[**GetAuthenticator**](AuthenticatorApi.md#GetAuthenticator) | **Get** /api/v1/authenticators/{authenticatorId} | Retrieve an Authenticator
[**GetAuthenticatorMethod**](AuthenticatorApi.md#GetAuthenticatorMethod) | **Get** /api/v1/authenticators/{authenticatorId}/methods/{methodType} | Retrieve a Method
[**GetWellKnownAppAuthenticatorConfiguration**](AuthenticatorApi.md#GetWellKnownAppAuthenticatorConfiguration) | **Get** /.well-known/app-authenticator-configuration | Retrieve the Well-Known App Authenticator Configuration
[**ListAuthenticatorMethods**](AuthenticatorApi.md#ListAuthenticatorMethods) | **Get** /api/v1/authenticators/{authenticatorId}/methods | List all Methods of an Authenticator
[**ListAuthenticators**](AuthenticatorApi.md#ListAuthenticators) | **Get** /api/v1/authenticators | List all Authenticators
[**ReplaceAuthenticator**](AuthenticatorApi.md#ReplaceAuthenticator) | **Put** /api/v1/authenticators/{authenticatorId} | Replace an Authenticator
[**ReplaceAuthenticatorMethod**](AuthenticatorApi.md#ReplaceAuthenticatorMethod) | **Put** /api/v1/authenticators/{authenticatorId}/methods/{methodType} | Replace a Method



## ActivateAuthenticator

> Authenticator ActivateAuthenticator(ctx, authenticatorId).Execute()

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
    resp, r, err := apiClient.AuthenticatorApi.ActivateAuthenticator(context.Background(), authenticatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.ActivateAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateAuthenticator`: Authenticator
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.ActivateAuthenticator`: %v\n", resp)
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

[**Authenticator**](Authenticator.md)

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
    methodType := openapiclient.AuthenticatorMethodType("cert") // AuthenticatorMethodType | Type of the authenticator method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticatorApi.ActivateAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.ActivateAuthenticatorMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.ActivateAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
**methodType** | [**AuthenticatorMethodType**](.md) | Type of the authenticator method | 

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

> Authenticator CreateAuthenticator(ctx).Authenticator(authenticator).Activate(activate).Execute()

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
    authenticator := *openapiclient.NewAuthenticator() // Authenticator | 
    activate := true // bool | Whether to execute the activation lifecycle operation when Okta creates the authenticator (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticatorApi.CreateAuthenticator(context.Background()).Authenticator(authenticator).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.CreateAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthenticator`: Authenticator
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.CreateAuthenticator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticator** | [**Authenticator**](Authenticator.md) |  | 
 **activate** | **bool** | Whether to execute the activation lifecycle operation when Okta creates the authenticator | [default to false]

### Return type

[**Authenticator**](Authenticator.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAuthenticator

> Authenticator DeactivateAuthenticator(ctx, authenticatorId).Execute()

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
    resp, r, err := apiClient.AuthenticatorApi.DeactivateAuthenticator(context.Background(), authenticatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.DeactivateAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateAuthenticator`: Authenticator
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.DeactivateAuthenticator`: %v\n", resp)
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

[**Authenticator**](Authenticator.md)

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
    methodType := openapiclient.AuthenticatorMethodType("cert") // AuthenticatorMethodType | Type of the authenticator method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticatorApi.DeactivateAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.DeactivateAuthenticatorMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.DeactivateAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
**methodType** | [**AuthenticatorMethodType**](.md) | Type of the authenticator method | 

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

> Authenticator GetAuthenticator(ctx, authenticatorId).Execute()

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
    resp, r, err := apiClient.AuthenticatorApi.GetAuthenticator(context.Background(), authenticatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.GetAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticator`: Authenticator
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.GetAuthenticator`: %v\n", resp)
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

[**Authenticator**](Authenticator.md)

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

Retrieve a Method



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
    methodType := openapiclient.AuthenticatorMethodType("cert") // AuthenticatorMethodType | Type of the authenticator method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticatorApi.GetAuthenticatorMethod(context.Background(), authenticatorId, methodType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.GetAuthenticatorMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.GetAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
**methodType** | [**AuthenticatorMethodType**](.md) | Type of the authenticator method | 

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
    resp, r, err := apiClient.AuthenticatorApi.GetWellKnownAppAuthenticatorConfiguration(context.Background()).OauthClientId(oauthClientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.GetWellKnownAppAuthenticatorConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWellKnownAppAuthenticatorConfiguration`: []WellKnownAppAuthenticatorConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.GetWellKnownAppAuthenticatorConfiguration`: %v\n", resp)
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
    resp, r, err := apiClient.AuthenticatorApi.ListAuthenticatorMethods(context.Background(), authenticatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.ListAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthenticatorMethods`: []ListAuthenticatorMethods200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.ListAuthenticatorMethods`: %v\n", resp)
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

> []Authenticator ListAuthenticators(ctx).Execute()

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
    resp, r, err := apiClient.AuthenticatorApi.ListAuthenticators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.ListAuthenticators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthenticators`: []Authenticator
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.ListAuthenticators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthenticatorsRequest struct via the builder pattern


### Return type

[**[]Authenticator**](Authenticator.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAuthenticator

> Authenticator ReplaceAuthenticator(ctx, authenticatorId).Authenticator(authenticator).Execute()

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
    authenticator := *openapiclient.NewAuthenticator() // Authenticator | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticatorApi.ReplaceAuthenticator(context.Background(), authenticatorId).Authenticator(authenticator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.ReplaceAuthenticator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthenticator`: Authenticator
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.ReplaceAuthenticator`: %v\n", resp)
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

 **authenticator** | [**Authenticator**](Authenticator.md) |  | 

### Return type

[**Authenticator**](Authenticator.md)

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

Replace a Method



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
    methodType := openapiclient.AuthenticatorMethodType("cert") // AuthenticatorMethodType | Type of the authenticator method
    listAuthenticatorMethods200ResponseInner := openapiclient.listAuthenticatorMethods_200_response_inner{AuthenticatorMethodOtp: openapiclient.NewAuthenticatorMethodOtp()} // ListAuthenticatorMethods200ResponseInner |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticatorApi.ReplaceAuthenticatorMethod(context.Background(), authenticatorId, methodType).ListAuthenticatorMethods200ResponseInner(listAuthenticatorMethods200ResponseInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorApi.ReplaceAuthenticatorMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthenticatorMethod`: ListAuthenticatorMethods200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorApi.ReplaceAuthenticatorMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticatorId** | **string** | &#x60;id&#x60; of the Authenticator | 
**methodType** | [**AuthenticatorMethodType**](.md) | Type of the authenticator method | 

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

