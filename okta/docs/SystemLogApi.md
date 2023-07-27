# \SystemLogApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogEvents**](SystemLogApi.md#ListLogEvents) | **Get** /api/v1/logs | List all System Log Events



## ListLogEvents

> []LogEvent ListLogEvents(ctx).Since(since).Until(until).Filter(filter).Q(q).Limit(limit).SortOrder(sortOrder).After(after).Execute()

List all System Log Events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filter := "filter_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 100)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "ASCENDING")
    after := "after_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemLogApi.ListLogEvents(context.Background()).Since(since).Until(until).Filter(filter).Q(q).Limit(limit).SortOrder(sortOrder).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemLogApi.ListLogEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogEvents`: []LogEvent
    fmt.Fprintf(os.Stdout, "Response from `SystemLogApi.ListLogEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filter** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **sortOrder** | **string** |  | [default to &quot;ASCENDING&quot;]
 **after** | **string** |  | 

### Return type

[**[]LogEvent**](LogEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

