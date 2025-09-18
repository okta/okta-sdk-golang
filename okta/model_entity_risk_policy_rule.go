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
	"reflect"
	"strings"
)

// checks if the EntityRiskPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRiskPolicyRule{}

// EntityRiskPolicyRule struct for EntityRiskPolicyRule
type EntityRiskPolicyRule struct {
	PolicyRule
	Actions              *EntityRiskPolicyRuleAllOfActions `json:"actions,omitempty"`
	Conditions           *EntityRiskPolicyRuleConditions   `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRule EntityRiskPolicyRule

// NewEntityRiskPolicyRule instantiates a new EntityRiskPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRule() *EntityRiskPolicyRule {
	this := EntityRiskPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewEntityRiskPolicyRuleWithDefaults instantiates a new EntityRiskPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleWithDefaults() *EntityRiskPolicyRule {
	this := EntityRiskPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *EntityRiskPolicyRule) GetActions() EntityRiskPolicyRuleAllOfActions {
	if o == nil || IsNil(o.Actions) {
		var ret EntityRiskPolicyRuleAllOfActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRule) GetActionsOk() (*EntityRiskPolicyRuleAllOfActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *EntityRiskPolicyRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given EntityRiskPolicyRuleAllOfActions and assigns it to the Actions field.
func (o *EntityRiskPolicyRule) SetActions(v EntityRiskPolicyRuleAllOfActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *EntityRiskPolicyRule) GetConditions() EntityRiskPolicyRuleConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret EntityRiskPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRule) GetConditionsOk() (*EntityRiskPolicyRuleConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *EntityRiskPolicyRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given EntityRiskPolicyRuleConditions and assigns it to the Conditions field.
func (o *EntityRiskPolicyRule) SetConditions(v EntityRiskPolicyRuleConditions) {
	o.Conditions = &v
}

func (o EntityRiskPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRiskPolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyRule, errPolicyRule := json.Marshal(o.PolicyRule)
	if errPolicyRule != nil {
		return map[string]interface{}{}, errPolicyRule
	}
	errPolicyRule = json.Unmarshal([]byte(serializedPolicyRule), &toSerialize)
	if errPolicyRule != nil {
		return map[string]interface{}{}, errPolicyRule
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntityRiskPolicyRule) UnmarshalJSON(data []byte) (err error) {
	type EntityRiskPolicyRuleWithoutEmbeddedStruct struct {
		Actions    *EntityRiskPolicyRuleAllOfActions `json:"actions,omitempty"`
		Conditions *EntityRiskPolicyRuleConditions   `json:"conditions,omitempty"`
	}

	varEntityRiskPolicyRuleWithoutEmbeddedStruct := EntityRiskPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEntityRiskPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varEntityRiskPolicyRule := _EntityRiskPolicyRule{}
		varEntityRiskPolicyRule.Actions = varEntityRiskPolicyRuleWithoutEmbeddedStruct.Actions
		varEntityRiskPolicyRule.Conditions = varEntityRiskPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = EntityRiskPolicyRule(varEntityRiskPolicyRule)
	} else {
		return err
	}

	varEntityRiskPolicyRule := _EntityRiskPolicyRule{}

	err = json.Unmarshal(data, &varEntityRiskPolicyRule)
	if err == nil {
		o.PolicyRule = varEntityRiskPolicyRule.PolicyRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")

		// remove fields from embedded structs
		reflectPolicyRule := reflect.ValueOf(o.PolicyRule)
		for i := 0; i < reflectPolicyRule.Type().NumField(); i++ {
			t := reflectPolicyRule.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityRiskPolicyRule struct {
	value *EntityRiskPolicyRule
	isSet bool
}

func (v NullableEntityRiskPolicyRule) Get() *EntityRiskPolicyRule {
	return v.value
}

func (v *NullableEntityRiskPolicyRule) Set(val *EntityRiskPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRule(val *EntityRiskPolicyRule) *NullableEntityRiskPolicyRule {
	return &NullableEntityRiskPolicyRule{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
