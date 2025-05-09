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

// InlineHookResponseCommands struct for InlineHookResponseCommands
type InlineHookResponseCommands struct {
	Type *string `json:"type,omitempty"`
	Value []InlineHookResponseCommandValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookResponseCommands InlineHookResponseCommands

// NewInlineHookResponseCommands instantiates a new InlineHookResponseCommands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookResponseCommands() *InlineHookResponseCommands {
	this := InlineHookResponseCommands{}
	return &this
}

// NewInlineHookResponseCommandsWithDefaults instantiates a new InlineHookResponseCommands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookResponseCommandsWithDefaults() *InlineHookResponseCommands {
	this := InlineHookResponseCommands{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineHookResponseCommands) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookResponseCommands) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookResponseCommands) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookResponseCommands) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineHookResponseCommands) GetValue() []InlineHookResponseCommandValue {
	if o == nil || o.Value == nil {
		var ret []InlineHookResponseCommandValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookResponseCommands) GetValueOk() ([]InlineHookResponseCommandValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineHookResponseCommands) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []InlineHookResponseCommandValue and assigns it to the Value field.
func (o *InlineHookResponseCommands) SetValue(v []InlineHookResponseCommandValue) {
	o.Value = v
}

func (o InlineHookResponseCommands) MarshalJSON() ([]byte, error) {
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

func (o *InlineHookResponseCommands) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookResponseCommands := _InlineHookResponseCommands{}

	err = json.Unmarshal(bytes, &varInlineHookResponseCommands)
	if err == nil {
		*o = InlineHookResponseCommands(varInlineHookResponseCommands)
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

type NullableInlineHookResponseCommands struct {
	value *InlineHookResponseCommands
	isSet bool
}

func (v NullableInlineHookResponseCommands) Get() *InlineHookResponseCommands {
	return v.value
}

func (v *NullableInlineHookResponseCommands) Set(val *InlineHookResponseCommands) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookResponseCommands) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookResponseCommands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookResponseCommands(val *InlineHookResponseCommands) *NullableInlineHookResponseCommands {
	return &NullableInlineHookResponseCommands{value: val, isSet: true}
}

func (v NullableInlineHookResponseCommands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookResponseCommands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

