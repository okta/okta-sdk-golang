# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateOrigin**](TrustedOriginApi.md#ActivateOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/activate | Activate Trusted Origin
[**CreateOrigin**](TrustedOriginApi.md#CreateOrigin) | **Post** /api/v1/trustedOrigins | Create Trusted Origin
[**DeactivateOrigin**](TrustedOriginApi.md#DeactivateOrigin) | **Post** /api/v1/trustedOrigins/{trustedOriginId}/lifecycle/deactivate | Deactivate Trusted Origin
[**DeleteOrigin**](TrustedOriginApi.md#DeleteOrigin) | **Delete** /api/v1/trustedOrigins/{trustedOriginId} | Delete Trusted Origin
[**GetOrigin**](TrustedOriginApi.md#GetOrigin) | **Get** /api/v1/trustedOrigins/{trustedOriginId} | Get Trusted Origin
[**ListOrigins**](TrustedOriginApi.md#ListOrigins) | **Get** /api/v1/trustedOrigins | List Trusted Origins
[**UpdateOrigin**](TrustedOriginApi.md#UpdateOrigin) | **Put** /api/v1/trustedOrigins/{trustedOriginId} | Update Trusted Origin

# **ActivateOrigin**
> TrustedOrigin ActivateOrigin(ctx, trustedOriginId)
Activate Trusted Origin

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trustedOriginId** | **string**|  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrigin**
> TrustedOrigin CreateOrigin(ctx, body)
Create Trusted Origin

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrustedOrigin**](TrustedOrigin.md)|  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateOrigin**
> TrustedOrigin DeactivateOrigin(ctx, trustedOriginId)
Deactivate Trusted Origin

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trustedOriginId** | **string**|  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrigin**
> DeleteOrigin(ctx, trustedOriginId)
Delete Trusted Origin

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trustedOriginId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrigin**
> TrustedOrigin GetOrigin(ctx, trustedOriginId)
Get Trusted Origin

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trustedOriginId** | **string**|  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrigins**
> []TrustedOrigin ListOrigins(ctx, optional)
List Trusted Origins

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TrustedOriginApiListOriginsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrustedOriginApiListOriginsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  | 
 **filter** | **optional.String**|  | 
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to -1]

### Return type

[**[]TrustedOrigin**](TrustedOrigin.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrigin**
> TrustedOrigin UpdateOrigin(ctx, body, trustedOriginId)
Update Trusted Origin

Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrustedOrigin**](TrustedOrigin.md)|  | 
  **trustedOriginId** | **string**|  | 

### Return type

[**TrustedOrigin**](TrustedOrigin.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

