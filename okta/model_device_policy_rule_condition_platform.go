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

// DevicePolicyRuleConditionPlatform struct for DevicePolicyRuleConditionPlatform
type DevicePolicyRuleConditionPlatform struct {
	SupportedMDMFrameworks []string `json:"supportedMDMFrameworks,omitempty"`
	Types []string `json:"types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePolicyRuleConditionPlatform DevicePolicyRuleConditionPlatform

// NewDevicePolicyRuleConditionPlatform instantiates a new DevicePolicyRuleConditionPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePolicyRuleConditionPlatform() *DevicePolicyRuleConditionPlatform {
	this := DevicePolicyRuleConditionPlatform{}
	return &this
}

// NewDevicePolicyRuleConditionPlatformWithDefaults instantiates a new DevicePolicyRuleConditionPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePolicyRuleConditionPlatformWithDefaults() *DevicePolicyRuleConditionPlatform {
	this := DevicePolicyRuleConditionPlatform{}
	return &this
}

// GetSupportedMDMFrameworks returns the SupportedMDMFrameworks field value if set, zero value otherwise.
func (o *DevicePolicyRuleConditionPlatform) GetSupportedMDMFrameworks() []string {
	if o == nil || o.SupportedMDMFrameworks == nil {
		var ret []string
		return ret
	}
	return o.SupportedMDMFrameworks
}

// GetSupportedMDMFrameworksOk returns a tuple with the SupportedMDMFrameworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleConditionPlatform) GetSupportedMDMFrameworksOk() ([]string, bool) {
	if o == nil || o.SupportedMDMFrameworks == nil {
		return nil, false
	}
	return o.SupportedMDMFrameworks, true
}

// HasSupportedMDMFrameworks returns a boolean if a field has been set.
func (o *DevicePolicyRuleConditionPlatform) HasSupportedMDMFrameworks() bool {
	if o != nil && o.SupportedMDMFrameworks != nil {
		return true
	}

	return false
}

// SetSupportedMDMFrameworks gets a reference to the given []string and assigns it to the SupportedMDMFrameworks field.
func (o *DevicePolicyRuleConditionPlatform) SetSupportedMDMFrameworks(v []string) {
	o.SupportedMDMFrameworks = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *DevicePolicyRuleConditionPlatform) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleConditionPlatform) GetTypesOk() ([]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *DevicePolicyRuleConditionPlatform) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *DevicePolicyRuleConditionPlatform) SetTypes(v []string) {
	o.Types = v
}

func (o DevicePolicyRuleConditionPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportedMDMFrameworks != nil {
		toSerialize["supportedMDMFrameworks"] = o.SupportedMDMFrameworks
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DevicePolicyRuleConditionPlatform) UnmarshalJSON(bytes []byte) (err error) {
	varDevicePolicyRuleConditionPlatform := _DevicePolicyRuleConditionPlatform{}

	err = json.Unmarshal(bytes, &varDevicePolicyRuleConditionPlatform)
	if err == nil {
		*o = DevicePolicyRuleConditionPlatform(varDevicePolicyRuleConditionPlatform)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "supportedMDMFrameworks")
		delete(additionalProperties, "types")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDevicePolicyRuleConditionPlatform struct {
	value *DevicePolicyRuleConditionPlatform
	isSet bool
}

func (v NullableDevicePolicyRuleConditionPlatform) Get() *DevicePolicyRuleConditionPlatform {
	return v.value
}

func (v *NullableDevicePolicyRuleConditionPlatform) Set(val *DevicePolicyRuleConditionPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePolicyRuleConditionPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePolicyRuleConditionPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePolicyRuleConditionPlatform(val *DevicePolicyRuleConditionPlatform) *NullableDevicePolicyRuleConditionPlatform {
	return &NullableDevicePolicyRuleConditionPlatform{value: val, isSet: true}
}

func (v NullableDevicePolicyRuleConditionPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePolicyRuleConditionPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

