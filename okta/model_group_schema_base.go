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

// GroupSchemaBase struct for GroupSchemaBase
type GroupSchemaBase struct {
	Id *string `json:"id,omitempty"`
	Properties *GroupSchemaBaseProperties `json:"properties,omitempty"`
	Required []string `json:"required,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupSchemaBase GroupSchemaBase

// NewGroupSchemaBase instantiates a new GroupSchemaBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSchemaBase() *GroupSchemaBase {
	this := GroupSchemaBase{}
	return &this
}

// NewGroupSchemaBaseWithDefaults instantiates a new GroupSchemaBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSchemaBaseWithDefaults() *GroupSchemaBase {
	this := GroupSchemaBase{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupSchemaBase) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaBase) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupSchemaBase) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupSchemaBase) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GroupSchemaBase) GetProperties() GroupSchemaBaseProperties {
	if o == nil || o.Properties == nil {
		var ret GroupSchemaBaseProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaBase) GetPropertiesOk() (*GroupSchemaBaseProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GroupSchemaBase) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given GroupSchemaBaseProperties and assigns it to the Properties field.
func (o *GroupSchemaBase) SetProperties(v GroupSchemaBaseProperties) {
	o.Properties = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *GroupSchemaBase) GetRequired() []string {
	if o == nil || o.Required == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaBase) GetRequiredOk() ([]string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *GroupSchemaBase) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *GroupSchemaBase) SetRequired(v []string) {
	o.Required = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupSchemaBase) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaBase) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupSchemaBase) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupSchemaBase) SetType(v string) {
	o.Type = &v
}

func (o GroupSchemaBase) MarshalJSON() ([]byte, error) {
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

func (o *GroupSchemaBase) UnmarshalJSON(bytes []byte) (err error) {
	varGroupSchemaBase := _GroupSchemaBase{}

	err = json.Unmarshal(bytes, &varGroupSchemaBase)
	if err == nil {
		*o = GroupSchemaBase(varGroupSchemaBase)
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

type NullableGroupSchemaBase struct {
	value *GroupSchemaBase
	isSet bool
}

func (v NullableGroupSchemaBase) Get() *GroupSchemaBase {
	return v.value
}

func (v *NullableGroupSchemaBase) Set(val *GroupSchemaBase) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSchemaBase) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSchemaBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSchemaBase(val *GroupSchemaBase) *NullableGroupSchemaBase {
	return &NullableGroupSchemaBase{value: val, isSet: true}
}

func (v NullableGroupSchemaBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSchemaBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

