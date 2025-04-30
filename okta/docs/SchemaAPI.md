# \SchemaAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationUserSchema**](SchemaAPI.md#GetApplicationUserSchema) | **Get** /api/v1/meta/schemas/apps/{appId}/default | Retrieve the default Application User Schema for an Application
[**GetGroupSchema**](SchemaAPI.md#GetGroupSchema) | **Get** /api/v1/meta/schemas/group/default | Retrieve the default Group Schema
[**GetLogStreamSchema**](SchemaAPI.md#GetLogStreamSchema) | **Get** /api/v1/meta/schemas/logStream/{logStreamType} | Retrieve the Log Stream Schema for the schema type
[**GetUserSchema**](SchemaAPI.md#GetUserSchema) | **Get** /api/v1/meta/schemas/user/{schemaId} | Retrieve a User Schema
[**ListLogStreamSchemas**](SchemaAPI.md#ListLogStreamSchemas) | **Get** /api/v1/meta/schemas/logStream | List the Log Stream Schemas
[**UpdateApplicationUserProfile**](SchemaAPI.md#UpdateApplicationUserProfile) | **Post** /api/v1/meta/schemas/apps/{appId}/default | Update the default Application User Schema for an Application
[**UpdateGroupSchema**](SchemaAPI.md#UpdateGroupSchema) | **Post** /api/v1/meta/schemas/group/default | Update the default Group Schema
[**UpdateUserProfile**](SchemaAPI.md#UpdateUserProfile) | **Post** /api/v1/meta/schemas/user/{schemaId} | Update a User Schema



## GetApplicationUserSchema

> UserSchema GetApplicationUserSchema(ctx, appId).Execute()

Retrieve the default Application User Schema for an Application



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.GetApplicationUserSchema(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.GetApplicationUserSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationUserSchema`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.GetApplicationUserSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.GetGroupSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.GetGroupSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupSchema`: GroupSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.GetGroupSchema`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    logStreamType := "logStreamType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.GetLogStreamSchema(context.Background(), logStreamType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.GetLogStreamSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogStreamSchema`: LogStreamSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.GetLogStreamSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logStreamType** | **string** |  | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    schemaId := "schemaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.GetUserSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.GetUserSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSchema`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.GetUserSchema`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.ListLogStreamSchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.ListLogStreamSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogStreamSchemas`: []LogStreamSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.ListLogStreamSchemas`: %v\n", resp)
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

> UserSchema UpdateApplicationUserProfile(ctx, appId).Body(body).Execute()

Update the default Application User Schema for an Application



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    body := *openapiclient.NewUserSchema() // UserSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.UpdateApplicationUserProfile(context.Background(), appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.UpdateApplicationUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationUserProfile`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.UpdateApplicationUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupSchema := *openapiclient.NewGroupSchema() // GroupSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.UpdateGroupSchema(context.Background()).GroupSchema(groupSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.UpdateGroupSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupSchema`: GroupSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.UpdateGroupSchema`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    schemaId := "schemaId_example" // string | 
    userSchema := *openapiclient.NewUserSchema() // UserSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaAPI.UpdateUserProfile(context.Background(), schemaId).UserSchema(userSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.UpdateUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserProfile`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.UpdateUserProfile`: %v\n", resp)
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

