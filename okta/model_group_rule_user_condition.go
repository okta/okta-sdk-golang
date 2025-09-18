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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GroupRuleUserCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRuleUserCondition{}

// GroupRuleUserCondition Defines conditions specific to user exclusion
type GroupRuleUserCondition struct {
	// Excluded `userIds` when processing rules
	Exclude              []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRuleUserCondition GroupRuleUserCondition

// NewGroupRuleUserCondition instantiates a new GroupRuleUserCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleUserCondition() *GroupRuleUserCondition {
	this := GroupRuleUserCondition{}
	return &this
}

// NewGroupRuleUserConditionWithDefaults instantiates a new GroupRuleUserCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleUserConditionWithDefaults() *GroupRuleUserCondition {
	this := GroupRuleUserCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *GroupRuleUserCondition) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleUserCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *GroupRuleUserCondition) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *GroupRuleUserCondition) SetExclude(v []string) {
	o.Exclude = v
}

func (o GroupRuleUserCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRuleUserCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupRuleUserCondition) UnmarshalJSON(data []byte) (err error) {
	varGroupRuleUserCondition := _GroupRuleUserCondition{}

	err = json.Unmarshal(data, &varGroupRuleUserCondition)

	if err != nil {
		return err
	}

	*o = GroupRuleUserCondition(varGroupRuleUserCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupRuleUserCondition struct {
	value *GroupRuleUserCondition
	isSet bool
}

func (v NullableGroupRuleUserCondition) Get() *GroupRuleUserCondition {
	return v.value
}

func (v *NullableGroupRuleUserCondition) Set(val *GroupRuleUserCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleUserCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleUserCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleUserCondition(val *GroupRuleUserCondition) *NullableGroupRuleUserCondition {
	return &NullableGroupRuleUserCondition{value: val, isSet: true}
}

func (v NullableGroupRuleUserCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleUserCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
