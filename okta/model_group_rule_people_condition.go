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

// GroupRulePeopleCondition struct for GroupRulePeopleCondition
type GroupRulePeopleCondition struct {
	Groups *GroupRuleGroupCondition `json:"groups,omitempty"`
	Users *GroupRuleUserCondition `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRulePeopleCondition GroupRulePeopleCondition

// NewGroupRulePeopleCondition instantiates a new GroupRulePeopleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRulePeopleCondition() *GroupRulePeopleCondition {
	this := GroupRulePeopleCondition{}
	return &this
}

// NewGroupRulePeopleConditionWithDefaults instantiates a new GroupRulePeopleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRulePeopleConditionWithDefaults() *GroupRulePeopleCondition {
	this := GroupRulePeopleCondition{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *GroupRulePeopleCondition) GetGroups() GroupRuleGroupCondition {
	if o == nil || o.Groups == nil {
		var ret GroupRuleGroupCondition
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRulePeopleCondition) GetGroupsOk() (*GroupRuleGroupCondition, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupRulePeopleCondition) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given GroupRuleGroupCondition and assigns it to the Groups field.
func (o *GroupRulePeopleCondition) SetGroups(v GroupRuleGroupCondition) {
	o.Groups = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GroupRulePeopleCondition) GetUsers() GroupRuleUserCondition {
	if o == nil || o.Users == nil {
		var ret GroupRuleUserCondition
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRulePeopleCondition) GetUsersOk() (*GroupRuleUserCondition, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupRulePeopleCondition) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given GroupRuleUserCondition and assigns it to the Users field.
func (o *GroupRulePeopleCondition) SetUsers(v GroupRuleUserCondition) {
	o.Users = &v
}

func (o GroupRulePeopleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupRulePeopleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varGroupRulePeopleCondition := _GroupRulePeopleCondition{}

	err = json.Unmarshal(bytes, &varGroupRulePeopleCondition)
	if err == nil {
		*o = GroupRulePeopleCondition(varGroupRulePeopleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupRulePeopleCondition struct {
	value *GroupRulePeopleCondition
	isSet bool
}

func (v NullableGroupRulePeopleCondition) Get() *GroupRulePeopleCondition {
	return v.value
}

func (v *NullableGroupRulePeopleCondition) Set(val *GroupRulePeopleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRulePeopleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRulePeopleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRulePeopleCondition(val *GroupRulePeopleCondition) *NullableGroupRulePeopleCondition {
	return &NullableGroupRulePeopleCondition{value: val, isSet: true}
}

func (v NullableGroupRulePeopleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRulePeopleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

