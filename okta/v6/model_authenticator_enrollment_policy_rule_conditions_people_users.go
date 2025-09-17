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

// checks if the AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers{}

// AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers Specifies a set of users to be included or excluded
type AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers struct {
	// Users to be excluded
	Exclude              []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers

// NewAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers instantiates a new AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers() *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers {
	this := AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsersWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsersWithDefaults() *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers {
	this := AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) SetExclude(v []string) {
	o.Exclude = v
}

func (o AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers := _AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers(varAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers struct {
	value *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) Get() *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) Set(val *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers(val *AuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers {
	return &NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditionsPeopleUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
