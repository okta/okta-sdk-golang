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

// PasswordDictionaryCommon struct for PasswordDictionaryCommon
type PasswordDictionaryCommon struct {
	Exclude *bool `json:"exclude,omitempty"`
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
	if o == nil || o.Exclude == nil {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordDictionaryCommon) GetExcludeOk() (*bool, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *PasswordDictionaryCommon) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *PasswordDictionaryCommon) SetExclude(v bool) {
	o.Exclude = &v
}

func (o PasswordDictionaryCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordDictionaryCommon) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordDictionaryCommon := _PasswordDictionaryCommon{}

	err = json.Unmarshal(bytes, &varPasswordDictionaryCommon)
	if err == nil {
		*o = PasswordDictionaryCommon(varPasswordDictionaryCommon)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

