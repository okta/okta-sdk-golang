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

// InlineHookResponseCommandValue struct for InlineHookResponseCommandValue
type InlineHookResponseCommandValue struct {
	Op *string `json:"op,omitempty"`
	Path *string `json:"path,omitempty"`
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookResponseCommandValue InlineHookResponseCommandValue

// NewInlineHookResponseCommandValue instantiates a new InlineHookResponseCommandValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookResponseCommandValue() *InlineHookResponseCommandValue {
	this := InlineHookResponseCommandValue{}
	return &this
}

// NewInlineHookResponseCommandValueWithDefaults instantiates a new InlineHookResponseCommandValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookResponseCommandValueWithDefaults() *InlineHookResponseCommandValue {
	this := InlineHookResponseCommandValue{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *InlineHookResponseCommandValue) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookResponseCommandValue) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *InlineHookResponseCommandValue) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *InlineHookResponseCommandValue) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *InlineHookResponseCommandValue) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookResponseCommandValue) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *InlineHookResponseCommandValue) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *InlineHookResponseCommandValue) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineHookResponseCommandValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookResponseCommandValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineHookResponseCommandValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineHookResponseCommandValue) SetValue(v string) {
	o.Value = &v
}

func (o InlineHookResponseCommandValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookResponseCommandValue) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookResponseCommandValue := _InlineHookResponseCommandValue{}

	err = json.Unmarshal(bytes, &varInlineHookResponseCommandValue)
	if err == nil {
		*o = InlineHookResponseCommandValue(varInlineHookResponseCommandValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookResponseCommandValue struct {
	value *InlineHookResponseCommandValue
	isSet bool
}

func (v NullableInlineHookResponseCommandValue) Get() *InlineHookResponseCommandValue {
	return v.value
}

func (v *NullableInlineHookResponseCommandValue) Set(val *InlineHookResponseCommandValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookResponseCommandValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookResponseCommandValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookResponseCommandValue(val *InlineHookResponseCommandValue) *NullableInlineHookResponseCommandValue {
	return &NullableInlineHookResponseCommandValue{value: val, isSet: true}
}

func (v NullableInlineHookResponseCommandValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookResponseCommandValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

