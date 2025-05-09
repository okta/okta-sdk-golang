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

// TempPassword struct for TempPassword
type TempPassword struct {
	TempPassword *string `json:"tempPassword,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TempPassword TempPassword

// NewTempPassword instantiates a new TempPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTempPassword() *TempPassword {
	this := TempPassword{}
	return &this
}

// NewTempPasswordWithDefaults instantiates a new TempPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTempPasswordWithDefaults() *TempPassword {
	this := TempPassword{}
	return &this
}

// GetTempPassword returns the TempPassword field value if set, zero value otherwise.
func (o *TempPassword) GetTempPassword() string {
	if o == nil || o.TempPassword == nil {
		var ret string
		return ret
	}
	return *o.TempPassword
}

// GetTempPasswordOk returns a tuple with the TempPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TempPassword) GetTempPasswordOk() (*string, bool) {
	if o == nil || o.TempPassword == nil {
		return nil, false
	}
	return o.TempPassword, true
}

// HasTempPassword returns a boolean if a field has been set.
func (o *TempPassword) HasTempPassword() bool {
	if o != nil && o.TempPassword != nil {
		return true
	}

	return false
}

// SetTempPassword gets a reference to the given string and assigns it to the TempPassword field.
func (o *TempPassword) SetTempPassword(v string) {
	o.TempPassword = &v
}

func (o TempPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TempPassword != nil {
		toSerialize["tempPassword"] = o.TempPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TempPassword) UnmarshalJSON(bytes []byte) (err error) {
	varTempPassword := _TempPassword{}

	err = json.Unmarshal(bytes, &varTempPassword)
	if err == nil {
		*o = TempPassword(varTempPassword)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "tempPassword")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTempPassword struct {
	value *TempPassword
	isSet bool
}

func (v NullableTempPassword) Get() *TempPassword {
	return v.value
}

func (v *NullableTempPassword) Set(val *TempPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableTempPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableTempPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTempPassword(val *TempPassword) *NullableTempPassword {
	return &NullableTempPassword{value: val, isSet: true}
}

func (v NullableTempPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTempPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

