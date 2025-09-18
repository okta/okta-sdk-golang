# \RoleBTargetBGroupAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAppInstanceTargetToAppAdminRoleForGroup**](RoleBTargetBGroupAPI.md#AssignAppInstanceTargetToAppAdminRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId} | Assign a group role app instance target
[**AssignAppTargetToAdminRoleForGroup**](RoleBTargetBGroupAPI.md#AssignAppTargetToAdminRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName} | Assign a group role app target
[**AssignGroupTargetToGroupAdminRole**](RoleBTargetBGroupAPI.md#AssignGroupTargetToGroupAdminRole) | **Put** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/groups/{targetGroupId} | Assign a group role group target
[**ListApplicationTargetsForApplicationAdministratorRoleForGroup**](RoleBTargetBGroupAPI.md#ListApplicationTargetsForApplicationAdministratorRoleForGroup) | **Get** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps | List all group role app targets
[**ListGroupTargetsForGroupRole**](RoleBTargetBGroupAPI.md#ListGroupTargetsForGroupRole) | **Get** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/groups | List all group role group targets
[**UnassignAppInstanceTargetToAppAdminRoleForGroup**](RoleBTargetBGroupAPI.md#UnassignAppInstanceTargetToAppAdminRoleForGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId} | Unassign a group role app instance target
[**UnassignAppTargetToAdminRoleForGroup**](RoleBTargetBGroupAPI.md#UnassignAppTargetToAdminRoleForGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName} | Unassign a group role app target
[**UnassignGroupTargetFromGroupAdminRole**](RoleBTargetBGroupAPI.md#UnassignGroupTargetFromGroupAdminRole) | **Delete** /api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/groups/{targetGroupId} | Unassign a group role group target



## AssignAppInstanceTargetToAppAdminRoleForGroup

> AssignAppInstanceTargetToAppAdminRoleForGroup(ctx, groupId, roleAssignmentId, appName, appId).Execute()

Assign a group role app instance target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetBGroupAPI.AssignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.AssignAppInstanceTargetToAppAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 
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


## AssignAppTargetToAdminRoleForGroup

> AssignAppTargetToAdminRoleForGroup(ctx, groupId, roleAssignmentId, appName).Execute()

Assign a group role app target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetBGroupAPI.AssignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.AssignAppTargetToAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 

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


## AssignGroupTargetToGroupAdminRole

> AssignGroupTargetToGroupAdminRole(ctx, groupId, roleAssignmentId, targetGroupId).Execute()

Assign a group role group target



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
	targetGroupId := "00g1e9dfjHeLAsdX983d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetBGroupAPI.AssignGroupTargetToGroupAdminRole(context.Background(), groupId, roleAssignmentId, targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.AssignGroupTargetToGroupAdminRole``: %v\n", err)
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


## ListApplicationTargetsForApplicationAdministratorRoleForGroup

> []CatalogApplication ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx, groupId, roleAssignmentId).After(after).Limit(limit).Execute()

List all group role app targets



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetBGroupAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup(context.Background(), groupId, roleAssignmentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationTargetsForApplicationAdministratorRoleForGroup`: []CatalogApplication
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetBGroupAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

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

> []Group ListGroupTargetsForGroupRole(ctx, groupId, roleAssignmentId).After(after).Limit(limit).Execute()

List all group role group targets



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetBGroupAPI.ListGroupTargetsForGroupRole(context.Background(), groupId, roleAssignmentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.ListGroupTargetsForGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupTargetsForGroupRole`: []Group
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetBGroupAPI.ListGroupTargetsForGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupTargetsForGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

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


## UnassignAppInstanceTargetToAppAdminRoleForGroup

> UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx, groupId, roleAssignmentId, appName, appId).Execute()

Unassign a group role app instance target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetBGroupAPI.UnassignAppInstanceTargetToAppAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.UnassignAppInstanceTargetToAppAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 
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


## UnassignAppTargetToAdminRoleForGroup

> UnassignAppTargetToAdminRoleForGroup(ctx, groupId, roleAssignmentId, appName).Execute()

Unassign a group role app target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetBGroupAPI.UnassignAppTargetToAdminRoleForGroup(context.Background(), groupId, roleAssignmentId, appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.UnassignAppTargetToAdminRoleForGroup``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 

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

> UnassignGroupTargetFromGroupAdminRole(ctx, groupId, roleAssignmentId, targetGroupId).Execute()

Unassign a group role group target



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
	targetGroupId := "00g1e9dfjHeLAsdX983d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetBGroupAPI.UnassignGroupTargetFromGroupAdminRole(context.Background(), groupId, roleAssignmentId, targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetBGroupAPI.UnassignGroupTargetFromGroupAdminRole``: %v\n", err)
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

