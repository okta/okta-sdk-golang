# \GroupAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroup**](GroupAPI.md#AddGroup) | **Post** /api/v1/groups | Add a group
[**AssignUserToGroup**](GroupAPI.md#AssignUserToGroup) | **Put** /api/v1/groups/{groupId}/users/{userId} | Assign a user to a group
[**DeleteGroup**](GroupAPI.md#DeleteGroup) | **Delete** /api/v1/groups/{groupId} | Delete a group
[**GetGroup**](GroupAPI.md#GetGroup) | **Get** /api/v1/groups/{groupId} | Retrieve a group
[**ListAssignedApplicationsForGroup**](GroupAPI.md#ListAssignedApplicationsForGroup) | **Get** /api/v1/groups/{groupId}/apps | List all assigned apps
[**ListGroupUsers**](GroupAPI.md#ListGroupUsers) | **Get** /api/v1/groups/{groupId}/users | List all member users
[**ListGroups**](GroupAPI.md#ListGroups) | **Get** /api/v1/groups | List all groups
[**ReplaceGroup**](GroupAPI.md#ReplaceGroup) | **Put** /api/v1/groups/{groupId} | Replace a group
[**UnassignUserFromGroup**](GroupAPI.md#UnassignUserFromGroup) | **Delete** /api/v1/groups/{groupId}/users/{userId} | Unassign a user from a group



## AddGroup

> Group AddGroup(ctx).Group(group).Execute()

Add a group



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
	group := *openapiclient.NewAddGroupRequest() // AddGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.AddGroup(context.Background()).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AddGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AddGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**AddGroupRequest**](AddGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignUserToGroup

> AssignUserToGroup(ctx, groupId, userId).Execute()

Assign a user to a group



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.AssignUserToGroup(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AssignUserToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUserToGroupRequest struct via the builder pattern


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


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()

Delete a group



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.DeleteGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## GetGroup

> Group GetGroup(ctx, groupId).Execute()

Retrieve a group



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssignedApplicationsForGroup

> []ListApplications200ResponseInner ListAssignedApplicationsForGroup(ctx, groupId).After(after).Limit(limit).Execute()

List all assigned apps



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
	after := "after_example" // string | Specifies the pagination cursor for the next page of apps (optional)
	limit := int32(56) // int32 | Specifies the number of app results for a page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListAssignedApplicationsForGroup(context.Background(), groupId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListAssignedApplicationsForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssignedApplicationsForGroup`: []ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListAssignedApplicationsForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssignedApplicationsForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | Specifies the pagination cursor for the next page of apps | 
 **limit** | **int32** | Specifies the number of app results for a page | [default to 20]

### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupUsers

> []User ListGroupUsers(ctx, groupId).After(after).Limit(limit).Execute()

List all member users



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)
	limit := int32(56) // int32 | Specifies the number of user results in a page (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListGroupUsers(context.Background(), groupId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
 **limit** | **int32** | Specifies the number of user results in a page | [default to 1000]

### Return type

[**[]User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> []Group ListGroups(ctx).Search(search).Filter(filter).Q(q).After(after).Limit(limit).Expand(expand).SortBy(sortBy).SortOrder(sortOrder).Execute()

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
	search := "type%20eq%20%22APP_GROUP%22" // string | Searches for groups with a supported [filtering](https://developer.okta.com/docs/api/#filter) expression for all properties except for `_embedded`, `_links`, and `objectClass`. Okta recommends this query parameter because it provides the largest range of search options and optimal performance.  This operation supports [pagination](https://developer.okta.com/docs/api/#pagination).  Using search requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding), for example, `search=type eq \"OKTA_GROUP\"` is encoded as `search=type+eq+%22OKTA_GROUP%22`.  This operation searches many properties:  * Any group profile attribute, including imported app group profile attributes. * The top-level properties: `id`, `created`, `lastMembershipUpdated`, `lastUpdated`, and `type`. * The [source](/openapi/okta-management/management/tag/Group/#tag/Group/operation/listGroups!c=200&path=_links/source&t=response) of groups with type of `APP_GROUP`, accessed as `source.id`.  You can also use the `sortBy` and `sortOrder` parameters.  Searches for groups can be filtered by the following operators: `sw`, `eq`, and `co`. You can only use `co` with these select profile attributes: `profile.name` and `profile.description`. See [Operators](https://developer.okta.com/docs/api/#operators). (optional)
	filter := "id%20eq%20%2200g1emaKYZTWRYYRRTSK%22" // string | Filter expression for groups. See [Filter](https://developer.okta.com/docs/api/#filter).  Filtering supports the following limited number of properties: `id`, `type`, `lastUpdated`, and `lastMembershipUpdated`.  > **Note:** All filters must be [URL encoded](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). For example, `filter=lastUpdated gt \"2013-06-01T00:00:00.000Z\"` is encoded as `filter=lastUpdated%20gt%20%222013-06-01T00:00:00.000Z%22`.  See [Special characters](https://developer.okta.com/docs/api/#special-characters). (optional)
	q := "West&limit=10" // string | Finds a group that matches the `name` property. > **Note:** Paging and searching are currently mutually exclusive. You can't page a query. The default limit for a query is 300 results. Query is intended for an auto-complete picker use case where users refine their search string to constrain the results. (optional)
	after := "after_example" // string | Specifies the pagination cursor for the next page of groups. The `after` cursor should be treated as an opaque value and obtained through the next link relation. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | Specifies the number of group results in a page.  Okta recommends using a specific value other than the default or maximum. If your request times out, retry your request with a smaller `limit` and [page the results](https://developer.okta.com/docs/api/#pagination).  The Okta default `Everyone` group isn't returned for users with a group admin role. (optional)
	expand := "expand_example" // string | If specified, additional metadata is included in the response. Possible values are `stats` and `app`. This additional metadata is listed in the [`_embedded`](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/addGroup!c=200&path=_embedded&t=response) property of the response.  > **Note:** You can use the `stats` value to return the number of users within a group. This is listed as the `_embedded.stats.usersCount` value in the response. See this [Knowledge Base article](https://support.okta.com/help/s/article/Is-there-an-API-that-returns-the-number-of-users-in-a-group?language=en_US) for more information and an example. (optional)
	sortBy := "lastUpdated" // string | Specifies the field to sort by (for search queries only). `sortBy` can be any single property, for example `sortBy=profile.name`. Groups with the same value for the `sortBy` property are ordered by `id`'. Use with `sortOrder` to control the order of results. (optional)
	sortOrder := "sortOrder_example" // string | Specifies sort order: `asc` or `desc` (for search queries only). This parameter is ignored if `sortBy` isn't present. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListGroups(context.Background()).Search(search).Filter(filter).Q(q).After(after).Limit(limit).Expand(expand).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroups`: []Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Searches for groups with a supported [filtering](https://developer.okta.com/docs/api/#filter) expression for all properties except for &#x60;_embedded&#x60;, &#x60;_links&#x60;, and &#x60;objectClass&#x60;. Okta recommends this query parameter because it provides the largest range of search options and optimal performance.  This operation supports [pagination](https://developer.okta.com/docs/api/#pagination).  Using search requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding), for example, &#x60;search&#x3D;type eq \&quot;OKTA_GROUP\&quot;&#x60; is encoded as &#x60;search&#x3D;type+eq+%22OKTA_GROUP%22&#x60;.  This operation searches many properties:  * Any group profile attribute, including imported app group profile attributes. * The top-level properties: &#x60;id&#x60;, &#x60;created&#x60;, &#x60;lastMembershipUpdated&#x60;, &#x60;lastUpdated&#x60;, and &#x60;type&#x60;. * The [source](/openapi/okta-management/management/tag/Group/#tag/Group/operation/listGroups!c&#x3D;200&amp;path&#x3D;_links/source&amp;t&#x3D;response) of groups with type of &#x60;APP_GROUP&#x60;, accessed as &#x60;source.id&#x60;.  You can also use the &#x60;sortBy&#x60; and &#x60;sortOrder&#x60; parameters.  Searches for groups can be filtered by the following operators: &#x60;sw&#x60;, &#x60;eq&#x60;, and &#x60;co&#x60;. You can only use &#x60;co&#x60; with these select profile attributes: &#x60;profile.name&#x60; and &#x60;profile.description&#x60;. See [Operators](https://developer.okta.com/docs/api/#operators). | 
 **filter** | **string** | Filter expression for groups. See [Filter](https://developer.okta.com/docs/api/#filter).  Filtering supports the following limited number of properties: &#x60;id&#x60;, &#x60;type&#x60;, &#x60;lastUpdated&#x60;, and &#x60;lastMembershipUpdated&#x60;.  &gt; **Note:** All filters must be [URL encoded](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). For example, &#x60;filter&#x3D;lastUpdated gt \&quot;2013-06-01T00:00:00.000Z\&quot;&#x60; is encoded as &#x60;filter&#x3D;lastUpdated%20gt%20%222013-06-01T00:00:00.000Z%22&#x60;.  See [Special characters](https://developer.okta.com/docs/api/#special-characters). | 
 **q** | **string** | Finds a group that matches the &#x60;name&#x60; property. &gt; **Note:** Paging and searching are currently mutually exclusive. You can&#39;t page a query. The default limit for a query is 300 results. Query is intended for an auto-complete picker use case where users refine their search string to constrain the results. | 
 **after** | **string** | Specifies the pagination cursor for the next page of groups. The &#x60;after&#x60; cursor should be treated as an opaque value and obtained through the next link relation. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | Specifies the number of group results in a page.  Okta recommends using a specific value other than the default or maximum. If your request times out, retry your request with a smaller &#x60;limit&#x60; and [page the results](https://developer.okta.com/docs/api/#pagination).  The Okta default &#x60;Everyone&#x60; group isn&#39;t returned for users with a group admin role. | 
 **expand** | **string** | If specified, additional metadata is included in the response. Possible values are &#x60;stats&#x60; and &#x60;app&#x60;. This additional metadata is listed in the [&#x60;_embedded&#x60;](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/addGroup!c&#x3D;200&amp;path&#x3D;_embedded&amp;t&#x3D;response) property of the response.  &gt; **Note:** You can use the &#x60;stats&#x60; value to return the number of users within a group. This is listed as the &#x60;_embedded.stats.usersCount&#x60; value in the response. See this [Knowledge Base article](https://support.okta.com/help/s/article/Is-there-an-API-that-returns-the-number-of-users-in-a-group?language&#x3D;en_US) for more information and an example. | 
 **sortBy** | **string** | Specifies the field to sort by (for search queries only). &#x60;sortBy&#x60; can be any single property, for example &#x60;sortBy&#x3D;profile.name&#x60;. Groups with the same value for the &#x60;sortBy&#x60; property are ordered by &#x60;id&#x60;&#39;. Use with &#x60;sortOrder&#x60; to control the order of results. | 
 **sortOrder** | **string** | Specifies sort order: &#x60;asc&#x60; or &#x60;desc&#x60; (for search queries only). This parameter is ignored if &#x60;sortBy&#x60; isn&#39;t present. | [default to &quot;asc&quot;]

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


## ReplaceGroup

> Group ReplaceGroup(ctx, groupId).Group(group).Execute()

Replace a group



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
	group := *openapiclient.NewAddGroupRequest() // AddGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ReplaceGroup(context.Background(), groupId).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ReplaceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ReplaceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**AddGroupRequest**](AddGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignUserFromGroup

> UnassignUserFromGroup(ctx, groupId, userId).Execute()

Unassign a user from a group



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.UnassignUserFromGroup(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.UnassignUserFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignUserFromGroupRequest struct via the builder pattern


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

