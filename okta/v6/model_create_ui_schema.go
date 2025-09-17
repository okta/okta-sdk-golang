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

// checks if the CreateUISchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUISchema{}

// CreateUISchema The request body properties for the new UI Schema
type CreateUISchema struct {
	UiSchema             *UISchemaObject `json:"uiSchema,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUISchema CreateUISchema

// NewCreateUISchema instantiates a new CreateUISchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUISchema() *CreateUISchema {
	this := CreateUISchema{}
	return &this
}

// NewCreateUISchemaWithDefaults instantiates a new CreateUISchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUISchemaWithDefaults() *CreateUISchema {
	this := CreateUISchema{}
	return &this
}

// GetUiSchema returns the UiSchema field value if set, zero value otherwise.
func (o *CreateUISchema) GetUiSchema() UISchemaObject {
	if o == nil || IsNil(o.UiSchema) {
		var ret UISchemaObject
		return ret
	}
	return *o.UiSchema
}

// GetUiSchemaOk returns a tuple with the UiSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUISchema) GetUiSchemaOk() (*UISchemaObject, bool) {
	if o == nil || IsNil(o.UiSchema) {
		return nil, false
	}
	return o.UiSchema, true
}

// HasUiSchema returns a boolean if a field has been set.
func (o *CreateUISchema) HasUiSchema() bool {
	if o != nil && !IsNil(o.UiSchema) {
		return true
	}

	return false
}

// SetUiSchema gets a reference to the given UISchemaObject and assigns it to the UiSchema field.
func (o *CreateUISchema) SetUiSchema(v UISchemaObject) {
	o.UiSchema = &v
}

func (o CreateUISchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUISchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UiSchema) {
		toSerialize["uiSchema"] = o.UiSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateUISchema) UnmarshalJSON(data []byte) (err error) {
	varCreateUISchema := _CreateUISchema{}

	err = json.Unmarshal(data, &varCreateUISchema)

	if err != nil {
		return err
	}

	*o = CreateUISchema(varCreateUISchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uiSchema")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUISchema struct {
	value *CreateUISchema
	isSet bool
}

func (v NullableCreateUISchema) Get() *CreateUISchema {
	return v.value
}

func (v *NullableCreateUISchema) Set(val *CreateUISchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUISchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUISchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUISchema(val *CreateUISchema) *NullableCreateUISchema {
	return &NullableCreateUISchema{value: val, isSet: true}
}

func (v NullableCreateUISchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUISchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
