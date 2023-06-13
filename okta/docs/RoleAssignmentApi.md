# \RoleAssignmentApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToGroup**](RoleAssignmentApi.md#AssignRoleToGroup) | **Post** /api/v1/groups/{groupId}/roles | Assign a Role to a Group
[**AssignRoleToUser**](RoleAssignmentApi.md#AssignRoleToUser) | **Post** /api/v1/users/{userId}/roles | Assign a Role to a User
[**GetGroupAssignedRole**](RoleAssignmentApi.md#GetGroupAssignedRole) | **Get** /api/v1/groups/{groupId}/roles/{roleId} | Retrieve a Role assigned to Group
[**GetUserAssignedRole**](RoleAssignmentApi.md#GetUserAssignedRole) | **Get** /api/v1/users/{userId}/roles/{roleId} | Retrieve a Role assigned to a User
[**ListAssignedRolesForUser**](RoleAssignmentApi.md#ListAssignedRolesForUser) | **Get** /api/v1/users/{userId}/roles | List all Roles assigned to a User
[**ListGroupAssignedRoles**](RoleAssignmentApi.md#ListGroupAssignedRoles) | **Get** /api/v1/groups/{groupId}/roles | List all Assigned Roles of Group
[**UnassignRoleFromGroup**](RoleAssignmentApi.md#UnassignRoleFromGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId} | Unassign a Role from a Group
[**UnassignRoleFromUser**](RoleAssignmentApi.md#UnassignRoleFromUser) | **Delete** /api/v1/users/{userId}/roles/{roleId} | Unassign a Role from a User



## AssignRoleToGroup

> Role AssignRoleToGroup(ctx, groupId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()

Assign a Role to a Group



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
    groupId := "groupId_example" // string | 
    assignRoleRequest := *openapiclient.NewAssignRoleRequest() // AssignRoleRequest | 
    disableNotifications := true // bool | Setting this to `true` grants the group third-party admin status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.AssignRoleToGroup(context.Background(), groupId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.AssignRoleToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRoleToGroup`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.AssignRoleToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRoleToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignRoleRequest** | [**AssignRoleRequest**](AssignRoleRequest.md) |  | 
 **disableNotifications** | **bool** | Setting this to &#x60;true&#x60; grants the group third-party admin status | 

### Return type

[**Role**](Role.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignRoleToUser

> Role AssignRoleToUser(ctx, userId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()

Assign a Role to a User



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
    userId := "userId_example" // string | 
    assignRoleRequest := *openapiclient.NewAssignRoleRequest() // AssignRoleRequest | 
    disableNotifications := true // bool | Setting this to `true` grants the user third-party admin status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.AssignRoleToUser(context.Background(), userId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.AssignRoleToUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRoleToUser`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.AssignRoleToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRoleToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignRoleRequest** | [**AssignRoleRequest**](AssignRoleRequest.md) |  | 
 **disableNotifications** | **bool** | Setting this to &#x60;true&#x60; grants the user third-party admin status | 

### Return type

[**Role**](Role.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAssignedRole

> Role GetGroupAssignedRole(ctx, groupId, roleId).Execute()

Retrieve a Role assigned to Group



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
    groupId := "groupId_example" // string | 
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.GetGroupAssignedRole(context.Background(), groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.GetGroupAssignedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupAssignedRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.GetGroupAssignedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAssignedRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAssignedRole

> Role GetUserAssignedRole(ctx, userId, roleId).Execute()

Retrieve a Role assigned to a User



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
    userId := "userId_example" // string | 
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.GetUserAssignedRole(context.Background(), userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.GetUserAssignedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAssignedRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.GetUserAssignedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAssignedRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssignedRolesForUser

> []Role ListAssignedRolesForUser(ctx, userId).Expand(expand).Execute()

List all Roles assigned to a User



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
    userId := "userId_example" // string | 
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.ListAssignedRolesForUser(context.Background(), userId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListAssignedRolesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssignedRolesForUser`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListAssignedRolesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssignedRolesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | 

### Return type

[**[]Role**](Role.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupAssignedRoles

> []Role ListGroupAssignedRoles(ctx, groupId).Expand(expand).Execute()

List all Assigned Roles of Group



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
    groupId := "groupId_example" // string | 
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.ListGroupAssignedRoles(context.Background(), groupId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListGroupAssignedRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupAssignedRoles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListGroupAssignedRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupAssignedRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | 

### Return type

[**[]Role**](Role.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignRoleFromGroup

> UnassignRoleFromGroup(ctx, groupId, roleId).Execute()

Unassign a Role from a Group



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
    groupId := "groupId_example" // string | 
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.UnassignRoleFromGroup(context.Background(), groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.UnassignRoleFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignRoleFromGroupRequest struct via the builder pattern


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


## UnassignRoleFromUser

> UnassignRoleFromUser(ctx, userId, roleId).Execute()

Unassign a Role from a User



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
    userId := "userId_example" // string | 
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentApi.UnassignRoleFromUser(context.Background(), userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.UnassignRoleFromUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignRoleFromUserRequest struct via the builder pattern


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

