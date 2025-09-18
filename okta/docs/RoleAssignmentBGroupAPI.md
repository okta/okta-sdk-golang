# \RoleAssignmentBGroupAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToGroup**](RoleAssignmentBGroupAPI.md#AssignRoleToGroup) | **Post** /api/v1/groups/{groupId}/roles | Assign a role to a group
[**GetGroupAssignedRole**](RoleAssignmentBGroupAPI.md#GetGroupAssignedRole) | **Get** /api/v1/groups/{groupId}/roles/{roleAssignmentId} | Retrieve a group role assignment
[**ListGroupAssignedRoles**](RoleAssignmentBGroupAPI.md#ListGroupAssignedRoles) | **Get** /api/v1/groups/{groupId}/roles | List all group role assignments
[**UnassignRoleFromGroup**](RoleAssignmentBGroupAPI.md#UnassignRoleFromGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleAssignmentId} | Unassign a group role



## AssignRoleToGroup

> ListGroupAssignedRoles200ResponseInner AssignRoleToGroup(ctx, groupId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()

Assign a role to a group



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
	assignRoleRequest := openapiclient.assignRoleToGroup_request{CustomRoleAssignmentSchema: openapiclient.NewCustomRoleAssignmentSchema()} // AssignRoleToGroupRequest | 
	disableNotifications := true // bool | Grants the group third-party admin status when set to `true` (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentBGroupAPI.AssignRoleToGroup(context.Background(), groupId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentBGroupAPI.AssignRoleToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRoleToGroup`: ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentBGroupAPI.AssignRoleToGroup`: %v\n", resp)
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

 **assignRoleRequest** | [**AssignRoleToGroupRequest**](AssignRoleToGroupRequest.md) |  | 
 **disableNotifications** | **bool** | Grants the group third-party admin status when set to &#x60;true&#x60; | [default to false]

### Return type

[**ListGroupAssignedRoles200ResponseInner**](ListGroupAssignedRoles200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAssignedRole

> ListGroupAssignedRoles200ResponseInner GetGroupAssignedRole(ctx, groupId, roleAssignmentId).Execute()

Retrieve a group role assignment



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentBGroupAPI.GetGroupAssignedRole(context.Background(), groupId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentBGroupAPI.GetGroupAssignedRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupAssignedRole`: ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentBGroupAPI.GetGroupAssignedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAssignedRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListGroupAssignedRoles200ResponseInner**](ListGroupAssignedRoles200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupAssignedRoles

> []ListGroupAssignedRoles200ResponseInner ListGroupAssignedRoles(ctx, groupId).Expand(expand).Execute()

List all group role assignments



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
	expand := "targets/groups" // string | An optional parameter used to return targets configured for the standard role assignment in the `embedded` property. Supported values: `targets/groups` or `targets/catalog/apps` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentBGroupAPI.ListGroupAssignedRoles(context.Background(), groupId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentBGroupAPI.ListGroupAssignedRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupAssignedRoles`: []ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentBGroupAPI.ListGroupAssignedRoles`: %v\n", resp)
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

 **expand** | **string** | An optional parameter used to return targets configured for the standard role assignment in the &#x60;embedded&#x60; property. Supported values: &#x60;targets/groups&#x60; or &#x60;targets/catalog/apps&#x60; | 

### Return type

[**[]ListGroupAssignedRoles200ResponseInner**](ListGroupAssignedRoles200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignRoleFromGroup

> UnassignRoleFromGroup(ctx, groupId, roleAssignmentId).Execute()

Unassign a group role



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleAssignmentBGroupAPI.UnassignRoleFromGroup(context.Background(), groupId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentBGroupAPI.UnassignRoleFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

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

