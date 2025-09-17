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

// checks if the StandardRoleAssignmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardRoleAssignmentSchema{}

// StandardRoleAssignmentSchema struct for StandardRoleAssignmentSchema
type StandardRoleAssignmentSchema struct {
	// Specify the standard or IAM-based role type. See [standard roles](/openapi/okta-management/guides/roles/#standard-roles).
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StandardRoleAssignmentSchema StandardRoleAssignmentSchema

// NewStandardRoleAssignmentSchema instantiates a new StandardRoleAssignmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardRoleAssignmentSchema() *StandardRoleAssignmentSchema {
	this := StandardRoleAssignmentSchema{}
	return &this
}

// NewStandardRoleAssignmentSchemaWithDefaults instantiates a new StandardRoleAssignmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardRoleAssignmentSchemaWithDefaults() *StandardRoleAssignmentSchema {
	this := StandardRoleAssignmentSchema{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StandardRoleAssignmentSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRoleAssignmentSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StandardRoleAssignmentSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StandardRoleAssignmentSchema) SetType(v string) {
	o.Type = &v
}

func (o StandardRoleAssignmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardRoleAssignmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StandardRoleAssignmentSchema) UnmarshalJSON(data []byte) (err error) {
	varStandardRoleAssignmentSchema := _StandardRoleAssignmentSchema{}

	err = json.Unmarshal(data, &varStandardRoleAssignmentSchema)

	if err != nil {
		return err
	}

	*o = StandardRoleAssignmentSchema(varStandardRoleAssignmentSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStandardRoleAssignmentSchema struct {
	value *StandardRoleAssignmentSchema
	isSet bool
}

func (v NullableStandardRoleAssignmentSchema) Get() *StandardRoleAssignmentSchema {
	return v.value
}

func (v *NullableStandardRoleAssignmentSchema) Set(val *StandardRoleAssignmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardRoleAssignmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardRoleAssignmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardRoleAssignmentSchema(val *StandardRoleAssignmentSchema) *NullableStandardRoleAssignmentSchema {
	return &NullableStandardRoleAssignmentSchema{value: val, isSet: true}
}

func (v NullableStandardRoleAssignmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardRoleAssignmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
