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

// checks if the CimdClientEntityEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdClientEntityEmbedded{}

// CimdClientEntityEmbedded Embedded resources that are included when the `expand` query parameter is specified
type CimdClientEntityEmbedded struct {
	// The binding objects attached to the CIMD client entity
	Bindings             []CimdClientEntityBinding `json:"bindings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CimdClientEntityEmbedded CimdClientEntityEmbedded

// NewCimdClientEntityEmbedded instantiates a new CimdClientEntityEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdClientEntityEmbedded() *CimdClientEntityEmbedded {
	this := CimdClientEntityEmbedded{}
	return &this
}

// NewCimdClientEntityEmbeddedWithDefaults instantiates a new CimdClientEntityEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdClientEntityEmbeddedWithDefaults() *CimdClientEntityEmbedded {
	this := CimdClientEntityEmbedded{}
	return &this
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *CimdClientEntityEmbedded) GetBindings() []CimdClientEntityBinding {
	if o == nil || IsNil(o.Bindings) {
		var ret []CimdClientEntityBinding
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityEmbedded) GetBindingsOk() ([]CimdClientEntityBinding, bool) {
	if o == nil || IsNil(o.Bindings) {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *CimdClientEntityEmbedded) HasBindings() bool {
	if o != nil && !IsNil(o.Bindings) {
		return true
	}

	return false
}

// SetBindings gets a reference to the given []CimdClientEntityBinding and assigns it to the Bindings field.
func (o *CimdClientEntityEmbedded) SetBindings(v []CimdClientEntityBinding) {
	o.Bindings = v
}

func (o CimdClientEntityEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdClientEntityEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bindings) {
		toSerialize["bindings"] = o.Bindings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CimdClientEntityEmbedded) UnmarshalJSON(data []byte) (err error) {
	varCimdClientEntityEmbedded := _CimdClientEntityEmbedded{}

	err = json.Unmarshal(data, &varCimdClientEntityEmbedded)

	if err != nil {
		return err
	}

	*o = CimdClientEntityEmbedded(varCimdClientEntityEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bindings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdClientEntityEmbedded struct {
	value *CimdClientEntityEmbedded
	isSet bool
}

func (v NullableCimdClientEntityEmbedded) Get() *CimdClientEntityEmbedded {
	return v.value
}

func (v *NullableCimdClientEntityEmbedded) Set(val *CimdClientEntityEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdClientEntityEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdClientEntityEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdClientEntityEmbedded(val *CimdClientEntityEmbedded) *NullableCimdClientEntityEmbedded {
	return &NullableCimdClientEntityEmbedded{value: val, isSet: true}
}

func (v NullableCimdClientEntityEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdClientEntityEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
