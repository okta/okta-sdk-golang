# \SystemLogAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogEvents**](SystemLogAPI.md#ListLogEvents) | **Get** /api/v1/logs | List all System Log events



## ListLogEvents

> []LogEvent ListLogEvents(ctx).Since(since).Until(until).After(after).Filter(filter).Q(q).Limit(limit).SortOrder(sortOrder).Execute()

List all System Log events



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
	since := "since_example" // string | Filters the lower time bound of the log events `published` property for bounded queries or persistence time for polling queries (optional) (default to "7 days prior to until")
	until := "until_example" // string | Filters the upper time bound of the log events `published` property for bounded queries or persistence time for polling queries. (optional) (default to "current time")
	after := "after_example" // string | Retrieves the next page of results. Okta returns a link in the HTTP Header (`rel=next`) that includes the after query parameter (optional)
	filter := "filter_example" // string | Filter expression that filters the results. All operators except [ ] are supported. See [Filter](https://developer.okta.com/docs/api/#filter) and [Operators](https://developer.okta.com/docs/api/#operators). (optional)
	q := "q_example" // string | Filters log events results by one or more case insensitive keywords. (optional)
	limit := int32(56) // int32 | Sets the number of results that are returned in the response (optional) (default to 100)
	sortOrder := "sortOrder_example" // string | The order of the returned events that are sorted by the `published` property (optional) (default to "ASCENDING")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemLogAPI.ListLogEvents(context.Background()).Since(since).Until(until).After(after).Filter(filter).Q(q).Limit(limit).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemLogAPI.ListLogEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogEvents`: []LogEvent
	fmt.Fprintf(os.Stdout, "Response from `SystemLogAPI.ListLogEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Filters the lower time bound of the log events &#x60;published&#x60; property for bounded queries or persistence time for polling queries | [default to &quot;7 days prior to until&quot;]
 **until** | **string** | Filters the upper time bound of the log events &#x60;published&#x60; property for bounded queries or persistence time for polling queries. | [default to &quot;current time&quot;]
 **after** | **string** | Retrieves the next page of results. Okta returns a link in the HTTP Header (&#x60;rel&#x3D;next&#x60;) that includes the after query parameter | 
 **filter** | **string** | Filter expression that filters the results. All operators except [ ] are supported. See [Filter](https://developer.okta.com/docs/api/#filter) and [Operators](https://developer.okta.com/docs/api/#operators). | 
 **q** | **string** | Filters log events results by one or more case insensitive keywords. | 
 **limit** | **int32** | Sets the number of results that are returned in the response | [default to 100]
 **sortOrder** | **string** | The order of the returned events that are sorted by the &#x60;published&#x60; property | [default to &quot;ASCENDING&quot;]

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

