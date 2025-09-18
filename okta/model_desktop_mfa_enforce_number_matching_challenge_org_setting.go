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

// checks if the DesktopMFAEnforceNumberMatchingChallengeOrgSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DesktopMFAEnforceNumberMatchingChallengeOrgSetting{}

// DesktopMFAEnforceNumberMatchingChallengeOrgSetting struct for DesktopMFAEnforceNumberMatchingChallengeOrgSetting
type DesktopMFAEnforceNumberMatchingChallengeOrgSetting struct {
	// Indicates whether or not the Desktop MFA Enforce Number Matching Challenge push notifications feature is enabled
	DesktopMFAEnforceNumberMatchingChallengeEnabled *bool `json:"desktopMFAEnforceNumberMatchingChallengeEnabled,omitempty"`
	AdditionalProperties                            map[string]interface{}
}

type _DesktopMFAEnforceNumberMatchingChallengeOrgSetting DesktopMFAEnforceNumberMatchingChallengeOrgSetting

// NewDesktopMFAEnforceNumberMatchingChallengeOrgSetting instantiates a new DesktopMFAEnforceNumberMatchingChallengeOrgSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopMFAEnforceNumberMatchingChallengeOrgSetting() *DesktopMFAEnforceNumberMatchingChallengeOrgSetting {
	this := DesktopMFAEnforceNumberMatchingChallengeOrgSetting{}
	var desktopMFAEnforceNumberMatchingChallengeEnabled bool = false
	this.DesktopMFAEnforceNumberMatchingChallengeEnabled = &desktopMFAEnforceNumberMatchingChallengeEnabled
	return &this
}

// NewDesktopMFAEnforceNumberMatchingChallengeOrgSettingWithDefaults instantiates a new DesktopMFAEnforceNumberMatchingChallengeOrgSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopMFAEnforceNumberMatchingChallengeOrgSettingWithDefaults() *DesktopMFAEnforceNumberMatchingChallengeOrgSetting {
	this := DesktopMFAEnforceNumberMatchingChallengeOrgSetting{}
	var desktopMFAEnforceNumberMatchingChallengeEnabled bool = false
	this.DesktopMFAEnforceNumberMatchingChallengeEnabled = &desktopMFAEnforceNumberMatchingChallengeEnabled
	return &this
}

// GetDesktopMFAEnforceNumberMatchingChallengeEnabled returns the DesktopMFAEnforceNumberMatchingChallengeEnabled field value if set, zero value otherwise.
func (o *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) GetDesktopMFAEnforceNumberMatchingChallengeEnabled() bool {
	if o == nil || IsNil(o.DesktopMFAEnforceNumberMatchingChallengeEnabled) {
		var ret bool
		return ret
	}
	return *o.DesktopMFAEnforceNumberMatchingChallengeEnabled
}

// GetDesktopMFAEnforceNumberMatchingChallengeEnabledOk returns a tuple with the DesktopMFAEnforceNumberMatchingChallengeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) GetDesktopMFAEnforceNumberMatchingChallengeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DesktopMFAEnforceNumberMatchingChallengeEnabled) {
		return nil, false
	}
	return o.DesktopMFAEnforceNumberMatchingChallengeEnabled, true
}

// HasDesktopMFAEnforceNumberMatchingChallengeEnabled returns a boolean if a field has been set.
func (o *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) HasDesktopMFAEnforceNumberMatchingChallengeEnabled() bool {
	if o != nil && !IsNil(o.DesktopMFAEnforceNumberMatchingChallengeEnabled) {
		return true
	}

	return false
}

// SetDesktopMFAEnforceNumberMatchingChallengeEnabled gets a reference to the given bool and assigns it to the DesktopMFAEnforceNumberMatchingChallengeEnabled field.
func (o *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) SetDesktopMFAEnforceNumberMatchingChallengeEnabled(v bool) {
	o.DesktopMFAEnforceNumberMatchingChallengeEnabled = &v
}

func (o DesktopMFAEnforceNumberMatchingChallengeOrgSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DesktopMFAEnforceNumberMatchingChallengeOrgSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DesktopMFAEnforceNumberMatchingChallengeEnabled) {
		toSerialize["desktopMFAEnforceNumberMatchingChallengeEnabled"] = o.DesktopMFAEnforceNumberMatchingChallengeEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) UnmarshalJSON(data []byte) (err error) {
	varDesktopMFAEnforceNumberMatchingChallengeOrgSetting := _DesktopMFAEnforceNumberMatchingChallengeOrgSetting{}

	err = json.Unmarshal(data, &varDesktopMFAEnforceNumberMatchingChallengeOrgSetting)

	if err != nil {
		return err
	}

	*o = DesktopMFAEnforceNumberMatchingChallengeOrgSetting(varDesktopMFAEnforceNumberMatchingChallengeOrgSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "desktopMFAEnforceNumberMatchingChallengeEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting struct {
	value *DesktopMFAEnforceNumberMatchingChallengeOrgSetting
	isSet bool
}

func (v NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting) Get() *DesktopMFAEnforceNumberMatchingChallengeOrgSetting {
	return v.value
}

func (v *NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting) Set(val *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting(val *DesktopMFAEnforceNumberMatchingChallengeOrgSetting) *NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting {
	return &NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting{value: val, isSet: true}
}

func (v NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopMFAEnforceNumberMatchingChallengeOrgSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
