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
)

// checks if the UserLifecycleAttributePolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLifecycleAttributePolicyRuleCondition{}

// UserLifecycleAttributePolicyRuleCondition struct for UserLifecycleAttributePolicyRuleCondition
type UserLifecycleAttributePolicyRuleCondition struct {
	AttributeName        *string `json:"attributeName,omitempty"`
	MatchingValue        *string `json:"matchingValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserLifecycleAttributePolicyRuleCondition UserLifecycleAttributePolicyRuleCondition

// NewUserLifecycleAttributePolicyRuleCondition instantiates a new UserLifecycleAttributePolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLifecycleAttributePolicyRuleCondition() *UserLifecycleAttributePolicyRuleCondition {
	this := UserLifecycleAttributePolicyRuleCondition{}
	return &this
}

// NewUserLifecycleAttributePolicyRuleConditionWithDefaults instantiates a new UserLifecycleAttributePolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLifecycleAttributePolicyRuleConditionWithDefaults() *UserLifecycleAttributePolicyRuleCondition {
	this := UserLifecycleAttributePolicyRuleCondition{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *UserLifecycleAttributePolicyRuleCondition) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLifecycleAttributePolicyRuleCondition) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *UserLifecycleAttributePolicyRuleCondition) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *UserLifecycleAttributePolicyRuleCondition) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetMatchingValue returns the MatchingValue field value if set, zero value otherwise.
func (o *UserLifecycleAttributePolicyRuleCondition) GetMatchingValue() string {
	if o == nil || IsNil(o.MatchingValue) {
		var ret string
		return ret
	}
	return *o.MatchingValue
}

// GetMatchingValueOk returns a tuple with the MatchingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLifecycleAttributePolicyRuleCondition) GetMatchingValueOk() (*string, bool) {
	if o == nil || IsNil(o.MatchingValue) {
		return nil, false
	}
	return o.MatchingValue, true
}

// HasMatchingValue returns a boolean if a field has been set.
func (o *UserLifecycleAttributePolicyRuleCondition) HasMatchingValue() bool {
	if o != nil && !IsNil(o.MatchingValue) {
		return true
	}

	return false
}

// SetMatchingValue gets a reference to the given string and assigns it to the MatchingValue field.
func (o *UserLifecycleAttributePolicyRuleCondition) SetMatchingValue(v string) {
	o.MatchingValue = &v
}

func (o UserLifecycleAttributePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLifecycleAttributePolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.MatchingValue) {
		toSerialize["matchingValue"] = o.MatchingValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserLifecycleAttributePolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varUserLifecycleAttributePolicyRuleCondition := _UserLifecycleAttributePolicyRuleCondition{}

	err = json.Unmarshal(data, &varUserLifecycleAttributePolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = UserLifecycleAttributePolicyRuleCondition(varUserLifecycleAttributePolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "matchingValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserLifecycleAttributePolicyRuleCondition struct {
	value *UserLifecycleAttributePolicyRuleCondition
	isSet bool
}

func (v NullableUserLifecycleAttributePolicyRuleCondition) Get() *UserLifecycleAttributePolicyRuleCondition {
	return v.value
}

func (v *NullableUserLifecycleAttributePolicyRuleCondition) Set(val *UserLifecycleAttributePolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLifecycleAttributePolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLifecycleAttributePolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLifecycleAttributePolicyRuleCondition(val *UserLifecycleAttributePolicyRuleCondition) *NullableUserLifecycleAttributePolicyRuleCondition {
	return &NullableUserLifecycleAttributePolicyRuleCondition{value: val, isSet: true}
}

func (v NullableUserLifecycleAttributePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLifecycleAttributePolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
