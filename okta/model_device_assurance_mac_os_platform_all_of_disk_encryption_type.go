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

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType struct for DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType
type DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType struct {
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType

// NewDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType instantiates a new DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType() *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType {
	this := DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType{}
	return &this
}

// NewDeviceAssuranceMacOSPlatformAllOfDiskEncryptionTypeWithDefaults instantiates a new DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceMacOSPlatformAllOfDiskEncryptionTypeWithDefaults() *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType {
	this := DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) SetInclude(v []string) {
	o.Include = v
}

func (o DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType := _DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType)
	if err == nil {
		*o = DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType(varDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType)
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

type NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType struct {
	value *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType
	isSet bool
}

func (v NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) Get() *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType {
	return v.value
}

func (v *NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) Set(val *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType(val *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) *NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType {
	return &NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType{value: val, isSet: true}
}

func (v NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
