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

// checks if the GroupEmbeddedStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupEmbeddedStats{}

// GroupEmbeddedStats Statistics about the group
type GroupEmbeddedStats struct {
	// Number of users in the group
	UsersCount *int32 `json:"usersCount,omitempty"`
	// Number of apps associated with the group
	AppsCount *int32 `json:"appsCount,omitempty"`
	// Number of group push mappings associated with the group
	GroupPushMappingsCount *int32 `json:"groupPushMappingsCount,omitempty"`
	// Indicates if the group has admin privileges via a group-level role assignment
	HasAdminPrivlege     *bool `json:"hasAdminPrivlege,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupEmbeddedStats GroupEmbeddedStats

// NewGroupEmbeddedStats instantiates a new GroupEmbeddedStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupEmbeddedStats() *GroupEmbeddedStats {
	this := GroupEmbeddedStats{}
	return &this
}

// NewGroupEmbeddedStatsWithDefaults instantiates a new GroupEmbeddedStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupEmbeddedStatsWithDefaults() *GroupEmbeddedStats {
	this := GroupEmbeddedStats{}
	return &this
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise.
func (o *GroupEmbeddedStats) GetUsersCount() int32 {
	if o == nil || IsNil(o.UsersCount) {
		var ret int32
		return ret
	}
	return *o.UsersCount
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedStats) GetUsersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UsersCount) {
		return nil, false
	}
	return o.UsersCount, true
}

// HasUsersCount returns a boolean if a field has been set.
func (o *GroupEmbeddedStats) HasUsersCount() bool {
	if o != nil && !IsNil(o.UsersCount) {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given int32 and assigns it to the UsersCount field.
func (o *GroupEmbeddedStats) SetUsersCount(v int32) {
	o.UsersCount = &v
}

// GetAppsCount returns the AppsCount field value if set, zero value otherwise.
func (o *GroupEmbeddedStats) GetAppsCount() int32 {
	if o == nil || IsNil(o.AppsCount) {
		var ret int32
		return ret
	}
	return *o.AppsCount
}

// GetAppsCountOk returns a tuple with the AppsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedStats) GetAppsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AppsCount) {
		return nil, false
	}
	return o.AppsCount, true
}

// HasAppsCount returns a boolean if a field has been set.
func (o *GroupEmbeddedStats) HasAppsCount() bool {
	if o != nil && !IsNil(o.AppsCount) {
		return true
	}

	return false
}

// SetAppsCount gets a reference to the given int32 and assigns it to the AppsCount field.
func (o *GroupEmbeddedStats) SetAppsCount(v int32) {
	o.AppsCount = &v
}

// GetGroupPushMappingsCount returns the GroupPushMappingsCount field value if set, zero value otherwise.
func (o *GroupEmbeddedStats) GetGroupPushMappingsCount() int32 {
	if o == nil || IsNil(o.GroupPushMappingsCount) {
		var ret int32
		return ret
	}
	return *o.GroupPushMappingsCount
}

// GetGroupPushMappingsCountOk returns a tuple with the GroupPushMappingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedStats) GetGroupPushMappingsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupPushMappingsCount) {
		return nil, false
	}
	return o.GroupPushMappingsCount, true
}

// HasGroupPushMappingsCount returns a boolean if a field has been set.
func (o *GroupEmbeddedStats) HasGroupPushMappingsCount() bool {
	if o != nil && !IsNil(o.GroupPushMappingsCount) {
		return true
	}

	return false
}

// SetGroupPushMappingsCount gets a reference to the given int32 and assigns it to the GroupPushMappingsCount field.
func (o *GroupEmbeddedStats) SetGroupPushMappingsCount(v int32) {
	o.GroupPushMappingsCount = &v
}

// GetHasAdminPrivlege returns the HasAdminPrivlege field value if set, zero value otherwise.
func (o *GroupEmbeddedStats) GetHasAdminPrivlege() bool {
	if o == nil || IsNil(o.HasAdminPrivlege) {
		var ret bool
		return ret
	}
	return *o.HasAdminPrivlege
}

// GetHasAdminPrivlegeOk returns a tuple with the HasAdminPrivlege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedStats) GetHasAdminPrivlegeOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAdminPrivlege) {
		return nil, false
	}
	return o.HasAdminPrivlege, true
}

// HasHasAdminPrivlege returns a boolean if a field has been set.
func (o *GroupEmbeddedStats) HasHasAdminPrivlege() bool {
	if o != nil && !IsNil(o.HasAdminPrivlege) {
		return true
	}

	return false
}

// SetHasAdminPrivlege gets a reference to the given bool and assigns it to the HasAdminPrivlege field.
func (o *GroupEmbeddedStats) SetHasAdminPrivlege(v bool) {
	o.HasAdminPrivlege = &v
}

func (o GroupEmbeddedStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupEmbeddedStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UsersCount) {
		toSerialize["usersCount"] = o.UsersCount
	}
	if !IsNil(o.AppsCount) {
		toSerialize["appsCount"] = o.AppsCount
	}
	if !IsNil(o.GroupPushMappingsCount) {
		toSerialize["groupPushMappingsCount"] = o.GroupPushMappingsCount
	}
	if !IsNil(o.HasAdminPrivlege) {
		toSerialize["hasAdminPrivlege"] = o.HasAdminPrivlege
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupEmbeddedStats) UnmarshalJSON(data []byte) (err error) {
	varGroupEmbeddedStats := _GroupEmbeddedStats{}

	err = json.Unmarshal(data, &varGroupEmbeddedStats)

	if err != nil {
		return err
	}

	*o = GroupEmbeddedStats(varGroupEmbeddedStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "usersCount")
		delete(additionalProperties, "appsCount")
		delete(additionalProperties, "groupPushMappingsCount")
		delete(additionalProperties, "hasAdminPrivlege")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupEmbeddedStats struct {
	value *GroupEmbeddedStats
	isSet bool
}

func (v NullableGroupEmbeddedStats) Get() *GroupEmbeddedStats {
	return v.value
}

func (v *NullableGroupEmbeddedStats) Set(val *GroupEmbeddedStats) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupEmbeddedStats) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupEmbeddedStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupEmbeddedStats(val *GroupEmbeddedStats) *NullableGroupEmbeddedStats {
	return &NullableGroupEmbeddedStats{value: val, isSet: true}
}

func (v NullableGroupEmbeddedStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupEmbeddedStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
