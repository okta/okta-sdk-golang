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

// checks if the GroupRuleGroupCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRuleGroupCondition{}

// GroupRuleGroupCondition Currently not supported
type GroupRuleGroupCondition struct {
	// Currently not supported
	Exclude              []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRuleGroupCondition GroupRuleGroupCondition

// NewGroupRuleGroupCondition instantiates a new GroupRuleGroupCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleGroupCondition() *GroupRuleGroupCondition {
	this := GroupRuleGroupCondition{}
	return &this
}

// NewGroupRuleGroupConditionWithDefaults instantiates a new GroupRuleGroupCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleGroupConditionWithDefaults() *GroupRuleGroupCondition {
	this := GroupRuleGroupCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *GroupRuleGroupCondition) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleGroupCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *GroupRuleGroupCondition) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *GroupRuleGroupCondition) SetExclude(v []string) {
	o.Exclude = v
}

func (o GroupRuleGroupCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRuleGroupCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupRuleGroupCondition) UnmarshalJSON(data []byte) (err error) {
	varGroupRuleGroupCondition := _GroupRuleGroupCondition{}

	err = json.Unmarshal(data, &varGroupRuleGroupCondition)

	if err != nil {
		return err
	}

	*o = GroupRuleGroupCondition(varGroupRuleGroupCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupRuleGroupCondition struct {
	value *GroupRuleGroupCondition
	isSet bool
}

func (v NullableGroupRuleGroupCondition) Get() *GroupRuleGroupCondition {
	return v.value
}

func (v *NullableGroupRuleGroupCondition) Set(val *GroupRuleGroupCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleGroupCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleGroupCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleGroupCondition(val *GroupRuleGroupCondition) *NullableGroupRuleGroupCondition {
	return &NullableGroupRuleGroupCondition{value: val, isSet: true}
}

func (v NullableGroupRuleGroupCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleGroupCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
