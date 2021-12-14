# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserType**](UserTypeApi.md#CreateUserType) | **Post** /api/v1/meta/types/user | 
[**DeleteUserType**](UserTypeApi.md#DeleteUserType) | **Delete** /api/v1/meta/types/user/{typeId} | 
[**GetUserType**](UserTypeApi.md#GetUserType) | **Get** /api/v1/meta/types/user/{typeId} | 
[**ListUserTypes**](UserTypeApi.md#ListUserTypes) | **Get** /api/v1/meta/types/user | 
[**ReplaceUserType**](UserTypeApi.md#ReplaceUserType) | **Put** /api/v1/meta/types/user/{typeId} | 
[**UpdateUserType**](UserTypeApi.md#UpdateUserType) | **Post** /api/v1/meta/types/user/{typeId} | 

# **CreateUserType**
> UserType CreateUserType(ctx, body)


Creates a new User Type. A default User Type is automatically created along with your org, and you may add another 9 User Types for a maximum of 10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserType**](UserType.md)|  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserType**
> DeleteUserType(ctx, typeId)


Deletes a User Type permanently. This operation is not permitted for the default type, nor for any User Type that has existing users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserType**
> UserType GetUserType(ctx, typeId)


Fetches a User Type by ID. The special identifier `default` may be used to fetch the default User Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeId** | **string**|  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserTypes**
> []UserType ListUserTypes(ctx, )


Fetches all User Types in your org

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserType**](UserType.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceUserType**
> UserType ReplaceUserType(ctx, body, typeId)


Replace an existing User Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserType**](UserType.md)|  | 
  **typeId** | **string**|  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserType**
> UserType UpdateUserType(ctx, body, typeId)


Updates an existing User Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserType**](UserType.md)|  | 
  **typeId** | **string**|  | 

### Return type

[**UserType**](UserType.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

