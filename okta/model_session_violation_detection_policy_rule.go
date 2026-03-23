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

// checks if the SessionViolationDetectionPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionViolationDetectionPolicyRule{}

// SessionViolationDetectionPolicyRule struct for SessionViolationDetectionPolicyRule
type SessionViolationDetectionPolicyRule struct {
	PolicyRule
	Actions              *SessionViolationDetectionPolicyRuleAllOfActions    `json:"actions,omitempty"`
	Conditions           *SessionViolationDetectionPolicyRuleAllOfConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionViolationDetectionPolicyRule SessionViolationDetectionPolicyRule

// NewSessionViolationDetectionPolicyRule instantiates a new SessionViolationDetectionPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionViolationDetectionPolicyRule() *SessionViolationDetectionPolicyRule {
	this := SessionViolationDetectionPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewSessionViolationDetectionPolicyRuleWithDefaults instantiates a new SessionViolationDetectionPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionViolationDetectionPolicyRuleWithDefaults() *SessionViolationDetectionPolicyRule {
	this := SessionViolationDetectionPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *SessionViolationDetectionPolicyRule) GetActions() SessionViolationDetectionPolicyRuleAllOfActions {
	if o == nil || IsNil(o.Actions) {
		var ret SessionViolationDetectionPolicyRuleAllOfActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionPolicyRule) GetActionsOk() (*SessionViolationDetectionPolicyRuleAllOfActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *SessionViolationDetectionPolicyRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given SessionViolationDetectionPolicyRuleAllOfActions and assigns it to the Actions field.
func (o *SessionViolationDetectionPolicyRule) SetActions(v SessionViolationDetectionPolicyRuleAllOfActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SessionViolationDetectionPolicyRule) GetConditions() SessionViolationDetectionPolicyRuleAllOfConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret SessionViolationDetectionPolicyRuleAllOfConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionPolicyRule) GetConditionsOk() (*SessionViolationDetectionPolicyRuleAllOfConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SessionViolationDetectionPolicyRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given SessionViolationDetectionPolicyRuleAllOfConditions and assigns it to the Conditions field.
func (o *SessionViolationDetectionPolicyRule) SetConditions(v SessionViolationDetectionPolicyRuleAllOfConditions) {
	o.Conditions = &v
}

func (o SessionViolationDetectionPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionViolationDetectionPolicyRule) ToMap() (map[string]interface{}, error) {
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

func (o *SessionViolationDetectionPolicyRule) UnmarshalJSON(data []byte) (err error) {
	type SessionViolationDetectionPolicyRuleWithoutEmbeddedStruct struct {
		Actions    *SessionViolationDetectionPolicyRuleAllOfActions    `json:"actions,omitempty"`
		Conditions *SessionViolationDetectionPolicyRuleAllOfConditions `json:"conditions,omitempty"`
	}

	varSessionViolationDetectionPolicyRuleWithoutEmbeddedStruct := SessionViolationDetectionPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSessionViolationDetectionPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varSessionViolationDetectionPolicyRule := _SessionViolationDetectionPolicyRule{}
		varSessionViolationDetectionPolicyRule.Actions = varSessionViolationDetectionPolicyRuleWithoutEmbeddedStruct.Actions
		varSessionViolationDetectionPolicyRule.Conditions = varSessionViolationDetectionPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = SessionViolationDetectionPolicyRule(varSessionViolationDetectionPolicyRule)
	} else {
		return err
	}

	varSessionViolationDetectionPolicyRule := _SessionViolationDetectionPolicyRule{}

	err = json.Unmarshal(data, &varSessionViolationDetectionPolicyRule)
	if err == nil {
		o.PolicyRule = varSessionViolationDetectionPolicyRule.PolicyRule
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

type NullableSessionViolationDetectionPolicyRule struct {
	value *SessionViolationDetectionPolicyRule
	isSet bool
}

func (v NullableSessionViolationDetectionPolicyRule) Get() *SessionViolationDetectionPolicyRule {
	return v.value
}

func (v *NullableSessionViolationDetectionPolicyRule) Set(val *SessionViolationDetectionPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionViolationDetectionPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionViolationDetectionPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionViolationDetectionPolicyRule(val *SessionViolationDetectionPolicyRule) *NullableSessionViolationDetectionPolicyRule {
	return &NullableSessionViolationDetectionPolicyRule{value: val, isSet: true}
}

func (v NullableSessionViolationDetectionPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionViolationDetectionPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
