# WsFederationApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeStatements** | Pointer to **string** | You can federate user attributes such as Okta profile fields, LDAP, Active Directory, and Workday values. The SP uses the federated WS-Fed attribute values accordingly. | [optional] 
**AudienceRestriction** | **string** | The entity ID of the SP. Use the entity ID value exactly as provided by the SP. | 
**AuthnContextClassRef** | **string** | Identifies the SAML authentication context class for the assertion&#39;s authentication statement | 
**GroupFilter** | Pointer to **string** | A regular expression that filters for the User Groups you want included with the &#x60;groupName&#x60; attribute. If the matching User Group has a corresponding AD group, then the attribute statement includes the value of the attribute specified by &#x60;groupValueFormat&#x60;. If the matching User Group doesn&#39;t contain a corresponding AD group, then the &#x60;groupName&#x60; is used in the attribute statement. | [optional] 
**GroupName** | Pointer to **string** | The group name to include in the WS-Fed response attribute statement. This property is used in conjunction with the &#x60;groupFilter&#x60; property.  Groups that are filtered through the &#x60;groupFilter&#x60; expression are included with the &#x60;groupName&#x60; in the attribute statement. Any users that belong to the group you&#39;ve filtered are included in the WS-Fed response attribute statement. | [optional] 
**GroupValueFormat** | **string** | Specifies the WS-Fed assertion attribute value for filtered groups. This attribute is only applied to Active Directory groups. | 
**NameIDFormat** | **string** | The username format that you send in the WS-Fed response | 
**Realm** | Pointer to **string** | The uniform resource identifier (URI) of the WS-Fed app that&#39;s used to share resources securely within a domain. It&#39;s the identity that&#39;s sent to the Okta IdP when signing in. See [Realm name](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;ext_Apps_Configure_Okta_Template_WS_Federation#Realm). | [optional] 
**SiteURL** | **string** | Launch URL for the web app | 
**UsernameAttribute** | **string** | Specifies additional username attribute statements to include in the WS-Fed assertion | 
**WReplyOverride** | Pointer to **bool** | Enables a web app to override the &#x60;wReplyURL&#x60; URL with a reply parameter. | [optional] 
**WReplyURL** | **string** | The WS-Fed SP endpoint where your users sign in | 

## Methods

### NewWsFederationApplicationSettingsApplication

`func NewWsFederationApplicationSettingsApplication(audienceRestriction string, authnContextClassRef string, groupValueFormat string, nameIDFormat string, siteURL string, usernameAttribute string, wReplyURL string, ) *WsFederationApplicationSettingsApplication`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


