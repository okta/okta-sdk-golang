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

// checks if the PostAuthSessionPolicyRuleAllOfConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAuthSessionPolicyRuleAllOfConditions{}

// PostAuthSessionPolicyRuleAllOfConditions Specifies conditions that must be met during policy evaluation to apply the rule. All policy conditions and conditions for at least one rule must be met to apply the settings specified in the policy and the associated rule.
type PostAuthSessionPolicyRuleAllOfConditions struct {
	People               *PolicyPeopleCondition `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionPolicyRuleAllOfConditions PostAuthSessionPolicyRuleAllOfConditions

// NewPostAuthSessionPolicyRuleAllOfConditions instantiates a new PostAuthSessionPolicyRuleAllOfConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionPolicyRuleAllOfConditions() *PostAuthSessionPolicyRuleAllOfConditions {
	this := PostAuthSessionPolicyRuleAllOfConditions{}
	return &this
}

// NewPostAuthSessionPolicyRuleAllOfConditionsWithDefaults instantiates a new PostAuthSessionPolicyRuleAllOfConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionPolicyRuleAllOfConditionsWithDefaults() *PostAuthSessionPolicyRuleAllOfConditions {
	this := PostAuthSessionPolicyRuleAllOfConditions{}
	return &this
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRuleAllOfConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || IsNil(o.People) {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRuleAllOfConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRuleAllOfConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *PostAuthSessionPolicyRuleAllOfConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

func (o PostAuthSessionPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthSessionPolicyRuleAllOfConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAuthSessionPolicyRuleAllOfConditions) UnmarshalJSON(data []byte) (err error) {
	varPostAuthSessionPolicyRuleAllOfConditions := _PostAuthSessionPolicyRuleAllOfConditions{}

	err = json.Unmarshal(data, &varPostAuthSessionPolicyRuleAllOfConditions)

	if err != nil {
		return err
	}

	*o = PostAuthSessionPolicyRuleAllOfConditions(varPostAuthSessionPolicyRuleAllOfConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAuthSessionPolicyRuleAllOfConditions struct {
	value *PostAuthSessionPolicyRuleAllOfConditions
	isSet bool
}

func (v NullablePostAuthSessionPolicyRuleAllOfConditions) Get() *PostAuthSessionPolicyRuleAllOfConditions {
	return v.value
}

func (v *NullablePostAuthSessionPolicyRuleAllOfConditions) Set(val *PostAuthSessionPolicyRuleAllOfConditions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionPolicyRuleAllOfConditions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionPolicyRuleAllOfConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionPolicyRuleAllOfConditions(val *PostAuthSessionPolicyRuleAllOfConditions) *NullablePostAuthSessionPolicyRuleAllOfConditions {
	return &NullablePostAuthSessionPolicyRuleAllOfConditions{value: val, isSet: true}
}

func (v NullablePostAuthSessionPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionPolicyRuleAllOfConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
