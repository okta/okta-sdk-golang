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

// checks if the DeviceDisplayName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceDisplayName{}

// DeviceDisplayName Display name of the device
type DeviceDisplayName struct {
	// Indicates whether the associated value is Personal Identifiable Information (PII) and requires masking
	Sensitive *bool `json:"sensitive,omitempty"`
	// Display name of the device
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceDisplayName DeviceDisplayName

// NewDeviceDisplayName instantiates a new DeviceDisplayName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceDisplayName() *DeviceDisplayName {
	this := DeviceDisplayName{}
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// NewDeviceDisplayNameWithDefaults instantiates a new DeviceDisplayName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceDisplayNameWithDefaults() *DeviceDisplayName {
	this := DeviceDisplayName{}
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *DeviceDisplayName) GetSensitive() bool {
	if o == nil || IsNil(o.Sensitive) {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDisplayName) GetSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Sensitive) {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *DeviceDisplayName) HasSensitive() bool {
	if o != nil && !IsNil(o.Sensitive) {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *DeviceDisplayName) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceDisplayName) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDisplayName) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceDisplayName) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DeviceDisplayName) SetValue(v string) {
	o.Value = &v
}

func (o DeviceDisplayName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceDisplayName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sensitive) {
		toSerialize["sensitive"] = o.Sensitive
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceDisplayName) UnmarshalJSON(data []byte) (err error) {
	varDeviceDisplayName := _DeviceDisplayName{}

	err = json.Unmarshal(data, &varDeviceDisplayName)

	if err != nil {
		return err
	}

	*o = DeviceDisplayName(varDeviceDisplayName)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sensitive")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceDisplayName struct {
	value *DeviceDisplayName
	isSet bool
}

func (v NullableDeviceDisplayName) Get() *DeviceDisplayName {
	return v.value
}

func (v *NullableDeviceDisplayName) Set(val *DeviceDisplayName) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceDisplayName) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceDisplayName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceDisplayName(val *DeviceDisplayName) *NullableDeviceDisplayName {
	return &NullableDeviceDisplayName{value: val, isSet: true}
}

func (v NullableDeviceDisplayName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceDisplayName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
