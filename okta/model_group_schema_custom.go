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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GroupSchemaCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSchemaCustom{}

// GroupSchemaCustom All custom profile properties are defined in a profile subschema with the resolution scope `#custom`
type GroupSchemaCustom struct {
	// The subschema name
	Id *string `json:"id,omitempty"`
	// The `#custom` object properties
	Properties *map[string]GroupSchemaAttribute `json:"properties,omitempty"`
	// A collection indicating required property names
	Required []string `json:"required,omitempty"`
	// The object type
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupSchemaCustom GroupSchemaCustom

// NewGroupSchemaCustom instantiates a new GroupSchemaCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSchemaCustom() *GroupSchemaCustom {
	this := GroupSchemaCustom{}
	return &this
}

// NewGroupSchemaCustomWithDefaults instantiates a new GroupSchemaCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSchemaCustomWithDefaults() *GroupSchemaCustom {
	this := GroupSchemaCustom{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupSchemaCustom) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaCustom) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupSchemaCustom) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupSchemaCustom) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GroupSchemaCustom) GetProperties() map[string]GroupSchemaAttribute {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]GroupSchemaAttribute
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaCustom) GetPropertiesOk() (*map[string]GroupSchemaAttribute, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GroupSchemaCustom) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]GroupSchemaAttribute and assigns it to the Properties field.
func (o *GroupSchemaCustom) SetProperties(v map[string]GroupSchemaAttribute) {
	o.Properties = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *GroupSchemaCustom) GetRequired() []string {
	if o == nil || IsNil(o.Required) {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaCustom) GetRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *GroupSchemaCustom) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *GroupSchemaCustom) SetRequired(v []string) {
	o.Required = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupSchemaCustom) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaCustom) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupSchemaCustom) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupSchemaCustom) SetType(v string) {
	o.Type = &v
}

func (o GroupSchemaCustom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSchemaCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupSchemaCustom) UnmarshalJSON(data []byte) (err error) {
	varGroupSchemaCustom := _GroupSchemaCustom{}

	err = json.Unmarshal(data, &varGroupSchemaCustom)

	if err != nil {
		return err
	}

	*o = GroupSchemaCustom(varGroupSchemaCustom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "required")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupSchemaCustom struct {
	value *GroupSchemaCustom
	isSet bool
}

func (v NullableGroupSchemaCustom) Get() *GroupSchemaCustom {
	return v.value
}

func (v *NullableGroupSchemaCustom) Set(val *GroupSchemaCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSchemaCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSchemaCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSchemaCustom(val *GroupSchemaCustom) *NullableGroupSchemaCustom {
	return &NullableGroupSchemaCustom{value: val, isSet: true}
}

func (v NullableGroupSchemaCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSchemaCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
