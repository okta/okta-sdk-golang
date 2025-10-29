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

// checks if the DevicePostureChecks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureChecks{}

// DevicePostureChecks <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Represents the Device Posture Checks configuration for the device assurance policy
type DevicePostureChecks struct {
	// An array of key value pairs including Device Posture Check `variableNames`
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureChecks DevicePostureChecks

// NewDevicePostureChecks instantiates a new DevicePostureChecks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureChecks() *DevicePostureChecks {
	this := DevicePostureChecks{}
	return &this
}

// NewDevicePostureChecksWithDefaults instantiates a new DevicePostureChecks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureChecksWithDefaults() *DevicePostureChecks {
	this := DevicePostureChecks{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *DevicePostureChecks) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecks) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *DevicePostureChecks) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *DevicePostureChecks) SetInclude(v []string) {
	o.Include = v
}

func (o DevicePostureChecks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostureChecks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureChecks) UnmarshalJSON(data []byte) (err error) {
	varDevicePostureChecks := _DevicePostureChecks{}

	err = json.Unmarshal(data, &varDevicePostureChecks)

	if err != nil {
		return err
	}

	*o = DevicePostureChecks(varDevicePostureChecks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePostureChecks struct {
	value *DevicePostureChecks
	isSet bool
}

func (v NullableDevicePostureChecks) Get() *DevicePostureChecks {
	return v.value
}

func (v *NullableDevicePostureChecks) Set(val *DevicePostureChecks) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureChecks) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureChecks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureChecks(val *DevicePostureChecks) *NullableDevicePostureChecks {
	return &NullableDevicePostureChecks{value: val, isSet: true}
}

func (v NullableDevicePostureChecks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureChecks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
