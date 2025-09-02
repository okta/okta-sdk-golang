# \RoleBTargetAdminAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAllAppsAsTargetToRoleForUser**](RoleBTargetAdminAPI.md#AssignAllAppsAsTargetToRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps | Assign all apps as target to admin role
[**AssignAppInstanceTargetToAppAdminRoleForUser**](RoleBTargetAdminAPI.md#AssignAppInstanceTargetToAppAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId} | Assign an admin role app instance target
[**AssignAppTargetToAdminRoleForUser**](RoleBTargetAdminAPI.md#AssignAppTargetToAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName} | Assign an admin role app target
[**AssignGroupTargetToUserRole**](RoleBTargetAdminAPI.md#AssignGroupTargetToUserRole) | **Put** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/groups/{groupId} | Assign an admin role group target
[**GetRoleTargetsByUserIdAndRoleId**](RoleBTargetAdminAPI.md#GetRoleTargetsByUserIdAndRoleId) | **Get** /api/v1/users/{userId}/roles/{roleIdOrEncodedRoleId}/targets | Retrieve a role target by assignment type
[**ListApplicationTargetsForApplicationAdministratorRoleForUser**](RoleBTargetAdminAPI.md#ListApplicationTargetsForApplicationAdministratorRoleForUser) | **Get** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps | List all admin role app targets
[**ListGroupTargetsForRole**](RoleBTargetAdminAPI.md#ListGroupTargetsForRole) | **Get** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/groups | List all admin role group targets
[**UnassignAppInstanceTargetFromAdminRoleForUser**](RoleBTargetAdminAPI.md#UnassignAppInstanceTargetFromAdminRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId} | Unassign an admin role app instance target
[**UnassignAppTargetFromAppAdminRoleForUser**](RoleBTargetAdminAPI.md#UnassignAppTargetFromAppAdminRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName} | Unassign an admin role app target
[**UnassignGroupTargetFromUserAdminRole**](RoleBTargetAdminAPI.md#UnassignGroupTargetFromUserAdminRole) | **Delete** /api/v1/users/{userId}/roles/{roleAssignmentId}/targets/groups/{groupId} | Unassign an admin role group target



## AssignAllAppsAsTargetToRoleForUser

> AssignAllAppsAsTargetToRoleForUser(ctx, userId, roleAssignmentId).Execute()

Assign all apps as target to admin role



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
	r, err := apiClient.RoleBTargetAdminAPI.AssignAllAppsAsTargetToRoleForUser(context.Background(), userId, roleAssignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.AssignAllAppsAsTargetToRoleForUser``: %v\n", err)
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


## AssignAppInstanceTargetToAppAdminRoleForUser

> AssignAppInstanceTargetToAppAdminRoleForUser(ctx, userId, roleAssignmentId, appName, appId).Execute()

Assign an admin role app instance target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetAdminAPI.AssignAppInstanceTargetToAppAdminRoleForUser(context.Background(), userId, roleAssignmentId, appName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.AssignAppInstanceTargetToAppAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 
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


## AssignAppTargetToAdminRoleForUser

> AssignAppTargetToAdminRoleForUser(ctx, userId, roleAssignmentId, appName).Execute()

Assign an admin role app target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetAdminAPI.AssignAppTargetToAdminRoleForUser(context.Background(), userId, roleAssignmentId, appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.AssignAppTargetToAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 

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


## AssignGroupTargetToUserRole

> AssignGroupTargetToUserRole(ctx, userId, roleAssignmentId, groupId).Execute()

Assign an admin role group target



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
	groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetAdminAPI.AssignGroupTargetToUserRole(context.Background(), userId, roleAssignmentId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.AssignGroupTargetToUserRole``: %v\n", err)
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


## GetRoleTargetsByUserIdAndRoleId

> []RoleTarget GetRoleTargetsByUserIdAndRoleId(ctx, userId, roleIdOrEncodedRoleId).AssignmentType(assignmentType).After(after).Limit(limit).Execute()

Retrieve a role target by assignment type



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
	roleIdOrEncodedRoleId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role or Base32 encoded `id` of the role name
	assignmentType := "GROUP" // string | Specifies the assignment type of the user (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetAdminAPI.GetRoleTargetsByUserIdAndRoleId(context.Background(), userId, roleIdOrEncodedRoleId).AssignmentType(assignmentType).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.GetRoleTargetsByUserIdAndRoleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleTargetsByUserIdAndRoleId`: []RoleTarget
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetAdminAPI.GetRoleTargetsByUserIdAndRoleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleIdOrEncodedRoleId** | **string** | The &#x60;id&#x60; of the role or Base32 encoded &#x60;id&#x60; of the role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleTargetsByUserIdAndRoleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignmentType** | **string** | Specifies the assignment type of the user | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**[]RoleTarget**](RoleTarget.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTargetsForApplicationAdministratorRoleForUser

> []CatalogApplication ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx, userId, roleAssignmentId).After(after).Limit(limit).Execute()

List all admin role app targets



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetAdminAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser(context.Background(), userId, roleAssignmentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationTargetsForApplicationAdministratorRoleForUser`: []CatalogApplication
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetAdminAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTargetsForApplicationAdministratorRoleForUserRequest struct via the builder pattern


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


## ListGroupTargetsForRole

> []Group ListGroupTargetsForRole(ctx, userId, roleAssignmentId).After(after).Limit(limit).Execute()

List all admin role group targets



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetAdminAPI.ListGroupTargetsForRole(context.Background(), userId, roleAssignmentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.ListGroupTargetsForRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupTargetsForRole`: []Group
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetAdminAPI.ListGroupTargetsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupTargetsForRoleRequest struct via the builder pattern


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


## UnassignAppInstanceTargetFromAdminRoleForUser

> UnassignAppInstanceTargetFromAdminRoleForUser(ctx, userId, roleAssignmentId, appName, appId).Execute()

Unassign an admin role app instance target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetAdminAPI.UnassignAppInstanceTargetFromAdminRoleForUser(context.Background(), userId, roleAssignmentId, appName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.UnassignAppInstanceTargetFromAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 
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


## UnassignAppTargetFromAppAdminRoleForUser

> UnassignAppTargetFromAppAdminRoleForUser(ctx, userId, roleAssignmentId, appName).Execute()

Unassign an admin role app target



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
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetAdminAPI.UnassignAppTargetFromAppAdminRoleForUser(context.Background(), userId, roleAssignmentId, appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.UnassignAppTargetFromAppAdminRoleForUser``: %v\n", err)
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
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 

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


## UnassignGroupTargetFromUserAdminRole

> UnassignGroupTargetFromUserAdminRole(ctx, userId, roleAssignmentId, groupId).Execute()

Unassign an admin role group target



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
	groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetAdminAPI.UnassignGroupTargetFromUserAdminRole(context.Background(), userId, roleAssignmentId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetAdminAPI.UnassignGroupTargetFromUserAdminRole``: %v\n", err)
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

