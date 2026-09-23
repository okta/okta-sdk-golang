/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ProximityProviderAuthenticatorCharacteristics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProximityProviderAuthenticatorCharacteristics{}

// ProximityProviderAuthenticatorCharacteristics The security characteristics of the proximity provider
type ProximityProviderAuthenticatorCharacteristics struct {
	// Indicates if this provider adds the hardware-protected authenticator characteristic. The hardware-protected characteristic means that authenticators require a physical device to authenticate.
	HardwareProtected    *bool `json:"hardwareProtected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProximityProviderAuthenticatorCharacteristics ProximityProviderAuthenticatorCharacteristics

// NewProximityProviderAuthenticatorCharacteristics instantiates a new ProximityProviderAuthenticatorCharacteristics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProximityProviderAuthenticatorCharacteristics() *ProximityProviderAuthenticatorCharacteristics {
	this := ProximityProviderAuthenticatorCharacteristics{}
	return &this
}

// NewProximityProviderAuthenticatorCharacteristicsWithDefaults instantiates a new ProximityProviderAuthenticatorCharacteristics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProximityProviderAuthenticatorCharacteristicsWithDefaults() *ProximityProviderAuthenticatorCharacteristics {
	this := ProximityProviderAuthenticatorCharacteristics{}
	return &this
}

// GetHardwareProtected returns the HardwareProtected field value if set, zero value otherwise.
func (o *ProximityProviderAuthenticatorCharacteristics) GetHardwareProtected() bool {
	if o == nil || IsNil(o.HardwareProtected) {
		var ret bool
		return ret
	}
	return *o.HardwareProtected
}

// GetHardwareProtectedOk returns a tuple with the HardwareProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProximityProviderAuthenticatorCharacteristics) GetHardwareProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.HardwareProtected) {
		return nil, false
	}
	return o.HardwareProtected, true
}

// HasHardwareProtected returns a boolean if a field has been set.
func (o *ProximityProviderAuthenticatorCharacteristics) HasHardwareProtected() bool {
	if o != nil && !IsNil(o.HardwareProtected) {
		return true
	}

	return false
}

// SetHardwareProtected gets a reference to the given bool and assigns it to the HardwareProtected field.
func (o *ProximityProviderAuthenticatorCharacteristics) SetHardwareProtected(v bool) {
	o.HardwareProtected = &v
}

func (o ProximityProviderAuthenticatorCharacteristics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProximityProviderAuthenticatorCharacteristics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HardwareProtected) {
		toSerialize["hardwareProtected"] = o.HardwareProtected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProximityProviderAuthenticatorCharacteristics) UnmarshalJSON(data []byte) (err error) {
	varProximityProviderAuthenticatorCharacteristics := _ProximityProviderAuthenticatorCharacteristics{}

	err = json.Unmarshal(data, &varProximityProviderAuthenticatorCharacteristics)

	if err != nil {
		return err
	}

	*o = ProximityProviderAuthenticatorCharacteristics(varProximityProviderAuthenticatorCharacteristics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hardwareProtected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProximityProviderAuthenticatorCharacteristics struct {
	value *ProximityProviderAuthenticatorCharacteristics
	isSet bool
}

func (v NullableProximityProviderAuthenticatorCharacteristics) Get() *ProximityProviderAuthenticatorCharacteristics {
	return v.value
}

func (v *NullableProximityProviderAuthenticatorCharacteristics) Set(val *ProximityProviderAuthenticatorCharacteristics) {
	v.value = val
	v.isSet = true
}

func (v NullableProximityProviderAuthenticatorCharacteristics) IsSet() bool {
	return v.isSet
}

func (v *NullableProximityProviderAuthenticatorCharacteristics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProximityProviderAuthenticatorCharacteristics(val *ProximityProviderAuthenticatorCharacteristics) *NullableProximityProviderAuthenticatorCharacteristics {
	return &NullableProximityProviderAuthenticatorCharacteristics{value: val, isSet: true}
}

func (v NullableProximityProviderAuthenticatorCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProximityProviderAuthenticatorCharacteristics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
