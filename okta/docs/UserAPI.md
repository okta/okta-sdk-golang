# \UserAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserAPI.md#CreateUser) | **Post** /api/v1/users | Create a user
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /api/v1/users/{id} | Delete a user
[**GetUser**](UserAPI.md#GetUser) | **Get** /api/v1/users/{id} | Retrieve a user
[**ListUserBlocks**](UserAPI.md#ListUserBlocks) | **Get** /api/v1/users/{id}/blocks | List all user blocks
[**ListUsers**](UserAPI.md#ListUsers) | **Get** /api/v1/users | List all users
[**ReplaceUser**](UserAPI.md#ReplaceUser) | **Put** /api/v1/users/{id} | Replace a user
[**UpdateUser**](UserAPI.md#UpdateUser) | **Post** /api/v1/users/{id} | Update a user



## CreateUser

> User CreateUser(ctx).Body(body).Activate(activate).Provider(provider).NextLogin(nextLogin).Execute()

Create a user



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
	body := *openapiclient.NewCreateUserRequest(*openapiclient.NewUserProfile()) // CreateUserRequest | 
	activate := true // bool | Executes an [activation lifecycle](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserLifecycle/#tag/UserLifecycle/operation/activateUser) operation when creating the user (optional) (default to true)
	provider := true // bool | Indicates whether to create a user with a specified authentication provider (optional) (default to false)
	nextLogin := "nextLogin_example" // string | With `activate=true`, if `nextLogin=changePassword`, a user is created, activated, and the password is set to `EXPIRED`. The user must change it the next time they sign in. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUser(context.Background()).Body(body).Activate(activate).Provider(provider).NextLogin(nextLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateUserRequest**](CreateUserRequest.md) |  | 
 **activate** | **bool** | Executes an [activation lifecycle](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserLifecycle/#tag/UserLifecycle/operation/activateUser) operation when creating the user | [default to true]
 **provider** | **bool** | Indicates whether to create a user with a specified authentication provider | [default to false]
 **nextLogin** | **string** | With &#x60;activate&#x3D;true&#x60;, if &#x60;nextLogin&#x3D;changePassword&#x60;, a user is created, activated, and the password is set to &#x60;EXPIRED&#x60;. The user must change it the next time they sign in. | 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).SendEmail(sendEmail).Prefer(prefer).Execute()

Delete a user



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
	sendEmail := true // bool | Sends a deactivation email to the admin if `true` (optional) (default to false)
	prefer := "prefer_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), id).SendEmail(sendEmail).Prefer(prefer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **bool** | Sends a deactivation email to the admin if &#x60;true&#x60; | [default to false]
 **prefer** | **string** |  | 

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


## GetUser

> UserGetSingleton GetUser(ctx, id).ContentType(contentType).Expand(expand).Execute()

Retrieve a user



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
	contentType := "application/json; okta-response=omitCredentials,omitCredentialsLinks" // string | Specifies the media type of the resource. Optional `okta-response` value can be included for performance optimization.  Complex DelAuth configurations may degrade performance when fetching specific parts of the response, and passing this parameter can omit these parts, bypassing the bottleneck.  Enum values for `okta-response`:   * `omitCredentials`: Omits the credentials subobject from the response.   * `omitCredentialsLinks`: Omits the following HAL links from the response: Update password, Change recovery question, Start forgot password flow, Reset password, Reset factors, Unlock.   * `omitTransitioningToStatus`: Omits the `transitioningToStatus` field from the response. (optional)
	expand := "blocks" // string | An optional parameter to include metadata in the `_embedded` attribute. Valid values: `blocks` or <x-lifecycle class=\"ea\"></x-lifecycle> `classification`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUser(context.Background(), id).ContentType(contentType).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserGetSingleton
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** | Specifies the media type of the resource. Optional &#x60;okta-response&#x60; value can be included for performance optimization.  Complex DelAuth configurations may degrade performance when fetching specific parts of the response, and passing this parameter can omit these parts, bypassing the bottleneck.  Enum values for &#x60;okta-response&#x60;:   * &#x60;omitCredentials&#x60;: Omits the credentials subobject from the response.   * &#x60;omitCredentialsLinks&#x60;: Omits the following HAL links from the response: Update password, Change recovery question, Start forgot password flow, Reset password, Reset factors, Unlock.   * &#x60;omitTransitioningToStatus&#x60;: Omits the &#x60;transitioningToStatus&#x60; field from the response. | 
 **expand** | **string** | An optional parameter to include metadata in the &#x60;_embedded&#x60; attribute. Valid values: &#x60;blocks&#x60; or &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &#x60;classification&#x60;. | 

### Return type

[**UserGetSingleton**](UserGetSingleton.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBlocks

> []UserBlock ListUserBlocks(ctx, id).Execute()

List all user blocks



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
	resp, r, err := apiClient.UserAPI.ListUserBlocks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserBlocks`: []UserBlock
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserBlock**](UserBlock.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx).ContentType(contentType).Search(search).Filter(filter).Q(q).After(after).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Fields(fields).Expand(expand).Execute()

List all users



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
	contentType := "application/json; okta-response=omitCredentials,omitCredentialsLinks" // string | Specifies the media type of the resource. Optional `okta-response` value can be included for performance optimization.  Complex DelAuth configurations may degrade performance when fetching specific parts of the response, and passing this parameter can omit these parts, bypassing the bottleneck.  Enum values for `okta-response`:   * `omitCredentials`: Omits the credentials subobject from the response.   * `omitCredentialsLinks`: Omits the following HAL links from the response: Update password, Change recovery question, Start forgot password flow, Reset password, Reset factors, Unlock.   * `omitTransitioningToStatus`: Omits the `transitioningToStatus` field from the response. (optional)
	search := "status%20eq%20%22STAGED%22" // string | Searches for users with a supported filtering expression for most properties. Okta recommends this query parameter because it provides the largest range of search options and optimal performance.  > **Note:** Using an overly complex or long search query can result in an error.  This operation supports [pagination](https://developer.okta.com/docs/api/#pagination). Use an ID lookup for records that you update to ensure your results contain the latest data. Returned users include those with the `DEPROVISIONED` status.  Property names in the search parameter are case sensitive, whereas operators (`eq`, `sw`, and so on) and string values are case insensitive. Unlike with user logins, diacritical marks are significant in search string values: a search for `isaac.brock` finds `Isaac.Brock`, but doesn't find a property whose value is `isáàc.bröck`.  This operation requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). See [Special characters](https://developer.okta.com/docs/api/#special-characters).  This operation searches many properties:   * Any user profile attribute, including custom-defined attributes   * The top-level properties: `id`, `status`, `created`, `activated`, `statusChanged`, and `lastUpdated`   * The [user type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/#tag/UserType/operation/updateUserType) accessed as `type.id`   * Properties that have array values  > **Note:** <x-lifecycle class=\"ea\"></x-lifecycle> The ability to search by user classification is available as an [Early Access](https://developer.okta.com/docs/api/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature. The `classification.type` property cannot be used in conjunction with other search terms. You can search using `classification.type eq \"LITE\"` or `classification.type eq \"STANDARD\"`.  You can also use `sortBy` and `sortOrder` parameters. The `ne` (not equal) operator isn't supported, but you can obtain the same result by using `lt ... or ... gt`. For example, to see all users except those that have a status of `STAGED`, use `(status lt \"STAGED\" or status gt \"STAGED\")`.  You can search properties that are arrays. If any element matches the search term, the entire array (object) is returned. Okta follows the [SCIM Protocol Specification](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) for searching arrays. You can search multiple arrays, multiple values in an array, as well as using the standard logical and filtering operators. See [Filter](https://developer.okta.com/docs/reference/core-okta-api/#filter).  Searches for users can be filtered by the following operators: `sw`, `eq`, and `co`. You can only use `co` with these select user profile attributes: `profile.firstName`, `profile.lastName`, `profile.email`, and `profile.login`. See [Operators](https://developer.okta.com/docs/api/#operators). (optional)
	filter := "status%20eq%20%22LOCKED_OUT%22" // string | Filters users with a supported expression for a subset of properties.  > **Note:** Returned users include those with the `DEPROVISIONED` status.  This requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). For example, `filter=lastUpdated gt \"2013-06-01T00:00:00.000Z\"` is encoded as `filter=lastUpdated%20gt%20%222013-06-01T00:00:00.000Z%22`. Filtering is case-sensitive for property names and query values, while operators are case-insensitive.  Filtering supports the following limited number of properties: `status`, `lastUpdated`, `id`, `profile.login`, `profile.email`, `profile.firstName`, and `profile.lastName`.  Additionally, filtering supports only the equal `eq` operator from the standard Okta API filtering semantics, except in the case of the `lastUpdated` property. This property can also use the inequality operators (`gt`, `ge`, `lt`, and `le`). For logical operators, only the logical operators `and` and `or` are supported. The `not` operator isn't supported. See [Filter](https://developer.okta.com/docs/api/#filter) and [Operators](https://developer.okta.com/docs/api/#operators). (optional)
	q := "q_example" // string | Finds users who match the specified query. Use the `q` parameter for simple queries, such as a lookup of users by name when creating a people picker.  The value of `q` is matched against `firstName`, `lastName`, or `email`. This performs a `startsWith` match, but this is an implementation detail and can change without notice. You don't need to specify `firstName`, `lastName`, or `email`.  > **Notes:** > * Using the `q` parameter in a request omits users that have a status of `DEPROVISIONED`. To return all users, use a `filter` or `search` query instead. > * This doesn't support pagination, but you can use `limit`. > * This isn't designed for large data sets. For optimal performance, use the `search` parameter instead. (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)
	limit := int32(56) // int32 | Specifies the number of results returned. Defaults to 10 if `q` is provided.  You can use `limit` with `after` to define the cursor location in the data set and manage the user records per page. (optional) (default to 200)
	sortBy := "sortBy_example" // string | Specifies the field to sort by (for search queries only). This can be any single property, for example `sortBy=profile.lastName`. Users with the same value for the `sortBy` property are ordered by `id`. Use with `sortOrder` to control the order of results. (optional)
	sortOrder := "sortOrder_example" // string | Specifies sort order: `asc` or `desc` (for search queries only). This parameter is ignored if `sortBy` isn't present. (optional) (default to "asc")
	fields := "id,status,profile:(firstName,lastName,city)" // string | Specifies a select set of user properties to query. Any other properties will be filtered out of the returned users. This is often called field projections in APIs, which can reduce payload size, improve performance, and limit unneccessary data exposure.  Requested fields should be comma-separated. Comma-separate the fields and place sub-fields in the profile object inside a `profile:()` directive, for example `profile:(firstName, city)`. The `id` field is always included, regardless of whether it's specified in the `fields` parameter. (optional)
	expand := "classification" // string | <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>A parameter to include metadata in the `_embedded` property. Supported value: `classification`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ListUsers(context.Background()).ContentType(contentType).Search(search).Filter(filter).Q(q).After(after).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | Specifies the media type of the resource. Optional &#x60;okta-response&#x60; value can be included for performance optimization.  Complex DelAuth configurations may degrade performance when fetching specific parts of the response, and passing this parameter can omit these parts, bypassing the bottleneck.  Enum values for &#x60;okta-response&#x60;:   * &#x60;omitCredentials&#x60;: Omits the credentials subobject from the response.   * &#x60;omitCredentialsLinks&#x60;: Omits the following HAL links from the response: Update password, Change recovery question, Start forgot password flow, Reset password, Reset factors, Unlock.   * &#x60;omitTransitioningToStatus&#x60;: Omits the &#x60;transitioningToStatus&#x60; field from the response. | 
 **search** | **string** | Searches for users with a supported filtering expression for most properties. Okta recommends this query parameter because it provides the largest range of search options and optimal performance.  &gt; **Note:** Using an overly complex or long search query can result in an error.  This operation supports [pagination](https://developer.okta.com/docs/api/#pagination). Use an ID lookup for records that you update to ensure your results contain the latest data. Returned users include those with the &#x60;DEPROVISIONED&#x60; status.  Property names in the search parameter are case sensitive, whereas operators (&#x60;eq&#x60;, &#x60;sw&#x60;, and so on) and string values are case insensitive. Unlike with user logins, diacritical marks are significant in search string values: a search for &#x60;isaac.brock&#x60; finds &#x60;Isaac.Brock&#x60;, but doesn&#39;t find a property whose value is &#x60;isáàc.bröck&#x60;.  This operation requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). See [Special characters](https://developer.okta.com/docs/api/#special-characters).  This operation searches many properties:   * Any user profile attribute, including custom-defined attributes   * The top-level properties: &#x60;id&#x60;, &#x60;status&#x60;, &#x60;created&#x60;, &#x60;activated&#x60;, &#x60;statusChanged&#x60;, and &#x60;lastUpdated&#x60;   * The [user type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/#tag/UserType/operation/updateUserType) accessed as &#x60;type.id&#x60;   * Properties that have array values  &gt; **Note:** &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; The ability to search by user classification is available as an [Early Access](https://developer.okta.com/docs/api/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature. The &#x60;classification.type&#x60; property cannot be used in conjunction with other search terms. You can search using &#x60;classification.type eq \&quot;LITE\&quot;&#x60; or &#x60;classification.type eq \&quot;STANDARD\&quot;&#x60;.  You can also use &#x60;sortBy&#x60; and &#x60;sortOrder&#x60; parameters. The &#x60;ne&#x60; (not equal) operator isn&#39;t supported, but you can obtain the same result by using &#x60;lt ... or ... gt&#x60;. For example, to see all users except those that have a status of &#x60;STAGED&#x60;, use &#x60;(status lt \&quot;STAGED\&quot; or status gt \&quot;STAGED\&quot;)&#x60;.  You can search properties that are arrays. If any element matches the search term, the entire array (object) is returned. Okta follows the [SCIM Protocol Specification](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) for searching arrays. You can search multiple arrays, multiple values in an array, as well as using the standard logical and filtering operators. See [Filter](https://developer.okta.com/docs/reference/core-okta-api/#filter).  Searches for users can be filtered by the following operators: &#x60;sw&#x60;, &#x60;eq&#x60;, and &#x60;co&#x60;. You can only use &#x60;co&#x60; with these select user profile attributes: &#x60;profile.firstName&#x60;, &#x60;profile.lastName&#x60;, &#x60;profile.email&#x60;, and &#x60;profile.login&#x60;. See [Operators](https://developer.okta.com/docs/api/#operators). | 
 **filter** | **string** | Filters users with a supported expression for a subset of properties.  &gt; **Note:** Returned users include those with the &#x60;DEPROVISIONED&#x60; status.  This requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). For example, &#x60;filter&#x3D;lastUpdated gt \&quot;2013-06-01T00:00:00.000Z\&quot;&#x60; is encoded as &#x60;filter&#x3D;lastUpdated%20gt%20%222013-06-01T00:00:00.000Z%22&#x60;. Filtering is case-sensitive for property names and query values, while operators are case-insensitive.  Filtering supports the following limited number of properties: &#x60;status&#x60;, &#x60;lastUpdated&#x60;, &#x60;id&#x60;, &#x60;profile.login&#x60;, &#x60;profile.email&#x60;, &#x60;profile.firstName&#x60;, and &#x60;profile.lastName&#x60;.  Additionally, filtering supports only the equal &#x60;eq&#x60; operator from the standard Okta API filtering semantics, except in the case of the &#x60;lastUpdated&#x60; property. This property can also use the inequality operators (&#x60;gt&#x60;, &#x60;ge&#x60;, &#x60;lt&#x60;, and &#x60;le&#x60;). For logical operators, only the logical operators &#x60;and&#x60; and &#x60;or&#x60; are supported. The &#x60;not&#x60; operator isn&#39;t supported. See [Filter](https://developer.okta.com/docs/api/#filter) and [Operators](https://developer.okta.com/docs/api/#operators). | 
 **q** | **string** | Finds users who match the specified query. Use the &#x60;q&#x60; parameter for simple queries, such as a lookup of users by name when creating a people picker.  The value of &#x60;q&#x60; is matched against &#x60;firstName&#x60;, &#x60;lastName&#x60;, or &#x60;email&#x60;. This performs a &#x60;startsWith&#x60; match, but this is an implementation detail and can change without notice. You don&#39;t need to specify &#x60;firstName&#x60;, &#x60;lastName&#x60;, or &#x60;email&#x60;.  &gt; **Notes:** &gt; * Using the &#x60;q&#x60; parameter in a request omits users that have a status of &#x60;DEPROVISIONED&#x60;. To return all users, use a &#x60;filter&#x60; or &#x60;search&#x60; query instead. &gt; * This doesn&#39;t support pagination, but you can use &#x60;limit&#x60;. &gt; * This isn&#39;t designed for large data sets. For optimal performance, use the &#x60;search&#x60; parameter instead. | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
 **limit** | **int32** | Specifies the number of results returned. Defaults to 10 if &#x60;q&#x60; is provided.  You can use &#x60;limit&#x60; with &#x60;after&#x60; to define the cursor location in the data set and manage the user records per page. | [default to 200]
 **sortBy** | **string** | Specifies the field to sort by (for search queries only). This can be any single property, for example &#x60;sortBy&#x3D;profile.lastName&#x60;. Users with the same value for the &#x60;sortBy&#x60; property are ordered by &#x60;id&#x60;. Use with &#x60;sortOrder&#x60; to control the order of results. | 
 **sortOrder** | **string** | Specifies sort order: &#x60;asc&#x60; or &#x60;desc&#x60; (for search queries only). This parameter is ignored if &#x60;sortBy&#x60; isn&#39;t present. | [default to &quot;asc&quot;]
 **fields** | **string** | Specifies a select set of user properties to query. Any other properties will be filtered out of the returned users. This is often called field projections in APIs, which can reduce payload size, improve performance, and limit unneccessary data exposure.  Requested fields should be comma-separated. Comma-separate the fields and place sub-fields in the profile object inside a &#x60;profile:()&#x60; directive, for example &#x60;profile:(firstName, city)&#x60;. The &#x60;id&#x60; field is always included, regardless of whether it&#39;s specified in the &#x60;fields&#x60; parameter. | 
 **expand** | **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;A parameter to include metadata in the &#x60;_embedded&#x60; property. Supported value: &#x60;classification&#x60;. | 

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


## ReplaceUser

> User ReplaceUser(ctx, id).User(user).Strict(strict).IfMatch(ifMatch).Execute()

Replace a user



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
	user := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 
	strict := true // bool | If `true`, validates against minimum age and history password policy (optional)
	ifMatch := "W/"1234567890abcdef"" // string | The ETag value of the user's expected current state. This becomes a conditional request used for concurrency control. See [Conditional Requests and Entity Tags](/#conditional-requests-and-entity-tags). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ReplaceUser(context.Background(), id).User(user).Strict(strict).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ReplaceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ReplaceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 
 **strict** | **bool** | If &#x60;true&#x60;, validates against minimum age and history password policy | 
 **ifMatch** | **string** | The ETag value of the user&#39;s expected current state. This becomes a conditional request used for concurrency control. See [Conditional Requests and Entity Tags](/#conditional-requests-and-entity-tags). | 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, id).User(user).Strict(strict).IfMatch(ifMatch).Execute()

Update a user



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
	user := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 
	strict := true // bool | If true, validates against minimum age and history password policy (optional)
	ifMatch := "W/"1234567890abcdef"" // string | The ETag value of the user's expected current state. This becomes a conditional request used for concurrency control. See [Conditional Requests and Entity Tags](/#conditional-requests-and-entity-tags). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateUser(context.Background(), id).User(user).Strict(strict).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 
 **strict** | **bool** | If true, validates against minimum age and history password policy | 
 **ifMatch** | **string** | The ETag value of the user&#39;s expected current state. This becomes a conditional request used for concurrency control. See [Conditional Requests and Entity Tags](/#conditional-requests-and-entity-tags). | 

### Return type

[**User**](User.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

