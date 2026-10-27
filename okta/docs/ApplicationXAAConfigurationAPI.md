# \ApplicationXAAConfigurationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetXAAResourceServer**](ApplicationXAAConfigurationAPI.md#GetXAAResourceServer) | **Get** /api/v1/apps/{appId}/xaa/resource-server | Retrieve the XAA resource server metadata



## GetXAAResourceServer

> XAAResourceServer GetXAAResourceServer(ctx, appId).Execute()

Retrieve the XAA resource server metadata



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationXAAConfigurationAPI.GetXAAResourceServer(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationXAAConfigurationAPI.GetXAAResourceServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetXAAResourceServer`: XAAResourceServer
	fmt.Fprintf(os.Stdout, "Response from `ApplicationXAAConfigurationAPI.GetXAAResourceServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXAAResourceServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XAAResourceServer**](XAAResourceServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

