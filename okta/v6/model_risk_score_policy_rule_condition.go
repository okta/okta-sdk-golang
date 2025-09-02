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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// RiskScorePolicyRuleCondition Specifies a particular level of risk to match on
type RiskScorePolicyRuleCondition struct {
	// The level to match
	Level string `json:"level"`
	AdditionalProperties map[string]interface{}
}

type _RiskScorePolicyRuleCondition RiskScorePolicyRuleCondition

// NewRiskScorePolicyRuleCondition instantiates a new RiskScorePolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskScorePolicyRuleCondition(level string) *RiskScorePolicyRuleCondition {
	this := RiskScorePolicyRuleCondition{}
	this.Level = level
	return &this
}

// NewRiskScorePolicyRuleConditionWithDefaults instantiates a new RiskScorePolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskScorePolicyRuleConditionWithDefaults() *RiskScorePolicyRuleCondition {
	this := RiskScorePolicyRuleCondition{}
	return &this
}

// GetLevel returns the Level field value
func (o *RiskScorePolicyRuleCondition) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *RiskScorePolicyRuleCondition) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *RiskScorePolicyRuleCondition) SetLevel(v string) {
	o.Level = v
}

func (o RiskScorePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskScorePolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varRiskScorePolicyRuleCondition := _RiskScorePolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varRiskScorePolicyRuleCondition)
	if err == nil {
		*o = RiskScorePolicyRuleCondition(varRiskScorePolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "level")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskScorePolicyRuleCondition struct {
	value *RiskScorePolicyRuleCondition
	isSet bool
}

func (v NullableRiskScorePolicyRuleCondition) Get() *RiskScorePolicyRuleCondition {
	return v.value
}

func (v *NullableRiskScorePolicyRuleCondition) Set(val *RiskScorePolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskScorePolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskScorePolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskScorePolicyRuleCondition(val *RiskScorePolicyRuleCondition) *NullableRiskScorePolicyRuleCondition {
	return &NullableRiskScorePolicyRuleCondition{value: val, isSet: true}
}

func (v NullableRiskScorePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskScorePolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

