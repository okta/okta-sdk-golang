# \RoleCResourceSetResourceAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourceSetResource**](RoleCResourceSetResourceAPI.md#AddResourceSetResource) | **Post** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/resources | Add a resource set resource with conditions
[**AddResourceSetResources**](RoleCResourceSetResourceAPI.md#AddResourceSetResources) | **Patch** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/resources | Add more resources to a resource set
[**DeleteResourceSetResource**](RoleCResourceSetResourceAPI.md#DeleteResourceSetResource) | **Delete** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/resources/{resourceId} | Delete a resource set resource
[**GetResourceSetResource**](RoleCResourceSetResourceAPI.md#GetResourceSetResource) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/resources/{resourceId} | Retrieve a resource set resource
[**ListResourceSetResources**](RoleCResourceSetResourceAPI.md#ListResourceSetResources) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/resources | List all resource set resources
[**ReplaceResourceSetResource**](RoleCResourceSetResourceAPI.md#ReplaceResourceSetResource) | **Put** /api/v1/iam/resource-sets/{resourceSetIdOrLabel}/resources/{resourceId} | Replace the resource set resource conditions



## AddResourceSetResource

> ResourceSetResource AddResourceSetResource(ctx, resourceSetIdOrLabel).Instance(instance).Execute()

Add a resource set resource with conditions



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
	instance := *openapiclient.NewResourceSetResourcePostRequest(*openapiclient.NewResourceConditions(), "ResourceOrnOrUrl_example") // ResourceSetResourcePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetResourceAPI.AddResourceSetResource(context.Background(), resourceSetIdOrLabel).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetResourceAPI.AddResourceSetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddResourceSetResource`: ResourceSetResource
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetResourceAPI.AddResourceSetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceSetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**ResourceSetResourcePostRequest**](ResourceSetResourcePostRequest.md) |  | 

### Return type

[**ResourceSetResource**](ResourceSetResource.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddResourceSetResources

> ResourceSet AddResourceSetResources(ctx, resourceSetIdOrLabel).Instance(instance).Execute()

Add more resources to a resource set



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
	instance := *openapiclient.NewResourceSetResourcePatchRequest() // ResourceSetResourcePatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetResourceAPI.AddResourceSetResources(context.Background(), resourceSetIdOrLabel).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetResourceAPI.AddResourceSetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddResourceSetResources`: ResourceSet
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetResourceAPI.AddResourceSetResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceSetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**ResourceSetResourcePatchRequest**](ResourceSetResourcePatchRequest.md) |  | 

### Return type

[**ResourceSet**](ResourceSet.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceSetResource

> DeleteResourceSetResource(ctx, resourceSetIdOrLabel, resourceId).Execute()

Delete a resource set resource



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
	resourceId := "ire106sQKoHoXXsAe0g4" // string | `id` of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleCResourceSetResourceAPI.DeleteResourceSetResource(context.Background(), resourceSetIdOrLabel, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetResourceAPI.DeleteResourceSetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**resourceId** | **string** | &#x60;id&#x60; of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceSetResourceRequest struct via the builder pattern


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


## GetResourceSetResource

> ResourceSetResource GetResourceSetResource(ctx, resourceSetIdOrLabel, resourceId).Execute()

Retrieve a resource set resource



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
	resourceId := "ire106sQKoHoXXsAe0g4" // string | `id` of the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetResourceAPI.GetResourceSetResource(context.Background(), resourceSetIdOrLabel, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetResourceAPI.GetResourceSetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceSetResource`: ResourceSetResource
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetResourceAPI.GetResourceSetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**resourceId** | **string** | &#x60;id&#x60; of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceSetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResourceSetResource**](ResourceSetResource.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceSetResources

> ResourceSetResources ListResourceSetResources(ctx, resourceSetIdOrLabel).After(after).Limit(limit).Execute()

List all resource set resources



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
	after := "after_example" // string | Specifies the pagination cursor for the next page of targets (optional)
	limit := int32(56) // int32 | Specifies the number of results returned. Defaults to `100`. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetResourceAPI.ListResourceSetResources(context.Background(), resourceSetIdOrLabel).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetResourceAPI.ListResourceSetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceSetResources`: ResourceSetResources
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetResourceAPI.ListResourceSetResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceSetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | Specifies the pagination cursor for the next page of targets | 
 **limit** | **int32** | Specifies the number of results returned. Defaults to &#x60;100&#x60;. | [default to 100]

### Return type

[**ResourceSetResources**](ResourceSetResources.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceResourceSetResource

> ResourceSetResource ReplaceResourceSetResource(ctx, resourceSetIdOrLabel, resourceId).ResourceSetResourcePutRequest(resourceSetResourcePutRequest).Execute()

Replace the resource set resource conditions



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
	resourceId := "ire106sQKoHoXXsAe0g4" // string | `id` of the resource
	resourceSetResourcePutRequest := *openapiclient.NewResourceSetResourcePutRequest() // ResourceSetResourcePutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetResourceAPI.ReplaceResourceSetResource(context.Background(), resourceSetIdOrLabel, resourceId).ResourceSetResourcePutRequest(resourceSetResourcePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetResourceAPI.ReplaceResourceSetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceResourceSetResource`: ResourceSetResource
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetResourceAPI.ReplaceResourceSetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 
**resourceId** | **string** | &#x60;id&#x60; of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceResourceSetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceSetResourcePutRequest** | [**ResourceSetResourcePutRequest**](ResourceSetResourcePutRequest.md) |  | 

### Return type

[**ResourceSetResource**](ResourceSetResource.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

