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

// checks if the UserSchemaAttributeEnum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSchemaAttributeEnum{}

// UserSchemaAttributeEnum struct for UserSchemaAttributeEnum
type UserSchemaAttributeEnum struct {
	// The enumerated value
	Const *string `json:"const,omitempty"`
	// The display label for the enumerated value
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaAttributeEnum UserSchemaAttributeEnum

// NewUserSchemaAttributeEnum instantiates a new UserSchemaAttributeEnum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaAttributeEnum() *UserSchemaAttributeEnum {
	this := UserSchemaAttributeEnum{}
	return &this
}

// NewUserSchemaAttributeEnumWithDefaults instantiates a new UserSchemaAttributeEnum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaAttributeEnumWithDefaults() *UserSchemaAttributeEnum {
	this := UserSchemaAttributeEnum{}
	return &this
}

// GetConst returns the Const field value if set, zero value otherwise.
func (o *UserSchemaAttributeEnum) GetConst() string {
	if o == nil || IsNil(o.Const) {
		var ret string
		return ret
	}
	return *o.Const
}

// GetConstOk returns a tuple with the Const field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributeEnum) GetConstOk() (*string, bool) {
	if o == nil || IsNil(o.Const) {
		return nil, false
	}
	return o.Const, true
}

// HasConst returns a boolean if a field has been set.
func (o *UserSchemaAttributeEnum) HasConst() bool {
	if o != nil && !IsNil(o.Const) {
		return true
	}

	return false
}

// SetConst gets a reference to the given string and assigns it to the Const field.
func (o *UserSchemaAttributeEnum) SetConst(v string) {
	o.Const = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserSchemaAttributeEnum) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributeEnum) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserSchemaAttributeEnum) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserSchemaAttributeEnum) SetTitle(v string) {
	o.Title = &v
}

func (o UserSchemaAttributeEnum) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSchemaAttributeEnum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Const) {
		toSerialize["const"] = o.Const
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSchemaAttributeEnum) UnmarshalJSON(data []byte) (err error) {
	varUserSchemaAttributeEnum := _UserSchemaAttributeEnum{}

	err = json.Unmarshal(data, &varUserSchemaAttributeEnum)

	if err != nil {
		return err
	}

	*o = UserSchemaAttributeEnum(varUserSchemaAttributeEnum)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "const")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSchemaAttributeEnum struct {
	value *UserSchemaAttributeEnum
	isSet bool
}

func (v NullableUserSchemaAttributeEnum) Get() *UserSchemaAttributeEnum {
	return v.value
}

func (v *NullableUserSchemaAttributeEnum) Set(val *UserSchemaAttributeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaAttributeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaAttributeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaAttributeEnum(val *UserSchemaAttributeEnum) *NullableUserSchemaAttributeEnum {
	return &NullableUserSchemaAttributeEnum{value: val, isSet: true}
}

func (v NullableUserSchemaAttributeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaAttributeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
