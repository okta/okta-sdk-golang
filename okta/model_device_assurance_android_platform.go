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

// DeviceAssuranceAndroidPlatform struct for DeviceAssuranceAndroidPlatform
type DeviceAssuranceAndroidPlatform struct {
	DeviceAssurance
	DiskEncryptionType *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
	Jailbreak *bool `json:"jailbreak,omitempty"`
	OsVersion *OSVersion `json:"osVersion,omitempty"`
	ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
	SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceAndroidPlatform DeviceAssuranceAndroidPlatform

// NewDeviceAssuranceAndroidPlatform instantiates a new DeviceAssuranceAndroidPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceAndroidPlatform() *DeviceAssuranceAndroidPlatform {
	this := DeviceAssuranceAndroidPlatform{}
	return &this
}

// NewDeviceAssuranceAndroidPlatformWithDefaults instantiates a new DeviceAssuranceAndroidPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceAndroidPlatformWithDefaults() *DeviceAssuranceAndroidPlatform {
	this := DeviceAssuranceAndroidPlatform{}
	return &this
}

// GetDiskEncryptionType returns the DiskEncryptionType field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatform) GetDiskEncryptionType() DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType {
	if o == nil || o.DiskEncryptionType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetDiskEncryptionTypeOk() (*DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType, bool) {
	if o == nil || o.DiskEncryptionType == nil {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasDiskEncryptionType() bool {
	if o != nil && o.DiskEncryptionType != nil {
		return true
	}

	return false
}

// SetDiskEncryptionType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType and assigns it to the DiskEncryptionType field.
func (o *DeviceAssuranceAndroidPlatform) SetDiskEncryptionType(v DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType) {
	o.DiskEncryptionType = &v
}

// GetJailbreak returns the Jailbreak field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatform) GetJailbreak() bool {
	if o == nil || o.Jailbreak == nil {
		var ret bool
		return ret
	}
	return *o.Jailbreak
}

// GetJailbreakOk returns a tuple with the Jailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetJailbreakOk() (*bool, bool) {
	if o == nil || o.Jailbreak == nil {
		return nil, false
	}
	return o.Jailbreak, true
}

// HasJailbreak returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasJailbreak() bool {
	if o != nil && o.Jailbreak != nil {
		return true
	}

	return false
}

// SetJailbreak gets a reference to the given bool and assigns it to the Jailbreak field.
func (o *DeviceAssuranceAndroidPlatform) SetJailbreak(v bool) {
	o.Jailbreak = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatform) GetOsVersion() OSVersion {
	if o == nil || o.OsVersion == nil {
		var ret OSVersion
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetOsVersionOk() (*OSVersion, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersion and assigns it to the OsVersion field.
func (o *DeviceAssuranceAndroidPlatform) SetOsVersion(v OSVersion) {
	o.OsVersion = &v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatform) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType {
	if o == nil || o.ScreenLockType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfScreenLockType
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool) {
	if o == nil || o.ScreenLockType == nil {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasScreenLockType() bool {
	if o != nil && o.ScreenLockType != nil {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfScreenLockType and assigns it to the ScreenLockType field.
func (o *DeviceAssuranceAndroidPlatform) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType) {
	o.ScreenLockType = &v
}

// GetSecureHardwarePresent returns the SecureHardwarePresent field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatform) GetSecureHardwarePresent() bool {
	if o == nil || o.SecureHardwarePresent == nil {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || o.SecureHardwarePresent == nil {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasSecureHardwarePresent() bool {
	if o != nil && o.SecureHardwarePresent != nil {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *DeviceAssuranceAndroidPlatform) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

func (o DeviceAssuranceAndroidPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedDeviceAssurance, errDeviceAssurance := json.Marshal(o.DeviceAssurance)
	if errDeviceAssurance != nil {
		return []byte{}, errDeviceAssurance
	}
	errDeviceAssurance = json.Unmarshal([]byte(serializedDeviceAssurance), &toSerialize)
	if errDeviceAssurance != nil {
		return []byte{}, errDeviceAssurance
	}
	if o.DiskEncryptionType != nil {
		toSerialize["diskEncryptionType"] = o.DiskEncryptionType
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
	if o.SecureHardwarePresent != nil {
		toSerialize["secureHardwarePresent"] = o.SecureHardwarePresent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceAndroidPlatform) UnmarshalJSON(bytes []byte) (err error) {
	type DeviceAssuranceAndroidPlatformWithoutEmbeddedStruct struct {
		DiskEncryptionType *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
		Jailbreak *bool `json:"jailbreak,omitempty"`
		OsVersion *OSVersion `json:"osVersion,omitempty"`
		ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
		SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
	}

	varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct := DeviceAssuranceAndroidPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct)
	if err == nil {
		varDeviceAssuranceAndroidPlatform := _DeviceAssuranceAndroidPlatform{}
		varDeviceAssuranceAndroidPlatform.DiskEncryptionType = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.DiskEncryptionType
		varDeviceAssuranceAndroidPlatform.Jailbreak = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.Jailbreak
		varDeviceAssuranceAndroidPlatform.OsVersion = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.OsVersion
		varDeviceAssuranceAndroidPlatform.ScreenLockType = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.ScreenLockType
		varDeviceAssuranceAndroidPlatform.SecureHardwarePresent = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.SecureHardwarePresent
		*o = DeviceAssuranceAndroidPlatform(varDeviceAssuranceAndroidPlatform)
	} else {
		return err
	}

	varDeviceAssuranceAndroidPlatform := _DeviceAssuranceAndroidPlatform{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceAndroidPlatform)
	if err == nil {
		o.DeviceAssurance = varDeviceAssuranceAndroidPlatform.DeviceAssurance
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "diskEncryptionType")
		delete(additionalProperties, "jailbreak")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "screenLockType")
		delete(additionalProperties, "secureHardwarePresent")

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

type NullableDeviceAssuranceAndroidPlatform struct {
	value *DeviceAssuranceAndroidPlatform
	isSet bool
}

func (v NullableDeviceAssuranceAndroidPlatform) Get() *DeviceAssuranceAndroidPlatform {
	return v.value
}

func (v *NullableDeviceAssuranceAndroidPlatform) Set(val *DeviceAssuranceAndroidPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceAndroidPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceAndroidPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceAndroidPlatform(val *DeviceAssuranceAndroidPlatform) *NullableDeviceAssuranceAndroidPlatform {
	return &NullableDeviceAssuranceAndroidPlatform{value: val, isSet: true}
}

func (v NullableDeviceAssuranceAndroidPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceAndroidPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

