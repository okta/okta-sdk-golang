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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordDictionaryCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordDictionaryCommon{}

// PasswordDictionaryCommon Lookup settings for commonly used passwords
type PasswordDictionaryCommon struct {
	// Indicates whether to check passwords against the common password dictionary
	Exclude              *bool `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordDictionaryCommon PasswordDictionaryCommon

// NewPasswordDictionaryCommon instantiates a new PasswordDictionaryCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordDictionaryCommon() *PasswordDictionaryCommon {
	this := PasswordDictionaryCommon{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// NewPasswordDictionaryCommonWithDefaults instantiates a new PasswordDictionaryCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordDictionaryCommonWithDefaults() *PasswordDictionaryCommon {
	this := PasswordDictionaryCommon{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *PasswordDictionaryCommon) GetExclude() bool {
	if o == nil || IsNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDictionaryCommon) GetExcludeOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *PasswordDictionaryCommon) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *PasswordDictionaryCommon) SetExclude(v bool) {
	o.Exclude = &v
}

func (o PasswordDictionaryCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordDictionaryCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordDictionaryCommon) UnmarshalJSON(data []byte) (err error) {
	varPasswordDictionaryCommon := _PasswordDictionaryCommon{}

	err = json.Unmarshal(data, &varPasswordDictionaryCommon)

	if err != nil {
		return err
	}

	*o = PasswordDictionaryCommon(varPasswordDictionaryCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordDictionaryCommon struct {
	value *PasswordDictionaryCommon
	isSet bool
}

func (v NullablePasswordDictionaryCommon) Get() *PasswordDictionaryCommon {
	return v.value
}

func (v *NullablePasswordDictionaryCommon) Set(val *PasswordDictionaryCommon) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordDictionaryCommon) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordDictionaryCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordDictionaryCommon(val *PasswordDictionaryCommon) *NullablePasswordDictionaryCommon {
	return &NullablePasswordDictionaryCommon{value: val, isSet: true}
}

func (v NullablePasswordDictionaryCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordDictionaryCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
