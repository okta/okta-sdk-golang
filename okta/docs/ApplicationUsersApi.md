# \ApplicationUsersApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignUserToApplication**](ApplicationUsersApi.md#AssignUserToApplication) | **Post** /api/v1/apps/{appId}/users | Assign a User
[**GetApplicationUser**](ApplicationUsersApi.md#GetApplicationUser) | **Get** /api/v1/apps/{appId}/users/{userId} | Retrieve an assigned User
[**ListApplicationUsers**](ApplicationUsersApi.md#ListApplicationUsers) | **Get** /api/v1/apps/{appId}/users | List all assigned Users
[**UnassignUserFromApplication**](ApplicationUsersApi.md#UnassignUserFromApplication) | **Delete** /api/v1/apps/{appId}/users/{userId} | Unassign an App User
[**UpdateApplicationUser**](ApplicationUsersApi.md#UpdateApplicationUser) | **Post** /api/v1/apps/{appId}/users/{userId} | Update an App Profile for an assigned User



## AssignUserToApplication

> AppUser AssignUserToApplication(ctx, appId).AppUser(appUser).Execute()

Assign a User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    appUser := *openapiclient.NewAppUser(time.Now(), time.Now(), "USER", openapiclient.AppUserStatus("ACTIVE"), time.Now(), *openapiclient.NewLinksAppAndUser()) // AppUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersApi.AssignUserToApplication(context.Background(), appId).AppUser(appUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersApi.AssignUserToApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignUserToApplication`: AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersApi.AssignUserToApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUserToApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUser** | [**AppUser**](AppUser.md) |  | 

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

Retrieve an assigned User



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
    userId := "userId_example" // string | 
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersApi.GetApplicationUser(context.Background(), appId, userId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersApi.GetApplicationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationUser`: AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersApi.GetApplicationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 

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

> []AppUser ListApplicationUsers(ctx, appId).Q(q).QueryScope(queryScope).After(after).Limit(limit).Filter(filter).Expand(expand).Execute()

List all assigned Users



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
    queryScope := "queryScope_example" // string |  (optional)
    after := "after_example" // string | specifies the pagination cursor for the next page of assignments (optional)
    limit := int32(56) // int32 | specifies the number of results for a page (optional) (default to -1)
    filter := "filter_example" // string |  (optional)
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersApi.ListApplicationUsers(context.Background(), appId).Q(q).QueryScope(queryScope).After(after).Limit(limit).Filter(filter).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersApi.ListApplicationUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationUsers`: []AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersApi.ListApplicationUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | 
 **queryScope** | **string** |  | 
 **after** | **string** | specifies the pagination cursor for the next page of assignments | 
 **limit** | **int32** | specifies the number of results for a page | [default to -1]
 **filter** | **string** |  | 
 **expand** | **string** |  | 

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

Unassign an App User



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
    userId := "userId_example" // string | 
    sendEmail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationUsersApi.UnassignUserFromApplication(context.Background(), appId, userId).SendEmail(sendEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersApi.UnassignUserFromApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignUserFromApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendEmail** | **bool** |  | [default to false]

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

Update an App Profile for an assigned User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    userId := "userId_example" // string | 
    appUser := *openapiclient.NewAppUser(time.Now(), time.Now(), "USER", openapiclient.AppUserStatus("ACTIVE"), time.Now(), *openapiclient.NewLinksAppAndUser()) // AppUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationUsersApi.UpdateApplicationUser(context.Background(), appId, userId).AppUser(appUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationUsersApi.UpdateApplicationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationUser`: AppUser
    fmt.Fprintf(os.Stdout, "Response from `ApplicationUsersApi.UpdateApplicationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appUser** | [**AppUser**](AppUser.md) |  | 

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

