# \TrustedOriginApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTrustedOrigin**](TrustedOriginApi.md#ActivateTrustedOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/activate | Activate a Trusted Origin
[**CreateTrustedOrigin**](TrustedOriginApi.md#CreateTrustedOrigin) | **Post** /api/v1/trustedOrigins | Create a Trusted Origin
[**DeactivateTrustedOrigin**](TrustedOriginApi.md#DeactivateTrustedOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/deactivate | Deactivate a Trusted Origin
[**DeleteTrustedOrigin**](TrustedOriginApi.md#DeleteTrustedOrigin) | **Delete** /api/v1/trustedOrigins/{trustedOriginId} | Delete a Trusted Origin
[**GetTrustedOrigin**](TrustedOriginApi.md#GetTrustedOrigin) | **Get** /api/v1/trustedOrigins/{trustedOriginId} | Retrieve a Trusted Origin
[**ListTrustedOrigins**](TrustedOriginApi.md#ListTrustedOrigins) | **Get** /api/v1/trustedOrigins | List all Trusted Origins
[**ReplaceTrustedOrigin**](TrustedOriginApi.md#ReplaceTrustedOrigin) | **Put** /api/v1/trustedOrigins/{trustedOriginId} | Replace a Trusted Origin



## ActivateTrustedOrigin

> TrustedOrigin ActivateTrustedOrigin(ctx, trustedOriginId).Execute()

Activate a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.ActivateTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.ActivateTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.ActivateTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateTrustedOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrustedOrigin

> TrustedOrigin CreateTrustedOrigin(ctx).TrustedOrigin(trustedOrigin).Execute()

Create a Trusted Origin



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
    trustedOrigin := *openapiclient.NewTrustedOrigin() // TrustedOrigin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.CreateTrustedOrigin(context.Background()).TrustedOrigin(trustedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.CreateTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.CreateTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrustedOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trustedOrigin** | [**TrustedOrigin**](TrustedOrigin.md) |  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateTrustedOrigin

> TrustedOrigin DeactivateTrustedOrigin(ctx, trustedOriginId).Execute()

Deactivate a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.DeactivateTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.DeactivateTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.DeactivateTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateTrustedOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustedOrigin

> DeleteTrustedOrigin(ctx, trustedOriginId).Execute()

Delete a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.DeleteTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.DeleteTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustedOriginRequest struct via the builder pattern


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


## GetTrustedOrigin

> TrustedOrigin GetTrustedOrigin(ctx, trustedOriginId).Execute()

Retrieve a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.GetTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.GetTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.GetTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustedOrigins

> []TrustedOrigin ListTrustedOrigins(ctx).Q(q).Filter(filter).After(after).Limit(limit).Execute()

List all Trusted Origins



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
    q := "q_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.ListTrustedOrigins(context.Background()).Q(q).Filter(filter).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.ListTrustedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrustedOrigins`: []TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.ListTrustedOrigins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrustedOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **filter** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]

### Return type

[**[]TrustedOrigin**](TrustedOrigin.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceTrustedOrigin

> TrustedOrigin ReplaceTrustedOrigin(ctx, trustedOriginId).TrustedOrigin(trustedOrigin).Execute()

Replace a Trusted Origin



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
    trustedOriginId := "trustedOriginId_example" // string | 
    trustedOrigin := *openapiclient.NewTrustedOrigin() // TrustedOrigin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginApi.ReplaceTrustedOrigin(context.Background(), trustedOriginId).TrustedOrigin(trustedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginApi.ReplaceTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginApi.ReplaceTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceTrustedOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedOrigin** | [**TrustedOrigin**](TrustedOrigin.md) |  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

