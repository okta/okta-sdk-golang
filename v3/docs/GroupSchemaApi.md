# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupSchema**](GroupSchemaApi.md#GetGroupSchema) | **Get** /api/v1/meta/schemas/group/default | Fetches the group schema
[**UpdateGroupSchema**](GroupSchemaApi.md#UpdateGroupSchema) | **Post** /api/v1/meta/schemas/group/default | Updates, adds ore removes one or more custom Group Profile properties in the schema

# **GetGroupSchema**
> GroupSchema GetGroupSchema(ctx, )
Fetches the group schema

Fetches the group schema

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupSchema**
> GroupSchema UpdateGroupSchema(ctx, optional)
Updates, adds ore removes one or more custom Group Profile properties in the schema

Updates, adds ore removes one or more custom Group Profile properties in the schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupSchemaApiUpdateGroupSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupSchemaApiUpdateGroupSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupSchema**](GroupSchema.md)|  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

