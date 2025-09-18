# \ApplicationCrossAppAccessConnectionsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrossAppAccessConnection**](ApplicationCrossAppAccessConnectionsAPI.md#CreateCrossAppAccessConnection) | **Post** /api/v1/apps/{appId}/cwo/connections | Create a Cross App Access connection
[**DeleteCrossAppAccessConnection**](ApplicationCrossAppAccessConnectionsAPI.md#DeleteCrossAppAccessConnection) | **Delete** /api/v1/apps/{appId}/cwo/connections/{connectionId} | Delete a Cross App Access connection
[**GetAllCrossAppAccessConnections**](ApplicationCrossAppAccessConnectionsAPI.md#GetAllCrossAppAccessConnections) | **Get** /api/v1/apps/{appId}/cwo/connections | Retrieve all Cross App Access connections
[**GetCrossAppAccessConnection**](ApplicationCrossAppAccessConnectionsAPI.md#GetCrossAppAccessConnection) | **Get** /api/v1/apps/{appId}/cwo/connections/{connectionId} | Retrieve a Cross App Access connection
[**UpdateCrossAppAccessConnection**](ApplicationCrossAppAccessConnectionsAPI.md#UpdateCrossAppAccessConnection) | **Patch** /api/v1/apps/{appId}/cwo/connections/{connectionId} | Update a Cross App Access connection



## CreateCrossAppAccessConnection

> OrgCrossAppAccessConnection CreateCrossAppAccessConnection(ctx, appId).OrgCrossAppAccessConnection(orgCrossAppAccessConnection).Execute()

Create a Cross App Access connection



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	orgCrossAppAccessConnection := *openapiclient.NewOrgCrossAppAccessConnection() // OrgCrossAppAccessConnection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.CreateCrossAppAccessConnection(context.Background(), appId).OrgCrossAppAccessConnection(orgCrossAppAccessConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCrossAppAccessConnectionsAPI.CreateCrossAppAccessConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCrossAppAccessConnection`: OrgCrossAppAccessConnection
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCrossAppAccessConnectionsAPI.CreateCrossAppAccessConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCrossAppAccessConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgCrossAppAccessConnection** | [**OrgCrossAppAccessConnection**](OrgCrossAppAccessConnection.md) |  | 

### Return type

[**OrgCrossAppAccessConnection**](OrgCrossAppAccessConnection.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrossAppAccessConnection

> DeleteCrossAppAccessConnection(ctx, appId, connectionId).Execute()

Delete a Cross App Access connection



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	connectionId := "0oafxqCAJWWGELFTYASJ" // string | Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.DeleteCrossAppAccessConnection(context.Background(), appId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCrossAppAccessConnectionsAPI.DeleteCrossAppAccessConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**connectionId** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrossAppAccessConnectionRequest struct via the builder pattern


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


## GetAllCrossAppAccessConnections

> []OrgCrossAppAccessConnection GetAllCrossAppAccessConnections(ctx, appId).After(after).Limit(limit).Execute()

Retrieve all Cross App Access connections



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	after := "after_example" // string | Specifies the pagination cursor for the next page of connection results (optional)
	limit := int32(56) // int32 | Specifies the number of results to return per page. The values:   * -1: Return all results (up to system maximum)   * 0: Return an empty result set   * Positive integer: Return up to that many results (capped at system maximum)  (optional) (default to -1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.GetAllCrossAppAccessConnections(context.Background(), appId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCrossAppAccessConnectionsAPI.GetAllCrossAppAccessConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCrossAppAccessConnections`: []OrgCrossAppAccessConnection
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCrossAppAccessConnectionsAPI.GetAllCrossAppAccessConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCrossAppAccessConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | Specifies the pagination cursor for the next page of connection results | 
 **limit** | **int32** | Specifies the number of results to return per page. The values:   * -1: Return all results (up to system maximum)   * 0: Return an empty result set   * Positive integer: Return up to that many results (capped at system maximum)  | [default to -1]

### Return type

[**[]OrgCrossAppAccessConnection**](OrgCrossAppAccessConnection.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossAppAccessConnection

> OrgCrossAppAccessConnection GetCrossAppAccessConnection(ctx, appId, connectionId).Execute()

Retrieve a Cross App Access connection



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	connectionId := "0oafxqCAJWWGELFTYASJ" // string | Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.GetCrossAppAccessConnection(context.Background(), appId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCrossAppAccessConnectionsAPI.GetCrossAppAccessConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossAppAccessConnection`: OrgCrossAppAccessConnection
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCrossAppAccessConnectionsAPI.GetCrossAppAccessConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**connectionId** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossAppAccessConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgCrossAppAccessConnection**](OrgCrossAppAccessConnection.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCrossAppAccessConnection

> OrgCrossAppAccessConnection UpdateCrossAppAccessConnection(ctx, appId, connectionId).OrgCrossAppAccessConnectionPatchRequest(orgCrossAppAccessConnectionPatchRequest).Execute()

Update a Cross App Access connection



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	connectionId := "0oafxqCAJWWGELFTYASJ" // string | Connection ID
	orgCrossAppAccessConnectionPatchRequest := *openapiclient.NewOrgCrossAppAccessConnectionPatchRequest("ACTIVE") // OrgCrossAppAccessConnectionPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCrossAppAccessConnectionsAPI.UpdateCrossAppAccessConnection(context.Background(), appId, connectionId).OrgCrossAppAccessConnectionPatchRequest(orgCrossAppAccessConnectionPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCrossAppAccessConnectionsAPI.UpdateCrossAppAccessConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCrossAppAccessConnection`: OrgCrossAppAccessConnection
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCrossAppAccessConnectionsAPI.UpdateCrossAppAccessConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**connectionId** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCrossAppAccessConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgCrossAppAccessConnectionPatchRequest** | [**OrgCrossAppAccessConnectionPatchRequest**](OrgCrossAppAccessConnectionPatchRequest.md) |  | 

### Return type

[**OrgCrossAppAccessConnection**](OrgCrossAppAccessConnection.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

