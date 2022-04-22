# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateEventHook**](EventHookApi.md#ActivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/activate | Activate Event Hook
[**CreateEventHook**](EventHookApi.md#CreateEventHook) | **Post** /api/v1/eventHooks | Create Event Hook
[**DeactivateEventHook**](EventHookApi.md#DeactivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/deactivate | Deactivate Event Hook
[**DeleteEventHook**](EventHookApi.md#DeleteEventHook) | **Delete** /api/v1/eventHooks/{eventHookId} | Delete Event Hook
[**GetEventHook**](EventHookApi.md#GetEventHook) | **Get** /api/v1/eventHooks/{eventHookId} | Get Event Hook
[**ListEventHooks**](EventHookApi.md#ListEventHooks) | **Get** /api/v1/eventHooks | List Event Hooks
[**UpdateEventHook**](EventHookApi.md#UpdateEventHook) | **Put** /api/v1/eventHooks/{eventHookId} | Update Event Hook
[**VerifyEventHook**](EventHookApi.md#VerifyEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/verify | Verify Event Hook

# **ActivateEventHook**
> EventHook ActivateEventHook(ctx, eventHookId)
Activate Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventHook**
> EventHook CreateEventHook(ctx, body)
Create Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventHook**](EventHook.md)|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateEventHook**
> EventHook DeactivateEventHook(ctx, eventHookId)
Deactivate Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventHook**
> DeleteEventHook(ctx, eventHookId)
Delete Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventHook**
> EventHook GetEventHook(ctx, eventHookId)
Get Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEventHooks**
> []EventHook ListEventHooks(ctx, )
List Event Hooks

Success

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventHook**
> EventHook UpdateEventHook(ctx, body, eventHookId)
Update Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventHook**](EventHook.md)|  | 
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEventHook**
> EventHook VerifyEventHook(ctx, eventHookId)
Verify Event Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

