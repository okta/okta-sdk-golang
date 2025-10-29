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

// checks if the IdentitySourceGroupMembershipsUpsertProfileInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySourceGroupMembershipsUpsertProfileInner{}

// IdentitySourceGroupMembershipsUpsertProfileInner struct for IdentitySourceGroupMembershipsUpsertProfileInner
type IdentitySourceGroupMembershipsUpsertProfileInner struct {
	// The external ID of the group whose memberships need to be inserted or updated in Okta
	GroupExternalId *string `json:"groupExternalId,omitempty"`
	// Array of external IDs of member profiles that need to be inserted in this group in Okta
	MemberExternalIds    []string `json:"memberExternalIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySourceGroupMembershipsUpsertProfileInner IdentitySourceGroupMembershipsUpsertProfileInner

// NewIdentitySourceGroupMembershipsUpsertProfileInner instantiates a new IdentitySourceGroupMembershipsUpsertProfileInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourceGroupMembershipsUpsertProfileInner() *IdentitySourceGroupMembershipsUpsertProfileInner {
	this := IdentitySourceGroupMembershipsUpsertProfileInner{}
	return &this
}

// NewIdentitySourceGroupMembershipsUpsertProfileInnerWithDefaults instantiates a new IdentitySourceGroupMembershipsUpsertProfileInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourceGroupMembershipsUpsertProfileInnerWithDefaults() *IdentitySourceGroupMembershipsUpsertProfileInner {
	this := IdentitySourceGroupMembershipsUpsertProfileInner{}
	return &this
}

// GetGroupExternalId returns the GroupExternalId field value if set, zero value otherwise.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetGroupExternalId() string {
	if o == nil || IsNil(o.GroupExternalId) {
		var ret string
		return ret
	}
	return *o.GroupExternalId
}

// GetGroupExternalIdOk returns a tuple with the GroupExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetGroupExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupExternalId) {
		return nil, false
	}
	return o.GroupExternalId, true
}

// HasGroupExternalId returns a boolean if a field has been set.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) HasGroupExternalId() bool {
	if o != nil && !IsNil(o.GroupExternalId) {
		return true
	}

	return false
}

// SetGroupExternalId gets a reference to the given string and assigns it to the GroupExternalId field.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) SetGroupExternalId(v string) {
	o.GroupExternalId = &v
}

// GetMemberExternalIds returns the MemberExternalIds field value if set, zero value otherwise.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetMemberExternalIds() []string {
	if o == nil || IsNil(o.MemberExternalIds) {
		var ret []string
		return ret
	}
	return o.MemberExternalIds
}

// GetMemberExternalIdsOk returns a tuple with the MemberExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetMemberExternalIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MemberExternalIds) {
		return nil, false
	}
	return o.MemberExternalIds, true
}

// HasMemberExternalIds returns a boolean if a field has been set.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) HasMemberExternalIds() bool {
	if o != nil && !IsNil(o.MemberExternalIds) {
		return true
	}

	return false
}

// SetMemberExternalIds gets a reference to the given []string and assigns it to the MemberExternalIds field.
func (o *IdentitySourceGroupMembershipsUpsertProfileInner) SetMemberExternalIds(v []string) {
	o.MemberExternalIds = v
}

func (o IdentitySourceGroupMembershipsUpsertProfileInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySourceGroupMembershipsUpsertProfileInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupExternalId) {
		toSerialize["groupExternalId"] = o.GroupExternalId
	}
	if !IsNil(o.MemberExternalIds) {
		toSerialize["memberExternalIds"] = o.MemberExternalIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySourceGroupMembershipsUpsertProfileInner) UnmarshalJSON(data []byte) (err error) {
	varIdentitySourceGroupMembershipsUpsertProfileInner := _IdentitySourceGroupMembershipsUpsertProfileInner{}

	err = json.Unmarshal(data, &varIdentitySourceGroupMembershipsUpsertProfileInner)

	if err != nil {
		return err
	}

	*o = IdentitySourceGroupMembershipsUpsertProfileInner(varIdentitySourceGroupMembershipsUpsertProfileInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groupExternalId")
		delete(additionalProperties, "memberExternalIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySourceGroupMembershipsUpsertProfileInner struct {
	value *IdentitySourceGroupMembershipsUpsertProfileInner
	isSet bool
}

func (v NullableIdentitySourceGroupMembershipsUpsertProfileInner) Get() *IdentitySourceGroupMembershipsUpsertProfileInner {
	return v.value
}

func (v *NullableIdentitySourceGroupMembershipsUpsertProfileInner) Set(val *IdentitySourceGroupMembershipsUpsertProfileInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourceGroupMembershipsUpsertProfileInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourceGroupMembershipsUpsertProfileInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourceGroupMembershipsUpsertProfileInner(val *IdentitySourceGroupMembershipsUpsertProfileInner) *NullableIdentitySourceGroupMembershipsUpsertProfileInner {
	return &NullableIdentitySourceGroupMembershipsUpsertProfileInner{value: val, isSet: true}
}

func (v NullableIdentitySourceGroupMembershipsUpsertProfileInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourceGroupMembershipsUpsertProfileInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
