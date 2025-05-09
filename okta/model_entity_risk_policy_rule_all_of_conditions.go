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

// EntityRiskPolicyRuleAllOfConditions struct for EntityRiskPolicyRuleAllOfConditions
type EntityRiskPolicyRuleAllOfConditions struct {
	People *PolicyPeopleCondition `json:"people,omitempty"`
	RiskDetectionTypes *EntityRiskPolicyRuleAllOfConditionsRiskDetectionTypes `json:"riskDetectionTypes,omitempty"`
	EntityRisk *EntityRiskPolicyRuleAllOfConditionsEntityRisk `json:"EntityRisk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleAllOfConditions EntityRiskPolicyRuleAllOfConditions

// NewEntityRiskPolicyRuleAllOfConditions instantiates a new EntityRiskPolicyRuleAllOfConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleAllOfConditions() *EntityRiskPolicyRuleAllOfConditions {
	this := EntityRiskPolicyRuleAllOfConditions{}
	return &this
}

// NewEntityRiskPolicyRuleAllOfConditionsWithDefaults instantiates a new EntityRiskPolicyRuleAllOfConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleAllOfConditionsWithDefaults() *EntityRiskPolicyRuleAllOfConditions {
	this := EntityRiskPolicyRuleAllOfConditions{}
	return &this
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleAllOfConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || o.People == nil {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleAllOfConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleAllOfConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *EntityRiskPolicyRuleAllOfConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

// GetRiskDetectionTypes returns the RiskDetectionTypes field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleAllOfConditions) GetRiskDetectionTypes() EntityRiskPolicyRuleAllOfConditionsRiskDetectionTypes {
	if o == nil || o.RiskDetectionTypes == nil {
		var ret EntityRiskPolicyRuleAllOfConditionsRiskDetectionTypes
		return ret
	}
	return *o.RiskDetectionTypes
}

// GetRiskDetectionTypesOk returns a tuple with the RiskDetectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleAllOfConditions) GetRiskDetectionTypesOk() (*EntityRiskPolicyRuleAllOfConditionsRiskDetectionTypes, bool) {
	if o == nil || o.RiskDetectionTypes == nil {
		return nil, false
	}
	return o.RiskDetectionTypes, true
}

// HasRiskDetectionTypes returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleAllOfConditions) HasRiskDetectionTypes() bool {
	if o != nil && o.RiskDetectionTypes != nil {
		return true
	}

	return false
}

// SetRiskDetectionTypes gets a reference to the given EntityRiskPolicyRuleAllOfConditionsRiskDetectionTypes and assigns it to the RiskDetectionTypes field.
func (o *EntityRiskPolicyRuleAllOfConditions) SetRiskDetectionTypes(v EntityRiskPolicyRuleAllOfConditionsRiskDetectionTypes) {
	o.RiskDetectionTypes = &v
}

// GetEntityRisk returns the EntityRisk field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleAllOfConditions) GetEntityRisk() EntityRiskPolicyRuleAllOfConditionsEntityRisk {
	if o == nil || o.EntityRisk == nil {
		var ret EntityRiskPolicyRuleAllOfConditionsEntityRisk
		return ret
	}
	return *o.EntityRisk
}

// GetEntityRiskOk returns a tuple with the EntityRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleAllOfConditions) GetEntityRiskOk() (*EntityRiskPolicyRuleAllOfConditionsEntityRisk, bool) {
	if o == nil || o.EntityRisk == nil {
		return nil, false
	}
	return o.EntityRisk, true
}

// HasEntityRisk returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleAllOfConditions) HasEntityRisk() bool {
	if o != nil && o.EntityRisk != nil {
		return true
	}

	return false
}

// SetEntityRisk gets a reference to the given EntityRiskPolicyRuleAllOfConditionsEntityRisk and assigns it to the EntityRisk field.
func (o *EntityRiskPolicyRuleAllOfConditions) SetEntityRisk(v EntityRiskPolicyRuleAllOfConditionsEntityRisk) {
	o.EntityRisk = &v
}

func (o EntityRiskPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	if o.RiskDetectionTypes != nil {
		toSerialize["riskDetectionTypes"] = o.RiskDetectionTypes
	}
	if o.EntityRisk != nil {
		toSerialize["EntityRisk"] = o.EntityRisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityRiskPolicyRuleAllOfConditions) UnmarshalJSON(bytes []byte) (err error) {
	varEntityRiskPolicyRuleAllOfConditions := _EntityRiskPolicyRuleAllOfConditions{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicyRuleAllOfConditions)
	if err == nil {
		*o = EntityRiskPolicyRuleAllOfConditions(varEntityRiskPolicyRuleAllOfConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "people")
		delete(additionalProperties, "riskDetectionTypes")
		delete(additionalProperties, "EntityRisk")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntityRiskPolicyRuleAllOfConditions struct {
	value *EntityRiskPolicyRuleAllOfConditions
	isSet bool
}

func (v NullableEntityRiskPolicyRuleAllOfConditions) Get() *EntityRiskPolicyRuleAllOfConditions {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleAllOfConditions) Set(val *EntityRiskPolicyRuleAllOfConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleAllOfConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleAllOfConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleAllOfConditions(val *EntityRiskPolicyRuleAllOfConditions) *NullableEntityRiskPolicyRuleAllOfConditions {
	return &NullableEntityRiskPolicyRuleAllOfConditions{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleAllOfConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

