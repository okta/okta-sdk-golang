# \GroupPushMappingAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupPushMapping**](GroupPushMappingAPI.md#CreateGroupPushMapping) | **Post** /api/v1/apps/{appId}/group-push/mappings | Create a group push mapping
[**DeleteGroupPushMapping**](GroupPushMappingAPI.md#DeleteGroupPushMapping) | **Delete** /api/v1/apps/{appId}/group-push/mappings/{mappingId} | Delete a group push mapping
[**GetGroupPushMapping**](GroupPushMappingAPI.md#GetGroupPushMapping) | **Get** /api/v1/apps/{appId}/group-push/mappings/{mappingId} | Retrieve a group push mapping
[**ListGroupPushMappings**](GroupPushMappingAPI.md#ListGroupPushMappings) | **Get** /api/v1/apps/{appId}/group-push/mappings | List all group push mappings
[**UpdateGroupPushMapping**](GroupPushMappingAPI.md#UpdateGroupPushMapping) | **Patch** /api/v1/apps/{appId}/group-push/mappings/{mappingId} | Update a group push mapping



## CreateGroupPushMapping

> GroupPushMapping CreateGroupPushMapping(ctx, appId).Body(body).Execute()

Create a group push mapping



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
	body := *openapiclient.NewCreateGroupPushMappingRequest("SourceGroupId_example") // CreateGroupPushMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPushMappingAPI.CreateGroupPushMapping(context.Background(), appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPushMappingAPI.CreateGroupPushMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupPushMapping`: GroupPushMapping
	fmt.Fprintf(os.Stdout, "Response from `GroupPushMappingAPI.CreateGroupPushMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupPushMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateGroupPushMappingRequest**](CreateGroupPushMappingRequest.md) |  | 

### Return type

[**GroupPushMapping**](GroupPushMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupPushMapping

> DeleteGroupPushMapping(ctx, appId, mappingId).DeleteTargetGroup(deleteTargetGroup).Execute()

Delete a group push mapping



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
	mappingId := "gPm00000000000000000" // string | Group push mapping ID
	deleteTargetGroup := true // bool | If set to `true`, the target group is also deleted. If set to `false`, the target group isn't deleted. (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupPushMappingAPI.DeleteGroupPushMapping(context.Background(), appId, mappingId).DeleteTargetGroup(deleteTargetGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPushMappingAPI.DeleteGroupPushMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**mappingId** | **string** | Group push mapping ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupPushMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteTargetGroup** | **bool** | If set to &#x60;true&#x60;, the target group is also deleted. If set to &#x60;false&#x60;, the target group isn&#39;t deleted. | [default to false]

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


## GetGroupPushMapping

> GroupPushMapping GetGroupPushMapping(ctx, appId, mappingId).Execute()

Retrieve a group push mapping



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
	mappingId := "gPm00000000000000000" // string | Group push mapping ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPushMappingAPI.GetGroupPushMapping(context.Background(), appId, mappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPushMappingAPI.GetGroupPushMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupPushMapping`: GroupPushMapping
	fmt.Fprintf(os.Stdout, "Response from `GroupPushMappingAPI.GetGroupPushMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**mappingId** | **string** | Group push mapping ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPushMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupPushMapping**](GroupPushMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupPushMappings

> []GroupPushMapping ListGroupPushMappings(ctx, appId).After(after).Limit(limit).LastUpdated(lastUpdated).SourceGroupId(sourceGroupId).Status(status).Execute()

List all group push mappings



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
	after := "after_example" // string | Specifies the pagination cursor for the next page of mappings (optional)
	limit := int32(56) // int32 | Specifies the number of results returned (optional) (default to 100)
	lastUpdated := "2025-01-01T00:00:00Z" // string | Filters group push mappings by last updated date. The `lastUpdated` parameter supports the following format: `YYYY-MM-DDTHH:mm:ssZ`. This filters mappings updated on or after the specified date and time in UTC.  If you don't specify a value, all group push mappings are returned. (optional)
	sourceGroupId := "00g00000000000000000" // string | Filters group push mappings by source group ID. If you don't specify a value, all group push mappings are returned. (optional)
	status := "status_example" // string | Filters group push mappings by status. If you don't specify a value, all group push mappings are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPushMappingAPI.ListGroupPushMappings(context.Background(), appId).After(after).Limit(limit).LastUpdated(lastUpdated).SourceGroupId(sourceGroupId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPushMappingAPI.ListGroupPushMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupPushMappings`: []GroupPushMapping
	fmt.Fprintf(os.Stdout, "Response from `GroupPushMappingAPI.ListGroupPushMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupPushMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | Specifies the pagination cursor for the next page of mappings | 
 **limit** | **int32** | Specifies the number of results returned | [default to 100]
 **lastUpdated** | **string** | Filters group push mappings by last updated date. The &#x60;lastUpdated&#x60; parameter supports the following format: &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60;. This filters mappings updated on or after the specified date and time in UTC.  If you don&#39;t specify a value, all group push mappings are returned. | 
 **sourceGroupId** | **string** | Filters group push mappings by source group ID. If you don&#39;t specify a value, all group push mappings are returned. | 
 **status** | **string** | Filters group push mappings by status. If you don&#39;t specify a value, all group push mappings are returned. | 

### Return type

[**[]GroupPushMapping**](GroupPushMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupPushMapping

> GroupPushMapping UpdateGroupPushMapping(ctx, appId, mappingId).Body(body).Execute()

Update a group push mapping



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
	mappingId := "gPm00000000000000000" // string | Group push mapping ID
	body := *openapiclient.NewUpdateGroupPushMappingRequest("Status_example") // UpdateGroupPushMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPushMappingAPI.UpdateGroupPushMapping(context.Background(), appId, mappingId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPushMappingAPI.UpdateGroupPushMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupPushMapping`: GroupPushMapping
	fmt.Fprintf(os.Stdout, "Response from `GroupPushMappingAPI.UpdateGroupPushMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**mappingId** | **string** | Group push mapping ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupPushMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateGroupPushMappingRequest**](UpdateGroupPushMappingRequest.md) |  | 

### Return type

[**GroupPushMapping**](GroupPushMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

