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

// DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders Settings for third-party signal providers (based on the `MACOS` platform)
type DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders struct {
	Dtc *DTCMacOS `json:"dtc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders

// NewDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders instantiates a new DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders() *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// NewDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProvidersWithDefaults instantiates a new DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProvidersWithDefaults() *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// GetDtc returns the Dtc field value if set, zero value otherwise.
func (o *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) GetDtc() DTCMacOS {
	if o == nil || o.Dtc == nil {
		var ret DTCMacOS
		return ret
	}
	return *o.Dtc
}

// GetDtcOk returns a tuple with the Dtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) GetDtcOk() (*DTCMacOS, bool) {
	if o == nil || o.Dtc == nil {
		return nil, false
	}
	return o.Dtc, true
}

// HasDtc returns a boolean if a field has been set.
func (o *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) HasDtc() bool {
	if o != nil && o.Dtc != nil {
		return true
	}

	return false
}

// SetDtc gets a reference to the given DTCMacOS and assigns it to the Dtc field.
func (o *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) SetDtc(v DTCMacOS) {
	o.Dtc = &v
}

func (o DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dtc != nil {
		toSerialize["dtc"] = o.Dtc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders := _DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders)
	if err == nil {
		*o = DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders(varDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "dtc")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders struct {
	value *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders
	isSet bool
}

func (v NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) Get() *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders {
	return v.value
}

func (v *NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) Set(val *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders(val *DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) *NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders {
	return &NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders{value: val, isSet: true}
}

func (v NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

