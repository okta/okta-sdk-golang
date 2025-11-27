# \RoleBTargetClientAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAppTargetInstanceRoleForClient**](RoleBTargetClientAPI.md#AssignAppTargetInstanceRoleForClient) | **Put** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId} | Assign a client role app instance target
[**AssignAppTargetRoleToClient**](RoleBTargetClientAPI.md#AssignAppTargetRoleToClient) | **Put** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName} | Assign a client role app target
[**AssignGroupTargetRoleForClient**](RoleBTargetClientAPI.md#AssignGroupTargetRoleForClient) | **Put** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/groups/{groupId} | Assign a client role group target
[**ListAppTargetRoleToClient**](RoleBTargetClientAPI.md#ListAppTargetRoleToClient) | **Get** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps | List all client role app targets
[**ListGroupTargetRoleForClient**](RoleBTargetClientAPI.md#ListGroupTargetRoleForClient) | **Get** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/groups | List all client role group targets
[**RemoveAppTargetInstanceRoleForClient**](RoleBTargetClientAPI.md#RemoveAppTargetInstanceRoleForClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId} | Unassign a client role app instance target
[**RemoveAppTargetRoleFromClient**](RoleBTargetClientAPI.md#RemoveAppTargetRoleFromClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName} | Unassign a client role app target
[**RemoveGroupTargetRoleFromClient**](RoleBTargetClientAPI.md#RemoveGroupTargetRoleFromClient) | **Delete** /oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/groups/{groupId} | Unassign a client role group target



## AssignAppTargetInstanceRoleForClient

> AssignAppTargetInstanceRoleForClient(ctx, clientId, roleAssignmentId, appName, appId).Execute()

Assign a client role app instance target



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetClientAPI.AssignAppTargetInstanceRoleForClient(context.Background(), clientId, roleAssignmentId, appName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.AssignAppTargetInstanceRoleForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 
**appId** | **string** | Application ID | 

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

> AssignAppTargetRoleToClient(ctx, clientId, roleAssignmentId, appName).Execute()

Assign a client role app target



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetClientAPI.AssignAppTargetRoleToClient(context.Background(), clientId, roleAssignmentId, appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.AssignAppTargetRoleToClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 

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


## AssignGroupTargetRoleForClient

> AssignGroupTargetRoleForClient(ctx, clientId, roleAssignmentId, groupId).Execute()

Assign a client role group target



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetClientAPI.AssignGroupTargetRoleForClient(context.Background(), clientId, roleAssignmentId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.AssignGroupTargetRoleForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
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


## ListAppTargetRoleToClient

> []CatalogApplication ListAppTargetRoleToClient(ctx, clientId, roleAssignmentId).After(after).Limit(limit).Execute()

List all client role app targets



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetClientAPI.ListAppTargetRoleToClient(context.Background(), clientId, roleAssignmentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.ListAppTargetRoleToClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppTargetRoleToClient`: []CatalogApplication
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetClientAPI.ListAppTargetRoleToClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppTargetRoleToClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
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


## ListGroupTargetRoleForClient

> []Group ListGroupTargetRoleForClient(ctx, clientId, roleAssignmentId).After(after).Limit(limit).Execute()

List all client role group targets



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleBTargetClientAPI.ListGroupTargetRoleForClient(context.Background(), clientId, roleAssignmentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.ListGroupTargetRoleForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupTargetRoleForClient`: []Group
	fmt.Fprintf(os.Stdout, "Response from `RoleBTargetClientAPI.ListGroupTargetRoleForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupTargetRoleForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
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


## RemoveAppTargetInstanceRoleForClient

> RemoveAppTargetInstanceRoleForClient(ctx, clientId, roleAssignmentId, appName, appId).Execute()

Unassign a client role app instance target



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetClientAPI.RemoveAppTargetInstanceRoleForClient(context.Background(), clientId, roleAssignmentId, appName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.RemoveAppTargetInstanceRoleForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 
**appId** | **string** | Application ID | 

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

> RemoveAppTargetRoleFromClient(ctx, clientId, roleAssignmentId, appName).Execute()

Unassign a client role app target



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	appName := "google" // string | Name of the app definition (the OIN catalog app key name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetClientAPI.RemoveAppTargetRoleFromClient(context.Background(), clientId, roleAssignmentId, appName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.RemoveAppTargetRoleFromClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
**appName** | **string** | Name of the app definition (the OIN catalog app key name) | 

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

> RemoveGroupTargetRoleFromClient(ctx, clientId, roleAssignmentId, groupId).Execute()

Unassign a client role group target



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
	roleAssignmentId := "JBCUYUC7IRCVGS27IFCE2SKO" // string | The `id` of the role assignment
	groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleBTargetClientAPI.RemoveGroupTargetRoleFromClient(context.Background(), clientId, roleAssignmentId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleBTargetClientAPI.RemoveGroupTargetRoleFromClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**roleAssignmentId** | **string** | The &#x60;id&#x60; of the role assignment | 
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

