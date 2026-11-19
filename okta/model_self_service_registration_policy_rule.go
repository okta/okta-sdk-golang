/*
Okta Admin Management API

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

// checks if the SelfServiceRegistrationPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceRegistrationPolicyRule{}

// SelfServiceRegistrationPolicyRule struct for SelfServiceRegistrationPolicyRule
type SelfServiceRegistrationPolicyRule struct {
	PolicyRule
	Actions              *SelfServiceRegistrationPolicyRuleActions    `json:"actions,omitempty"`
	Conditions           *SelfServiceRegistrationPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceRegistrationPolicyRule SelfServiceRegistrationPolicyRule

// NewSelfServiceRegistrationPolicyRule instantiates a new SelfServiceRegistrationPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceRegistrationPolicyRule() *SelfServiceRegistrationPolicyRule {
	this := SelfServiceRegistrationPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewSelfServiceRegistrationPolicyRuleWithDefaults instantiates a new SelfServiceRegistrationPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceRegistrationPolicyRuleWithDefaults() *SelfServiceRegistrationPolicyRule {
	this := SelfServiceRegistrationPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *SelfServiceRegistrationPolicyRule) GetActions() SelfServiceRegistrationPolicyRuleActions {
	if o == nil || IsNil(o.Actions) {
		var ret SelfServiceRegistrationPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationPolicyRule) GetActionsOk() (*SelfServiceRegistrationPolicyRuleActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *SelfServiceRegistrationPolicyRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given SelfServiceRegistrationPolicyRuleActions and assigns it to the Actions field.
func (o *SelfServiceRegistrationPolicyRule) SetActions(v SelfServiceRegistrationPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SelfServiceRegistrationPolicyRule) GetConditions() SelfServiceRegistrationPolicyRuleConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret SelfServiceRegistrationPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationPolicyRule) GetConditionsOk() (*SelfServiceRegistrationPolicyRuleConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SelfServiceRegistrationPolicyRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given SelfServiceRegistrationPolicyRuleConditions and assigns it to the Conditions field.
func (o *SelfServiceRegistrationPolicyRule) SetConditions(v SelfServiceRegistrationPolicyRuleConditions) {
	o.Conditions = &v
}

func (o SelfServiceRegistrationPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceRegistrationPolicyRule) ToMap() (map[string]interface{}, error) {
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

func (o *SelfServiceRegistrationPolicyRule) UnmarshalJSON(data []byte) (err error) {
	type SelfServiceRegistrationPolicyRuleWithoutEmbeddedStruct struct {
		Actions    *SelfServiceRegistrationPolicyRuleActions    `json:"actions,omitempty"`
		Conditions *SelfServiceRegistrationPolicyRuleConditions `json:"conditions,omitempty"`
	}

	varSelfServiceRegistrationPolicyRuleWithoutEmbeddedStruct := SelfServiceRegistrationPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varSelfServiceRegistrationPolicyRule := _SelfServiceRegistrationPolicyRule{}
		varSelfServiceRegistrationPolicyRule.Actions = varSelfServiceRegistrationPolicyRuleWithoutEmbeddedStruct.Actions
		varSelfServiceRegistrationPolicyRule.Conditions = varSelfServiceRegistrationPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = SelfServiceRegistrationPolicyRule(varSelfServiceRegistrationPolicyRule)
	} else {
		return err
	}

	varSelfServiceRegistrationPolicyRule := _SelfServiceRegistrationPolicyRule{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationPolicyRule)
	if err == nil {
		o.PolicyRule = varSelfServiceRegistrationPolicyRule.PolicyRule
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

type NullableSelfServiceRegistrationPolicyRule struct {
	value *SelfServiceRegistrationPolicyRule
	isSet bool
}

func (v NullableSelfServiceRegistrationPolicyRule) Get() *SelfServiceRegistrationPolicyRule {
	return v.value
}

func (v *NullableSelfServiceRegistrationPolicyRule) Set(val *SelfServiceRegistrationPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRegistrationPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRegistrationPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRegistrationPolicyRule(val *SelfServiceRegistrationPolicyRule) *NullableSelfServiceRegistrationPolicyRule {
	return &NullableSelfServiceRegistrationPolicyRule{value: val, isSet: true}
}

func (v NullableSelfServiceRegistrationPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRegistrationPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
