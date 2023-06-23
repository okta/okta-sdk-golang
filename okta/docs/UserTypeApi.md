# \UserTypeApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserType**](UserTypeApi.md#CreateUserType) | **Post** /api/v1/meta/types/user | Create a User Type
[**DeleteUserType**](UserTypeApi.md#DeleteUserType) | **Delete** /api/v1/meta/types/user/{typeId} | Delete a User Type
[**GetUserType**](UserTypeApi.md#GetUserType) | **Get** /api/v1/meta/types/user/{typeId} | Retrieve a User Type
[**ListUserTypes**](UserTypeApi.md#ListUserTypes) | **Get** /api/v1/meta/types/user | List all User Types
[**ReplaceUserType**](UserTypeApi.md#ReplaceUserType) | **Put** /api/v1/meta/types/user/{typeId} | Replace a User Type
[**UpdateUserType**](UserTypeApi.md#UpdateUserType) | **Post** /api/v1/meta/types/user/{typeId} | Update a User Type



## CreateUserType

> UserType CreateUserType(ctx).UserType(userType).Execute()

Create a User Type



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
    userType := *openapiclient.NewUserType() // UserType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTypeApi.CreateUserType(context.Background()).UserType(userType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTypeApi.CreateUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserType`: UserType
    fmt.Fprintf(os.Stdout, "Response from `UserTypeApi.CreateUserType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userType** | [**UserType**](UserType.md) |  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserType

> DeleteUserType(ctx, typeId).Execute()

Delete a User Type



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
    typeId := "typeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTypeApi.DeleteUserType(context.Background(), typeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTypeApi.DeleteUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserTypeRequest struct via the builder pattern


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


## GetUserType

> UserType GetUserType(ctx, typeId).Execute()

Retrieve a User Type



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
    typeId := "typeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTypeApi.GetUserType(context.Background(), typeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTypeApi.GetUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserType`: UserType
    fmt.Fprintf(os.Stdout, "Response from `UserTypeApi.GetUserType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserType**](UserType.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserTypes

> []UserType ListUserTypes(ctx).Execute()

List all User Types



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
    resp, r, err := apiClient.UserTypeApi.ListUserTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTypeApi.ListUserTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserTypes`: []UserType
    fmt.Fprintf(os.Stdout, "Response from `UserTypeApi.ListUserTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserTypesRequest struct via the builder pattern


### Return type

[**[]UserType**](UserType.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUserType

> UserType ReplaceUserType(ctx, typeId).UserType(userType).Execute()

Replace a User Type



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
    typeId := "typeId_example" // string | 
    userType := *openapiclient.NewUserType() // UserType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTypeApi.ReplaceUserType(context.Background(), typeId).UserType(userType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTypeApi.ReplaceUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceUserType`: UserType
    fmt.Fprintf(os.Stdout, "Response from `UserTypeApi.ReplaceUserType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userType** | [**UserType**](UserType.md) |  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserType

> UserType UpdateUserType(ctx, typeId).UserType(userType).Execute()

Update a User Type



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
    typeId := "typeId_example" // string | 
    userType := *openapiclient.NewUserType() // UserType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTypeApi.UpdateUserType(context.Background(), typeId).UserType(userType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTypeApi.UpdateUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserType`: UserType
    fmt.Fprintf(os.Stdout, "Response from `UserTypeApi.UpdateUserType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userType** | [**UserType**](UserType.md) |  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

