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

// checks if the PasswordDictionary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordDictionary{}

// PasswordDictionary Weak password dictionary lookup settings
type PasswordDictionary struct {
	Common               *PasswordDictionaryCommon `json:"common,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordDictionary PasswordDictionary

// NewPasswordDictionary instantiates a new PasswordDictionary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordDictionary() *PasswordDictionary {
	this := PasswordDictionary{}
	return &this
}

// NewPasswordDictionaryWithDefaults instantiates a new PasswordDictionary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordDictionaryWithDefaults() *PasswordDictionary {
	this := PasswordDictionary{}
	return &this
}

// GetCommon returns the Common field value if set, zero value otherwise.
func (o *PasswordDictionary) GetCommon() PasswordDictionaryCommon {
	if o == nil || IsNil(o.Common) {
		var ret PasswordDictionaryCommon
		return ret
	}
	return *o.Common
}

// GetCommonOk returns a tuple with the Common field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDictionary) GetCommonOk() (*PasswordDictionaryCommon, bool) {
	if o == nil || IsNil(o.Common) {
		return nil, false
	}
	return o.Common, true
}

// HasCommon returns a boolean if a field has been set.
func (o *PasswordDictionary) HasCommon() bool {
	if o != nil && !IsNil(o.Common) {
		return true
	}

	return false
}

// SetCommon gets a reference to the given PasswordDictionaryCommon and assigns it to the Common field.
func (o *PasswordDictionary) SetCommon(v PasswordDictionaryCommon) {
	o.Common = &v
}

func (o PasswordDictionary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordDictionary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Common) {
		toSerialize["common"] = o.Common
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordDictionary) UnmarshalJSON(data []byte) (err error) {
	varPasswordDictionary := _PasswordDictionary{}

	err = json.Unmarshal(data, &varPasswordDictionary)

	if err != nil {
		return err
	}

	*o = PasswordDictionary(varPasswordDictionary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "common")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordDictionary struct {
	value *PasswordDictionary
	isSet bool
}

func (v NullablePasswordDictionary) Get() *PasswordDictionary {
	return v.value
}

func (v *NullablePasswordDictionary) Set(val *PasswordDictionary) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordDictionary) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordDictionary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordDictionary(val *PasswordDictionary) *NullablePasswordDictionary {
	return &NullablePasswordDictionary{value: val, isSet: true}
}

func (v NullablePasswordDictionary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordDictionary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
