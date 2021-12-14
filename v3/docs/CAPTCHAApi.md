# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCaptchaInstance**](CAPTCHAApi.md#CreateCaptchaInstance) | **Post** /api/v1/captchas | Create new CAPTCHA instance
[**DeleteCaptchaInstance**](CAPTCHAApi.md#DeleteCaptchaInstance) | **Delete** /api/v1/captchas/{captchaId} | Delete CAPTCHA Instance
[**GetCaptchaInstance**](CAPTCHAApi.md#GetCaptchaInstance) | **Get** /api/v1/captchas/{captchaId} | Get CAPTCHA Instance
[**ListCaptchaInstances**](CAPTCHAApi.md#ListCaptchaInstances) | **Get** /api/v1/captchas | List CAPTCHA instances
[**PartialUpdateCaptchaInstance**](CAPTCHAApi.md#PartialUpdateCaptchaInstance) | **Post** /api/v1/captchas/{captchaId} | Partial Update CAPTCHA instance
[**UpdateCaptchaInstance**](CAPTCHAApi.md#UpdateCaptchaInstance) | **Put** /api/v1/captchas/{captchaId} | Update CAPTCHA instance

# **CreateCaptchaInstance**
> CaptchaInstance CreateCaptchaInstance(ctx, optional)
Create new CAPTCHA instance

Adds a new CAPTCHA instance to your organization. In current release, we only allow one CAPTCHA instance per org.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CAPTCHAApiCreateCaptchaInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CAPTCHAApiCreateCaptchaInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CaptchaInstance**](CaptchaInstance.md)|  | 

### Return type

[**CaptchaInstance**](CAPTCHAInstance.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCaptchaInstance**
> DeleteCaptchaInstance(ctx, captchaId)
Delete CAPTCHA Instance

Delete a CAPTCHA instance by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **captchaId** | **string**| id of the CAPTCHA | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCaptchaInstance**
> CaptchaInstance GetCaptchaInstance(ctx, captchaId)
Get CAPTCHA Instance

Fetches a CAPTCHA instance by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **captchaId** | **string**| id of the CAPTCHA | 

### Return type

[**CaptchaInstance**](CAPTCHAInstance.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCaptchaInstances**
> []CaptchaInstance ListCaptchaInstances(ctx, )
List CAPTCHA instances

Enumerates CAPTCHA instances in your organization with pagination. A subset of CAPTCHA instances can be returned that match a supported filter expression or query.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CaptchaInstance**](CAPTCHAInstance.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateCaptchaInstance**
> CaptchaInstance PartialUpdateCaptchaInstance(ctx, body, captchaId)
Partial Update CAPTCHA instance

Partially update a CAPTCHA instance by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CaptchaInstance**](CaptchaInstance.md)|  | 
  **captchaId** | **string**| id of the CAPTCHA | 

### Return type

[**CaptchaInstance**](CAPTCHAInstance.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCaptchaInstance**
> CaptchaInstance UpdateCaptchaInstance(ctx, body, captchaId)
Update CAPTCHA instance

Update a CAPTCHA instance by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CaptchaInstance**](CaptchaInstance.md)|  | 
  **captchaId** | **string**| id of the CAPTCHA | 

### Return type

[**CaptchaInstance**](CAPTCHAInstance.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

