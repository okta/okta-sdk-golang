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

// checks if the ClientUpdatePolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientUpdatePolicyRule{}

// ClientUpdatePolicyRule struct for ClientUpdatePolicyRule
type ClientUpdatePolicyRule struct {
	PolicyRule
	Actions *ClientUpdatePolicyRuleAllOfActions `json:"actions,omitempty"`
	// Rule conditions aren't supported for this policy type.
	Conditions           NullableString `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientUpdatePolicyRule ClientUpdatePolicyRule

// NewClientUpdatePolicyRule instantiates a new ClientUpdatePolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientUpdatePolicyRule() *ClientUpdatePolicyRule {
	this := ClientUpdatePolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewClientUpdatePolicyRuleWithDefaults instantiates a new ClientUpdatePolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientUpdatePolicyRuleWithDefaults() *ClientUpdatePolicyRule {
	this := ClientUpdatePolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ClientUpdatePolicyRule) GetActions() ClientUpdatePolicyRuleAllOfActions {
	if o == nil || IsNil(o.Actions) {
		var ret ClientUpdatePolicyRuleAllOfActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUpdatePolicyRule) GetActionsOk() (*ClientUpdatePolicyRuleAllOfActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ClientUpdatePolicyRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given ClientUpdatePolicyRuleAllOfActions and assigns it to the Actions field.
func (o *ClientUpdatePolicyRule) SetActions(v ClientUpdatePolicyRuleAllOfActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientUpdatePolicyRule) GetConditions() string {
	if o == nil || IsNil(o.Conditions.Get()) {
		var ret string
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientUpdatePolicyRule) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *ClientUpdatePolicyRule) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableString and assigns it to the Conditions field.
func (o *ClientUpdatePolicyRule) SetConditions(v string) {
	o.Conditions.Set(&v)
}

// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *ClientUpdatePolicyRule) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *ClientUpdatePolicyRule) UnsetConditions() {
	o.Conditions.Unset()
}

func (o ClientUpdatePolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientUpdatePolicyRule) ToMap() (map[string]interface{}, error) {
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
	if o.Conditions.IsSet() {
		toSerialize["conditions"] = o.Conditions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientUpdatePolicyRule) UnmarshalJSON(data []byte) (err error) {
	type ClientUpdatePolicyRuleWithoutEmbeddedStruct struct {
		Actions *ClientUpdatePolicyRuleAllOfActions `json:"actions,omitempty"`
		// Rule conditions aren't supported for this policy type.
		Conditions NullableString `json:"conditions,omitempty"`
	}

	varClientUpdatePolicyRuleWithoutEmbeddedStruct := ClientUpdatePolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varClientUpdatePolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varClientUpdatePolicyRule := _ClientUpdatePolicyRule{}
		varClientUpdatePolicyRule.Actions = varClientUpdatePolicyRuleWithoutEmbeddedStruct.Actions
		varClientUpdatePolicyRule.Conditions = varClientUpdatePolicyRuleWithoutEmbeddedStruct.Conditions
		*o = ClientUpdatePolicyRule(varClientUpdatePolicyRule)
	} else {
		return err
	}

	varClientUpdatePolicyRule := _ClientUpdatePolicyRule{}

	err = json.Unmarshal(data, &varClientUpdatePolicyRule)
	if err == nil {
		o.PolicyRule = varClientUpdatePolicyRule.PolicyRule
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

type NullableClientUpdatePolicyRule struct {
	value *ClientUpdatePolicyRule
	isSet bool
}

func (v NullableClientUpdatePolicyRule) Get() *ClientUpdatePolicyRule {
	return v.value
}

func (v *NullableClientUpdatePolicyRule) Set(val *ClientUpdatePolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableClientUpdatePolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableClientUpdatePolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientUpdatePolicyRule(val *ClientUpdatePolicyRule) *NullableClientUpdatePolicyRule {
	return &NullableClientUpdatePolicyRule{value: val, isSet: true}
}

func (v NullableClientUpdatePolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientUpdatePolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
