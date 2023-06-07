# \EventHookApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateEventHook**](EventHookApi.md#ActivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/activate | Activate an Event Hook
[**CreateEventHook**](EventHookApi.md#CreateEventHook) | **Post** /api/v1/eventHooks | Create an Event Hook
[**DeactivateEventHook**](EventHookApi.md#DeactivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/deactivate | Deactivate an Event Hook
[**DeleteEventHook**](EventHookApi.md#DeleteEventHook) | **Delete** /api/v1/eventHooks/{eventHookId} | Delete an Event Hook
[**GetEventHook**](EventHookApi.md#GetEventHook) | **Get** /api/v1/eventHooks/{eventHookId} | Retrieve an Event Hook
[**ListEventHooks**](EventHookApi.md#ListEventHooks) | **Get** /api/v1/eventHooks | List all Event Hooks
[**ReplaceEventHook**](EventHookApi.md#ReplaceEventHook) | **Put** /api/v1/eventHooks/{eventHookId} | Replace an Event Hook
[**VerifyEventHook**](EventHookApi.md#VerifyEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/verify | Verify an Event Hook



## ActivateEventHook

> EventHook ActivateEventHook(ctx, eventHookId).Execute()

Activate an Event Hook



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
    eventHookId := "eventHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.ActivateEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.ActivateEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.ActivateEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateEventHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEventHook

> EventHook CreateEventHook(ctx).EventHook(eventHook).Execute()

Create an Event Hook



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
    eventHook := *openapiclient.NewEventHook() // EventHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.CreateEventHook(context.Background()).EventHook(eventHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.CreateEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.CreateEventHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventHook** | [**EventHook**](EventHook.md) |  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateEventHook

> EventHook DeactivateEventHook(ctx, eventHookId).Execute()

Deactivate an Event Hook



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
    eventHookId := "eventHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.DeactivateEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.DeactivateEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.DeactivateEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateEventHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEventHook

> DeleteEventHook(ctx, eventHookId).Execute()

Delete an Event Hook



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
    eventHookId := "eventHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.DeleteEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.DeleteEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventHookRequest struct via the builder pattern


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


## GetEventHook

> EventHook GetEventHook(ctx, eventHookId).Execute()

Retrieve an Event Hook



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
    eventHookId := "eventHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.GetEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.GetEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.GetEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventHooks

> []EventHook ListEventHooks(ctx).Execute()

List all Event Hooks



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
    resp, r, err := apiClient.EventHookApi.ListEventHooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.ListEventHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventHooks`: []EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.ListEventHooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEventHooksRequest struct via the builder pattern


### Return type

[**[]EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEventHook

> EventHook ReplaceEventHook(ctx, eventHookId).EventHook(eventHook).Execute()

Replace an Event Hook



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
    eventHookId := "eventHookId_example" // string | 
    eventHook := *openapiclient.NewEventHook() // EventHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.ReplaceEventHook(context.Background(), eventHookId).EventHook(eventHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.ReplaceEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.ReplaceEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEventHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventHook** | [**EventHook**](EventHook.md) |  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEventHook

> EventHook VerifyEventHook(ctx, eventHookId).Execute()

Verify an Event Hook



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
    eventHookId := "eventHookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookApi.VerifyEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookApi.VerifyEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookApi.VerifyEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEventHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventHook**](EventHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

