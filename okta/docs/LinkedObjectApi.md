# \LinkedObjectApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLinkedObjectDefinition**](LinkedObjectApi.md#CreateLinkedObjectDefinition) | **Post** /api/v1/meta/schemas/user/linkedObjects | Create a Linked Object Definition
[**DeleteLinkedObjectDefinition**](LinkedObjectApi.md#DeleteLinkedObjectDefinition) | **Delete** /api/v1/meta/schemas/user/linkedObjects/{linkedObjectName} | Delete a Linked Object Definition
[**GetLinkedObjectDefinition**](LinkedObjectApi.md#GetLinkedObjectDefinition) | **Get** /api/v1/meta/schemas/user/linkedObjects/{linkedObjectName} | Retrieve a Linked Object Definition
[**ListLinkedObjectDefinitions**](LinkedObjectApi.md#ListLinkedObjectDefinitions) | **Get** /api/v1/meta/schemas/user/linkedObjects | List all Linked Object Definitions



## CreateLinkedObjectDefinition

> LinkedObject CreateLinkedObjectDefinition(ctx).LinkedObject(linkedObject).Execute()

Create a Linked Object Definition



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
    linkedObject := *openapiclient.NewLinkedObject() // LinkedObject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkedObjectApi.CreateLinkedObjectDefinition(context.Background()).LinkedObject(linkedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedObjectApi.CreateLinkedObjectDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLinkedObjectDefinition`: LinkedObject
    fmt.Fprintf(os.Stdout, "Response from `LinkedObjectApi.CreateLinkedObjectDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkedObjectDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkedObject** | [**LinkedObject**](LinkedObject.md) |  | 

### Return type

[**LinkedObject**](LinkedObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkedObjectDefinition

> DeleteLinkedObjectDefinition(ctx, linkedObjectName).Execute()

Delete a Linked Object Definition



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
    linkedObjectName := "linkedObjectName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkedObjectApi.DeleteLinkedObjectDefinition(context.Background(), linkedObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedObjectApi.DeleteLinkedObjectDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedObjectName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkedObjectDefinitionRequest struct via the builder pattern


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


## GetLinkedObjectDefinition

> LinkedObject GetLinkedObjectDefinition(ctx, linkedObjectName).Execute()

Retrieve a Linked Object Definition



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
    linkedObjectName := "linkedObjectName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkedObjectApi.GetLinkedObjectDefinition(context.Background(), linkedObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedObjectApi.GetLinkedObjectDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinkedObjectDefinition`: LinkedObject
    fmt.Fprintf(os.Stdout, "Response from `LinkedObjectApi.GetLinkedObjectDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedObjectName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedObjectDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkedObject**](LinkedObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLinkedObjectDefinitions

> []LinkedObject ListLinkedObjectDefinitions(ctx).Execute()

List all Linked Object Definitions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkedObjectApi.ListLinkedObjectDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedObjectApi.ListLinkedObjectDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLinkedObjectDefinitions`: []LinkedObject
    fmt.Fprintf(os.Stdout, "Response from `LinkedObjectApi.ListLinkedObjectDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLinkedObjectDefinitionsRequest struct via the builder pattern


### Return type

[**[]LinkedObject**](LinkedObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

