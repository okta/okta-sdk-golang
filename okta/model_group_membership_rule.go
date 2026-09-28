/*
Okta Admin Management API

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

// checks if the GroupMembershipRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMembershipRule{}

// GroupMembershipRule A reference to a group rule that manages a user's membership in a group
type GroupMembershipRule struct {
	// The ID of the group rule
	Id *string `json:"id,omitempty"`
	// The name of the group rule
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupMembershipRule GroupMembershipRule

// NewGroupMembershipRule instantiates a new GroupMembershipRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMembershipRule() *GroupMembershipRule {
	this := GroupMembershipRule{}
	return &this
}

// NewGroupMembershipRuleWithDefaults instantiates a new GroupMembershipRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMembershipRuleWithDefaults() *GroupMembershipRule {
	this := GroupMembershipRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupMembershipRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembershipRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupMembershipRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupMembershipRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupMembershipRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembershipRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupMembershipRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupMembershipRule) SetName(v string) {
	o.Name = &v
}

func (o GroupMembershipRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMembershipRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupMembershipRule) UnmarshalJSON(data []byte) (err error) {
	varGroupMembershipRule := _GroupMembershipRule{}

	err = json.Unmarshal(data, &varGroupMembershipRule)

	if err != nil {
		return err
	}

	*o = GroupMembershipRule(varGroupMembershipRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupMembershipRule struct {
	value *GroupMembershipRule
	isSet bool
}

func (v NullableGroupMembershipRule) Get() *GroupMembershipRule {
	return v.value
}

func (v *NullableGroupMembershipRule) Set(val *GroupMembershipRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembershipRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembershipRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembershipRule(val *GroupMembershipRule) *NullableGroupMembershipRule {
	return &NullableGroupMembershipRule{value: val, isSet: true}
}

func (v NullableGroupMembershipRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembershipRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
