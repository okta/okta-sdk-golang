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
	"fmt"
)

// checks if the AccessPolicyRuleCustomCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicyRuleCustomCondition{}

// AccessPolicyRuleCustomCondition Specifies [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language-in-identity-engine/) expressions
type AccessPolicyRuleCustomCondition struct {
	// expression to match
	Condition            string `json:"condition"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleCustomCondition AccessPolicyRuleCustomCondition

// NewAccessPolicyRuleCustomCondition instantiates a new AccessPolicyRuleCustomCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleCustomCondition(condition string) *AccessPolicyRuleCustomCondition {
	this := AccessPolicyRuleCustomCondition{}
	this.Condition = condition
	return &this
}

// NewAccessPolicyRuleCustomConditionWithDefaults instantiates a new AccessPolicyRuleCustomCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleCustomConditionWithDefaults() *AccessPolicyRuleCustomCondition {
	this := AccessPolicyRuleCustomCondition{}
	return &this
}

// GetCondition returns the Condition field value
func (o *AccessPolicyRuleCustomCondition) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleCustomCondition) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *AccessPolicyRuleCustomCondition) SetCondition(v string) {
	o.Condition = v
}

func (o AccessPolicyRuleCustomCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicyRuleCustomCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition"] = o.Condition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicyRuleCustomCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"condition",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessPolicyRuleCustomCondition := _AccessPolicyRuleCustomCondition{}

	err = json.Unmarshal(data, &varAccessPolicyRuleCustomCondition)

	if err != nil {
		return err
	}

	*o = AccessPolicyRuleCustomCondition(varAccessPolicyRuleCustomCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "condition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicyRuleCustomCondition struct {
	value *AccessPolicyRuleCustomCondition
	isSet bool
}

func (v NullableAccessPolicyRuleCustomCondition) Get() *AccessPolicyRuleCustomCondition {
	return v.value
}

func (v *NullableAccessPolicyRuleCustomCondition) Set(val *AccessPolicyRuleCustomCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleCustomCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleCustomCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleCustomCondition(val *AccessPolicyRuleCustomCondition) *NullableAccessPolicyRuleCustomCondition {
	return &NullableAccessPolicyRuleCustomCondition{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleCustomCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleCustomCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
