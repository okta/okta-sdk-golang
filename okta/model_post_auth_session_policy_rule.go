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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the PostAuthSessionPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAuthSessionPolicyRule{}

// PostAuthSessionPolicyRule struct for PostAuthSessionPolicyRule
type PostAuthSessionPolicyRule struct {
	PolicyRule
	Actions              *PostAuthSessionPolicyRuleAllOfActions    `json:"actions,omitempty"`
	Conditions           *PostAuthSessionPolicyRuleAllOfConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionPolicyRule PostAuthSessionPolicyRule

// NewPostAuthSessionPolicyRule instantiates a new PostAuthSessionPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionPolicyRule() *PostAuthSessionPolicyRule {
	this := PostAuthSessionPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewPostAuthSessionPolicyRuleWithDefaults instantiates a new PostAuthSessionPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionPolicyRuleWithDefaults() *PostAuthSessionPolicyRule {
	this := PostAuthSessionPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRule) GetActions() PostAuthSessionPolicyRuleAllOfActions {
	if o == nil || IsNil(o.Actions) {
		var ret PostAuthSessionPolicyRuleAllOfActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRule) GetActionsOk() (*PostAuthSessionPolicyRuleAllOfActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given PostAuthSessionPolicyRuleAllOfActions and assigns it to the Actions field.
func (o *PostAuthSessionPolicyRule) SetActions(v PostAuthSessionPolicyRuleAllOfActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRule) GetConditions() PostAuthSessionPolicyRuleAllOfConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret PostAuthSessionPolicyRuleAllOfConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRule) GetConditionsOk() (*PostAuthSessionPolicyRuleAllOfConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PostAuthSessionPolicyRuleAllOfConditions and assigns it to the Conditions field.
func (o *PostAuthSessionPolicyRule) SetConditions(v PostAuthSessionPolicyRuleAllOfConditions) {
	o.Conditions = &v
}

func (o PostAuthSessionPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthSessionPolicyRule) ToMap() (map[string]interface{}, error) {
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

func (o *PostAuthSessionPolicyRule) UnmarshalJSON(data []byte) (err error) {
	type PostAuthSessionPolicyRuleWithoutEmbeddedStruct struct {
		Actions    *PostAuthSessionPolicyRuleAllOfActions    `json:"actions,omitempty"`
		Conditions *PostAuthSessionPolicyRuleAllOfConditions `json:"conditions,omitempty"`
	}

	varPostAuthSessionPolicyRuleWithoutEmbeddedStruct := PostAuthSessionPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPostAuthSessionPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varPostAuthSessionPolicyRule := _PostAuthSessionPolicyRule{}
		varPostAuthSessionPolicyRule.Actions = varPostAuthSessionPolicyRuleWithoutEmbeddedStruct.Actions
		varPostAuthSessionPolicyRule.Conditions = varPostAuthSessionPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = PostAuthSessionPolicyRule(varPostAuthSessionPolicyRule)
	} else {
		return err
	}

	varPostAuthSessionPolicyRule := _PostAuthSessionPolicyRule{}

	err = json.Unmarshal(data, &varPostAuthSessionPolicyRule)
	if err == nil {
		o.PolicyRule = varPostAuthSessionPolicyRule.PolicyRule
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

type NullablePostAuthSessionPolicyRule struct {
	value *PostAuthSessionPolicyRule
	isSet bool
}

func (v NullablePostAuthSessionPolicyRule) Get() *PostAuthSessionPolicyRule {
	return v.value
}

func (v *NullablePostAuthSessionPolicyRule) Set(val *PostAuthSessionPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionPolicyRule(val *PostAuthSessionPolicyRule) *NullablePostAuthSessionPolicyRule {
	return &NullablePostAuthSessionPolicyRule{value: val, isSet: true}
}

func (v NullablePostAuthSessionPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
