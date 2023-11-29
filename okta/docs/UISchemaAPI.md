# \UISchemaAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUISchema**](UISchemaAPI.md#CreateUISchema) | **Post** /api/v1/meta/uischemas | Create a UI Schema
[**DeleteUISchemas**](UISchemaAPI.md#DeleteUISchemas) | **Delete** /api/v1/meta/uischemas/{id} | Delete a UI Schema
[**GetUISchema**](UISchemaAPI.md#GetUISchema) | **Get** /api/v1/meta/uischemas/{id} | Retrieve a UI Schema
[**ListUISchemas**](UISchemaAPI.md#ListUISchemas) | **Get** /api/v1/meta/uischemas | List all UI Schemas
[**ReplaceUISchemas**](UISchemaAPI.md#ReplaceUISchemas) | **Put** /api/v1/meta/uischemas/{id} | Replace a UI Schema



## CreateUISchema

> UISchemasResponseObject CreateUISchema(ctx).Uischemabody(uischemabody).Execute()

Create a UI Schema



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
    uischemabody := *openapiclient.NewCreateUISchema() // CreateUISchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UISchemaAPI.CreateUISchema(context.Background()).Uischemabody(uischemabody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UISchemaAPI.CreateUISchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUISchema`: UISchemasResponseObject
    fmt.Fprintf(os.Stdout, "Response from `UISchemaAPI.CreateUISchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUISchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uischemabody** | [**CreateUISchema**](CreateUISchema.md) |  | 

### Return type

[**UISchemasResponseObject**](UISchemasResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUISchemas

> DeleteUISchemas(ctx, id).Execute()

Delete a UI Schema



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
    id := "uis4a7liocgcRgcxZ0g7" // string | The unique ID of the UI Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UISchemaAPI.DeleteUISchemas(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UISchemaAPI.DeleteUISchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the UI Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUISchemasRequest struct via the builder pattern


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


## GetUISchema

> UISchemasResponseObject GetUISchema(ctx, id).Execute()

Retrieve a UI Schema



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
    id := "uis4a7liocgcRgcxZ0g7" // string | The unique ID of the UI Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UISchemaAPI.GetUISchema(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UISchemaAPI.GetUISchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUISchema`: UISchemasResponseObject
    fmt.Fprintf(os.Stdout, "Response from `UISchemaAPI.GetUISchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the UI Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUISchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UISchemasResponseObject**](UISchemasResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUISchemas

> []UISchemasResponseObject ListUISchemas(ctx).Execute()

List all UI Schemas



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
    resp, r, err := apiClient.UISchemaAPI.ListUISchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UISchemaAPI.ListUISchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUISchemas`: []UISchemasResponseObject
    fmt.Fprintf(os.Stdout, "Response from `UISchemaAPI.ListUISchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUISchemasRequest struct via the builder pattern


### Return type

[**[]UISchemasResponseObject**](UISchemasResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUISchemas

> UISchemasResponseObject ReplaceUISchemas(ctx, id).UpdateUISchemaBody(updateUISchemaBody).Execute()

Replace a UI Schema



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
    id := "uis4a7liocgcRgcxZ0g7" // string | The unique ID of the UI Schema
    updateUISchemaBody := *openapiclient.NewUpdateUISchema() // UpdateUISchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UISchemaAPI.ReplaceUISchemas(context.Background(), id).UpdateUISchemaBody(updateUISchemaBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UISchemaAPI.ReplaceUISchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceUISchemas`: UISchemasResponseObject
    fmt.Fprintf(os.Stdout, "Response from `UISchemaAPI.ReplaceUISchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the UI Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUISchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUISchemaBody** | [**UpdateUISchema**](UpdateUISchema.md) |  | 

### Return type

[**UISchemasResponseObject**](UISchemasResponseObject.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

