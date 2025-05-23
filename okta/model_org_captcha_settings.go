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

// OrgCAPTCHASettings 
type OrgCAPTCHASettings struct {
	// The unique key of the associated CAPTCHA instance
	CaptchaId *string `json:"captchaId,omitempty"`
	// An array of pages that have CAPTCHA enabled
	EnabledPages []string `json:"enabledPages,omitempty"`
	Links *OrgCAPTCHASettingsLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCAPTCHASettings OrgCAPTCHASettings

// NewOrgCAPTCHASettings instantiates a new OrgCAPTCHASettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCAPTCHASettings() *OrgCAPTCHASettings {
	this := OrgCAPTCHASettings{}
	return &this
}

// NewOrgCAPTCHASettingsWithDefaults instantiates a new OrgCAPTCHASettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCAPTCHASettingsWithDefaults() *OrgCAPTCHASettings {
	this := OrgCAPTCHASettings{}
	return &this
}

// GetCaptchaId returns the CaptchaId field value if set, zero value otherwise.
func (o *OrgCAPTCHASettings) GetCaptchaId() string {
	if o == nil || o.CaptchaId == nil {
		var ret string
		return ret
	}
	return *o.CaptchaId
}

// GetCaptchaIdOk returns a tuple with the CaptchaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCAPTCHASettings) GetCaptchaIdOk() (*string, bool) {
	if o == nil || o.CaptchaId == nil {
		return nil, false
	}
	return o.CaptchaId, true
}

// HasCaptchaId returns a boolean if a field has been set.
func (o *OrgCAPTCHASettings) HasCaptchaId() bool {
	if o != nil && o.CaptchaId != nil {
		return true
	}

	return false
}

// SetCaptchaId gets a reference to the given string and assigns it to the CaptchaId field.
func (o *OrgCAPTCHASettings) SetCaptchaId(v string) {
	o.CaptchaId = &v
}

// GetEnabledPages returns the EnabledPages field value if set, zero value otherwise.
func (o *OrgCAPTCHASettings) GetEnabledPages() []string {
	if o == nil || o.EnabledPages == nil {
		var ret []string
		return ret
	}
	return o.EnabledPages
}

// GetEnabledPagesOk returns a tuple with the EnabledPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCAPTCHASettings) GetEnabledPagesOk() ([]string, bool) {
	if o == nil || o.EnabledPages == nil {
		return nil, false
	}
	return o.EnabledPages, true
}

// HasEnabledPages returns a boolean if a field has been set.
func (o *OrgCAPTCHASettings) HasEnabledPages() bool {
	if o != nil && o.EnabledPages != nil {
		return true
	}

	return false
}

// SetEnabledPages gets a reference to the given []string and assigns it to the EnabledPages field.
func (o *OrgCAPTCHASettings) SetEnabledPages(v []string) {
	o.EnabledPages = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgCAPTCHASettings) GetLinks() OrgCAPTCHASettingsLinks {
	if o == nil || o.Links == nil {
		var ret OrgCAPTCHASettingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCAPTCHASettings) GetLinksOk() (*OrgCAPTCHASettingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgCAPTCHASettings) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgCAPTCHASettingsLinks and assigns it to the Links field.
func (o *OrgCAPTCHASettings) SetLinks(v OrgCAPTCHASettingsLinks) {
	o.Links = &v
}

func (o OrgCAPTCHASettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaptchaId != nil {
		toSerialize["captchaId"] = o.CaptchaId
	}
	if o.EnabledPages != nil {
		toSerialize["enabledPages"] = o.EnabledPages
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgCAPTCHASettings) UnmarshalJSON(bytes []byte) (err error) {
	varOrgCAPTCHASettings := _OrgCAPTCHASettings{}

	err = json.Unmarshal(bytes, &varOrgCAPTCHASettings)
	if err == nil {
		*o = OrgCAPTCHASettings(varOrgCAPTCHASettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "captchaId")
		delete(additionalProperties, "enabledPages")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgCAPTCHASettings struct {
	value *OrgCAPTCHASettings
	isSet bool
}

func (v NullableOrgCAPTCHASettings) Get() *OrgCAPTCHASettings {
	return v.value
}

func (v *NullableOrgCAPTCHASettings) Set(val *OrgCAPTCHASettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCAPTCHASettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCAPTCHASettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCAPTCHASettings(val *OrgCAPTCHASettings) *NullableOrgCAPTCHASettings {
	return &NullableOrgCAPTCHASettings{value: val, isSet: true}
}

func (v NullableOrgCAPTCHASettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCAPTCHASettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

