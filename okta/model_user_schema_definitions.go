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

// UserSchemaDefinitions struct for UserSchemaDefinitions
type UserSchemaDefinitions struct {
	Base *UserSchemaBase `json:"base,omitempty"`
	Custom *UserSchemaPublic `json:"custom,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaDefinitions UserSchemaDefinitions

// NewUserSchemaDefinitions instantiates a new UserSchemaDefinitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaDefinitions() *UserSchemaDefinitions {
	this := UserSchemaDefinitions{}
	return &this
}

// NewUserSchemaDefinitionsWithDefaults instantiates a new UserSchemaDefinitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaDefinitionsWithDefaults() *UserSchemaDefinitions {
	this := UserSchemaDefinitions{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *UserSchemaDefinitions) GetBase() UserSchemaBase {
	if o == nil || o.Base == nil {
		var ret UserSchemaBase
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaDefinitions) GetBaseOk() (*UserSchemaBase, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *UserSchemaDefinitions) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given UserSchemaBase and assigns it to the Base field.
func (o *UserSchemaDefinitions) SetBase(v UserSchemaBase) {
	o.Base = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *UserSchemaDefinitions) GetCustom() UserSchemaPublic {
	if o == nil || o.Custom == nil {
		var ret UserSchemaPublic
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaDefinitions) GetCustomOk() (*UserSchemaPublic, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *UserSchemaDefinitions) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given UserSchemaPublic and assigns it to the Custom field.
func (o *UserSchemaDefinitions) SetCustom(v UserSchemaPublic) {
	o.Custom = &v
}

func (o UserSchemaDefinitions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaDefinitions) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaDefinitions := _UserSchemaDefinitions{}

	err = json.Unmarshal(bytes, &varUserSchemaDefinitions)
	if err == nil {
		*o = UserSchemaDefinitions(varUserSchemaDefinitions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "base")
		delete(additionalProperties, "custom")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaDefinitions struct {
	value *UserSchemaDefinitions
	isSet bool
}

func (v NullableUserSchemaDefinitions) Get() *UserSchemaDefinitions {
	return v.value
}

func (v *NullableUserSchemaDefinitions) Set(val *UserSchemaDefinitions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaDefinitions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaDefinitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaDefinitions(val *UserSchemaDefinitions) *NullableUserSchemaDefinitions {
	return &NullableUserSchemaDefinitions{value: val, isSet: true}
}

func (v NullableUserSchemaDefinitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaDefinitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

