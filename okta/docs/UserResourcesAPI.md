# \UserResourcesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAppLinks**](UserResourcesAPI.md#ListAppLinks) | **Get** /api/v1/users/{id}/appLinks | List all assigned app links
[**ListUserClients**](UserResourcesAPI.md#ListUserClients) | **Get** /api/v1/users/{userId}/clients | List all clients
[**ListUserDevices**](UserResourcesAPI.md#ListUserDevices) | **Get** /api/v1/users/{userId}/devices | List all devices
[**ListUserGroups**](UserResourcesAPI.md#ListUserGroups) | **Get** /api/v1/users/{id}/groups | List all groups



## ListAppLinks

> []AssignedAppLink ListAppLinks(ctx, id).Execute()

List all assigned app links



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
	id := "id_example" // string | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserResourcesAPI.ListAppLinks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserResourcesAPI.ListAppLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppLinks`: []AssignedAppLink
	fmt.Fprintf(os.Stdout, "Response from `UserResourcesAPI.ListAppLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AssignedAppLink**](AssignedAppLink.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserClients

> []OAuth2Client ListUserClients(ctx, userId).Execute()

List all clients



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserResourcesAPI.ListUserClients(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserResourcesAPI.ListUserClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserClients`: []OAuth2Client
	fmt.Fprintf(os.Stdout, "Response from `UserResourcesAPI.ListUserClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserDevices

> []UserDevice ListUserDevices(ctx, userId).Execute()

List all devices



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserResourcesAPI.ListUserDevices(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserResourcesAPI.ListUserDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserDevices`: []UserDevice
	fmt.Fprintf(os.Stdout, "Response from `UserResourcesAPI.ListUserDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserDevice**](UserDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroups

> []Group ListUserGroups(ctx, id).Execute()

List all groups



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
	id := "id_example" // string | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserResourcesAPI.ListUserGroups(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserResourcesAPI.ListUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserGroups`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UserResourcesAPI.ListUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

