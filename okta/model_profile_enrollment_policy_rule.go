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

// ProfileEnrollmentPolicyRule struct for ProfileEnrollmentPolicyRule
type ProfileEnrollmentPolicyRule struct {
	PolicyRule
	Actions *ProfileEnrollmentPolicyRuleActions `json:"actions,omitempty"`
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileEnrollmentPolicyRule ProfileEnrollmentPolicyRule

// NewProfileEnrollmentPolicyRule instantiates a new ProfileEnrollmentPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentPolicyRule() *ProfileEnrollmentPolicyRule {
	this := ProfileEnrollmentPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewProfileEnrollmentPolicyRuleWithDefaults instantiates a new ProfileEnrollmentPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentPolicyRuleWithDefaults() *ProfileEnrollmentPolicyRule {
	this := ProfileEnrollmentPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRule) GetActions() ProfileEnrollmentPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret ProfileEnrollmentPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRule) GetActionsOk() (*ProfileEnrollmentPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given ProfileEnrollmentPolicyRuleActions and assigns it to the Actions field.
func (o *ProfileEnrollmentPolicyRule) SetActions(v ProfileEnrollmentPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRule) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRule) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *ProfileEnrollmentPolicyRule) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

func (o ProfileEnrollmentPolicyRule) MarshalJSON() ([]byte, error) {
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

func (o *ProfileEnrollmentPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type ProfileEnrollmentPolicyRuleWithoutEmbeddedStruct struct {
		Actions *ProfileEnrollmentPolicyRuleActions `json:"actions,omitempty"`
		Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	}

	varProfileEnrollmentPolicyRuleWithoutEmbeddedStruct := ProfileEnrollmentPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varProfileEnrollmentPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varProfileEnrollmentPolicyRule := _ProfileEnrollmentPolicyRule{}
		varProfileEnrollmentPolicyRule.Actions = varProfileEnrollmentPolicyRuleWithoutEmbeddedStruct.Actions
		varProfileEnrollmentPolicyRule.Conditions = varProfileEnrollmentPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = ProfileEnrollmentPolicyRule(varProfileEnrollmentPolicyRule)
	} else {
		return err
	}

	varProfileEnrollmentPolicyRule := _ProfileEnrollmentPolicyRule{}

	err = json.Unmarshal(bytes, &varProfileEnrollmentPolicyRule)
	if err == nil {
		o.PolicyRule = varProfileEnrollmentPolicyRule.PolicyRule
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

type NullableProfileEnrollmentPolicyRule struct {
	value *ProfileEnrollmentPolicyRule
	isSet bool
}

func (v NullableProfileEnrollmentPolicyRule) Get() *ProfileEnrollmentPolicyRule {
	return v.value
}

func (v *NullableProfileEnrollmentPolicyRule) Set(val *ProfileEnrollmentPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentPolicyRule(val *ProfileEnrollmentPolicyRule) *NullableProfileEnrollmentPolicyRule {
	return &NullableProfileEnrollmentPolicyRule{value: val, isSet: true}
}

func (v NullableProfileEnrollmentPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

