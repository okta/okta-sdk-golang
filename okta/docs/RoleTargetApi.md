# \RoleTargetApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAllAppsAsTargetToRoleForUser**](RoleTargetApi.md#AssignAllAppsAsTargetToRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps | Assign all Apps as Target to Role
[**AssignAppInstanceTargetToAppAdminRoleForGroup**](RoleTargetApi.md#AssignAppInstanceTargetToAppAdminRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Assign an Application Instance Target to Application Administrator Role
[**AssignAppInstanceTargetToAppAdminRoleForUser**](RoleTargetApi.md#AssignAppInstanceTargetToAppAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Assign an Application Instance Target to an Application Administrator Role
[**AssignAppTargetToAdminRoleForGroup**](RoleTargetApi.md#AssignAppTargetToAdminRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName} | Assign an Application Target to Administrator Role
[**AssignAppTargetToAdminRoleForUser**](RoleTargetApi.md#AssignAppTargetToAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName} | Assign an Application Target to Administrator Role
[**AssignGroupTargetToGroupAdminRole**](RoleTargetApi.md#AssignGroupTargetToGroupAdminRole) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId} | Assign a Group Target to a Group Role
[**AssignGroupTargetToUserRole**](RoleTargetApi.md#AssignGroupTargetToUserRole) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId} | Assign a Group Target to Role
[**ListApplicationTargetsForApplicationAdministratorRoleForGroup**](RoleTargetApi.md#ListApplicationTargetsForApplicationAdministratorRoleForGroup) | **Get** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps | List all Application Targets for an Application Administrator Role
[**ListApplicationTargetsForApplicationAdministratorRoleForUser**](RoleTargetApi.md#ListApplicationTargetsForApplicationAdministratorRoleForUser) | **Get** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps | List all Application Targets for Application Administrator Role
[**ListGroupTargetsForGroupRole**](RoleTargetApi.md#ListGroupTargetsForGroupRole) | **Get** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups | List all Group Targets for a Group Role
[**ListGroupTargetsForRole**](RoleTargetApi.md#ListGroupTargetsForRole) | **Get** /api/v1/users/{userId}/roles/{roleId}/targets/groups | List all Group Targets for Role
[**UnassignAppInstanceTargetFromAdminRoleForUser**](RoleTargetApi.md#UnassignAppInstanceTargetFromAdminRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Unassign an Application Instance Target from an Application Administrator Role
[**UnassignAppInstanceTargetToAppAdminRoleForGroup**](RoleTargetApi.md#UnassignAppInstanceTargetToAppAdminRoleForGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Unassign an Application Instance Target from an Application Administrator Role
[**UnassignAppTargetFromAppAdminRoleForUser**](RoleTargetApi.md#UnassignAppTargetFromAppAdminRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName} | Unassign an Application Target from an Application Administrator Role
[**UnassignAppTargetToAdminRoleForGroup**](RoleTargetApi.md#UnassignAppTargetToAdminRoleForGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName} | Unassign an Application Target from Application Administrator Role
[**UnassignGroupTargetFromGroupAdminRole**](RoleTargetApi.md#UnassignGroupTargetFromGroupAdminRole) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId} | Unassign a Group Target from a Group Role
[**UnassignGroupTargetFromUserAdminRole**](RoleTargetApi.md#UnassignGroupTargetFromUserAdminRole) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId} | Unassign a Group Target from Role



## AssignAllAppsAsTargetToRoleForUser

> AssignAllAppsAsTargetToRoleForUser(ctx, userId, roleId).Execute()

Assign all Apps as Target to Role



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
    resp, r, err := apiClient.RoleTargetApi.AssignAllAppsAsTargetToRoleForUser(context.Background(), userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignAllAppsAsTargetToRoleForUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAssignAllAppsAsTargetToRoleForUserRequest struct via the builder pattern


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


## AssignAppInstanceTargetToAppAdminRoleForGroup

> AssignAppInstanceTargetToAppAdminRoleForGroup(ctx, groupId, roleId, appName, applicationId).Execute()

Assign an Application Instance Target to Application Administrator Role



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
    appName := "appName_example" // string | 
    applicationId := "applicationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.AssignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleId, appName, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignAppInstanceTargetToAppAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAppInstanceTargetToAppAdminRoleForGroupRequest struct via the builder pattern


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


## AssignAppInstanceTargetToAppAdminRoleForUser

> AssignAppInstanceTargetToAppAdminRoleForUser(ctx, userId, roleId, appName, applicationId).Execute()

Assign an Application Instance Target to an Application Administrator Role



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
    appName := "appName_example" // string | 
    applicationId := "applicationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.AssignAppInstanceTargetToAppAdminRoleForUser(context.Background(), userId, roleId, appName, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignAppInstanceTargetToAppAdminRoleForUser``: %v\n", err)
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
**appName** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAppInstanceTargetToAppAdminRoleForUserRequest struct via the builder pattern


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


## AssignAppTargetToAdminRoleForGroup

> AssignAppTargetToAdminRoleForGroup(ctx, groupId, roleId, appName).Execute()

Assign an Application Target to Administrator Role



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.AssignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignAppTargetToAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAppTargetToAdminRoleForGroupRequest struct via the builder pattern


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


## AssignAppTargetToAdminRoleForUser

> AssignAppTargetToAdminRoleForUser(ctx, userId, roleId, appName).Execute()

Assign an Application Target to Administrator Role



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.AssignAppTargetToAdminRoleForUser(context.Background(), userId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignAppTargetToAdminRoleForUser``: %v\n", err)
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
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAppTargetToAdminRoleForUserRequest struct via the builder pattern


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


## AssignGroupTargetToGroupAdminRole

> AssignGroupTargetToGroupAdminRole(ctx, groupId, roleId, targetGroupId).Execute()

Assign a Group Target to a Group Role



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
    targetGroupId := "targetGroupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.AssignGroupTargetToGroupAdminRole(context.Background(), groupId, roleId, targetGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignGroupTargetToGroupAdminRole``: %v\n", err)
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
**targetGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupTargetToGroupAdminRoleRequest struct via the builder pattern


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


## AssignGroupTargetToUserRole

> AssignGroupTargetToUserRole(ctx, userId, roleId, groupId).Execute()

Assign a Group Target to Role



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.AssignGroupTargetToUserRole(context.Background(), userId, roleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.AssignGroupTargetToUserRole``: %v\n", err)
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
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupTargetToUserRoleRequest struct via the builder pattern


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


## ListApplicationTargetsForApplicationAdministratorRoleForGroup

> []CatalogApplication ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx, groupId, roleId).After(after).Limit(limit).Execute()

List all Application Targets for an Application Administrator Role



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
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.ListApplicationTargetsForApplicationAdministratorRoleForGroup(context.Background(), groupId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.ListApplicationTargetsForApplicationAdministratorRoleForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationTargetsForApplicationAdministratorRoleForGroup`: []CatalogApplication
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetApi.ListApplicationTargetsForApplicationAdministratorRoleForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]CatalogApplication**](CatalogApplication.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTargetsForApplicationAdministratorRoleForUser

> []CatalogApplication ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx, userId, roleId).After(after).Limit(limit).Execute()

List all Application Targets for Application Administrator Role



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
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.ListApplicationTargetsForApplicationAdministratorRoleForUser(context.Background(), userId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.ListApplicationTargetsForApplicationAdministratorRoleForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationTargetsForApplicationAdministratorRoleForUser`: []CatalogApplication
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetApi.ListApplicationTargetsForApplicationAdministratorRoleForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTargetsForApplicationAdministratorRoleForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]CatalogApplication**](CatalogApplication.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupTargetsForGroupRole

> []Group ListGroupTargetsForGroupRole(ctx, groupId, roleId).After(after).Limit(limit).Execute()

List all Group Targets for a Group Role



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
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.ListGroupTargetsForGroupRole(context.Background(), groupId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.ListGroupTargetsForGroupRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupTargetsForGroupRole`: []Group
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetApi.ListGroupTargetsForGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupTargetsForGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]Group**](Group.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupTargetsForRole

> []Group ListGroupTargetsForRole(ctx, userId, roleId).After(after).Limit(limit).Execute()

List all Group Targets for Role



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
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.ListGroupTargetsForRole(context.Background(), userId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.ListGroupTargetsForRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupTargetsForRole`: []Group
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetApi.ListGroupTargetsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupTargetsForRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]Group**](Group.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignAppInstanceTargetFromAdminRoleForUser

> UnassignAppInstanceTargetFromAdminRoleForUser(ctx, userId, roleId, appName, applicationId).Execute()

Unassign an Application Instance Target from an Application Administrator Role



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
    appName := "appName_example" // string | 
    applicationId := "applicationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.UnassignAppInstanceTargetFromAdminRoleForUser(context.Background(), userId, roleId, appName, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.UnassignAppInstanceTargetFromAdminRoleForUser``: %v\n", err)
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
**appName** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignAppInstanceTargetFromAdminRoleForUserRequest struct via the builder pattern


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


## UnassignAppInstanceTargetToAppAdminRoleForGroup

> UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx, groupId, roleId, appName, applicationId).Execute()

Unassign an Application Instance Target from an Application Administrator Role



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
    appName := "appName_example" // string | 
    applicationId := "applicationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.UnassignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleId, appName, applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.UnassignAppInstanceTargetToAppAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest struct via the builder pattern


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


## UnassignAppTargetFromAppAdminRoleForUser

> UnassignAppTargetFromAppAdminRoleForUser(ctx, userId, roleId, appName).Execute()

Unassign an Application Target from an Application Administrator Role



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.UnassignAppTargetFromAppAdminRoleForUser(context.Background(), userId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.UnassignAppTargetFromAppAdminRoleForUser``: %v\n", err)
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
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignAppTargetFromAppAdminRoleForUserRequest struct via the builder pattern


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


## UnassignAppTargetToAdminRoleForGroup

> UnassignAppTargetToAdminRoleForGroup(ctx, groupId, roleId, appName).Execute()

Unassign an Application Target from Application Administrator Role



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.UnassignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.UnassignAppTargetToAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignAppTargetToAdminRoleForGroupRequest struct via the builder pattern


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


## UnassignGroupTargetFromGroupAdminRole

> UnassignGroupTargetFromGroupAdminRole(ctx, groupId, roleId, targetGroupId).Execute()

Unassign a Group Target from a Group Role



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
    targetGroupId := "targetGroupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.UnassignGroupTargetFromGroupAdminRole(context.Background(), groupId, roleId, targetGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.UnassignGroupTargetFromGroupAdminRole``: %v\n", err)
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
**targetGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignGroupTargetFromGroupAdminRoleRequest struct via the builder pattern


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


## UnassignGroupTargetFromUserAdminRole

> UnassignGroupTargetFromUserAdminRole(ctx, userId, roleId, groupId).Execute()

Unassign a Group Target from Role



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetApi.UnassignGroupTargetFromUserAdminRole(context.Background(), userId, roleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetApi.UnassignGroupTargetFromUserAdminRole``: %v\n", err)
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
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignGroupTargetFromUserAdminRoleRequest struct via the builder pattern


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

