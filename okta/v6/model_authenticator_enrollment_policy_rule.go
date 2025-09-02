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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AuthenticatorEnrollmentPolicyRule struct for AuthenticatorEnrollmentPolicyRule
type AuthenticatorEnrollmentPolicyRule struct {
	PolicyRule
	Actions *AuthenticatorEnrollmentPolicyRuleActions `json:"actions,omitempty"`
	Conditions *AuthenticatorEnrollmentPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyRule AuthenticatorEnrollmentPolicyRule

// NewAuthenticatorEnrollmentPolicyRule instantiates a new AuthenticatorEnrollmentPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyRule() *AuthenticatorEnrollmentPolicyRule {
	this := AuthenticatorEnrollmentPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewAuthenticatorEnrollmentPolicyRuleWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyRuleWithDefaults() *AuthenticatorEnrollmentPolicyRule {
	this := AuthenticatorEnrollmentPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRule) GetActions() AuthenticatorEnrollmentPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret AuthenticatorEnrollmentPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRule) GetActionsOk() (*AuthenticatorEnrollmentPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given AuthenticatorEnrollmentPolicyRuleActions and assigns it to the Actions field.
func (o *AuthenticatorEnrollmentPolicyRule) SetActions(v AuthenticatorEnrollmentPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRule) GetConditions() AuthenticatorEnrollmentPolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret AuthenticatorEnrollmentPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRule) GetConditionsOk() (*AuthenticatorEnrollmentPolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AuthenticatorEnrollmentPolicyRuleConditions and assigns it to the Conditions field.
func (o *AuthenticatorEnrollmentPolicyRule) SetConditions(v AuthenticatorEnrollmentPolicyRuleConditions) {
	o.Conditions = &v
}

func (o AuthenticatorEnrollmentPolicyRule) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorEnrollmentPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorEnrollmentPolicyRuleWithoutEmbeddedStruct struct {
		Actions *AuthenticatorEnrollmentPolicyRuleActions `json:"actions,omitempty"`
		Conditions *AuthenticatorEnrollmentPolicyRuleConditions `json:"conditions,omitempty"`
	}

	varAuthenticatorEnrollmentPolicyRuleWithoutEmbeddedStruct := AuthenticatorEnrollmentPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorEnrollmentPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorEnrollmentPolicyRule := _AuthenticatorEnrollmentPolicyRule{}
		varAuthenticatorEnrollmentPolicyRule.Actions = varAuthenticatorEnrollmentPolicyRuleWithoutEmbeddedStruct.Actions
		varAuthenticatorEnrollmentPolicyRule.Conditions = varAuthenticatorEnrollmentPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = AuthenticatorEnrollmentPolicyRule(varAuthenticatorEnrollmentPolicyRule)
	} else {
		return err
	}

	varAuthenticatorEnrollmentPolicyRule := _AuthenticatorEnrollmentPolicyRule{}

	err = json.Unmarshal(bytes, &varAuthenticatorEnrollmentPolicyRule)
	if err == nil {
		o.PolicyRule = varAuthenticatorEnrollmentPolicyRule.PolicyRule
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

type NullableAuthenticatorEnrollmentPolicyRule struct {
	value *AuthenticatorEnrollmentPolicyRule
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyRule) Get() *AuthenticatorEnrollmentPolicyRule {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyRule) Set(val *AuthenticatorEnrollmentPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyRule(val *AuthenticatorEnrollmentPolicyRule) *NullableAuthenticatorEnrollmentPolicyRule {
	return &NullableAuthenticatorEnrollmentPolicyRule{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

