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

// checks if the UserSchemaPropertiesProfileItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSchemaPropertiesProfileItem{}

// UserSchemaPropertiesProfileItem struct for UserSchemaPropertiesProfileItem
type UserSchemaPropertiesProfileItem struct {
	Ref                  *string `json:"$ref,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaPropertiesProfileItem UserSchemaPropertiesProfileItem

// NewUserSchemaPropertiesProfileItem instantiates a new UserSchemaPropertiesProfileItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaPropertiesProfileItem() *UserSchemaPropertiesProfileItem {
	this := UserSchemaPropertiesProfileItem{}
	return &this
}

// NewUserSchemaPropertiesProfileItemWithDefaults instantiates a new UserSchemaPropertiesProfileItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaPropertiesProfileItemWithDefaults() *UserSchemaPropertiesProfileItem {
	this := UserSchemaPropertiesProfileItem{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *UserSchemaPropertiesProfileItem) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaPropertiesProfileItem) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *UserSchemaPropertiesProfileItem) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *UserSchemaPropertiesProfileItem) SetRef(v string) {
	o.Ref = &v
}

func (o UserSchemaPropertiesProfileItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSchemaPropertiesProfileItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["$ref"] = o.Ref
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSchemaPropertiesProfileItem) UnmarshalJSON(data []byte) (err error) {
	varUserSchemaPropertiesProfileItem := _UserSchemaPropertiesProfileItem{}

	err = json.Unmarshal(data, &varUserSchemaPropertiesProfileItem)

	if err != nil {
		return err
	}

	*o = UserSchemaPropertiesProfileItem(varUserSchemaPropertiesProfileItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "$ref")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSchemaPropertiesProfileItem struct {
	value *UserSchemaPropertiesProfileItem
	isSet bool
}

func (v NullableUserSchemaPropertiesProfileItem) Get() *UserSchemaPropertiesProfileItem {
	return v.value
}

func (v *NullableUserSchemaPropertiesProfileItem) Set(val *UserSchemaPropertiesProfileItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaPropertiesProfileItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaPropertiesProfileItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaPropertiesProfileItem(val *UserSchemaPropertiesProfileItem) *NullableUserSchemaPropertiesProfileItem {
	return &NullableUserSchemaPropertiesProfileItem{value: val, isSet: true}
}

func (v NullableUserSchemaPropertiesProfileItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaPropertiesProfileItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
