# ProtocolOidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | Pointer to [**OidcAlgorithms**](OidcAlgorithms.md) |  | [optional] 
**Credentials** | Pointer to [**OAuthCredentials**](OAuthCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**OAuthEndpoints**](OAuthEndpoints.md) |  | [optional] 
**OktaIdpOrgUrl** | Pointer to **string** | URL of the IdP org | [optional] 
**Scopes** | Pointer to **[]string** | OpenID Connect and IdP-defined permission bundles to request delegated access from the user &gt; **Note:** The [IdP type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;type&amp;t&#x3D;request) table lists the scopes that are supported for each IdP. | [optional] 
**Settings** | Pointer to [**OidcSettings**](OidcSettings.md) |  | [optional] 
**Type** | Pointer to **string** | OpenID Connect Authorization Code flow | [optional] 

## Methods

### NewProtocolOidc

`func NewProtocolOidc() *ProtocolOidc`

NewProtocolOidc instantiates a new ProtocolOidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolOidcWithDefaults

`func NewProtocolOidcWithDefaults() *ProtocolOidc`

NewProtocolOidcWithDefaults instantiates a new ProtocolOidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *ProtocolOidc) GetAlgorithms() OidcAlgorithms`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *ProtocolOidc) GetAlgorithmsOk() (*OidcAlgorithms, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *ProtocolOidc) SetAlgorithms(v OidcAlgorithms)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *ProtocolOidc) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetCredentials

`func (o *ProtocolOidc) GetCredentials() OAuthCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ProtocolOidc) GetCredentialsOk() (*OAuthCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ProtocolOidc) SetCredentials(v OAuthCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ProtocolOidc) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *ProtocolOidc) GetEndpoints() OAuthEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ProtocolOidc) GetEndpointsOk() (*OAuthEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ProtocolOidc) SetEndpoints(v OAuthEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ProtocolOidc) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetOktaIdpOrgUrl

`func (o *ProtocolOidc) GetOktaIdpOrgUrl() string`

GetOktaIdpOrgUrl returns the OktaIdpOrgUrl field if non-nil, zero value otherwise.

### GetOktaIdpOrgUrlOk

`func (o *ProtocolOidc) GetOktaIdpOrgUrlOk() (*string, bool)`

GetOktaIdpOrgUrlOk returns a tuple with the OktaIdpOrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaIdpOrgUrl

`func (o *ProtocolOidc) SetOktaIdpOrgUrl(v string)`

SetOktaIdpOrgUrl sets OktaIdpOrgUrl field to given value.

### HasOktaIdpOrgUrl

`func (o *ProtocolOidc) HasOktaIdpOrgUrl() bool`

HasOktaIdpOrgUrl returns a boolean if a field has been set.

### GetScopes

`func (o *ProtocolOidc) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProtocolOidc) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProtocolOidc) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProtocolOidc) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSettings

`func (o *ProtocolOidc) GetSettings() OidcSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProtocolOidc) GetSettingsOk() (*OidcSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProtocolOidc) SetSettings(v OidcSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ProtocolOidc) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *ProtocolOidc) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtocolOidc) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtocolOidc) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtocolOidc) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


