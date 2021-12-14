# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLinkedObjectDefinition**](LinkedObjectApi.md#AddLinkedObjectDefinition) | **Post** /api/v1/meta/schemas/user/linkedObjects | 
[**DeleteLinkedObjectDefinition**](LinkedObjectApi.md#DeleteLinkedObjectDefinition) | **Delete** /api/v1/meta/schemas/user/linkedObjects/{linkedObjectName} | 
[**GetLinkedObjectDefinition**](LinkedObjectApi.md#GetLinkedObjectDefinition) | **Get** /api/v1/meta/schemas/user/linkedObjects/{linkedObjectName} | 
[**ListLinkedObjectDefinitions**](LinkedObjectApi.md#ListLinkedObjectDefinitions) | **Get** /api/v1/meta/schemas/user/linkedObjects | 

# **AddLinkedObjectDefinition**
> LinkedObject AddLinkedObjectDefinition(ctx, body)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LinkedObject**](LinkedObject.md)|  | 

### Return type

[**LinkedObject**](LinkedObject.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLinkedObjectDefinition**
> DeleteLinkedObjectDefinition(ctx, linkedObjectName)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkedObjectName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedObjectDefinition**
> LinkedObject GetLinkedObjectDefinition(ctx, linkedObjectName)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkedObjectName** | **string**|  | 

### Return type

[**LinkedObject**](LinkedObject.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLinkedObjectDefinitions**
> []LinkedObject ListLinkedObjectDefinitions(ctx, )


Success

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LinkedObject**](LinkedObject.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

