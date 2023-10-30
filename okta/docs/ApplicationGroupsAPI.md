# \ApplicationGroupsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignGroupToApplication**](ApplicationGroupsAPI.md#AssignGroupToApplication) | **Put** /api/v1/apps/{appId}/groups/{groupId} | Assign a Group
[**GetApplicationGroupAssignment**](ApplicationGroupsAPI.md#GetApplicationGroupAssignment) | **Get** /api/v1/apps/{appId}/groups/{groupId} | Retrieve an Assigned Group
[**ListApplicationGroupAssignments**](ApplicationGroupsAPI.md#ListApplicationGroupAssignments) | **Get** /api/v1/apps/{appId}/groups | List all Assigned Groups
[**UnassignApplicationFromGroup**](ApplicationGroupsAPI.md#UnassignApplicationFromGroup) | **Delete** /api/v1/apps/{appId}/groups/{groupId} | Unassign a Group



## AssignGroupToApplication

> ApplicationGroupAssignment AssignGroupToApplication(ctx, appId, groupId).ApplicationGroupAssignment(applicationGroupAssignment).Execute()

Assign a Group



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
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
**appId** | **string** | ID of the Application | 
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

Retrieve an Assigned Group



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    groupId := "00g1emaKYZTWRYYRRTSK" // string | The `id` of the group
    expand := "expand_example" // string |  (optional)

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
**appId** | **string** | ID of the Application | 
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationGroupAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 

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

List all Assigned Groups



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    q := "q_example" // string |  (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of assignments (optional)
    limit := int32(56) // int32 | Specifies the number of results for a page (optional) (default to -1)
    expand := "expand_example" // string |  (optional)

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
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationGroupAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | 
 **after** | **string** | Specifies the pagination cursor for the next page of assignments | 
 **limit** | **int32** | Specifies the number of results for a page | [default to -1]
 **expand** | **string** |  | 

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

Unassign a Group



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
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
**appId** | **string** | ID of the Application | 
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

