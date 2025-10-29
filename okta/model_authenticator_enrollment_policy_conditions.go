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

// checks if the AuthenticatorEnrollmentPolicyConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyConditions{}

// AuthenticatorEnrollmentPolicyConditions Specifies the conditions that must be met during policy evaluation to apply the policy
type AuthenticatorEnrollmentPolicyConditions struct {
	People               *AuthenticatorEnrollmentPolicyConditionsAllOfPeople `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyConditions AuthenticatorEnrollmentPolicyConditions

// NewAuthenticatorEnrollmentPolicyConditions instantiates a new AuthenticatorEnrollmentPolicyConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyConditions() *AuthenticatorEnrollmentPolicyConditions {
	this := AuthenticatorEnrollmentPolicyConditions{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyConditionsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyConditionsWithDefaults() *AuthenticatorEnrollmentPolicyConditions {
	this := AuthenticatorEnrollmentPolicyConditions{}
	return &this
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyConditions) GetPeople() AuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	if o == nil || IsNil(o.People) {
		var ret AuthenticatorEnrollmentPolicyConditionsAllOfPeople
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyConditions) GetPeopleOk() (*AuthenticatorEnrollmentPolicyConditionsAllOfPeople, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given AuthenticatorEnrollmentPolicyConditionsAllOfPeople and assigns it to the People field.
func (o *AuthenticatorEnrollmentPolicyConditions) SetPeople(v AuthenticatorEnrollmentPolicyConditionsAllOfPeople) {
	o.People = &v
}

func (o AuthenticatorEnrollmentPolicyConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyConditions) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyConditions := _AuthenticatorEnrollmentPolicyConditions{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyConditions)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyConditions(varAuthenticatorEnrollmentPolicyConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyConditions struct {
	value *AuthenticatorEnrollmentPolicyConditions
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyConditions) Get() *AuthenticatorEnrollmentPolicyConditions {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyConditions) Set(val *AuthenticatorEnrollmentPolicyConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyConditions(val *AuthenticatorEnrollmentPolicyConditions) *NullableAuthenticatorEnrollmentPolicyConditions {
	return &NullableAuthenticatorEnrollmentPolicyConditions{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
