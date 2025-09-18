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

// checks if the EntityRiskPolicyRuleConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRiskPolicyRuleConditions{}

// EntityRiskPolicyRuleConditions struct for EntityRiskPolicyRuleConditions
type EntityRiskPolicyRuleConditions struct {
	EntityRisk           *EntityRiskScorePolicyRuleCondition    `json:"entityRisk,omitempty"`
	People               *PolicyPeopleCondition                 `json:"people,omitempty"`
	RiskDetectionTypes   *RiskDetectionTypesPolicyRuleCondition `json:"riskDetectionTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleConditions EntityRiskPolicyRuleConditions

// NewEntityRiskPolicyRuleConditions instantiates a new EntityRiskPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleConditions() *EntityRiskPolicyRuleConditions {
	this := EntityRiskPolicyRuleConditions{}
	return &this
}

// NewEntityRiskPolicyRuleConditionsWithDefaults instantiates a new EntityRiskPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleConditionsWithDefaults() *EntityRiskPolicyRuleConditions {
	this := EntityRiskPolicyRuleConditions{}
	return &this
}

// GetEntityRisk returns the EntityRisk field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleConditions) GetEntityRisk() EntityRiskScorePolicyRuleCondition {
	if o == nil || IsNil(o.EntityRisk) {
		var ret EntityRiskScorePolicyRuleCondition
		return ret
	}
	return *o.EntityRisk
}

// GetEntityRiskOk returns a tuple with the EntityRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleConditions) GetEntityRiskOk() (*EntityRiskScorePolicyRuleCondition, bool) {
	if o == nil || IsNil(o.EntityRisk) {
		return nil, false
	}
	return o.EntityRisk, true
}

// HasEntityRisk returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleConditions) HasEntityRisk() bool {
	if o != nil && !IsNil(o.EntityRisk) {
		return true
	}

	return false
}

// SetEntityRisk gets a reference to the given EntityRiskScorePolicyRuleCondition and assigns it to the EntityRisk field.
func (o *EntityRiskPolicyRuleConditions) SetEntityRisk(v EntityRiskScorePolicyRuleCondition) {
	o.EntityRisk = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || IsNil(o.People) {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *EntityRiskPolicyRuleConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

// GetRiskDetectionTypes returns the RiskDetectionTypes field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleConditions) GetRiskDetectionTypes() RiskDetectionTypesPolicyRuleCondition {
	if o == nil || IsNil(o.RiskDetectionTypes) {
		var ret RiskDetectionTypesPolicyRuleCondition
		return ret
	}
	return *o.RiskDetectionTypes
}

// GetRiskDetectionTypesOk returns a tuple with the RiskDetectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleConditions) GetRiskDetectionTypesOk() (*RiskDetectionTypesPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.RiskDetectionTypes) {
		return nil, false
	}
	return o.RiskDetectionTypes, true
}

// HasRiskDetectionTypes returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleConditions) HasRiskDetectionTypes() bool {
	if o != nil && !IsNil(o.RiskDetectionTypes) {
		return true
	}

	return false
}

// SetRiskDetectionTypes gets a reference to the given RiskDetectionTypesPolicyRuleCondition and assigns it to the RiskDetectionTypes field.
func (o *EntityRiskPolicyRuleConditions) SetRiskDetectionTypes(v RiskDetectionTypesPolicyRuleCondition) {
	o.RiskDetectionTypes = &v
}

func (o EntityRiskPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRiskPolicyRuleConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityRisk) {
		toSerialize["entityRisk"] = o.EntityRisk
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	if !IsNil(o.RiskDetectionTypes) {
		toSerialize["riskDetectionTypes"] = o.RiskDetectionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntityRiskPolicyRuleConditions) UnmarshalJSON(data []byte) (err error) {
	varEntityRiskPolicyRuleConditions := _EntityRiskPolicyRuleConditions{}

	err = json.Unmarshal(data, &varEntityRiskPolicyRuleConditions)

	if err != nil {
		return err
	}

	*o = EntityRiskPolicyRuleConditions(varEntityRiskPolicyRuleConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entityRisk")
		delete(additionalProperties, "people")
		delete(additionalProperties, "riskDetectionTypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityRiskPolicyRuleConditions struct {
	value *EntityRiskPolicyRuleConditions
	isSet bool
}

func (v NullableEntityRiskPolicyRuleConditions) Get() *EntityRiskPolicyRuleConditions {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleConditions) Set(val *EntityRiskPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleConditions(val *EntityRiskPolicyRuleConditions) *NullableEntityRiskPolicyRuleConditions {
	return &NullableEntityRiskPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
