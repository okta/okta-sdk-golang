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

// checks if the DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders{}

// DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders Settings for third-party signal providers (based on the `ANDROID` platform)
type DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders struct {
	AndroidDeviceTrust   *AndroidDeviceTrust `json:"androidDeviceTrust,omitempty"`
	DevicePostureIdP     *DevicePostureIdP   `json:"devicePostureIdP,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders

// NewDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders instantiates a new DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders() *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// NewDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProvidersWithDefaults instantiates a new DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProvidersWithDefaults() *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// GetAndroidDeviceTrust returns the AndroidDeviceTrust field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) GetAndroidDeviceTrust() AndroidDeviceTrust {
	if o == nil || IsNil(o.AndroidDeviceTrust) {
		var ret AndroidDeviceTrust
		return ret
	}
	return *o.AndroidDeviceTrust
}

// GetAndroidDeviceTrustOk returns a tuple with the AndroidDeviceTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) GetAndroidDeviceTrustOk() (*AndroidDeviceTrust, bool) {
	if o == nil || IsNil(o.AndroidDeviceTrust) {
		return nil, false
	}
	return o.AndroidDeviceTrust, true
}

// HasAndroidDeviceTrust returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) HasAndroidDeviceTrust() bool {
	if o != nil && !IsNil(o.AndroidDeviceTrust) {
		return true
	}

	return false
}

// SetAndroidDeviceTrust gets a reference to the given AndroidDeviceTrust and assigns it to the AndroidDeviceTrust field.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) SetAndroidDeviceTrust(v AndroidDeviceTrust) {
	o.AndroidDeviceTrust = &v
}

// GetDevicePostureIdP returns the DevicePostureIdP field value if set, zero value otherwise.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) GetDevicePostureIdP() DevicePostureIdP {
	if o == nil || IsNil(o.DevicePostureIdP) {
		var ret DevicePostureIdP
		return ret
	}
	return *o.DevicePostureIdP
}

// GetDevicePostureIdPOk returns a tuple with the DevicePostureIdP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) GetDevicePostureIdPOk() (*DevicePostureIdP, bool) {
	if o == nil || IsNil(o.DevicePostureIdP) {
		return nil, false
	}
	return o.DevicePostureIdP, true
}

// HasDevicePostureIdP returns a boolean if a field has been set.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) HasDevicePostureIdP() bool {
	if o != nil && !IsNil(o.DevicePostureIdP) {
		return true
	}

	return false
}

// SetDevicePostureIdP gets a reference to the given DevicePostureIdP and assigns it to the DevicePostureIdP field.
func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) SetDevicePostureIdP(v DevicePostureIdP) {
	o.DevicePostureIdP = &v
}

func (o DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AndroidDeviceTrust) {
		toSerialize["androidDeviceTrust"] = o.AndroidDeviceTrust
	}
	if !IsNil(o.DevicePostureIdP) {
		toSerialize["devicePostureIdP"] = o.DevicePostureIdP
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(data []byte) (err error) {
	varDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders := _DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders{}

	err = json.Unmarshal(data, &varDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders)

	if err != nil {
		return err
	}

	*o = DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders(varDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "androidDeviceTrust")
		delete(additionalProperties, "devicePostureIdP")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders struct {
	value *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders
	isSet bool
}

func (v NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) Get() *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders {
	return v.value
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) Set(val *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders(val *DeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) *NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders {
	return &NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders{value: val, isSet: true}
}

func (v NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceAndroidPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
