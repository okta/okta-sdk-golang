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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// RoleGovernance List of all user role governance sources
type RoleGovernance struct {
	Grants []RoleGovernanceSource `json:"grants,omitempty"`
	Links *LinksGovernanceSources `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGovernance RoleGovernance

// NewRoleGovernance instantiates a new RoleGovernance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGovernance() *RoleGovernance {
	this := RoleGovernance{}
	return &this
}

// NewRoleGovernanceWithDefaults instantiates a new RoleGovernance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGovernanceWithDefaults() *RoleGovernance {
	this := RoleGovernance{}
	return &this
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *RoleGovernance) GetGrants() []RoleGovernanceSource {
	if o == nil || o.Grants == nil {
		var ret []RoleGovernanceSource
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernance) GetGrantsOk() ([]RoleGovernanceSource, bool) {
	if o == nil || o.Grants == nil {
		return nil, false
	}
	return o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *RoleGovernance) HasGrants() bool {
	if o != nil && o.Grants != nil {
		return true
	}

	return false
}

// SetGrants gets a reference to the given []RoleGovernanceSource and assigns it to the Grants field.
func (o *RoleGovernance) SetGrants(v []RoleGovernanceSource) {
	o.Grants = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoleGovernance) GetLinks() LinksGovernanceSources {
	if o == nil || o.Links == nil {
		var ret LinksGovernanceSources
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernance) GetLinksOk() (*LinksGovernanceSources, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoleGovernance) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksGovernanceSources and assigns it to the Links field.
func (o *RoleGovernance) SetLinks(v LinksGovernanceSources) {
	o.Links = &v
}

func (o RoleGovernance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Grants != nil {
		toSerialize["grants"] = o.Grants
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleGovernance) UnmarshalJSON(bytes []byte) (err error) {
	varRoleGovernance := _RoleGovernance{}

	err = json.Unmarshal(bytes, &varRoleGovernance)
	if err == nil {
		*o = RoleGovernance(varRoleGovernance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "grants")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRoleGovernance struct {
	value *RoleGovernance
	isSet bool
}

func (v NullableRoleGovernance) Get() *RoleGovernance {
	return v.value
}

func (v *NullableRoleGovernance) Set(val *RoleGovernance) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGovernance) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGovernance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGovernance(val *RoleGovernance) *NullableRoleGovernance {
	return &NullableRoleGovernance{value: val, isSet: true}
}

func (v NullableRoleGovernance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGovernance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

