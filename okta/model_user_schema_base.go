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

// UserSchemaBase struct for UserSchemaBase
type UserSchemaBase struct {
	Id *string `json:"id,omitempty"`
	Properties *UserSchemaBaseProperties `json:"properties,omitempty"`
	Required []string `json:"required,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaBase UserSchemaBase

// NewUserSchemaBase instantiates a new UserSchemaBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaBase() *UserSchemaBase {
	this := UserSchemaBase{}
	return &this
}

// NewUserSchemaBaseWithDefaults instantiates a new UserSchemaBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaBaseWithDefaults() *UserSchemaBase {
	this := UserSchemaBase{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSchemaBase) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSchemaBase) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSchemaBase) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserSchemaBase) GetProperties() UserSchemaBaseProperties {
	if o == nil || o.Properties == nil {
		var ret UserSchemaBaseProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetPropertiesOk() (*UserSchemaBaseProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserSchemaBase) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given UserSchemaBaseProperties and assigns it to the Properties field.
func (o *UserSchemaBase) SetProperties(v UserSchemaBaseProperties) {
	o.Properties = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UserSchemaBase) GetRequired() []string {
	if o == nil || o.Required == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetRequiredOk() ([]string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UserSchemaBase) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *UserSchemaBase) SetRequired(v []string) {
	o.Required = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserSchemaBase) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSchemaBase) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSchemaBase) SetType(v string) {
	o.Type = &v
}

func (o UserSchemaBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaBase) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaBase := _UserSchemaBase{}

	err = json.Unmarshal(bytes, &varUserSchemaBase)
	if err == nil {
		*o = UserSchemaBase(varUserSchemaBase)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "required")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaBase struct {
	value *UserSchemaBase
	isSet bool
}

func (v NullableUserSchemaBase) Get() *UserSchemaBase {
	return v.value
}

func (v *NullableUserSchemaBase) Set(val *UserSchemaBase) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaBase) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaBase(val *UserSchemaBase) *NullableUserSchemaBase {
	return &NullableUserSchemaBase{value: val, isSet: true}
}

func (v NullableUserSchemaBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

