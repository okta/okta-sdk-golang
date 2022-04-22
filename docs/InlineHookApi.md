# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateInlineHook**](InlineHookApi.md#ActivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/activate | Activate Inline Hook
[**CreateInlineHook**](InlineHookApi.md#CreateInlineHook) | **Post** /api/v1/inlineHooks | Create Inline Hook
[**DeactivateInlineHook**](InlineHookApi.md#DeactivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/deactivate | Deactivate Inline Hook
[**DeleteInlineHook**](InlineHookApi.md#DeleteInlineHook) | **Delete** /api/v1/inlineHooks/{inlineHookId} | Delete Inline Hook
[**ExecuteInlineHook**](InlineHookApi.md#ExecuteInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/execute | Execute Inline Hook
[**GetInlineHook**](InlineHookApi.md#GetInlineHook) | **Get** /api/v1/inlineHooks/{inlineHookId} | Get Inline Hook
[**ListInlineHooks**](InlineHookApi.md#ListInlineHooks) | **Get** /api/v1/inlineHooks | List Inline Hooks
[**UpdateInlineHook**](InlineHookApi.md#UpdateInlineHook) | **Put** /api/v1/inlineHooks/{inlineHookId} | Update Inline Hook

# **ActivateInlineHook**
> InlineHook ActivateInlineHook(ctx, inlineHookId)
Activate Inline Hook

Activates the Inline Hook matching the provided id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inlineHookId** | **string**|  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInlineHook**
> InlineHook CreateInlineHook(ctx, body)
Create Inline Hook

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InlineHook**](InlineHook.md)|  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateInlineHook**
> InlineHook DeactivateInlineHook(ctx, inlineHookId)
Deactivate Inline Hook

Deactivates the Inline Hook matching the provided id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inlineHookId** | **string**|  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInlineHook**
> DeleteInlineHook(ctx, inlineHookId)
Delete Inline Hook

Deletes the Inline Hook matching the provided id. Once deleted, the Inline Hook is unrecoverable. As a safety precaution, only Inline Hooks with a status of INACTIVE are eligible for deletion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inlineHookId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteInlineHook**
> InlineHookResponse ExecuteInlineHook(ctx, body, inlineHookId)
Execute Inline Hook

Executes the Inline Hook matching the provided inlineHookId using the request body as the input. This will send the provided data through the Channel and return a response if it matches the correct data contract. This execution endpoint should only be used for testing purposes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InlineHookPayload**](InlineHookPayload.md)|  | 
  **inlineHookId** | **string**|  | 

### Return type

[**InlineHookResponse**](InlineHookResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInlineHook**
> InlineHook GetInlineHook(ctx, inlineHookId)
Get Inline Hook

Gets an inline hook by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inlineHookId** | **string**|  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInlineHooks**
> []InlineHook ListInlineHooks(ctx, optional)
List Inline Hooks

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InlineHookApiListInlineHooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InlineHookApiListInlineHooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 

### Return type

[**[]InlineHook**](InlineHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInlineHook**
> InlineHook UpdateInlineHook(ctx, body, inlineHookId)
Update Inline Hook

Updates an inline hook by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InlineHook**](InlineHook.md)|  | 
  **inlineHookId** | **string**|  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

