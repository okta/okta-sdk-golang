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

// checks if the DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders{}

// DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders Settings for third-party signal providers (based on the `IOS` platform)
type DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders struct {
	DevicePostureIdP     *DevicePostureIdP `json:"devicePostureIdP,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders

// NewDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders instantiates a new DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders() *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// NewDeviceAssuranceIOSPlatformAllOfThirdPartySignalProvidersWithDefaults instantiates a new DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceIOSPlatformAllOfThirdPartySignalProvidersWithDefaults() *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// GetDevicePostureIdP returns the DevicePostureIdP field value if set, zero value otherwise.
func (o *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) GetDevicePostureIdP() DevicePostureIdP {
	if o == nil || IsNil(o.DevicePostureIdP) {
		var ret DevicePostureIdP
		return ret
	}
	return *o.DevicePostureIdP
}

// GetDevicePostureIdPOk returns a tuple with the DevicePostureIdP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) GetDevicePostureIdPOk() (*DevicePostureIdP, bool) {
	if o == nil || IsNil(o.DevicePostureIdP) {
		return nil, false
	}
	return o.DevicePostureIdP, true
}

// HasDevicePostureIdP returns a boolean if a field has been set.
func (o *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) HasDevicePostureIdP() bool {
	if o != nil && !IsNil(o.DevicePostureIdP) {
		return true
	}

	return false
}

// SetDevicePostureIdP gets a reference to the given DevicePostureIdP and assigns it to the DevicePostureIdP field.
func (o *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) SetDevicePostureIdP(v DevicePostureIdP) {
	o.DevicePostureIdP = &v
}

func (o DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DevicePostureIdP) {
		toSerialize["devicePostureIdP"] = o.DevicePostureIdP
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(data []byte) (err error) {
	varDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders := _DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders{}

	err = json.Unmarshal(data, &varDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders)

	if err != nil {
		return err
	}

	*o = DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders(varDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "devicePostureIdP")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders struct {
	value *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders
	isSet bool
}

func (v NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) Get() *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders {
	return v.value
}

func (v *NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) Set(val *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders(val *DeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) *NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders {
	return &NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders{value: val, isSet: true}
}

func (v NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceIOSPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
