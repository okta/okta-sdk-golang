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

// checks if the PolicyAccountLinkFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAccountLinkFilter{}

// PolicyAccountLinkFilter Specifies filters on which users are available for account linking by an IdP
type PolicyAccountLinkFilter struct {
	Groups               *PolicyAccountLinkFilterGroups `json:"groups,omitempty"`
	Users                *PolicyAccountLinkFilterUsers  `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAccountLinkFilter PolicyAccountLinkFilter

// NewPolicyAccountLinkFilter instantiates a new PolicyAccountLinkFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAccountLinkFilter() *PolicyAccountLinkFilter {
	this := PolicyAccountLinkFilter{}
	return &this
}

// NewPolicyAccountLinkFilterWithDefaults instantiates a new PolicyAccountLinkFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAccountLinkFilterWithDefaults() *PolicyAccountLinkFilter {
	this := PolicyAccountLinkFilter{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilter) GetGroups() PolicyAccountLinkFilterGroups {
	if o == nil || IsNil(o.Groups) {
		var ret PolicyAccountLinkFilterGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilter) GetGroupsOk() (*PolicyAccountLinkFilterGroups, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilter) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given PolicyAccountLinkFilterGroups and assigns it to the Groups field.
func (o *PolicyAccountLinkFilter) SetGroups(v PolicyAccountLinkFilterGroups) {
	o.Groups = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilter) GetUsers() PolicyAccountLinkFilterUsers {
	if o == nil || IsNil(o.Users) {
		var ret PolicyAccountLinkFilterUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilter) GetUsersOk() (*PolicyAccountLinkFilterUsers, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilter) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given PolicyAccountLinkFilterUsers and assigns it to the Users field.
func (o *PolicyAccountLinkFilter) SetUsers(v PolicyAccountLinkFilterUsers) {
	o.Users = &v
}

func (o PolicyAccountLinkFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAccountLinkFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAccountLinkFilter) UnmarshalJSON(data []byte) (err error) {
	varPolicyAccountLinkFilter := _PolicyAccountLinkFilter{}

	err = json.Unmarshal(data, &varPolicyAccountLinkFilter)

	if err != nil {
		return err
	}

	*o = PolicyAccountLinkFilter(varPolicyAccountLinkFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAccountLinkFilter struct {
	value *PolicyAccountLinkFilter
	isSet bool
}

func (v NullablePolicyAccountLinkFilter) Get() *PolicyAccountLinkFilter {
	return v.value
}

func (v *NullablePolicyAccountLinkFilter) Set(val *PolicyAccountLinkFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccountLinkFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccountLinkFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccountLinkFilter(val *PolicyAccountLinkFilter) *NullablePolicyAccountLinkFilter {
	return &NullablePolicyAccountLinkFilter{value: val, isSet: true}
}

func (v NullablePolicyAccountLinkFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccountLinkFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
