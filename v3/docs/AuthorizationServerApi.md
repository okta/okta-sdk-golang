# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthorizationServer**](AuthorizationServerApi.md#ActivateAuthorizationServer) | **Post** /api/v1/authorizationServers/{authServerId}/lifecycle/activate | 
[**ActivateAuthorizationServerPolicy**](AuthorizationServerApi.md#ActivateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/lifecycle/activate | 
[**ActivateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#ActivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/activate | 
[**CreateAuthorizationServer**](AuthorizationServerApi.md#CreateAuthorizationServer) | **Post** /api/v1/authorizationServers | 
[**CreateAuthorizationServerPolicy**](AuthorizationServerApi.md#CreateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies | 
[**CreateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#CreateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | 
[**CreateOAuth2Claim**](AuthorizationServerApi.md#CreateOAuth2Claim) | **Post** /api/v1/authorizationServers/{authServerId}/claims | 
[**CreateOAuth2Scope**](AuthorizationServerApi.md#CreateOAuth2Scope) | **Post** /api/v1/authorizationServers/{authServerId}/scopes | 
[**DeactivateAuthorizationServer**](AuthorizationServerApi.md#DeactivateAuthorizationServer) | **Post** /api/v1/authorizationServers/{authServerId}/lifecycle/deactivate | 
[**DeactivateAuthorizationServerPolicy**](AuthorizationServerApi.md#DeactivateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/lifecycle/deactivate | 
[**DeactivateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#DeactivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | 
[**DeleteAuthorizationServer**](AuthorizationServerApi.md#DeleteAuthorizationServer) | **Delete** /api/v1/authorizationServers/{authServerId} | 
[**DeleteAuthorizationServerPolicy**](AuthorizationServerApi.md#DeleteAuthorizationServerPolicy) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | 
[**DeleteAuthorizationServerPolicyRule**](AuthorizationServerApi.md#DeleteAuthorizationServerPolicyRule) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | 
[**DeleteOAuth2Claim**](AuthorizationServerApi.md#DeleteOAuth2Claim) | **Delete** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | 
[**DeleteOAuth2Scope**](AuthorizationServerApi.md#DeleteOAuth2Scope) | **Delete** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | 
[**GetAuthorizationServer**](AuthorizationServerApi.md#GetAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId} | 
[**GetAuthorizationServerPolicy**](AuthorizationServerApi.md#GetAuthorizationServerPolicy) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | 
[**GetAuthorizationServerPolicyRule**](AuthorizationServerApi.md#GetAuthorizationServerPolicyRule) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | 
[**GetOAuth2Claim**](AuthorizationServerApi.md#GetOAuth2Claim) | **Get** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | 
[**GetOAuth2Scope**](AuthorizationServerApi.md#GetOAuth2Scope) | **Get** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | 
[**GetRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerApi.md#GetRefreshTokenForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | 
[**ListAuthorizationServerKeys**](AuthorizationServerApi.md#ListAuthorizationServerKeys) | **Get** /api/v1/authorizationServers/{authServerId}/credentials/keys | 
[**ListAuthorizationServerPolicies**](AuthorizationServerApi.md#ListAuthorizationServerPolicies) | **Get** /api/v1/authorizationServers/{authServerId}/policies | 
[**ListAuthorizationServerPolicyRules**](AuthorizationServerApi.md#ListAuthorizationServerPolicyRules) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | 
[**ListAuthorizationServers**](AuthorizationServerApi.md#ListAuthorizationServers) | **Get** /api/v1/authorizationServers | 
[**ListOAuth2Claims**](AuthorizationServerApi.md#ListOAuth2Claims) | **Get** /api/v1/authorizationServers/{authServerId}/claims | 
[**ListOAuth2ClientsForAuthorizationServer**](AuthorizationServerApi.md#ListOAuth2ClientsForAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId}/clients | 
[**ListOAuth2Scopes**](AuthorizationServerApi.md#ListOAuth2Scopes) | **Get** /api/v1/authorizationServers/{authServerId}/scopes | 
[**ListRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerApi.md#ListRefreshTokensForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | 
[**RevokeRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerApi.md#RevokeRefreshTokenForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | 
[**RevokeRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerApi.md#RevokeRefreshTokensForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | 
[**RotateAuthorizationServerKeys**](AuthorizationServerApi.md#RotateAuthorizationServerKeys) | **Post** /api/v1/authorizationServers/{authServerId}/credentials/lifecycle/keyRotate | 
[**UpdateAuthorizationServer**](AuthorizationServerApi.md#UpdateAuthorizationServer) | **Put** /api/v1/authorizationServers/{authServerId} | 
[**UpdateAuthorizationServerPolicy**](AuthorizationServerApi.md#UpdateAuthorizationServerPolicy) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | 
[**UpdateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#UpdateAuthorizationServerPolicyRule) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | 
[**UpdateOAuth2Claim**](AuthorizationServerApi.md#UpdateOAuth2Claim) | **Put** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | 
[**UpdateOAuth2Scope**](AuthorizationServerApi.md#UpdateOAuth2Scope) | **Put** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | 

# **ActivateAuthorizationServer**
> ActivateAuthorizationServer(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivateAuthorizationServerPolicy**
> ActivateAuthorizationServerPolicy(ctx, authServerId, policyId)


Activate Authorization Server Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivateAuthorizationServerPolicyRule**
> ActivateAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId)


Activate Authorization Server Policy Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
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

# **CreateAuthorizationServer**
> AuthorizationServer CreateAuthorizationServer(ctx, body)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationServer**](AuthorizationServer.md)|  | 

### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthorizationServerPolicy**
> AuthorizationServerPolicy CreateAuthorizationServerPolicy(ctx, body, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)|  | 
  **authServerId** | **string**|  | 

### Return type

[**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthorizationServerPolicyRule**
> AuthorizationServerPolicyRule CreateAuthorizationServerPolicyRule(ctx, body, policyId, authServerId)


Creates a policy rule for the specified Custom Authorization Server and Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)|  | 
  **policyId** | **string**|  | 
  **authServerId** | **string**|  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOAuth2Claim**
> OAuth2Claim CreateOAuth2Claim(ctx, body, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OAuth2Claim**](OAuth2Claim.md)|  | 
  **authServerId** | **string**|  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOAuth2Scope**
> OAuth2Scope CreateOAuth2Scope(ctx, body, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OAuth2Scope**](OAuth2Scope.md)|  | 
  **authServerId** | **string**|  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateAuthorizationServer**
> DeactivateAuthorizationServer(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateAuthorizationServerPolicy**
> DeactivateAuthorizationServerPolicy(ctx, authServerId, policyId)


Deactivate Authorization Server Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateAuthorizationServerPolicyRule**
> DeactivateAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId)


Deactivate Authorization Server Policy Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
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

# **DeleteAuthorizationServer**
> DeleteAuthorizationServer(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthorizationServerPolicy**
> DeleteAuthorizationServerPolicy(ctx, authServerId, policyId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthorizationServerPolicyRule**
> DeleteAuthorizationServerPolicyRule(ctx, policyId, authServerId, ruleId)


Deletes a Policy Rule defined in the specified Custom Authorization Server and Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **authServerId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOAuth2Claim**
> DeleteOAuth2Claim(ctx, authServerId, claimId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **claimId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOAuth2Scope**
> DeleteOAuth2Scope(ctx, authServerId, scopeId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorizationServer**
> AuthorizationServer GetAuthorizationServer(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorizationServerPolicy**
> AuthorizationServerPolicy GetAuthorizationServerPolicy(ctx, authServerId, policyId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

[**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorizationServerPolicyRule**
> AuthorizationServerPolicyRule GetAuthorizationServerPolicyRule(ctx, policyId, authServerId, ruleId)


Returns a Policy Rule by ID that is defined in the specified Custom Authorization Server and Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **authServerId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOAuth2Claim**
> OAuth2Claim GetOAuth2Claim(ctx, authServerId, claimId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **claimId** | **string**|  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOAuth2Scope**
> OAuth2Scope GetOAuth2Scope(ctx, authServerId, scopeId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **scopeId** | **string**|  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRefreshTokenForAuthorizationServerAndClient**
> OAuth2RefreshToken GetRefreshTokenForAuthorizationServerAndClient(ctx, authServerId, clientId, tokenId, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **clientId** | **string**|  | 
  **tokenId** | **string**|  | 
 **optional** | ***AuthorizationServerApiGetRefreshTokenForAuthorizationServerAndClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationServerApiGetRefreshTokenForAuthorizationServerAndClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **optional.String**|  | 

### Return type

[**OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthorizationServerKeys**
> []JsonWebKey ListAuthorizationServerKeys(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthorizationServerPolicies**
> []AuthorizationServerPolicy ListAuthorizationServerPolicies(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

[**[]AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthorizationServerPolicyRules**
> []AuthorizationServerPolicyRule ListAuthorizationServerPolicyRules(ctx, policyId, authServerId)


Enumerates all policy rules for the specified Custom Authorization Server and Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
  **authServerId** | **string**|  | 

### Return type

[**[]AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthorizationServers**
> []AuthorizationServer ListAuthorizationServers(ctx, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationServerApiListAuthorizationServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationServerApiListAuthorizationServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  | 
 **limit** | **optional.String**|  | 
 **after** | **optional.String**|  | 

### Return type

[**[]AuthorizationServer**](AuthorizationServer.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOAuth2Claims**
> []OAuth2Claim ListOAuth2Claims(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

[**[]OAuth2Claim**](OAuth2Claim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOAuth2ClientsForAuthorizationServer**
> []OAuth2Client ListOAuth2ClientsForAuthorizationServer(ctx, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 

### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOAuth2Scopes**
> []OAuth2Scope ListOAuth2Scopes(ctx, authServerId, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
 **optional** | ***AuthorizationServerApiListOAuth2ScopesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationServerApiListOAuth2ScopesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  | 
 **filter** | **optional.String**|  | 
 **cursor** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to -1]

### Return type

[**[]OAuth2Scope**](OAuth2Scope.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRefreshTokensForAuthorizationServerAndClient**
> []OAuth2RefreshToken ListRefreshTokensForAuthorizationServerAndClient(ctx, authServerId, clientId, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **clientId** | **string**|  | 
 **optional** | ***AuthorizationServerApiListRefreshTokensForAuthorizationServerAndClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationServerApiListRefreshTokensForAuthorizationServerAndClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to -1]

### Return type

[**[]OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeRefreshTokenForAuthorizationServerAndClient**
> RevokeRefreshTokenForAuthorizationServerAndClient(ctx, authServerId, clientId, tokenId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **clientId** | **string**|  | 
  **tokenId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeRefreshTokensForAuthorizationServerAndClient**
> RevokeRefreshTokensForAuthorizationServerAndClient(ctx, authServerId, clientId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authServerId** | **string**|  | 
  **clientId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateAuthorizationServerKeys**
> []JsonWebKey RotateAuthorizationServerKeys(ctx, body, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JwkUse**](JwkUse.md)|  | 
  **authServerId** | **string**|  | 

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthorizationServer**
> AuthorizationServer UpdateAuthorizationServer(ctx, body, authServerId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationServer**](AuthorizationServer.md)|  | 
  **authServerId** | **string**|  | 

### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthorizationServerPolicy**
> AuthorizationServerPolicy UpdateAuthorizationServerPolicy(ctx, body, authServerId, policyId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)|  | 
  **authServerId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

[**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthorizationServerPolicyRule**
> AuthorizationServerPolicyRule UpdateAuthorizationServerPolicyRule(ctx, body, policyId, authServerId, ruleId)


Updates the configuration of the Policy Rule defined in the specified Custom Authorization Server and Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)|  | 
  **policyId** | **string**|  | 
  **authServerId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOAuth2Claim**
> OAuth2Claim UpdateOAuth2Claim(ctx, body, authServerId, claimId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OAuth2Claim**](OAuth2Claim.md)|  | 
  **authServerId** | **string**|  | 
  **claimId** | **string**|  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOAuth2Scope**
> OAuth2Scope UpdateOAuth2Scope(ctx, body, authServerId, scopeId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OAuth2Scope**](OAuth2Scope.md)|  | 
  **authServerId** | **string**|  | 
  **scopeId** | **string**|  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

