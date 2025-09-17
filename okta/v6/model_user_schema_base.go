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

// checks if the UserSchemaBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSchemaBase{}

// UserSchemaBase All Okta-defined profile properties are defined in a profile subschema with the resolution scope `#base`. You can't modify these properties, except to update permissions, to change the nullability of `firstName` and `lastName`, or to specify a pattern for `login`. They can't be removed.  The base user profile is based on the [System for Cross-domain Identity Management: Core Schema](https://tools.ietf.org/html/draft-ietf-scim-core-schema-22#section-4.1.1) and has the standard properties detailed below.
type UserSchemaBase struct {
	// The subschema name
	Id *string `json:"id,omitempty"`
	// The `#base` object properties
	Properties *UserSchemaBaseProperties `json:"properties,omitempty"`
	// A collection indicating required property names
	Required []string `json:"required,omitempty"`
	// The object type
	Type                 *string `json:"type,omitempty"`
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSchemaBase) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Properties) {
		var ret UserSchemaBaseProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetPropertiesOk() (*UserSchemaBaseProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserSchemaBase) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
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
	if o == nil || IsNil(o.Required) {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UserSchemaBase) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBase) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSchemaBase) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSchemaBase) SetType(v string) {
	o.Type = &v
}

func (o UserSchemaBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSchemaBase) ToMap() (map[string]interface{}, error) {
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

func (o *UserSchemaBase) UnmarshalJSON(data []byte) (err error) {
	varUserSchemaBase := _UserSchemaBase{}

	err = json.Unmarshal(data, &varUserSchemaBase)

	if err != nil {
		return err
	}

	*o = UserSchemaBase(varUserSchemaBase)

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
