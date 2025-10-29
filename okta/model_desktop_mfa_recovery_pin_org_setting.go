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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DesktopMFARecoveryPinOrgSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DesktopMFARecoveryPinOrgSetting{}

// DesktopMFARecoveryPinOrgSetting struct for DesktopMFARecoveryPinOrgSetting
type DesktopMFARecoveryPinOrgSetting struct {
	// Indicates whether or not the Desktop MFA Recovery PIN feature is enabled
	DesktopMFARecoveryPinEnabled *bool `json:"desktopMFARecoveryPinEnabled,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _DesktopMFARecoveryPinOrgSetting DesktopMFARecoveryPinOrgSetting

// NewDesktopMFARecoveryPinOrgSetting instantiates a new DesktopMFARecoveryPinOrgSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopMFARecoveryPinOrgSetting() *DesktopMFARecoveryPinOrgSetting {
	this := DesktopMFARecoveryPinOrgSetting{}
	var desktopMFARecoveryPinEnabled bool = false
	this.DesktopMFARecoveryPinEnabled = &desktopMFARecoveryPinEnabled
	return &this
}

// NewDesktopMFARecoveryPinOrgSettingWithDefaults instantiates a new DesktopMFARecoveryPinOrgSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopMFARecoveryPinOrgSettingWithDefaults() *DesktopMFARecoveryPinOrgSetting {
	this := DesktopMFARecoveryPinOrgSetting{}
	var desktopMFARecoveryPinEnabled bool = false
	this.DesktopMFARecoveryPinEnabled = &desktopMFARecoveryPinEnabled
	return &this
}

// GetDesktopMFARecoveryPinEnabled returns the DesktopMFARecoveryPinEnabled field value if set, zero value otherwise.
func (o *DesktopMFARecoveryPinOrgSetting) GetDesktopMFARecoveryPinEnabled() bool {
	if o == nil || IsNil(o.DesktopMFARecoveryPinEnabled) {
		var ret bool
		return ret
	}
	return *o.DesktopMFARecoveryPinEnabled
}

// GetDesktopMFARecoveryPinEnabledOk returns a tuple with the DesktopMFARecoveryPinEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopMFARecoveryPinOrgSetting) GetDesktopMFARecoveryPinEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DesktopMFARecoveryPinEnabled) {
		return nil, false
	}
	return o.DesktopMFARecoveryPinEnabled, true
}

// HasDesktopMFARecoveryPinEnabled returns a boolean if a field has been set.
func (o *DesktopMFARecoveryPinOrgSetting) HasDesktopMFARecoveryPinEnabled() bool {
	if o != nil && !IsNil(o.DesktopMFARecoveryPinEnabled) {
		return true
	}

	return false
}

// SetDesktopMFARecoveryPinEnabled gets a reference to the given bool and assigns it to the DesktopMFARecoveryPinEnabled field.
func (o *DesktopMFARecoveryPinOrgSetting) SetDesktopMFARecoveryPinEnabled(v bool) {
	o.DesktopMFARecoveryPinEnabled = &v
}

func (o DesktopMFARecoveryPinOrgSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DesktopMFARecoveryPinOrgSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DesktopMFARecoveryPinEnabled) {
		toSerialize["desktopMFARecoveryPinEnabled"] = o.DesktopMFARecoveryPinEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DesktopMFARecoveryPinOrgSetting) UnmarshalJSON(data []byte) (err error) {
	varDesktopMFARecoveryPinOrgSetting := _DesktopMFARecoveryPinOrgSetting{}

	err = json.Unmarshal(data, &varDesktopMFARecoveryPinOrgSetting)

	if err != nil {
		return err
	}

	*o = DesktopMFARecoveryPinOrgSetting(varDesktopMFARecoveryPinOrgSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "desktopMFARecoveryPinEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDesktopMFARecoveryPinOrgSetting struct {
	value *DesktopMFARecoveryPinOrgSetting
	isSet bool
}

func (v NullableDesktopMFARecoveryPinOrgSetting) Get() *DesktopMFARecoveryPinOrgSetting {
	return v.value
}

func (v *NullableDesktopMFARecoveryPinOrgSetting) Set(val *DesktopMFARecoveryPinOrgSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopMFARecoveryPinOrgSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopMFARecoveryPinOrgSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopMFARecoveryPinOrgSetting(val *DesktopMFARecoveryPinOrgSetting) *NullableDesktopMFARecoveryPinOrgSetting {
	return &NullableDesktopMFARecoveryPinOrgSetting{value: val, isSet: true}
}

func (v NullableDesktopMFARecoveryPinOrgSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopMFARecoveryPinOrgSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
