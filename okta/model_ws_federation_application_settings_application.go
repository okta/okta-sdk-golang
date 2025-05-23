/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// WsFederationApplicationSettingsApplication struct for WsFederationApplicationSettingsApplication
type WsFederationApplicationSettingsApplication struct {
	AttributeStatements *string `json:"attributeStatements,omitempty"`
	AudienceRestriction *string `json:"audienceRestriction,omitempty"`
	AuthnContextClassRef *string `json:"authnContextClassRef,omitempty"`
	GroupFilter *string `json:"groupFilter,omitempty"`
	GroupName *string `json:"groupName,omitempty"`
	GroupValueFormat *string `json:"groupValueFormat,omitempty"`
	NameIDFormat *string `json:"nameIDFormat,omitempty"`
	Realm *string `json:"realm,omitempty"`
	SiteURL *string `json:"siteURL,omitempty"`
	UsernameAttribute *string `json:"usernameAttribute,omitempty"`
	WReplyOverride *bool `json:"wReplyOverride,omitempty"`
	WReplyURL *string `json:"wReplyURL,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WsFederationApplicationSettingsApplication WsFederationApplicationSettingsApplication

// NewWsFederationApplicationSettingsApplication instantiates a new WsFederationApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsFederationApplicationSettingsApplication() *WsFederationApplicationSettingsApplication {
	this := WsFederationApplicationSettingsApplication{}
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
	if o == nil || o.AttributeStatements == nil {
		var ret string
		return ret
	}
	return *o.AttributeStatements
}

// GetAttributeStatementsOk returns a tuple with the AttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetAttributeStatementsOk() (*string, bool) {
	if o == nil || o.AttributeStatements == nil {
		return nil, false
	}
	return o.AttributeStatements, true
}

// HasAttributeStatements returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasAttributeStatements() bool {
	if o != nil && o.AttributeStatements != nil {
		return true
	}

	return false
}

// SetAttributeStatements gets a reference to the given string and assigns it to the AttributeStatements field.
func (o *WsFederationApplicationSettingsApplication) SetAttributeStatements(v string) {
	o.AttributeStatements = &v
}

// GetAudienceRestriction returns the AudienceRestriction field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetAudienceRestriction() string {
	if o == nil || o.AudienceRestriction == nil {
		var ret string
		return ret
	}
	return *o.AudienceRestriction
}

// GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetAudienceRestrictionOk() (*string, bool) {
	if o == nil || o.AudienceRestriction == nil {
		return nil, false
	}
	return o.AudienceRestriction, true
}

// HasAudienceRestriction returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasAudienceRestriction() bool {
	if o != nil && o.AudienceRestriction != nil {
		return true
	}

	return false
}

// SetAudienceRestriction gets a reference to the given string and assigns it to the AudienceRestriction field.
func (o *WsFederationApplicationSettingsApplication) SetAudienceRestriction(v string) {
	o.AudienceRestriction = &v
}

// GetAuthnContextClassRef returns the AuthnContextClassRef field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetAuthnContextClassRef() string {
	if o == nil || o.AuthnContextClassRef == nil {
		var ret string
		return ret
	}
	return *o.AuthnContextClassRef
}

// GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetAuthnContextClassRefOk() (*string, bool) {
	if o == nil || o.AuthnContextClassRef == nil {
		return nil, false
	}
	return o.AuthnContextClassRef, true
}

// HasAuthnContextClassRef returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasAuthnContextClassRef() bool {
	if o != nil && o.AuthnContextClassRef != nil {
		return true
	}

	return false
}

// SetAuthnContextClassRef gets a reference to the given string and assigns it to the AuthnContextClassRef field.
func (o *WsFederationApplicationSettingsApplication) SetAuthnContextClassRef(v string) {
	o.AuthnContextClassRef = &v
}

// GetGroupFilter returns the GroupFilter field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetGroupFilter() string {
	if o == nil || o.GroupFilter == nil {
		var ret string
		return ret
	}
	return *o.GroupFilter
}

// GetGroupFilterOk returns a tuple with the GroupFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetGroupFilterOk() (*string, bool) {
	if o == nil || o.GroupFilter == nil {
		return nil, false
	}
	return o.GroupFilter, true
}

// HasGroupFilter returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasGroupFilter() bool {
	if o != nil && o.GroupFilter != nil {
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
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *WsFederationApplicationSettingsApplication) SetGroupName(v string) {
	o.GroupName = &v
}

// GetGroupValueFormat returns the GroupValueFormat field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetGroupValueFormat() string {
	if o == nil || o.GroupValueFormat == nil {
		var ret string
		return ret
	}
	return *o.GroupValueFormat
}

// GetGroupValueFormatOk returns a tuple with the GroupValueFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetGroupValueFormatOk() (*string, bool) {
	if o == nil || o.GroupValueFormat == nil {
		return nil, false
	}
	return o.GroupValueFormat, true
}

// HasGroupValueFormat returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasGroupValueFormat() bool {
	if o != nil && o.GroupValueFormat != nil {
		return true
	}

	return false
}

// SetGroupValueFormat gets a reference to the given string and assigns it to the GroupValueFormat field.
func (o *WsFederationApplicationSettingsApplication) SetGroupValueFormat(v string) {
	o.GroupValueFormat = &v
}

// GetNameIDFormat returns the NameIDFormat field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetNameIDFormat() string {
	if o == nil || o.NameIDFormat == nil {
		var ret string
		return ret
	}
	return *o.NameIDFormat
}

// GetNameIDFormatOk returns a tuple with the NameIDFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetNameIDFormatOk() (*string, bool) {
	if o == nil || o.NameIDFormat == nil {
		return nil, false
	}
	return o.NameIDFormat, true
}

// HasNameIDFormat returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasNameIDFormat() bool {
	if o != nil && o.NameIDFormat != nil {
		return true
	}

	return false
}

// SetNameIDFormat gets a reference to the given string and assigns it to the NameIDFormat field.
func (o *WsFederationApplicationSettingsApplication) SetNameIDFormat(v string) {
	o.NameIDFormat = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetRealm() string {
	if o == nil || o.Realm == nil {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetRealmOk() (*string, bool) {
	if o == nil || o.Realm == nil {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasRealm() bool {
	if o != nil && o.Realm != nil {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *WsFederationApplicationSettingsApplication) SetRealm(v string) {
	o.Realm = &v
}

// GetSiteURL returns the SiteURL field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetSiteURL() string {
	if o == nil || o.SiteURL == nil {
		var ret string
		return ret
	}
	return *o.SiteURL
}

// GetSiteURLOk returns a tuple with the SiteURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetSiteURLOk() (*string, bool) {
	if o == nil || o.SiteURL == nil {
		return nil, false
	}
	return o.SiteURL, true
}

// HasSiteURL returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasSiteURL() bool {
	if o != nil && o.SiteURL != nil {
		return true
	}

	return false
}

// SetSiteURL gets a reference to the given string and assigns it to the SiteURL field.
func (o *WsFederationApplicationSettingsApplication) SetSiteURL(v string) {
	o.SiteURL = &v
}

// GetUsernameAttribute returns the UsernameAttribute field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetUsernameAttribute() string {
	if o == nil || o.UsernameAttribute == nil {
		var ret string
		return ret
	}
	return *o.UsernameAttribute
}

// GetUsernameAttributeOk returns a tuple with the UsernameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetUsernameAttributeOk() (*string, bool) {
	if o == nil || o.UsernameAttribute == nil {
		return nil, false
	}
	return o.UsernameAttribute, true
}

// HasUsernameAttribute returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasUsernameAttribute() bool {
	if o != nil && o.UsernameAttribute != nil {
		return true
	}

	return false
}

// SetUsernameAttribute gets a reference to the given string and assigns it to the UsernameAttribute field.
func (o *WsFederationApplicationSettingsApplication) SetUsernameAttribute(v string) {
	o.UsernameAttribute = &v
}

// GetWReplyOverride returns the WReplyOverride field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetWReplyOverride() bool {
	if o == nil || o.WReplyOverride == nil {
		var ret bool
		return ret
	}
	return *o.WReplyOverride
}

// GetWReplyOverrideOk returns a tuple with the WReplyOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetWReplyOverrideOk() (*bool, bool) {
	if o == nil || o.WReplyOverride == nil {
		return nil, false
	}
	return o.WReplyOverride, true
}

// HasWReplyOverride returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasWReplyOverride() bool {
	if o != nil && o.WReplyOverride != nil {
		return true
	}

	return false
}

// SetWReplyOverride gets a reference to the given bool and assigns it to the WReplyOverride field.
func (o *WsFederationApplicationSettingsApplication) SetWReplyOverride(v bool) {
	o.WReplyOverride = &v
}

// GetWReplyURL returns the WReplyURL field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsApplication) GetWReplyURL() string {
	if o == nil || o.WReplyURL == nil {
		var ret string
		return ret
	}
	return *o.WReplyURL
}

// GetWReplyURLOk returns a tuple with the WReplyURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsApplication) GetWReplyURLOk() (*string, bool) {
	if o == nil || o.WReplyURL == nil {
		return nil, false
	}
	return o.WReplyURL, true
}

// HasWReplyURL returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsApplication) HasWReplyURL() bool {
	if o != nil && o.WReplyURL != nil {
		return true
	}

	return false
}

// SetWReplyURL gets a reference to the given string and assigns it to the WReplyURL field.
func (o *WsFederationApplicationSettingsApplication) SetWReplyURL(v string) {
	o.WReplyURL = &v
}

func (o WsFederationApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeStatements != nil {
		toSerialize["attributeStatements"] = o.AttributeStatements
	}
	if o.AudienceRestriction != nil {
		toSerialize["audienceRestriction"] = o.AudienceRestriction
	}
	if o.AuthnContextClassRef != nil {
		toSerialize["authnContextClassRef"] = o.AuthnContextClassRef
	}
	if o.GroupFilter != nil {
		toSerialize["groupFilter"] = o.GroupFilter
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.GroupValueFormat != nil {
		toSerialize["groupValueFormat"] = o.GroupValueFormat
	}
	if o.NameIDFormat != nil {
		toSerialize["nameIDFormat"] = o.NameIDFormat
	}
	if o.Realm != nil {
		toSerialize["realm"] = o.Realm
	}
	if o.SiteURL != nil {
		toSerialize["siteURL"] = o.SiteURL
	}
	if o.UsernameAttribute != nil {
		toSerialize["usernameAttribute"] = o.UsernameAttribute
	}
	if o.WReplyOverride != nil {
		toSerialize["wReplyOverride"] = o.WReplyOverride
	}
	if o.WReplyURL != nil {
		toSerialize["wReplyURL"] = o.WReplyURL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WsFederationApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varWsFederationApplicationSettingsApplication := _WsFederationApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varWsFederationApplicationSettingsApplication)
	if err == nil {
		*o = WsFederationApplicationSettingsApplication(varWsFederationApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
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

