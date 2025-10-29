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

// checks if the GroupsResponseSchemaProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsResponseSchemaProfile{}

// GroupsResponseSchemaProfile The profile information of the group
type GroupsResponseSchemaProfile struct {
	Profile              *IdentitySourceGroupProfileForUpsert `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupsResponseSchemaProfile GroupsResponseSchemaProfile

// NewGroupsResponseSchemaProfile instantiates a new GroupsResponseSchemaProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsResponseSchemaProfile() *GroupsResponseSchemaProfile {
	this := GroupsResponseSchemaProfile{}
	return &this
}

// NewGroupsResponseSchemaProfileWithDefaults instantiates a new GroupsResponseSchemaProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsResponseSchemaProfileWithDefaults() *GroupsResponseSchemaProfile {
	this := GroupsResponseSchemaProfile{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *GroupsResponseSchemaProfile) GetProfile() IdentitySourceGroupProfileForUpsert {
	if o == nil || IsNil(o.Profile) {
		var ret IdentitySourceGroupProfileForUpsert
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsResponseSchemaProfile) GetProfileOk() (*IdentitySourceGroupProfileForUpsert, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *GroupsResponseSchemaProfile) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given IdentitySourceGroupProfileForUpsert and assigns it to the Profile field.
func (o *GroupsResponseSchemaProfile) SetProfile(v IdentitySourceGroupProfileForUpsert) {
	o.Profile = &v
}

func (o GroupsResponseSchemaProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsResponseSchemaProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsResponseSchemaProfile) UnmarshalJSON(data []byte) (err error) {
	varGroupsResponseSchemaProfile := _GroupsResponseSchemaProfile{}

	err = json.Unmarshal(data, &varGroupsResponseSchemaProfile)

	if err != nil {
		return err
	}

	*o = GroupsResponseSchemaProfile(varGroupsResponseSchemaProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsResponseSchemaProfile struct {
	value *GroupsResponseSchemaProfile
	isSet bool
}

func (v NullableGroupsResponseSchemaProfile) Get() *GroupsResponseSchemaProfile {
	return v.value
}

func (v *NullableGroupsResponseSchemaProfile) Set(val *GroupsResponseSchemaProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsResponseSchemaProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsResponseSchemaProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsResponseSchemaProfile(val *GroupsResponseSchemaProfile) *NullableGroupsResponseSchemaProfile {
	return &NullableGroupsResponseSchemaProfile{value: val, isSet: true}
}

func (v NullableGroupsResponseSchemaProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsResponseSchemaProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
