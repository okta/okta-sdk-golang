# \ApplicationUsersAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignUserToApplication**](ApplicationUsersAPI.md#AssignUserToApplication) | **Post** /api/v1/apps/{appId}/users | Assign an Application User
[**GetApplicationUser**](ApplicationUsersAPI.md#GetApplicationUser) | **Get** /api/v1/apps/{appId}/users/{userId} | Retrieve an Application User
[**ListApplicationUsers**](ApplicationUsersAPI.md#ListApplicationUsers) | **Get** /api/v1/apps/{appId}/users | List all Application Users
[**UnassignUserFromApplication**](ApplicationUsersAPI.md#UnassignUserFromApplication) | **Delete** /api/v1/apps/{appId}/users/{userId} | Unassign an Application User
[**UpdateApplicationUser**](ApplicationUsersAPI.md#UpdateApplicationUser) | **Post** /api/v1/apps/{appId}/users/{userId} | Update an Application User



## AssignUserToApplication

> AppUser AssignUserToApplication(ctx, appId).AppUser(appUser).Execute()

Assign an Application User



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
    appUser := *openapiclient.NewAppUserAssignRequest("00u11z6WHMYCGPCHCRFK") // AppUserAssignRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersAPI.AssignUserToApplication(context.Background(), appId).AppUser(appUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersAPI.AssignUserToApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignUserToApplication`: AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersAPI.AssignUserToApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUserToApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUser** | [**AppUserAssignRequest**](AppUserAssignRequest.md) |  | 

### Return type

[**AppUser**](AppUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationUser

> AppUser GetApplicationUser(ctx, appId, userId).Expand(expand).Execute()

Retrieve an Application User



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
    userId := "00u13okQOVWZJGDOAUVR" // string | ID of an existing Okta user
    expand := "user" // string | An optional query parameter to return the corresponding [User](/openapi/okta-management/management/tag/User/) object in the `_embedded` property. Valid value: `user` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersAPI.GetApplicationUser(context.Background(), appId, userId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersAPI.GetApplicationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationUser`: AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersAPI.GetApplicationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | An optional query parameter to return the corresponding [User](/openapi/okta-management/management/tag/User/) object in the &#x60;_embedded&#x60; property. Valid value: &#x60;user&#x60; | 

### Return type

[**AppUser**](AppUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationUsers

> []AppUser ListApplicationUsers(ctx, appId).After(after).Limit(limit).Q(q).Expand(expand).Execute()

List all Application Users



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
    after := "16275000448691" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the next link relationship. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | Specifies the number of objects to return per page. If there are multiple pages of results, the Link header contains a `next` link that you need to use as an opaque value (follow it, don't parse it). See [Pagination](/#pagination).  (optional) (default to 50)
    q := "sam" // string | Specifies a filter for the list of Application Users returned based on their profile attributes. The value of `q` is matched against the beginning of the following profile attributes: `userName`, `firstName`, `lastName`, and `email`. This filter only supports the `startsWith` operation that matches the `q` string against the beginning of the attribute values. > **Note:** For OIDC apps, user profiles don't contain the `firstName` or `lastName` attributes. Therefore, the query only matches against the `userName` or `email` attributes.  (optional)
    expand := "user" // string | An optional query parameter to return the corresponding [User](/openapi/okta-management/management/tag/User/) object in the `_embedded` property. Valid value: `user` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersAPI.ListApplicationUsers(context.Background(), appId).After(after).Limit(limit).Q(q).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersAPI.ListApplicationUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationUsers`: []AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersAPI.ListApplicationUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the next link relationship. See [Pagination](/#pagination). | 
 **limit** | **int32** | Specifies the number of objects to return per page. If there are multiple pages of results, the Link header contains a &#x60;next&#x60; link that you need to use as an opaque value (follow it, don&#39;t parse it). See [Pagination](/#pagination).  | [default to 50]
 **q** | **string** | Specifies a filter for the list of Application Users returned based on their profile attributes. The value of &#x60;q&#x60; is matched against the beginning of the following profile attributes: &#x60;userName&#x60;, &#x60;firstName&#x60;, &#x60;lastName&#x60;, and &#x60;email&#x60;. This filter only supports the &#x60;startsWith&#x60; operation that matches the &#x60;q&#x60; string against the beginning of the attribute values. &gt; **Note:** For OIDC apps, user profiles don&#39;t contain the &#x60;firstName&#x60; or &#x60;lastName&#x60; attributes. Therefore, the query only matches against the &#x60;userName&#x60; or &#x60;email&#x60; attributes.  | 
 **expand** | **string** | An optional query parameter to return the corresponding [User](/openapi/okta-management/management/tag/User/) object in the &#x60;_embedded&#x60; property. Valid value: &#x60;user&#x60; | 

### Return type

[**[]AppUser**](AppUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignUserFromApplication

> UnassignUserFromApplication(ctx, appId, userId).SendEmail(sendEmail).Execute()

Unassign an Application User



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
    userId := "00u13okQOVWZJGDOAUVR" // string | ID of an existing Okta user
    sendEmail := true // bool | Sends a deactivation email to the administrator if `true` (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationUsersAPI.UnassignUserFromApplication(context.Background(), appId, userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersAPI.UnassignUserFromApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignUserFromApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendEmail** | **bool** | Sends a deactivation email to the administrator if &#x60;true&#x60; | [default to false]

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


## UpdateApplicationUser

> AppUser UpdateApplicationUser(ctx, appId, userId).AppUser(appUser).Execute()

Update an Application User



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
    userId := "00u13okQOVWZJGDOAUVR" // string | ID of an existing Okta user
    appUser := openapiclient.AppUserUpdateRequest{AppUserCredentialsRequestPayload: openapiclient.NewAppUserCredentialsRequestPayload()} // AppUserUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersAPI.UpdateApplicationUser(context.Background(), appId, userId).AppUser(appUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersAPI.UpdateApplicationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationUser`: AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersAPI.UpdateApplicationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appUser** | [**AppUserUpdateRequest**](AppUserUpdateRequest.md) |  | 

### Return type

[**AppUser**](AppUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

