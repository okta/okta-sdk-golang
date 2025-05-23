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

// RoleAssignedUser struct for RoleAssignedUser
type RoleAssignedUser struct {
	Id *string `json:"id,omitempty"`
	Orn *string `json:"orn,omitempty"`
	Links *LinksSelfAndRoles `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleAssignedUser RoleAssignedUser

// NewRoleAssignedUser instantiates a new RoleAssignedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignedUser() *RoleAssignedUser {
	this := RoleAssignedUser{}
	return &this
}

// NewRoleAssignedUserWithDefaults instantiates a new RoleAssignedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignedUserWithDefaults() *RoleAssignedUser {
	this := RoleAssignedUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAssignedUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignedUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAssignedUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleAssignedUser) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *RoleAssignedUser) GetOrn() string {
	if o == nil || o.Orn == nil {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignedUser) GetOrnOk() (*string, bool) {
	if o == nil || o.Orn == nil {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *RoleAssignedUser) HasOrn() bool {
	if o != nil && o.Orn != nil {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *RoleAssignedUser) SetOrn(v string) {
	o.Orn = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoleAssignedUser) GetLinks() LinksSelfAndRoles {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndRoles
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignedUser) GetLinksOk() (*LinksSelfAndRoles, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoleAssignedUser) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndRoles and assigns it to the Links field.
func (o *RoleAssignedUser) SetLinks(v LinksSelfAndRoles) {
	o.Links = &v
}

func (o RoleAssignedUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Orn != nil {
		toSerialize["orn"] = o.Orn
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleAssignedUser) UnmarshalJSON(bytes []byte) (err error) {
	varRoleAssignedUser := _RoleAssignedUser{}

	err = json.Unmarshal(bytes, &varRoleAssignedUser)
	if err == nil {
		*o = RoleAssignedUser(varRoleAssignedUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRoleAssignedUser struct {
	value *RoleAssignedUser
	isSet bool
}

func (v NullableRoleAssignedUser) Get() *RoleAssignedUser {
	return v.value
}

func (v *NullableRoleAssignedUser) Set(val *RoleAssignedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignedUser(val *RoleAssignedUser) *NullableRoleAssignedUser {
	return &NullableRoleAssignedUser{value: val, isSet: true}
}

func (v NullableRoleAssignedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

