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

// checks if the AuthenticatorEnrollmentPolicyRuleConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyRuleConditions{}

// AuthenticatorEnrollmentPolicyRuleConditions Specifies conditions that must be met during policy evaluation to apply the rule. All policy conditions and conditions for at least one rule must be met to apply the settings specified in the policy and the associated rule.
type AuthenticatorEnrollmentPolicyRuleConditions struct {
	Network              *PolicyNetworkCondition                            `json:"network,omitempty"`
	People               *AuthenticatorEnrollmentPolicyRuleConditionsPeople `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyRuleConditions AuthenticatorEnrollmentPolicyRuleConditions

// NewAuthenticatorEnrollmentPolicyRuleConditions instantiates a new AuthenticatorEnrollmentPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyRuleConditions() *AuthenticatorEnrollmentPolicyRuleConditions {
	this := AuthenticatorEnrollmentPolicyRuleConditions{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyRuleConditionsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyRuleConditionsWithDefaults() *AuthenticatorEnrollmentPolicyRuleConditions {
	this := AuthenticatorEnrollmentPolicyRuleConditions{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) GetPeople() AuthenticatorEnrollmentPolicyRuleConditionsPeople {
	if o == nil || IsNil(o.People) {
		var ret AuthenticatorEnrollmentPolicyRuleConditionsPeople
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) GetPeopleOk() (*AuthenticatorEnrollmentPolicyRuleConditionsPeople, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given AuthenticatorEnrollmentPolicyRuleConditionsPeople and assigns it to the People field.
func (o *AuthenticatorEnrollmentPolicyRuleConditions) SetPeople(v AuthenticatorEnrollmentPolicyRuleConditionsPeople) {
	o.People = &v
}

func (o AuthenticatorEnrollmentPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyRuleConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyRuleConditions) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyRuleConditions := _AuthenticatorEnrollmentPolicyRuleConditions{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyRuleConditions)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyRuleConditions(varAuthenticatorEnrollmentPolicyRuleConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyRuleConditions struct {
	value *AuthenticatorEnrollmentPolicyRuleConditions
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditions) Get() *AuthenticatorEnrollmentPolicyRuleConditions {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditions) Set(val *AuthenticatorEnrollmentPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyRuleConditions(val *AuthenticatorEnrollmentPolicyRuleConditions) *NullableAuthenticatorEnrollmentPolicyRuleConditions {
	return &NullableAuthenticatorEnrollmentPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
