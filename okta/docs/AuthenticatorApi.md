# \AuthenticatorApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthenticator**](AuthenticatorApi.md#ActivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/activate | Activate an Authenticator
[**CreateAuthenticator**](AuthenticatorApi.md#CreateAuthenticator) | **Post** /api/v1/authenticators | Create an Authenticator
[**DeactivateAuthenticator**](AuthenticatorApi.md#DeactivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/deactivate | Deactivate an Authenticator
[**GetAuthenticator**](AuthenticatorApi.md#GetAuthenticator) | **Get** /api/v1/authenticators/{authenticatorId} | Retrieve an Authenticator
[**ListAuthenticators**](AuthenticatorApi.md#ListAuthenticators) | **Get** /api/v1/authenticators | List all Authenticators
[**ReplaceAuthenticator**](AuthenticatorApi.md#ReplaceAuthenticator) | **Put** /api/v1/authenticators/{authenticatorId} | Replace an Authenticator



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
    openapiclient "./openapi"
)

func main() {
    authenticatorId := "authenticatorId_example" // string | 

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
**authenticatorId** | **string** |  | 

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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    authenticatorId := "authenticatorId_example" // string | 

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
**authenticatorId** | **string** |  | 

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
    openapiclient "./openapi"
)

func main() {
    authenticatorId := "authenticatorId_example" // string | 

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
**authenticatorId** | **string** |  | 

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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    authenticatorId := "authenticatorId_example" // string | 
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
**authenticatorId** | **string** |  | 

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

