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

// checks if the CustomRoleAssignmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRoleAssignmentSchema{}

// CustomRoleAssignmentSchema struct for CustomRoleAssignmentSchema
type CustomRoleAssignmentSchema struct {
	// Resource set ID
	ResourceSet *string `json:"resource-set,omitempty"`
	// Custom role ID
	Role *string `json:"role,omitempty"`
	// The type of role. Specify `CUSTOM` for a custom role.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomRoleAssignmentSchema CustomRoleAssignmentSchema

// NewCustomRoleAssignmentSchema instantiates a new CustomRoleAssignmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRoleAssignmentSchema() *CustomRoleAssignmentSchema {
	this := CustomRoleAssignmentSchema{}
	return &this
}

// NewCustomRoleAssignmentSchemaWithDefaults instantiates a new CustomRoleAssignmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRoleAssignmentSchemaWithDefaults() *CustomRoleAssignmentSchema {
	this := CustomRoleAssignmentSchema{}
	return &this
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *CustomRoleAssignmentSchema) GetResourceSet() string {
	if o == nil || IsNil(o.ResourceSet) {
		var ret string
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRoleAssignmentSchema) GetResourceSetOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceSet) {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *CustomRoleAssignmentSchema) HasResourceSet() bool {
	if o != nil && !IsNil(o.ResourceSet) {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given string and assigns it to the ResourceSet field.
func (o *CustomRoleAssignmentSchema) SetResourceSet(v string) {
	o.ResourceSet = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *CustomRoleAssignmentSchema) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRoleAssignmentSchema) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *CustomRoleAssignmentSchema) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *CustomRoleAssignmentSchema) SetRole(v string) {
	o.Role = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomRoleAssignmentSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRoleAssignmentSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomRoleAssignmentSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomRoleAssignmentSchema) SetType(v string) {
	o.Type = &v
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
	if !IsNil(o.ResourceSet) {
		toSerialize["resource-set"] = o.ResourceSet
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomRoleAssignmentSchema) UnmarshalJSON(data []byte) (err error) {
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
