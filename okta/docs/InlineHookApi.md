# \InlineHookApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateInlineHook**](InlineHookApi.md#ActivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/activate | Activate an Inline Hook
[**CreateInlineHook**](InlineHookApi.md#CreateInlineHook) | **Post** /api/v1/inlineHooks | Create an Inline Hook
[**DeactivateInlineHook**](InlineHookApi.md#DeactivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/deactivate | Deactivate an Inline Hook
[**DeleteInlineHook**](InlineHookApi.md#DeleteInlineHook) | **Delete** /api/v1/inlineHooks/{inlineHookId} | Delete an Inline Hook
[**ExecuteInlineHook**](InlineHookApi.md#ExecuteInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/execute | Execute an Inline Hook
[**GetInlineHook**](InlineHookApi.md#GetInlineHook) | **Get** /api/v1/inlineHooks/{inlineHookId} | Retrieve an Inline Hook
[**ListInlineHooks**](InlineHookApi.md#ListInlineHooks) | **Get** /api/v1/inlineHooks | List all Inline Hooks
[**ReplaceInlineHook**](InlineHookApi.md#ReplaceInlineHook) | **Put** /api/v1/inlineHooks/{inlineHookId} | Replace an Inline Hook



## ActivateInlineHook

> InlineHook ActivateInlineHook(ctx, inlineHookId).Execute()

Activate an Inline Hook



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
    inlineHookId := "inlineHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.ActivateInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.ActivateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.ActivateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInlineHook

> InlineHook CreateInlineHook(ctx).InlineHook(inlineHook).Execute()

Create an Inline Hook



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
    inlineHook := *openapiclient.NewInlineHook() // InlineHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.CreateInlineHook(context.Background()).InlineHook(inlineHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.CreateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.CreateInlineHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineHook** | [**InlineHook**](InlineHook.md) |  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateInlineHook

> InlineHook DeactivateInlineHook(ctx, inlineHookId).Execute()

Deactivate an Inline Hook



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
    inlineHookId := "inlineHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.DeactivateInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.DeactivateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.DeactivateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInlineHook

> DeleteInlineHook(ctx, inlineHookId).Execute()

Delete an Inline Hook



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
    inlineHookId := "inlineHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.DeleteInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.DeleteInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInlineHookRequest struct via the builder pattern


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


## ExecuteInlineHook

> InlineHookResponse ExecuteInlineHook(ctx, inlineHookId).PayloadData(payloadData).Execute()

Execute an Inline Hook



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
    inlineHookId := "inlineHookId_example" // string | 
    payloadData := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.ExecuteInlineHook(context.Background(), inlineHookId).PayloadData(payloadData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.ExecuteInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteInlineHook`: InlineHookResponse
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.ExecuteInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payloadData** | **map[string]interface{}** |  | 

### Return type

[**InlineHookResponse**](InlineHookResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineHook

> InlineHook GetInlineHook(ctx, inlineHookId).Execute()

Retrieve an Inline Hook



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
    inlineHookId := "inlineHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.GetInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.GetInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.GetInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInlineHooks

> []InlineHook ListInlineHooks(ctx).Type_(type_).Execute()

List all Inline Hooks



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
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.ListInlineHooks(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.ListInlineHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInlineHooks`: []InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.ListInlineHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInlineHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 

### Return type

[**[]InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceInlineHook

> InlineHook ReplaceInlineHook(ctx, inlineHookId).InlineHook(inlineHook).Execute()

Replace an Inline Hook



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
    inlineHookId := "inlineHookId_example" // string | 
    inlineHook := *openapiclient.NewInlineHook() // InlineHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookApi.ReplaceInlineHook(context.Background(), inlineHookId).InlineHook(inlineHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookApi.ReplaceInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookApi.ReplaceInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineHook** | [**InlineHook**](InlineHook.md) |  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

