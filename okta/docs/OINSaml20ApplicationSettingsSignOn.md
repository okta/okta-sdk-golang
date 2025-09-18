# OINSaml20ApplicationSettingsSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeStatements** | Pointer to [**[]SamlAttributeStatement**](SamlAttributeStatement.md) | A list of custom attribute statements for the app&#39;s SAML assertion. See [SAML 2.0 Technical Overview](https://docs.oasis-open.org/security/saml/Post2.0/sstc-saml-tech-overview-2.0-cd-02.html).  There are two types of attribute statements: | Type | Description | | ---- | ----------- | | EXPRESSION | Generic attribute statement that can be dynamic and supports [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/) | | GROUP | Group attribute statement |  | [optional] 
**AudienceOverride** | Pointer to **NullableString** | Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 
**DefaultRelayState** | Pointer to **NullableString** | Identifies a specific application resource in an IdP-initiated SSO scenario | [optional] 
**DestinationOverride** | Pointer to **NullableString** | Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 
**RecipientOverride** | Pointer to **NullableString** | Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 
**SamlAssertionLifetimeSeconds** | Pointer to **int32** | Determines the SAML app session lifetimes with Okta | [optional] 
**SsoAcsUrlOverride** | Pointer to **NullableString** | Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm). | [optional] 

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

### GetAttributeStatements

`func (o *OINSaml20ApplicationSettingsSignOn) GetAttributeStatements() []SamlAttributeStatement`

GetAttributeStatements returns the AttributeStatements field if non-nil, zero value otherwise.

### GetAttributeStatementsOk

`func (o *OINSaml20ApplicationSettingsSignOn) GetAttributeStatementsOk() (*[]SamlAttributeStatement, bool)`

GetAttributeStatementsOk returns a tuple with the AttributeStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeStatements

`func (o *OINSaml20ApplicationSettingsSignOn) SetAttributeStatements(v []SamlAttributeStatement)`

SetAttributeStatements sets AttributeStatements field to given value.

### HasAttributeStatements

`func (o *OINSaml20ApplicationSettingsSignOn) HasAttributeStatements() bool`

HasAttributeStatements returns a boolean if a field has been set.

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

### SetAudienceOverrideNil

`func (o *OINSaml20ApplicationSettingsSignOn) SetAudienceOverrideNil(b bool)`

 SetAudienceOverrideNil sets the value for AudienceOverride to be an explicit nil

### UnsetAudienceOverride
`func (o *OINSaml20ApplicationSettingsSignOn) UnsetAudienceOverride()`

UnsetAudienceOverride ensures that no value is present for AudienceOverride, not even an explicit nil
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

### SetDefaultRelayStateNil

`func (o *OINSaml20ApplicationSettingsSignOn) SetDefaultRelayStateNil(b bool)`

 SetDefaultRelayStateNil sets the value for DefaultRelayState to be an explicit nil

### UnsetDefaultRelayState
`func (o *OINSaml20ApplicationSettingsSignOn) UnsetDefaultRelayState()`

UnsetDefaultRelayState ensures that no value is present for DefaultRelayState, not even an explicit nil
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

### SetDestinationOverrideNil

`func (o *OINSaml20ApplicationSettingsSignOn) SetDestinationOverrideNil(b bool)`

 SetDestinationOverrideNil sets the value for DestinationOverride to be an explicit nil

### UnsetDestinationOverride
`func (o *OINSaml20ApplicationSettingsSignOn) UnsetDestinationOverride()`

UnsetDestinationOverride ensures that no value is present for DestinationOverride, not even an explicit nil
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

### SetRecipientOverrideNil

`func (o *OINSaml20ApplicationSettingsSignOn) SetRecipientOverrideNil(b bool)`

 SetRecipientOverrideNil sets the value for RecipientOverride to be an explicit nil

### UnsetRecipientOverride
`func (o *OINSaml20ApplicationSettingsSignOn) UnsetRecipientOverride()`

UnsetRecipientOverride ensures that no value is present for RecipientOverride, not even an explicit nil
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

### SetSsoAcsUrlOverrideNil

`func (o *OINSaml20ApplicationSettingsSignOn) SetSsoAcsUrlOverrideNil(b bool)`

 SetSsoAcsUrlOverrideNil sets the value for SsoAcsUrlOverride to be an explicit nil

### UnsetSsoAcsUrlOverride
`func (o *OINSaml20ApplicationSettingsSignOn) UnsetSsoAcsUrlOverride()`

UnsetSsoAcsUrlOverride ensures that no value is present for SsoAcsUrlOverride, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


