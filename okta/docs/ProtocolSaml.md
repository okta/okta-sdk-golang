# ProtocolSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | Pointer to [**SamlAlgorithms**](SamlAlgorithms.md) |  | [optional] 
**Credentials** | Pointer to [**SamlCredentials**](SamlCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**SamlEndpoints**](SamlEndpoints.md) |  | [optional] 
**RelayState** | Pointer to [**SamlRelayState**](SamlRelayState.md) |  | [optional] 
**Settings** | Pointer to [**SamlSettings**](SamlSettings.md) |  | [optional] 
**Type** | Pointer to **string** | SAML 2.0 protocol | [optional] 

## Methods

### NewProtocolSaml

`func NewProtocolSaml() *ProtocolSaml`

NewProtocolSaml instantiates a new ProtocolSaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolSamlWithDefaults

`func NewProtocolSamlWithDefaults() *ProtocolSaml`

NewProtocolSamlWithDefaults instantiates a new ProtocolSaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *ProtocolSaml) GetAlgorithms() SamlAlgorithms`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *ProtocolSaml) GetAlgorithmsOk() (*SamlAlgorithms, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *ProtocolSaml) SetAlgorithms(v SamlAlgorithms)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *ProtocolSaml) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetCredentials

`func (o *ProtocolSaml) GetCredentials() SamlCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ProtocolSaml) GetCredentialsOk() (*SamlCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ProtocolSaml) SetCredentials(v SamlCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ProtocolSaml) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *ProtocolSaml) GetEndpoints() SamlEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ProtocolSaml) GetEndpointsOk() (*SamlEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ProtocolSaml) SetEndpoints(v SamlEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ProtocolSaml) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetRelayState

`func (o *ProtocolSaml) GetRelayState() SamlRelayState`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *ProtocolSaml) GetRelayStateOk() (*SamlRelayState, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *ProtocolSaml) SetRelayState(v SamlRelayState)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *ProtocolSaml) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetSettings

`func (o *ProtocolSaml) GetSettings() SamlSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProtocolSaml) GetSettingsOk() (*SamlSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProtocolSaml) SetSettings(v SamlSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ProtocolSaml) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *ProtocolSaml) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtocolSaml) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtocolSaml) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtocolSaml) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


