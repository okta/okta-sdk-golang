# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateFactor**](UserFactorApi.md#ActivateFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/lifecycle/activate | Activate Factor
[**DeleteFactor**](UserFactorApi.md#DeleteFactor) | **Delete** /api/v1/users/{userId}/factors/{factorId} | 
[**EnrollFactor**](UserFactorApi.md#EnrollFactor) | **Post** /api/v1/users/{userId}/factors | Enroll Factor
[**GetFactor**](UserFactorApi.md#GetFactor) | **Get** /api/v1/users/{userId}/factors/{factorId} | 
[**GetFactorTransactionStatus**](UserFactorApi.md#GetFactorTransactionStatus) | **Get** /api/v1/users/{userId}/factors/{factorId}/transactions/{transactionId} | 
[**ListFactors**](UserFactorApi.md#ListFactors) | **Get** /api/v1/users/{userId}/factors | 
[**ListSupportedFactors**](UserFactorApi.md#ListSupportedFactors) | **Get** /api/v1/users/{userId}/factors/catalog | 
[**ListSupportedSecurityQuestions**](UserFactorApi.md#ListSupportedSecurityQuestions) | **Get** /api/v1/users/{userId}/factors/questions | 
[**VerifyFactor**](UserFactorApi.md#VerifyFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/verify | Verify MFA Factor

# **ActivateFactor**
> UserFactor ActivateFactor(ctx, userId, factorId, optional)
Activate Factor

The `sms` and `token:software:totp` factor types require activation to complete the enrollment process.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **factorId** | **string**|  | 
 **optional** | ***UserFactorApiActivateFactorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserFactorApiActivateFactorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ActivateFactorRequest**](ActivateFactorRequest.md)|  | 

### Return type

[**UserFactor**](UserFactor.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFactor**
> DeleteFactor(ctx, userId, factorId)


Unenrolls an existing factor for the specified user, allowing the user to enroll a new factor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **factorId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnrollFactor**
> UserFactor EnrollFactor(ctx, body, userId, optional)
Enroll Factor

Enrolls a user with a supported factor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserFactor**](UserFactor.md)| Factor | 
  **userId** | **string**|  | 
 **optional** | ***UserFactorApiEnrollFactorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserFactorApiEnrollFactorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePhone** | **optional.**|  | [default to false]
 **templateId** | **optional.**| id of SMS template (only for SMS factor) | 
 **tokenLifetimeSeconds** | **optional.**|  | [default to 300]
 **activate** | **optional.**|  | [default to false]

### Return type

[**UserFactor**](UserFactor.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFactor**
> UserFactor GetFactor(ctx, userId, factorId)


Fetches a factor for the specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **factorId** | **string**|  | 

### Return type

[**UserFactor**](UserFactor.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFactorTransactionStatus**
> VerifyUserFactorResponse GetFactorTransactionStatus(ctx, userId, factorId, transactionId)


Polls factors verification transaction for status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **factorId** | **string**|  | 
  **transactionId** | **string**|  | 

### Return type

[**VerifyUserFactorResponse**](VerifyUserFactorResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFactors**
> []UserFactor ListFactors(ctx, userId)


Enumerates all the enrolled factors for the specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]UserFactor**](UserFactor.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSupportedFactors**
> []UserFactor ListSupportedFactors(ctx, userId)


Enumerates all the supported factors that can be enrolled for the specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]UserFactor**](UserFactor.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSupportedSecurityQuestions**
> []SecurityQuestion ListSupportedSecurityQuestions(ctx, userId)


Enumerates all available security questions for a user's `question` factor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]SecurityQuestion**](SecurityQuestion.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyFactor**
> VerifyUserFactorResponse VerifyFactor(ctx, userId, factorId, optional)
Verify MFA Factor

Verifies an OTP for a `token` or `token:hardware` factor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **factorId** | **string**|  | 
 **optional** | ***UserFactorApiVerifyFactorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserFactorApiVerifyFactorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VerifyFactorRequest**](VerifyFactorRequest.md)|  | 
 **templateId** | **optional.**|  | 
 **tokenLifetimeSeconds** | **optional.**|  | [default to 300]
 **xForwardedFor** | **optional.**|  | 
 **userAgent** | **optional.**|  | 
 **acceptLanguage** | **optional.**|  | 

### Return type

[**VerifyUserFactorResponse**](VerifyUserFactorResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

