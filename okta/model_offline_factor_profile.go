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

// checks if the OfflineFactorProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfflineFactorProfile{}

// OfflineFactorProfile Profile details for an offline factor. The returned properties depend on the factor `type`, and the object is empty for factor types that have no additional detail.
type OfflineFactorProfile struct {
	// Serial number of the factor's hardware. Only returned for `SECURITY_KEY` factors.
	SerialNumber         *string `json:"serialNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OfflineFactorProfile OfflineFactorProfile

// NewOfflineFactorProfile instantiates a new OfflineFactorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfflineFactorProfile() *OfflineFactorProfile {
	this := OfflineFactorProfile{}
	return &this
}

// NewOfflineFactorProfileWithDefaults instantiates a new OfflineFactorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfflineFactorProfileWithDefaults() *OfflineFactorProfile {
	this := OfflineFactorProfile{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OfflineFactorProfile) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineFactorProfile) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OfflineFactorProfile) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OfflineFactorProfile) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o OfflineFactorProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfflineFactorProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OfflineFactorProfile) UnmarshalJSON(data []byte) (err error) {
	varOfflineFactorProfile := _OfflineFactorProfile{}

	err = json.Unmarshal(data, &varOfflineFactorProfile)

	if err != nil {
		return err
	}

	*o = OfflineFactorProfile(varOfflineFactorProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serialNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOfflineFactorProfile struct {
	value *OfflineFactorProfile
	isSet bool
}

func (v NullableOfflineFactorProfile) Get() *OfflineFactorProfile {
	return v.value
}

func (v *NullableOfflineFactorProfile) Set(val *OfflineFactorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOfflineFactorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOfflineFactorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfflineFactorProfile(val *OfflineFactorProfile) *NullableOfflineFactorProfile {
	return &NullableOfflineFactorProfile{value: val, isSet: true}
}

func (v NullableOfflineFactorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfflineFactorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
