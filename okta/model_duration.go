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

// checks if the Duration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Duration{}

// Duration struct for Duration
type Duration struct {
	Number               *int32  `json:"number,omitempty"`
	Unit                 *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Duration Duration

// NewDuration instantiates a new Duration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuration() *Duration {
	this := Duration{}
	return &this
}

// NewDurationWithDefaults instantiates a new Duration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDurationWithDefaults() *Duration {
	this := Duration{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Duration) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Duration) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Duration) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *Duration) SetNumber(v int32) {
	o.Number = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Duration) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Duration) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Duration) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Duration) SetUnit(v string) {
	o.Unit = &v
}

func (o Duration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Duration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Duration) UnmarshalJSON(data []byte) (err error) {
	varDuration := _Duration{}

	err = json.Unmarshal(data, &varDuration)

	if err != nil {
		return err
	}

	*o = Duration(varDuration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "number")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDuration struct {
	value *Duration
	isSet bool
}

func (v NullableDuration) Get() *Duration {
	return v.value
}

func (v *NullableDuration) Set(val *Duration) {
	v.value = val
	v.isSet = true
}

func (v NullableDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuration(val *Duration) *NullableDuration {
	return &NullableDuration{value: val, isSet: true}
}

func (v NullableDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
