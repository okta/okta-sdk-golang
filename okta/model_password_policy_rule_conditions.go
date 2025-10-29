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

// checks if the PasswordPolicyRuleConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRuleConditions{}

// PasswordPolicyRuleConditions Specifies conditions that must be met during policy evaluation to apply the rule. All policy conditions and conditions for at least one rule must be met to apply the settings specified in the policy and the associated rule.
type PasswordPolicyRuleConditions struct {
	Network              *PolicyNetworkCondition `json:"network,omitempty"`
	People               *PolicyPeopleCondition  `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRuleConditions PasswordPolicyRuleConditions

// NewPasswordPolicyRuleConditions instantiates a new PasswordPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRuleConditions() *PasswordPolicyRuleConditions {
	this := PasswordPolicyRuleConditions{}
	return &this
}

// NewPasswordPolicyRuleConditionsWithDefaults instantiates a new PasswordPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRuleConditionsWithDefaults() *PasswordPolicyRuleConditions {
	this := PasswordPolicyRuleConditions{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *PasswordPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *PasswordPolicyRuleConditions) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *PasswordPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *PasswordPolicyRuleConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || IsNil(o.People) {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *PasswordPolicyRuleConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *PasswordPolicyRuleConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

func (o PasswordPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRuleConditions) ToMap() (map[string]interface{}, error) {
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

func (o *PasswordPolicyRuleConditions) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRuleConditions := _PasswordPolicyRuleConditions{}

	err = json.Unmarshal(data, &varPasswordPolicyRuleConditions)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRuleConditions(varPasswordPolicyRuleConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRuleConditions struct {
	value *PasswordPolicyRuleConditions
	isSet bool
}

func (v NullablePasswordPolicyRuleConditions) Get() *PasswordPolicyRuleConditions {
	return v.value
}

func (v *NullablePasswordPolicyRuleConditions) Set(val *PasswordPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRuleConditions(val *PasswordPolicyRuleConditions) *NullablePasswordPolicyRuleConditions {
	return &NullablePasswordPolicyRuleConditions{value: val, isSet: true}
}

func (v NullablePasswordPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
