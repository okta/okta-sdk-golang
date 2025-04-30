# \GroupAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateGroupRule**](GroupAPI.md#ActivateGroupRule) | **Post** /api/v1/groups/rules/{groupRuleId}/lifecycle/activate | Activate a Group Rule
[**AssignUserToGroup**](GroupAPI.md#AssignUserToGroup) | **Put** /api/v1/groups/{groupId}/users/{userId} | Assign a User
[**CreateGroup**](GroupAPI.md#CreateGroup) | **Post** /api/v1/groups | Create a Group
[**CreateGroupRule**](GroupAPI.md#CreateGroupRule) | **Post** /api/v1/groups/rules | Create a Group Rule
[**DeactivateGroupRule**](GroupAPI.md#DeactivateGroupRule) | **Post** /api/v1/groups/rules/{groupRuleId}/lifecycle/deactivate | Deactivate a Group Rule
[**DeleteGroup**](GroupAPI.md#DeleteGroup) | **Delete** /api/v1/groups/{groupId} | Delete a Group
[**DeleteGroupRule**](GroupAPI.md#DeleteGroupRule) | **Delete** /api/v1/groups/rules/{groupRuleId} | Delete a group Rule
[**GetGroup**](GroupAPI.md#GetGroup) | **Get** /api/v1/groups/{groupId} | Retrieve a Group
[**GetGroupRule**](GroupAPI.md#GetGroupRule) | **Get** /api/v1/groups/rules/{groupRuleId} | Retrieve a Group Rule
[**ListAssignedApplicationsForGroup**](GroupAPI.md#ListAssignedApplicationsForGroup) | **Get** /api/v1/groups/{groupId}/apps | List all Assigned Applications
[**ListGroupRules**](GroupAPI.md#ListGroupRules) | **Get** /api/v1/groups/rules | List all Group Rules
[**ListGroupUsers**](GroupAPI.md#ListGroupUsers) | **Get** /api/v1/groups/{groupId}/users | List all Member Users
[**ListGroups**](GroupAPI.md#ListGroups) | **Get** /api/v1/groups | List all Groups
[**ReplaceGroup**](GroupAPI.md#ReplaceGroup) | **Put** /api/v1/groups/{groupId} | Replace a Group
[**ReplaceGroupRule**](GroupAPI.md#ReplaceGroupRule) | **Put** /api/v1/groups/rules/{groupRuleId} | Replace a Group Rule
[**UnassignUserFromGroup**](GroupAPI.md#UnassignUserFromGroup) | **Delete** /api/v1/groups/{groupId}/users/{userId} | Unassign a User



## ActivateGroupRule

> ActivateGroupRule(ctx, groupRuleId).Execute()

Activate a Group Rule



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
    groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupAPI.ActivateGroupRule(context.Background(), groupRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ActivateGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateGroupRuleRequest struct via the builder pattern


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


## AssignUserToGroup

> AssignUserToGroup(ctx, groupId, userId).Execute()

Assign a User



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
    userId := "userId_example" // string | ID of an existing Okta user

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


## CreateGroup

> Group CreateGroup(ctx).Group(group).Execute()

Create a Group



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
    group := *openapiclient.NewGroup() // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.CreateGroup(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) |  | 

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


## CreateGroupRule

> GroupRule CreateGroupRule(ctx).GroupRule(groupRule).Execute()

Create a Group Rule



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
    groupRule := *openapiclient.NewGroupRule() // GroupRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.CreateGroupRule(context.Background()).GroupRule(groupRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupRule`: GroupRule
    fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateGroupRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRule** | [**GroupRule**](GroupRule.md) |  | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateGroupRule

> DeactivateGroupRule(ctx, groupRuleId).Execute()

Deactivate a Group Rule



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
    groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupAPI.DeactivateGroupRule(context.Background(), groupRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeactivateGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateGroupRuleRequest struct via the builder pattern


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

Delete a Group



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


## DeleteGroupRule

> DeleteGroupRule(ctx, groupRuleId).RemoveUsers(removeUsers).Execute()

Delete a group Rule



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
    groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule
    removeUsers := true // bool | Indicates whether to keep or remove users from groups assigned by this rule. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupAPI.DeleteGroupRule(context.Background(), groupRuleId).RemoveUsers(removeUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeUsers** | **bool** | Indicates whether to keep or remove users from groups assigned by this rule. | 

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

Retrieve a Group



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


## GetGroupRule

> GroupRule GetGroupRule(ctx, groupRuleId).Expand(expand).Execute()

Retrieve a Group Rule



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
    groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.GetGroupRule(context.Background(), groupRuleId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupRule`: GroupRule
    fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | 

### Return type

[**GroupRule**](GroupRule.md)

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

List all Assigned Applications



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


## ListGroupRules

> []GroupRule ListGroupRules(ctx).Limit(limit).After(after).Search(search).Expand(expand).Execute()

List all Group Rules



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
    limit := int32(56) // int32 | Specifies the number of rule results in a page (optional) (default to 50)
    after := "after_example" // string | Specifies the pagination cursor for the next page of rules (optional)
    search := "search_example" // string | Specifies the keyword to search fules for (optional)
    expand := "expand_example" // string | If specified as `groupIdToGroupNameMap`, then show group names (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.ListGroupRules(context.Background()).Limit(limit).After(after).Search(search).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupRules`: []GroupRule
    fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroupRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Specifies the number of rule results in a page | [default to 50]
 **after** | **string** | Specifies the pagination cursor for the next page of rules | 
 **search** | **string** | Specifies the keyword to search fules for | 
 **expand** | **string** | If specified as &#x60;groupIdToGroupNameMap&#x60;, then show group names | 

### Return type

[**[]GroupRule**](GroupRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupUsers

> []GroupMember ListGroupUsers(ctx, groupId).Search(search).SortBy(sortBy).SortOrder(sortOrder).After(after).Limit(limit).Execute()

List all Member Users



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
    search := "search_example" // string | Searches for users with a supported filtering expression for user name, primary email, or user name (optional)
    sortBy := "id" // string | Specifies which field to sort by. This can be any single property. (optional)
    sortOrder := "sortOrder_example" // string | Specifies sort order: `asc` or `desc`. This parameter is ignored if `sortBy` is not present. Users with the same value for the `sortBy` parameter are ordered by `id`. (optional) (default to "asc")
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | Specifies the number of user results in a page (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.ListGroupUsers(context.Background(), groupId).Search(search).SortBy(sortBy).SortOrder(sortOrder).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupUsers`: []GroupMember
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

 **search** | **string** | Searches for users with a supported filtering expression for user name, primary email, or user name | 
 **sortBy** | **string** | Specifies which field to sort by. This can be any single property. | 
 **sortOrder** | **string** | Specifies sort order: &#x60;asc&#x60; or &#x60;desc&#x60;. This parameter is ignored if &#x60;sortBy&#x60; is not present. Users with the same value for the &#x60;sortBy&#x60; parameter are ordered by &#x60;id&#x60;. | [default to &quot;asc&quot;]
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | Specifies the number of user results in a page | [default to 1000]

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> []Group ListGroups(ctx).Q(q).Filter(filter).After(after).Limit(limit).Expand(expand).Search(search).SortBy(sortBy).SortOrder(sortOrder).Execute()

List all Groups



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
    q := "q_example" // string | Searches the name property of groups for matching value (optional)
    filter := "filter_example" // string | Filter expression for groups (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of groups (optional)
    limit := int32(56) // int32 | Specifies the number of group results in a page (optional) (default to 10000)
    expand := "expand_example" // string | If specified, it causes additional metadata to be included in the response. (optional)
    search := "search_example" // string | Searches for groups with a supported filtering expression for all attributes except for _embedded, _links, and objectClass (optional)
    sortBy := "lastUpdated" // string | Specifies field to sort by and can be any single property (for search queries only). (optional)
    sortOrder := "sortOrder_example" // string | Specifies sort order `asc` or `desc` (for search queries only). This parameter is ignored if `sortBy` is not present. Groups with the same value for the `sortBy` parameter are ordered by `id`. (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.ListGroups(context.Background()).Q(q).Filter(filter).After(after).Limit(limit).Expand(expand).Search(search).SortBy(sortBy).SortOrder(sortOrder).Execute()
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
 **q** | **string** | Searches the name property of groups for matching value | 
 **filter** | **string** | Filter expression for groups | 
 **after** | **string** | Specifies the pagination cursor for the next page of groups | 
 **limit** | **int32** | Specifies the number of group results in a page | [default to 10000]
 **expand** | **string** | If specified, it causes additional metadata to be included in the response. | 
 **search** | **string** | Searches for groups with a supported filtering expression for all attributes except for _embedded, _links, and objectClass | 
 **sortBy** | **string** | Specifies field to sort by and can be any single property (for search queries only). | 
 **sortOrder** | **string** | Specifies sort order &#x60;asc&#x60; or &#x60;desc&#x60; (for search queries only). This parameter is ignored if &#x60;sortBy&#x60; is not present. Groups with the same value for the &#x60;sortBy&#x60; parameter are ordered by &#x60;id&#x60;. | [default to &quot;asc&quot;]

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

Replace a Group



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
    group := *openapiclient.NewGroup() // Group | 

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

 **group** | [**Group**](Group.md) |  | 

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


## ReplaceGroupRule

> GroupRule ReplaceGroupRule(ctx, groupRuleId).GroupRule(groupRule).Execute()

Replace a Group Rule



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
    groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule
    groupRule := *openapiclient.NewGroupRule() // GroupRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupAPI.ReplaceGroupRule(context.Background(), groupRuleId).GroupRule(groupRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ReplaceGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceGroupRule`: GroupRule
    fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ReplaceGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRule** | [**GroupRule**](GroupRule.md) |  | 

### Return type

[**GroupRule**](GroupRule.md)

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

Unassign a User



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
    userId := "userId_example" // string | ID of an existing Okta user

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

