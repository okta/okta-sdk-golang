# WsFederationApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeStatements** | Pointer to **string** |  | [optional] 
**AudienceRestriction** | Pointer to **string** |  | [optional] 
**AuthnContextClassRef** | Pointer to **string** |  | [optional] 
**GroupFilter** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**GroupValueFormat** | Pointer to **string** |  | [optional] 
**NameIDFormat** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**SiteURL** | Pointer to **string** |  | [optional] 
**UsernameAttribute** | Pointer to **string** |  | [optional] 
**WReplyOverride** | Pointer to **bool** |  | [optional] 
**WReplyURL** | Pointer to **string** |  | [optional] 

## Methods

### NewWsFederationApplicationSettingsApplication

`func NewWsFederationApplicationSettingsApplication() *WsFederationApplicationSettingsApplication`

NewWsFederationApplicationSettingsApplication instantiates a new WsFederationApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsFederationApplicationSettingsApplicationWithDefaults

`func NewWsFederationApplicationSettingsApplicationWithDefaults() *WsFederationApplicationSettingsApplication`

NewWsFederationApplicationSettingsApplicationWithDefaults instantiates a new WsFederationApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeStatements

`func (o *WsFederationApplicationSettingsApplication) GetAttributeStatements() string`

GetAttributeStatements returns the AttributeStatements field if non-nil, zero value otherwise.

### GetAttributeStatementsOk

`func (o *WsFederationApplicationSettingsApplication) GetAttributeStatementsOk() (*string, bool)`

GetAttributeStatementsOk returns a tuple with the AttributeStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeStatements

`func (o *WsFederationApplicationSettingsApplication) SetAttributeStatements(v string)`

SetAttributeStatements sets AttributeStatements field to given value.

### HasAttributeStatements

`func (o *WsFederationApplicationSettingsApplication) HasAttributeStatements() bool`

HasAttributeStatements returns a boolean if a field has been set.

### GetAudienceRestriction

`func (o *WsFederationApplicationSettingsApplication) GetAudienceRestriction() string`

GetAudienceRestriction returns the AudienceRestriction field if non-nil, zero value otherwise.

### GetAudienceRestrictionOk

`func (o *WsFederationApplicationSettingsApplication) GetAudienceRestrictionOk() (*string, bool)`

GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceRestriction

`func (o *WsFederationApplicationSettingsApplication) SetAudienceRestriction(v string)`

SetAudienceRestriction sets AudienceRestriction field to given value.

### HasAudienceRestriction

`func (o *WsFederationApplicationSettingsApplication) HasAudienceRestriction() bool`

HasAudienceRestriction returns a boolean if a field has been set.

### GetAuthnContextClassRef

`func (o *WsFederationApplicationSettingsApplication) GetAuthnContextClassRef() string`

GetAuthnContextClassRef returns the AuthnContextClassRef field if non-nil, zero value otherwise.

### GetAuthnContextClassRefOk

`func (o *WsFederationApplicationSettingsApplication) GetAuthnContextClassRefOk() (*string, bool)`

GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextClassRef

`func (o *WsFederationApplicationSettingsApplication) SetAuthnContextClassRef(v string)`

SetAuthnContextClassRef sets AuthnContextClassRef field to given value.

### HasAuthnContextClassRef

`func (o *WsFederationApplicationSettingsApplication) HasAuthnContextClassRef() bool`

HasAuthnContextClassRef returns a boolean if a field has been set.

### GetGroupFilter

`func (o *WsFederationApplicationSettingsApplication) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *WsFederationApplicationSettingsApplication) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *WsFederationApplicationSettingsApplication) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *WsFederationApplicationSettingsApplication) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

### GetGroupName

`func (o *WsFederationApplicationSettingsApplication) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *WsFederationApplicationSettingsApplication) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *WsFederationApplicationSettingsApplication) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *WsFederationApplicationSettingsApplication) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupValueFormat

`func (o *WsFederationApplicationSettingsApplication) GetGroupValueFormat() string`

GetGroupValueFormat returns the GroupValueFormat field if non-nil, zero value otherwise.

### GetGroupValueFormatOk

`func (o *WsFederationApplicationSettingsApplication) GetGroupValueFormatOk() (*string, bool)`

GetGroupValueFormatOk returns a tuple with the GroupValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupValueFormat

`func (o *WsFederationApplicationSettingsApplication) SetGroupValueFormat(v string)`

SetGroupValueFormat sets GroupValueFormat field to given value.

### HasGroupValueFormat

`func (o *WsFederationApplicationSettingsApplication) HasGroupValueFormat() bool`

HasGroupValueFormat returns a boolean if a field has been set.

### GetNameIDFormat

`func (o *WsFederationApplicationSettingsApplication) GetNameIDFormat() string`

GetNameIDFormat returns the NameIDFormat field if non-nil, zero value otherwise.

### GetNameIDFormatOk

`func (o *WsFederationApplicationSettingsApplication) GetNameIDFormatOk() (*string, bool)`

GetNameIDFormatOk returns a tuple with the NameIDFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIDFormat

`func (o *WsFederationApplicationSettingsApplication) SetNameIDFormat(v string)`

SetNameIDFormat sets NameIDFormat field to given value.

### HasNameIDFormat

`func (o *WsFederationApplicationSettingsApplication) HasNameIDFormat() bool`

HasNameIDFormat returns a boolean if a field has been set.

### GetRealm

`func (o *WsFederationApplicationSettingsApplication) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *WsFederationApplicationSettingsApplication) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *WsFederationApplicationSettingsApplication) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *WsFederationApplicationSettingsApplication) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetSiteURL

`func (o *WsFederationApplicationSettingsApplication) GetSiteURL() string`

GetSiteURL returns the SiteURL field if non-nil, zero value otherwise.

### GetSiteURLOk

`func (o *WsFederationApplicationSettingsApplication) GetSiteURLOk() (*string, bool)`

GetSiteURLOk returns a tuple with the SiteURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteURL

`func (o *WsFederationApplicationSettingsApplication) SetSiteURL(v string)`

SetSiteURL sets SiteURL field to given value.

### HasSiteURL

`func (o *WsFederationApplicationSettingsApplication) HasSiteURL() bool`

HasSiteURL returns a boolean if a field has been set.

### GetUsernameAttribute

`func (o *WsFederationApplicationSettingsApplication) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *WsFederationApplicationSettingsApplication) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *WsFederationApplicationSettingsApplication) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.

### HasUsernameAttribute

`func (o *WsFederationApplicationSettingsApplication) HasUsernameAttribute() bool`

HasUsernameAttribute returns a boolean if a field has been set.

### GetWReplyOverride

`func (o *WsFederationApplicationSettingsApplication) GetWReplyOverride() bool`

GetWReplyOverride returns the WReplyOverride field if non-nil, zero value otherwise.

### GetWReplyOverrideOk

`func (o *WsFederationApplicationSettingsApplication) GetWReplyOverrideOk() (*bool, bool)`

GetWReplyOverrideOk returns a tuple with the WReplyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWReplyOverride

`func (o *WsFederationApplicationSettingsApplication) SetWReplyOverride(v bool)`

SetWReplyOverride sets WReplyOverride field to given value.

### HasWReplyOverride

`func (o *WsFederationApplicationSettingsApplication) HasWReplyOverride() bool`

HasWReplyOverride returns a boolean if a field has been set.

### GetWReplyURL

`func (o *WsFederationApplicationSettingsApplication) GetWReplyURL() string`

GetWReplyURL returns the WReplyURL field if non-nil, zero value otherwise.

### GetWReplyURLOk

`func (o *WsFederationApplicationSettingsApplication) GetWReplyURLOk() (*string, bool)`

GetWReplyURLOk returns a tuple with the WReplyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWReplyURL

`func (o *WsFederationApplicationSettingsApplication) SetWReplyURL(v string)`

SetWReplyURL sets WReplyURL field to given value.

### HasWReplyURL

`func (o *WsFederationApplicationSettingsApplication) HasWReplyURL() bool`

HasWReplyURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


