# \TrustedOriginAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTrustedOrigin**](TrustedOriginAPI.md#ActivateTrustedOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/activate | Activate a Trusted Origin
[**CreateTrustedOrigin**](TrustedOriginAPI.md#CreateTrustedOrigin) | **Post** /api/v1/trustedOrigins | Create a Trusted Origin
[**DeactivateTrustedOrigin**](TrustedOriginAPI.md#DeactivateTrustedOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/deactivate | Deactivate a Trusted Origin
[**DeleteTrustedOrigin**](TrustedOriginAPI.md#DeleteTrustedOrigin) | **Delete** /api/v1/trustedOrigins/{trustedOriginId} | Delete a Trusted Origin
[**GetTrustedOrigin**](TrustedOriginAPI.md#GetTrustedOrigin) | **Get** /api/v1/trustedOrigins/{trustedOriginId} | Retrieve a Trusted Origin
[**ListTrustedOrigins**](TrustedOriginAPI.md#ListTrustedOrigins) | **Get** /api/v1/trustedOrigins | List all Trusted Origins
[**ReplaceTrustedOrigin**](TrustedOriginAPI.md#ReplaceTrustedOrigin) | **Put** /api/v1/trustedOrigins/{trustedOriginId} | Replace a Trusted Origin



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    trustedOriginId := "7j2PkU1nyNIDe26ZNufR" // string | `id` of the Trusted Origin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginAPI.ActivateTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.ActivateTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginAPI.ActivateTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** | &#x60;id&#x60; of the Trusted Origin | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    trustedOrigin := *openapiclient.NewTrustedOriginWrite() // TrustedOriginWrite | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginAPI.CreateTrustedOrigin(context.Background()).TrustedOrigin(trustedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.CreateTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginAPI.CreateTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrustedOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trustedOrigin** | [**TrustedOriginWrite**](TrustedOriginWrite.md) |  | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    trustedOriginId := "7j2PkU1nyNIDe26ZNufR" // string | `id` of the Trusted Origin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginAPI.DeactivateTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.DeactivateTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginAPI.DeactivateTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** | &#x60;id&#x60; of the Trusted Origin | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    trustedOriginId := "7j2PkU1nyNIDe26ZNufR" // string | `id` of the Trusted Origin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TrustedOriginAPI.DeleteTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.DeleteTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** | &#x60;id&#x60; of the Trusted Origin | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    trustedOriginId := "7j2PkU1nyNIDe26ZNufR" // string | `id` of the Trusted Origin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginAPI.GetTrustedOrigin(context.Background(), trustedOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.GetTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginAPI.GetTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** | &#x60;id&#x60; of the Trusted Origin | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    q := "q_example" // string | A search string that will prefix match against the `name` and `origin` (optional)
    filter := "name eq "Example Trusted Origin"" // string | [Filter](/#filter) Trusted Origins with a supported expression for a subset of properties. You can filter on the following properties: `name`, `origin`, `status`, and `type` (type of scopes).  (optional)
    after := "after_example" // string | The after cursor provided by a prior request. (optional)
    limit := int32(56) // int32 | Specifies the number of results. (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginAPI.ListTrustedOrigins(context.Background()).Q(q).Filter(filter).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.ListTrustedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrustedOrigins`: []TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginAPI.ListTrustedOrigins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrustedOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | A search string that will prefix match against the &#x60;name&#x60; and &#x60;origin&#x60; | 
 **filter** | **string** | [Filter](/#filter) Trusted Origins with a supported expression for a subset of properties. You can filter on the following properties: &#x60;name&#x60;, &#x60;origin&#x60;, &#x60;status&#x60;, and &#x60;type&#x60; (type of scopes).  | 
 **after** | **string** | The after cursor provided by a prior request. | 
 **limit** | **int32** | Specifies the number of results. | [default to 20]

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    trustedOriginId := "7j2PkU1nyNIDe26ZNufR" // string | `id` of the Trusted Origin
    trustedOrigin := *openapiclient.NewTrustedOrigin() // TrustedOrigin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedOriginAPI.ReplaceTrustedOrigin(context.Background(), trustedOriginId).TrustedOrigin(trustedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedOriginAPI.ReplaceTrustedOrigin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceTrustedOrigin`: TrustedOrigin
    fmt.Fprintf(os.Stdout, "Response from `TrustedOriginAPI.ReplaceTrustedOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedOriginId** | **string** | &#x60;id&#x60; of the Trusted Origin | 

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

