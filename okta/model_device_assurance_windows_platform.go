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

// DeviceAssuranceWindowsPlatform struct for DeviceAssuranceWindowsPlatform
type DeviceAssuranceWindowsPlatform struct {
	DeviceAssurance
	DiskEncryptionType *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
	OsVersion *OSVersionFourComponents `json:"osVersion,omitempty"`
	// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>Specifies the Windows version requirements for the assurance policy. Each requirement must correspond to a different major version (Windows 11 or Windows 10). If a requirement isn't specified for a major version, then devices on that major version satisfy the condition.  There are two types of OS requirements: * **Static**: A specific Windows version requirement that doesn't change until you update the policy. A static OS Windows requirement is specified with `majorVersionConstraint` and `minimum`. * **Dynamic**: A Windows version requirement that is relative to the latest major release and security patch. A dynamic OS Windows requirement is specified with `majorVersionConstraint` and `dynamicVersionRequirement`.  > **Note:** Dynamic OS requirements are available only if the **Dynamic OS version compliance** [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled. The `osVersionConstraints` property is only supported for the Windows platform. You can't specify both `osVersion.minimum` and `osVersionConstraints` properties at the same time. 
	OsVersionConstraints []OSVersionConstraint `json:"osVersionConstraints,omitempty"`
	ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
	SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
	ThirdPartySignalProviders *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceWindowsPlatform DeviceAssuranceWindowsPlatform

// NewDeviceAssuranceWindowsPlatform instantiates a new DeviceAssuranceWindowsPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceWindowsPlatform() *DeviceAssuranceWindowsPlatform {
	this := DeviceAssuranceWindowsPlatform{}
	return &this
}

// NewDeviceAssuranceWindowsPlatformWithDefaults instantiates a new DeviceAssuranceWindowsPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceWindowsPlatformWithDefaults() *DeviceAssuranceWindowsPlatform {
	this := DeviceAssuranceWindowsPlatform{}
	return &this
}

// GetDiskEncryptionType returns the DiskEncryptionType field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatform) GetDiskEncryptionType() DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType {
	if o == nil || o.DiskEncryptionType == nil {
		var ret DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatform) GetDiskEncryptionTypeOk() (*DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType, bool) {
	if o == nil || o.DiskEncryptionType == nil {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatform) HasDiskEncryptionType() bool {
	if o != nil && o.DiskEncryptionType != nil {
		return true
	}

	return false
}

// SetDiskEncryptionType gets a reference to the given DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType and assigns it to the DiskEncryptionType field.
func (o *DeviceAssuranceWindowsPlatform) SetDiskEncryptionType(v DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType) {
	o.DiskEncryptionType = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatform) GetOsVersion() OSVersionFourComponents {
	if o == nil || o.OsVersion == nil {
		var ret OSVersionFourComponents
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatform) GetOsVersionOk() (*OSVersionFourComponents, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatform) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersionFourComponents and assigns it to the OsVersion field.
func (o *DeviceAssuranceWindowsPlatform) SetOsVersion(v OSVersionFourComponents) {
	o.OsVersion = &v
}

// GetOsVersionConstraints returns the OsVersionConstraints field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatform) GetOsVersionConstraints() []OSVersionConstraint {
	if o == nil || o.OsVersionConstraints == nil {
		var ret []OSVersionConstraint
		return ret
	}
	return o.OsVersionConstraints
}

// GetOsVersionConstraintsOk returns a tuple with the OsVersionConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatform) GetOsVersionConstraintsOk() ([]OSVersionConstraint, bool) {
	if o == nil || o.OsVersionConstraints == nil {
		return nil, false
	}
	return o.OsVersionConstraints, true
}

// HasOsVersionConstraints returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatform) HasOsVersionConstraints() bool {
	if o != nil && o.OsVersionConstraints != nil {
		return true
	}

	return false
}

// SetOsVersionConstraints gets a reference to the given []OSVersionConstraint and assigns it to the OsVersionConstraints field.
func (o *DeviceAssuranceWindowsPlatform) SetOsVersionConstraints(v []OSVersionConstraint) {
	o.OsVersionConstraints = v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatform) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType {
	if o == nil || o.ScreenLockType == nil {
		var ret DeviceAssuranceAndroidPlatformAllOfScreenLockType
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatform) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool) {
	if o == nil || o.ScreenLockType == nil {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatform) HasScreenLockType() bool {
	if o != nil && o.ScreenLockType != nil {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given DeviceAssuranceAndroidPlatformAllOfScreenLockType and assigns it to the ScreenLockType field.
func (o *DeviceAssuranceWindowsPlatform) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType) {
	o.ScreenLockType = &v
}

// GetSecureHardwarePresent returns the SecureHardwarePresent field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatform) GetSecureHardwarePresent() bool {
	if o == nil || o.SecureHardwarePresent == nil {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatform) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || o.SecureHardwarePresent == nil {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatform) HasSecureHardwarePresent() bool {
	if o != nil && o.SecureHardwarePresent != nil {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *DeviceAssuranceWindowsPlatform) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

// GetThirdPartySignalProviders returns the ThirdPartySignalProviders field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatform) GetThirdPartySignalProviders() DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders {
	if o == nil || o.ThirdPartySignalProviders == nil {
		var ret DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders
		return ret
	}
	return *o.ThirdPartySignalProviders
}

// GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatform) GetThirdPartySignalProvidersOk() (*DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders, bool) {
	if o == nil || o.ThirdPartySignalProviders == nil {
		return nil, false
	}
	return o.ThirdPartySignalProviders, true
}

// HasThirdPartySignalProviders returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatform) HasThirdPartySignalProviders() bool {
	if o != nil && o.ThirdPartySignalProviders != nil {
		return true
	}

	return false
}

// SetThirdPartySignalProviders gets a reference to the given DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders and assigns it to the ThirdPartySignalProviders field.
func (o *DeviceAssuranceWindowsPlatform) SetThirdPartySignalProviders(v DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) {
	o.ThirdPartySignalProviders = &v
}

func (o DeviceAssuranceWindowsPlatform) MarshalJSON() ([]byte, error) {
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
	if o.OsVersionConstraints != nil {
		toSerialize["osVersionConstraints"] = o.OsVersionConstraints
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

func (o *DeviceAssuranceWindowsPlatform) UnmarshalJSON(bytes []byte) (err error) {
	type DeviceAssuranceWindowsPlatformWithoutEmbeddedStruct struct {
		DiskEncryptionType *DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType `json:"diskEncryptionType,omitempty"`
		OsVersion *OSVersionFourComponents `json:"osVersion,omitempty"`
		// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>Specifies the Windows version requirements for the assurance policy. Each requirement must correspond to a different major version (Windows 11 or Windows 10). If a requirement isn't specified for a major version, then devices on that major version satisfy the condition.  There are two types of OS requirements: * **Static**: A specific Windows version requirement that doesn't change until you update the policy. A static OS Windows requirement is specified with `majorVersionConstraint` and `minimum`. * **Dynamic**: A Windows version requirement that is relative to the latest major release and security patch. A dynamic OS Windows requirement is specified with `majorVersionConstraint` and `dynamicVersionRequirement`.  > **Note:** Dynamic OS requirements are available only if the **Dynamic OS version compliance** [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled. The `osVersionConstraints` property is only supported for the Windows platform. You can't specify both `osVersion.minimum` and `osVersionConstraints` properties at the same time. 
		OsVersionConstraints []OSVersionConstraint `json:"osVersionConstraints,omitempty"`
		ScreenLockType *DeviceAssuranceAndroidPlatformAllOfScreenLockType `json:"screenLockType,omitempty"`
		SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
		ThirdPartySignalProviders *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	}

	varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct := DeviceAssuranceWindowsPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct)
	if err == nil {
		varDeviceAssuranceWindowsPlatform := _DeviceAssuranceWindowsPlatform{}
		varDeviceAssuranceWindowsPlatform.DiskEncryptionType = varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct.DiskEncryptionType
		varDeviceAssuranceWindowsPlatform.OsVersion = varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct.OsVersion
		varDeviceAssuranceWindowsPlatform.OsVersionConstraints = varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct.OsVersionConstraints
		varDeviceAssuranceWindowsPlatform.ScreenLockType = varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct.ScreenLockType
		varDeviceAssuranceWindowsPlatform.SecureHardwarePresent = varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct.SecureHardwarePresent
		varDeviceAssuranceWindowsPlatform.ThirdPartySignalProviders = varDeviceAssuranceWindowsPlatformWithoutEmbeddedStruct.ThirdPartySignalProviders
		*o = DeviceAssuranceWindowsPlatform(varDeviceAssuranceWindowsPlatform)
	} else {
		return err
	}

	varDeviceAssuranceWindowsPlatform := _DeviceAssuranceWindowsPlatform{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceWindowsPlatform)
	if err == nil {
		o.DeviceAssurance = varDeviceAssuranceWindowsPlatform.DeviceAssurance
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "diskEncryptionType")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "osVersionConstraints")
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

type NullableDeviceAssuranceWindowsPlatform struct {
	value *DeviceAssuranceWindowsPlatform
	isSet bool
}

func (v NullableDeviceAssuranceWindowsPlatform) Get() *DeviceAssuranceWindowsPlatform {
	return v.value
}

func (v *NullableDeviceAssuranceWindowsPlatform) Set(val *DeviceAssuranceWindowsPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceWindowsPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceWindowsPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceWindowsPlatform(val *DeviceAssuranceWindowsPlatform) *NullableDeviceAssuranceWindowsPlatform {
	return &NullableDeviceAssuranceWindowsPlatform{value: val, isSet: true}
}

func (v NullableDeviceAssuranceWindowsPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceWindowsPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

