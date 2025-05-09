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

// ContinuousAccessPolicyRuleAllOfConditions struct for ContinuousAccessPolicyRuleAllOfConditions
type ContinuousAccessPolicyRuleAllOfConditions struct {
	People *PolicyPeopleCondition `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleAllOfConditions ContinuousAccessPolicyRuleAllOfConditions

// NewContinuousAccessPolicyRuleAllOfConditions instantiates a new ContinuousAccessPolicyRuleAllOfConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleAllOfConditions() *ContinuousAccessPolicyRuleAllOfConditions {
	this := ContinuousAccessPolicyRuleAllOfConditions{}
	return &this
}

// NewContinuousAccessPolicyRuleAllOfConditionsWithDefaults instantiates a new ContinuousAccessPolicyRuleAllOfConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleAllOfConditionsWithDefaults() *ContinuousAccessPolicyRuleAllOfConditions {
	this := ContinuousAccessPolicyRuleAllOfConditions{}
	return &this
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleAllOfConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || o.People == nil {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleAllOfConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleAllOfConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *ContinuousAccessPolicyRuleAllOfConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

func (o ContinuousAccessPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.People != nil {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleAllOfConditions) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleAllOfConditions := _ContinuousAccessPolicyRuleAllOfConditions{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleAllOfConditions)
	if err == nil {
		*o = ContinuousAccessPolicyRuleAllOfConditions(varContinuousAccessPolicyRuleAllOfConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleAllOfConditions struct {
	value *ContinuousAccessPolicyRuleAllOfConditions
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleAllOfConditions) Get() *ContinuousAccessPolicyRuleAllOfConditions {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleAllOfConditions) Set(val *ContinuousAccessPolicyRuleAllOfConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleAllOfConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleAllOfConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleAllOfConditions(val *ContinuousAccessPolicyRuleAllOfConditions) *NullableContinuousAccessPolicyRuleAllOfConditions {
	return &NullableContinuousAccessPolicyRuleAllOfConditions{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleAllOfConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

