# Protocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | Pointer to [**ProtocolAlgorithms**](ProtocolAlgorithms.md) |  | [optional] 
**Credentials** | Pointer to [**IdentityProviderCredentials**](IdentityProviderCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**ProtocolEndpoints**](ProtocolEndpoints.md) |  | [optional] 
**Issuer** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**RelayState** | Pointer to [**ProtocolRelayState**](ProtocolRelayState.md) |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Settings** | Pointer to [**ProtocolSettings**](ProtocolSettings.md) |  | [optional] 
**Type** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] 

## Methods

### NewProtocol

`func NewProtocol() *Protocol`

NewProtocol instantiates a new Protocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolWithDefaults

`func NewProtocolWithDefaults() *Protocol`

NewProtocolWithDefaults instantiates a new Protocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *Protocol) GetAlgorithms() ProtocolAlgorithms`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *Protocol) GetAlgorithmsOk() (*ProtocolAlgorithms, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *Protocol) SetAlgorithms(v ProtocolAlgorithms)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *Protocol) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetCredentials

`func (o *Protocol) GetCredentials() IdentityProviderCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Protocol) GetCredentialsOk() (*IdentityProviderCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Protocol) SetCredentials(v IdentityProviderCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Protocol) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *Protocol) GetEndpoints() ProtocolEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *Protocol) GetEndpointsOk() (*ProtocolEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *Protocol) SetEndpoints(v ProtocolEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *Protocol) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetIssuer

`func (o *Protocol) GetIssuer() ProtocolEndpoint`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Protocol) GetIssuerOk() (*ProtocolEndpoint, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Protocol) SetIssuer(v ProtocolEndpoint)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Protocol) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetRelayState

`func (o *Protocol) GetRelayState() ProtocolRelayState`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *Protocol) GetRelayStateOk() (*ProtocolRelayState, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *Protocol) SetRelayState(v ProtocolRelayState)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *Protocol) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetScopes

`func (o *Protocol) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Protocol) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Protocol) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Protocol) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSettings

`func (o *Protocol) GetSettings() ProtocolSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Protocol) GetSettingsOk() (*ProtocolSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Protocol) SetSettings(v ProtocolSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Protocol) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *Protocol) GetType() ProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Protocol) GetTypeOk() (*ProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Protocol) SetType(v ProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *Protocol) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


