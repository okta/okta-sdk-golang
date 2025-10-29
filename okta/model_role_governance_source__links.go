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

// checks if the RoleGovernanceSourceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleGovernanceSourceLinks{}

// RoleGovernanceSourceLinks struct for RoleGovernanceSourceLinks
type RoleGovernanceSourceLinks struct {
	Resources            *HrefObjectGovernanceResourcesLink `json:"resources,omitempty"`
	Self                 *HrefObjectSelfLink                `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGovernanceSourceLinks RoleGovernanceSourceLinks

// NewRoleGovernanceSourceLinks instantiates a new RoleGovernanceSourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGovernanceSourceLinks() *RoleGovernanceSourceLinks {
	this := RoleGovernanceSourceLinks{}
	return &this
}

// NewRoleGovernanceSourceLinksWithDefaults instantiates a new RoleGovernanceSourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGovernanceSourceLinksWithDefaults() *RoleGovernanceSourceLinks {
	this := RoleGovernanceSourceLinks{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *RoleGovernanceSourceLinks) GetResources() HrefObjectGovernanceResourcesLink {
	if o == nil || IsNil(o.Resources) {
		var ret HrefObjectGovernanceResourcesLink
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSourceLinks) GetResourcesOk() (*HrefObjectGovernanceResourcesLink, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *RoleGovernanceSourceLinks) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given HrefObjectGovernanceResourcesLink and assigns it to the Resources field.
func (o *RoleGovernanceSourceLinks) SetResources(v HrefObjectGovernanceResourcesLink) {
	o.Resources = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *RoleGovernanceSourceLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSourceLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *RoleGovernanceSourceLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *RoleGovernanceSourceLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o RoleGovernanceSourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleGovernanceSourceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleGovernanceSourceLinks) UnmarshalJSON(data []byte) (err error) {
	varRoleGovernanceSourceLinks := _RoleGovernanceSourceLinks{}

	err = json.Unmarshal(data, &varRoleGovernanceSourceLinks)

	if err != nil {
		return err
	}

	*o = RoleGovernanceSourceLinks(varRoleGovernanceSourceLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resources")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleGovernanceSourceLinks struct {
	value *RoleGovernanceSourceLinks
	isSet bool
}

func (v NullableRoleGovernanceSourceLinks) Get() *RoleGovernanceSourceLinks {
	return v.value
}

func (v *NullableRoleGovernanceSourceLinks) Set(val *RoleGovernanceSourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGovernanceSourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGovernanceSourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGovernanceSourceLinks(val *RoleGovernanceSourceLinks) *NullableRoleGovernanceSourceLinks {
	return &NullableRoleGovernanceSourceLinks{value: val, isSet: true}
}

func (v NullableRoleGovernanceSourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGovernanceSourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
