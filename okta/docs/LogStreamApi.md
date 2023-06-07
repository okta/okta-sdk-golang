# \LogStreamApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateLogStream**](LogStreamApi.md#ActivateLogStream) | **Post** /api/v1/logStreams/{logStreamId}/lifecycle/activate | Activate a Log Stream
[**CreateLogStream**](LogStreamApi.md#CreateLogStream) | **Post** /api/v1/logStreams | Create a Log Stream
[**DeactivateLogStream**](LogStreamApi.md#DeactivateLogStream) | **Post** /api/v1/logStreams/{logStreamId}/lifecycle/deactivate | Deactivate a Log Stream
[**DeleteLogStream**](LogStreamApi.md#DeleteLogStream) | **Delete** /api/v1/logStreams/{logStreamId} | Delete a Log Stream
[**GetLogStream**](LogStreamApi.md#GetLogStream) | **Get** /api/v1/logStreams/{logStreamId} | Retrieve a Log Stream
[**ListLogStreams**](LogStreamApi.md#ListLogStreams) | **Get** /api/v1/logStreams | List all Log Streams
[**ReplaceLogStream**](LogStreamApi.md#ReplaceLogStream) | **Put** /api/v1/logStreams/{logStreamId} | Replace a Log Stream



## ActivateLogStream

> ListLogStreams200ResponseInner ActivateLogStream(ctx, logStreamId).Execute()

Activate a Log Stream



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
    logStreamId := "abcd1234" // string | id of the log stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.ActivateLogStream(context.Background(), logStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.ActivateLogStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateLogStream`: ListLogStreams200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LogStreamApi.ActivateLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamId** | **string** | id of the log stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateLogStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogStream

> ListLogStreams200ResponseInner CreateLogStream(ctx).Instance(instance).Execute()

Create a Log Stream



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
    instance := openapiclient.listLogStreams_200_response_inner{LogStreamAws: openapiclient.NewLogStreamAws()} // ListLogStreams200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.CreateLogStream(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.CreateLogStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogStream`: ListLogStreams200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LogStreamApi.CreateLogStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md) |  | 

### Return type

[**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateLogStream

> ListLogStreams200ResponseInner DeactivateLogStream(ctx, logStreamId).Execute()

Deactivate a Log Stream



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
    logStreamId := "abcd1234" // string | id of the log stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.DeactivateLogStream(context.Background(), logStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.DeactivateLogStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateLogStream`: ListLogStreams200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LogStreamApi.DeactivateLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamId** | **string** | id of the log stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateLogStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogStream

> DeleteLogStream(ctx, logStreamId).Execute()

Delete a Log Stream



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
    logStreamId := "abcd1234" // string | id of the log stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.DeleteLogStream(context.Background(), logStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.DeleteLogStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamId** | **string** | id of the log stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogStreamRequest struct via the builder pattern


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


## GetLogStream

> ListLogStreams200ResponseInner GetLogStream(ctx, logStreamId).Execute()

Retrieve a Log Stream



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
    logStreamId := "abcd1234" // string | id of the log stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.GetLogStream(context.Background(), logStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.GetLogStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogStream`: ListLogStreams200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LogStreamApi.GetLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamId** | **string** | id of the log stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogStreams

> []ListLogStreams200ResponseInner ListLogStreams(ctx).After(after).Limit(limit).Filter(filter).Execute()

List all Log Streams



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
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination) for more information. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return. (optional) (default to 20)
    filter := "type eq "aws_eventbridge"" // string | SCIM filter expression that filters the results. This expression only supports the `eq` operator on either the `status` or `type`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.ListLogStreams(context.Background()).After(after).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.ListLogStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogStreams`: []ListLogStreams200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LogStreamApi.ListLogStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination) for more information. | 
 **limit** | **int32** | A limit on the number of objects to return. | [default to 20]
 **filter** | **string** | SCIM filter expression that filters the results. This expression only supports the &#x60;eq&#x60; operator on either the &#x60;status&#x60; or &#x60;type&#x60;. | 

### Return type

[**[]ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceLogStream

> ListLogStreams200ResponseInner ReplaceLogStream(ctx, logStreamId).Instance(instance).Execute()

Replace a Log Stream



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
    logStreamId := "abcd1234" // string | id of the log stream
    instance := openapiclient.listLogStreams_200_response_inner{LogStreamAws: openapiclient.NewLogStreamAws()} // ListLogStreams200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogStreamApi.ReplaceLogStream(context.Background(), logStreamId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogStreamApi.ReplaceLogStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceLogStream`: ListLogStreams200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LogStreamApi.ReplaceLogStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamId** | **string** | id of the log stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceLogStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md) |  | 

### Return type

[**ListLogStreams200ResponseInner**](ListLogStreams200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

