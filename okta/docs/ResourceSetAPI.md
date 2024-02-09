# \ResourceSetAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMembersToBinding**](ResourceSetAPI.md#AddMembersToBinding) | **Patch** /api/v1/iam/resource-sets/{resourceSetId}/bindings/{roleIdOrLabel}/members | Add more Members to a binding
[**AddResourceSetResource**](ResourceSetAPI.md#AddResourceSetResource) | **Patch** /api/v1/iam/resource-sets/{resourceSetId}/resources | Add more Resource to a Resource Set
[**CreateResourceSet**](ResourceSetAPI.md#CreateResourceSet) | **Post** /api/v1/iam/resource-sets | Create a Resource Set
[**CreateResourceSetBinding**](ResourceSetAPI.md#CreateResourceSetBinding) | **Post** /api/v1/iam/resource-sets/{resourceSetId}/bindings | Create a Resource Set Binding
[**DeleteBinding**](ResourceSetAPI.md#DeleteBinding) | **Delete** /api/v1/iam/resource-sets/{resourceSetId}/bindings/{roleIdOrLabel} | Delete a Binding
[**DeleteResourceSet**](ResourceSetAPI.md#DeleteResourceSet) | **Delete** /api/v1/iam/resource-sets/{resourceSetId} | Delete a Resource Set
[**DeleteResourceSetResource**](ResourceSetAPI.md#DeleteResourceSetResource) | **Delete** /api/v1/iam/resource-sets/{resourceSetId}/resources/{resourceId} | Delete a Resource from a Resource Set
[**GetBinding**](ResourceSetAPI.md#GetBinding) | **Get** /api/v1/iam/resource-sets/{resourceSetId}/bindings/{roleIdOrLabel} | Retrieve a Binding
[**GetMemberOfBinding**](ResourceSetAPI.md#GetMemberOfBinding) | **Get** /api/v1/iam/resource-sets/{resourceSetId}/bindings/{roleIdOrLabel}/members/{memberId} | Retrieve a Member of a binding
[**GetResourceSet**](ResourceSetAPI.md#GetResourceSet) | **Get** /api/v1/iam/resource-sets/{resourceSetId} | Retrieve a Resource Set
[**ListBindings**](ResourceSetAPI.md#ListBindings) | **Get** /api/v1/iam/resource-sets/{resourceSetId}/bindings | List all Bindings
[**ListMembersOfBinding**](ResourceSetAPI.md#ListMembersOfBinding) | **Get** /api/v1/iam/resource-sets/{resourceSetId}/bindings/{roleIdOrLabel}/members | List all Members of a binding
[**ListResourceSetResources**](ResourceSetAPI.md#ListResourceSetResources) | **Get** /api/v1/iam/resource-sets/{resourceSetId}/resources | List all Resources of a Resource Set
[**ListResourceSets**](ResourceSetAPI.md#ListResourceSets) | **Get** /api/v1/iam/resource-sets | List all Resource Sets
[**ReplaceResourceSet**](ResourceSetAPI.md#ReplaceResourceSet) | **Put** /api/v1/iam/resource-sets/{resourceSetId} | Replace a Resource Set
[**UnassignMemberFromBinding**](ResourceSetAPI.md#UnassignMemberFromBinding) | **Delete** /api/v1/iam/resource-sets/{resourceSetId}/bindings/{roleIdOrLabel}/members/{memberId} | Unassign a Member from a binding



## AddMembersToBinding

> ResourceSetBindingResponse AddMembersToBinding(ctx, resourceSetId, roleIdOrLabel).Instance(instance).Execute()

Add more Members to a binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    instance := *openapiclient.NewResourceSetBindingAddMembersRequest() // ResourceSetBindingAddMembersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.AddMembersToBinding(context.Background(), resourceSetId, roleIdOrLabel).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.AddMembersToBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMembersToBinding`: ResourceSetBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.AddMembersToBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMembersToBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**ResourceSetBindingAddMembersRequest**](ResourceSetBindingAddMembersRequest.md) |  | 

### Return type

[**ResourceSetBindingResponse**](ResourceSetBindingResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddResourceSetResource

> ResourceSet AddResourceSetResource(ctx, resourceSetId).Instance(instance).Execute()

Add more Resource to a Resource Set



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    instance := *openapiclient.NewResourceSetResourcePatchRequest() // ResourceSetResourcePatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.AddResourceSetResource(context.Background(), resourceSetId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.AddResourceSetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResourceSetResource`: ResourceSet
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.AddResourceSetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceSetResourceRequest struct via the builder pattern


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


## CreateResourceSet

> ResourceSet CreateResourceSet(ctx).Instance(instance).Execute()

Create a Resource Set



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
    instance := *openapiclient.NewCreateResourceSetRequest() // CreateResourceSetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.CreateResourceSet(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.CreateResourceSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceSet`: ResourceSet
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.CreateResourceSet`: %v\n", resp)
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


## CreateResourceSetBinding

> ResourceSetBindingResponse CreateResourceSetBinding(ctx, resourceSetId).Instance(instance).Execute()

Create a Resource Set Binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    instance := *openapiclient.NewResourceSetBindingCreateRequest() // ResourceSetBindingCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.CreateResourceSetBinding(context.Background(), resourceSetId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.CreateResourceSetBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceSetBinding`: ResourceSetBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.CreateResourceSetBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceSetBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**ResourceSetBindingCreateRequest**](ResourceSetBindingCreateRequest.md) |  | 

### Return type

[**ResourceSetBindingResponse**](ResourceSetBindingResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBinding

> DeleteBinding(ctx, resourceSetId, roleIdOrLabel).Execute()

Delete a Binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceSetAPI.DeleteBinding(context.Background(), resourceSetId, roleIdOrLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.DeleteBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
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


## DeleteResourceSet

> DeleteResourceSet(ctx, resourceSetId).Execute()

Delete a Resource Set



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceSetAPI.DeleteResourceSet(context.Background(), resourceSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.DeleteResourceSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

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


## DeleteResourceSetResource

> DeleteResourceSetResource(ctx, resourceSetId, resourceId).Execute()

Delete a Resource from a Resource Set



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    resourceId := "ire106sQKoHoXXsAe0g4" // string | `id` of a resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceSetAPI.DeleteResourceSetResource(context.Background(), resourceSetId, resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.DeleteResourceSetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
**resourceId** | **string** | &#x60;id&#x60; of a resource | 

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


## GetBinding

> ResourceSetBindingResponse GetBinding(ctx, resourceSetId, roleIdOrLabel).Execute()

Retrieve a Binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.GetBinding(context.Background(), resourceSetId, roleIdOrLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.GetBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBinding`: ResourceSetBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.GetBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
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


## GetMemberOfBinding

> ResourceSetBindingMember GetMemberOfBinding(ctx, resourceSetId, roleIdOrLabel, memberId).Execute()

Retrieve a Member of a binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    memberId := "irb1qe6PGuMc7Oh8N0g4" // string | `id` of a member

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.GetMemberOfBinding(context.Background(), resourceSetId, roleIdOrLabel, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.GetMemberOfBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemberOfBinding`: ResourceSetBindingMember
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.GetMemberOfBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**memberId** | **string** | &#x60;id&#x60; of a member | 

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


## GetResourceSet

> ResourceSet GetResourceSet(ctx, resourceSetId).Execute()

Retrieve a Resource Set



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.GetResourceSet(context.Background(), resourceSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.GetResourceSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceSet`: ResourceSet
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.GetResourceSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

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


## ListBindings

> ResourceSetBindings ListBindings(ctx, resourceSetId).After(after).Execute()

List all Bindings



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.ListBindings(context.Background(), resourceSetId).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.ListBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBindings`: ResourceSetBindings
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.ListBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 

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


## ListMembersOfBinding

> ResourceSetBindingMembers ListMembersOfBinding(ctx, resourceSetId, roleIdOrLabel).After(after).Execute()

List all Members of a binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.ListMembersOfBinding(context.Background(), resourceSetId, roleIdOrLabel).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.ListMembersOfBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembersOfBinding`: ResourceSetBindingMembers
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.ListMembersOfBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersOfBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 

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


## ListResourceSetResources

> ResourceSetResources ListResourceSetResources(ctx, resourceSetId).Execute()

List all Resources of a Resource Set



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.ListResourceSetResources(context.Background(), resourceSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.ListResourceSetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResourceSetResources`: ResourceSetResources
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.ListResourceSetResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceSetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListResourceSets

> ResourceSets ListResourceSets(ctx).After(after).Execute()

List all Resource Sets



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
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.ListResourceSets(context.Background()).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.ListResourceSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResourceSets`: ResourceSets
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.ListResourceSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 

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

> ResourceSet ReplaceResourceSet(ctx, resourceSetId).Instance(instance).Execute()

Replace a Resource Set



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    instance := *openapiclient.NewResourceSet() // ResourceSet | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSetAPI.ReplaceResourceSet(context.Background(), resourceSetId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.ReplaceResourceSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceResourceSet`: ResourceSet
    fmt.Fprintf(os.Stdout, "Response from `ResourceSetAPI.ReplaceResourceSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 

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


## UnassignMemberFromBinding

> UnassignMemberFromBinding(ctx, resourceSetId, roleIdOrLabel, memberId).Execute()

Unassign a Member from a binding



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
    resourceSetId := "iamoJDFKaJxGIr0oamd9g" // string | `id` of a Resource Set
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    memberId := "irb1qe6PGuMc7Oh8N0g4" // string | `id` of a member

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceSetAPI.UnassignMemberFromBinding(context.Background(), resourceSetId, roleIdOrLabel, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSetAPI.UnassignMemberFromBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSetId** | **string** | &#x60;id&#x60; of a Resource Set | 
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**memberId** | **string** | &#x60;id&#x60; of a member | 

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

