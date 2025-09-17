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
	"fmt"
)

// checks if the EntityRiskScorePolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRiskScorePolicyRuleCondition{}

// EntityRiskScorePolicyRuleCondition <x-lifecycle class=\"oie\"></x-lifecycle> The risk score level of the entity risk policy rule
type EntityRiskScorePolicyRuleCondition struct {
	Level                string `json:"level"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRiskScorePolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntityRiskScorePolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEntityRiskScorePolicyRuleCondition := _EntityRiskScorePolicyRuleCondition{}

	err = json.Unmarshal(data, &varEntityRiskScorePolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = EntityRiskScorePolicyRuleCondition(varEntityRiskScorePolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "level")
		o.AdditionalProperties = additionalProperties
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
