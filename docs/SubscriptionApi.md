# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoleSubscriptionByNotificationType**](SubscriptionApi.md#GetRoleSubscriptionByNotificationType) | **Get** /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType} | Get subscriptions of a Custom Role with a specific notification type
[**GetUserSubscriptionByNotificationType**](SubscriptionApi.md#GetUserSubscriptionByNotificationType) | **Get** /api/v1/users/{userId}/subscriptions/{notificationType} | Get the subscription of a User with a specific notification type
[**ListRoleSubscriptions**](SubscriptionApi.md#ListRoleSubscriptions) | **Get** /api/v1/roles/{roleTypeOrRoleId}/subscriptions | List all subscriptions of a Custom Role
[**ListUserSubscriptions**](SubscriptionApi.md#ListUserSubscriptions) | **Get** /api/v1/users/{userId}/subscriptions | List subscriptions of a User
[**SubscribeRoleSubscriptionByNotificationType**](SubscriptionApi.md#SubscribeRoleSubscriptionByNotificationType) | **Post** /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}/subscribe | Subscribe a Custom Role to a specific notification type
[**SubscribeUserSubscriptionByNotificationType**](SubscriptionApi.md#SubscribeUserSubscriptionByNotificationType) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/subscribe | Subscribe to a specific notification type
[**UnsubscribeRoleSubscriptionByNotificationType**](SubscriptionApi.md#UnsubscribeRoleSubscriptionByNotificationType) | **Post** /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}/unsubscribe | Unsubscribe a Custom Role from a specific notification type
[**UnsubscribeUserSubscriptionByNotificationType**](SubscriptionApi.md#UnsubscribeUserSubscriptionByNotificationType) | **Post** /api/v1/users/{userId}/subscriptions/{notificationType}/unsubscribe | Unsubscribe from a specific notification type

# **GetRoleSubscriptionByNotificationType**
> Subscription GetRoleSubscriptionByNotificationType(ctx, roleTypeOrRoleId, notificationType)
Get subscriptions of a Custom Role with a specific notification type

When roleType Get subscriptions of a Role with a specific notification type. Else when roleId Get subscription of a Custom Role with a specific notification type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleTypeOrRoleId** | **string**|  | 
  **notificationType** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSubscriptionByNotificationType**
> Subscription GetUserSubscriptionByNotificationType(ctx, userId, notificationType)
Get the subscription of a User with a specific notification type

Get the subscriptions of a User with a specific notification type. Only gets subscriptions for current user. An AccessDeniedException message is sent if requests are made from other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **notificationType** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoleSubscriptions**
> []Subscription ListRoleSubscriptions(ctx, roleTypeOrRoleId)
List all subscriptions of a Custom Role

When roleType List all subscriptions of a Role. Else when roleId List subscriptions of a Custom Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleTypeOrRoleId** | **string**|  | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserSubscriptions**
> []Subscription ListUserSubscriptions(ctx, userId)
List subscriptions of a User

List subscriptions of a User. Only lists subscriptions for current user. An AccessDeniedException message is sent if requests are made from other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeRoleSubscriptionByNotificationType**
> SubscribeRoleSubscriptionByNotificationType(ctx, roleTypeOrRoleId, notificationType)
Subscribe a Custom Role to a specific notification type

When roleType Subscribes a Role to a specific notification type. When you change the subscription status of a Role, it overrides the subscription of any individual user of that Role. Else when roleId Subscribes a Custom Role to a specific notification type. When you change the subscription status of a Custom Role, it overrides the subscription of any individual user of that Custom Role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleTypeOrRoleId** | **string**|  | 
  **notificationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeUserSubscriptionByNotificationType**
> SubscribeUserSubscriptionByNotificationType(ctx, userId, notificationType)
Subscribe to a specific notification type

Subscribes a User to a specific notification type. Only the current User can subscribe to a specific notification type. An AccessDeniedException message is sent if requests are made from other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **notificationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeRoleSubscriptionByNotificationType**
> UnsubscribeRoleSubscriptionByNotificationType(ctx, roleTypeOrRoleId, notificationType)
Unsubscribe a Custom Role from a specific notification type

When roleType Unsubscribes a Role from a specific notification type. When you change the subscription status of a Role, it overrides the subscription of any individual user of that Role. Else when roleId Unsubscribes a Custom Role from a specific notification type. When you change the subscription status of a Custom Role, it overrides the subscription of any individual user of that Custom Role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleTypeOrRoleId** | **string**|  | 
  **notificationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeUserSubscriptionByNotificationType**
> UnsubscribeUserSubscriptionByNotificationType(ctx, userId, notificationType)
Unsubscribe from a specific notification type

Unsubscribes a User from a specific notification type. Only the current User can unsubscribe from a specific notification type. An AccessDeniedException message is sent if requests are made from other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **notificationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

