# \RoleCResourceSetAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceSet**](RoleCResourceSetAPI.md#CreateResourceSet) | **Post** /api/v1/iam/resource-sets | Create a resource set
[**DeleteResourceSet**](RoleCResourceSetAPI.md#DeleteResourceSet) | **Delete** /api/v1/iam/resource-sets/{resourceSetIdOrLabel} | Delete a resource set
[**GetResourceSet**](RoleCResourceSetAPI.md#GetResourceSet) | **Get** /api/v1/iam/resource-sets/{resourceSetIdOrLabel} | Retrieve a resource set
[**ListResourceSets**](RoleCResourceSetAPI.md#ListResourceSets) | **Get** /api/v1/iam/resource-sets | List all resource sets
[**ReplaceResourceSet**](RoleCResourceSetAPI.md#ReplaceResourceSet) | **Put** /api/v1/iam/resource-sets/{resourceSetIdOrLabel} | Replace a resource set



## CreateResourceSet

> ResourceSet CreateResourceSet(ctx).Instance(instance).Execute()

Create a resource set



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
	instance := *openapiclient.NewCreateResourceSetRequest("Description_example", "Label_example", []string{"Resources_example"}) // CreateResourceSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetAPI.CreateResourceSet(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetAPI.CreateResourceSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourceSet`: ResourceSet
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetAPI.CreateResourceSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**CreateResourceSetRequest**](CreateResourceSetRequest.md) |  | 

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


## DeleteResourceSet

> DeleteResourceSet(ctx, resourceSetIdOrLabel).Execute()

Delete a resource set



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleCResourceSetAPI.DeleteResourceSet(context.Background(), resourceSetIdOrLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetAPI.DeleteResourceSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceSetRequest struct via the builder pattern


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


## GetResourceSet

> ResourceSet GetResourceSet(ctx, resourceSetIdOrLabel).Execute()

Retrieve a resource set



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetAPI.GetResourceSet(context.Background(), resourceSetIdOrLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetAPI.GetResourceSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceSet`: ResourceSet
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetAPI.GetResourceSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceSet**](ResourceSet.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceSets

> ResourceSets ListResourceSets(ctx).After(after).Execute()

List all resource sets



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetAPI.ListResourceSets(context.Background()).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetAPI.ListResourceSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceSets`: ResourceSets
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetAPI.ListResourceSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**ResourceSets**](ResourceSets.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceResourceSet

> ResourceSet ReplaceResourceSet(ctx, resourceSetIdOrLabel).Instance(instance).Execute()

Replace a resource set



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
	instance := *openapiclient.NewResourceSet() // ResourceSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleCResourceSetAPI.ReplaceResourceSet(context.Background(), resourceSetIdOrLabel).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleCResourceSetAPI.ReplaceResourceSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceResourceSet`: ResourceSet
	fmt.Fprintf(os.Stdout, "Response from `RoleCResourceSetAPI.ReplaceResourceSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the resource set | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceResourceSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**ResourceSet**](ResourceSet.md) |  | 

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

