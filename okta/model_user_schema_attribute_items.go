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

// UserSchemaAttributeItems struct for UserSchemaAttributeItems
type UserSchemaAttributeItems struct {
	Enum []string `json:"enum,omitempty"`
	OneOf []UserSchemaAttributeEnum `json:"oneOf,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaAttributeItems UserSchemaAttributeItems

// NewUserSchemaAttributeItems instantiates a new UserSchemaAttributeItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaAttributeItems() *UserSchemaAttributeItems {
	this := UserSchemaAttributeItems{}
	return &this
}

// NewUserSchemaAttributeItemsWithDefaults instantiates a new UserSchemaAttributeItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaAttributeItemsWithDefaults() *UserSchemaAttributeItems {
	this := UserSchemaAttributeItems{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *UserSchemaAttributeItems) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributeItems) GetEnumOk() ([]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *UserSchemaAttributeItems) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *UserSchemaAttributeItems) SetEnum(v []string) {
	o.Enum = v
}

// GetOneOf returns the OneOf field value if set, zero value otherwise.
func (o *UserSchemaAttributeItems) GetOneOf() []UserSchemaAttributeEnum {
	if o == nil || o.OneOf == nil {
		var ret []UserSchemaAttributeEnum
		return ret
	}
	return o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributeItems) GetOneOfOk() ([]UserSchemaAttributeEnum, bool) {
	if o == nil || o.OneOf == nil {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *UserSchemaAttributeItems) HasOneOf() bool {
	if o != nil && o.OneOf != nil {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given []UserSchemaAttributeEnum and assigns it to the OneOf field.
func (o *UserSchemaAttributeItems) SetOneOf(v []UserSchemaAttributeEnum) {
	o.OneOf = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserSchemaAttributeItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributeItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSchemaAttributeItems) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSchemaAttributeItems) SetType(v string) {
	o.Type = &v
}

func (o UserSchemaAttributeItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.OneOf != nil {
		toSerialize["oneOf"] = o.OneOf
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaAttributeItems) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaAttributeItems := _UserSchemaAttributeItems{}

	err = json.Unmarshal(bytes, &varUserSchemaAttributeItems)
	if err == nil {
		*o = UserSchemaAttributeItems(varUserSchemaAttributeItems)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enum")
		delete(additionalProperties, "oneOf")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaAttributeItems struct {
	value *UserSchemaAttributeItems
	isSet bool
}

func (v NullableUserSchemaAttributeItems) Get() *UserSchemaAttributeItems {
	return v.value
}

func (v *NullableUserSchemaAttributeItems) Set(val *UserSchemaAttributeItems) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaAttributeItems) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaAttributeItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaAttributeItems(val *UserSchemaAttributeItems) *NullableUserSchemaAttributeItems {
	return &NullableUserSchemaAttributeItems{value: val, isSet: true}
}

func (v NullableUserSchemaAttributeItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaAttributeItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

