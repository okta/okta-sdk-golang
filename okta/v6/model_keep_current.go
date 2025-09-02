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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// KeepCurrent struct for KeepCurrent
type KeepCurrent struct {
	// Skip deleting the user's current session when set to `true`
	KeepCurrent *bool `json:"keepCurrent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeepCurrent KeepCurrent

// NewKeepCurrent instantiates a new KeepCurrent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeepCurrent() *KeepCurrent {
	this := KeepCurrent{}
	var keepCurrent bool = true
	this.KeepCurrent = &keepCurrent
	return &this
}

// NewKeepCurrentWithDefaults instantiates a new KeepCurrent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeepCurrentWithDefaults() *KeepCurrent {
	this := KeepCurrent{}
	var keepCurrent bool = true
	this.KeepCurrent = &keepCurrent
	return &this
}

// GetKeepCurrent returns the KeepCurrent field value if set, zero value otherwise.
func (o *KeepCurrent) GetKeepCurrent() bool {
	if o == nil || o.KeepCurrent == nil {
		var ret bool
		return ret
	}
	return *o.KeepCurrent
}

// GetKeepCurrentOk returns a tuple with the KeepCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepCurrent) GetKeepCurrentOk() (*bool, bool) {
	if o == nil || o.KeepCurrent == nil {
		return nil, false
	}
	return o.KeepCurrent, true
}

// HasKeepCurrent returns a boolean if a field has been set.
func (o *KeepCurrent) HasKeepCurrent() bool {
	if o != nil && o.KeepCurrent != nil {
		return true
	}

	return false
}

// SetKeepCurrent gets a reference to the given bool and assigns it to the KeepCurrent field.
func (o *KeepCurrent) SetKeepCurrent(v bool) {
	o.KeepCurrent = &v
}

func (o KeepCurrent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeepCurrent != nil {
		toSerialize["keepCurrent"] = o.KeepCurrent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KeepCurrent) UnmarshalJSON(bytes []byte) (err error) {
	varKeepCurrent := _KeepCurrent{}

	err = json.Unmarshal(bytes, &varKeepCurrent)
	if err == nil {
		*o = KeepCurrent(varKeepCurrent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "keepCurrent")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableKeepCurrent struct {
	value *KeepCurrent
	isSet bool
}

func (v NullableKeepCurrent) Get() *KeepCurrent {
	return v.value
}

func (v *NullableKeepCurrent) Set(val *KeepCurrent) {
	v.value = val
	v.isSet = true
}

func (v NullableKeepCurrent) IsSet() bool {
	return v.isSet
}

func (v *NullableKeepCurrent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeepCurrent(val *KeepCurrent) *NullableKeepCurrent {
	return &NullableKeepCurrent{value: val, isSet: true}
}

func (v NullableKeepCurrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeepCurrent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

