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

// checks if the DeviceSignalCollectionPlatformConditionEvaluatorPlatform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSignalCollectionPlatformConditionEvaluatorPlatform{}

// DeviceSignalCollectionPlatformConditionEvaluatorPlatform struct for DeviceSignalCollectionPlatformConditionEvaluatorPlatform
type DeviceSignalCollectionPlatformConditionEvaluatorPlatform struct {
	Os *DeviceSignalCollectionPlatformConditionEvaluatorPlatformOperatingSystem `json:"os,omitempty"`
	// The type of platform
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceSignalCollectionPlatformConditionEvaluatorPlatform DeviceSignalCollectionPlatformConditionEvaluatorPlatform

// NewDeviceSignalCollectionPlatformConditionEvaluatorPlatform instantiates a new DeviceSignalCollectionPlatformConditionEvaluatorPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSignalCollectionPlatformConditionEvaluatorPlatform() *DeviceSignalCollectionPlatformConditionEvaluatorPlatform {
	this := DeviceSignalCollectionPlatformConditionEvaluatorPlatform{}
	return &this
}

// NewDeviceSignalCollectionPlatformConditionEvaluatorPlatformWithDefaults instantiates a new DeviceSignalCollectionPlatformConditionEvaluatorPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSignalCollectionPlatformConditionEvaluatorPlatformWithDefaults() *DeviceSignalCollectionPlatformConditionEvaluatorPlatform {
	this := DeviceSignalCollectionPlatformConditionEvaluatorPlatform{}
	return &this
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) GetOs() DeviceSignalCollectionPlatformConditionEvaluatorPlatformOperatingSystem {
	if o == nil || IsNil(o.Os) {
		var ret DeviceSignalCollectionPlatformConditionEvaluatorPlatformOperatingSystem
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) GetOsOk() (*DeviceSignalCollectionPlatformConditionEvaluatorPlatformOperatingSystem, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given DeviceSignalCollectionPlatformConditionEvaluatorPlatformOperatingSystem and assigns it to the Os field.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) SetOs(v DeviceSignalCollectionPlatformConditionEvaluatorPlatformOperatingSystem) {
	o.Os = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) SetType(v string) {
	o.Type = &v
}

func (o DeviceSignalCollectionPlatformConditionEvaluatorPlatform) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSignalCollectionPlatformConditionEvaluatorPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) UnmarshalJSON(data []byte) (err error) {
	varDeviceSignalCollectionPlatformConditionEvaluatorPlatform := _DeviceSignalCollectionPlatformConditionEvaluatorPlatform{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPlatformConditionEvaluatorPlatform)

	if err != nil {
		return err
	}

	*o = DeviceSignalCollectionPlatformConditionEvaluatorPlatform(varDeviceSignalCollectionPlatformConditionEvaluatorPlatform)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "os")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform struct {
	value *DeviceSignalCollectionPlatformConditionEvaluatorPlatform
	isSet bool
}

func (v NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform) Get() *DeviceSignalCollectionPlatformConditionEvaluatorPlatform {
	return v.value
}

func (v *NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform) Set(val *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform(val *DeviceSignalCollectionPlatformConditionEvaluatorPlatform) *NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform {
	return &NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform{value: val, isSet: true}
}

func (v NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSignalCollectionPlatformConditionEvaluatorPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
