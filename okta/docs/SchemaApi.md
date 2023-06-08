# \SchemaApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationLayout**](SchemaApi.md#GetApplicationLayout) | **Get** /api/v1/meta/layouts/apps/{appName} | Retrieve the UI Layout for an Application
[**GetApplicationUserSchema**](SchemaApi.md#GetApplicationUserSchema) | **Get** /api/v1/meta/schemas/apps/{appInstanceId}/default | Retrieve the default Application User Schema for an Application
[**GetGroupSchema**](SchemaApi.md#GetGroupSchema) | **Get** /api/v1/meta/schemas/group/default | Retrieve the default Group Schema
[**GetLogStreamSchema**](SchemaApi.md#GetLogStreamSchema) | **Get** /api/v1/meta/schemas/logStream/{logStreamType} | Retrieve the Log Stream Schema for the schema type
[**GetUserSchema**](SchemaApi.md#GetUserSchema) | **Get** /api/v1/meta/schemas/user/{schemaId} | Retrieve a User Schema
[**ListLogStreamSchemas**](SchemaApi.md#ListLogStreamSchemas) | **Get** /api/v1/meta/schemas/logStream | List the Log Stream Schemas
[**UpdateApplicationUserProfile**](SchemaApi.md#UpdateApplicationUserProfile) | **Post** /api/v1/meta/schemas/apps/{appInstanceId}/default | Update the default Application User Schema for an Application
[**UpdateGroupSchema**](SchemaApi.md#UpdateGroupSchema) | **Post** /api/v1/meta/schemas/group/default | Update the default Group Schema
[**UpdateUserProfile**](SchemaApi.md#UpdateUserProfile) | **Post** /api/v1/meta/schemas/user/{schemaId} | Update a User Schema



## GetApplicationLayout

> ApplicationLayout GetApplicationLayout(ctx, appName).Execute()

Retrieve the UI Layout for an Application



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.GetApplicationLayout(context.Background(), appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.GetApplicationLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationLayout`: ApplicationLayout
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.GetApplicationLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationLayout**](ApplicationLayout.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationUserSchema

> UserSchema GetApplicationUserSchema(ctx, appInstanceId).Execute()

Retrieve the default Application User Schema for an Application



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
    appInstanceId := "appInstanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.GetApplicationUserSchema(context.Background(), appInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.GetApplicationUserSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationUserSchema`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.GetApplicationUserSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationUserSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupSchema

> GroupSchema GetGroupSchema(ctx).Execute()

Retrieve the default Group Schema



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
    resp, r, err := apiClient.SchemaApi.GetGroupSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.GetGroupSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupSchema`: GroupSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.GetGroupSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupSchemaRequest struct via the builder pattern


### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogStreamSchema

> LogStreamSchema GetLogStreamSchema(ctx, logStreamType).Execute()

Retrieve the Log Stream Schema for the schema type



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
    logStreamType := openapiclient.LogStreamType("aws_eventbridge") // LogStreamType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.GetLogStreamSchema(context.Background(), logStreamType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.GetLogStreamSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogStreamSchema`: LogStreamSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.GetLogStreamSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamType** | [**LogStreamType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogStreamSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogStreamSchema**](LogStreamSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSchema

> UserSchema GetUserSchema(ctx, schemaId).Execute()

Retrieve a User Schema



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
    schemaId := "schemaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.GetUserSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.GetUserSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSchema`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.GetUserSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogStreamSchemas

> []LogStreamSchema ListLogStreamSchemas(ctx).Execute()

List the Log Stream Schemas



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
    resp, r, err := apiClient.SchemaApi.ListLogStreamSchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.ListLogStreamSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogStreamSchemas`: []LogStreamSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.ListLogStreamSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogStreamSchemasRequest struct via the builder pattern


### Return type

[**[]LogStreamSchema**](LogStreamSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationUserProfile

> UserSchema UpdateApplicationUserProfile(ctx, appInstanceId).Body(body).Execute()

Update the default Application User Schema for an Application



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
    appInstanceId := "appInstanceId_example" // string | 
    body := *openapiclient.NewUserSchema() // UserSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.UpdateApplicationUserProfile(context.Background(), appInstanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.UpdateApplicationUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationUserProfile`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.UpdateApplicationUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserSchema**](UserSchema.md) |  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupSchema

> GroupSchema UpdateGroupSchema(ctx).GroupSchema(groupSchema).Execute()

Update the default Group Schema



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
    groupSchema := *openapiclient.NewGroupSchema() // GroupSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.UpdateGroupSchema(context.Background()).GroupSchema(groupSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.UpdateGroupSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupSchema`: GroupSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.UpdateGroupSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupSchema** | [**GroupSchema**](GroupSchema.md) |  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserProfile

> UserSchema UpdateUserProfile(ctx, schemaId).UserSchema(userSchema).Execute()

Update a User Schema



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
    schemaId := "schemaId_example" // string | 
    userSchema := *openapiclient.NewUserSchema() // UserSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaApi.UpdateUserProfile(context.Background(), schemaId).UserSchema(userSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.UpdateUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserProfile`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.UpdateUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSchema** | [**UserSchema**](UserSchema.md) |  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

