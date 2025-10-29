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

// checks if the GroupCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupCondition{}

// GroupCondition Specifies a set of groups whose users are to be included or excluded
type GroupCondition struct {
	// Groups to be excluded
	Exclude []string `json:"exclude"`
	// Groups to be included
	Include              []string `json:"include"`
	AdditionalProperties map[string]interface{}
}

type _GroupCondition GroupCondition

// NewGroupCondition instantiates a new GroupCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCondition(exclude []string, include []string) *GroupCondition {
	this := GroupCondition{}
	this.Exclude = exclude
	this.Include = include
	return &this
}

// NewGroupConditionWithDefaults instantiates a new GroupCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupConditionWithDefaults() *GroupCondition {
	this := GroupCondition{}
	return &this
}

// GetExclude returns the Exclude field value
func (o *GroupCondition) GetExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value
// and a boolean to check if the value has been set.
func (o *GroupCondition) GetExcludeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exclude, true
}

// SetExclude sets field value
func (o *GroupCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value
func (o *GroupCondition) GetInclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *GroupCondition) GetIncludeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *GroupCondition) SetInclude(v []string) {
	o.Include = v
}

func (o GroupCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exclude"] = o.Exclude
	toSerialize["include"] = o.Include

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exclude",
		"include",
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

	varGroupCondition := _GroupCondition{}

	err = json.Unmarshal(data, &varGroupCondition)

	if err != nil {
		return err
	}

	*o = GroupCondition(varGroupCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupCondition struct {
	value *GroupCondition
	isSet bool
}

func (v NullableGroupCondition) Get() *GroupCondition {
	return v.value
}

func (v *NullableGroupCondition) Set(val *GroupCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCondition(val *GroupCondition) *NullableGroupCondition {
	return &NullableGroupCondition{value: val, isSet: true}
}

func (v NullableGroupCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
