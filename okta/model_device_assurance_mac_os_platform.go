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

// DeviceAssuranceMacOSPlatform struct for DeviceAssuranceMacOSPlatform
type DeviceAssuranceMacOSPlatform struct {
	DeviceAssurance
	DiskEncryptionType *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
	OsVersion *OSVersion `json:"osVersion,omitempty"`
	ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
	SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
	ThirdPartySignalProviders *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceMacOSPlatform DeviceAssuranceMacOSPlatform

// NewDeviceAssuranceMacOSPlatform instantiates a new DeviceAssuranceMacOSPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceMacOSPlatform() *DeviceAssuranceMacOSPlatform {
	this := DeviceAssuranceMacOSPlatform{}
	return &this
}

// NewDeviceAssuranceMacOSPlatformWithDefaults instantiates a new DeviceAssuranceMacOSPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceMacOSPlatformWithDefaults() *DeviceAssuranceMacOSPlatform {
	this := DeviceAssuranceMacOSPlatform{}
	return &this
}

// GetDiskEncryptionType returns the DiskEncryptionType field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatform) GetDiskEncryptionType() DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType {
	if o == nil || o.DiskEncryptionType == nil {
		var ret DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatform) GetDiskEncryptionTypeOk() (*DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType, bool) {
	if o == nil || o.DiskEncryptionType == nil {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatform) HasDiskEncryptionType() bool {
	if o != nil && o.DiskEncryptionType != nil {
		return true
	}

	return false
}

// SetDiskEncryptionType gets a reference to the given DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType and assigns it to the DiskEncryptionType field.
func (o *DeviceAssuranceMacOSPlatform) SetDiskEncryptionType(v DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) {
	o.DiskEncryptionType = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatform) GetOsVersion() OSVersion {
	if o == nil || o.OsVersion == nil {
		var ret OSVersion
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatform) GetOsVersionOk() (*OSVersion, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatform) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersion and assigns it to the OsVersion field.
func (o *DeviceAssuranceMacOSPlatform) SetOsVersion(v OSVersion) {
	o.OsVersion = &v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatform) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType {
	if o == nil || o.ScreenLockType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfScreenLockType
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatform) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool) {
	if o == nil || o.ScreenLockType == nil {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatform) HasScreenLockType() bool {
	if o != nil && o.ScreenLockType != nil {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfScreenLockType and assigns it to the ScreenLockType field.
func (o *DeviceAssuranceMacOSPlatform) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType) {
	o.ScreenLockType = &v
}

// GetSecureHardwarePresent returns the SecureHardwarePresent field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatform) GetSecureHardwarePresent() bool {
	if o == nil || o.SecureHardwarePresent == nil {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatform) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || o.SecureHardwarePresent == nil {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatform) HasSecureHardwarePresent() bool {
	if o != nil && o.SecureHardwarePresent != nil {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *DeviceAssuranceMacOSPlatform) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

// GetThirdPartySignalProviders returns the ThirdPartySignalProviders field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatform) GetThirdPartySignalProviders() DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders {
	if o == nil || o.ThirdPartySignalProviders == nil {
		var ret DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders
		return ret
	}
	return *o.ThirdPartySignalProviders
}

// GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatform) GetThirdPartySignalProvidersOk() (*DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders, bool) {
	if o == nil || o.ThirdPartySignalProviders == nil {
		return nil, false
	}
	return o.ThirdPartySignalProviders, true
}

// HasThirdPartySignalProviders returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatform) HasThirdPartySignalProviders() bool {
	if o != nil && o.ThirdPartySignalProviders != nil {
		return true
	}

	return false
}

// SetThirdPartySignalProviders gets a reference to the given DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders and assigns it to the ThirdPartySignalProviders field.
func (o *DeviceAssuranceMacOSPlatform) SetThirdPartySignalProviders(v DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) {
	o.ThirdPartySignalProviders = &v
}

func (o DeviceAssuranceMacOSPlatform) MarshalJSON() ([]byte, error) {
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
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.ScreenLockType != nil {
		toSerialize["screenLockType"] = o.ScreenLockType
	}
	if o.SecureHardwarePresent != nil {
		toSerialize["secureHardwarePresent"] = o.SecureHardwarePresent
	}
	if o.ThirdPartySignalProviders != nil {
		toSerialize["thirdPartySignalProviders"] = o.ThirdPartySignalProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceMacOSPlatform) UnmarshalJSON(bytes []byte) (err error) {
	type DeviceAssuranceMacOSPlatformWithoutEmbeddedStruct struct {
		DiskEncryptionType *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
		OsVersion *OSVersion `json:"osVersion,omitempty"`
		ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
		SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
		ThirdPartySignalProviders *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	}

	varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct := DeviceAssuranceMacOSPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct)
	if err == nil {
		varDeviceAssuranceMacOSPlatform := _DeviceAssuranceMacOSPlatform{}
		varDeviceAssuranceMacOSPlatform.DiskEncryptionType = varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct.DiskEncryptionType
		varDeviceAssuranceMacOSPlatform.OsVersion = varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct.OsVersion
		varDeviceAssuranceMacOSPlatform.ScreenLockType = varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct.ScreenLockType
		varDeviceAssuranceMacOSPlatform.SecureHardwarePresent = varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct.SecureHardwarePresent
		varDeviceAssuranceMacOSPlatform.ThirdPartySignalProviders = varDeviceAssuranceMacOSPlatformWithoutEmbeddedStruct.ThirdPartySignalProviders
		*o = DeviceAssuranceMacOSPlatform(varDeviceAssuranceMacOSPlatform)
	} else {
		return err
	}

	varDeviceAssuranceMacOSPlatform := _DeviceAssuranceMacOSPlatform{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceMacOSPlatform)
	if err == nil {
		o.DeviceAssurance = varDeviceAssuranceMacOSPlatform.DeviceAssurance
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "diskEncryptionType")
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
	} else {
		return err
	}

	return err
}

type NullableDeviceAssuranceMacOSPlatform struct {
	value *DeviceAssuranceMacOSPlatform
	isSet bool
}

func (v NullableDeviceAssuranceMacOSPlatform) Get() *DeviceAssuranceMacOSPlatform {
	return v.value
}

func (v *NullableDeviceAssuranceMacOSPlatform) Set(val *DeviceAssuranceMacOSPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceMacOSPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceMacOSPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceMacOSPlatform(val *DeviceAssuranceMacOSPlatform) *NullableDeviceAssuranceMacOSPlatform {
	return &NullableDeviceAssuranceMacOSPlatform{value: val, isSet: true}
}

func (v NullableDeviceAssuranceMacOSPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceMacOSPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

