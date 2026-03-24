/*
Okta Admin Management API

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

// checks if the DevicePostureChecksIncludeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureChecksIncludeInner{}

// DevicePostureChecksIncludeInner struct for DevicePostureChecksIncludeInner
type DevicePostureChecksIncludeInner struct {
	// The device posture check key
	VariableName *string `json:"variableName,omitempty"`
	// The device posture check value
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureChecksIncludeInner DevicePostureChecksIncludeInner

// NewDevicePostureChecksIncludeInner instantiates a new DevicePostureChecksIncludeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureChecksIncludeInner() *DevicePostureChecksIncludeInner {
	this := DevicePostureChecksIncludeInner{}
	return &this
}

// NewDevicePostureChecksIncludeInnerWithDefaults instantiates a new DevicePostureChecksIncludeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureChecksIncludeInnerWithDefaults() *DevicePostureChecksIncludeInner {
	this := DevicePostureChecksIncludeInner{}
	return &this
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *DevicePostureChecksIncludeInner) GetVariableName() string {
	if o == nil || IsNil(o.VariableName) {
		var ret string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksIncludeInner) GetVariableNameOk() (*string, bool) {
	if o == nil || IsNil(o.VariableName) {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *DevicePostureChecksIncludeInner) HasVariableName() bool {
	if o != nil && !IsNil(o.VariableName) {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given string and assigns it to the VariableName field.
func (o *DevicePostureChecksIncludeInner) SetVariableName(v string) {
	o.VariableName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DevicePostureChecksIncludeInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksIncludeInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DevicePostureChecksIncludeInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DevicePostureChecksIncludeInner) SetValue(v string) {
	o.Value = &v
}

func (o DevicePostureChecksIncludeInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostureChecksIncludeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VariableName) {
		toSerialize["variableName"] = o.VariableName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureChecksIncludeInner) UnmarshalJSON(data []byte) (err error) {
	varDevicePostureChecksIncludeInner := _DevicePostureChecksIncludeInner{}

	err = json.Unmarshal(data, &varDevicePostureChecksIncludeInner)

	if err != nil {
		return err
	}

	*o = DevicePostureChecksIncludeInner(varDevicePostureChecksIncludeInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "variableName")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePostureChecksIncludeInner struct {
	value *DevicePostureChecksIncludeInner
	isSet bool
}

func (v NullableDevicePostureChecksIncludeInner) Get() *DevicePostureChecksIncludeInner {
	return v.value
}

func (v *NullableDevicePostureChecksIncludeInner) Set(val *DevicePostureChecksIncludeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureChecksIncludeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureChecksIncludeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureChecksIncludeInner(val *DevicePostureChecksIncludeInner) *NullableDevicePostureChecksIncludeInner {
	return &NullableDevicePostureChecksIncludeInner{value: val, isSet: true}
}

func (v NullableDevicePostureChecksIncludeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureChecksIncludeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
