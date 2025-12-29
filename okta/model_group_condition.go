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

// checks if the GroupCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupCondition{}

// GroupCondition Specifies a set of groups whose users are to be included or excluded
type GroupCondition struct {
	// Groups to be excluded
	Exclude []string `json:"exclude,omitempty"`
	// Groups to be included
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupCondition GroupCondition

// NewGroupCondition instantiates a new GroupCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCondition() *GroupCondition {
	this := GroupCondition{}
	return &this
}

// NewGroupConditionWithDefaults instantiates a new GroupCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupConditionWithDefaults() *GroupCondition {
	this := GroupCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *GroupCondition) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *GroupCondition) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *GroupCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *GroupCondition) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *GroupCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
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
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupCondition) UnmarshalJSON(data []byte) (err error) {
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
