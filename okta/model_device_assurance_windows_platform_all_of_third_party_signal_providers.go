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

// DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders Settings for third-party signal providers (based on the `WINDOWS` platform)
type DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders struct {
	Dtc *DTCWindows `json:"dtc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders

// NewDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders instantiates a new DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders() *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// NewDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProvidersWithDefaults instantiates a new DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProvidersWithDefaults() *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders {
	this := DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders{}
	return &this
}

// GetDtc returns the Dtc field value if set, zero value otherwise.
func (o *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) GetDtc() DTCWindows {
	if o == nil || o.Dtc == nil {
		var ret DTCWindows
		return ret
	}
	return *o.Dtc
}

// GetDtcOk returns a tuple with the Dtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) GetDtcOk() (*DTCWindows, bool) {
	if o == nil || o.Dtc == nil {
		return nil, false
	}
	return o.Dtc, true
}

// HasDtc returns a boolean if a field has been set.
func (o *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) HasDtc() bool {
	if o != nil && o.Dtc != nil {
		return true
	}

	return false
}

// SetDtc gets a reference to the given DTCWindows and assigns it to the Dtc field.
func (o *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) SetDtc(v DTCWindows) {
	o.Dtc = &v
}

func (o DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dtc != nil {
		toSerialize["dtc"] = o.Dtc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders := _DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders)
	if err == nil {
		*o = DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders(varDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders)
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

type NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders struct {
	value *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders
	isSet bool
}

func (v NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) Get() *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders {
	return v.value
}

func (v *NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) Set(val *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders(val *DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) *NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders {
	return &NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders{value: val, isSet: true}
}

func (v NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

