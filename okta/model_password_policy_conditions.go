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

// checks if the PasswordPolicyConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyConditions{}

// PasswordPolicyConditions Specifies the conditions that must be met during policy evaluation to apply the policy
type PasswordPolicyConditions struct {
	AuthProvider         *PasswordPolicyAuthenticationProviderCondition      `json:"authProvider,omitempty"`
	People               *AuthenticatorEnrollmentPolicyConditionsAllOfPeople `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyConditions PasswordPolicyConditions

// NewPasswordPolicyConditions instantiates a new PasswordPolicyConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyConditions() *PasswordPolicyConditions {
	this := PasswordPolicyConditions{}
	return &this
}

// NewPasswordPolicyConditionsWithDefaults instantiates a new PasswordPolicyConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyConditionsWithDefaults() *PasswordPolicyConditions {
	this := PasswordPolicyConditions{}
	return &this
}

// GetAuthProvider returns the AuthProvider field value if set, zero value otherwise.
func (o *PasswordPolicyConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition {
	if o == nil || IsNil(o.AuthProvider) {
		var ret PasswordPolicyAuthenticationProviderCondition
		return ret
	}
	return *o.AuthProvider
}

// GetAuthProviderOk returns a tuple with the AuthProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool) {
	if o == nil || IsNil(o.AuthProvider) {
		return nil, false
	}
	return o.AuthProvider, true
}

// HasAuthProvider returns a boolean if a field has been set.
func (o *PasswordPolicyConditions) HasAuthProvider() bool {
	if o != nil && !IsNil(o.AuthProvider) {
		return true
	}

	return false
}

// SetAuthProvider gets a reference to the given PasswordPolicyAuthenticationProviderCondition and assigns it to the AuthProvider field.
func (o *PasswordPolicyConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition) {
	o.AuthProvider = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *PasswordPolicyConditions) GetPeople() AuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	if o == nil || IsNil(o.People) {
		var ret AuthenticatorEnrollmentPolicyConditionsAllOfPeople
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyConditions) GetPeopleOk() (*AuthenticatorEnrollmentPolicyConditionsAllOfPeople, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *PasswordPolicyConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given AuthenticatorEnrollmentPolicyConditionsAllOfPeople and assigns it to the People field.
func (o *PasswordPolicyConditions) SetPeople(v AuthenticatorEnrollmentPolicyConditionsAllOfPeople) {
	o.People = &v
}

func (o PasswordPolicyConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthProvider) {
		toSerialize["authProvider"] = o.AuthProvider
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyConditions) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyConditions := _PasswordPolicyConditions{}

	err = json.Unmarshal(data, &varPasswordPolicyConditions)

	if err != nil {
		return err
	}

	*o = PasswordPolicyConditions(varPasswordPolicyConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authProvider")
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyConditions struct {
	value *PasswordPolicyConditions
	isSet bool
}

func (v NullablePasswordPolicyConditions) Get() *PasswordPolicyConditions {
	return v.value
}

func (v *NullablePasswordPolicyConditions) Set(val *PasswordPolicyConditions) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyConditions) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyConditions(val *PasswordPolicyConditions) *NullablePasswordPolicyConditions {
	return &NullablePasswordPolicyConditions{value: val, isSet: true}
}

func (v NullablePasswordPolicyConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
