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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the LifecycleExpirationPolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleExpirationPolicyRuleCondition{}

// LifecycleExpirationPolicyRuleCondition struct for LifecycleExpirationPolicyRuleCondition
type LifecycleExpirationPolicyRuleCondition struct {
	LifecycleStatus      *string `json:"lifecycleStatus,omitempty"`
	Number               *int32  `json:"number,omitempty"`
	Unit                 *string `json:"unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleExpirationPolicyRuleCondition LifecycleExpirationPolicyRuleCondition

// NewLifecycleExpirationPolicyRuleCondition instantiates a new LifecycleExpirationPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleExpirationPolicyRuleCondition() *LifecycleExpirationPolicyRuleCondition {
	this := LifecycleExpirationPolicyRuleCondition{}
	return &this
}

// NewLifecycleExpirationPolicyRuleConditionWithDefaults instantiates a new LifecycleExpirationPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleExpirationPolicyRuleConditionWithDefaults() *LifecycleExpirationPolicyRuleCondition {
	this := LifecycleExpirationPolicyRuleCondition{}
	return &this
}

// GetLifecycleStatus returns the LifecycleStatus field value if set, zero value otherwise.
func (o *LifecycleExpirationPolicyRuleCondition) GetLifecycleStatus() string {
	if o == nil || IsNil(o.LifecycleStatus) {
		var ret string
		return ret
	}
	return *o.LifecycleStatus
}

// GetLifecycleStatusOk returns a tuple with the LifecycleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleExpirationPolicyRuleCondition) GetLifecycleStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LifecycleStatus) {
		return nil, false
	}
	return o.LifecycleStatus, true
}

// HasLifecycleStatus returns a boolean if a field has been set.
func (o *LifecycleExpirationPolicyRuleCondition) HasLifecycleStatus() bool {
	if o != nil && !IsNil(o.LifecycleStatus) {
		return true
	}

	return false
}

// SetLifecycleStatus gets a reference to the given string and assigns it to the LifecycleStatus field.
func (o *LifecycleExpirationPolicyRuleCondition) SetLifecycleStatus(v string) {
	o.LifecycleStatus = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *LifecycleExpirationPolicyRuleCondition) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleExpirationPolicyRuleCondition) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *LifecycleExpirationPolicyRuleCondition) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *LifecycleExpirationPolicyRuleCondition) SetNumber(v int32) {
	o.Number = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *LifecycleExpirationPolicyRuleCondition) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleExpirationPolicyRuleCondition) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *LifecycleExpirationPolicyRuleCondition) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *LifecycleExpirationPolicyRuleCondition) SetUnit(v string) {
	o.Unit = &v
}

func (o LifecycleExpirationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleExpirationPolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LifecycleStatus) {
		toSerialize["lifecycleStatus"] = o.LifecycleStatus
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleExpirationPolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varLifecycleExpirationPolicyRuleCondition := _LifecycleExpirationPolicyRuleCondition{}

	err = json.Unmarshal(data, &varLifecycleExpirationPolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = LifecycleExpirationPolicyRuleCondition(varLifecycleExpirationPolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lifecycleStatus")
		delete(additionalProperties, "number")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleExpirationPolicyRuleCondition struct {
	value *LifecycleExpirationPolicyRuleCondition
	isSet bool
}

func (v NullableLifecycleExpirationPolicyRuleCondition) Get() *LifecycleExpirationPolicyRuleCondition {
	return v.value
}

func (v *NullableLifecycleExpirationPolicyRuleCondition) Set(val *LifecycleExpirationPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleExpirationPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleExpirationPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleExpirationPolicyRuleCondition(val *LifecycleExpirationPolicyRuleCondition) *NullableLifecycleExpirationPolicyRuleCondition {
	return &NullableLifecycleExpirationPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableLifecycleExpirationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleExpirationPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
