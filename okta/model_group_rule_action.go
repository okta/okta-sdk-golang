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

// GroupRuleAction struct for GroupRuleAction
type GroupRuleAction struct {
	AssignUserToGroups *GroupRuleGroupAssignment `json:"assignUserToGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRuleAction GroupRuleAction

// NewGroupRuleAction instantiates a new GroupRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleAction() *GroupRuleAction {
	this := GroupRuleAction{}
	return &this
}

// NewGroupRuleActionWithDefaults instantiates a new GroupRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleActionWithDefaults() *GroupRuleAction {
	this := GroupRuleAction{}
	return &this
}

// GetAssignUserToGroups returns the AssignUserToGroups field value if set, zero value otherwise.
func (o *GroupRuleAction) GetAssignUserToGroups() GroupRuleGroupAssignment {
	if o == nil || o.AssignUserToGroups == nil {
		var ret GroupRuleGroupAssignment
		return ret
	}
	return *o.AssignUserToGroups
}

// GetAssignUserToGroupsOk returns a tuple with the AssignUserToGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleAction) GetAssignUserToGroupsOk() (*GroupRuleGroupAssignment, bool) {
	if o == nil || o.AssignUserToGroups == nil {
		return nil, false
	}
	return o.AssignUserToGroups, true
}

// HasAssignUserToGroups returns a boolean if a field has been set.
func (o *GroupRuleAction) HasAssignUserToGroups() bool {
	if o != nil && o.AssignUserToGroups != nil {
		return true
	}

	return false
}

// SetAssignUserToGroups gets a reference to the given GroupRuleGroupAssignment and assigns it to the AssignUserToGroups field.
func (o *GroupRuleAction) SetAssignUserToGroups(v GroupRuleGroupAssignment) {
	o.AssignUserToGroups = &v
}

func (o GroupRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignUserToGroups != nil {
		toSerialize["assignUserToGroups"] = o.AssignUserToGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	varGroupRuleAction := _GroupRuleAction{}

	err = json.Unmarshal(bytes, &varGroupRuleAction)
	if err == nil {
		*o = GroupRuleAction(varGroupRuleAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "assignUserToGroups")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupRuleAction struct {
	value *GroupRuleAction
	isSet bool
}

func (v NullableGroupRuleAction) Get() *GroupRuleAction {
	return v.value
}

func (v *NullableGroupRuleAction) Set(val *GroupRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleAction(val *GroupRuleAction) *NullableGroupRuleAction {
	return &NullableGroupRuleAction{value: val, isSet: true}
}

func (v NullableGroupRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

