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
	"reflect"
	"strings"
)

// checks if the InlineHookChannelHttpCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookChannelHttpCreate{}

// InlineHookChannelHttpCreate struct for InlineHookChannelHttpCreate
type InlineHookChannelHttpCreate struct {
	InlineHookChannelCreate
	Config               *InlineHookHttpConfigCreate `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelHttpCreate InlineHookChannelHttpCreate

// NewInlineHookChannelHttpCreate instantiates a new InlineHookChannelHttpCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelHttpCreate() *InlineHookChannelHttpCreate {
	this := InlineHookChannelHttpCreate{}
	return &this
}

// NewInlineHookChannelHttpCreateWithDefaults instantiates a new InlineHookChannelHttpCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelHttpCreateWithDefaults() *InlineHookChannelHttpCreate {
	this := InlineHookChannelHttpCreate{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InlineHookChannelHttpCreate) GetConfig() InlineHookHttpConfigCreate {
	if o == nil || IsNil(o.Config) {
		var ret InlineHookHttpConfigCreate
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelHttpCreate) GetConfigOk() (*InlineHookHttpConfigCreate, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *InlineHookChannelHttpCreate) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given InlineHookHttpConfigCreate and assigns it to the Config field.
func (o *InlineHookChannelHttpCreate) SetConfig(v InlineHookHttpConfigCreate) {
	o.Config = &v
}

func (o InlineHookChannelHttpCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookChannelHttpCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInlineHookChannelCreate, errInlineHookChannelCreate := json.Marshal(o.InlineHookChannelCreate)
	if errInlineHookChannelCreate != nil {
		return map[string]interface{}{}, errInlineHookChannelCreate
	}
	errInlineHookChannelCreate = json.Unmarshal([]byte(serializedInlineHookChannelCreate), &toSerialize)
	if errInlineHookChannelCreate != nil {
		return map[string]interface{}{}, errInlineHookChannelCreate
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookChannelHttpCreate) UnmarshalJSON(data []byte) (err error) {
	type InlineHookChannelHttpCreateWithoutEmbeddedStruct struct {
		Config *InlineHookHttpConfigCreate `json:"config,omitempty"`
	}

	varInlineHookChannelHttpCreateWithoutEmbeddedStruct := InlineHookChannelHttpCreateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varInlineHookChannelHttpCreateWithoutEmbeddedStruct)
	if err == nil {
		varInlineHookChannelHttpCreate := _InlineHookChannelHttpCreate{}
		varInlineHookChannelHttpCreate.Config = varInlineHookChannelHttpCreateWithoutEmbeddedStruct.Config
		*o = InlineHookChannelHttpCreate(varInlineHookChannelHttpCreate)
	} else {
		return err
	}

	varInlineHookChannelHttpCreate := _InlineHookChannelHttpCreate{}

	err = json.Unmarshal(data, &varInlineHookChannelHttpCreate)
	if err == nil {
		o.InlineHookChannelCreate = varInlineHookChannelHttpCreate.InlineHookChannelCreate
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")

		// remove fields from embedded structs
		reflectInlineHookChannelCreate := reflect.ValueOf(o.InlineHookChannelCreate)
		for i := 0; i < reflectInlineHookChannelCreate.Type().NumField(); i++ {
			t := reflectInlineHookChannelCreate.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookChannelHttpCreate struct {
	value *InlineHookChannelHttpCreate
	isSet bool
}

func (v NullableInlineHookChannelHttpCreate) Get() *InlineHookChannelHttpCreate {
	return v.value
}

func (v *NullableInlineHookChannelHttpCreate) Set(val *InlineHookChannelHttpCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelHttpCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelHttpCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelHttpCreate(val *InlineHookChannelHttpCreate) *NullableInlineHookChannelHttpCreate {
	return &NullableInlineHookChannelHttpCreate{value: val, isSet: true}
}

func (v NullableInlineHookChannelHttpCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelHttpCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
