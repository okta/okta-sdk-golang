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

// checks if the RegistrationAuthorityList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationAuthorityList{}

// RegistrationAuthorityList A collection of registration authority configurations
type RegistrationAuthorityList struct {
	// The configurations, grouped by the certificate they're bound to, newest certificate first
	Data                 []RegistrationAuthority `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationAuthorityList RegistrationAuthorityList

// NewRegistrationAuthorityList instantiates a new RegistrationAuthorityList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationAuthorityList() *RegistrationAuthorityList {
	this := RegistrationAuthorityList{}
	return &this
}

// NewRegistrationAuthorityListWithDefaults instantiates a new RegistrationAuthorityList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationAuthorityListWithDefaults() *RegistrationAuthorityList {
	this := RegistrationAuthorityList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RegistrationAuthorityList) GetData() []RegistrationAuthority {
	if o == nil || IsNil(o.Data) {
		var ret []RegistrationAuthority
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthorityList) GetDataOk() ([]RegistrationAuthority, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RegistrationAuthorityList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RegistrationAuthority and assigns it to the Data field.
func (o *RegistrationAuthorityList) SetData(v []RegistrationAuthority) {
	o.Data = v
}

func (o RegistrationAuthorityList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationAuthorityList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationAuthorityList) UnmarshalJSON(data []byte) (err error) {
	varRegistrationAuthorityList := _RegistrationAuthorityList{}

	err = json.Unmarshal(data, &varRegistrationAuthorityList)

	if err != nil {
		return err
	}

	*o = RegistrationAuthorityList(varRegistrationAuthorityList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationAuthorityList struct {
	value *RegistrationAuthorityList
	isSet bool
}

func (v NullableRegistrationAuthorityList) Get() *RegistrationAuthorityList {
	return v.value
}

func (v *NullableRegistrationAuthorityList) Set(val *RegistrationAuthorityList) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationAuthorityList) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationAuthorityList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationAuthorityList(val *RegistrationAuthorityList) *NullableRegistrationAuthorityList {
	return &NullableRegistrationAuthorityList{value: val, isSet: true}
}

func (v NullableRegistrationAuthorityList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationAuthorityList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
