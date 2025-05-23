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
)

// ApplicationLayoutRule struct for ApplicationLayoutRule
type ApplicationLayoutRule struct {
	Effect *string `json:"effect,omitempty"`
	Condition *ApplicationLayoutRuleCondition `json:"condition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationLayoutRule ApplicationLayoutRule

// NewApplicationLayoutRule instantiates a new ApplicationLayoutRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLayoutRule() *ApplicationLayoutRule {
	this := ApplicationLayoutRule{}
	return &this
}

// NewApplicationLayoutRuleWithDefaults instantiates a new ApplicationLayoutRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLayoutRuleWithDefaults() *ApplicationLayoutRule {
	this := ApplicationLayoutRule{}
	return &this
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *ApplicationLayoutRule) GetEffect() string {
	if o == nil || o.Effect == nil {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutRule) GetEffectOk() (*string, bool) {
	if o == nil || o.Effect == nil {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *ApplicationLayoutRule) HasEffect() bool {
	if o != nil && o.Effect != nil {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *ApplicationLayoutRule) SetEffect(v string) {
	o.Effect = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ApplicationLayoutRule) GetCondition() ApplicationLayoutRuleCondition {
	if o == nil || o.Condition == nil {
		var ret ApplicationLayoutRuleCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutRule) GetConditionOk() (*ApplicationLayoutRuleCondition, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ApplicationLayoutRule) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given ApplicationLayoutRuleCondition and assigns it to the Condition field.
func (o *ApplicationLayoutRule) SetCondition(v ApplicationLayoutRuleCondition) {
	o.Condition = &v
}

func (o ApplicationLayoutRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Effect != nil {
		toSerialize["effect"] = o.Effect
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationLayoutRule) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationLayoutRule := _ApplicationLayoutRule{}

	err = json.Unmarshal(bytes, &varApplicationLayoutRule)
	if err == nil {
		*o = ApplicationLayoutRule(varApplicationLayoutRule)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "effect")
		delete(additionalProperties, "condition")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationLayoutRule struct {
	value *ApplicationLayoutRule
	isSet bool
}

func (v NullableApplicationLayoutRule) Get() *ApplicationLayoutRule {
	return v.value
}

func (v *NullableApplicationLayoutRule) Set(val *ApplicationLayoutRule) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLayoutRule) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLayoutRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLayoutRule(val *ApplicationLayoutRule) *NullableApplicationLayoutRule {
	return &NullableApplicationLayoutRule{value: val, isSet: true}
}

func (v NullableApplicationLayoutRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLayoutRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

