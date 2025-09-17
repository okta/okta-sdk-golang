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

// checks if the SAMLHookResponseCommandsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLHookResponseCommandsInner{}

// SAMLHookResponseCommandsInner struct for SAMLHookResponseCommandsInner
type SAMLHookResponseCommandsInner struct {
	// One of the supported commands `com.okta.assertion.patch`
	Type                 *string                                   `json:"type,omitempty"`
	Value                []SAMLHookResponseCommandsInnerValueInner `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLHookResponseCommandsInner SAMLHookResponseCommandsInner

// NewSAMLHookResponseCommandsInner instantiates a new SAMLHookResponseCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLHookResponseCommandsInner() *SAMLHookResponseCommandsInner {
	this := SAMLHookResponseCommandsInner{}
	return &this
}

// NewSAMLHookResponseCommandsInnerWithDefaults instantiates a new SAMLHookResponseCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLHookResponseCommandsInnerWithDefaults() *SAMLHookResponseCommandsInner {
	this := SAMLHookResponseCommandsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SAMLHookResponseCommandsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponseCommandsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SAMLHookResponseCommandsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SAMLHookResponseCommandsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SAMLHookResponseCommandsInner) GetValue() []SAMLHookResponseCommandsInnerValueInner {
	if o == nil || IsNil(o.Value) {
		var ret []SAMLHookResponseCommandsInnerValueInner
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponseCommandsInner) GetValueOk() ([]SAMLHookResponseCommandsInnerValueInner, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SAMLHookResponseCommandsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []SAMLHookResponseCommandsInnerValueInner and assigns it to the Value field.
func (o *SAMLHookResponseCommandsInner) SetValue(v []SAMLHookResponseCommandsInnerValueInner) {
	o.Value = v
}

func (o SAMLHookResponseCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLHookResponseCommandsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLHookResponseCommandsInner) UnmarshalJSON(data []byte) (err error) {
	varSAMLHookResponseCommandsInner := _SAMLHookResponseCommandsInner{}

	err = json.Unmarshal(data, &varSAMLHookResponseCommandsInner)

	if err != nil {
		return err
	}

	*o = SAMLHookResponseCommandsInner(varSAMLHookResponseCommandsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLHookResponseCommandsInner struct {
	value *SAMLHookResponseCommandsInner
	isSet bool
}

func (v NullableSAMLHookResponseCommandsInner) Get() *SAMLHookResponseCommandsInner {
	return v.value
}

func (v *NullableSAMLHookResponseCommandsInner) Set(val *SAMLHookResponseCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLHookResponseCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLHookResponseCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLHookResponseCommandsInner(val *SAMLHookResponseCommandsInner) *NullableSAMLHookResponseCommandsInner {
	return &NullableSAMLHookResponseCommandsInner{value: val, isSet: true}
}

func (v NullableSAMLHookResponseCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLHookResponseCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
