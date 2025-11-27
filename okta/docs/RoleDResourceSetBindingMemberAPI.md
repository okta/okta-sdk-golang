# \RoleDResourceSetBindingMemberAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMembersToBinding**](RoleDResourceSetBindingMemberAPI.md#AddMembersToBinding) | **Patch** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members | Add more role resource set binding members
[**GetMemberOfBinding**](RoleDResourceSetBindingMemberAPI.md#GetMemberOfBinding) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members/{memberId} | Retrieve a role resource set binding member
[**ListMembersOfBinding**](RoleDResourceSetBindingMemberAPI.md#ListMembersOfBinding) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members | List all role resource set binding members
[**UnassignMemberFromBinding**](RoleDResourceSetBindingMemberAPI.md#UnassignMemberFromBinding) | **Delete** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members/{memberId} | Unassign a role resource set binding member



## AddMembersToBinding

> ResourceSetBindingEditResponse AddMembersToBinding(ctx, resourceSetIdOrLabel, roleIdOrLabel).Instance(instance).Execute()

Add more role resource set binding members



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
	resourceSetIdOrLabel := "iamoJDFKaJxGIr0oamd9g" // string | `id` or `label` of the resource set
	roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
	instance := *openapiclient.NewResourceSetBindingAddMembersRequest() // ResourceSetBindingAddMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleDResourceSetBindingMemberAPI.AddMembersToBinding(context.Background(), resourceSetIdOrLabel, roleIdOrLabel).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingMemberAPI.AddMembersToBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMembersToBinding`: ResourceSetBindingEditResponse
	fmt.Fprintf(os.Stdout, "Response from `RoleDResourceSetBindingMemberAPI.AddMembersToBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMembersToBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**ResourceSetBindingAddMembersRequest**](ResourceSetBindingAddMembersRequest.md) |  | 

### Return type

[**ResourceSetBindingEditResponse**](ResourceSetBindingEditResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemberOfBinding

> ResourceSetBindingMember GetMemberOfBinding(ctx, resourceSetIdOrLabel, roleIdOrLabel, memberId).Execute()

Retrieve a role resource set binding member



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
	resourceSetIdOrLabel := "iamoJDFKaJxGIr0oamd9g" // string | `id` or `label` of the resource set
	roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
	memberId := "irb1qe6PGuMc7Oh8N0g4" // string | `id` of the member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleDResourceSetBindingMemberAPI.GetMemberOfBinding(context.Background(), resourceSetIdOrLabel, roleIdOrLabel, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingMemberAPI.GetMemberOfBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMemberOfBinding`: ResourceSetBindingMember
	fmt.Fprintf(os.Stdout, "Response from `RoleDResourceSetBindingMemberAPI.GetMemberOfBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**memberId** | **string** | &#x60;id&#x60; of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberOfBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceSetBindingMember**](ResourceSetBindingMember.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMembersOfBinding

> ResourceSetBindingMembers ListMembersOfBinding(ctx, resourceSetIdOrLabel, roleIdOrLabel).After(after).Execute()

List all role resource set binding members



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
	resourceSetIdOrLabel := "iamoJDFKaJxGIr0oamd9g" // string | `id` or `label` of the resource set
	roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleDResourceSetBindingMemberAPI.ListMembersOfBinding(context.Background(), resourceSetIdOrLabel, roleIdOrLabel).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingMemberAPI.ListMembersOfBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMembersOfBinding`: ResourceSetBindingMembers
	fmt.Fprintf(os.Stdout, "Response from `RoleDResourceSetBindingMemberAPI.ListMembersOfBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersOfBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 

### Return type

[**ResourceSetBindingMembers**](ResourceSetBindingMembers.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignMemberFromBinding

> UnassignMemberFromBinding(ctx, resourceSetIdOrLabel, roleIdOrLabel, memberId).Execute()

Unassign a role resource set binding member



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
	resourceSetIdOrLabel := "iamoJDFKaJxGIr0oamd9g" // string | `id` or `label` of the resource set
	roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
	memberId := "irb1qe6PGuMc7Oh8N0g4" // string | `id` of the member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleDResourceSetBindingMemberAPI.UnassignMemberFromBinding(context.Background(), resourceSetIdOrLabel, roleIdOrLabel, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingMemberAPI.UnassignMemberFromBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**memberId** | **string** | &#x60;id&#x60; of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignMemberFromBindingRequest struct via the builder pattern


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

