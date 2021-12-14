# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthenticator**](AuthenticatorApi.md#ActivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/activate | Activate Authenticator
[**DeactivateAuthenticator**](AuthenticatorApi.md#DeactivateAuthenticator) | **Post** /api/v1/authenticators/{authenticatorId}/lifecycle/deactivate | Deactivate Authenticator
[**GetAuthenticator**](AuthenticatorApi.md#GetAuthenticator) | **Get** /api/v1/authenticators/{authenticatorId} | Get Authenticator
[**ListAuthenticators**](AuthenticatorApi.md#ListAuthenticators) | **Get** /api/v1/authenticators | List Authenticators

# **ActivateAuthenticator**
> Authenticator ActivateAuthenticator(ctx, authenticatorId)
Activate Authenticator

Activates an authenticator by `authenticatorId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authenticatorId** | **string**|  | 

### Return type

[**Authenticator**](Authenticator.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateAuthenticator**
> Authenticator DeactivateAuthenticator(ctx, authenticatorId)
Deactivate Authenticator

Deactivates an authenticator by `authenticatorId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authenticatorId** | **string**|  | 

### Return type

[**Authenticator**](Authenticator.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthenticator**
> Authenticator GetAuthenticator(ctx, authenticatorId)
Get Authenticator

Fetches an authenticator from your Okta organization by `authenticatorId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authenticatorId** | **string**|  | 

### Return type

[**Authenticator**](Authenticator.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthenticators**
> []Authenticator ListAuthenticators(ctx, )
List Authenticators

Enumerates authenticators in your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Authenticator**](Authenticator.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

