# \SubscriptionAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptionsNotificationTypeRole**](SubscriptionAPI.md#GetSubscriptionsNotificationTypeRole) | **Get** /api/v1/roles/{roleRef}/subscriptions/{notificationType} | Retrieve a Subscription for a Role
[**GetSubscriptionsNotificationTypeUser**](SubscriptionAPI.md#GetSubscriptionsNotificationTypeUser) | **Get** /api/v1/users/{userId}/subscriptions/{notificationType} | Retrieve a Subscription for a User
[**ListSubscriptionsRole**](SubscriptionAPI.md#ListSubscriptionsRole) | **Get** /api/v1/roles/{roleRef}/subscriptions | List all Subscriptions for a Role
[**ListSubscriptionsUser**](SubscriptionAPI.md#ListSubscriptionsUser) | **Get** /api/v1/users/{userId}/subscriptions | List all Subscriptions for a User
[**SubscribeByNotificationTypeRole**](SubscriptionAPI.md#SubscribeByNotificationTypeRole) | **Post** /api/v1/roles/{roleRef}/subscriptions/{notificationType}/subscribe | Subscribe a Role to a Specific Notification Type
[**SubscribeByNotificationTypeUser**](SubscriptionAPI.md#SubscribeByNotificationTypeUser) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/subscribe | Subscribe a User to a Specific Notification Type
[**UnsubscribeByNotificationTypeRole**](SubscriptionAPI.md#UnsubscribeByNotificationTypeRole) | **Post** /api/v1/roles/{roleRef}/subscriptions/{notificationType}/unsubscribe | Unsubscribe a Role from a Specific Notification Type
[**UnsubscribeByNotificationTypeUser**](SubscriptionAPI.md#UnsubscribeByNotificationTypeUser) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/unsubscribe | Unsubscribe a User from a Specific Notification Type



## GetSubscriptionsNotificationTypeRole

> Subscription GetSubscriptionsNotificationTypeRole(ctx, roleRef, notificationType).Execute()

Retrieve a Subscription for a Role



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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{String: new(string)} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.GetSubscriptionsNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.GetSubscriptionsNotificationTypeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionsNotificationTypeRole`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.GetSubscriptionsNotificationTypeRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles). | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsNotificationTypeRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Subscription**](Subscription.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionsNotificationTypeUser

> Subscription GetSubscriptionsNotificationTypeUser(ctx, notificationType, userId).Execute()

Retrieve a Subscription for a User



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
    notificationType := "notificationType_example" // string | 
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.GetSubscriptionsNotificationTypeUser(context.Background(), notificationType, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.GetSubscriptionsNotificationTypeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionsNotificationTypeUser`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.GetSubscriptionsNotificationTypeUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | **string** |  | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsNotificationTypeUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Subscription**](Subscription.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionsRole

> []Subscription ListSubscriptionsRole(ctx, roleRef).Execute()

List all Subscriptions for a Role



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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{String: new(string)} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.ListSubscriptionsRole(context.Background(), roleRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ListSubscriptionsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionsRole`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.ListSubscriptionsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionsUser

> []Subscription ListSubscriptionsUser(ctx, userId).Execute()

List all Subscriptions for a User



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionAPI.ListSubscriptionsUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.ListSubscriptionsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionsUser`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.ListSubscriptionsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeByNotificationTypeRole

> SubscribeByNotificationTypeRole(ctx, roleRef, notificationType).Execute()

Subscribe a Role to a Specific Notification Type



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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{String: new(string)} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.SubscribeByNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscribeByNotificationTypeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles). | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeByNotificationTypeRoleRequest struct via the builder pattern


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


## SubscribeByNotificationTypeUser

> SubscribeByNotificationTypeUser(ctx, notificationType, userId).Execute()

Subscribe a User to a Specific Notification Type



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
    notificationType := "notificationType_example" // string | 
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.SubscribeByNotificationTypeUser(context.Background(), notificationType, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscribeByNotificationTypeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | **string** |  | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeByNotificationTypeUserRequest struct via the builder pattern


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


## UnsubscribeByNotificationTypeRole

> UnsubscribeByNotificationTypeRole(ctx, roleRef, notificationType).Execute()

Unsubscribe a Role from a Specific Notification Type



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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{String: new(string)} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.UnsubscribeByNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.UnsubscribeByNotificationTypeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles). | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeByNotificationTypeRoleRequest struct via the builder pattern


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


## UnsubscribeByNotificationTypeUser

> UnsubscribeByNotificationTypeUser(ctx, notificationType, userId).Execute()

Unsubscribe a User from a Specific Notification Type



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
    notificationType := "notificationType_example" // string | 
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionAPI.UnsubscribeByNotificationTypeUser(context.Background(), notificationType, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.UnsubscribeByNotificationTypeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | **string** |  | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeByNotificationTypeUserRequest struct via the builder pattern


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

