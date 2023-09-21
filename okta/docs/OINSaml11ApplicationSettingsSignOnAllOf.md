# OINSaml11ApplicationSettingsSignOnAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to SAML_1_1]
**DefaultRelayState** | Pointer to **string** | Identifies a specific application resource in an IDP-initiated SSO scenario | [optional] 
**SsoAcsUrlOverride** | Pointer to **string** | Assertion Consumer Service URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm) | [optional] 
**AudienceOverride** | Pointer to **string** | Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm) | [optional] 
**RecipientOverride** | Pointer to **string** | Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm) | [optional] 

## Methods

### NewOINSaml11ApplicationSettingsSignOnAllOf

`func NewOINSaml11ApplicationSettingsSignOnAllOf() *OINSaml11ApplicationSettingsSignOnAllOf`

NewOINSaml11ApplicationSettingsSignOnAllOf instantiates a new OINSaml11ApplicationSettingsSignOnAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOINSaml11ApplicationSettingsSignOnAllOfWithDefaults

`func NewOINSaml11ApplicationSettingsSignOnAllOfWithDefaults() *OINSaml11ApplicationSettingsSignOnAllOf`

NewOINSaml11ApplicationSettingsSignOnAllOfWithDefaults instantiates a new OINSaml11ApplicationSettingsSignOnAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignOnMode

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *OINSaml11ApplicationSettingsSignOnAllOf) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetDefaultRelayState

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetDefaultRelayState() string`

GetDefaultRelayState returns the DefaultRelayState field if non-nil, zero value otherwise.

### GetDefaultRelayStateOk

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetDefaultRelayStateOk() (*string, bool)`

GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRelayState

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) SetDefaultRelayState(v string)`

SetDefaultRelayState sets DefaultRelayState field to given value.

### HasDefaultRelayState

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) HasDefaultRelayState() bool`

HasDefaultRelayState returns a boolean if a field has been set.

### GetSsoAcsUrlOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetSsoAcsUrlOverride() string`

GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field if non-nil, zero value otherwise.

### GetSsoAcsUrlOverrideOk

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetSsoAcsUrlOverrideOk() (*string, bool)`

GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAcsUrlOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) SetSsoAcsUrlOverride(v string)`

SetSsoAcsUrlOverride sets SsoAcsUrlOverride field to given value.

### HasSsoAcsUrlOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) HasSsoAcsUrlOverride() bool`

HasSsoAcsUrlOverride returns a boolean if a field has been set.

### GetAudienceOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetAudienceOverride() string`

GetAudienceOverride returns the AudienceOverride field if non-nil, zero value otherwise.

### GetAudienceOverrideOk

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetAudienceOverrideOk() (*string, bool)`

GetAudienceOverrideOk returns a tuple with the AudienceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) SetAudienceOverride(v string)`

SetAudienceOverride sets AudienceOverride field to given value.

### HasAudienceOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) HasAudienceOverride() bool`

HasAudienceOverride returns a boolean if a field has been set.

### GetRecipientOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetRecipientOverride() string`

GetRecipientOverride returns the RecipientOverride field if non-nil, zero value otherwise.

### GetRecipientOverrideOk

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) GetRecipientOverrideOk() (*string, bool)`

GetRecipientOverrideOk returns a tuple with the RecipientOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) SetRecipientOverride(v string)`

SetRecipientOverride sets RecipientOverride field to given value.

### HasRecipientOverride

`func (o *OINSaml11ApplicationSettingsSignOnAllOf) HasRecipientOverride() bool`

HasRecipientOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


