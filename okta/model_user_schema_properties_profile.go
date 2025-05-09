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

// UserSchemaPropertiesProfile struct for UserSchemaPropertiesProfile
type UserSchemaPropertiesProfile struct {
	AllOf []UserSchemaPropertiesProfileItem `json:"allOf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaPropertiesProfile UserSchemaPropertiesProfile

// NewUserSchemaPropertiesProfile instantiates a new UserSchemaPropertiesProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaPropertiesProfile() *UserSchemaPropertiesProfile {
	this := UserSchemaPropertiesProfile{}
	return &this
}

// NewUserSchemaPropertiesProfileWithDefaults instantiates a new UserSchemaPropertiesProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaPropertiesProfileWithDefaults() *UserSchemaPropertiesProfile {
	this := UserSchemaPropertiesProfile{}
	return &this
}

// GetAllOf returns the AllOf field value if set, zero value otherwise.
func (o *UserSchemaPropertiesProfile) GetAllOf() []UserSchemaPropertiesProfileItem {
	if o == nil || o.AllOf == nil {
		var ret []UserSchemaPropertiesProfileItem
		return ret
	}
	return o.AllOf
}

// GetAllOfOk returns a tuple with the AllOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaPropertiesProfile) GetAllOfOk() ([]UserSchemaPropertiesProfileItem, bool) {
	if o == nil || o.AllOf == nil {
		return nil, false
	}
	return o.AllOf, true
}

// HasAllOf returns a boolean if a field has been set.
func (o *UserSchemaPropertiesProfile) HasAllOf() bool {
	if o != nil && o.AllOf != nil {
		return true
	}

	return false
}

// SetAllOf gets a reference to the given []UserSchemaPropertiesProfileItem and assigns it to the AllOf field.
func (o *UserSchemaPropertiesProfile) SetAllOf(v []UserSchemaPropertiesProfileItem) {
	o.AllOf = v
}

func (o UserSchemaPropertiesProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllOf != nil {
		toSerialize["allOf"] = o.AllOf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaPropertiesProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaPropertiesProfile := _UserSchemaPropertiesProfile{}

	err = json.Unmarshal(bytes, &varUserSchemaPropertiesProfile)
	if err == nil {
		*o = UserSchemaPropertiesProfile(varUserSchemaPropertiesProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allOf")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaPropertiesProfile struct {
	value *UserSchemaPropertiesProfile
	isSet bool
}

func (v NullableUserSchemaPropertiesProfile) Get() *UserSchemaPropertiesProfile {
	return v.value
}

func (v *NullableUserSchemaPropertiesProfile) Set(val *UserSchemaPropertiesProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaPropertiesProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaPropertiesProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaPropertiesProfile(val *UserSchemaPropertiesProfile) *NullableUserSchemaPropertiesProfile {
	return &NullableUserSchemaPropertiesProfile{value: val, isSet: true}
}

func (v NullableUserSchemaPropertiesProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaPropertiesProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

