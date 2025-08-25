# \ApplicationGroupsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignGroupToApplication**](ApplicationGroupsAPI.md#AssignGroupToApplication) | **Put** /api/v1/apps/{appId}/groups/{groupId} | Assign an Application Group
[**GetApplicationGroupAssignment**](ApplicationGroupsAPI.md#GetApplicationGroupAssignment) | **Get** /api/v1/apps/{appId}/groups/{groupId} | Retrieve an Application Group
[**ListApplicationGroupAssignments**](ApplicationGroupsAPI.md#ListApplicationGroupAssignments) | **Get** /api/v1/apps/{appId}/groups | List all Application Groups
[**UnassignApplicationFromGroup**](ApplicationGroupsAPI.md#UnassignApplicationFromGroup) | **Delete** /api/v1/apps/{appId}/groups/{groupId} | Unassign an Application Group
[**UpdateGroupAssignmentToApplication**](ApplicationGroupsAPI.md#UpdateGroupAssignmentToApplication) | **Patch** /api/v1/apps/{appId}/groups/{groupId} | Update an Application Group



## AssignGroupToApplication

> ApplicationGroupAssignment AssignGroupToApplication(ctx, appId, groupId).ApplicationGroupAssignment(applicationGroupAssignment).Execute()

Assign an Application Group



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
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    applicationGroupAssignment := *openapiclient.NewApplicationGroupAssignment() // ApplicationGroupAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPI.AssignGroupToApplication(context.Background(), appId, groupId).ApplicationGroupAssignment(applicationGroupAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPI.AssignGroupToApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignGroupToApplication`: ApplicationGroupAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPI.AssignGroupToApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupToApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationGroupAssignment** | [**ApplicationGroupAssignment**](ApplicationGroupAssignment.md) |  | 

### Return type

[**ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationGroupAssignment

> ApplicationGroupAssignment GetApplicationGroupAssignment(ctx, appId, groupId).Expand(expand).Execute()

Retrieve an Application Group



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
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    expand := "group" // string | An optional query parameter to return the corresponding assigned [Group](/openapi/okta-management/management/tag/Group/) or the group assignment metadata details in the `_embedded` property. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPI.GetApplicationGroupAssignment(context.Background(), appId, groupId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPI.GetApplicationGroupAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationGroupAssignment`: ApplicationGroupAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPI.GetApplicationGroupAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationGroupAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | An optional query parameter to return the corresponding assigned [Group](/openapi/okta-management/management/tag/Group/) or the group assignment metadata details in the &#x60;_embedded&#x60; property. | 

### Return type

[**ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationGroupAssignments

> []ApplicationGroupAssignment ListApplicationGroupAssignments(ctx, appId).Q(q).After(after).Limit(limit).Expand(expand).Execute()

List all Application Groups



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
    q := "test" // string | Specifies a filter for a list of assigned groups returned based on their names. The value of `q` is matched against the group `name`. This filter only supports the `startsWith` operation that matches the `q` string against the beginning of the [Group name](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/listGroups!c=200&path=profile/name&t=response). (optional)
    after := "16275000448691" // string | Specifies the pagination cursor for the `next` page of results. Treat this as an opaque value obtained through the next link relationship. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
    limit := int32(20) // int32 | Specifies the number of objects to return per page. If there are multiple pages of results, the Link header contains a `next` link that you need to use as an opaque value (follow it, don't parse it). See [Pagination](/#pagination). (optional) (default to 20)
    expand := "group" // string | An optional query parameter to return the corresponding assigned [Group](/openapi/okta-management/management/tag/Group/) or the group assignment metadata details in the `_embedded` property. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPI.ListApplicationGroupAssignments(context.Background(), appId).Q(q).After(after).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPI.ListApplicationGroupAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationGroupAssignments`: []ApplicationGroupAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPI.ListApplicationGroupAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationGroupAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Specifies a filter for a list of assigned groups returned based on their names. The value of &#x60;q&#x60; is matched against the group &#x60;name&#x60;. This filter only supports the &#x60;startsWith&#x60; operation that matches the &#x60;q&#x60; string against the beginning of the [Group name](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/listGroups!c&#x3D;200&amp;path&#x3D;profile/name&amp;t&#x3D;response). | 
 **after** | **string** | Specifies the pagination cursor for the &#x60;next&#x60; page of results. Treat this as an opaque value obtained through the next link relationship. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | Specifies the number of objects to return per page. If there are multiple pages of results, the Link header contains a &#x60;next&#x60; link that you need to use as an opaque value (follow it, don&#39;t parse it). See [Pagination](/#pagination). | [default to 20]
 **expand** | **string** | An optional query parameter to return the corresponding assigned [Group](/openapi/okta-management/management/tag/Group/) or the group assignment metadata details in the &#x60;_embedded&#x60; property. | 

### Return type

[**[]ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignApplicationFromGroup

> UnassignApplicationFromGroup(ctx, appId, groupId).Execute()

Unassign an Application Group



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
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPI.UnassignApplicationFromGroup(context.Background(), appId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPI.UnassignApplicationFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignApplicationFromGroupRequest struct via the builder pattern


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


## UpdateGroupAssignmentToApplication

> ApplicationGroupAssignment UpdateGroupAssignmentToApplication(ctx, appId, groupId).JsonPatchOperation(jsonPatchOperation).Execute()

Update an Application Group



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
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation()} // []JsonPatchOperation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPI.UpdateGroupAssignmentToApplication(context.Background(), appId, groupId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPI.UpdateGroupAssignmentToApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupAssignmentToApplication`: ApplicationGroupAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPI.UpdateGroupAssignmentToApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupAssignmentToApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

