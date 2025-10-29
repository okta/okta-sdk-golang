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
	"reflect"
	"strings"
)

// checks if the IdpDiscoveryPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpDiscoveryPolicyRule{}

// IdpDiscoveryPolicyRule struct for IdpDiscoveryPolicyRule
type IdpDiscoveryPolicyRule struct {
	PolicyRule
	Actions              *IdpPolicyRuleAction             `json:"actions,omitempty"`
	Conditions           *IdpDiscoveryPolicyRuleCondition `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpDiscoveryPolicyRule IdpDiscoveryPolicyRule

// NewIdpDiscoveryPolicyRule instantiates a new IdpDiscoveryPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpDiscoveryPolicyRule() *IdpDiscoveryPolicyRule {
	this := IdpDiscoveryPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewIdpDiscoveryPolicyRuleWithDefaults instantiates a new IdpDiscoveryPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpDiscoveryPolicyRuleWithDefaults() *IdpDiscoveryPolicyRule {
	this := IdpDiscoveryPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRule) GetActions() IdpPolicyRuleAction {
	if o == nil || IsNil(o.Actions) {
		var ret IdpPolicyRuleAction
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRule) GetActionsOk() (*IdpPolicyRuleAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given IdpPolicyRuleAction and assigns it to the Actions field.
func (o *IdpDiscoveryPolicyRule) SetActions(v IdpPolicyRuleAction) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRule) GetConditions() IdpDiscoveryPolicyRuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret IdpDiscoveryPolicyRuleCondition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRule) GetConditionsOk() (*IdpDiscoveryPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given IdpDiscoveryPolicyRuleCondition and assigns it to the Conditions field.
func (o *IdpDiscoveryPolicyRule) SetConditions(v IdpDiscoveryPolicyRuleCondition) {
	o.Conditions = &v
}

func (o IdpDiscoveryPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpDiscoveryPolicyRule) ToMap() (map[string]interface{}, error) {
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

func (o *IdpDiscoveryPolicyRule) UnmarshalJSON(data []byte) (err error) {
	type IdpDiscoveryPolicyRuleWithoutEmbeddedStruct struct {
		Actions    *IdpPolicyRuleAction             `json:"actions,omitempty"`
		Conditions *IdpDiscoveryPolicyRuleCondition `json:"conditions,omitempty"`
	}

	varIdpDiscoveryPolicyRuleWithoutEmbeddedStruct := IdpDiscoveryPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIdpDiscoveryPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varIdpDiscoveryPolicyRule := _IdpDiscoveryPolicyRule{}
		varIdpDiscoveryPolicyRule.Actions = varIdpDiscoveryPolicyRuleWithoutEmbeddedStruct.Actions
		varIdpDiscoveryPolicyRule.Conditions = varIdpDiscoveryPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = IdpDiscoveryPolicyRule(varIdpDiscoveryPolicyRule)
	} else {
		return err
	}

	varIdpDiscoveryPolicyRule := _IdpDiscoveryPolicyRule{}

	err = json.Unmarshal(data, &varIdpDiscoveryPolicyRule)
	if err == nil {
		o.PolicyRule = varIdpDiscoveryPolicyRule.PolicyRule
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

type NullableIdpDiscoveryPolicyRule struct {
	value *IdpDiscoveryPolicyRule
	isSet bool
}

func (v NullableIdpDiscoveryPolicyRule) Get() *IdpDiscoveryPolicyRule {
	return v.value
}

func (v *NullableIdpDiscoveryPolicyRule) Set(val *IdpDiscoveryPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpDiscoveryPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpDiscoveryPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpDiscoveryPolicyRule(val *IdpDiscoveryPolicyRule) *NullableIdpDiscoveryPolicyRule {
	return &NullableIdpDiscoveryPolicyRule{value: val, isSet: true}
}

func (v NullableIdpDiscoveryPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpDiscoveryPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
