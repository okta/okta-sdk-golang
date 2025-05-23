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

// PolicyPeopleCondition Identifies Users and Groups that are used together
type PolicyPeopleCondition struct {
	Groups *GroupCondition `json:"groups,omitempty"`
	Users *UserCondition `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyPeopleCondition PolicyPeopleCondition

// NewPolicyPeopleCondition instantiates a new PolicyPeopleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyPeopleCondition() *PolicyPeopleCondition {
	this := PolicyPeopleCondition{}
	return &this
}

// NewPolicyPeopleConditionWithDefaults instantiates a new PolicyPeopleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyPeopleConditionWithDefaults() *PolicyPeopleCondition {
	this := PolicyPeopleCondition{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PolicyPeopleCondition) GetGroups() GroupCondition {
	if o == nil || o.Groups == nil {
		var ret GroupCondition
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPeopleCondition) GetGroupsOk() (*GroupCondition, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PolicyPeopleCondition) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given GroupCondition and assigns it to the Groups field.
func (o *PolicyPeopleCondition) SetGroups(v GroupCondition) {
	o.Groups = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PolicyPeopleCondition) GetUsers() UserCondition {
	if o == nil || o.Users == nil {
		var ret UserCondition
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPeopleCondition) GetUsersOk() (*UserCondition, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PolicyPeopleCondition) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given UserCondition and assigns it to the Users field.
func (o *PolicyPeopleCondition) SetUsers(v UserCondition) {
	o.Users = &v
}

func (o PolicyPeopleCondition) MarshalJSON() ([]byte, error) {
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

func (o *PolicyPeopleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyPeopleCondition := _PolicyPeopleCondition{}

	err = json.Unmarshal(bytes, &varPolicyPeopleCondition)
	if err == nil {
		*o = PolicyPeopleCondition(varPolicyPeopleCondition)
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

type NullablePolicyPeopleCondition struct {
	value *PolicyPeopleCondition
	isSet bool
}

func (v NullablePolicyPeopleCondition) Get() *PolicyPeopleCondition {
	return v.value
}

func (v *NullablePolicyPeopleCondition) Set(val *PolicyPeopleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyPeopleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyPeopleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyPeopleCondition(val *PolicyPeopleCondition) *NullablePolicyPeopleCondition {
	return &NullablePolicyPeopleCondition{value: val, isSet: true}
}

func (v NullablePolicyPeopleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyPeopleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

