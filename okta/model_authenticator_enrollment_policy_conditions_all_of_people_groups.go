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

// checks if the AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups{}

// AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups Specifies a set of groups whose users are to be included or excluded
type AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups struct {
	// Groups to be included
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups

// NewAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups instantiates a new AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups() *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups {
	this := AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroupsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroupsWithDefaults() *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups {
	this := AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) SetInclude(v []string) {
	o.Include = v
}

func (o AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups := _AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups(varAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups struct {
	value *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) Get() *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) Set(val *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups(val *AuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups {
	return &NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyConditionsAllOfPeopleGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
