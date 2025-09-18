# \RoleAssignmentAUserAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToUser**](RoleAssignmentAUserAPI.md#AssignRoleToUser) | **Post** /api/v1/users/{userId}/roles | Assign a user role
[**GetRoleAssignmentGovernanceGrant**](RoleAssignmentAUserAPI.md#GetRoleAssignmentGovernanceGrant) | **Get** /api/v1/users/{userId}/roles/{roleAssignmentId}/governance/{grantId} | Retrieve a user role governance source
[**GetRoleAssignmentGovernanceGrantResources**](RoleAssignmentAUserAPI.md#GetRoleAssignmentGovernanceGrantResources) | **Get** /api/v1/users/{userId}/roles/{roleAssignmentId}/governance/{grantId}/resources | Retrieve the user role governance source resources
[**GetUserAssignedRole**](RoleAssignmentAUserAPI.md#GetUserAssignedRole) | **Get** /api/v1/users/{userId}/roles/{roleAssignmentId} | Retrieve a user role assignment
[**GetUserAssignedRoleGovernance**](RoleAssignmentAUserAPI.md#GetUserAssignedRoleGovernance) | **Get** /api/v1/users/{userId}/roles/{roleAssignmentId}/governance | Retrieve all user role governance sources
[**ListAssignedRolesForUser**](RoleAssignmentAUserAPI.md#ListAssignedRolesForUser) | **Get** /api/v1/users/{userId}/roles | List all user role assignments
[**ListUsersWithRoleAssignments**](RoleAssignmentAUserAPI.md#ListUsersWithRoleAssignments) | **Get** /api/v1/iam/assignees/users | List all users with role assignments
[**UnassignRoleFromUser**](RoleAssignmentAUserAPI.md#UnassignRoleFromUser) | **Delete** /api/v1/users/{userId}/roles/{roleAssignmentId} | Unassign a user role



## AssignRoleToUser

> AssignRoleToUser201Response AssignRoleToUser(ctx, userId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()

Assign a user role



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	assignRoleRequest := openapiclient.assignRoleToUser_request{CustomRoleAssignmentSchema: openapiclient.NewCustomRoleAssignmentSchema()} // AssignRoleToUserRequest | 
	disableNotifications := true // bool | Setting this to `true` grants the user third-party admin status (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.AssignRoleToUser(context.Background(), userId).AssignRoleRequest(assignRoleRequest).DisableNotifications(disableNotifications).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.AssignRoleToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRoleToUser`: AssignRoleToUser201Response
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.AssignRoleToUser`: %v\n", resp)
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

 **assignRoleRequest** | [**AssignRoleToUserRequest**](AssignRoleToUserRequest.md) |  | 
 **disableNotifications** | **bool** | Setting this to &#x60;true&#x60; grants the user third-party admin status | [default to false]

### Return type

[**AssignRoleToUser201Response**](AssignRoleToUser201Response.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAssignmentGovernanceGrant

> RoleGovernanceSource GetRoleAssignmentGovernanceGrant(ctx, userId, roleAssignmentId, grantId).Execute()

Retrieve a user role governance source



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.GetRoleAssignmentGovernanceGrant(context.Background(), userId, roleAssignmentId, grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.GetRoleAssignmentGovernanceGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleAssignmentGovernanceGrant`: RoleGovernanceSource
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.GetRoleAssignmentGovernanceGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAssignmentGovernanceGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleGovernanceSource**](RoleGovernanceSource.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAssignmentGovernanceGrantResources

> RoleGovernanceResources GetRoleAssignmentGovernanceGrantResources(ctx, userId, roleAssignmentId, grantId).Execute()

Retrieve the user role governance source resources



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.GetRoleAssignmentGovernanceGrantResources(context.Background(), userId, roleAssignmentId, grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.GetRoleAssignmentGovernanceGrantResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleAssignmentGovernanceGrantResources`: RoleGovernanceResources
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.GetRoleAssignmentGovernanceGrantResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAssignmentGovernanceGrantResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleGovernanceResources**](RoleGovernanceResources.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAssignedRole

> ListGroupAssignedRoles200ResponseInner GetUserAssignedRole(ctx, userId, roleAssignmentId).Execute()

Retrieve a user role assignment



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.GetUserAssignedRole(context.Background(), userId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.GetUserAssignedRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAssignedRole`: ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.GetUserAssignedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAssignedRoleRequest struct via the builder pattern


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


## GetUserAssignedRoleGovernance

> RoleGovernance GetUserAssignedRoleGovernance(ctx, userId, roleAssignmentId).Execute()

Retrieve all user role governance sources



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.GetUserAssignedRoleGovernance(context.Background(), userId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.GetUserAssignedRoleGovernance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAssignedRoleGovernance`: RoleGovernance
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.GetUserAssignedRoleGovernance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAssignedRoleGovernanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleGovernance**](RoleGovernance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssignedRolesForUser

> []ListGroupAssignedRoles200ResponseInner ListAssignedRolesForUser(ctx, userId).Expand(expand).Execute()

List all user role assignments



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	expand := "targets/groups" // string | An optional parameter used to return targets configured for the standard role assignment in the `embedded` property. Supported values: `targets/groups` or `targets/catalog/apps` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.ListAssignedRolesForUser(context.Background(), userId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.ListAssignedRolesForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssignedRolesForUser`: []ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.ListAssignedRolesForUser`: %v\n", resp)
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


## ListUsersWithRoleAssignments

> RoleAssignedUsers ListUsersWithRoleAssignments(ctx).After(after).Limit(limit).Execute()

List all users with role assignments



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
	after := "after_example" // string | Specifies the pagination cursor for the next page of targets (optional)
	limit := int32(56) // int32 | Specifies the number of results returned. Defaults to `100`. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentAUserAPI.ListUsersWithRoleAssignments(context.Background()).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.ListUsersWithRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersWithRoleAssignments`: RoleAssignedUsers
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentAUserAPI.ListUsersWithRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersWithRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Specifies the pagination cursor for the next page of targets | 
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


## UnassignRoleFromUser

> UnassignRoleFromUser(ctx, userId, roleAssignmentId).Execute()

Unassign a user role



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleAssignmentAUserAPI.UnassignRoleFromUser(context.Background(), userId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentAUserAPI.UnassignRoleFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

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

