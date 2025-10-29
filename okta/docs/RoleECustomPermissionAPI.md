# \RoleECustomPermissionAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRolePermission**](RoleECustomPermissionAPI.md#CreateRolePermission) | **Post** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Create a custom role permission
[**DeleteRolePermission**](RoleECustomPermissionAPI.md#DeleteRolePermission) | **Delete** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Delete a custom role permission
[**GetRolePermission**](RoleECustomPermissionAPI.md#GetRolePermission) | **Get** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Retrieve a custom role permission
[**ListRolePermissions**](RoleECustomPermissionAPI.md#ListRolePermissions) | **Get** /api/v1/iam/roles/{roleIdOrLabel}/permissions | List all custom role permissions
[**ReplaceRolePermission**](RoleECustomPermissionAPI.md#ReplaceRolePermission) | **Put** /api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType} | Replace a custom role permission



## CreateRolePermission

> CreateRolePermission(ctx, roleIdOrLabel, permissionType).Instance(instance).Execute()

Create a custom role permission



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
	permissionType := "okta.users.manage" // string | An Okta [permission](/openapi/okta-management/guides/permissions)
	instance := *openapiclient.NewCreateUpdateIamRolePermissionRequest() // CreateUpdateIamRolePermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleECustomPermissionAPI.CreateRolePermission(context.Background(), roleIdOrLabel, permissionType).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleECustomPermissionAPI.CreateRolePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An Okta [permission](/openapi/okta-management/guides/permissions) | 

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


## DeleteRolePermission

> DeleteRolePermission(ctx, roleIdOrLabel, permissionType).Execute()

Delete a custom role permission



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
	permissionType := "okta.users.manage" // string | An Okta [permission](/openapi/okta-management/guides/permissions)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleECustomPermissionAPI.DeleteRolePermission(context.Background(), roleIdOrLabel, permissionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleECustomPermissionAPI.DeleteRolePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An Okta [permission](/openapi/okta-management/guides/permissions) | 

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


## GetRolePermission

> Permission GetRolePermission(ctx, roleIdOrLabel, permissionType).Execute()

Retrieve a custom role permission



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
	permissionType := "okta.users.manage" // string | An Okta [permission](/openapi/okta-management/guides/permissions)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleECustomPermissionAPI.GetRolePermission(context.Background(), roleIdOrLabel, permissionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleECustomPermissionAPI.GetRolePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolePermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `RoleECustomPermissionAPI.GetRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An Okta [permission](/openapi/okta-management/guides/permissions) | 

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

List all custom role permissions



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
	resp, r, err := apiClient.RoleECustomPermissionAPI.ListRolePermissions(context.Background(), roleIdOrLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleECustomPermissionAPI.ListRolePermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRolePermissions`: Permissions
	fmt.Fprintf(os.Stdout, "Response from `RoleECustomPermissionAPI.ListRolePermissions`: %v\n", resp)
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


## ReplaceRolePermission

> Permission ReplaceRolePermission(ctx, roleIdOrLabel, permissionType).Instance(instance).Execute()

Replace a custom role permission



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
	permissionType := "okta.users.manage" // string | An Okta [permission](/openapi/okta-management/guides/permissions)
	instance := *openapiclient.NewCreateUpdateIamRolePermissionRequest() // CreateUpdateIamRolePermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleECustomPermissionAPI.ReplaceRolePermission(context.Background(), roleIdOrLabel, permissionType).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleECustomPermissionAPI.ReplaceRolePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRolePermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `RoleECustomPermissionAPI.ReplaceRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIdOrLabel** | **string** | &#x60;id&#x60; or &#x60;label&#x60; of the role | 
**permissionType** | **string** | An Okta [permission](/openapi/okta-management/guides/permissions) | 

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

