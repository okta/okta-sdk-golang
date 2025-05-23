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

// CreateUISchema The request body properties for the new UI Schema
type CreateUISchema struct {
	UiSchema *UISchemaObject `json:"uiSchema,omitempty"`
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
	if o == nil || o.UiSchema == nil {
		var ret UISchemaObject
		return ret
	}
	return *o.UiSchema
}

// GetUiSchemaOk returns a tuple with the UiSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUISchema) GetUiSchemaOk() (*UISchemaObject, bool) {
	if o == nil || o.UiSchema == nil {
		return nil, false
	}
	return o.UiSchema, true
}

// HasUiSchema returns a boolean if a field has been set.
func (o *CreateUISchema) HasUiSchema() bool {
	if o != nil && o.UiSchema != nil {
		return true
	}

	return false
}

// SetUiSchema gets a reference to the given UISchemaObject and assigns it to the UiSchema field.
func (o *CreateUISchema) SetUiSchema(v UISchemaObject) {
	o.UiSchema = &v
}

func (o CreateUISchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UiSchema != nil {
		toSerialize["uiSchema"] = o.UiSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateUISchema) UnmarshalJSON(bytes []byte) (err error) {
	varCreateUISchema := _CreateUISchema{}

	err = json.Unmarshal(bytes, &varCreateUISchema)
	if err == nil {
		*o = CreateUISchema(varCreateUISchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "uiSchema")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

