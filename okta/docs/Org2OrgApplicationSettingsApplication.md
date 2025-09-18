# Org2OrgApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsUrl** | Pointer to **string** | The Assertion Consumer Service (ACS) URL of the source org (for &#x60;SAML_2_0&#x60; sign-on mode) | [optional] 
**AudRestriction** | Pointer to **string** | The entity ID of the SP (for &#x60;SAML_2_0&#x60; sign-on mode) | [optional] 
**BaseUrl** | **string** | The base URL of the target Okta org (for &#x60;SAML_2_0&#x60; sign-on mode) | 
**CreationState** | Pointer to **string** | Used to track and manage the state of the app&#39;s creation or the provisioning process between two Okta orgs | [optional] 
**PreferUsernameOverEmail** | Pointer to **bool** | Indicates that you don&#39;t want to use an email address as the username | [optional] 
**Token** | Pointer to **string** | An API token from the target org that&#39;s used to secure the connection between the orgs | [optional] 
**TokenEncrypted** | Pointer to **string** | Encrypted token to enhance security | [optional] 

## Methods

### NewOrg2OrgApplicationSettingsApplication

`func NewOrg2OrgApplicationSettingsApplication(baseUrl string, ) *Org2OrgApplicationSettingsApplication`

NewOrg2OrgApplicationSettingsApplication instantiates a new Org2OrgApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrg2OrgApplicationSettingsApplicationWithDefaults

`func NewOrg2OrgApplicationSettingsApplicationWithDefaults() *Org2OrgApplicationSettingsApplication`

NewOrg2OrgApplicationSettingsApplicationWithDefaults instantiates a new Org2OrgApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsUrl

`func (o *Org2OrgApplicationSettingsApplication) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *Org2OrgApplicationSettingsApplication) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *Org2OrgApplicationSettingsApplication) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *Org2OrgApplicationSettingsApplication) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetAudRestriction

`func (o *Org2OrgApplicationSettingsApplication) GetAudRestriction() string`

GetAudRestriction returns the AudRestriction field if non-nil, zero value otherwise.

### GetAudRestrictionOk

`func (o *Org2OrgApplicationSettingsApplication) GetAudRestrictionOk() (*string, bool)`

GetAudRestrictionOk returns a tuple with the AudRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudRestriction

`func (o *Org2OrgApplicationSettingsApplication) SetAudRestriction(v string)`

SetAudRestriction sets AudRestriction field to given value.

### HasAudRestriction

`func (o *Org2OrgApplicationSettingsApplication) HasAudRestriction() bool`

HasAudRestriction returns a boolean if a field has been set.

### GetBaseUrl

`func (o *Org2OrgApplicationSettingsApplication) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *Org2OrgApplicationSettingsApplication) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *Org2OrgApplicationSettingsApplication) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetCreationState

`func (o *Org2OrgApplicationSettingsApplication) GetCreationState() string`

GetCreationState returns the CreationState field if non-nil, zero value otherwise.

### GetCreationStateOk

`func (o *Org2OrgApplicationSettingsApplication) GetCreationStateOk() (*string, bool)`

GetCreationStateOk returns a tuple with the CreationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationState

`func (o *Org2OrgApplicationSettingsApplication) SetCreationState(v string)`

SetCreationState sets CreationState field to given value.

### HasCreationState

`func (o *Org2OrgApplicationSettingsApplication) HasCreationState() bool`

HasCreationState returns a boolean if a field has been set.

### GetPreferUsernameOverEmail

`func (o *Org2OrgApplicationSettingsApplication) GetPreferUsernameOverEmail() bool`

GetPreferUsernameOverEmail returns the PreferUsernameOverEmail field if non-nil, zero value otherwise.

### GetPreferUsernameOverEmailOk

`func (o *Org2OrgApplicationSettingsApplication) GetPreferUsernameOverEmailOk() (*bool, bool)`

GetPreferUsernameOverEmailOk returns a tuple with the PreferUsernameOverEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferUsernameOverEmail

`func (o *Org2OrgApplicationSettingsApplication) SetPreferUsernameOverEmail(v bool)`

SetPreferUsernameOverEmail sets PreferUsernameOverEmail field to given value.

### HasPreferUsernameOverEmail

`func (o *Org2OrgApplicationSettingsApplication) HasPreferUsernameOverEmail() bool`

HasPreferUsernameOverEmail returns a boolean if a field has been set.

### GetToken

`func (o *Org2OrgApplicationSettingsApplication) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Org2OrgApplicationSettingsApplication) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Org2OrgApplicationSettingsApplication) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Org2OrgApplicationSettingsApplication) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenEncrypted

`func (o *Org2OrgApplicationSettingsApplication) GetTokenEncrypted() string`

GetTokenEncrypted returns the TokenEncrypted field if non-nil, zero value otherwise.

### GetTokenEncryptedOk

`func (o *Org2OrgApplicationSettingsApplication) GetTokenEncryptedOk() (*string, bool)`

GetTokenEncryptedOk returns a tuple with the TokenEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEncrypted

`func (o *Org2OrgApplicationSettingsApplication) SetTokenEncrypted(v string)`

SetTokenEncrypted sets TokenEncrypted field to given value.

### HasTokenEncrypted

`func (o *Org2OrgApplicationSettingsApplication) HasTokenEncrypted() bool`

HasTokenEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


