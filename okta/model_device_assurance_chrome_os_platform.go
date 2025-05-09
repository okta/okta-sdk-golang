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

// DeviceAssuranceChromeOSPlatform struct for DeviceAssuranceChromeOSPlatform
type DeviceAssuranceChromeOSPlatform struct {
	DeviceAssurance
	ThirdPartySignalProviders *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAssuranceChromeOSPlatform DeviceAssuranceChromeOSPlatform

// NewDeviceAssuranceChromeOSPlatform instantiates a new DeviceAssuranceChromeOSPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAssuranceChromeOSPlatform() *DeviceAssuranceChromeOSPlatform {
	this := DeviceAssuranceChromeOSPlatform{}
	return &this
}

// NewDeviceAssuranceChromeOSPlatformWithDefaults instantiates a new DeviceAssuranceChromeOSPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAssuranceChromeOSPlatformWithDefaults() *DeviceAssuranceChromeOSPlatform {
	this := DeviceAssuranceChromeOSPlatform{}
	return &this
}

// GetThirdPartySignalProviders returns the ThirdPartySignalProviders field value if set, zero value otherwise.
func (o *DeviceAssuranceChromeOSPlatform) GetThirdPartySignalProviders() DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders {
	if o == nil || o.ThirdPartySignalProviders == nil {
		var ret DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders
		return ret
	}
	return *o.ThirdPartySignalProviders
}

// GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAssuranceChromeOSPlatform) GetThirdPartySignalProvidersOk() (*DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders, bool) {
	if o == nil || o.ThirdPartySignalProviders == nil {
		return nil, false
	}
	return o.ThirdPartySignalProviders, true
}

// HasThirdPartySignalProviders returns a boolean if a field has been set.
func (o *DeviceAssuranceChromeOSPlatform) HasThirdPartySignalProviders() bool {
	if o != nil && o.ThirdPartySignalProviders != nil {
		return true
	}

	return false
}

// SetThirdPartySignalProviders gets a reference to the given DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders and assigns it to the ThirdPartySignalProviders field.
func (o *DeviceAssuranceChromeOSPlatform) SetThirdPartySignalProviders(v DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders) {
	o.ThirdPartySignalProviders = &v
}

func (o DeviceAssuranceChromeOSPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedDeviceAssurance, errDeviceAssurance := json.Marshal(o.DeviceAssurance)
	if errDeviceAssurance != nil {
		return []byte{}, errDeviceAssurance
	}
	errDeviceAssurance = json.Unmarshal([]byte(serializedDeviceAssurance), &toSerialize)
	if errDeviceAssurance != nil {
		return []byte{}, errDeviceAssurance
	}
	if o.ThirdPartySignalProviders != nil {
		toSerialize["thirdPartySignalProviders"] = o.ThirdPartySignalProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAssuranceChromeOSPlatform) UnmarshalJSON(bytes []byte) (err error) {
	type DeviceAssuranceChromeOSPlatformWithoutEmbeddedStruct struct {
		ThirdPartySignalProviders *DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders `json:"thirdPartySignalProviders,omitempty"`
	}

	varDeviceAssuranceChromeOSPlatformWithoutEmbeddedStruct := DeviceAssuranceChromeOSPlatformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceChromeOSPlatformWithoutEmbeddedStruct)
	if err == nil {
		varDeviceAssuranceChromeOSPlatform := _DeviceAssuranceChromeOSPlatform{}
		varDeviceAssuranceChromeOSPlatform.ThirdPartySignalProviders = varDeviceAssuranceChromeOSPlatformWithoutEmbeddedStruct.ThirdPartySignalProviders
		*o = DeviceAssuranceChromeOSPlatform(varDeviceAssuranceChromeOSPlatform)
	} else {
		return err
	}

	varDeviceAssuranceChromeOSPlatform := _DeviceAssuranceChromeOSPlatform{}

	err = json.Unmarshal(bytes, &varDeviceAssuranceChromeOSPlatform)
	if err == nil {
		o.DeviceAssurance = varDeviceAssuranceChromeOSPlatform.DeviceAssurance
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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

type NullableDeviceAssuranceChromeOSPlatform struct {
	value *DeviceAssuranceChromeOSPlatform
	isSet bool
}

func (v NullableDeviceAssuranceChromeOSPlatform) Get() *DeviceAssuranceChromeOSPlatform {
	return v.value
}

func (v *NullableDeviceAssuranceChromeOSPlatform) Set(val *DeviceAssuranceChromeOSPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAssuranceChromeOSPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAssuranceChromeOSPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAssuranceChromeOSPlatform(val *DeviceAssuranceChromeOSPlatform) *NullableDeviceAssuranceChromeOSPlatform {
	return &NullableDeviceAssuranceChromeOSPlatform{value: val, isSet: true}
}

func (v NullableDeviceAssuranceChromeOSPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAssuranceChromeOSPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

