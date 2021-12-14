# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsTemplate**](TemplateApi.md#CreateSmsTemplate) | **Post** /api/v1/templates/sms | Add SMS Template
[**DeleteSmsTemplate**](TemplateApi.md#DeleteSmsTemplate) | **Delete** /api/v1/templates/sms/{templateId} | Remove SMS Template
[**GetSmsTemplate**](TemplateApi.md#GetSmsTemplate) | **Get** /api/v1/templates/sms/{templateId} | Get SMS Template
[**ListSmsTemplates**](TemplateApi.md#ListSmsTemplates) | **Get** /api/v1/templates/sms | List SMS Templates
[**PartialUpdateSmsTemplate**](TemplateApi.md#PartialUpdateSmsTemplate) | **Post** /api/v1/templates/sms/{templateId} | Partial SMS Template Update
[**UpdateSmsTemplate**](TemplateApi.md#UpdateSmsTemplate) | **Put** /api/v1/templates/sms/{templateId} | Update SMS Template

# **CreateSmsTemplate**
> SmsTemplate CreateSmsTemplate(ctx, body)
Add SMS Template

Adds a new custom SMS template to your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmsTemplate**](SmsTemplate.md)|  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmsTemplate**
> DeleteSmsTemplate(ctx, templateId)
Remove SMS Template

Removes an SMS template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmsTemplate**
> SmsTemplate GetSmsTemplate(ctx, templateId)
Get SMS Template

Fetches a specific template by `id`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**|  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSmsTemplates**
> []SmsTemplate ListSmsTemplates(ctx, optional)
List SMS Templates

Enumerates custom SMS templates in your organization. A subset of templates can be returned that match a template type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TemplateApiListSmsTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplateApiListSmsTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateType** | [**optional.Interface of SmsTemplateType**](.md)|  | 

### Return type

[**[]SmsTemplate**](SmsTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateSmsTemplate**
> SmsTemplate PartialUpdateSmsTemplate(ctx, body, templateId)
Partial SMS Template Update

Updates only some of the SMS template properties:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmsTemplate**](SmsTemplate.md)|  | 
  **templateId** | **string**|  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmsTemplate**
> SmsTemplate UpdateSmsTemplate(ctx, body, templateId)
Update SMS Template

Updates the SMS template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmsTemplate**](SmsTemplate.md)|  | 
  **templateId** | **string**|  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

