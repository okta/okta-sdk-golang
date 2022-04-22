# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfileMapping**](ProfileMappingApi.md#GetProfileMapping) | **Get** /api/v1/mappings/{mappingId} | Get Profile Mapping
[**ListProfileMappings**](ProfileMappingApi.md#ListProfileMappings) | **Get** /api/v1/mappings | List Profile Mappings
[**UpdateProfileMapping**](ProfileMappingApi.md#UpdateProfileMapping) | **Post** /api/v1/mappings/{mappingId} | Update Profile Mapping

# **GetProfileMapping**
> ProfileMapping GetProfileMapping(ctx, mappingId)
Get Profile Mapping

Fetches a single Profile Mapping referenced by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mappingId** | **string**|  | 

### Return type

[**ProfileMapping**](ProfileMapping.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProfileMappings**
> []ProfileMapping ListProfileMappings(ctx, optional)
List Profile Mappings

Enumerates Profile Mappings in your organization with pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProfileMappingApiListProfileMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProfileMappingApiListProfileMappingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to -1]
 **sourceId** | **optional.String**|  | 
 **targetId** | **optional.String**|  | 

### Return type

[**[]ProfileMapping**](ProfileMapping.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProfileMapping**
> ProfileMapping UpdateProfileMapping(ctx, body, mappingId)
Update Profile Mapping

Updates an existing Profile Mapping by adding, updating, or removing one or many Property Mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProfileMapping**](ProfileMapping.md)|  | 
  **mappingId** | **string**|  | 

### Return type

[**ProfileMapping**](ProfileMapping.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

