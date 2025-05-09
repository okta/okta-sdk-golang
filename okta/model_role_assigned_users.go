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

// RoleAssignedUsers struct for RoleAssignedUsers
type RoleAssignedUsers struct {
	Value []RoleAssignedUser `json:"value,omitempty"`
	Links *LinksNext `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleAssignedUsers RoleAssignedUsers

// NewRoleAssignedUsers instantiates a new RoleAssignedUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignedUsers() *RoleAssignedUsers {
	this := RoleAssignedUsers{}
	return &this
}

// NewRoleAssignedUsersWithDefaults instantiates a new RoleAssignedUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignedUsersWithDefaults() *RoleAssignedUsers {
	this := RoleAssignedUsers{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RoleAssignedUsers) GetValue() []RoleAssignedUser {
	if o == nil || o.Value == nil {
		var ret []RoleAssignedUser
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignedUsers) GetValueOk() ([]RoleAssignedUser, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RoleAssignedUsers) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []RoleAssignedUser and assigns it to the Value field.
func (o *RoleAssignedUsers) SetValue(v []RoleAssignedUser) {
	o.Value = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoleAssignedUsers) GetLinks() LinksNext {
	if o == nil || o.Links == nil {
		var ret LinksNext
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignedUsers) GetLinksOk() (*LinksNext, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoleAssignedUsers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksNext and assigns it to the Links field.
func (o *RoleAssignedUsers) SetLinks(v LinksNext) {
	o.Links = &v
}

func (o RoleAssignedUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleAssignedUsers) UnmarshalJSON(bytes []byte) (err error) {
	varRoleAssignedUsers := _RoleAssignedUsers{}

	err = json.Unmarshal(bytes, &varRoleAssignedUsers)
	if err == nil {
		*o = RoleAssignedUsers(varRoleAssignedUsers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRoleAssignedUsers struct {
	value *RoleAssignedUsers
	isSet bool
}

func (v NullableRoleAssignedUsers) Get() *RoleAssignedUsers {
	return v.value
}

func (v *NullableRoleAssignedUsers) Set(val *RoleAssignedUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignedUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignedUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignedUsers(val *RoleAssignedUsers) *NullableRoleAssignedUsers {
	return &NullableRoleAssignedUsers{value: val, isSet: true}
}

func (v NullableRoleAssignedUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignedUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

