# \RoleDResourceSetBindingAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceSetBinding**](RoleDResourceSetBindingAPI.md#CreateResourceSetBinding) | **Post** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings | Create a role resource set binding
[**DeleteBinding**](RoleDResourceSetBindingAPI.md#DeleteBinding) | **Delete** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel} | Delete a role resource set binding
[**GetBinding**](RoleDResourceSetBindingAPI.md#GetBinding) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel} | Retrieve a role resource set binding
[**ListBindings**](RoleDResourceSetBindingAPI.md#ListBindings) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings | List all role resource set bindings



## CreateResourceSetBinding

> ResourceSetBindingEditResponse CreateResourceSetBinding(ctx, resourceSetIdOrLabel).Instance(instance).Execute()

Create a role resource set binding



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
	instance := *openapiclient.NewResourceSetBindingCreateRequest() // ResourceSetBindingCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleDResourceSetBindingAPI.CreateResourceSetBinding(context.Background(), resourceSetIdOrLabel).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingAPI.CreateResourceSetBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourceSetBinding`: ResourceSetBindingEditResponse
	fmt.Fprintf(os.Stdout, "Response from `RoleDResourceSetBindingAPI.CreateResourceSetBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceSetBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**ResourceSetBindingCreateRequest**](ResourceSetBindingCreateRequest.md) |  | 

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


## DeleteBinding

> DeleteBinding(ctx, resourceSetIdOrLabel, roleIdOrLabel).Execute()

Delete a role resource set binding



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleDResourceSetBindingAPI.DeleteBinding(context.Background(), resourceSetIdOrLabel, roleIdOrLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingAPI.DeleteBinding``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBindingRequest struct via the builder pattern


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


## GetBinding

> ResourceSetBindingResponse GetBinding(ctx, resourceSetIdOrLabel, roleIdOrLabel).Execute()

Retrieve a role resource set binding



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleDResourceSetBindingAPI.GetBinding(context.Background(), resourceSetIdOrLabel, roleIdOrLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingAPI.GetBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinding`: ResourceSetBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `RoleDResourceSetBindingAPI.GetBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResourceSetBindingResponse**](ResourceSetBindingResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBindings

> ResourceSetBindings ListBindings(ctx, resourceSetIdOrLabel).After(after).Execute()

List all role resource set bindings



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleDResourceSetBindingAPI.ListBindings(context.Background(), resourceSetIdOrLabel).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleDResourceSetBindingAPI.ListBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBindings`: ResourceSetBindings
	fmt.Fprintf(os.Stdout, "Response from `RoleDResourceSetBindingAPI.ListBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 

### Return type

[**ResourceSetBindings**](ResourceSetBindings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

