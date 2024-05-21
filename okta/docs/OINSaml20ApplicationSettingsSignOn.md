# OINSaml20ApplicationSettingsSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to SAML_2_0]
**DestinationOverride** | **string** | Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm) | 
**HonorForceAuthn** | Pointer to **bool** | Set to &#x60;true&#x60; to prompt users for their credentials when a SAML request has the &#x60;ForceAuthn&#x60; attribute set to &#x60;true&#x60; | [optional] 
**ConfiguredAttributeStatements** | Pointer to [**[]SamlAttributeStatement**](SamlAttributeStatement.md) |  | [optional] 

## Methods

### NewOINSaml20ApplicationSettingsSignOn

`func NewOINSaml20ApplicationSettingsSignOn(destinationOverride string, ) *OINSaml20ApplicationSettingsSignOn`

NewOINSaml20ApplicationSettingsSignOn instantiates a new OINSaml20ApplicationSettingsSignOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOINSaml20ApplicationSettingsSignOnWithDefaults

`func NewOINSaml20ApplicationSettingsSignOnWithDefaults() *OINSaml20ApplicationSettingsSignOn`

NewOINSaml20ApplicationSettingsSignOnWithDefaults instantiates a new OINSaml20ApplicationSettingsSignOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignOnMode

`func (o *OINSaml20ApplicationSettingsSignOn) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *OINSaml20ApplicationSettingsSignOn) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *OINSaml20ApplicationSettingsSignOn) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *OINSaml20ApplicationSettingsSignOn) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *OINSaml20ApplicationSettingsSignOn) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetDestinationOverride

`func (o *OINSaml20ApplicationSettingsSignOn) GetDestinationOverride() string`

GetDestinationOverride returns the DestinationOverride field if non-nil, zero value otherwise.

### GetDestinationOverrideOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetDestinationOverrideOk() (*string, bool)`

GetDestinationOverrideOk returns a tuple with the DestinationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOverride

`func (o *OINSaml20ApplicationSettingsSignOn) SetDestinationOverride(v string)`

SetDestinationOverride sets DestinationOverride field to given value.


### GetHonorForceAuthn

`func (o *OINSaml20ApplicationSettingsSignOn) GetHonorForceAuthn() bool`

GetHonorForceAuthn returns the HonorForceAuthn field if non-nil, zero value otherwise.

### GetHonorForceAuthnOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetHonorForceAuthnOk() (*bool, bool)`

GetHonorForceAuthnOk returns a tuple with the HonorForceAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorForceAuthn

`func (o *OINSaml20ApplicationSettingsSignOn) SetHonorForceAuthn(v bool)`

SetHonorForceAuthn sets HonorForceAuthn field to given value.

### HasHonorForceAuthn

`func (o *OINSaml20ApplicationSettingsSignOn) HasHonorForceAuthn() bool`

HasHonorForceAuthn returns a boolean if a field has been set.

### GetConfiguredAttributeStatements

`func (o *OINSaml20ApplicationSettingsSignOn) GetConfiguredAttributeStatements() []SamlAttributeStatement`

GetConfiguredAttributeStatements returns the ConfiguredAttributeStatements field if non-nil, zero value otherwise.

### GetConfiguredAttributeStatementsOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetConfiguredAttributeStatementsOk() (*[]SamlAttributeStatement, bool)`

GetConfiguredAttributeStatementsOk returns a tuple with the ConfiguredAttributeStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredAttributeStatements

`func (o *OINSaml20ApplicationSettingsSignOn) SetConfiguredAttributeStatements(v []SamlAttributeStatement)`

SetConfiguredAttributeStatements sets ConfiguredAttributeStatements field to given value.

### HasConfiguredAttributeStatements

`func (o *OINSaml20ApplicationSettingsSignOn) HasConfiguredAttributeStatements() bool`

HasConfiguredAttributeStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


