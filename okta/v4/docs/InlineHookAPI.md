# \InlineHookAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateInlineHook**](InlineHookAPI.md#ActivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/activate | Activate an Inline Hook
[**CreateInlineHook**](InlineHookAPI.md#CreateInlineHook) | **Post** /api/v1/inlineHooks | Create an Inline Hook
[**DeactivateInlineHook**](InlineHookAPI.md#DeactivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/deactivate | Deactivate an Inline Hook
[**DeleteInlineHook**](InlineHookAPI.md#DeleteInlineHook) | **Delete** /api/v1/inlineHooks/{inlineHookId} | Delete an Inline Hook
[**ExecuteInlineHook**](InlineHookAPI.md#ExecuteInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/execute | Execute an Inline Hook
[**GetInlineHook**](InlineHookAPI.md#GetInlineHook) | **Get** /api/v1/inlineHooks/{inlineHookId} | Retrieve an Inline Hook
[**ListInlineHooks**](InlineHookAPI.md#ListInlineHooks) | **Get** /api/v1/inlineHooks | List all Inline Hooks
[**ReplaceInlineHook**](InlineHookAPI.md#ReplaceInlineHook) | **Put** /api/v1/inlineHooks/{inlineHookId} | Replace an Inline Hook
[**UpdateInlineHook**](InlineHookAPI.md#UpdateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId} | Update an Inline Hook



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.ActivateInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ActivateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ActivateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHook := *openapiclient.NewInlineHook() // InlineHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.CreateInlineHook(context.Background()).InlineHook(inlineHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.CreateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.CreateInlineHook`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.DeactivateInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.DeactivateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.DeactivateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InlineHookAPI.DeleteInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.DeleteInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook
    payloadData := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.ExecuteInlineHook(context.Background(), inlineHookId).PayloadData(payloadData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ExecuteInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteInlineHook`: InlineHookResponse
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ExecuteInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.GetInlineHook(context.Background(), inlineHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.GetInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.GetInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.ListInlineHooks(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ListInlineHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInlineHooks`: []InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ListInlineHooks`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook
    inlineHook := *openapiclient.NewInlineHook() // InlineHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.ReplaceInlineHook(context.Background(), inlineHookId).InlineHook(inlineHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ReplaceInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ReplaceInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

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


## UpdateInlineHook

> InlineHook UpdateInlineHook(ctx, inlineHookId).InlineHook(inlineHook).Execute()

Update an Inline Hook



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
    inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the Inline Hook
    inlineHook := *openapiclient.NewInlineHook() // InlineHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InlineHookAPI.UpdateInlineHook(context.Background(), inlineHookId).InlineHook(inlineHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.UpdateInlineHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInlineHook`: InlineHook
    fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.UpdateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the Inline Hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInlineHookRequest struct via the builder pattern


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

