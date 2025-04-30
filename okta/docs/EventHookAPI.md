# \EventHookAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateEventHook**](EventHookAPI.md#ActivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/activate | Activate an Event Hook
[**CreateEventHook**](EventHookAPI.md#CreateEventHook) | **Post** /api/v1/eventHooks | Create an Event Hook
[**DeactivateEventHook**](EventHookAPI.md#DeactivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/deactivate | Deactivate an Event Hook
[**DeleteEventHook**](EventHookAPI.md#DeleteEventHook) | **Delete** /api/v1/eventHooks/{eventHookId} | Delete an Event Hook
[**GetEventHook**](EventHookAPI.md#GetEventHook) | **Get** /api/v1/eventHooks/{eventHookId} | Retrieve an Event Hook
[**ListEventHooks**](EventHookAPI.md#ListEventHooks) | **Get** /api/v1/eventHooks | List all Event Hooks
[**ReplaceEventHook**](EventHookAPI.md#ReplaceEventHook) | **Put** /api/v1/eventHooks/{eventHookId} | Replace an Event Hook
[**VerifyEventHook**](EventHookAPI.md#VerifyEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/verify | Verify an Event Hook



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHookId := "who8vt36qfNpCGz9H1e6" // string | `id` of the Event Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.ActivateEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.ActivateEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.ActivateEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** | &#x60;id&#x60; of the Event Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHook := *openapiclient.NewEventHook(*openapiclient.NewEventHookChannel(*openapiclient.NewEventHookChannelConfig("Uri_example"), "Type_example", "Version_example"), *openapiclient.NewEventSubscriptions([]string{"Items_example"}, "Type_example"), "Name_example") // EventHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.CreateEventHook(context.Background()).EventHook(eventHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.CreateEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.CreateEventHook`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHookId := "who8vt36qfNpCGz9H1e6" // string | `id` of the Event Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.DeactivateEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.DeactivateEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.DeactivateEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** | &#x60;id&#x60; of the Event Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHookId := "who8vt36qfNpCGz9H1e6" // string | `id` of the Event Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventHookAPI.DeleteEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.DeleteEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** | &#x60;id&#x60; of the Event Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHookId := "who8vt36qfNpCGz9H1e6" // string | `id` of the Event Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.GetEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.GetEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.GetEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** | &#x60;id&#x60; of the Event Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.ListEventHooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.ListEventHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventHooks`: []EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.ListEventHooks`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHookId := "who8vt36qfNpCGz9H1e6" // string | `id` of the Event Hook
    eventHook := *openapiclient.NewEventHook(*openapiclient.NewEventHookChannel(*openapiclient.NewEventHookChannelConfig("Uri_example"), "Type_example", "Version_example"), *openapiclient.NewEventSubscriptions([]string{"Items_example"}, "Type_example"), "Name_example") // EventHook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.ReplaceEventHook(context.Background(), eventHookId).EventHook(eventHook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.ReplaceEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.ReplaceEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** | &#x60;id&#x60; of the Event Hook | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    eventHookId := "who8vt36qfNpCGz9H1e6" // string | `id` of the Event Hook

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHookAPI.VerifyEventHook(context.Background(), eventHookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHookAPI.VerifyEventHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEventHook`: EventHook
    fmt.Fprintf(os.Stdout, "Response from `EventHookAPI.VerifyEventHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHookId** | **string** | &#x60;id&#x60; of the Event Hook | 

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

