# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateIdentityProvider**](IdentityProviderApi.md#ActivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/activate | Activate Identity Provider
[**CloneIdentityProviderKey**](IdentityProviderApi.md#CloneIdentityProviderKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/{keyId}/clone | Clone Signing Key Credential for IdP
[**CreateIdentityProvider**](IdentityProviderApi.md#CreateIdentityProvider) | **Post** /api/v1/idps | Add Identity Provider
[**CreateIdentityProviderKey**](IdentityProviderApi.md#CreateIdentityProviderKey) | **Post** /api/v1/idps/credentials/keys | Add X.509 Certificate Public Key
[**DeactivateIdentityProvider**](IdentityProviderApi.md#DeactivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/deactivate | Deactivate Identity Provider
[**DeleteIdentityProvider**](IdentityProviderApi.md#DeleteIdentityProvider) | **Delete** /api/v1/idps/{idpId} | Delete Identity Provider
[**DeleteIdentityProviderKey**](IdentityProviderApi.md#DeleteIdentityProviderKey) | **Delete** /api/v1/idps/credentials/keys/{keyId} | Delete Key
[**GenerateCsrForIdentityProvider**](IdentityProviderApi.md#GenerateCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs | Generate Certificate Signing Request for IdP
[**GenerateIdentityProviderSigningKey**](IdentityProviderApi.md#GenerateIdentityProviderSigningKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/generate | Generate New IdP Signing Key Credential
[**GetCsrForIdentityProvider**](IdentityProviderApi.md#GetCsrForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs/{csrId} | 
[**GetIdentityProvider**](IdentityProviderApi.md#GetIdentityProvider) | **Get** /api/v1/idps/{idpId} | Get Identity Provider
[**GetIdentityProviderApplicationUser**](IdentityProviderApi.md#GetIdentityProviderApplicationUser) | **Get** /api/v1/idps/{idpId}/users/{userId} | 
[**GetIdentityProviderKey**](IdentityProviderApi.md#GetIdentityProviderKey) | **Get** /api/v1/idps/credentials/keys/{keyId} | Get Key
[**GetIdentityProviderSigningKey**](IdentityProviderApi.md#GetIdentityProviderSigningKey) | **Get** /api/v1/idps/{idpId}/credentials/keys/{keyId} | Get Signing Key Credential for IdP
[**LinkUserToIdentityProvider**](IdentityProviderApi.md#LinkUserToIdentityProvider) | **Post** /api/v1/idps/{idpId}/users/{userId} | Link a user to a Social IdP without a transaction
[**ListCsrsForIdentityProvider**](IdentityProviderApi.md#ListCsrsForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs | List Certificate Signing Requests for IdP
[**ListIdentityProviderApplicationUsers**](IdentityProviderApi.md#ListIdentityProviderApplicationUsers) | **Get** /api/v1/idps/{idpId}/users | Find Users
[**ListIdentityProviderKeys**](IdentityProviderApi.md#ListIdentityProviderKeys) | **Get** /api/v1/idps/credentials/keys | List Keys
[**ListIdentityProviderSigningKeys**](IdentityProviderApi.md#ListIdentityProviderSigningKeys) | **Get** /api/v1/idps/{idpId}/credentials/keys | List Signing Key Credentials for IdP
[**ListIdentityProviders**](IdentityProviderApi.md#ListIdentityProviders) | **Get** /api/v1/idps | List Identity Providers
[**ListSocialAuthTokens**](IdentityProviderApi.md#ListSocialAuthTokens) | **Get** /api/v1/idps/{idpId}/users/{userId}/credentials/tokens | Social Authentication Token Operation
[**PublishCsrForIdentityProvider**](IdentityProviderApi.md#PublishCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs/{csrId}/lifecycle/publish | 
[**RevokeCsrForIdentityProvider**](IdentityProviderApi.md#RevokeCsrForIdentityProvider) | **Delete** /api/v1/idps/{idpId}/credentials/csrs/{csrId} | 
[**UnlinkUserFromIdentityProvider**](IdentityProviderApi.md#UnlinkUserFromIdentityProvider) | **Delete** /api/v1/idps/{idpId}/users/{userId} | Unlink User from IdP
[**UpdateIdentityProvider**](IdentityProviderApi.md#UpdateIdentityProvider) | **Put** /api/v1/idps/{idpId} | Update Identity Provider

# **ActivateIdentityProvider**
> IdentityProvider ActivateIdentityProvider(ctx, idpId)
Activate Identity Provider

Activates an inactive IdP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloneIdentityProviderKey**
> JsonWebKey CloneIdentityProviderKey(ctx, idpId, keyId, targetIdpId)
Clone Signing Key Credential for IdP

Clones a X.509 certificate for an IdP signing key credential from a source IdP to target IdP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **keyId** | **string**|  | 
  **targetIdpId** | **string**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIdentityProvider**
> IdentityProvider CreateIdentityProvider(ctx, body)
Add Identity Provider

Adds a new IdP to your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentityProvider**](IdentityProvider.md)|  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIdentityProviderKey**
> JsonWebKey CreateIdentityProviderKey(ctx, body)
Add X.509 Certificate Public Key

Adds a new X.509 certificate credential to the IdP key store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonWebKey**](JsonWebKey.md)|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateIdentityProvider**
> IdentityProvider DeactivateIdentityProvider(ctx, idpId)
Deactivate Identity Provider

Deactivates an active IdP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdentityProvider**
> DeleteIdentityProvider(ctx, idpId)
Delete Identity Provider

Removes an IdP from your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdentityProviderKey**
> DeleteIdentityProviderKey(ctx, keyId)
Delete Key

Deletes a specific IdP Key Credential by `kid` if it is not currently being used by an Active or Inactive IdP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateCsrForIdentityProvider**
> Csr GenerateCsrForIdentityProvider(ctx, body, idpId)
Generate Certificate Signing Request for IdP

Generates a new key pair and returns a Certificate Signing Request for it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CsrMetadata**](CsrMetadata.md)|  | 
  **idpId** | **string**|  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateIdentityProviderSigningKey**
> JsonWebKey GenerateIdentityProviderSigningKey(ctx, idpId, validityYears)
Generate New IdP Signing Key Credential

Generates a new X.509 certificate for an IdP signing key credential to be used for signing assertions sent to the IdP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **validityYears** | **int32**| expiry of the IdP Key Credential | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCsrForIdentityProvider**
> Csr GetCsrForIdentityProvider(ctx, idpId, csrId)


Gets a specific Certificate Signing Request model by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **csrId** | **string**|  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityProvider**
> IdentityProvider GetIdentityProvider(ctx, idpId)
Get Identity Provider

Fetches an IdP by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityProviderApplicationUser**
> IdentityProviderApplicationUser GetIdentityProviderApplicationUser(ctx, idpId, userId)


Fetches a linked IdP user by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**IdentityProviderApplicationUser**](IdentityProviderApplicationUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityProviderKey**
> JsonWebKey GetIdentityProviderKey(ctx, keyId)
Get Key

Gets a specific IdP Key Credential by `kid`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityProviderSigningKey**
> JsonWebKey GetIdentityProviderSigningKey(ctx, idpId, keyId)
Get Signing Key Credential for IdP

Gets a specific IdP Key Credential by `kid`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **keyId** | **string**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkUserToIdentityProvider**
> IdentityProviderApplicationUser LinkUserToIdentityProvider(ctx, body, idpId, userId)
Link a user to a Social IdP without a transaction

Links an Okta user to an existing Social Identity Provider. This does not support the SAML2 Identity Provider Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserIdentityProviderLinkRequest**](UserIdentityProviderLinkRequest.md)|  | 
  **idpId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**IdentityProviderApplicationUser**](IdentityProviderApplicationUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCsrsForIdentityProvider**
> []Csr ListCsrsForIdentityProvider(ctx, idpId)
List Certificate Signing Requests for IdP

Enumerates Certificate Signing Requests for an IdP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

[**[]Csr**](Csr.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdentityProviderApplicationUsers**
> []IdentityProviderApplicationUser ListIdentityProviderApplicationUsers(ctx, idpId)
Find Users

Find all the users linked to an identity provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

[**[]IdentityProviderApplicationUser**](IdentityProviderApplicationUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdentityProviderKeys**
> []JsonWebKey ListIdentityProviderKeys(ctx, optional)
List Keys

Enumerates IdP key credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityProviderApiListIdentityProviderKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityProviderApiListIdentityProviderKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **optional.String**| Specifies the pagination cursor for the next page of keys | 
 **limit** | **optional.Int32**| Specifies the number of key results in a page | [default to 20]

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdentityProviderSigningKeys**
> []JsonWebKey ListIdentityProviderSigningKeys(ctx, idpId)
List Signing Key Credentials for IdP

Enumerates signing key credentials for an IdP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdentityProviders**
> []IdentityProvider ListIdentityProviders(ctx, optional)
List Identity Providers

Enumerates IdPs in your organization with pagination. A subset of IdPs can be returned that match a supported filter expression or query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IdentityProviderApiListIdentityProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IdentityProviderApiListIdentityProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Searches the name property of IdPs for matching value | 
 **after** | **optional.String**| Specifies the pagination cursor for the next page of IdPs | 
 **limit** | **optional.Int32**| Specifies the number of IdP results in a page | [default to 20]
 **type_** | **optional.String**| Filters IdPs by type | 

### Return type

[**[]IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSocialAuthTokens**
> []SocialAuthToken ListSocialAuthTokens(ctx, idpId, userId)
Social Authentication Token Operation

Fetches the tokens minted by the Social Authentication Provider when the user authenticates with Okta via Social Auth.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**[]SocialAuthToken**](SocialAuthToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishCsrForIdentityProvider**
> JsonWebKey PublishCsrForIdentityProvider(ctx, body, idpId, csrId)


Update the Certificate Signing Request with a signed X.509 certificate and add it into the signing key credentials for the IdP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](Object.md)|  | 
  **idpId** | **string**|  | 
  **csrId** | **string**|  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/octet-stream, application/x-x509-ca-cert, application/pkix-cert, application/x-pem-file
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeCsrForIdentityProvider**
> RevokeCsrForIdentityProvider(ctx, idpId, csrId)


Revoke a Certificate Signing Request and delete the key pair from the IdP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **csrId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkUserFromIdentityProvider**
> UnlinkUserFromIdentityProvider(ctx, idpId, userId)
Unlink User from IdP

Removes the link between the Okta user and the IdP user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIdentityProvider**
> IdentityProvider UpdateIdentityProvider(ctx, body, idpId)
Update Identity Provider

Updates the configuration for an IdP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentityProvider**](IdentityProvider.md)|  | 
  **idpId** | **string**|  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

