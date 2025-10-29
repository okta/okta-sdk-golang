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

// checks if the AAGUIDAuthenticatorCharacteristics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AAGUIDAuthenticatorCharacteristics{}

// AAGUIDAuthenticatorCharacteristics Contains additional properties about custom AAGUID.
type AAGUIDAuthenticatorCharacteristics struct {
	// Indicates whether the authenticator meets Federal Information Processing Standards (FIPS) compliance requirements
	FipsCompliant *bool `json:"fipsCompliant,omitempty"`
	// Indicates whether the authenticator stores the private key on a hardware component
	HardwareProtected *bool `json:"hardwareProtected,omitempty"`
	// Indicates whether the custom AAGUID is built into the authenticator (`true`) or if it's a separate, external authenticator
	PlatformAttached     *bool `json:"platformAttached,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AAGUIDAuthenticatorCharacteristics AAGUIDAuthenticatorCharacteristics

// NewAAGUIDAuthenticatorCharacteristics instantiates a new AAGUIDAuthenticatorCharacteristics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAAGUIDAuthenticatorCharacteristics() *AAGUIDAuthenticatorCharacteristics {
	this := AAGUIDAuthenticatorCharacteristics{}
	return &this
}

// NewAAGUIDAuthenticatorCharacteristicsWithDefaults instantiates a new AAGUIDAuthenticatorCharacteristics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAAGUIDAuthenticatorCharacteristicsWithDefaults() *AAGUIDAuthenticatorCharacteristics {
	this := AAGUIDAuthenticatorCharacteristics{}
	return &this
}

// GetFipsCompliant returns the FipsCompliant field value if set, zero value otherwise.
func (o *AAGUIDAuthenticatorCharacteristics) GetFipsCompliant() bool {
	if o == nil || IsNil(o.FipsCompliant) {
		var ret bool
		return ret
	}
	return *o.FipsCompliant
}

// GetFipsCompliantOk returns a tuple with the FipsCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AAGUIDAuthenticatorCharacteristics) GetFipsCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.FipsCompliant) {
		return nil, false
	}
	return o.FipsCompliant, true
}

// HasFipsCompliant returns a boolean if a field has been set.
func (o *AAGUIDAuthenticatorCharacteristics) HasFipsCompliant() bool {
	if o != nil && !IsNil(o.FipsCompliant) {
		return true
	}

	return false
}

// SetFipsCompliant gets a reference to the given bool and assigns it to the FipsCompliant field.
func (o *AAGUIDAuthenticatorCharacteristics) SetFipsCompliant(v bool) {
	o.FipsCompliant = &v
}

// GetHardwareProtected returns the HardwareProtected field value if set, zero value otherwise.
func (o *AAGUIDAuthenticatorCharacteristics) GetHardwareProtected() bool {
	if o == nil || IsNil(o.HardwareProtected) {
		var ret bool
		return ret
	}
	return *o.HardwareProtected
}

// GetHardwareProtectedOk returns a tuple with the HardwareProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AAGUIDAuthenticatorCharacteristics) GetHardwareProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.HardwareProtected) {
		return nil, false
	}
	return o.HardwareProtected, true
}

// HasHardwareProtected returns a boolean if a field has been set.
func (o *AAGUIDAuthenticatorCharacteristics) HasHardwareProtected() bool {
	if o != nil && !IsNil(o.HardwareProtected) {
		return true
	}

	return false
}

// SetHardwareProtected gets a reference to the given bool and assigns it to the HardwareProtected field.
func (o *AAGUIDAuthenticatorCharacteristics) SetHardwareProtected(v bool) {
	o.HardwareProtected = &v
}

// GetPlatformAttached returns the PlatformAttached field value if set, zero value otherwise.
func (o *AAGUIDAuthenticatorCharacteristics) GetPlatformAttached() bool {
	if o == nil || IsNil(o.PlatformAttached) {
		var ret bool
		return ret
	}
	return *o.PlatformAttached
}

// GetPlatformAttachedOk returns a tuple with the PlatformAttached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AAGUIDAuthenticatorCharacteristics) GetPlatformAttachedOk() (*bool, bool) {
	if o == nil || IsNil(o.PlatformAttached) {
		return nil, false
	}
	return o.PlatformAttached, true
}

// HasPlatformAttached returns a boolean if a field has been set.
func (o *AAGUIDAuthenticatorCharacteristics) HasPlatformAttached() bool {
	if o != nil && !IsNil(o.PlatformAttached) {
		return true
	}

	return false
}

// SetPlatformAttached gets a reference to the given bool and assigns it to the PlatformAttached field.
func (o *AAGUIDAuthenticatorCharacteristics) SetPlatformAttached(v bool) {
	o.PlatformAttached = &v
}

func (o AAGUIDAuthenticatorCharacteristics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AAGUIDAuthenticatorCharacteristics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FipsCompliant) {
		toSerialize["fipsCompliant"] = o.FipsCompliant
	}
	if !IsNil(o.HardwareProtected) {
		toSerialize["hardwareProtected"] = o.HardwareProtected
	}
	if !IsNil(o.PlatformAttached) {
		toSerialize["platformAttached"] = o.PlatformAttached
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AAGUIDAuthenticatorCharacteristics) UnmarshalJSON(data []byte) (err error) {
	varAAGUIDAuthenticatorCharacteristics := _AAGUIDAuthenticatorCharacteristics{}

	err = json.Unmarshal(data, &varAAGUIDAuthenticatorCharacteristics)

	if err != nil {
		return err
	}

	*o = AAGUIDAuthenticatorCharacteristics(varAAGUIDAuthenticatorCharacteristics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fipsCompliant")
		delete(additionalProperties, "hardwareProtected")
		delete(additionalProperties, "platformAttached")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAAGUIDAuthenticatorCharacteristics struct {
	value *AAGUIDAuthenticatorCharacteristics
	isSet bool
}

func (v NullableAAGUIDAuthenticatorCharacteristics) Get() *AAGUIDAuthenticatorCharacteristics {
	return v.value
}

func (v *NullableAAGUIDAuthenticatorCharacteristics) Set(val *AAGUIDAuthenticatorCharacteristics) {
	v.value = val
	v.isSet = true
}

func (v NullableAAGUIDAuthenticatorCharacteristics) IsSet() bool {
	return v.isSet
}

func (v *NullableAAGUIDAuthenticatorCharacteristics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAAGUIDAuthenticatorCharacteristics(val *AAGUIDAuthenticatorCharacteristics) *NullableAAGUIDAuthenticatorCharacteristics {
	return &NullableAAGUIDAuthenticatorCharacteristics{value: val, isSet: true}
}

func (v NullableAAGUIDAuthenticatorCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAAGUIDAuthenticatorCharacteristics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
