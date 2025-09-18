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
)

// checks if the OrgPreferencesLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgPreferencesLinks{}

// OrgPreferencesLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for this object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgPreferencesLinks struct {
	// Link to hide the footer in the End-User Dashboard
	HideEndUserFooter *HrefObject `json:"hideEndUserFooter,omitempty"`
	// Link to show the footer on the End-User Dashboard
	ShowEndUserFooter    *HrefObject `json:"showEndUserFooter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgPreferencesLinks OrgPreferencesLinks

// NewOrgPreferencesLinks instantiates a new OrgPreferencesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgPreferencesLinks() *OrgPreferencesLinks {
	this := OrgPreferencesLinks{}
	return &this
}

// NewOrgPreferencesLinksWithDefaults instantiates a new OrgPreferencesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgPreferencesLinksWithDefaults() *OrgPreferencesLinks {
	this := OrgPreferencesLinks{}
	return &this
}

// GetHideEndUserFooter returns the HideEndUserFooter field value if set, zero value otherwise.
func (o *OrgPreferencesLinks) GetHideEndUserFooter() HrefObject {
	if o == nil || IsNil(o.HideEndUserFooter) {
		var ret HrefObject
		return ret
	}
	return *o.HideEndUserFooter
}

// GetHideEndUserFooterOk returns a tuple with the HideEndUserFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPreferencesLinks) GetHideEndUserFooterOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.HideEndUserFooter) {
		return nil, false
	}
	return o.HideEndUserFooter, true
}

// HasHideEndUserFooter returns a boolean if a field has been set.
func (o *OrgPreferencesLinks) HasHideEndUserFooter() bool {
	if o != nil && !IsNil(o.HideEndUserFooter) {
		return true
	}

	return false
}

// SetHideEndUserFooter gets a reference to the given HrefObject and assigns it to the HideEndUserFooter field.
func (o *OrgPreferencesLinks) SetHideEndUserFooter(v HrefObject) {
	o.HideEndUserFooter = &v
}

// GetShowEndUserFooter returns the ShowEndUserFooter field value if set, zero value otherwise.
func (o *OrgPreferencesLinks) GetShowEndUserFooter() HrefObject {
	if o == nil || IsNil(o.ShowEndUserFooter) {
		var ret HrefObject
		return ret
	}
	return *o.ShowEndUserFooter
}

// GetShowEndUserFooterOk returns a tuple with the ShowEndUserFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPreferencesLinks) GetShowEndUserFooterOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.ShowEndUserFooter) {
		return nil, false
	}
	return o.ShowEndUserFooter, true
}

// HasShowEndUserFooter returns a boolean if a field has been set.
func (o *OrgPreferencesLinks) HasShowEndUserFooter() bool {
	if o != nil && !IsNil(o.ShowEndUserFooter) {
		return true
	}

	return false
}

// SetShowEndUserFooter gets a reference to the given HrefObject and assigns it to the ShowEndUserFooter field.
func (o *OrgPreferencesLinks) SetShowEndUserFooter(v HrefObject) {
	o.ShowEndUserFooter = &v
}

func (o OrgPreferencesLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgPreferencesLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HideEndUserFooter) {
		toSerialize["hideEndUserFooter"] = o.HideEndUserFooter
	}
	if !IsNil(o.ShowEndUserFooter) {
		toSerialize["showEndUserFooter"] = o.ShowEndUserFooter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgPreferencesLinks) UnmarshalJSON(data []byte) (err error) {
	varOrgPreferencesLinks := _OrgPreferencesLinks{}

	err = json.Unmarshal(data, &varOrgPreferencesLinks)

	if err != nil {
		return err
	}

	*o = OrgPreferencesLinks(varOrgPreferencesLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hideEndUserFooter")
		delete(additionalProperties, "showEndUserFooter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgPreferencesLinks struct {
	value *OrgPreferencesLinks
	isSet bool
}

func (v NullableOrgPreferencesLinks) Get() *OrgPreferencesLinks {
	return v.value
}

func (v *NullableOrgPreferencesLinks) Set(val *OrgPreferencesLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgPreferencesLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgPreferencesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgPreferencesLinks(val *OrgPreferencesLinks) *NullableOrgPreferencesLinks {
	return &NullableOrgPreferencesLinks{value: val, isSet: true}
}

func (v NullableOrgPreferencesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgPreferencesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
