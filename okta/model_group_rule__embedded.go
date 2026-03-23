/*
Okta Admin Management API

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

// checks if the GroupRuleEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRuleEmbedded{}

// GroupRuleEmbedded This object appears with embedded resources related to the group rule if you use the `expand` query parameter
type GroupRuleEmbedded struct {
	// A mapping of group IDs to group names
	GroupIdToGroupNameMap *map[string]string `json:"groupIdToGroupNameMap,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GroupRuleEmbedded GroupRuleEmbedded

// NewGroupRuleEmbedded instantiates a new GroupRuleEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRuleEmbedded() *GroupRuleEmbedded {
	this := GroupRuleEmbedded{}
	return &this
}

// NewGroupRuleEmbeddedWithDefaults instantiates a new GroupRuleEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRuleEmbeddedWithDefaults() *GroupRuleEmbedded {
	this := GroupRuleEmbedded{}
	return &this
}

// GetGroupIdToGroupNameMap returns the GroupIdToGroupNameMap field value if set, zero value otherwise.
func (o *GroupRuleEmbedded) GetGroupIdToGroupNameMap() map[string]string {
	if o == nil || IsNil(o.GroupIdToGroupNameMap) {
		var ret map[string]string
		return ret
	}
	return *o.GroupIdToGroupNameMap
}

// GetGroupIdToGroupNameMapOk returns a tuple with the GroupIdToGroupNameMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRuleEmbedded) GetGroupIdToGroupNameMapOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.GroupIdToGroupNameMap) {
		return nil, false
	}
	return o.GroupIdToGroupNameMap, true
}

// HasGroupIdToGroupNameMap returns a boolean if a field has been set.
func (o *GroupRuleEmbedded) HasGroupIdToGroupNameMap() bool {
	if o != nil && !IsNil(o.GroupIdToGroupNameMap) {
		return true
	}

	return false
}

// SetGroupIdToGroupNameMap gets a reference to the given map[string]string and assigns it to the GroupIdToGroupNameMap field.
func (o *GroupRuleEmbedded) SetGroupIdToGroupNameMap(v map[string]string) {
	o.GroupIdToGroupNameMap = &v
}

func (o GroupRuleEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRuleEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupIdToGroupNameMap) {
		toSerialize["groupIdToGroupNameMap"] = o.GroupIdToGroupNameMap
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupRuleEmbedded) UnmarshalJSON(data []byte) (err error) {
	varGroupRuleEmbedded := _GroupRuleEmbedded{}

	err = json.Unmarshal(data, &varGroupRuleEmbedded)

	if err != nil {
		return err
	}

	*o = GroupRuleEmbedded(varGroupRuleEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groupIdToGroupNameMap")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupRuleEmbedded struct {
	value *GroupRuleEmbedded
	isSet bool
}

func (v NullableGroupRuleEmbedded) Get() *GroupRuleEmbedded {
	return v.value
}

func (v *NullableGroupRuleEmbedded) Set(val *GroupRuleEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRuleEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRuleEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRuleEmbedded(val *GroupRuleEmbedded) *NullableGroupRuleEmbedded {
	return &NullableGroupRuleEmbedded{value: val, isSet: true}
}

func (v NullableGroupRuleEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRuleEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
