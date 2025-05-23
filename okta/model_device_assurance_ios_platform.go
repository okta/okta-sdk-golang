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
	"reflect"
	"strings"
)

// DeviceAssuranceIOSPlatform struct for DeviceAssuranceIOSPlatform
type DeviceAssuranceIOSPlatform struct {
	DeviceAssurance
	Jailbreak *bool `json:"jailbreak,omitempty"`
	OsVersion *OSVersion `json:"osVersion,omitempty"`
	ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceIOSPlatform DeviceAssuranceIOSPlatform

// NewDeviceAssuranceIOSPlatform instantiates a new DeviceAssuranceIOSPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceIOSPlatform() *DeviceAssuranceIOSPlatform {
	this := DeviceAssuranceIOSPlatform{}
	return &this
}

// NewDeviceAssuranceIOSPlatformWithDefaults instantiates a new DeviceAssuranceIOSPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceIOSPlatformWithDefaults() *DeviceAssuranceIOSPlatform {
	this := DeviceAssuranceIOSPlatform{}
	return &this
}

// GetJailbreak returns the Jailbreak field value if set, zero value otherwise.
func (o *DeviceAssuranceIOSPlatform) GetJailbreak() bool {
	if o == nil || o.Jailbreak == nil {
		var ret bool
		return ret
	}
	return *o.Jailbreak
}

// GetJailbreakOk returns a tuple with the Jailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceIOSPlatform) GetJailbreakOk() (*bool, bool) {
	if o == nil || o.Jailbreak == nil {
		return nil, false
	}
	return o.Jailbreak, true
}

// HasJailbreak returns a boolean if a field has been set.
func (o *DeviceAssuranceIOSPlatform) HasJailbreak() bool {
	if o != nil && o.Jailbreak != nil {
		return true
	}

	return false
}

// SetJailbreak gets a reference to the given bool and assigns it to the Jailbreak field.
func (o *DeviceAssuranceIOSPlatform) SetJailbreak(v bool) {
	o.Jailbreak = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceAssuranceIOSPlatform) GetOsVersion() OSVersion {
	if o == nil || o.OsVersion == nil {
		var ret OSVersion
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceIOSPlatform) GetOsVersionOk() (*OSVersion, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAssuranceIOSPlatform) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersion and assigns it to the OsVersion field.
func (o *DeviceAssuranceIOSPlatform) SetOsVersion(v OSVersion) {
	o.OsVersion = &v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *DeviceAssuranceIOSPlatform) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType {
	if o == nil || o.ScreenLockType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfScreenLockType
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceIOSPlatform) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool) {
	if o == nil || o.ScreenLockType == nil {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceAssuranceIOSPlatform) HasScreenLockType() bool {
	if o != nil && o.ScreenLockType != nil {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfScreenLockType and assigns it to the ScreenLockType field.
func (o *DeviceAssuranceIOSPlatform) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType) {
	o.ScreenLockType = &v
}

func (o DeviceAssuranceIOSPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedDeviceAssurance, errDeviceAssurance := json.Marshal(o.DeviceAssurance)
	if errDeviceAssurance != nil {
		return []byte{}, errDeviceAssurance
	}
	errDeviceAssurance = json.Unmarshal([]byte(serializedDeviceAssurance), &toSerialize)
	if errDeviceAssurance != nil {
		return []byte{}, errDeviceAssurance
	}
	if o.Jailbreak != nil {
		toSerialize["jailbreak"] = o.Jailbreak
	}
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.ScreenLockType != nil {
		toSerialize["screenLockType"] = o.ScreenLockType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceIOSPlatform) UnmarshalJSON(bytes []byte) (err error) {
	type DeviceAssuranceIOSPlatformWithoutEmbeddedStruct struct {
		Jailbreak *bool `json:"jailbreak,omitempty"`
		OsVersion *OSVersion `json:"osVersion,omitempty"`
		ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
	}

	varDeviceAssuranceIOSPlatformWithoutEmbeddedStruct := DeviceAssuranceIOSPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceIOSPlatformWithoutEmbeddedStruct)
	if err == nil {
		varDeviceAssuranceIOSPlatform := _DeviceAssuranceIOSPlatform{}
		varDeviceAssuranceIOSPlatform.Jailbreak = varDeviceAssuranceIOSPlatformWithoutEmbeddedStruct.Jailbreak
		varDeviceAssuranceIOSPlatform.OsVersion = varDeviceAssuranceIOSPlatformWithoutEmbeddedStruct.OsVersion
		varDeviceAssuranceIOSPlatform.ScreenLockType = varDeviceAssuranceIOSPlatformWithoutEmbeddedStruct.ScreenLockType
		*o = DeviceAssuranceIOSPlatform(varDeviceAssuranceIOSPlatform)
	} else {
		return err
	}

	varDeviceAssuranceIOSPlatform := _DeviceAssuranceIOSPlatform{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceIOSPlatform)
	if err == nil {
		o.DeviceAssurance = varDeviceAssuranceIOSPlatform.DeviceAssurance
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "jailbreak")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "screenLockType")

		// remove fields from embedded structs
		reflectDeviceAssurance := reflect.ValueOf(o.DeviceAssurance)
		for i := 0; i < reflectDeviceAssurance.Type().NumField(); i++ {
			t := reflectDeviceAssurance.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceAssuranceIOSPlatform struct {
	value *DeviceAssuranceIOSPlatform
	isSet bool
}

func (v NullableDeviceAssuranceIOSPlatform) Get() *DeviceAssuranceIOSPlatform {
	return v.value
}

func (v *NullableDeviceAssuranceIOSPlatform) Set(val *DeviceAssuranceIOSPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceIOSPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceIOSPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceIOSPlatform(val *DeviceAssuranceIOSPlatform) *NullableDeviceAssuranceIOSPlatform {
	return &NullableDeviceAssuranceIOSPlatform{value: val, isSet: true}
}

func (v NullableDeviceAssuranceIOSPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceIOSPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

