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
	"reflect"
	"strings"
)

// checks if the DeviceAssuranceAndroidPlatform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAssuranceAndroidPlatform{}

// DeviceAssuranceAndroidPlatform struct for DeviceAssuranceAndroidPlatform
type DeviceAssuranceAndroidPlatform struct {
	DeviceAssurance
	DiskEncryptionType        *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType        `json:"diskEncryptionType,omitempty"`
	Jailbreak                 *bool                                                         `json:"jailbreak,omitempty"`
	OsVersion                 *OSVersion                                                    `json:"osVersion,omitempty"`
	ScreenLockType            *DeviceAssuranceAndroidPlatformAllOfScreenLockType            `json:"screenLockType,omitempty"`
	SecureHardwarePresent     *bool                                                         `json:"secureHardwarePresent,omitempty"`
	ThirdPartySignalProviders *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	AdditionalProperties      map[string]interface{}
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
	if o == nil || IsNil(o.DiskEncryptionType) {
		var ret DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetDiskEncryptionTypeOk() (*DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType, bool) {
	if o == nil || IsNil(o.DiskEncryptionType) {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasDiskEncryptionType() bool {
	if o != nil && !IsNil(o.DiskEncryptionType) {
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
	if o == nil || IsNil(o.Jailbreak) {
		var ret bool
		return ret
	}
	return *o.Jailbreak
}

// GetJailbreakOk returns a tuple with the Jailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetJailbreakOk() (*bool, bool) {
	if o == nil || IsNil(o.Jailbreak) {
		return nil, false
	}
	return o.Jailbreak, true
}

// HasJailbreak returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasJailbreak() bool {
	if o != nil && !IsNil(o.Jailbreak) {
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
	if o == nil || IsNil(o.OsVersion) {
		var ret OSVersion
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetOsVersionOk() (*OSVersion, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
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
	if o == nil || IsNil(o.ScreenLockType) {
		var ret DeviceAssuranceAndroidPlatformAllOfScreenLockType
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool) {
	if o == nil || IsNil(o.ScreenLockType) {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasScreenLockType() bool {
	if o != nil && !IsNil(o.ScreenLockType) {
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
	if o == nil || IsNil(o.SecureHardwarePresent) {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureHardwarePresent) {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasSecureHardwarePresent() bool {
	if o != nil && !IsNil(o.SecureHardwarePresent) {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *DeviceAssuranceAndroidPlatform) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

// GetThirdPartySignalProviders returns the ThirdPartySignalProviders field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatform) GetThirdPartySignalProviders() DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders {
	if o == nil || IsNil(o.ThirdPartySignalProviders) {
		var ret DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders
		return ret
	}
	return *o.ThirdPartySignalProviders
}

// GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatform) GetThirdPartySignalProvidersOk() (*DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders, bool) {
	if o == nil || IsNil(o.ThirdPartySignalProviders) {
		return nil, false
	}
	return o.ThirdPartySignalProviders, true
}

// HasThirdPartySignalProviders returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatform) HasThirdPartySignalProviders() bool {
	if o != nil && !IsNil(o.ThirdPartySignalProviders) {
		return true
	}

	return false
}

// SetThirdPartySignalProviders gets a reference to the given DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders and assigns it to the ThirdPartySignalProviders field.
func (o *DeviceAssuranceAndroidPlatform) SetThirdPartySignalProviders(v DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) {
	o.ThirdPartySignalProviders = &v
}

func (o DeviceAssuranceAndroidPlatform) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAssuranceAndroidPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDeviceAssurance, errDeviceAssurance := json.Marshal(o.DeviceAssurance)
	if errDeviceAssurance != nil {
		return map[string]interface{}{}, errDeviceAssurance
	}
	errDeviceAssurance = json.Unmarshal([]byte(serializedDeviceAssurance), &toSerialize)
	if errDeviceAssurance != nil {
		return map[string]interface{}{}, errDeviceAssurance
	}
	if !IsNil(o.DiskEncryptionType) {
		toSerialize["diskEncryptionType"] = o.DiskEncryptionType
	}
	if !IsNil(o.Jailbreak) {
		toSerialize["jailbreak"] = o.Jailbreak
	}
	if !IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !IsNil(o.ScreenLockType) {
		toSerialize["screenLockType"] = o.ScreenLockType
	}
	if !IsNil(o.SecureHardwarePresent) {
		toSerialize["secureHardwarePresent"] = o.SecureHardwarePresent
	}
	if !IsNil(o.ThirdPartySignalProviders) {
		toSerialize["thirdPartySignalProviders"] = o.ThirdPartySignalProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceAssuranceAndroidPlatform) UnmarshalJSON(data []byte) (err error) {
	type DeviceAssuranceAndroidPlatformWithoutEmbeddedStruct struct {
		DiskEncryptionType        *DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType        `json:"diskEncryptionType,omitempty"`
		Jailbreak                 *bool                                                         `json:"jailbreak,omitempty"`
		OsVersion                 *OSVersion                                                    `json:"osVersion,omitempty"`
		ScreenLockType            *DeviceAssuranceAndroidPlatformAllOfScreenLockType            `json:"screenLockType,omitempty"`
		SecureHardwarePresent     *bool                                                         `json:"secureHardwarePresent,omitempty"`
		ThirdPartySignalProviders *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	}

	varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct := DeviceAssuranceAndroidPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct)
	if err == nil {
		varDeviceAssuranceAndroidPlatform := _DeviceAssuranceAndroidPlatform{}
		varDeviceAssuranceAndroidPlatform.DiskEncryptionType = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.DiskEncryptionType
		varDeviceAssuranceAndroidPlatform.Jailbreak = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.Jailbreak
		varDeviceAssuranceAndroidPlatform.OsVersion = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.OsVersion
		varDeviceAssuranceAndroidPlatform.ScreenLockType = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.ScreenLockType
		varDeviceAssuranceAndroidPlatform.SecureHardwarePresent = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.SecureHardwarePresent
		varDeviceAssuranceAndroidPlatform.ThirdPartySignalProviders = varDeviceAssuranceAndroidPlatformWithoutEmbeddedStruct.ThirdPartySignalProviders
		*o = DeviceAssuranceAndroidPlatform(varDeviceAssuranceAndroidPlatform)
	} else {
		return err
	}

	varDeviceAssuranceAndroidPlatform := _DeviceAssuranceAndroidPlatform{}

	err = json.Unmarshal(data, &varDeviceAssuranceAndroidPlatform)
	if err == nil {
		o.DeviceAssurance = varDeviceAssuranceAndroidPlatform.DeviceAssurance
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "diskEncryptionType")
		delete(additionalProperties, "jailbreak")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "screenLockType")
		delete(additionalProperties, "secureHardwarePresent")
		delete(additionalProperties, "thirdPartySignalProviders")

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
