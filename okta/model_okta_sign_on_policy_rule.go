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
	"reflect"
	"strings"
)

// OktaSignOnPolicyRule struct for OktaSignOnPolicyRule
type OktaSignOnPolicyRule struct {
	PolicyRule
	Actions *OktaSignOnPolicyRuleActions `json:"actions,omitempty"`
	Conditions *OktaSignOnPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyRule OktaSignOnPolicyRule

// NewOktaSignOnPolicyRule instantiates a new OktaSignOnPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyRule() *OktaSignOnPolicyRule {
	this := OktaSignOnPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewOktaSignOnPolicyRuleWithDefaults instantiates a new OktaSignOnPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyRuleWithDefaults() *OktaSignOnPolicyRule {
	this := OktaSignOnPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRule) GetActions() OktaSignOnPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret OktaSignOnPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRule) GetActionsOk() (*OktaSignOnPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given OktaSignOnPolicyRuleActions and assigns it to the Actions field.
func (o *OktaSignOnPolicyRule) SetActions(v OktaSignOnPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRule) GetConditions() OktaSignOnPolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret OktaSignOnPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRule) GetConditionsOk() (*OktaSignOnPolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given OktaSignOnPolicyRuleConditions and assigns it to the Conditions field.
func (o *OktaSignOnPolicyRule) SetConditions(v OktaSignOnPolicyRuleConditions) {
	o.Conditions = &v
}

func (o OktaSignOnPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyRule, errPolicyRule := json.Marshal(o.PolicyRule)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	errPolicyRule = json.Unmarshal([]byte(serializedPolicyRule), &toSerialize)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type OktaSignOnPolicyRuleWithoutEmbeddedStruct struct {
		Actions *OktaSignOnPolicyRuleActions `json:"actions,omitempty"`
		Conditions *OktaSignOnPolicyRuleConditions `json:"conditions,omitempty"`
	}

	varOktaSignOnPolicyRuleWithoutEmbeddedStruct := OktaSignOnPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varOktaSignOnPolicyRule := _OktaSignOnPolicyRule{}
		varOktaSignOnPolicyRule.Actions = varOktaSignOnPolicyRuleWithoutEmbeddedStruct.Actions
		varOktaSignOnPolicyRule.Conditions = varOktaSignOnPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = OktaSignOnPolicyRule(varOktaSignOnPolicyRule)
	} else {
		return err
	}

	varOktaSignOnPolicyRule := _OktaSignOnPolicyRule{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyRule)
	if err == nil {
		o.PolicyRule = varOktaSignOnPolicyRule.PolicyRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyRule struct {
	value *OktaSignOnPolicyRule
	isSet bool
}

func (v NullableOktaSignOnPolicyRule) Get() *OktaSignOnPolicyRule {
	return v.value
}

func (v *NullableOktaSignOnPolicyRule) Set(val *OktaSignOnPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyRule(val *OktaSignOnPolicyRule) *NullableOktaSignOnPolicyRule {
	return &NullableOktaSignOnPolicyRule{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

