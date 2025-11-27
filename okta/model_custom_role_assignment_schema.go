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

// checks if the CustomRoleAssignmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRoleAssignmentSchema{}

// CustomRoleAssignmentSchema struct for CustomRoleAssignmentSchema
type CustomRoleAssignmentSchema struct {
	// Resource set ID
	ResourceSet string `json:"resource-set"`
	// Custom role ID
	Role string `json:"role"`
	// Specify a [standard admin role](/openapi/okta-management/guides/roles/#standard-roles), an [IAM-based standard role](/openapi/okta-management/guides/roles/#iam-based-standard-roles), or `CUSTOM` for a custom role type:
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _CustomRoleAssignmentSchema CustomRoleAssignmentSchema

// NewCustomRoleAssignmentSchema instantiates a new CustomRoleAssignmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRoleAssignmentSchema(resourceSet string, role string, type_ string) *CustomRoleAssignmentSchema {
	this := CustomRoleAssignmentSchema{}
	this.ResourceSet = resourceSet
	this.Role = role
	this.Type = type_
	return &this
}

// NewCustomRoleAssignmentSchemaWithDefaults instantiates a new CustomRoleAssignmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRoleAssignmentSchemaWithDefaults() *CustomRoleAssignmentSchema {
	this := CustomRoleAssignmentSchema{}
	return &this
}

// GetResourceSet returns the ResourceSet field value
func (o *CustomRoleAssignmentSchema) GetResourceSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value
// and a boolean to check if the value has been set.
func (o *CustomRoleAssignmentSchema) GetResourceSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSet, true
}

// SetResourceSet sets field value
func (o *CustomRoleAssignmentSchema) SetResourceSet(v string) {
	o.ResourceSet = v
}

// GetRole returns the Role field value
func (o *CustomRoleAssignmentSchema) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CustomRoleAssignmentSchema) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CustomRoleAssignmentSchema) SetRole(v string) {
	o.Role = v
}

// GetType returns the Type field value
func (o *CustomRoleAssignmentSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomRoleAssignmentSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomRoleAssignmentSchema) SetType(v string) {
	o.Type = v
}

func (o CustomRoleAssignmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomRoleAssignmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource-set"] = o.ResourceSet
	toSerialize["role"] = o.Role
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomRoleAssignmentSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource-set",
		"role",
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

	varCustomRoleAssignmentSchema := _CustomRoleAssignmentSchema{}

	err = json.Unmarshal(data, &varCustomRoleAssignmentSchema)

	if err != nil {
		return err
	}

	*o = CustomRoleAssignmentSchema(varCustomRoleAssignmentSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resource-set")
		delete(additionalProperties, "role")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomRoleAssignmentSchema struct {
	value *CustomRoleAssignmentSchema
	isSet bool
}

func (v NullableCustomRoleAssignmentSchema) Get() *CustomRoleAssignmentSchema {
	return v.value
}

func (v *NullableCustomRoleAssignmentSchema) Set(val *CustomRoleAssignmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomRoleAssignmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomRoleAssignmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomRoleAssignmentSchema(val *CustomRoleAssignmentSchema) *NullableCustomRoleAssignmentSchema {
	return &NullableCustomRoleAssignmentSchema{value: val, isSet: true}
}

func (v NullableCustomRoleAssignmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomRoleAssignmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
