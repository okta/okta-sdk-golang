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

// GroupRuleGroupAssignment struct for GroupRuleGroupAssignment
type GroupRuleGroupAssignment struct {
	GroupIds []string `json:"groupIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRuleGroupAssignment GroupRuleGroupAssignment

// NewGroupRuleGroupAssignment instantiates a new GroupRuleGroupAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleGroupAssignment() *GroupRuleGroupAssignment {
	this := GroupRuleGroupAssignment{}
	return &this
}

// NewGroupRuleGroupAssignmentWithDefaults instantiates a new GroupRuleGroupAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleGroupAssignmentWithDefaults() *GroupRuleGroupAssignment {
	this := GroupRuleGroupAssignment{}
	return &this
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *GroupRuleGroupAssignment) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleGroupAssignment) GetGroupIdsOk() ([]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *GroupRuleGroupAssignment) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *GroupRuleGroupAssignment) SetGroupIds(v []string) {
	o.GroupIds = v
}

func (o GroupRuleGroupAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupRuleGroupAssignment) UnmarshalJSON(bytes []byte) (err error) {
	varGroupRuleGroupAssignment := _GroupRuleGroupAssignment{}

	err = json.Unmarshal(bytes, &varGroupRuleGroupAssignment)
	if err == nil {
		*o = GroupRuleGroupAssignment(varGroupRuleGroupAssignment)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "groupIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupRuleGroupAssignment struct {
	value *GroupRuleGroupAssignment
	isSet bool
}

func (v NullableGroupRuleGroupAssignment) Get() *GroupRuleGroupAssignment {
	return v.value
}

func (v *NullableGroupRuleGroupAssignment) Set(val *GroupRuleGroupAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleGroupAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleGroupAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleGroupAssignment(val *GroupRuleGroupAssignment) *NullableGroupRuleGroupAssignment {
	return &NullableGroupRuleGroupAssignment{value: val, isSet: true}
}

func (v NullableGroupRuleGroupAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleGroupAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

