# Saml11ApplicationSettingsSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceOverride** | Pointer to **string** | The intended audience of the SAML assertion. This is usually the Entity ID of your application. | [optional] 
**DefaultRelayState** | Pointer to **string** | The URL of the resource to direct users after they successfully sign in to the SP using SAML. See the SP documentation to check if you need to specify a RelayState. In most instances, you can leave this field blank. | [optional] 
**RecipientOverride** | Pointer to **string** | The location where the application can present the SAML assertion. This is usually the Single Sign-On (SSO) URL. | [optional] 
**SsoAcsUrlOverride** | Pointer to **string** | Assertion Consumer Services (ACS) URL value for the Service Provider (SP). This URL is always used for Identity Provider (IdP) initiated sign-on requests. | [optional] 

## Methods

### NewSaml11ApplicationSettingsSignOn

`func NewSaml11ApplicationSettingsSignOn() *Saml11ApplicationSettingsSignOn`

NewSaml11ApplicationSettingsSignOn instantiates a new Saml11ApplicationSettingsSignOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaml11ApplicationSettingsSignOnWithDefaults

`func NewSaml11ApplicationSettingsSignOnWithDefaults() *Saml11ApplicationSettingsSignOn`

NewSaml11ApplicationSettingsSignOnWithDefaults instantiates a new Saml11ApplicationSettingsSignOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceOverride

`func (o *Saml11ApplicationSettingsSignOn) GetAudienceOverride() string`

GetAudienceOverride returns the AudienceOverride field if non-nil, zero value otherwise.

### GetAudienceOverrideOk

`func (o *Saml11ApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool)`

GetAudienceOverrideOk returns a tuple with the AudienceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceOverride

`func (o *Saml11ApplicationSettingsSignOn) SetAudienceOverride(v string)`

SetAudienceOverride sets AudienceOverride field to given value.

### HasAudienceOverride

`func (o *Saml11ApplicationSettingsSignOn) HasAudienceOverride() bool`

HasAudienceOverride returns a boolean if a field has been set.

### GetDefaultRelayState

`func (o *Saml11ApplicationSettingsSignOn) GetDefaultRelayState() string`

GetDefaultRelayState returns the DefaultRelayState field if non-nil, zero value otherwise.

### GetDefaultRelayStateOk

`func (o *Saml11ApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool)`

GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRelayState

`func (o *Saml11ApplicationSettingsSignOn) SetDefaultRelayState(v string)`

SetDefaultRelayState sets DefaultRelayState field to given value.

### HasDefaultRelayState

`func (o *Saml11ApplicationSettingsSignOn) HasDefaultRelayState() bool`

HasDefaultRelayState returns a boolean if a field has been set.

### GetRecipientOverride

`func (o *Saml11ApplicationSettingsSignOn) GetRecipientOverride() string`

GetRecipientOverride returns the RecipientOverride field if non-nil, zero value otherwise.

### GetRecipientOverrideOk

`func (o *Saml11ApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool)`

GetRecipientOverrideOk returns a tuple with the RecipientOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientOverride

`func (o *Saml11ApplicationSettingsSignOn) SetRecipientOverride(v string)`

SetRecipientOverride sets RecipientOverride field to given value.

### HasRecipientOverride

`func (o *Saml11ApplicationSettingsSignOn) HasRecipientOverride() bool`

HasRecipientOverride returns a boolean if a field has been set.

### GetSsoAcsUrlOverride

`func (o *Saml11ApplicationSettingsSignOn) GetSsoAcsUrlOverride() string`

GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field if non-nil, zero value otherwise.

### GetSsoAcsUrlOverrideOk

`func (o *Saml11ApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool)`

GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAcsUrlOverride

`func (o *Saml11ApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string)`

SetSsoAcsUrlOverride sets SsoAcsUrlOverride field to given value.

### HasSsoAcsUrlOverride

`func (o *Saml11ApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool`

HasSsoAcsUrlOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


