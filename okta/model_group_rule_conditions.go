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

// GroupRuleConditions struct for GroupRuleConditions
type GroupRuleConditions struct {
	Expression *GroupRuleExpression `json:"expression,omitempty"`
	People *GroupRulePeopleCondition `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRuleConditions GroupRuleConditions

// NewGroupRuleConditions instantiates a new GroupRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleConditions() *GroupRuleConditions {
	this := GroupRuleConditions{}
	return &this
}

// NewGroupRuleConditionsWithDefaults instantiates a new GroupRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleConditionsWithDefaults() *GroupRuleConditions {
	this := GroupRuleConditions{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *GroupRuleConditions) GetExpression() GroupRuleExpression {
	if o == nil || o.Expression == nil {
		var ret GroupRuleExpression
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleConditions) GetExpressionOk() (*GroupRuleExpression, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *GroupRuleConditions) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given GroupRuleExpression and assigns it to the Expression field.
func (o *GroupRuleConditions) SetExpression(v GroupRuleExpression) {
	o.Expression = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *GroupRuleConditions) GetPeople() GroupRulePeopleCondition {
	if o == nil || o.People == nil {
		var ret GroupRulePeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleConditions) GetPeopleOk() (*GroupRulePeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *GroupRuleConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given GroupRulePeopleCondition and assigns it to the People field.
func (o *GroupRuleConditions) SetPeople(v GroupRulePeopleCondition) {
	o.People = &v
}

func (o GroupRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupRuleConditions) UnmarshalJSON(bytes []byte) (err error) {
	varGroupRuleConditions := _GroupRuleConditions{}

	err = json.Unmarshal(bytes, &varGroupRuleConditions)
	if err == nil {
		*o = GroupRuleConditions(varGroupRuleConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupRuleConditions struct {
	value *GroupRuleConditions
	isSet bool
}

func (v NullableGroupRuleConditions) Get() *GroupRuleConditions {
	return v.value
}

func (v *NullableGroupRuleConditions) Set(val *GroupRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleConditions(val *GroupRuleConditions) *NullableGroupRuleConditions {
	return &NullableGroupRuleConditions{value: val, isSet: true}
}

func (v NullableGroupRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

