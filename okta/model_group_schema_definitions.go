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

// GroupSchemaDefinitions struct for GroupSchemaDefinitions
type GroupSchemaDefinitions struct {
	Base *GroupSchemaBase `json:"base,omitempty"`
	Custom *GroupSchemaCustom `json:"custom,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupSchemaDefinitions GroupSchemaDefinitions

// NewGroupSchemaDefinitions instantiates a new GroupSchemaDefinitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSchemaDefinitions() *GroupSchemaDefinitions {
	this := GroupSchemaDefinitions{}
	return &this
}

// NewGroupSchemaDefinitionsWithDefaults instantiates a new GroupSchemaDefinitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSchemaDefinitionsWithDefaults() *GroupSchemaDefinitions {
	this := GroupSchemaDefinitions{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *GroupSchemaDefinitions) GetBase() GroupSchemaBase {
	if o == nil || o.Base == nil {
		var ret GroupSchemaBase
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaDefinitions) GetBaseOk() (*GroupSchemaBase, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *GroupSchemaDefinitions) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given GroupSchemaBase and assigns it to the Base field.
func (o *GroupSchemaDefinitions) SetBase(v GroupSchemaBase) {
	o.Base = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *GroupSchemaDefinitions) GetCustom() GroupSchemaCustom {
	if o == nil || o.Custom == nil {
		var ret GroupSchemaCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaDefinitions) GetCustomOk() (*GroupSchemaCustom, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *GroupSchemaDefinitions) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given GroupSchemaCustom and assigns it to the Custom field.
func (o *GroupSchemaDefinitions) SetCustom(v GroupSchemaCustom) {
	o.Custom = &v
}

func (o GroupSchemaDefinitions) MarshalJSON() ([]byte, error) {
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

func (o *GroupSchemaDefinitions) UnmarshalJSON(bytes []byte) (err error) {
	varGroupSchemaDefinitions := _GroupSchemaDefinitions{}

	err = json.Unmarshal(bytes, &varGroupSchemaDefinitions)
	if err == nil {
		*o = GroupSchemaDefinitions(varGroupSchemaDefinitions)
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

type NullableGroupSchemaDefinitions struct {
	value *GroupSchemaDefinitions
	isSet bool
}

func (v NullableGroupSchemaDefinitions) Get() *GroupSchemaDefinitions {
	return v.value
}

func (v *NullableGroupSchemaDefinitions) Set(val *GroupSchemaDefinitions) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSchemaDefinitions) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSchemaDefinitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSchemaDefinitions(val *GroupSchemaDefinitions) *NullableGroupSchemaDefinitions {
	return &NullableGroupSchemaDefinitions{value: val, isSet: true}
}

func (v NullableGroupSchemaDefinitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSchemaDefinitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

