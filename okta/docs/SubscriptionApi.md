# \SubscriptionApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptionsNotificationTypeRole**](SubscriptionApi.md#GetSubscriptionsNotificationTypeRole) | **Get** /api/v1/roles/{roleRef}/subscriptions/{notificationType} | Retrieve a Subscription for a Role
[**GetSubscriptionsNotificationTypeUser**](SubscriptionApi.md#GetSubscriptionsNotificationTypeUser) | **Get** /api/v1/users/{userId}/subscriptions/{notificationType} | Retrieve a Subscription for a User
[**ListSubscriptionsRole**](SubscriptionApi.md#ListSubscriptionsRole) | **Get** /api/v1/roles/{roleRef}/subscriptions | List all Subscriptions for a Role
[**ListSubscriptionsUser**](SubscriptionApi.md#ListSubscriptionsUser) | **Get** /api/v1/users/{userId}/subscriptions | List all Subscriptions for a User
[**SubscribeByNotificationTypeRole**](SubscriptionApi.md#SubscribeByNotificationTypeRole) | **Post** /api/v1/roles/{roleRef}/subscriptions/{notificationType}/subscribe | Subscribe a Role to a Specific Notification Type
[**SubscribeByNotificationTypeUser**](SubscriptionApi.md#SubscribeByNotificationTypeUser) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/subscribe | Subscribe a User to a Specific Notification Type
[**UnsubscribeByNotificationTypeRole**](SubscriptionApi.md#UnsubscribeByNotificationTypeRole) | **Post** /api/v1/roles/{roleRef}/subscriptions/{notificationType}/unsubscribe | Unsubscribe a Role from a Specific Notification Type
[**UnsubscribeByNotificationTypeUser**](SubscriptionApi.md#UnsubscribeByNotificationTypeUser) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/unsubscribe | Unsubscribe a User from a Specific Notification Type



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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{RoleType: penapiclient.RoleType("API_ACCESS_MANAGEMENT_ADMIN")} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types).
    notificationType := openapiclient.NotificationType("AD_AGENT") // NotificationType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetSubscriptionsNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetSubscriptionsNotificationTypeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionsNotificationTypeRole`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetSubscriptionsNotificationTypeRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types). | 
**notificationType** | [**NotificationType**](.md) |  | 

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
    notificationType := openapiclient.NotificationType("AD_AGENT") // NotificationType | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetSubscriptionsNotificationTypeUser(context.Background(), notificationType, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetSubscriptionsNotificationTypeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionsNotificationTypeUser`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetSubscriptionsNotificationTypeUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | [**NotificationType**](.md) |  | 
**userId** | **string** |  | 

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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{RoleType: penapiclient.RoleType("API_ACCESS_MANAGEMENT_ADMIN")} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListSubscriptionsRole(context.Background(), roleRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListSubscriptionsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionsRole`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListSubscriptionsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types). | 

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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListSubscriptionsUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListSubscriptionsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionsUser`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListSubscriptionsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{RoleType: penapiclient.RoleType("API_ACCESS_MANAGEMENT_ADMIN")} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types).
    notificationType := openapiclient.NotificationType("AD_AGENT") // NotificationType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.SubscribeByNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscribeByNotificationTypeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types). | 
**notificationType** | [**NotificationType**](.md) |  | 

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
    notificationType := openapiclient.NotificationType("AD_AGENT") // NotificationType | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.SubscribeByNotificationTypeUser(context.Background(), notificationType, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscribeByNotificationTypeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | [**NotificationType**](.md) |  | 
**userId** | **string** |  | 

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
    roleRef := openapiclient.listSubscriptionsRole_roleRef_parameter{RoleType: penapiclient.RoleType("API_ACCESS_MANAGEMENT_ADMIN")} // ListSubscriptionsRoleRoleRefParameter | A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types).
    notificationType := openapiclient.NotificationType("AD_AGENT") // NotificationType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.UnsubscribeByNotificationTypeRole(context.Background(), roleRef, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.UnsubscribeByNotificationTypeRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleRef** | [**ListSubscriptionsRoleRoleRefParameter**](.md) | A reference to an existing role. Standard roles require a &#x60;roleType&#x60;, while Custom Roles require a &#x60;roleId&#x60;. See [Standard Role Types](https://developer.okta.com/docs/concepts/role-assignment/#standard-role-types). | 
**notificationType** | [**NotificationType**](.md) |  | 

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
    notificationType := openapiclient.NotificationType("AD_AGENT") // NotificationType | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.UnsubscribeByNotificationTypeUser(context.Background(), notificationType, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.UnsubscribeByNotificationTypeUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationType** | [**NotificationType**](.md) |  | 
**userId** | **string** |  | 

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

