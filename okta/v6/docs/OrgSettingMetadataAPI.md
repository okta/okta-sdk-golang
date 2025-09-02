# \OrgSettingMetadataAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWellknownOrgMetadata**](OrgSettingMetadataAPI.md#GetWellknownOrgMetadata) | **Get** /.well-known/okta-organization | Retrieve the Org metadata



## GetWellknownOrgMetadata

> WellKnownOrgMetadata GetWellknownOrgMetadata(ctx).Execute()

Retrieve the Org metadata



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
	resp, r, err := apiClient.OrgSettingMetadataAPI.GetWellknownOrgMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingMetadataAPI.GetWellknownOrgMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWellknownOrgMetadata`: WellKnownOrgMetadata
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingMetadataAPI.GetWellknownOrgMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWellknownOrgMetadataRequest struct via the builder pattern


### Return type

[**WellKnownOrgMetadata**](WellKnownOrgMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

