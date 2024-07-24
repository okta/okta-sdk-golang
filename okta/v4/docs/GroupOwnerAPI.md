# \GroupOwnerAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignGroupOwner**](GroupOwnerAPI.md#AssignGroupOwner) | **Post** /api/v1/groups/{groupId}/owners | Assign a Group Owner
[**DeleteGroupOwner**](GroupOwnerAPI.md#DeleteGroupOwner) | **Delete** /api/v1/groups/{groupId}/owners/{ownerId} | Delete a Group Owner
[**ListGroupOwners**](GroupOwnerAPI.md#ListGroupOwners) | **Get** /api/v1/groups/{groupId}/owners | List all Group Owners



## AssignGroupOwner

> GroupOwner AssignGroupOwner(ctx, groupId).AssignGroupOwnerRequestBody(assignGroupOwnerRequestBody).Execute()

Assign a Group Owner



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
    assignGroupOwnerRequestBody := *openapiclient.NewAssignGroupOwnerRequestBody() // AssignGroupOwnerRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupOwnerAPI.AssignGroupOwner(context.Background(), groupId).AssignGroupOwnerRequestBody(assignGroupOwnerRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupOwnerAPI.AssignGroupOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignGroupOwner`: GroupOwner
    fmt.Fprintf(os.Stdout, "Response from `GroupOwnerAPI.AssignGroupOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignGroupOwnerRequestBody** | [**AssignGroupOwnerRequestBody**](AssignGroupOwnerRequestBody.md) |  | 

### Return type

[**GroupOwner**](GroupOwner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupOwner

> DeleteGroupOwner(ctx, groupId, ownerId).Execute()

Delete a Group Owner



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
    ownerId := "00u1emaK22TWRYd3TtG" // string | The `id` of the group owner

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupOwnerAPI.DeleteGroupOwner(context.Background(), groupId, ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupOwnerAPI.DeleteGroupOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 
**ownerId** | **string** | The &#x60;id&#x60; of the group owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupOwnerRequest struct via the builder pattern


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


## ListGroupOwners

> []GroupOwner ListGroupOwners(ctx, groupId).Search(search).After(after).Limit(limit).Execute()

List all Group Owners



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
    search := "search_example" // string | SCIM Filter expression for group owners. Allows to filter owners by type. (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of owners (optional)
    limit := int32(56) // int32 | Specifies the number of owner results in a page (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupOwnerAPI.ListGroupOwners(context.Background(), groupId).Search(search).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupOwnerAPI.ListGroupOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupOwners`: []GroupOwner
    fmt.Fprintf(os.Stdout, "Response from `GroupOwnerAPI.ListGroupOwners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The &#x60;id&#x60; of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | SCIM Filter expression for group owners. Allows to filter owners by type. | 
 **after** | **string** | Specifies the pagination cursor for the next page of owners | 
 **limit** | **int32** | Specifies the number of owner results in a page | [default to 1000]

### Return type

[**[]GroupOwner**](GroupOwner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

