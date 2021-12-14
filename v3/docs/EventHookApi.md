# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateEventHook**](EventHookApi.md#ActivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/activate | 
[**CreateEventHook**](EventHookApi.md#CreateEventHook) | **Post** /api/v1/eventHooks | 
[**DeactivateEventHook**](EventHookApi.md#DeactivateEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/deactivate | 
[**DeleteEventHook**](EventHookApi.md#DeleteEventHook) | **Delete** /api/v1/eventHooks/{eventHookId} | 
[**GetEventHook**](EventHookApi.md#GetEventHook) | **Get** /api/v1/eventHooks/{eventHookId} | 
[**ListEventHooks**](EventHookApi.md#ListEventHooks) | **Get** /api/v1/eventHooks | 
[**UpdateEventHook**](EventHookApi.md#UpdateEventHook) | **Put** /api/v1/eventHooks/{eventHookId} | 
[**VerifyEventHook**](EventHookApi.md#VerifyEventHook) | **Post** /api/v1/eventHooks/{eventHookId}/lifecycle/verify | 

# **ActivateEventHook**
> EventHook ActivateEventHook(ctx, eventHookId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventHook**
> EventHook CreateEventHook(ctx, body)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventHook**](EventHook.md)|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateEventHook**
> EventHook DeactivateEventHook(ctx, eventHookId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventHook**
> DeleteEventHook(ctx, eventHookId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventHook**
> EventHook GetEventHook(ctx, eventHookId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEventHooks**
> []EventHook ListEventHooks(ctx, )


Success

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventHook**
> EventHook UpdateEventHook(ctx, body, eventHookId)


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

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEventHook**
> EventHook VerifyEventHook(ctx, eventHookId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventHookId** | **string**|  | 

### Return type

[**EventHook**](EventHook.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

