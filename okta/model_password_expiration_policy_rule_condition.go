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

// PasswordExpirationPolicyRuleCondition struct for PasswordExpirationPolicyRuleCondition
type PasswordExpirationPolicyRuleCondition struct {
	Number *int32 `json:"number,omitempty"`
	Unit *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordExpirationPolicyRuleCondition PasswordExpirationPolicyRuleCondition

// NewPasswordExpirationPolicyRuleCondition instantiates a new PasswordExpirationPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordExpirationPolicyRuleCondition() *PasswordExpirationPolicyRuleCondition {
	this := PasswordExpirationPolicyRuleCondition{}
	return &this
}

// NewPasswordExpirationPolicyRuleConditionWithDefaults instantiates a new PasswordExpirationPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordExpirationPolicyRuleConditionWithDefaults() *PasswordExpirationPolicyRuleCondition {
	this := PasswordExpirationPolicyRuleCondition{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PasswordExpirationPolicyRuleCondition) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpirationPolicyRuleCondition) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PasswordExpirationPolicyRuleCondition) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *PasswordExpirationPolicyRuleCondition) SetNumber(v int32) {
	o.Number = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *PasswordExpirationPolicyRuleCondition) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpirationPolicyRuleCondition) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *PasswordExpirationPolicyRuleCondition) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *PasswordExpirationPolicyRuleCondition) SetUnit(v string) {
	o.Unit = &v
}

func (o PasswordExpirationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordExpirationPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordExpirationPolicyRuleCondition := _PasswordExpirationPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varPasswordExpirationPolicyRuleCondition)
	if err == nil {
		*o = PasswordExpirationPolicyRuleCondition(varPasswordExpirationPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "number")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordExpirationPolicyRuleCondition struct {
	value *PasswordExpirationPolicyRuleCondition
	isSet bool
}

func (v NullablePasswordExpirationPolicyRuleCondition) Get() *PasswordExpirationPolicyRuleCondition {
	return v.value
}

func (v *NullablePasswordExpirationPolicyRuleCondition) Set(val *PasswordExpirationPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordExpirationPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordExpirationPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordExpirationPolicyRuleCondition(val *PasswordExpirationPolicyRuleCondition) *NullablePasswordExpirationPolicyRuleCondition {
	return &NullablePasswordExpirationPolicyRuleCondition{value: val, isSet: true}
}

func (v NullablePasswordExpirationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordExpirationPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

