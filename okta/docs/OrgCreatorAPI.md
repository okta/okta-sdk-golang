# \OrgCreatorAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChildOrg**](OrgCreatorAPI.md#CreateChildOrg) | **Post** /api/v1/orgs | Create an org



## CreateChildOrg

> ChildOrg CreateChildOrg(ctx).ChildOrg(childOrg).Execute()

Create an org



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
	childOrg := *openapiclient.NewChildOrg(*openapiclient.NewOrgCreationAdmin(*openapiclient.NewOrgCreationAdminProfile("FirstName_example", "LastName_example", "Email_example", "Login_example")), "SKU", "My Child Org 1", "my-child-org-1") // ChildOrg |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgCreatorAPI.CreateChildOrg(context.Background()).ChildOrg(childOrg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgCreatorAPI.CreateChildOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChildOrg`: ChildOrg
	fmt.Fprintf(os.Stdout, "Response from `OrgCreatorAPI.CreateChildOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChildOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childOrg** | [**ChildOrg**](ChildOrg.md) |  | 

### Return type

[**ChildOrg**](ChildOrg.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

