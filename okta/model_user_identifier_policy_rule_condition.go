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

// UserIdentifierPolicyRuleCondition struct for UserIdentifierPolicyRuleCondition
type UserIdentifierPolicyRuleCondition struct {
	Attribute *string `json:"attribute,omitempty"`
	Patterns []UserIdentifierConditionEvaluatorPattern `json:"patterns,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentifierPolicyRuleCondition UserIdentifierPolicyRuleCondition

// NewUserIdentifierPolicyRuleCondition instantiates a new UserIdentifierPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentifierPolicyRuleCondition() *UserIdentifierPolicyRuleCondition {
	this := UserIdentifierPolicyRuleCondition{}
	return &this
}

// NewUserIdentifierPolicyRuleConditionWithDefaults instantiates a new UserIdentifierPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentifierPolicyRuleConditionWithDefaults() *UserIdentifierPolicyRuleCondition {
	this := UserIdentifierPolicyRuleCondition{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *UserIdentifierPolicyRuleCondition) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentifierPolicyRuleCondition) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *UserIdentifierPolicyRuleCondition) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *UserIdentifierPolicyRuleCondition) SetAttribute(v string) {
	o.Attribute = &v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
func (o *UserIdentifierPolicyRuleCondition) GetPatterns() []UserIdentifierConditionEvaluatorPattern {
	if o == nil || o.Patterns == nil {
		var ret []UserIdentifierConditionEvaluatorPattern
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentifierPolicyRuleCondition) GetPatternsOk() ([]UserIdentifierConditionEvaluatorPattern, bool) {
	if o == nil || o.Patterns == nil {
		return nil, false
	}
	return o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *UserIdentifierPolicyRuleCondition) HasPatterns() bool {
	if o != nil && o.Patterns != nil {
		return true
	}

	return false
}

// SetPatterns gets a reference to the given []UserIdentifierConditionEvaluatorPattern and assigns it to the Patterns field.
func (o *UserIdentifierPolicyRuleCondition) SetPatterns(v []UserIdentifierConditionEvaluatorPattern) {
	o.Patterns = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserIdentifierPolicyRuleCondition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentifierPolicyRuleCondition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserIdentifierPolicyRuleCondition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserIdentifierPolicyRuleCondition) SetType(v string) {
	o.Type = &v
}

func (o UserIdentifierPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Patterns != nil {
		toSerialize["patterns"] = o.Patterns
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserIdentifierPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varUserIdentifierPolicyRuleCondition := _UserIdentifierPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varUserIdentifierPolicyRuleCondition)
	if err == nil {
		*o = UserIdentifierPolicyRuleCondition(varUserIdentifierPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "patterns")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserIdentifierPolicyRuleCondition struct {
	value *UserIdentifierPolicyRuleCondition
	isSet bool
}

func (v NullableUserIdentifierPolicyRuleCondition) Get() *UserIdentifierPolicyRuleCondition {
	return v.value
}

func (v *NullableUserIdentifierPolicyRuleCondition) Set(val *UserIdentifierPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentifierPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentifierPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentifierPolicyRuleCondition(val *UserIdentifierPolicyRuleCondition) *NullableUserIdentifierPolicyRuleCondition {
	return &NullableUserIdentifierPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableUserIdentifierPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentifierPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

