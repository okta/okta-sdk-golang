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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AppPropertiesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPropertiesValue{}

// AppPropertiesValue struct for AppPropertiesValue
type AppPropertiesValue struct {
	// Name of the property
	Name *string `json:"name,omitempty"`
	// Value of the property
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPropertiesValue AppPropertiesValue

// NewAppPropertiesValue instantiates a new AppPropertiesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPropertiesValue() *AppPropertiesValue {
	this := AppPropertiesValue{}
	return &this
}

// NewAppPropertiesValueWithDefaults instantiates a new AppPropertiesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPropertiesValueWithDefaults() *AppPropertiesValue {
	this := AppPropertiesValue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppPropertiesValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPropertiesValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppPropertiesValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppPropertiesValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AppPropertiesValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPropertiesValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AppPropertiesValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AppPropertiesValue) SetValue(v string) {
	o.Value = &v
}

func (o AppPropertiesValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPropertiesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppPropertiesValue) UnmarshalJSON(data []byte) (err error) {
	varAppPropertiesValue := _AppPropertiesValue{}

	err = json.Unmarshal(data, &varAppPropertiesValue)

	if err != nil {
		return err
	}

	*o = AppPropertiesValue(varAppPropertiesValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPropertiesValue struct {
	value *AppPropertiesValue
	isSet bool
}

func (v NullableAppPropertiesValue) Get() *AppPropertiesValue {
	return v.value
}

func (v *NullableAppPropertiesValue) Set(val *AppPropertiesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPropertiesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPropertiesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPropertiesValue(val *AppPropertiesValue) *NullableAppPropertiesValue {
	return &NullableAppPropertiesValue{value: val, isSet: true}
}

func (v NullableAppPropertiesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPropertiesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
