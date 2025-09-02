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

// EntityRiskScorePolicyRuleCondition <x-lifecycle class=\"oie\"></x-lifecycle> The risk score level of the entity risk policy rule
type EntityRiskScorePolicyRuleCondition struct {
	Level string `json:"level"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskScorePolicyRuleCondition EntityRiskScorePolicyRuleCondition

// NewEntityRiskScorePolicyRuleCondition instantiates a new EntityRiskScorePolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskScorePolicyRuleCondition(level string) *EntityRiskScorePolicyRuleCondition {
	this := EntityRiskScorePolicyRuleCondition{}
	this.Level = level
	return &this
}

// NewEntityRiskScorePolicyRuleConditionWithDefaults instantiates a new EntityRiskScorePolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskScorePolicyRuleConditionWithDefaults() *EntityRiskScorePolicyRuleCondition {
	this := EntityRiskScorePolicyRuleCondition{}
	return &this
}

// GetLevel returns the Level field value
func (o *EntityRiskScorePolicyRuleCondition) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *EntityRiskScorePolicyRuleCondition) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *EntityRiskScorePolicyRuleCondition) SetLevel(v string) {
	o.Level = v
}

func (o EntityRiskScorePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityRiskScorePolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varEntityRiskScorePolicyRuleCondition := _EntityRiskScorePolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varEntityRiskScorePolicyRuleCondition)
	if err == nil {
		*o = EntityRiskScorePolicyRuleCondition(varEntityRiskScorePolicyRuleCondition)
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

type NullableEntityRiskScorePolicyRuleCondition struct {
	value *EntityRiskScorePolicyRuleCondition
	isSet bool
}

func (v NullableEntityRiskScorePolicyRuleCondition) Get() *EntityRiskScorePolicyRuleCondition {
	return v.value
}

func (v *NullableEntityRiskScorePolicyRuleCondition) Set(val *EntityRiskScorePolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskScorePolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskScorePolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskScorePolicyRuleCondition(val *EntityRiskScorePolicyRuleCondition) *NullableEntityRiskScorePolicyRuleCondition {
	return &NullableEntityRiskScorePolicyRuleCondition{value: val, isSet: true}
}

func (v NullableEntityRiskScorePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskScorePolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

