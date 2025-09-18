# \RoleAssignmentClientAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToClient**](RoleAssignmentClientAPI.md#AssignRoleToClient) | **Post** /oauth2/v1/clients/{clientId}/roles | Assign a client role
[**DeleteRoleFromClient**](RoleAssignmentClientAPI.md#DeleteRoleFromClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId} | Unassign a client role
[**ListRolesForClient**](RoleAssignmentClientAPI.md#ListRolesForClient) | **Get** /oauth2/v1/clients/{clientId}/roles | List all client role assignments
[**RetrieveClientRole**](RoleAssignmentClientAPI.md#RetrieveClientRole) | **Get** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId} | Retrieve a client role



## AssignRoleToClient

> ListGroupAssignedRoles200ResponseInner AssignRoleToClient(ctx, clientId).AssignRoleToGroupRequest(assignRoleToGroupRequest).Execute()

Assign a client role



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID
	assignRoleToGroupRequest := openapiclient.assignRoleToGroup_request{CustomRoleAssignmentSchema: openapiclient.NewCustomRoleAssignmentSchema()} // AssignRoleToGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentClientAPI.AssignRoleToClient(context.Background(), clientId).AssignRoleToGroupRequest(assignRoleToGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentClientAPI.AssignRoleToClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRoleToClient`: ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentClientAPI.AssignRoleToClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | Client app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRoleToClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignRoleToGroupRequest** | [**AssignRoleToGroupRequest**](AssignRoleToGroupRequest.md) |  | 

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


## DeleteRoleFromClient

> DeleteRoleFromClient(ctx, clientId, roleAssignmentId).Execute()

Unassign a client role



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleAssignmentClientAPI.DeleteRoleFromClient(context.Background(), clientId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentClientAPI.DeleteRoleFromClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | Client app ID | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

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


## ListRolesForClient

> ListGroupAssignedRoles200ResponseInner ListRolesForClient(ctx, clientId).Execute()

List all client role assignments



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentClientAPI.ListRolesForClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentClientAPI.ListRolesForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRolesForClient`: ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentClientAPI.ListRolesForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | Client app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesForClientRequest struct via the builder pattern


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


## RetrieveClientRole

> ListGroupAssignedRoles200ResponseInner RetrieveClientRole(ctx, clientId, roleAssignmentId).Execute()

Retrieve a client role



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleAssignmentClientAPI.RetrieveClientRole(context.Background(), clientId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentClientAPI.RetrieveClientRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveClientRole`: ListGroupAssignedRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentClientAPI.RetrieveClientRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | Client app ID | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveClientRoleRequest struct via the builder pattern


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

