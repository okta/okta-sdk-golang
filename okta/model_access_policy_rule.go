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

// AccessPolicyRule struct for AccessPolicyRule
type AccessPolicyRule struct {
	PolicyRule
	Actions *AccessPolicyRuleActions `json:"actions,omitempty"`
	Conditions *AccessPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRule AccessPolicyRule

// NewAccessPolicyRule instantiates a new AccessPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRule() *AccessPolicyRule {
	this := AccessPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewAccessPolicyRuleWithDefaults instantiates a new AccessPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleWithDefaults() *AccessPolicyRule {
	this := AccessPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AccessPolicyRule) GetActions() AccessPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret AccessPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRule) GetActionsOk() (*AccessPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AccessPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given AccessPolicyRuleActions and assigns it to the Actions field.
func (o *AccessPolicyRule) SetActions(v AccessPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AccessPolicyRule) GetConditions() AccessPolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret AccessPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRule) GetConditionsOk() (*AccessPolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AccessPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AccessPolicyRuleConditions and assigns it to the Conditions field.
func (o *AccessPolicyRule) SetConditions(v AccessPolicyRuleConditions) {
	o.Conditions = &v
}

func (o AccessPolicyRule) MarshalJSON() ([]byte, error) {
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

func (o *AccessPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type AccessPolicyRuleWithoutEmbeddedStruct struct {
		Actions *AccessPolicyRuleActions `json:"actions,omitempty"`
		Conditions *AccessPolicyRuleConditions `json:"conditions,omitempty"`
	}

	varAccessPolicyRuleWithoutEmbeddedStruct := AccessPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAccessPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varAccessPolicyRule := _AccessPolicyRule{}
		varAccessPolicyRule.Actions = varAccessPolicyRuleWithoutEmbeddedStruct.Actions
		varAccessPolicyRule.Conditions = varAccessPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = AccessPolicyRule(varAccessPolicyRule)
	} else {
		return err
	}

	varAccessPolicyRule := _AccessPolicyRule{}

	err = json.Unmarshal(bytes, &varAccessPolicyRule)
	if err == nil {
		o.PolicyRule = varAccessPolicyRule.PolicyRule
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

type NullableAccessPolicyRule struct {
	value *AccessPolicyRule
	isSet bool
}

func (v NullableAccessPolicyRule) Get() *AccessPolicyRule {
	return v.value
}

func (v *NullableAccessPolicyRule) Set(val *AccessPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRule(val *AccessPolicyRule) *NullableAccessPolicyRule {
	return &NullableAccessPolicyRule{value: val, isSet: true}
}

func (v NullableAccessPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

