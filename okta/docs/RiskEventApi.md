# \RiskEventApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendRiskEvents**](RiskEventApi.md#SendRiskEvents) | **Post** /api/v1/risk/events/ip | Send multiple Risk Events



## SendRiskEvents

> SendRiskEvents(ctx).Instance(instance).Execute()

Send multiple Risk Events



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
    instance := []openapiclient.RiskEvent{*openapiclient.NewRiskEvent([]openapiclient.RiskEventSubject{*openapiclient.NewRiskEventSubject("Ip_example")})} // []RiskEvent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskEventApi.SendRiskEvents(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskEventApi.SendRiskEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRiskEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**[]RiskEvent**](RiskEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

