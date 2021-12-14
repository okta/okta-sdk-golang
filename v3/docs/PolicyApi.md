# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePolicy**](PolicyApi.md#ActivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/activate | 
[**ActivatePolicyRule**](PolicyApi.md#ActivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/activate | 
[**CreatePolicy**](PolicyApi.md#CreatePolicy) | **Post** /api/v1/policies | 
[**CreatePolicyRule**](PolicyApi.md#CreatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules | 
[**DeactivatePolicy**](PolicyApi.md#DeactivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/deactivate | 
[**DeactivatePolicyRule**](PolicyApi.md#DeactivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | 
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Delete** /api/v1/policies/{policyId} | 
[**DeletePolicyRule**](PolicyApi.md#DeletePolicyRule) | **Delete** /api/v1/policies/{policyId}/rules/{ruleId} | 
[**GetPolicy**](PolicyApi.md#GetPolicy) | **Get** /api/v1/policies/{policyId} | 
[**GetPolicyRule**](PolicyApi.md#GetPolicyRule) | **Get** /api/v1/policies/{policyId}/rules/{ruleId} | 
[**ListPolicies**](PolicyApi.md#ListPolicies) | **Get** /api/v1/policies | 
[**ListPolicyRules**](PolicyApi.md#ListPolicyRules) | **Get** /api/v1/policies/{policyId}/rules | 
[**UpdatePolicy**](PolicyApi.md#UpdatePolicy) | **Put** /api/v1/policies/{policyId} | 
[**UpdatePolicyRule**](PolicyApi.md#UpdatePolicyRule) | **Put** /api/v1/policies/{policyId}/rules/{ruleId} | 

# **ActivatePolicy**
> ActivatePolicy(ctx, policyId)


Activates a policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivatePolicyRule**
> ActivatePolicyRule(ctx, policyId, ruleId)


Activates a policy rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicy**
> Policy CreatePolicy(ctx, body, optional)


Creates a policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Policy**](Policy.md)|  | 
 **optional** | ***PolicyApiCreatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiCreatePolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activate** | **optional.**|  | [default to true]

### Return type

[**Policy**](Policy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyRule**
> PolicyRule CreatePolicyRule(ctx, body, policyId)


Creates a policy rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyRule**](PolicyRule.md)|  | 
  **policyId** | **string**|  | 

### Return type

[**PolicyRule**](PolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivatePolicy**
> DeactivatePolicy(ctx, policyId)


Deactivates a policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivatePolicyRule**
> DeactivatePolicyRule(ctx, policyId, ruleId)


Deactivates a policy rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> DeletePolicy(ctx, policyId)


Removes a policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyRule**
> DeletePolicyRule(ctx, policyId, ruleId)


Removes a policy rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> Policy GetPolicy(ctx, policyId, optional)


Gets a policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
 **optional** | ***PolicyApiGetPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiGetPolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**|  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyRule**
> PolicyRule GetPolicyRule(ctx, policyId, ruleId)


Gets a policy rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**PolicyRule**](PolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicies**
> []Policy ListPolicies(ctx, type_, optional)


Gets all policies with the specified type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
 **optional** | ***PolicyApiListPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiListPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**|  | 
 **expand** | **optional.String**|  | 

### Return type

[**[]Policy**](Policy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicyRules**
> []PolicyRule ListPolicyRules(ctx, policyId)


Enumerates all policy rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**[]PolicyRule**](PolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicy**
> Policy UpdatePolicy(ctx, body, policyId)


Updates a policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Policy**](Policy.md)|  | 
  **policyId** | **string**|  | 

### Return type

[**Policy**](Policy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyRule**
> PolicyRule UpdatePolicyRule(ctx, body, policyId, ruleId)


Updates a policy rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PolicyRule**](PolicyRule.md)|  | 
  **policyId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**PolicyRule**](PolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

