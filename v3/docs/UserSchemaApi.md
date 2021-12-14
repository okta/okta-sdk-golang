# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationUserSchema**](UserSchemaApi.md#GetApplicationUserSchema) | **Get** /api/v1/meta/schemas/apps/{appInstanceId}/default | Fetches the Schema for an App User
[**GetUserSchema**](UserSchemaApi.md#GetUserSchema) | **Get** /api/v1/meta/schemas/user/{schemaId} | Fetches the schema for a Schema Id.
[**UpdateApplicationUserProfile**](UserSchemaApi.md#UpdateApplicationUserProfile) | **Post** /api/v1/meta/schemas/apps/{appInstanceId}/default | Partial updates on the User Profile properties of the Application User Schema.
[**UpdateUserProfile**](UserSchemaApi.md#UpdateUserProfile) | **Post** /api/v1/meta/schemas/user/{schemaId} | 

# **GetApplicationUserSchema**
> UserSchema GetApplicationUserSchema(ctx, appInstanceId)
Fetches the Schema for an App User

Fetches the Schema for an App User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appInstanceId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSchema**
> UserSchema GetUserSchema(ctx, schemaId)
Fetches the schema for a Schema Id.

Fetches the schema for a Schema Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schemaId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationUserProfile**
> UserSchema UpdateApplicationUserProfile(ctx, appInstanceId, optional)
Partial updates on the User Profile properties of the Application User Schema.

Partial updates on the User Profile properties of the Application User Schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appInstanceId** | **string**|  | 
 **optional** | ***UserSchemaApiUpdateApplicationUserProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserSchemaApiUpdateApplicationUserProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserProfile**
> UserSchema UpdateUserProfile(ctx, body, schemaId)


Partial updates on the User Profile properties of the user schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserSchema**](UserSchema.md)|  | 
  **schemaId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

