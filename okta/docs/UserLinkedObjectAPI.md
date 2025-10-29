# \UserLinkedObjectAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignLinkedObjectValueForPrimary**](UserLinkedObjectAPI.md#AssignLinkedObjectValueForPrimary) | **Put** /api/v1/users/{userIdOrLogin}/linkedObjects/{primaryRelationshipName}/{primaryUserId} | Assign a linked object value for primary
[**DeleteLinkedObjectForUser**](UserLinkedObjectAPI.md#DeleteLinkedObjectForUser) | **Delete** /api/v1/users/{userIdOrLogin}/linkedObjects/{relationshipName} | Delete a linked object value
[**ListLinkedObjectsForUser**](UserLinkedObjectAPI.md#ListLinkedObjectsForUser) | **Get** /api/v1/users/{userIdOrLogin}/linkedObjects/{relationshipName} | List the primary or all of the associated linked object values



## AssignLinkedObjectValueForPrimary

> AssignLinkedObjectValueForPrimary(ctx, userIdOrLogin, primaryRelationshipName, primaryUserId).Execute()

Assign a linked object value for primary



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
	userIdOrLogin := "00u5zex6ztMbOZhF50h7" // string | If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
	primaryRelationshipName := "manager" // string | Name of the `primary` relationship being assigned
	primaryUserId := "primaryUserId_example" // string | User ID to be assigned to the `primary` relationship for the `associated` user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserLinkedObjectAPI.AssignLinkedObjectValueForPrimary(context.Background(), userIdOrLogin, primaryRelationshipName, primaryUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLinkedObjectAPI.AssignLinkedObjectValueForPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrLogin** | **string** | If for the &#x60;self&#x60; link, this is the ID of the user for whom you want to get the primary user ID. If for the &#x60;associated&#x60; relation, this is the user ID or login value of the user assigned the associated relationship.  This can be &#x60;me&#x60; to represent the current session user. | 
**primaryRelationshipName** | **string** | Name of the &#x60;primary&#x60; relationship being assigned | 
**primaryUserId** | **string** | User ID to be assigned to the &#x60;primary&#x60; relationship for the &#x60;associated&#x60; user | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignLinkedObjectValueForPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkedObjectForUser

> DeleteLinkedObjectForUser(ctx, userIdOrLogin, relationshipName).Execute()

Delete a linked object value



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
	userIdOrLogin := "00u5zex6ztMbOZhF50h7" // string | If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
	relationshipName := "manager" // string | Name of the `primary` or `associated` relationship being queried

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserLinkedObjectAPI.DeleteLinkedObjectForUser(context.Background(), userIdOrLogin, relationshipName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLinkedObjectAPI.DeleteLinkedObjectForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrLogin** | **string** | If for the &#x60;self&#x60; link, this is the ID of the user for whom you want to get the primary user ID. If for the &#x60;associated&#x60; relation, this is the user ID or login value of the user assigned the associated relationship.  This can be &#x60;me&#x60; to represent the current session user. | 
**relationshipName** | **string** | Name of the &#x60;primary&#x60; or &#x60;associated&#x60; relationship being queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkedObjectForUserRequest struct via the builder pattern


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


## ListLinkedObjectsForUser

> []ResponseLinks ListLinkedObjectsForUser(ctx, userIdOrLogin, relationshipName).Execute()

List the primary or all of the associated linked object values



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
	userIdOrLogin := "00u5zex6ztMbOZhF50h7" // string | If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
	relationshipName := "manager" // string | Name of the `primary` or `associated` relationship being queried

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLinkedObjectAPI.ListLinkedObjectsForUser(context.Background(), userIdOrLogin, relationshipName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLinkedObjectAPI.ListLinkedObjectsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLinkedObjectsForUser`: []ResponseLinks
	fmt.Fprintf(os.Stdout, "Response from `UserLinkedObjectAPI.ListLinkedObjectsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userIdOrLogin** | **string** | If for the &#x60;self&#x60; link, this is the ID of the user for whom you want to get the primary user ID. If for the &#x60;associated&#x60; relation, this is the user ID or login value of the user assigned the associated relationship.  This can be &#x60;me&#x60; to represent the current session user. | 
**relationshipName** | **string** | Name of the &#x60;primary&#x60; or &#x60;associated&#x60; relationship being queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLinkedObjectsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ResponseLinks**](ResponseLinks.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

