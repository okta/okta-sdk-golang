# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateApplication**](ApplicationApi.md#ActivateApplication) | **Post** /api/v1/apps/{appId}/lifecycle/activate | Activate Application
[**AssignUserToApplication**](ApplicationApi.md#AssignUserToApplication) | **Post** /api/v1/apps/{appId}/users | Assign User to Application for SSO &amp; Provisioning
[**CloneApplicationKey**](ApplicationApi.md#CloneApplicationKey) | **Post** /api/v1/apps/{appId}/credentials/keys/{keyId}/clone | Clone Application Key Credential
[**CreateApplication**](ApplicationApi.md#CreateApplication) | **Post** /api/v1/apps | Add Application
[**CreateApplicationGroupAssignment**](ApplicationApi.md#CreateApplicationGroupAssignment) | **Put** /api/v1/apps/{appId}/groups/{groupId} | Assign Group to Application
[**DeactivateApplication**](ApplicationApi.md#DeactivateApplication) | **Post** /api/v1/apps/{appId}/lifecycle/deactivate | Deactivate Application
[**DeleteApplication**](ApplicationApi.md#DeleteApplication) | **Delete** /api/v1/apps/{appId} | Delete Application
[**DeleteApplicationGroupAssignment**](ApplicationApi.md#DeleteApplicationGroupAssignment) | **Delete** /api/v1/apps/{appId}/groups/{groupId} | Remove Group from Application
[**DeleteApplicationUser**](ApplicationApi.md#DeleteApplicationUser) | **Delete** /api/v1/apps/{appId}/users/{userId} | Remove User from Application
[**GenerateApplicationKey**](ApplicationApi.md#GenerateApplicationKey) | **Post** /api/v1/apps/{appId}/credentials/keys/generate | 
[**GenerateCsrForApplication**](ApplicationApi.md#GenerateCsrForApplication) | **Post** /api/v1/apps/{appId}/credentials/csrs | Generate Certificate Signing Request for Application
[**GetApplication**](ApplicationApi.md#GetApplication) | **Get** /api/v1/apps/{appId} | Get Application
[**GetApplicationGroupAssignment**](ApplicationApi.md#GetApplicationGroupAssignment) | **Get** /api/v1/apps/{appId}/groups/{groupId} | Get Assigned Group for Application
[**GetApplicationKey**](ApplicationApi.md#GetApplicationKey) | **Get** /api/v1/apps/{appId}/credentials/keys/{keyId} | Get Key Credential for Application
[**GetApplicationUser**](ApplicationApi.md#GetApplicationUser) | **Get** /api/v1/apps/{appId}/users/{userId} | Get Assigned User for Application
[**GetCsrForApplication**](ApplicationApi.md#GetCsrForApplication) | **Get** /api/v1/apps/{appId}/credentials/csrs/{csrId} | Get Certificate Signing Request
[**GetOAuth2TokenForApplication**](ApplicationApi.md#GetOAuth2TokenForApplication) | **Get** /api/v1/apps/{appId}/tokens/{tokenId} | 
[**GetScopeConsentGrant**](ApplicationApi.md#GetScopeConsentGrant) | **Get** /api/v1/apps/{appId}/grants/{grantId} | 
[**GrantConsentToScope**](ApplicationApi.md#GrantConsentToScope) | **Post** /api/v1/apps/{appId}/grants | 
[**ListApplicationGroupAssignments**](ApplicationApi.md#ListApplicationGroupAssignments) | **Get** /api/v1/apps/{appId}/groups | List Groups Assigned to Application
[**ListApplicationKeys**](ApplicationApi.md#ListApplicationKeys) | **Get** /api/v1/apps/{appId}/credentials/keys | List Key Credentials for Application
[**ListApplicationUsers**](ApplicationApi.md#ListApplicationUsers) | **Get** /api/v1/apps/{appId}/users | List Users Assigned to Application
[**ListApplications**](ApplicationApi.md#ListApplications) | **Get** /api/v1/apps | List Applications
[**ListCsrsForApplication**](ApplicationApi.md#ListCsrsForApplication) | **Get** /api/v1/apps/{appId}/credentials/csrs | List Certificate Signing Requests for Application
[**ListOAuth2TokensForApplication**](ApplicationApi.md#ListOAuth2TokensForApplication) | **Get** /api/v1/apps/{appId}/tokens | 
[**ListScopeConsentGrants**](ApplicationApi.md#ListScopeConsentGrants) | **Get** /api/v1/apps/{appId}/grants | 
[**PublishCsrFromApplication**](ApplicationApi.md#PublishCsrFromApplication) | **Post** /api/v1/apps/{appId}/credentials/csrs/{csrId}/lifecycle/publish | Publish Certificate Signing Request
[**RevokeCsrFromApplication**](ApplicationApi.md#RevokeCsrFromApplication) | **Delete** /api/v1/apps/{appId}/credentials/csrs/{csrId} | Revoke Certificate Signing Request
[**RevokeOAuth2TokenForApplication**](ApplicationApi.md#RevokeOAuth2TokenForApplication) | **Delete** /api/v1/apps/{appId}/tokens/{tokenId} | 
[**RevokeOAuth2TokensForApplication**](ApplicationApi.md#RevokeOAuth2TokensForApplication) | **Delete** /api/v1/apps/{appId}/tokens | 
[**RevokeScopeConsentGrant**](ApplicationApi.md#RevokeScopeConsentGrant) | **Delete** /api/v1/apps/{appId}/grants/{grantId} | 
[**UpdateApplication**](ApplicationApi.md#UpdateApplication) | **Put** /api/v1/apps/{appId} | Update Application
[**UpdateApplicationUser**](ApplicationApi.md#UpdateApplicationUser) | **Post** /api/v1/apps/{appId}/users/{userId} | Update Application Profile for Assigned User

# **ActivateApplication**
> ActivateApplication(ctx, appId)
Activate Application

Activates an inactive application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignUserToApplication**
> AppUser AssignUserToApplication(ctx, body, appId)
Assign User to Application for SSO & Provisioning

Assigns an user to an application with [credentials](#application-user-credentials-object) and an app-specific [profile](#application-user-profile-object). Profile mappings defined for the application are first applied before applying any profile properties specified in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppUser**](AppUser.md)|  | 
  **appId** | **string**|  | 

### Return type

[**AppUser**](AppUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloneApplicationKey**
> JsonWebKey CloneApplicationKey(ctx, appId, keyId, targetAid)
Clone Application Key Credential

Clones a X.509 certificate for an application key credential from a source application to target application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **keyId** | **string**|  | 
  **targetAid** | **string**| Unique key of the target Application | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplication**
> Application CreateApplication(ctx, body, optional)
Add Application

Adds a new application to your Okta organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Application**](Application.md)|  | 
 **optional** | ***ApplicationApiCreateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiCreateApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activate** | **optional.**| Executes activation lifecycle operation when creating the app | [default to true]
 **oktaAccessGatewayAgent** | **optional.**|  | 

### Return type

[**Application**](Application.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplicationGroupAssignment**
> ApplicationGroupAssignment CreateApplicationGroupAssignment(ctx, appId, groupId, optional)
Assign Group to Application

Assigns a group to an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***ApplicationApiCreateApplicationGroupAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiCreateApplicationGroupAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApplicationGroupAssignment**](ApplicationGroupAssignment.md)|  | 

### Return type

[**ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateApplication**
> DeactivateApplication(ctx, appId)
Deactivate Application

Deactivates an active application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplication**
> DeleteApplication(ctx, appId)
Delete Application

Removes an inactive application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationGroupAssignment**
> DeleteApplicationGroupAssignment(ctx, appId, groupId)
Remove Group from Application

Removes a group assignment from an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationUser**
> DeleteApplicationUser(ctx, appId, userId, optional)
Remove User from Application

Removes an assignment for a user from an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **userId** | **string**|  | 
 **optional** | ***ApplicationApiDeleteApplicationUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiDeleteApplicationUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendEmail** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateApplicationKey**
> JsonWebKey GenerateApplicationKey(ctx, appId, optional)


Generates a new X.509 certificate for an application key credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
 **optional** | ***ApplicationApiGenerateApplicationKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGenerateApplicationKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validityYears** | **optional.Int32**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateCsrForApplication**
> Csr GenerateCsrForApplication(ctx, body, appId)
Generate Certificate Signing Request for Application

Generates a new key pair and returns the Certificate Signing Request for it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CsrMetadata**](CsrMetadata.md)|  | 
  **appId** | **string**|  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplication**
> Application GetApplication(ctx, appId, optional)
Get Application

Fetches an application from your Okta organization by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
 **optional** | ***ApplicationApiGetApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**|  | 

### Return type

[**Application**](Application.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationGroupAssignment**
> ApplicationGroupAssignment GetApplicationGroupAssignment(ctx, appId, groupId, optional)
Get Assigned Group for Application

Fetches an application group assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***ApplicationApiGetApplicationGroupAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetApplicationGroupAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 

### Return type

[**ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationKey**
> JsonWebKey GetApplicationKey(ctx, appId, keyId)
Get Key Credential for Application

Gets a specific application key credential by kid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **keyId** | **string**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationUser**
> AppUser GetApplicationUser(ctx, appId, userId, optional)
Get Assigned User for Application

Fetches a specific user assignment for application by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **userId** | **string**|  | 
 **optional** | ***ApplicationApiGetApplicationUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetApplicationUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 

### Return type

[**AppUser**](AppUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCsrForApplication**
> Csr GetCsrForApplication(ctx, appId, csrId)
Get Certificate Signing Request

Fetches a certificate signing request for the app by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **csrId** | **string**|  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOAuth2TokenForApplication**
> OAuth2Token GetOAuth2TokenForApplication(ctx, appId, tokenId, optional)


Gets a token for the specified application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **tokenId** | **string**|  | 
 **optional** | ***ApplicationApiGetOAuth2TokenForApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetOAuth2TokenForApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 

### Return type

[**OAuth2Token**](OAuth2Token.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScopeConsentGrant**
> OAuth2ScopeConsentGrant GetScopeConsentGrant(ctx, appId, grantId, optional)


Fetches a single scope consent grant for the application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **grantId** | **string**|  | 
 **optional** | ***ApplicationApiGetScopeConsentGrantOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetScopeConsentGrantOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrantConsentToScope**
> OAuth2ScopeConsentGrant GrantConsentToScope(ctx, body, appId)


Grants consent for the application to request an OAuth 2.0 Okta scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)|  | 
  **appId** | **string**|  | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationGroupAssignments**
> []ApplicationGroupAssignment ListApplicationGroupAssignments(ctx, appId, optional)
List Groups Assigned to Application

Enumerates group assignments for an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
 **optional** | ***ApplicationApiListApplicationGroupAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiListApplicationGroupAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  | 
 **after** | **optional.String**| Specifies the pagination cursor for the next page of assignments | 
 **limit** | **optional.Int32**| Specifies the number of results for a page | [default to -1]
 **expand** | **optional.String**|  | 

### Return type

[**[]ApplicationGroupAssignment**](ApplicationGroupAssignment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationKeys**
> []JsonWebKey ListApplicationKeys(ctx, appId)
List Key Credentials for Application

Enumerates key credentials for an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationUsers**
> []AppUser ListApplicationUsers(ctx, appId, optional)
List Users Assigned to Application

Enumerates all assigned [application users](#application-user-model) for an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
 **optional** | ***ApplicationApiListApplicationUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiListApplicationUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  | 
 **queryScope** | **optional.String**|  | 
 **after** | **optional.String**| specifies the pagination cursor for the next page of assignments | 
 **limit** | **optional.Int32**| specifies the number of results for a page | [default to -1]
 **filter** | **optional.String**|  | 
 **expand** | **optional.String**|  | 

### Return type

[**[]AppUser**](AppUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplications**
> []Application ListApplications(ctx, optional)
List Applications

Enumerates apps added to your organization with pagination. A subset of apps can be returned that match a supported filter expression or query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiListApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiListApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  | 
 **after** | **optional.String**| Specifies the pagination cursor for the next page of apps | 
 **limit** | **optional.Int32**| Specifies the number of results for a page | [default to -1]
 **filter** | **optional.String**| Filters apps by status, user.id, group.id or credentials.signing.kid expression | 
 **expand** | **optional.String**| Traverses users link relationship and optionally embeds Application User resource | 
 **includeNonDeleted** | **optional.Bool**|  | [default to false]

### Return type

[**[]Application**](Application.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCsrsForApplication**
> []Csr ListCsrsForApplication(ctx, appId)
List Certificate Signing Requests for Application

Enumerates Certificate Signing Requests for an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

[**[]Csr**](Csr.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOAuth2TokensForApplication**
> []OAuth2Token ListOAuth2TokensForApplication(ctx, appId, optional)


Lists all tokens for the application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
 **optional** | ***ApplicationApiListOAuth2TokensForApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiListOAuth2TokensForApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**|  | 
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]

### Return type

[**[]OAuth2Token**](OAuth2Token.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListScopeConsentGrants**
> []OAuth2ScopeConsentGrant ListScopeConsentGrants(ctx, appId, optional)


Lists all scope consent grants for the application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
 **optional** | ***ApplicationApiListScopeConsentGrantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiListScopeConsentGrantsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**|  | 

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishCsrFromApplication**
> JsonWebKey PublishCsrFromApplication(ctx, body, appId, csrId)
Publish Certificate Signing Request

Updates a certificate signing request for the app with a signed X.509 certificate and adds it into the application key credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](Object.md)|  | 
  **appId** | **string**|  | 
  **csrId** | **string**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/octet-stream, application/x-x509-ca-cert, application/pkix-cert, application/x-pem-file
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeCsrFromApplication**
> RevokeCsrFromApplication(ctx, appId, csrId)
Revoke Certificate Signing Request

Revokes a certificate signing request and deletes the key pair from the application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **csrId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeOAuth2TokenForApplication**
> RevokeOAuth2TokenForApplication(ctx, appId, tokenId)


Revokes the specified token for the specified application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **tokenId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeOAuth2TokensForApplication**
> RevokeOAuth2TokensForApplication(ctx, appId)


Revokes all tokens for the specified application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeScopeConsentGrant**
> RevokeScopeConsentGrant(ctx, appId, grantId)


Revokes permission for the application to request the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **grantId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplication**
> Application UpdateApplication(ctx, body, appId)
Update Application

Updates an application in your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Application**](Application.md)|  | 
  **appId** | **string**|  | 

### Return type

[**Application**](Application.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationUser**
> AppUser UpdateApplicationUser(ctx, body, appId, userId)
Update Application Profile for Assigned User

Updates a user's profile for an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppUser**](AppUser.md)|  | 
  **appId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**AppUser**](AppUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

