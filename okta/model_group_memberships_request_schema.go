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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GroupMembershipsRequestSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMembershipsRequestSchema{}

// GroupMembershipsRequestSchema struct for GroupMembershipsRequestSchema
type GroupMembershipsRequestSchema struct {
	// A list of app user external IDs to be inserted in this group in Okta
	MemberExternalIds    []string `json:"memberExternalIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupMembershipsRequestSchema GroupMembershipsRequestSchema

// NewGroupMembershipsRequestSchema instantiates a new GroupMembershipsRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMembershipsRequestSchema() *GroupMembershipsRequestSchema {
	this := GroupMembershipsRequestSchema{}
	return &this
}

// NewGroupMembershipsRequestSchemaWithDefaults instantiates a new GroupMembershipsRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMembershipsRequestSchemaWithDefaults() *GroupMembershipsRequestSchema {
	this := GroupMembershipsRequestSchema{}
	return &this
}

// GetMemberExternalIds returns the MemberExternalIds field value if set, zero value otherwise.
func (o *GroupMembershipsRequestSchema) GetMemberExternalIds() []string {
	if o == nil || IsNil(o.MemberExternalIds) {
		var ret []string
		return ret
	}
	return o.MemberExternalIds
}

// GetMemberExternalIdsOk returns a tuple with the MemberExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembershipsRequestSchema) GetMemberExternalIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MemberExternalIds) {
		return nil, false
	}
	return o.MemberExternalIds, true
}

// HasMemberExternalIds returns a boolean if a field has been set.
func (o *GroupMembershipsRequestSchema) HasMemberExternalIds() bool {
	if o != nil && !IsNil(o.MemberExternalIds) {
		return true
	}

	return false
}

// SetMemberExternalIds gets a reference to the given []string and assigns it to the MemberExternalIds field.
func (o *GroupMembershipsRequestSchema) SetMemberExternalIds(v []string) {
	o.MemberExternalIds = v
}

func (o GroupMembershipsRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMembershipsRequestSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberExternalIds) {
		toSerialize["memberExternalIds"] = o.MemberExternalIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupMembershipsRequestSchema) UnmarshalJSON(data []byte) (err error) {
	varGroupMembershipsRequestSchema := _GroupMembershipsRequestSchema{}

	err = json.Unmarshal(data, &varGroupMembershipsRequestSchema)

	if err != nil {
		return err
	}

	*o = GroupMembershipsRequestSchema(varGroupMembershipsRequestSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "memberExternalIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupMembershipsRequestSchema struct {
	value *GroupMembershipsRequestSchema
	isSet bool
}

func (v NullableGroupMembershipsRequestSchema) Get() *GroupMembershipsRequestSchema {
	return v.value
}

func (v *NullableGroupMembershipsRequestSchema) Set(val *GroupMembershipsRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembershipsRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembershipsRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembershipsRequestSchema(val *GroupMembershipsRequestSchema) *NullableGroupMembershipsRequestSchema {
	return &NullableGroupMembershipsRequestSchema{value: val, isSet: true}
}

func (v NullableGroupMembershipsRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembershipsRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
