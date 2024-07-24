# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | Pointer to **string** | The type of client application. Default value: &#x60;web&#x60;. | [optional] 
**ClientId** | Pointer to **string** | Unique key for the client application. The &#x60;client_id&#x60; is immutable. When you create a client Application, you can&#39;t specify the &#x60;client_id&#x60; because Okta uses the application ID for the &#x60;client_id&#x60;. | [optional] [readonly] 
**ClientIdIssuedAt** | Pointer to **int32** | Time at which the &#x60;client_id&#x60; was issued (measured in unix seconds) | [optional] [readonly] 
**ClientName** | Pointer to **string** | Human-readable string name of the client application | [optional] 
**ClientSecret** | Pointer to **NullableString** | OAuth 2.0 client secret string (used for confidential clients). The &#x60;client_secret&#x60; is shown only on the response of the creation or update of a client Application (and only if the &#x60;token_endpoint_auth_method&#x60; is one that requires a client secret). You can&#39;t specify the &#x60;client_secret&#x60;. If the &#x60;token_endpoint_auth_method&#x60; requires one, Okta generates a random &#x60;client_secret&#x60; for the client Application. | [optional] [readonly] 
**ClientSecretExpiresAt** | Pointer to **NullableInt32** | Time at which the &#x60;client_secret&#x60; expires or 0 if it doesn&#39;t expire (measured in unix seconds) | [optional] [readonly] 
**FrontchannelLogoutSessionRequired** | Pointer to **bool** | Include user session details | [optional] 
**FrontchannelLogoutUri** | Pointer to **NullableString** | URL where Okta sends the logout request | [optional] 
**GrantTypes** | Pointer to **[]string** | Array of OAuth 2.0 grant type strings. Default value: &#x60;[authorization_code]&#x60; | [optional] 
**InitiateLoginUri** | Pointer to **string** | URL that a third party can use to initiate a login by the client | [optional] 
**JwksUri** | Pointer to **string** | URL string that references a [JSON Web Key Set](https://tools.ietf.org/html/rfc7517#section-5) for validating JWTs presented to Okta | [optional] 
**LogoUri** | Pointer to **NullableString** | URL string that references a logo for the client consent dialog (not the sign-in dialog) | [optional] 
**PolicyUri** | Pointer to **NullableString** | URL string of a web page providing the client&#39;s policy document | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** | Array of redirection URI strings for use for relying party initiated logouts | [optional] 
**RedirectUris** | Pointer to **[]string** | Array of redirection URI strings for use in redirect-based flows. All redirect URIs must be absolute URIs and must not include a fragment component. At least one redirect URI and response type is required for all client types, with the following exceptions: If the client uses the Resource Owner Password flow (if &#x60;grant_type&#x60; contains the value password) or the Client Credentials flow (if &#x60;grant_type&#x60; contains the value &#x60;client_credentials&#x60;), then no redirect URI or response type is necessary. In these cases, you can pass either null or an empty array for these attributes. | [optional] 
**RequestObjectSigningAlg** | Pointer to **[]string** | The type of [JSON Web Key Set](https://tools.ietf.org/html/rfc7517#section-5) algorithm that must be used for signing request objects | [optional] 
**ResponseTypes** | Pointer to **[]string** | Array of OAuth 2.0 response type strings. Default value: &#x60;[code]&#x60; | [optional] 
**TokenEndpointAuthMethod** | Pointer to **string** | Requested authentication method for OAuth 2.0 endpoints. | [optional] 
**TosUri** | Pointer to **NullableString** | URL string of a web page providing the client&#39;s terms of service document | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *Client) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *Client) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *Client) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *Client) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetClientId

`func (o *Client) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Client) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdIssuedAt

`func (o *Client) GetClientIdIssuedAt() int32`

GetClientIdIssuedAt returns the ClientIdIssuedAt field if non-nil, zero value otherwise.

### GetClientIdIssuedAtOk

`func (o *Client) GetClientIdIssuedAtOk() (*int32, bool)`

GetClientIdIssuedAtOk returns a tuple with the ClientIdIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdIssuedAt

`func (o *Client) SetClientIdIssuedAt(v int32)`

SetClientIdIssuedAt sets ClientIdIssuedAt field to given value.

### HasClientIdIssuedAt

`func (o *Client) HasClientIdIssuedAt() bool`

HasClientIdIssuedAt returns a boolean if a field has been set.

### GetClientName

`func (o *Client) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *Client) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *Client) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *Client) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientSecret

`func (o *Client) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Client) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Client) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Client) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *Client) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *Client) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetClientSecretExpiresAt

`func (o *Client) GetClientSecretExpiresAt() int32`

GetClientSecretExpiresAt returns the ClientSecretExpiresAt field if non-nil, zero value otherwise.

### GetClientSecretExpiresAtOk

`func (o *Client) GetClientSecretExpiresAtOk() (*int32, bool)`

GetClientSecretExpiresAtOk returns a tuple with the ClientSecretExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretExpiresAt

`func (o *Client) SetClientSecretExpiresAt(v int32)`

SetClientSecretExpiresAt sets ClientSecretExpiresAt field to given value.

### HasClientSecretExpiresAt

`func (o *Client) HasClientSecretExpiresAt() bool`

HasClientSecretExpiresAt returns a boolean if a field has been set.

### SetClientSecretExpiresAtNil

`func (o *Client) SetClientSecretExpiresAtNil(b bool)`

 SetClientSecretExpiresAtNil sets the value for ClientSecretExpiresAt to be an explicit nil

### UnsetClientSecretExpiresAt
`func (o *Client) UnsetClientSecretExpiresAt()`

UnsetClientSecretExpiresAt ensures that no value is present for ClientSecretExpiresAt, not even an explicit nil
### GetFrontchannelLogoutSessionRequired

`func (o *Client) GetFrontchannelLogoutSessionRequired() bool`

GetFrontchannelLogoutSessionRequired returns the FrontchannelLogoutSessionRequired field if non-nil, zero value otherwise.

### GetFrontchannelLogoutSessionRequiredOk

`func (o *Client) GetFrontchannelLogoutSessionRequiredOk() (*bool, bool)`

GetFrontchannelLogoutSessionRequiredOk returns a tuple with the FrontchannelLogoutSessionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontchannelLogoutSessionRequired

`func (o *Client) SetFrontchannelLogoutSessionRequired(v bool)`

SetFrontchannelLogoutSessionRequired sets FrontchannelLogoutSessionRequired field to given value.

### HasFrontchannelLogoutSessionRequired

`func (o *Client) HasFrontchannelLogoutSessionRequired() bool`

HasFrontchannelLogoutSessionRequired returns a boolean if a field has been set.

### GetFrontchannelLogoutUri

`func (o *Client) GetFrontchannelLogoutUri() string`

GetFrontchannelLogoutUri returns the FrontchannelLogoutUri field if non-nil, zero value otherwise.

### GetFrontchannelLogoutUriOk

`func (o *Client) GetFrontchannelLogoutUriOk() (*string, bool)`

GetFrontchannelLogoutUriOk returns a tuple with the FrontchannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontchannelLogoutUri

`func (o *Client) SetFrontchannelLogoutUri(v string)`

SetFrontchannelLogoutUri sets FrontchannelLogoutUri field to given value.

### HasFrontchannelLogoutUri

`func (o *Client) HasFrontchannelLogoutUri() bool`

HasFrontchannelLogoutUri returns a boolean if a field has been set.

### SetFrontchannelLogoutUriNil

`func (o *Client) SetFrontchannelLogoutUriNil(b bool)`

 SetFrontchannelLogoutUriNil sets the value for FrontchannelLogoutUri to be an explicit nil

### UnsetFrontchannelLogoutUri
`func (o *Client) UnsetFrontchannelLogoutUri()`

UnsetFrontchannelLogoutUri ensures that no value is present for FrontchannelLogoutUri, not even an explicit nil
### GetGrantTypes

`func (o *Client) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *Client) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *Client) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *Client) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetInitiateLoginUri

`func (o *Client) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *Client) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *Client) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *Client) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetJwksUri

`func (o *Client) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *Client) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *Client) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *Client) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetLogoUri

`func (o *Client) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *Client) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *Client) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *Client) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### SetLogoUriNil

`func (o *Client) SetLogoUriNil(b bool)`

 SetLogoUriNil sets the value for LogoUri to be an explicit nil

### UnsetLogoUri
`func (o *Client) UnsetLogoUri()`

UnsetLogoUri ensures that no value is present for LogoUri, not even an explicit nil
### GetPolicyUri

`func (o *Client) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *Client) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *Client) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *Client) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### SetPolicyUriNil

`func (o *Client) SetPolicyUriNil(b bool)`

 SetPolicyUriNil sets the value for PolicyUri to be an explicit nil

### UnsetPolicyUri
`func (o *Client) UnsetPolicyUri()`

UnsetPolicyUri ensures that no value is present for PolicyUri, not even an explicit nil
### GetPostLogoutRedirectUris

`func (o *Client) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *Client) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *Client) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *Client) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *Client) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *Client) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *Client) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *Client) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRequestObjectSigningAlg

`func (o *Client) GetRequestObjectSigningAlg() []string`

GetRequestObjectSigningAlg returns the RequestObjectSigningAlg field if non-nil, zero value otherwise.

### GetRequestObjectSigningAlgOk

`func (o *Client) GetRequestObjectSigningAlgOk() (*[]string, bool)`

GetRequestObjectSigningAlgOk returns a tuple with the RequestObjectSigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectSigningAlg

`func (o *Client) SetRequestObjectSigningAlg(v []string)`

SetRequestObjectSigningAlg sets RequestObjectSigningAlg field to given value.

### HasRequestObjectSigningAlg

`func (o *Client) HasRequestObjectSigningAlg() bool`

HasRequestObjectSigningAlg returns a boolean if a field has been set.

### GetResponseTypes

`func (o *Client) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *Client) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *Client) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *Client) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *Client) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *Client) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *Client) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *Client) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.

### GetTosUri

`func (o *Client) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *Client) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *Client) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *Client) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### SetTosUriNil

`func (o *Client) SetTosUriNil(b bool)`

 SetTosUriNil sets the value for TosUri to be an explicit nil

### UnsetTosUri
`func (o *Client) UnsetTosUri()`

UnsetTosUri ensures that no value is present for TosUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


