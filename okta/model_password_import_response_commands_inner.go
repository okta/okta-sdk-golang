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

// PasswordImportResponseCommandsInner struct for PasswordImportResponseCommandsInner
type PasswordImportResponseCommandsInner struct {
	// The location where you specify the command. For the password import inline hook, there's only one command, `com.okta.action.update`.
	Type *string `json:"type,omitempty"`
	Value *PasswordImportResponseCommandsInnerValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportResponseCommandsInner PasswordImportResponseCommandsInner

// NewPasswordImportResponseCommandsInner instantiates a new PasswordImportResponseCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportResponseCommandsInner() *PasswordImportResponseCommandsInner {
	this := PasswordImportResponseCommandsInner{}
	return &this
}

// NewPasswordImportResponseCommandsInnerWithDefaults instantiates a new PasswordImportResponseCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportResponseCommandsInnerWithDefaults() *PasswordImportResponseCommandsInner {
	this := PasswordImportResponseCommandsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PasswordImportResponseCommandsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportResponseCommandsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PasswordImportResponseCommandsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PasswordImportResponseCommandsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PasswordImportResponseCommandsInner) GetValue() PasswordImportResponseCommandsInnerValue {
	if o == nil || o.Value == nil {
		var ret PasswordImportResponseCommandsInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportResponseCommandsInner) GetValueOk() (*PasswordImportResponseCommandsInnerValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PasswordImportResponseCommandsInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given PasswordImportResponseCommandsInnerValue and assigns it to the Value field.
func (o *PasswordImportResponseCommandsInner) SetValue(v PasswordImportResponseCommandsInnerValue) {
	o.Value = &v
}

func (o PasswordImportResponseCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordImportResponseCommandsInner) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordImportResponseCommandsInner := _PasswordImportResponseCommandsInner{}

	err = json.Unmarshal(bytes, &varPasswordImportResponseCommandsInner)
	if err == nil {
		*o = PasswordImportResponseCommandsInner(varPasswordImportResponseCommandsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordImportResponseCommandsInner struct {
	value *PasswordImportResponseCommandsInner
	isSet bool
}

func (v NullablePasswordImportResponseCommandsInner) Get() *PasswordImportResponseCommandsInner {
	return v.value
}

func (v *NullablePasswordImportResponseCommandsInner) Set(val *PasswordImportResponseCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportResponseCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportResponseCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportResponseCommandsInner(val *PasswordImportResponseCommandsInner) *NullablePasswordImportResponseCommandsInner {
	return &NullablePasswordImportResponseCommandsInner{value: val, isSet: true}
}

func (v NullablePasswordImportResponseCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportResponseCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

