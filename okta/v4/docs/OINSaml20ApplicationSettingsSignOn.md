# OINSaml20ApplicationSettingsSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceOverride** | Pointer to **string** | Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 
**ConfiguredAttributeStatements** | Pointer to [**[]SamlAttributeStatement**](SamlAttributeStatement.md) |  | [optional] 
**DefaultRelayState** | Pointer to **string** | Identifies a specific application resource in an IdP-initiated SSO scenario | [optional] 
**DestinationOverride** | Pointer to **string** | Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 
**RecipientOverride** | Pointer to **string** | Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 
**SamlAssertionLifetimeSeconds** | Pointer to **int32** | Determines the SAML app session lifetimes with Okta | [optional] 
**SsoAcsUrlOverride** | Pointer to **string** | Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 

## Methods

### NewOINSaml20ApplicationSettingsSignOn

`func NewOINSaml20ApplicationSettingsSignOn() *OINSaml20ApplicationSettingsSignOn`

NewOINSaml20ApplicationSettingsSignOn instantiates a new OINSaml20ApplicationSettingsSignOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOINSaml20ApplicationSettingsSignOnWithDefaults

`func NewOINSaml20ApplicationSettingsSignOnWithDefaults() *OINSaml20ApplicationSettingsSignOn`

NewOINSaml20ApplicationSettingsSignOnWithDefaults instantiates a new OINSaml20ApplicationSettingsSignOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceOverride

`func (o *OINSaml20ApplicationSettingsSignOn) GetAudienceOverride() string`

GetAudienceOverride returns the AudienceOverride field if non-nil, zero value otherwise.

### GetAudienceOverrideOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool)`

GetAudienceOverrideOk returns a tuple with the AudienceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceOverride

`func (o *OINSaml20ApplicationSettingsSignOn) SetAudienceOverride(v string)`

SetAudienceOverride sets AudienceOverride field to given value.

### HasAudienceOverride

`func (o *OINSaml20ApplicationSettingsSignOn) HasAudienceOverride() bool`

HasAudienceOverride returns a boolean if a field has been set.

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

### GetDefaultRelayState

`func (o *OINSaml20ApplicationSettingsSignOn) GetDefaultRelayState() string`

GetDefaultRelayState returns the DefaultRelayState field if non-nil, zero value otherwise.

### GetDefaultRelayStateOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool)`

GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRelayState

`func (o *OINSaml20ApplicationSettingsSignOn) SetDefaultRelayState(v string)`

SetDefaultRelayState sets DefaultRelayState field to given value.

### HasDefaultRelayState

`func (o *OINSaml20ApplicationSettingsSignOn) HasDefaultRelayState() bool`

HasDefaultRelayState returns a boolean if a field has been set.

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

### HasDestinationOverride

`func (o *OINSaml20ApplicationSettingsSignOn) HasDestinationOverride() bool`

HasDestinationOverride returns a boolean if a field has been set.

### GetRecipientOverride

`func (o *OINSaml20ApplicationSettingsSignOn) GetRecipientOverride() string`

GetRecipientOverride returns the RecipientOverride field if non-nil, zero value otherwise.

### GetRecipientOverrideOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool)`

GetRecipientOverrideOk returns a tuple with the RecipientOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientOverride

`func (o *OINSaml20ApplicationSettingsSignOn) SetRecipientOverride(v string)`

SetRecipientOverride sets RecipientOverride field to given value.

### HasRecipientOverride

`func (o *OINSaml20ApplicationSettingsSignOn) HasRecipientOverride() bool`

HasRecipientOverride returns a boolean if a field has been set.

### GetSamlAssertionLifetimeSeconds

`func (o *OINSaml20ApplicationSettingsSignOn) GetSamlAssertionLifetimeSeconds() int32`

GetSamlAssertionLifetimeSeconds returns the SamlAssertionLifetimeSeconds field if non-nil, zero value otherwise.

### GetSamlAssertionLifetimeSecondsOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetSamlAssertionLifetimeSecondsOk() (*int32, bool)`

GetSamlAssertionLifetimeSecondsOk returns a tuple with the SamlAssertionLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAssertionLifetimeSeconds

`func (o *OINSaml20ApplicationSettingsSignOn) SetSamlAssertionLifetimeSeconds(v int32)`

SetSamlAssertionLifetimeSeconds sets SamlAssertionLifetimeSeconds field to given value.

### HasSamlAssertionLifetimeSeconds

`func (o *OINSaml20ApplicationSettingsSignOn) HasSamlAssertionLifetimeSeconds() bool`

HasSamlAssertionLifetimeSeconds returns a boolean if a field has been set.

### GetSsoAcsUrlOverride

`func (o *OINSaml20ApplicationSettingsSignOn) GetSsoAcsUrlOverride() string`

GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field if non-nil, zero value otherwise.

### GetSsoAcsUrlOverrideOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool)`

GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAcsUrlOverride

`func (o *OINSaml20ApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string)`

SetSsoAcsUrlOverride sets SsoAcsUrlOverride field to given value.

### HasSsoAcsUrlOverride

`func (o *OINSaml20ApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool`

HasSsoAcsUrlOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


