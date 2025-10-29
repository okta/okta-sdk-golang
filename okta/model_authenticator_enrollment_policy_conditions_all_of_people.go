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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AuthenticatorEnrollmentPolicyConditionsAllOfPeople type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyConditionsAllOfPeople{}

// AuthenticatorEnrollmentPolicyConditionsAllOfPeople Identifies users and groups that are used together
type AuthenticatorEnrollmentPolicyConditionsAllOfPeople struct {
	Groups               *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups `json:"groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyConditionsAllOfPeople AuthenticatorEnrollmentPolicyConditionsAllOfPeople

// NewAuthenticatorEnrollmentPolicyConditionsAllOfPeople instantiates a new AuthenticatorEnrollmentPolicyConditionsAllOfPeople object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyConditionsAllOfPeople() *AuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	this := AuthenticatorEnrollmentPolicyConditionsAllOfPeople{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyConditionsAllOfPeopleWithDefaults instantiates a new AuthenticatorEnrollmentPolicyConditionsAllOfPeople object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyConditionsAllOfPeopleWithDefaults() *AuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	this := AuthenticatorEnrollmentPolicyConditionsAllOfPeople{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) GetGroups() AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups {
	if o == nil || IsNil(o.Groups) {
		var ret AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) GetGroupsOk() (*AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups and assigns it to the Groups field.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) SetGroups(v AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) {
	o.Groups = &v
}

func (o AuthenticatorEnrollmentPolicyConditionsAllOfPeople) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyConditionsAllOfPeople) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyConditionsAllOfPeople := _AuthenticatorEnrollmentPolicyConditionsAllOfPeople{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyConditionsAllOfPeople)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyConditionsAllOfPeople(varAuthenticatorEnrollmentPolicyConditionsAllOfPeople)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople struct {
	value *AuthenticatorEnrollmentPolicyConditionsAllOfPeople
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople) Get() *AuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople) Set(val *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople(val *AuthenticatorEnrollmentPolicyConditionsAllOfPeople) *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	return &NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeople) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
