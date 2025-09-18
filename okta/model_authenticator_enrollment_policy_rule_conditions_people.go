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

// checks if the AuthenticatorEnrollmentPolicyRuleConditionsPeople type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyRuleConditionsPeople{}

// AuthenticatorEnrollmentPolicyRuleConditionsPeople Identifies users and groups that are used together
type AuthenticatorEnrollmentPolicyRuleConditionsPeople struct {
	Users                *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyRuleConditionsPeople AuthenticatorEnrollmentPolicyRuleConditionsPeople

// NewAuthenticatorEnrollmentPolicyRuleConditionsPeople instantiates a new AuthenticatorEnrollmentPolicyRuleConditionsPeople object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyRuleConditionsPeople() *AuthenticatorEnrollmentPolicyRuleConditionsPeople {
	this := AuthenticatorEnrollmentPolicyRuleConditionsPeople{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyRuleConditionsPeopleWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRuleConditionsPeople object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyRuleConditionsPeopleWithDefaults() *AuthenticatorEnrollmentPolicyRuleConditionsPeople {
	this := AuthenticatorEnrollmentPolicyRuleConditionsPeople{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeople) GetUsers() AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers {
	if o == nil || IsNil(o.Users) {
		var ret AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeople) GetUsersOk() (*AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeople) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers and assigns it to the Users field.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeople) SetUsers(v AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) {
	o.Users = &v
}

func (o AuthenticatorEnrollmentPolicyRuleConditionsPeople) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyRuleConditionsPeople) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeople) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyRuleConditionsPeople := _AuthenticatorEnrollmentPolicyRuleConditionsPeople{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyRuleConditionsPeople)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyRuleConditionsPeople(varAuthenticatorEnrollmentPolicyRuleConditionsPeople)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople struct {
	value *AuthenticatorEnrollmentPolicyRuleConditionsPeople
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople) Get() *AuthenticatorEnrollmentPolicyRuleConditionsPeople {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople) Set(val *AuthenticatorEnrollmentPolicyRuleConditionsPeople) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyRuleConditionsPeople(val *AuthenticatorEnrollmentPolicyRuleConditionsPeople) *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople {
	return &NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeople) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
