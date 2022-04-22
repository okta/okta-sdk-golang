# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationUserSchema**](SchemaApi.md#GetApplicationUserSchema) | **Get** /api/v1/meta/schemas/apps/{appInstanceId}/default | Fetches the Schema for an App User
[**GetGroupSchema**](SchemaApi.md#GetGroupSchema) | **Get** /api/v1/meta/schemas/group/default | Fetches the group schema
[**GetUserSchema**](SchemaApi.md#GetUserSchema) | **Get** /api/v1/meta/schemas/user/{schemaId} | Fetches the schema for a Schema Id.
[**UpdateApplicationUserProfile**](SchemaApi.md#UpdateApplicationUserProfile) | **Post** /api/v1/meta/schemas/apps/{appInstanceId}/default | Partial updates on the User Profile properties of the Application User Schema.
[**UpdateGroupSchema**](SchemaApi.md#UpdateGroupSchema) | **Post** /api/v1/meta/schemas/group/default | Updates, adds ore removes one or more custom Group Profile properties in the schema
[**UpdateUserProfile**](SchemaApi.md#UpdateUserProfile) | **Post** /api/v1/meta/schemas/user/{schemaId} | Update User Profile

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

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupSchema**
> GroupSchema GetGroupSchema(ctx, )
Fetches the group schema

Fetches the group schema

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

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

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

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
 **optional** | ***SchemaApiUpdateApplicationUserProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemaApiUpdateApplicationUserProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
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
 **optional** | ***SchemaApiUpdateGroupSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemaApiUpdateGroupSchemaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupSchema**](GroupSchema.md)|  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserProfile**
> UserSchema UpdateUserProfile(ctx, body, schemaId)
Update User Profile

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

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

