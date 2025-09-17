/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the WsFederationApplicationSettingsApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WsFederationApplicationSettingsApplication{}

// WsFederationApplicationSettingsApplication struct for WsFederationApplicationSettingsApplication
type WsFederationApplicationSettingsApplication struct {
	// You can federate user attributes such as Okta profile fields, LDAP, Active Directory, and Workday values. The SP uses the federated WS-Fed attribute values accordingly.
	AttributeStatements *string `json:"attributeStatements,omitempty"`
	// The entity ID of the SP. Use the entity ID value exactly as provided by the SP.
	AudienceRestriction string `json:"audienceRestriction"`
	// Identifies the SAML authentication context class for the assertion's authentication statement
	AuthnContextClassRef string `json:"authnContextClassRef"`
	// A regular expression that filters for the User Groups you want included with the `groupName` attribute. If the matching User Group has a corresponding AD group, then the attribute statement includes the value of the attribute specified by `groupValueFormat`. If the matching User Group doesn't contain a corresponding AD group, then the `groupName` is used in the attribute statement.
	GroupFilter *string `json:"groupFilter,omitempty"`
	// The group name to include in the WS-Fed response attribute statement. This property is used in conjunction with the `groupFilter` property.  Groups that are filtered through the `groupFilter` expression are included with the `groupName` in the attribute statement. Any users that belong to the group you've filtered are included in the WS-Fed response attribute statement.
	GroupName *string `json:"groupName,omitempty"`
	// Specifies the WS-Fed assertion attribute value for filtered groups. This attribute is only applied to Active Directory groups.
	GroupValueFormat string `json:"groupValueFormat"`
	// The username format that you send in the WS-Fed response
	NameIDFormat string `json:"nameIDFormat"`
	// The uniform resource identifier (URI) of the WS-Fed app that's used to share resources securely within a domain. It's the identity that's sent to the Okta IdP when signing in. See [Realm name](https://help.okta.com/okta_help.htm?type=oie&id=ext_Apps_Configure_Okta_Template_WS_Federation#Realm).
	Realm *string `json:"realm,omitempty"`
	// Launch URL for the web app
	SiteURL string `json:"siteURL"`
	// Specifies additional username attribute statements to include in the WS-Fed assertion
	UsernameAttribute string `json:"usernameAttribute"`
	// Enables a web app to override the `wReplyURL` URL with a reply parameter.
	WReplyOverride *bool `json:"wReplyOverride,omitempty"`
	// The WS-Fed SP endpoint where your users sign in
	WReplyURL            string `json:"wReplyURL"`
	AdditionalProperties map[string]interface{}
}

type _WsFederationApplicationSettingsApplication WsFederationApplicationSettingsApplication

// NewWsFederationApplicationSettingsApplication instantiates a new WsFederationApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsFederationApplicationSettingsApplication(audienceRestriction string, authnContextClassRef string, groupValueFormat string, nameIDFormat string, siteURL string, usernameAttribute string, wReplyURL string) *WsFederationApplicationSettingsApplication {
	this := WsFederationApplicationSettingsApplication{}
	this.AudienceRestriction = audienceRestriction
	this.AuthnContextClassRef = authnContextClassRef
	this.GroupValueFormat = groupValueFormat
	this.NameIDFormat = nameIDFormat
	this.SiteURL = siteURL
	this.UsernameAttribute = usernameAttribute
	this.WReplyURL = wReplyURL
	return &this
}

// NewWsFederationApplicationSettingsApplicationWithDefaults instantiates a new WsFederationApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsFederationApplicationSettingsApplicationWithDefaults() *WsFederationApplicationSettingsApplication {
	this := WsFederationApplicationSettingsApplication{}
	return &this
}

// GetAttributeStatements returns the AttributeStatements field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetAttributeStatements() string {
	if o == nil || IsNil(o.AttributeStatements) {
		var ret string
		return ret
	}
	return *o.AttributeStatements
}

// GetAttributeStatementsOk returns a tuple with the AttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetAttributeStatementsOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeStatements) {
		return nil, false
	}
	return o.AttributeStatements, true
}

// HasAttributeStatements returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasAttributeStatements() bool {
	if o != nil && !IsNil(o.AttributeStatements) {
		return true
	}

	return false
}

// SetAttributeStatements gets a reference to the given string and assigns it to the AttributeStatements field.
func (o *WsFederationApplicationSettingsApplication) SetAttributeStatements(v string) {
	o.AttributeStatements = &v
}

// GetAudienceRestriction returns the AudienceRestriction field value
func (o *WsFederationApplicationSettingsApplication) GetAudienceRestriction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudienceRestriction
}

// GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetAudienceRestrictionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceRestriction, true
}

// SetAudienceRestriction sets field value
func (o *WsFederationApplicationSettingsApplication) SetAudienceRestriction(v string) {
	o.AudienceRestriction = v
}

// GetAuthnContextClassRef returns the AuthnContextClassRef field value
func (o *WsFederationApplicationSettingsApplication) GetAuthnContextClassRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthnContextClassRef
}

// GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetAuthnContextClassRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthnContextClassRef, true
}

// SetAuthnContextClassRef sets field value
func (o *WsFederationApplicationSettingsApplication) SetAuthnContextClassRef(v string) {
	o.AuthnContextClassRef = v
}

// GetGroupFilter returns the GroupFilter field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetGroupFilter() string {
	if o == nil || IsNil(o.GroupFilter) {
		var ret string
		return ret
	}
	return *o.GroupFilter
}

// GetGroupFilterOk returns a tuple with the GroupFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetGroupFilterOk() (*string, bool) {
	if o == nil || IsNil(o.GroupFilter) {
		return nil, false
	}
	return o.GroupFilter, true
}

// HasGroupFilter returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasGroupFilter() bool {
	if o != nil && !IsNil(o.GroupFilter) {
		return true
	}

	return false
}

// SetGroupFilter gets a reference to the given string and assigns it to the GroupFilter field.
func (o *WsFederationApplicationSettingsApplication) SetGroupFilter(v string) {
	o.GroupFilter = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *WsFederationApplicationSettingsApplication) SetGroupName(v string) {
	o.GroupName = &v
}

// GetGroupValueFormat returns the GroupValueFormat field value
func (o *WsFederationApplicationSettingsApplication) GetGroupValueFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupValueFormat
}

// GetGroupValueFormatOk returns a tuple with the GroupValueFormat field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetGroupValueFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupValueFormat, true
}

// SetGroupValueFormat sets field value
func (o *WsFederationApplicationSettingsApplication) SetGroupValueFormat(v string) {
	o.GroupValueFormat = v
}

// GetNameIDFormat returns the NameIDFormat field value
func (o *WsFederationApplicationSettingsApplication) GetNameIDFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameIDFormat
}

// GetNameIDFormatOk returns a tuple with the NameIDFormat field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetNameIDFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameIDFormat, true
}

// SetNameIDFormat sets field value
func (o *WsFederationApplicationSettingsApplication) SetNameIDFormat(v string) {
	o.NameIDFormat = v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *WsFederationApplicationSettingsApplication) SetRealm(v string) {
	o.Realm = &v
}

// GetSiteURL returns the SiteURL field value
func (o *WsFederationApplicationSettingsApplication) GetSiteURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteURL
}

// GetSiteURLOk returns a tuple with the SiteURL field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetSiteURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteURL, true
}

// SetSiteURL sets field value
func (o *WsFederationApplicationSettingsApplication) SetSiteURL(v string) {
	o.SiteURL = v
}

// GetUsernameAttribute returns the UsernameAttribute field value
func (o *WsFederationApplicationSettingsApplication) GetUsernameAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameAttribute
}

// GetUsernameAttributeOk returns a tuple with the UsernameAttribute field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetUsernameAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameAttribute, true
}

// SetUsernameAttribute sets field value
func (o *WsFederationApplicationSettingsApplication) SetUsernameAttribute(v string) {
	o.UsernameAttribute = v
}

// GetWReplyOverride returns the WReplyOverride field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetWReplyOverride() bool {
	if o == nil || IsNil(o.WReplyOverride) {
		var ret bool
		return ret
	}
	return *o.WReplyOverride
}

// GetWReplyOverrideOk returns a tuple with the WReplyOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetWReplyOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.WReplyOverride) {
		return nil, false
	}
	return o.WReplyOverride, true
}

// HasWReplyOverride returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasWReplyOverride() bool {
	if o != nil && !IsNil(o.WReplyOverride) {
		return true
	}

	return false
}

// SetWReplyOverride gets a reference to the given bool and assigns it to the WReplyOverride field.
func (o *WsFederationApplicationSettingsApplication) SetWReplyOverride(v bool) {
	o.WReplyOverride = &v
}

// GetWReplyURL returns the WReplyURL field value
func (o *WsFederationApplicationSettingsApplication) GetWReplyURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WReplyURL
}

// GetWReplyURLOk returns a tuple with the WReplyURL field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetWReplyURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WReplyURL, true
}

// SetWReplyURL sets field value
func (o *WsFederationApplicationSettingsApplication) SetWReplyURL(v string) {
	o.WReplyURL = v
}

func (o WsFederationApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WsFederationApplicationSettingsApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeStatements) {
		toSerialize["attributeStatements"] = o.AttributeStatements
	}
	toSerialize["audienceRestriction"] = o.AudienceRestriction
	toSerialize["authnContextClassRef"] = o.AuthnContextClassRef
	if !IsNil(o.GroupFilter) {
		toSerialize["groupFilter"] = o.GroupFilter
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	toSerialize["groupValueFormat"] = o.GroupValueFormat
	toSerialize["nameIDFormat"] = o.NameIDFormat
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	toSerialize["siteURL"] = o.SiteURL
	toSerialize["usernameAttribute"] = o.UsernameAttribute
	if !IsNil(o.WReplyOverride) {
		toSerialize["wReplyOverride"] = o.WReplyOverride
	}
	toSerialize["wReplyURL"] = o.WReplyURL

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WsFederationApplicationSettingsApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audienceRestriction",
		"authnContextClassRef",
		"groupValueFormat",
		"nameIDFormat",
		"siteURL",
		"usernameAttribute",
		"wReplyURL",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWsFederationApplicationSettingsApplication := _WsFederationApplicationSettingsApplication{}

	err = json.Unmarshal(data, &varWsFederationApplicationSettingsApplication)

	if err != nil {
		return err
	}

	*o = WsFederationApplicationSettingsApplication(varWsFederationApplicationSettingsApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeStatements")
		delete(additionalProperties, "audienceRestriction")
		delete(additionalProperties, "authnContextClassRef")
		delete(additionalProperties, "groupFilter")
		delete(additionalProperties, "groupName")
		delete(additionalProperties, "groupValueFormat")
		delete(additionalProperties, "nameIDFormat")
		delete(additionalProperties, "realm")
		delete(additionalProperties, "siteURL")
		delete(additionalProperties, "usernameAttribute")
		delete(additionalProperties, "wReplyOverride")
		delete(additionalProperties, "wReplyURL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWsFederationApplicationSettingsApplication struct {
	value *WsFederationApplicationSettingsApplication
	isSet bool
}

func (v NullableWsFederationApplicationSettingsApplication) Get() *WsFederationApplicationSettingsApplication {
	return v.value
}

func (v *NullableWsFederationApplicationSettingsApplication) Set(val *WsFederationApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableWsFederationApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableWsFederationApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWsFederationApplicationSettingsApplication(val *WsFederationApplicationSettingsApplication) *NullableWsFederationApplicationSettingsApplication {
	return &NullableWsFederationApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableWsFederationApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWsFederationApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
