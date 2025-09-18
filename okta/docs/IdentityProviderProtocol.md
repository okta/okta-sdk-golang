# IdentityProviderProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | Pointer to [**OidcAlgorithms**](OidcAlgorithms.md) |  | [optional] 
**Credentials** | Pointer to [**IDVCredentials**](IDVCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**IDVEndpoints**](IDVEndpoints.md) |  | [optional] 
**RelayState** | Pointer to [**SamlRelayState**](SamlRelayState.md) |  | [optional] 
**Settings** | Pointer to [**OidcSettings**](OidcSettings.md) |  | [optional] 
**Type** | Pointer to **string** | SAML 2.0 protocol | [optional] 
**Scopes** | Pointer to **[]string** | IdP-defined permission bundles to request delegated access from the user. &gt; **Note:** The [identity provider type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;type&amp;t&#x3D;request) table lists the scopes that are supported for each IdP. | [optional] 
**OktaIdpOrgUrl** | Pointer to **string** | URL of the IdP org | [optional] 

## Methods

### NewIdentityProviderProtocol

`func NewIdentityProviderProtocol() *IdentityProviderProtocol`

NewIdentityProviderProtocol instantiates a new IdentityProviderProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderProtocolWithDefaults

`func NewIdentityProviderProtocolWithDefaults() *IdentityProviderProtocol`

NewIdentityProviderProtocolWithDefaults instantiates a new IdentityProviderProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *IdentityProviderProtocol) GetAlgorithms() OidcAlgorithms`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *IdentityProviderProtocol) GetAlgorithmsOk() (*OidcAlgorithms, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *IdentityProviderProtocol) SetAlgorithms(v OidcAlgorithms)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *IdentityProviderProtocol) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetCredentials

`func (o *IdentityProviderProtocol) GetCredentials() IDVCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IdentityProviderProtocol) GetCredentialsOk() (*IDVCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IdentityProviderProtocol) SetCredentials(v IDVCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *IdentityProviderProtocol) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *IdentityProviderProtocol) GetEndpoints() IDVEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *IdentityProviderProtocol) GetEndpointsOk() (*IDVEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *IdentityProviderProtocol) SetEndpoints(v IDVEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *IdentityProviderProtocol) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetRelayState

`func (o *IdentityProviderProtocol) GetRelayState() SamlRelayState`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *IdentityProviderProtocol) GetRelayStateOk() (*SamlRelayState, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *IdentityProviderProtocol) SetRelayState(v SamlRelayState)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *IdentityProviderProtocol) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetSettings

`func (o *IdentityProviderProtocol) GetSettings() OidcSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *IdentityProviderProtocol) GetSettingsOk() (*OidcSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *IdentityProviderProtocol) SetSettings(v OidcSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *IdentityProviderProtocol) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *IdentityProviderProtocol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProviderProtocol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProviderProtocol) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityProviderProtocol) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScopes

`func (o *IdentityProviderProtocol) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityProviderProtocol) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityProviderProtocol) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IdentityProviderProtocol) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetOktaIdpOrgUrl

`func (o *IdentityProviderProtocol) GetOktaIdpOrgUrl() string`

GetOktaIdpOrgUrl returns the OktaIdpOrgUrl field if non-nil, zero value otherwise.

### GetOktaIdpOrgUrlOk

`func (o *IdentityProviderProtocol) GetOktaIdpOrgUrlOk() (*string, bool)`

GetOktaIdpOrgUrlOk returns a tuple with the OktaIdpOrgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaIdpOrgUrl

`func (o *IdentityProviderProtocol) SetOktaIdpOrgUrl(v string)`

SetOktaIdpOrgUrl sets OktaIdpOrgUrl field to given value.

### HasOktaIdpOrgUrl

`func (o *IdentityProviderProtocol) HasOktaIdpOrgUrl() bool`

HasOktaIdpOrgUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


