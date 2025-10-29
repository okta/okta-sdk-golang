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

// checks if the GroupRuleExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRuleExpression{}

// GroupRuleExpression Defines Okta specific [group-rules expression](https://developer.okta.com/docs/reference/okta-expression-language/#expressions-in-group-rules)
type GroupRuleExpression struct {
	// Expression type. Only valid value is '`urn:okta:expression:1.0`'.
	Type *string `json:"type,omitempty"`
	// Okta expression that would result in a Boolean value
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRuleExpression GroupRuleExpression

// NewGroupRuleExpression instantiates a new GroupRuleExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleExpression() *GroupRuleExpression {
	this := GroupRuleExpression{}
	return &this
}

// NewGroupRuleExpressionWithDefaults instantiates a new GroupRuleExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleExpressionWithDefaults() *GroupRuleExpression {
	this := GroupRuleExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupRuleExpression) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleExpression) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupRuleExpression) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupRuleExpression) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GroupRuleExpression) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleExpression) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GroupRuleExpression) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GroupRuleExpression) SetValue(v string) {
	o.Value = &v
}

func (o GroupRuleExpression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRuleExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupRuleExpression) UnmarshalJSON(data []byte) (err error) {
	varGroupRuleExpression := _GroupRuleExpression{}

	err = json.Unmarshal(data, &varGroupRuleExpression)

	if err != nil {
		return err
	}

	*o = GroupRuleExpression(varGroupRuleExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupRuleExpression struct {
	value *GroupRuleExpression
	isSet bool
}

func (v NullableGroupRuleExpression) Get() *GroupRuleExpression {
	return v.value
}

func (v *NullableGroupRuleExpression) Set(val *GroupRuleExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleExpression(val *GroupRuleExpression) *NullableGroupRuleExpression {
	return &NullableGroupRuleExpression{value: val, isSet: true}
}

func (v NullableGroupRuleExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
