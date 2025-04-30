# \RoleTargetAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAllAppsAsTargetToRoleForUser**](RoleTargetAPI.md#AssignAllAppsAsTargetToRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps | Assign all Apps as Target to Role
[**AssignAppInstanceTargetToAppAdminRoleForGroup**](RoleTargetAPI.md#AssignAppInstanceTargetToAppAdminRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId} | Assign an Application Instance Target to Application Administrator Role
[**AssignAppInstanceTargetToAppAdminRoleForUser**](RoleTargetAPI.md#AssignAppInstanceTargetToAppAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId} | Assign an Application Instance Target to an Application Administrator Role
[**AssignAppTargetInstanceRoleForClient**](RoleTargetAPI.md#AssignAppTargetInstanceRoleForClient) | **Put** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName}/{appInstanceId} | Replace an App Instance Target for a Client
[**AssignAppTargetRoleToClient**](RoleTargetAPI.md#AssignAppTargetRoleToClient) | **Put** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName} | Assign an App Target to a Client
[**AssignAppTargetToAdminRoleForGroup**](RoleTargetAPI.md#AssignAppTargetToAdminRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName} | Assign an Application Target to Administrator Role
[**AssignAppTargetToAdminRoleForUser**](RoleTargetAPI.md#AssignAppTargetToAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName} | Assign an Application Target to Administrator Role
[**AssignGroupTargetRoleForClient**](RoleTargetAPI.md#AssignGroupTargetRoleForClient) | **Put** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/groups/{groupId} | Assign a Group Target to a Client
[**AssignGroupTargetToGroupAdminRole**](RoleTargetAPI.md#AssignGroupTargetToGroupAdminRole) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId} | Assign a Group Target to a Group Role
[**AssignGroupTargetToUserRole**](RoleTargetAPI.md#AssignGroupTargetToUserRole) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId} | Assign a Group Target to Role
[**ListAppTargetRoleToClient**](RoleTargetAPI.md#ListAppTargetRoleToClient) | **Get** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps | List all App Targets assigned to a Client
[**ListApplicationTargetsForApplicationAdministratorRoleForGroup**](RoleTargetAPI.md#ListApplicationTargetsForApplicationAdministratorRoleForGroup) | **Get** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps | List all Application Targets for an Application Administrator Role
[**ListApplicationTargetsForApplicationAdministratorRoleForUser**](RoleTargetAPI.md#ListApplicationTargetsForApplicationAdministratorRoleForUser) | **Get** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps | List all Application Targets for Application Administrator Role
[**ListGroupTargetRoleForClient**](RoleTargetAPI.md#ListGroupTargetRoleForClient) | **Get** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/groups | List all Group Targets for a Client
[**ListGroupTargetsForGroupRole**](RoleTargetAPI.md#ListGroupTargetsForGroupRole) | **Get** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups | List all Group Targets for a Group Role
[**ListGroupTargetsForRole**](RoleTargetAPI.md#ListGroupTargetsForRole) | **Get** /api/v1/users/{userId}/roles/{roleId}/targets/groups | List all Group Targets for Role
[**RemoveAppTargetInstanceRoleForClient**](RoleTargetAPI.md#RemoveAppTargetInstanceRoleForClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName}/{appInstanceId} | Delete an App Instance Target for a Client
[**RemoveAppTargetRoleFromClient**](RoleTargetAPI.md#RemoveAppTargetRoleFromClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName} | Remove an App Target from a Client
[**RemoveGroupTargetRoleFromClient**](RoleTargetAPI.md#RemoveGroupTargetRoleFromClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleId}/targets/groups/{groupId} | Delete a Group Target from a Client
[**UnassignAppInstanceTargetFromAdminRoleForUser**](RoleTargetAPI.md#UnassignAppInstanceTargetFromAdminRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId} | Unassign an Application Instance Target from an Application Administrator Role
[**UnassignAppInstanceTargetToAppAdminRoleForGroup**](RoleTargetAPI.md#UnassignAppInstanceTargetToAppAdminRoleForGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId} | Unassign an Application Instance Target from an Application Administrator Role
[**UnassignAppTargetFromAppAdminRoleForUser**](RoleTargetAPI.md#UnassignAppTargetFromAppAdminRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName} | Unassign an Application Target from an Application Administrator Role
[**UnassignAppTargetToAdminRoleForGroup**](RoleTargetAPI.md#UnassignAppTargetToAdminRoleForGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName} | Unassign an Application Target from Application Administrator Role
[**UnassignGroupTargetFromGroupAdminRole**](RoleTargetAPI.md#UnassignGroupTargetFromGroupAdminRole) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId} | Unassign a Group Target from a Group Role
[**UnassignGroupTargetFromUserAdminRole**](RoleTargetAPI.md#UnassignGroupTargetFromUserAdminRole) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId} | Unassign a Group Target from Role



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAllAppsAsTargetToRoleForUser(context.Background(), userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAllAppsAsTargetToRoleForUser``: %v\n", err)
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

> AssignAppInstanceTargetToAppAdminRoleForGroup(ctx, groupId, roleId, appName, appId).Execute()

Assign an Application Instance Target to Application Administrator Role



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
    appName := "oidc_client" // string | Application name for the app type
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleId, appName, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAppInstanceTargetToAppAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 
**appId** | **string** | Application ID | 

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

> AssignAppInstanceTargetToAppAdminRoleForUser(ctx, userId, roleId, appName, appId).Execute()

Assign an Application Instance Target to an Application Administrator Role



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
    appName := "oidc_client" // string | Application name for the app type
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAppInstanceTargetToAppAdminRoleForUser(context.Background(), userId, roleId, appName, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAppInstanceTargetToAppAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 
**appId** | **string** | Application ID | 

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


## AssignAppTargetInstanceRoleForClient

> AssignAppTargetInstanceRoleForClient(ctx, clientId, roleId, appName, appInstanceId).Execute()

Replace an App Instance Target for a Client



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
    appName := "oidc_client" // string | Application name for the app type
    appInstanceId := "appInstanceId_example" // string | `id` of the application instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAppTargetInstanceRoleForClient(context.Background(), clientId, roleId, appName, appInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAppTargetInstanceRoleForClient``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 
**appInstanceId** | **string** | &#x60;id&#x60; of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAppTargetInstanceRoleForClientRequest struct via the builder pattern


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


## AssignAppTargetRoleToClient

> AssignAppTargetRoleToClient(ctx, clientId, roleId, appName).Execute()

Assign an App Target to a Client



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
    appName := "oidc_client" // string | Application name for the app type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAppTargetRoleToClient(context.Background(), clientId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAppTargetRoleToClient``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAppTargetRoleToClientRequest struct via the builder pattern


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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    appName := "oidc_client" // string | Application name for the app type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAppTargetToAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    appName := "oidc_client" // string | Application name for the app type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignAppTargetToAdminRoleForUser(context.Background(), userId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignAppTargetToAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 

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


## AssignGroupTargetRoleForClient

> AssignGroupTargetRoleForClient(ctx, clientId, roleId, groupId).Execute()

Assign a Group Target to a Client



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
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignGroupTargetRoleForClient(context.Background(), clientId, roleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignGroupTargetRoleForClient``: %v\n", err)
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
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupTargetRoleForClientRequest struct via the builder pattern


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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    targetGroupId := "00g1e9dfjHeLAsdX983d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignGroupTargetToGroupAdminRole(context.Background(), groupId, roleId, targetGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignGroupTargetToGroupAdminRole``: %v\n", err)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.AssignGroupTargetToUserRole(context.Background(), userId, roleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.AssignGroupTargetToUserRole``: %v\n", err)
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
**groupId** | **string** | The &#x60;id&#x60; of the group | 

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


## ListAppTargetRoleToClient

> Client ListAppTargetRoleToClient(ctx, clientId, roleId).After(after).Limit(limit).Execute()

List all App Targets assigned to a Client



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
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetAPI.ListAppTargetRoleToClient(context.Background(), clientId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.ListAppTargetRoleToClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppTargetRoleToClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetAPI.ListAppTargetRoleToClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppTargetRoleToClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup(context.Background(), groupId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationTargetsForApplicationAdministratorRoleForGroup`: []CatalogApplication
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser(context.Background(), userId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationTargetsForApplicationAdministratorRoleForUser`: []CatalogApplication
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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


## ListGroupTargetRoleForClient

> Client ListGroupTargetRoleForClient(ctx, clientId, roleId).After(after).Limit(limit).Execute()

List all Group Targets for a Client



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
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetAPI.ListGroupTargetRoleForClient(context.Background(), clientId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.ListGroupTargetRoleForClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupTargetRoleForClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetAPI.ListGroupTargetRoleForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupTargetRoleForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetAPI.ListGroupTargetsForGroupRole(context.Background(), groupId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.ListGroupTargetsForGroupRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupTargetsForGroupRole`: []Group
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetAPI.ListGroupTargetsForGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleTargetAPI.ListGroupTargetsForRole(context.Background(), userId, roleId).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.ListGroupTargetsForRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupTargetsForRole`: []Group
    fmt.Fprintf(os.Stdout, "Response from `RoleTargetAPI.ListGroupTargetsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleId** | **string** | &#x60;id&#x60; of the Role | 

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


## RemoveAppTargetInstanceRoleForClient

> RemoveAppTargetInstanceRoleForClient(ctx, clientId, roleId, appName, appInstanceId).Execute()

Delete an App Instance Target for a Client



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
    appName := "oidc_client" // string | Application name for the app type
    appInstanceId := "appInstanceId_example" // string | `id` of the application instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.RemoveAppTargetInstanceRoleForClient(context.Background(), clientId, roleId, appName, appInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.RemoveAppTargetInstanceRoleForClient``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 
**appInstanceId** | **string** | &#x60;id&#x60; of the application instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAppTargetInstanceRoleForClientRequest struct via the builder pattern


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


## RemoveAppTargetRoleFromClient

> RemoveAppTargetRoleFromClient(ctx, clientId, roleId, appName).Execute()

Remove an App Target from a Client



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
    appName := "oidc_client" // string | Application name for the app type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.RemoveAppTargetRoleFromClient(context.Background(), clientId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.RemoveAppTargetRoleFromClient``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAppTargetRoleFromClientRequest struct via the builder pattern


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


## RemoveGroupTargetRoleFromClient

> RemoveGroupTargetRoleFromClient(ctx, clientId, roleId, groupId).Execute()

Delete a Group Target from a Client



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
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.RemoveGroupTargetRoleFromClient(context.Background(), clientId, roleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.RemoveGroupTargetRoleFromClient``: %v\n", err)
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
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupTargetRoleFromClientRequest struct via the builder pattern


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


## UnassignAppInstanceTargetFromAdminRoleForUser

> UnassignAppInstanceTargetFromAdminRoleForUser(ctx, userId, roleId, appName, appId).Execute()

Unassign an Application Instance Target from an Application Administrator Role



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
    appName := "oidc_client" // string | Application name for the app type
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.UnassignAppInstanceTargetFromAdminRoleForUser(context.Background(), userId, roleId, appName, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.UnassignAppInstanceTargetFromAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 
**appId** | **string** | Application ID | 

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

> UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx, groupId, roleId, appName, appId).Execute()

Unassign an Application Instance Target from an Application Administrator Role



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
    appName := "oidc_client" // string | Application name for the app type
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.UnassignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleId, appName, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.UnassignAppInstanceTargetToAppAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 
**appId** | **string** | Application ID | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    appName := "oidc_client" // string | Application name for the app type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.UnassignAppTargetFromAppAdminRoleForUser(context.Background(), userId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.UnassignAppTargetFromAppAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    appName := "oidc_client" // string | Application name for the app type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.UnassignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleId, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.UnassignAppTargetToAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Application name for the app type | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    targetGroupId := "00g1e9dfjHeLAsdX983d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.UnassignGroupTargetFromGroupAdminRole(context.Background(), groupId, roleId, targetGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.UnassignGroupTargetFromGroupAdminRole``: %v\n", err)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    userId := "userId_example" // string | ID of an existing Okta user
    roleId := "3Vg1Pjp3qzw4qcCK5EdO" // string | `id` of the Role
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleTargetAPI.UnassignGroupTargetFromUserAdminRole(context.Background(), userId, roleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleTargetAPI.UnassignGroupTargetFromUserAdminRole``: %v\n", err)
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
**groupId** | **string** | The &#x60;id&#x60; of the group | 

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

