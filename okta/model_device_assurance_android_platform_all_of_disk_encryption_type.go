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

// DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType struct for DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType
type DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType struct {
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType

// NewDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType instantiates a new DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType() *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType {
	this := DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType{}
	return &this
}

// NewDeviceAssuranceAndroidPlatformAllOfDiskEncryptionTypeWithDefaults instantiates a new DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceAndroidPlatformAllOfDiskEncryptionTypeWithDefaults() *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType {
	this := DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) SetInclude(v []string) {
	o.Include = v
}

func (o DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType := _DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType)
	if err == nil {
		*o = DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType(varDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType struct {
	value *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType
	isSet bool
}

func (v NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) Get() *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType {
	return v.value
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) Set(val *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType(val *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) *NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType {
	return &NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType{value: val, isSet: true}
}

func (v NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

