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
	"fmt"
)

// checks if the StandardRoleAssignmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardRoleAssignmentSchema{}

// StandardRoleAssignmentSchema struct for StandardRoleAssignmentSchema
type StandardRoleAssignmentSchema struct {
	// Specify a [standard admin role](/openapi/okta-management/guides/roles/#standard-roles), an [IAM-based standard role](/openapi/okta-management/guides/roles/#iam-based-standard-roles), or `CUSTOM` for a custom role type:
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _StandardRoleAssignmentSchema StandardRoleAssignmentSchema

// NewStandardRoleAssignmentSchema instantiates a new StandardRoleAssignmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardRoleAssignmentSchema(type_ string) *StandardRoleAssignmentSchema {
	this := StandardRoleAssignmentSchema{}
	this.Type = type_
	return &this
}

// NewStandardRoleAssignmentSchemaWithDefaults instantiates a new StandardRoleAssignmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardRoleAssignmentSchemaWithDefaults() *StandardRoleAssignmentSchema {
	this := StandardRoleAssignmentSchema{}
	return &this
}

// GetType returns the Type field value
func (o *StandardRoleAssignmentSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StandardRoleAssignmentSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StandardRoleAssignmentSchema) SetType(v string) {
	o.Type = v
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
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StandardRoleAssignmentSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

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
