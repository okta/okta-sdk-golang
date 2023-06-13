# \SubscriptionApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRoleSubscriptions**](SubscriptionApi.md#ListRoleSubscriptions) | **Get** /api/v1/roles/{roleTypeOrRoleId}/subscriptions | List all Subscriptions of a Custom Role
[**ListRoleSubscriptionsByNotificationType**](SubscriptionApi.md#ListRoleSubscriptionsByNotificationType) | **Get** /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType} | List all Subscriptions of a Custom Role with a specific notification type
[**ListUserSubscriptions**](SubscriptionApi.md#ListUserSubscriptions) | **Get** /api/v1/users/{userId}/subscriptions | List all Subscriptions
[**ListUserSubscriptionsByNotificationType**](SubscriptionApi.md#ListUserSubscriptionsByNotificationType) | **Get** /api/v1/users/{userId}/subscriptions/{notificationType} | List all Subscriptions by type
[**SubscribeRoleSubscriptionByNotificationType**](SubscriptionApi.md#SubscribeRoleSubscriptionByNotificationType) | **Post** /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}/subscribe | Subscribe a Custom Role to a specific notification type
[**SubscribeUserSubscriptionByNotificationType**](SubscriptionApi.md#SubscribeUserSubscriptionByNotificationType) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/subscribe | Subscribe to a specific notification type
[**UnsubscribeRoleSubscriptionByNotificationType**](SubscriptionApi.md#UnsubscribeRoleSubscriptionByNotificationType) | **Post** /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}/unsubscribe | Unsubscribe a Custom Role from a specific notification type
[**UnsubscribeUserSubscriptionByNotificationType**](SubscriptionApi.md#UnsubscribeUserSubscriptionByNotificationType) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/unsubscribe | Unsubscribe from a specific notification type



## ListRoleSubscriptions

> []Subscription ListRoleSubscriptions(ctx, roleTypeOrRoleId).Execute()

List all Subscriptions of a Custom Role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTypeOrRoleId := "roleTypeOrRoleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListRoleSubscriptions(context.Background(), roleTypeOrRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListRoleSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListRoleSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTypeOrRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleSubscriptionsRequest struct via the builder pattern


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


## ListRoleSubscriptionsByNotificationType

> Subscription ListRoleSubscriptionsByNotificationType(ctx, roleTypeOrRoleId, notificationType).Execute()

List all Subscriptions of a Custom Role with a specific notification type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTypeOrRoleId := "roleTypeOrRoleId_example" // string | 
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListRoleSubscriptionsByNotificationType(context.Background(), roleTypeOrRoleId, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListRoleSubscriptionsByNotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleSubscriptionsByNotificationType`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListRoleSubscriptionsByNotificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTypeOrRoleId** | **string** |  | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleSubscriptionsByNotificationTypeRequest struct via the builder pattern


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


## ListUserSubscriptions

> []Subscription ListUserSubscriptions(ctx, userId).Execute()

List all Subscriptions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListUserSubscriptions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListUserSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListUserSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserSubscriptionsRequest struct via the builder pattern


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


## ListUserSubscriptionsByNotificationType

> Subscription ListUserSubscriptionsByNotificationType(ctx, userId, notificationType).Execute()

List all Subscriptions by type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListUserSubscriptionsByNotificationType(context.Background(), userId, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListUserSubscriptionsByNotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserSubscriptionsByNotificationType`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListUserSubscriptionsByNotificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserSubscriptionsByNotificationTypeRequest struct via the builder pattern


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


## SubscribeRoleSubscriptionByNotificationType

> SubscribeRoleSubscriptionByNotificationType(ctx, roleTypeOrRoleId, notificationType).Execute()

Subscribe a Custom Role to a specific notification type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTypeOrRoleId := "roleTypeOrRoleId_example" // string | 
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.SubscribeRoleSubscriptionByNotificationType(context.Background(), roleTypeOrRoleId, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscribeRoleSubscriptionByNotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTypeOrRoleId** | **string** |  | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRoleSubscriptionByNotificationTypeRequest struct via the builder pattern


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


## SubscribeUserSubscriptionByNotificationType

> SubscribeUserSubscriptionByNotificationType(ctx, userId, notificationType).Execute()

Subscribe to a specific notification type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.SubscribeUserSubscriptionByNotificationType(context.Background(), userId, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscribeUserSubscriptionByNotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeUserSubscriptionByNotificationTypeRequest struct via the builder pattern


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


## UnsubscribeRoleSubscriptionByNotificationType

> UnsubscribeRoleSubscriptionByNotificationType(ctx, roleTypeOrRoleId, notificationType).Execute()

Unsubscribe a Custom Role from a specific notification type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleTypeOrRoleId := "roleTypeOrRoleId_example" // string | 
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.UnsubscribeRoleSubscriptionByNotificationType(context.Background(), roleTypeOrRoleId, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.UnsubscribeRoleSubscriptionByNotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleTypeOrRoleId** | **string** |  | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeRoleSubscriptionByNotificationTypeRequest struct via the builder pattern


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


## UnsubscribeUserSubscriptionByNotificationType

> UnsubscribeUserSubscriptionByNotificationType(ctx, userId, notificationType).Execute()

Unsubscribe from a specific notification type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 
    notificationType := "notificationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.UnsubscribeUserSubscriptionByNotificationType(context.Background(), userId, notificationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.UnsubscribeUserSubscriptionByNotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**notificationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeUserSubscriptionByNotificationTypeRequest struct via the builder pattern


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

