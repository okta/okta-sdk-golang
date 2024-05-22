# \RoleAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RoleAPI.md#CreateRole) | **Post** /api/v1/iam/roles | Create a Role
[**CreateRolePermission**](RoleAPI.md#CreateRolePermission) | **Post** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Create a Permission
[**DeleteRole**](RoleAPI.md#DeleteRole) | **Delete** /api/v1/iam/roles/{roleIdOrLabel} | Delete a Role
[**DeleteRolePermission**](RoleAPI.md#DeleteRolePermission) | **Delete** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Delete a Permission
[**GetRole**](RoleAPI.md#GetRole) | **Get** /api/v1/iam/roles/{roleIdOrLabel} | Retrieve a Role
[**GetRolePermission**](RoleAPI.md#GetRolePermission) | **Get** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Retrieve a Permission
[**ListRolePermissions**](RoleAPI.md#ListRolePermissions) | **Get** /api/v1/iam/roles/{roleIdOrLabel}/permissions | List all Permissions
[**ListRoles**](RoleAPI.md#ListRoles) | **Get** /api/v1/iam/roles | List all Roles
[**ReplaceRole**](RoleAPI.md#ReplaceRole) | **Put** /api/v1/iam/roles/{roleIdOrLabel} | Replace a Role
[**ReplaceRolePermission**](RoleAPI.md#ReplaceRolePermission) | **Put** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Replace a Permission



## CreateRole

> IamRole CreateRole(ctx).Instance(instance).Execute()

Create a Role



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
    instance := *openapiclient.NewCreateIamRoleRequest("Description_example", "Label_example", []string{"Permissions_example"}) // CreateIamRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAPI.CreateRole(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: IamRole
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**CreateIamRoleRequest**](CreateIamRoleRequest.md) |  | 

### Return type

[**IamRole**](IamRole.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRolePermission

> CreateRolePermission(ctx, roleIdOrLabel, permissionType).Instance(instance).Execute()

Create a Permission



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    permissionType := "okta.users.manage" // string | An okta permission type
    instance := *openapiclient.NewCreateUpdateIamRolePermissionRequest() // CreateUpdateIamRolePermissionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleAPI.CreateRolePermission(context.Background(), roleIdOrLabel, permissionType).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.CreateRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An okta permission type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRolePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**CreateUpdateIamRolePermissionRequest**](CreateUpdateIamRolePermissionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleIdOrLabel).Execute()

Delete a Role



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleAPI.DeleteRole(context.Background(), roleIdOrLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## DeleteRolePermission

> DeleteRolePermission(ctx, roleIdOrLabel, permissionType).Execute()

Delete a Permission



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    permissionType := "okta.users.manage" // string | An okta permission type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleAPI.DeleteRolePermission(context.Background(), roleIdOrLabel, permissionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.DeleteRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An okta permission type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRolePermissionRequest struct via the builder pattern


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


## GetRole

> IamRole GetRole(ctx, roleIdOrLabel).Execute()

Retrieve a Role



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAPI.GetRole(context.Background(), roleIdOrLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: IamRole
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamRole**](IamRole.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRolePermission

> Permission GetRolePermission(ctx, roleIdOrLabel, permissionType).Execute()

Retrieve a Permission



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    permissionType := "okta.users.manage" // string | An okta permission type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAPI.GetRolePermission(context.Background(), roleIdOrLabel, permissionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.GetRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRolePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.GetRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An okta permission type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Permission**](Permission.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolePermissions

> Permissions ListRolePermissions(ctx, roleIdOrLabel).Execute()

List all Permissions



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAPI.ListRolePermissions(context.Background(), roleIdOrLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.ListRolePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRolePermissions`: Permissions
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.ListRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Permissions**](Permissions.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> IamRoles ListRoles(ctx).After(after).Execute()

List all Roles



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
    resp, r, err := apiClient.RoleAPI.ListRoles(context.Background()).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: IamRoles
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 

### Return type

[**IamRoles**](IamRoles.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRole

> IamRole ReplaceRole(ctx, roleIdOrLabel).Instance(instance).Execute()

Replace a Role



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    instance := *openapiclient.NewUpdateIamRoleRequest("Description_example", "Label_example") // UpdateIamRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAPI.ReplaceRole(context.Background(), roleIdOrLabel).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.ReplaceRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRole`: IamRole
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.ReplaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**UpdateIamRoleRequest**](UpdateIamRoleRequest.md) |  | 

### Return type

[**IamRole**](IamRole.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRolePermission

> Permission ReplaceRolePermission(ctx, roleIdOrLabel, permissionType).Instance(instance).Execute()

Replace a Permission



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
    roleIdOrLabel := "cr0Yq6IJxGIr0ouum0g3" // string | `id` or `label` of the role
    permissionType := "okta.users.manage" // string | An okta permission type
    instance := *openapiclient.NewCreateUpdateIamRolePermissionRequest() // CreateUpdateIamRolePermissionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleAPI.ReplaceRolePermission(context.Background(), roleIdOrLabel, permissionType).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAPI.ReplaceRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRolePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `RoleAPI.ReplaceRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An okta permission type | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRolePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**CreateUpdateIamRolePermissionRequest**](CreateUpdateIamRolePermissionRequest.md) |  | 

### Return type

[**Permission**](Permission.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

