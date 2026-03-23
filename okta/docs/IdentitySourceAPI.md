# \IdentitySourceAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentitySourceGroups**](IdentitySourceAPI.md#CreateIdentitySourceGroups) | **Post** /api/v1/identity-sources/{identitySourceId}/groups | Create an identity source group
[**CreateIdentitySourceGroupsMemberships**](IdentitySourceAPI.md#CreateIdentitySourceGroupsMemberships) | **Post** /api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}/membership | Create the memberships for the given identity source group
[**CreateIdentitySourceSession**](IdentitySourceAPI.md#CreateIdentitySourceSession) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions | Create an identity source session
[**CreateIdentitySourceUser**](IdentitySourceAPI.md#CreateIdentitySourceUser) | **Post** /api/v1/identity-sources/{identitySourceId}/users | Create an identity source user
[**DeleteIdentitySourceGroup**](IdentitySourceAPI.md#DeleteIdentitySourceGroup) | **Delete** /api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId} | Delete an identity source group
[**DeleteIdentitySourceGroupMemberships**](IdentitySourceAPI.md#DeleteIdentitySourceGroupMemberships) | **Delete** /api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}/membership/{memberExternalId} | Delete the memberships for the specified identity source group
[**DeleteIdentitySourceSession**](IdentitySourceAPI.md#DeleteIdentitySourceSession) | **Delete** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId} | Delete an identity source session
[**DeleteIdentitySourceUser**](IdentitySourceAPI.md#DeleteIdentitySourceUser) | **Delete** /api/v1/identity-sources/{identitySourceId}/users/{externalId} | Delete an identity source user
[**GetIdentitySourceGroup**](IdentitySourceAPI.md#GetIdentitySourceGroup) | **Get** /api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId} | Retrieve an identity source group
[**GetIdentitySourceGroupMemberships**](IdentitySourceAPI.md#GetIdentitySourceGroupMemberships) | **Get** /api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}/membership | Retrieve the memberships for the given identity source group
[**GetIdentitySourceSession**](IdentitySourceAPI.md#GetIdentitySourceSession) | **Get** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId} | Retrieve an identity source session
[**GetIdentitySourceUser**](IdentitySourceAPI.md#GetIdentitySourceUser) | **Get** /api/v1/identity-sources/{identitySourceId}/users/{externalId} | Retrieve an identity source user
[**ListIdentitySourceSessions**](IdentitySourceAPI.md#ListIdentitySourceSessions) | **Get** /api/v1/identity-sources/{identitySourceId}/sessions | List all identity source sessions
[**ReplaceExistingIdentitySourceUser**](IdentitySourceAPI.md#ReplaceExistingIdentitySourceUser) | **Put** /api/v1/identity-sources/{identitySourceId}/users/{externalId} | Replace an existing identity source user
[**StartImportFromIdentitySource**](IdentitySourceAPI.md#StartImportFromIdentitySource) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/start-import | Start the import from the identity source
[**UpdateIdentitySourceGroups**](IdentitySourceAPI.md#UpdateIdentitySourceGroups) | **Post** /api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId} | Update an identity source group
[**UpdateIdentitySourceUsers**](IdentitySourceAPI.md#UpdateIdentitySourceUsers) | **Patch** /api/v1/identity-sources/{identitySourceId}/users/{externalId} | Update an identity source user
[**UploadIdentitySourceDataForDelete**](IdentitySourceAPI.md#UploadIdentitySourceDataForDelete) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-delete | Upload the data to be deleted in Okta
[**UploadIdentitySourceDataForUpsert**](IdentitySourceAPI.md#UploadIdentitySourceDataForUpsert) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-upsert | Upload the data to be upserted in Okta
[**UploadIdentitySourceGroupMembershipsForDelete**](IdentitySourceAPI.md#UploadIdentitySourceGroupMembershipsForDelete) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-group-memberships-delete | Upload the group memberships to be deleted in Okta
[**UploadIdentitySourceGroupMembershipsForUpsert**](IdentitySourceAPI.md#UploadIdentitySourceGroupMembershipsForUpsert) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-group-memberships-upsert | Upload the group memberships to be upserted in Okta
[**UploadIdentitySourceGroupsDataForDelete**](IdentitySourceAPI.md#UploadIdentitySourceGroupsDataForDelete) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-groups-delete | Upload the group external IDs to be deleted in Okta
[**UploadIdentitySourceGroupsForUpsert**](IdentitySourceAPI.md#UploadIdentitySourceGroupsForUpsert) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-groups-upsert | Upload the group profiles without memberships to be upserted in Okta



## CreateIdentitySourceGroups

> GroupsResponseSchema CreateIdentitySourceGroups(ctx, identitySourceId).GroupsRequestSchema(groupsRequestSchema).Execute()

Create an identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupsRequestSchema := *openapiclient.NewGroupsRequestSchema() // GroupsRequestSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.CreateIdentitySourceGroups(context.Background(), identitySourceId).GroupsRequestSchema(groupsRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.CreateIdentitySourceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentitySourceGroups`: GroupsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.CreateIdentitySourceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitySourceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsRequestSchema** | [**GroupsRequestSchema**](GroupsRequestSchema.md) |  | 

### Return type

[**GroupsResponseSchema**](GroupsResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentitySourceGroupsMemberships

> CreateIdentitySourceGroupsMemberships(ctx, identitySourceId, groupOrExternalId).MembershipRequestSchema(membershipRequestSchema).Execute()

Create the memberships for the given identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupOrExternalId := "00gsl4xM9ys8TdnbZ0g4 or GROUPEXT123456784C2IF" // string | The Okta group ID or external ID of the identity source group
	membershipRequestSchema := *openapiclient.NewMembershipRequestSchema() // MembershipRequestSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.CreateIdentitySourceGroupsMemberships(context.Background(), identitySourceId, groupOrExternalId).MembershipRequestSchema(membershipRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.CreateIdentitySourceGroupsMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**groupOrExternalId** | **string** | The Okta group ID or external ID of the identity source group | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitySourceGroupsMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **membershipRequestSchema** | [**MembershipRequestSchema**](MembershipRequestSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentitySourceSession

> IdentitySourceSession CreateIdentitySourceSession(ctx, identitySourceId).Execute()

Create an identity source session



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.CreateIdentitySourceSession(context.Background(), identitySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.CreateIdentitySourceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentitySourceSession`: IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.CreateIdentitySourceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitySourceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentitySourceUser

> CreateIdentitySourceUser(ctx, identitySourceId).UserRequestSchema(userRequestSchema).Execute()

Create an identity source user



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	userRequestSchema := *openapiclient.NewUserRequestSchema() // UserRequestSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.CreateIdentitySourceUser(context.Background(), identitySourceId).UserRequestSchema(userRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.CreateIdentitySourceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitySourceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRequestSchema** | [**UserRequestSchema**](UserRequestSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentitySourceGroup

> DeleteIdentitySourceGroup(ctx, identitySourceId, groupOrExternalId).Execute()

Delete an identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupOrExternalId := "00gsl4xM9ys8TdnbZ0g4 or GROUPEXT123456784C2IF" // string | The Okta group ID or external ID of the identity source group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.DeleteIdentitySourceGroup(context.Background(), identitySourceId, groupOrExternalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.DeleteIdentitySourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**groupOrExternalId** | **string** | The Okta group ID or external ID of the identity source group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitySourceGroupRequest struct via the builder pattern


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


## DeleteIdentitySourceGroupMemberships

> DeleteIdentitySourceGroupMemberships(ctx, identitySourceId, groupOrExternalId, memberExternalId).Execute()

Delete the memberships for the specified identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupOrExternalId := "00gsl4xM9ys8TdnbZ0g4 or GROUPEXT123456784C2IF" // string | The Okta group ID or external ID of the identity source group
	memberExternalId := "USEREXT123456784C2IFA" // string | The external ID of the identity source user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.DeleteIdentitySourceGroupMemberships(context.Background(), identitySourceId, groupOrExternalId, memberExternalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.DeleteIdentitySourceGroupMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**groupOrExternalId** | **string** | The Okta group ID or external ID of the identity source group | 
**memberExternalId** | **string** | The external ID of the identity source user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitySourceGroupMembershipsRequest struct via the builder pattern


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


## DeleteIdentitySourceSession

> DeleteIdentitySourceSession(ctx, identitySourceId, sessionId).Execute()

Delete an identity source session



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.DeleteIdentitySourceSession(context.Background(), identitySourceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.DeleteIdentitySourceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitySourceSessionRequest struct via the builder pattern


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


## DeleteIdentitySourceUser

> DeleteIdentitySourceUser(ctx, identitySourceId, externalId).Execute()

Delete an identity source user



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	externalId := "00u7m9p9ZT8k2S2EX1f7" // string | The external ID of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.DeleteIdentitySourceUser(context.Background(), identitySourceId, externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.DeleteIdentitySourceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**externalId** | **string** | The external ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitySourceUserRequest struct via the builder pattern


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


## GetIdentitySourceGroup

> GroupsResponseSchema GetIdentitySourceGroup(ctx, identitySourceId, groupOrExternalId).Execute()

Retrieve an identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupOrExternalId := "00gsl4xM9ys8TdnbZ0g4 or GROUPEXT123456784C2IF" // string | The Okta group ID or external ID of the identity source group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.GetIdentitySourceGroup(context.Background(), identitySourceId, groupOrExternalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.GetIdentitySourceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitySourceGroup`: GroupsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.GetIdentitySourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**groupOrExternalId** | **string** | The Okta group ID or external ID of the identity source group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupsResponseSchema**](GroupsResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitySourceGroupMemberships

> GroupMembershipsResponseSchema GetIdentitySourceGroupMemberships(ctx, identitySourceId, groupOrExternalId).After(after).Limit(limit).Execute()

Retrieve the memberships for the given identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupOrExternalId := "00gsl4xM9ys8TdnbZ0g4 or GROUPEXT123456784C2IF" // string | The Okta group ID or external ID of the identity source group
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)
	limit := int32(56) // int32 | Specifies the number of group membership results in a page. Okta recommends using a specific value other than the default or maximum. If your request times out, retry your request with a smaller `limit` and [page the results](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.GetIdentitySourceGroupMemberships(context.Background(), identitySourceId, groupOrExternalId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.GetIdentitySourceGroupMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitySourceGroupMemberships`: GroupMembershipsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.GetIdentitySourceGroupMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**groupOrExternalId** | **string** | The Okta group ID or external ID of the identity source group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySourceGroupMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
 **limit** | **int32** | Specifies the number of group membership results in a page. Okta recommends using a specific value other than the default or maximum. If your request times out, retry your request with a smaller &#x60;limit&#x60; and [page the results](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**GroupMembershipsResponseSchema**](GroupMembershipsResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitySourceSession

> IdentitySourceSession GetIdentitySourceSession(ctx, identitySourceId, sessionId).Execute()

Retrieve an identity source session



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.GetIdentitySourceSession(context.Background(), identitySourceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.GetIdentitySourceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitySourceSession`: IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.GetIdentitySourceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySourceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitySourceUser

> UserResponseSchema GetIdentitySourceUser(ctx, identitySourceId, externalId).Execute()

Retrieve an identity source user



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	externalId := "00u7m9p9ZT8k2S2EX1f7" // string | The external ID of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.GetIdentitySourceUser(context.Background(), identitySourceId, externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.GetIdentitySourceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitySourceUser`: UserResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.GetIdentitySourceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**externalId** | **string** | The external ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySourceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserResponseSchema**](UserResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentitySourceSessions

> []IdentitySourceSession ListIdentitySourceSessions(ctx, identitySourceId).Execute()

List all identity source sessions



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.ListIdentitySourceSessions(context.Background(), identitySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.ListIdentitySourceSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentitySourceSessions`: []IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.ListIdentitySourceSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitySourceSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceExistingIdentitySourceUser

> UserResponseSchema ReplaceExistingIdentitySourceUser(ctx, identitySourceId, externalId).UserRequestSchema(userRequestSchema).Execute()

Replace an existing identity source user



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	externalId := "00u7m9p9ZT8k2S2EX1f7" // string | The external ID of the user
	userRequestSchema := *openapiclient.NewUserRequestSchema() // UserRequestSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.ReplaceExistingIdentitySourceUser(context.Background(), identitySourceId, externalId).UserRequestSchema(userRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.ReplaceExistingIdentitySourceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceExistingIdentitySourceUser`: UserResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.ReplaceExistingIdentitySourceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**externalId** | **string** | The external ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceExistingIdentitySourceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userRequestSchema** | [**UserRequestSchema**](UserRequestSchema.md) |  | 

### Return type

[**UserResponseSchema**](UserResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartImportFromIdentitySource

> IdentitySourceSession StartImportFromIdentitySource(ctx, identitySourceId, sessionId).Execute()

Start the import from the identity source



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.StartImportFromIdentitySource(context.Background(), identitySourceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.StartImportFromIdentitySource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartImportFromIdentitySource`: IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.StartImportFromIdentitySource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartImportFromIdentitySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentitySourceGroups

> GroupsResponseSchema UpdateIdentitySourceGroups(ctx, identitySourceId, groupOrExternalId).GroupsRequestSchema(groupsRequestSchema).Execute()

Update an identity source group



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	groupOrExternalId := "00gsl4xM9ys8TdnbZ0g4 or GROUPEXT123456784C2IF" // string | The Okta group ID or external ID of the identity source group
	groupsRequestSchema := *openapiclient.NewGroupsRequestSchema() // GroupsRequestSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.UpdateIdentitySourceGroups(context.Background(), identitySourceId, groupOrExternalId).GroupsRequestSchema(groupsRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UpdateIdentitySourceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentitySourceGroups`: GroupsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.UpdateIdentitySourceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**groupOrExternalId** | **string** | The Okta group ID or external ID of the identity source group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitySourceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupsRequestSchema** | [**GroupsRequestSchema**](GroupsRequestSchema.md) |  | 

### Return type

[**GroupsResponseSchema**](GroupsResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentitySourceUsers

> UserResponseSchema UpdateIdentitySourceUsers(ctx, identitySourceId, externalId).UsersUpdateRequestSchema(usersUpdateRequestSchema).Execute()

Update an identity source user



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	externalId := "00u7m9p9ZT8k2S2EX1f7" // string | The external ID of the user
	usersUpdateRequestSchema := *openapiclient.NewUsersUpdateRequestSchema() // UsersUpdateRequestSchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.UpdateIdentitySourceUsers(context.Background(), identitySourceId, externalId).UsersUpdateRequestSchema(usersUpdateRequestSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UpdateIdentitySourceUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentitySourceUsers`: UserResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.UpdateIdentitySourceUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**externalId** | **string** | The external ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitySourceUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **usersUpdateRequestSchema** | [**UsersUpdateRequestSchema**](UsersUpdateRequestSchema.md) |  | 

### Return type

[**UserResponseSchema**](UserResponseSchema.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceDataForDelete

> UploadIdentitySourceDataForDelete(ctx, identitySourceId, sessionId).BulkDeleteRequestBody(bulkDeleteRequestBody).Execute()

Upload the data to be deleted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkDeleteRequestBody := *openapiclient.NewBulkDeleteRequestBody() // BulkDeleteRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceDataForDelete(context.Background(), identitySourceId, sessionId).BulkDeleteRequestBody(bulkDeleteRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceDataForDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceDataForDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkDeleteRequestBody** | [**BulkDeleteRequestBody**](BulkDeleteRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceDataForUpsert

> UploadIdentitySourceDataForUpsert(ctx, identitySourceId, sessionId).BulkUpsertRequestBody(bulkUpsertRequestBody).Execute()

Upload the data to be upserted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkUpsertRequestBody := *openapiclient.NewBulkUpsertRequestBody() // BulkUpsertRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceDataForUpsert(context.Background(), identitySourceId, sessionId).BulkUpsertRequestBody(bulkUpsertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceDataForUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceDataForUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkUpsertRequestBody** | [**BulkUpsertRequestBody**](BulkUpsertRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupMembershipsForDelete

> UploadIdentitySourceGroupMembershipsForDelete(ctx, identitySourceId, sessionId).BulkGroupMembershipsDeleteRequestBody(bulkGroupMembershipsDeleteRequestBody).Execute()

Upload the group memberships to be deleted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupMembershipsDeleteRequestBody := *openapiclient.NewBulkGroupMembershipsDeleteRequestBody() // BulkGroupMembershipsDeleteRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForDelete(context.Background(), identitySourceId, sessionId).BulkGroupMembershipsDeleteRequestBody(bulkGroupMembershipsDeleteRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupMembershipsForDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupMembershipsDeleteRequestBody** | [**BulkGroupMembershipsDeleteRequestBody**](BulkGroupMembershipsDeleteRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupMembershipsForUpsert

> UploadIdentitySourceGroupMembershipsForUpsert(ctx, identitySourceId, sessionId).BulkGroupMembershipsUpsertRequestBody(bulkGroupMembershipsUpsertRequestBody).Execute()

Upload the group memberships to be upserted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupMembershipsUpsertRequestBody := *openapiclient.NewBulkGroupMembershipsUpsertRequestBody() // BulkGroupMembershipsUpsertRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForUpsert(context.Background(), identitySourceId, sessionId).BulkGroupMembershipsUpsertRequestBody(bulkGroupMembershipsUpsertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupMembershipsForUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupMembershipsUpsertRequestBody** | [**BulkGroupMembershipsUpsertRequestBody**](BulkGroupMembershipsUpsertRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupsDataForDelete

> UploadIdentitySourceGroupsDataForDelete(ctx, identitySourceId, sessionId).BulkGroupDeleteRequestBody(bulkGroupDeleteRequestBody).Execute()

Upload the group external IDs to be deleted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupDeleteRequestBody := *openapiclient.NewBulkGroupDeleteRequestBody() // BulkGroupDeleteRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupsDataForDelete(context.Background(), identitySourceId, sessionId).BulkGroupDeleteRequestBody(bulkGroupDeleteRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupsDataForDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupsDataForDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupDeleteRequestBody** | [**BulkGroupDeleteRequestBody**](BulkGroupDeleteRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupsForUpsert

> UploadIdentitySourceGroupsForUpsert(ctx, identitySourceId, sessionId).BulkGroupUpsertRequestBody(bulkGroupUpsertRequestBody).Execute()

Upload the group profiles without memberships to be upserted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupUpsertRequestBody := *openapiclient.NewBulkGroupUpsertRequestBody() // BulkGroupUpsertRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupsForUpsert(context.Background(), identitySourceId, sessionId).BulkGroupUpsertRequestBody(bulkGroupUpsertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupsForUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupsForUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupUpsertRequestBody** | [**BulkGroupUpsertRequestBody**](BulkGroupUpsertRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

