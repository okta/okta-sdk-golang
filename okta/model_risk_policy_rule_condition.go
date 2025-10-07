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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the RiskPolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicyRuleCondition{}

// RiskPolicyRuleCondition struct for RiskPolicyRuleCondition
type RiskPolicyRuleCondition struct {
	Behaviors            []string `json:"behaviors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskPolicyRuleCondition RiskPolicyRuleCondition

// NewRiskPolicyRuleCondition instantiates a new RiskPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicyRuleCondition() *RiskPolicyRuleCondition {
	this := RiskPolicyRuleCondition{}
	return &this
}

// NewRiskPolicyRuleConditionWithDefaults instantiates a new RiskPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicyRuleConditionWithDefaults() *RiskPolicyRuleCondition {
	this := RiskPolicyRuleCondition{}
	return &this
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *RiskPolicyRuleCondition) GetBehaviors() []string {
	if o == nil || IsNil(o.Behaviors) {
		var ret []string
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyRuleCondition) GetBehaviorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *RiskPolicyRuleCondition) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []string and assigns it to the Behaviors field.
func (o *RiskPolicyRuleCondition) SetBehaviors(v []string) {
	o.Behaviors = v
}

func (o RiskPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskPolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varRiskPolicyRuleCondition := _RiskPolicyRuleCondition{}

	err = json.Unmarshal(data, &varRiskPolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = RiskPolicyRuleCondition(varRiskPolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "behaviors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskPolicyRuleCondition struct {
	value *RiskPolicyRuleCondition
	isSet bool
}

func (v NullableRiskPolicyRuleCondition) Get() *RiskPolicyRuleCondition {
	return v.value
}

func (v *NullableRiskPolicyRuleCondition) Set(val *RiskPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicyRuleCondition(val *RiskPolicyRuleCondition) *NullableRiskPolicyRuleCondition {
	return &NullableRiskPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableRiskPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
