# \RoleAssignmentAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToClient**](RoleAssignmentAPI.md#AssignRoleToClient) | **Post** /oauth2/v1/clients/{clientId}/roles | Assign Role to Client
[**AssignRoleToGroup**](RoleAssignmentAPI.md#AssignRoleToGroup) | **Post** /api/v1/groups/{groupId}/roles | Assign a Role to a Group
[**AssignRoleToUser**](RoleAssignmentAPI.md#AssignRoleToUser) | **Post** /api/v1/users/{userId}/roles | Assign a Role to a User
[**DeleteRoleFromClient**](RoleAssignmentAPI.md#DeleteRoleFromClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleId} | Unassign a Role from a Client
[**GetGroupAssignedRole**](RoleAssignmentAPI.md#GetGroupAssignedRole) | **Get** /api/v1/groups/{groupId}/roles/{roleId} | Retrieve a Role assigned to Group
[**GetUserAssignedRole**](RoleAssignmentAPI.md#GetUserAssignedRole) | **Get** /api/v1/users/{userId}/roles/{roleId} | Retrieve a Role assigned to a User
[**ListAssignedRolesForUser**](RoleAssignmentAPI.md#ListAssignedRolesForUser) | **Get** /api/v1/users/{userId}/roles | List all Roles assigned to a User
[**ListGroupAssignedRoles**](RoleAssignmentAPI.md#ListGroupAssignedRoles) | **Get** /api/v1/groups/{groupId}/roles | List all Assigned Roles of Group
[**ListRolesForClient**](RoleAssignmentAPI.md#ListRolesForClient) | **Get** /oauth2/v1/clients/{clientId}/roles | List all Roles for a Client
[**ListUsersWithRoleAssignments**](RoleAssignmentAPI.md#ListUsersWithRoleAssignments) | **Get** /api/v1/iam/assignees/users | List all Users with Role Assignments
[**RetrieveClientRole**](RoleAssignmentAPI.md#RetrieveClientRole) | **Get** /oauth2/v1/clients/{clientId}/roles/{roleId} | Retrieve a Client Role
[**UnassignRoleFromGroup**](RoleAssignmentAPI.md#UnassignRoleFromGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId} | Unassign a Role from a Group
[**UnassignRoleFromUser**](RoleAssignmentAPI.md#UnassignRoleFromUser) | **Delete** /api/v1/users/{userId}/roles/{roleId} | Unassign a Role from a User



## AssignRoleToClient

> Client AssignRoleToClient(ctx, clientId).AssignRoleToClientRequest(assignRoleToClientRequest).Execute()

Assign Role to Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    assignRoleToClientRequest := openapiclient.assignRoleToClient_request{CustomRoleAssignmentSchema: openapiclient.NewCustomRoleAssignmentSchema()} // AssignRoleToClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.AssignRoleToClient(context.Background(), clientId).AssignRoleToClientRequest(assignRoleToClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.AssignRoleToClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRoleToClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.AssignRoleToClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRoleToClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignRoleToClientRequest** | [**AssignRoleToClientRequest**](AssignRoleToClientRequest.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    assignRoleRequest := *openapiclient.NewAssignRoleRequest() // AssignRoleRequest | 
    disableNotifications := true // bool | Setting this to `true` grants the group third-party admin status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.AssignRoleToGroup(context.Background(), groupId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.AssignRoleToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRoleToGroup`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.AssignRoleToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    assignRoleRequest := *openapiclient.NewAssignRoleRequest() // AssignRoleRequest | 
    disableNotifications := true // bool | Setting this to `true` grants the user third-party admin status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.AssignRoleToUser(context.Background(), userId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.AssignRoleToUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRoleToUser`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.AssignRoleToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

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


## DeleteRoleFromClient

> DeleteRoleFromClient(ctx, clientId, roleId).Execute()

Unassign a Role from a Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleAssignmentAPI.DeleteRoleFromClient(context.Background(), clientId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.DeleteRoleFromClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleFromClientRequest struct via the builder pattern


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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.GetGroupAssignedRole(context.Background(), groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.GetGroupAssignedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupAssignedRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.GetGroupAssignedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.GetUserAssignedRole(context.Background(), userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.GetUserAssignedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAssignedRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.GetUserAssignedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.ListAssignedRolesForUser(context.Background(), userId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.ListAssignedRolesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssignedRolesForUser`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.ListAssignedRolesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.ListGroupAssignedRoles(context.Background(), groupId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.ListGroupAssignedRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupAssignedRoles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.ListGroupAssignedRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

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


## ListRolesForClient

> Client ListRolesForClient(ctx, clientId).Execute()

List all Roles for a Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.ListRolesForClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.ListRolesForClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRolesForClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.ListRolesForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Client**](Client.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersWithRoleAssignments

> RoleAssignedUsers ListUsersWithRoleAssignments(ctx).After(after).Limit(limit).Execute()

List all Users with Role Assignments



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
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 | Specifies the number of results returned. Defaults to `100`. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.ListUsersWithRoleAssignments(context.Background()).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.ListUsersWithRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsersWithRoleAssignments`: RoleAssignedUsers
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.ListUsersWithRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersWithRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** |  | 
 **limit** | **int32** | Specifies the number of results returned. Defaults to &#x60;100&#x60;. | [default to 100]

### Return type

[**RoleAssignedUsers**](RoleAssignedUsers.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveClientRole

> Client RetrieveClientRole(ctx, clientId, roleId).Execute()

Retrieve a Client Role



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAssignmentAPI.RetrieveClientRole(context.Background(), clientId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.RetrieveClientRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveClientRole`: Client
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAPI.RetrieveClientRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveClientRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Client**](Client.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleAssignmentAPI.UnassignRoleFromGroup(context.Background(), groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.UnassignRoleFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleAssignmentAPI.UnassignRoleFromUser(context.Background(), userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAPI.UnassignRoleFromUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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

