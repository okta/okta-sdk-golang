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

// AuthorizationServerPolicyPeopleCondition Identifies Users and Groups that are used together
type AuthorizationServerPolicyPeopleCondition struct {
	Groups *AuthorizationServerPolicyRuleGroupCondition `json:"groups,omitempty"`
	Users *AuthorizationServerPolicyRuleUserCondition `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyPeopleCondition AuthorizationServerPolicyPeopleCondition

// NewAuthorizationServerPolicyPeopleCondition instantiates a new AuthorizationServerPolicyPeopleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyPeopleCondition() *AuthorizationServerPolicyPeopleCondition {
	this := AuthorizationServerPolicyPeopleCondition{}
	return &this
}

// NewAuthorizationServerPolicyPeopleConditionWithDefaults instantiates a new AuthorizationServerPolicyPeopleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyPeopleConditionWithDefaults() *AuthorizationServerPolicyPeopleCondition {
	this := AuthorizationServerPolicyPeopleCondition{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyPeopleCondition) GetGroups() AuthorizationServerPolicyRuleGroupCondition {
	if o == nil || o.Groups == nil {
		var ret AuthorizationServerPolicyRuleGroupCondition
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyPeopleCondition) GetGroupsOk() (*AuthorizationServerPolicyRuleGroupCondition, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyPeopleCondition) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given AuthorizationServerPolicyRuleGroupCondition and assigns it to the Groups field.
func (o *AuthorizationServerPolicyPeopleCondition) SetGroups(v AuthorizationServerPolicyRuleGroupCondition) {
	o.Groups = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyPeopleCondition) GetUsers() AuthorizationServerPolicyRuleUserCondition {
	if o == nil || o.Users == nil {
		var ret AuthorizationServerPolicyRuleUserCondition
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyPeopleCondition) GetUsersOk() (*AuthorizationServerPolicyRuleUserCondition, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyPeopleCondition) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given AuthorizationServerPolicyRuleUserCondition and assigns it to the Users field.
func (o *AuthorizationServerPolicyPeopleCondition) SetUsers(v AuthorizationServerPolicyRuleUserCondition) {
	o.Users = &v
}

func (o AuthorizationServerPolicyPeopleCondition) MarshalJSON() ([]byte, error) {
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

func (o *AuthorizationServerPolicyPeopleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyPeopleCondition := _AuthorizationServerPolicyPeopleCondition{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyPeopleCondition)
	if err == nil {
		*o = AuthorizationServerPolicyPeopleCondition(varAuthorizationServerPolicyPeopleCondition)
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

type NullableAuthorizationServerPolicyPeopleCondition struct {
	value *AuthorizationServerPolicyPeopleCondition
	isSet bool
}

func (v NullableAuthorizationServerPolicyPeopleCondition) Get() *AuthorizationServerPolicyPeopleCondition {
	return v.value
}

func (v *NullableAuthorizationServerPolicyPeopleCondition) Set(val *AuthorizationServerPolicyPeopleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyPeopleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyPeopleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyPeopleCondition(val *AuthorizationServerPolicyPeopleCondition) *NullableAuthorizationServerPolicyPeopleCondition {
	return &NullableAuthorizationServerPolicyPeopleCondition{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyPeopleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyPeopleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

