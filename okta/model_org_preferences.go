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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OrgPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgPreferences{}

// OrgPreferences struct for OrgPreferences
type OrgPreferences struct {
	// Indicates if the footer is shown on the End-User Dashboard
	ShowEndUserFooter    *bool                `json:"showEndUserFooter,omitempty"`
	Links                *OrgPreferencesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgPreferences OrgPreferences

// NewOrgPreferences instantiates a new OrgPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgPreferences() *OrgPreferences {
	this := OrgPreferences{}
	return &this
}

// NewOrgPreferencesWithDefaults instantiates a new OrgPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgPreferencesWithDefaults() *OrgPreferences {
	this := OrgPreferences{}
	return &this
}

// GetShowEndUserFooter returns the ShowEndUserFooter field value if set, zero value otherwise.
func (o *OrgPreferences) GetShowEndUserFooter() bool {
	if o == nil || IsNil(o.ShowEndUserFooter) {
		var ret bool
		return ret
	}
	return *o.ShowEndUserFooter
}

// GetShowEndUserFooterOk returns a tuple with the ShowEndUserFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPreferences) GetShowEndUserFooterOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowEndUserFooter) {
		return nil, false
	}
	return o.ShowEndUserFooter, true
}

// HasShowEndUserFooter returns a boolean if a field has been set.
func (o *OrgPreferences) HasShowEndUserFooter() bool {
	if o != nil && !IsNil(o.ShowEndUserFooter) {
		return true
	}

	return false
}

// SetShowEndUserFooter gets a reference to the given bool and assigns it to the ShowEndUserFooter field.
func (o *OrgPreferences) SetShowEndUserFooter(v bool) {
	o.ShowEndUserFooter = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgPreferences) GetLinks() OrgPreferencesLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrgPreferencesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgPreferences) GetLinksOk() (*OrgPreferencesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgPreferences) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgPreferencesLinks and assigns it to the Links field.
func (o *OrgPreferences) SetLinks(v OrgPreferencesLinks) {
	o.Links = &v
}

func (o OrgPreferences) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShowEndUserFooter) {
		toSerialize["showEndUserFooter"] = o.ShowEndUserFooter
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgPreferences) UnmarshalJSON(data []byte) (err error) {
	varOrgPreferences := _OrgPreferences{}

	err = json.Unmarshal(data, &varOrgPreferences)

	if err != nil {
		return err
	}

	*o = OrgPreferences(varOrgPreferences)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "showEndUserFooter")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgPreferences struct {
	value *OrgPreferences
	isSet bool
}

func (v NullableOrgPreferences) Get() *OrgPreferences {
	return v.value
}

func (v *NullableOrgPreferences) Set(val *OrgPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgPreferences(val *OrgPreferences) *NullableOrgPreferences {
	return &NullableOrgPreferences{value: val, isSet: true}
}

func (v NullableOrgPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
