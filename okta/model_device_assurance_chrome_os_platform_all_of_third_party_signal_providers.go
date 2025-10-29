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

// checks if the DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders{}

// DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders Settings for third-party signal providers (based on the `CHROMEOS` platform)
type DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders struct {
	Dtc                  *DTCChromeOS      `json:"dtc,omitempty"`
	DevicePostureIdP     *DevicePostureIdP `json:"devicePostureIdP,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders

// NewDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders instantiates a new DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders() *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// NewDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProvidersWithDefaults instantiates a new DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProvidersWithDefaults() *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// GetDtc returns the Dtc field value if set, zero value otherwise.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) GetDtc() DTCChromeOS {
	if o == nil || IsNil(o.Dtc) {
		var ret DTCChromeOS
		return ret
	}
	return *o.Dtc
}

// GetDtcOk returns a tuple with the Dtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) GetDtcOk() (*DTCChromeOS, bool) {
	if o == nil || IsNil(o.Dtc) {
		return nil, false
	}
	return o.Dtc, true
}

// HasDtc returns a boolean if a field has been set.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) HasDtc() bool {
	if o != nil && !IsNil(o.Dtc) {
		return true
	}

	return false
}

// SetDtc gets a reference to the given DTCChromeOS and assigns it to the Dtc field.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) SetDtc(v DTCChromeOS) {
	o.Dtc = &v
}

// GetDevicePostureIdP returns the DevicePostureIdP field value if set, zero value otherwise.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) GetDevicePostureIdP() DevicePostureIdP {
	if o == nil || IsNil(o.DevicePostureIdP) {
		var ret DevicePostureIdP
		return ret
	}
	return *o.DevicePostureIdP
}

// GetDevicePostureIdPOk returns a tuple with the DevicePostureIdP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) GetDevicePostureIdPOk() (*DevicePostureIdP, bool) {
	if o == nil || IsNil(o.DevicePostureIdP) {
		return nil, false
	}
	return o.DevicePostureIdP, true
}

// HasDevicePostureIdP returns a boolean if a field has been set.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) HasDevicePostureIdP() bool {
	if o != nil && !IsNil(o.DevicePostureIdP) {
		return true
	}

	return false
}

// SetDevicePostureIdP gets a reference to the given DevicePostureIdP and assigns it to the DevicePostureIdP field.
func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) SetDevicePostureIdP(v DevicePostureIdP) {
	o.DevicePostureIdP = &v
}

func (o DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dtc) {
		toSerialize["dtc"] = o.Dtc
	}
	if !IsNil(o.DevicePostureIdP) {
		toSerialize["devicePostureIdP"] = o.DevicePostureIdP
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(data []byte) (err error) {
	varDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders := _DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders{}

	err = json.Unmarshal(data, &varDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders)

	if err != nil {
		return err
	}

	*o = DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders(varDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dtc")
		delete(additionalProperties, "devicePostureIdP")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders struct {
	value *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders
	isSet bool
}

func (v NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) Get() *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders {
	return v.value
}

func (v *NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) Set(val *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders(val *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) *NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders {
	return &NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders{value: val, isSet: true}
}

func (v NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
