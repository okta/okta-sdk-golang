# \EmailCustomizationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkRemoveEmailAddressBounces**](EmailCustomizationAPI.md#BulkRemoveEmailAddressBounces) | **Post** /api/v1/org/email/bounces/remove-list | Remove bounced emails



## BulkRemoveEmailAddressBounces

> BouncesRemoveListResult BulkRemoveEmailAddressBounces(ctx).BouncesRemoveListObj(bouncesRemoveListObj).Execute()

Remove bounced emails



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
	bouncesRemoveListObj := *openapiclient.NewBouncesRemoveListObj() // BouncesRemoveListObj |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailCustomizationAPI.BulkRemoveEmailAddressBounces(context.Background()).BouncesRemoveListObj(bouncesRemoveListObj).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailCustomizationAPI.BulkRemoveEmailAddressBounces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkRemoveEmailAddressBounces`: BouncesRemoveListResult
	fmt.Fprintf(os.Stdout, "Response from `EmailCustomizationAPI.BulkRemoveEmailAddressBounces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveEmailAddressBouncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bouncesRemoveListObj** | [**BouncesRemoveListObj**](BouncesRemoveListObj.md) |  | 

### Return type

[**BouncesRemoveListResult**](BouncesRemoveListResult.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

