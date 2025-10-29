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

// checks if the RoleGovernanceResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleGovernanceResources{}

// RoleGovernanceResources The resources of a grant
type RoleGovernanceResources struct {
	Resources            []RoleGovernanceResource           `json:"resources,omitempty"`
	Links                *AIAgentOperationListResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGovernanceResources RoleGovernanceResources

// NewRoleGovernanceResources instantiates a new RoleGovernanceResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGovernanceResources() *RoleGovernanceResources {
	this := RoleGovernanceResources{}
	return &this
}

// NewRoleGovernanceResourcesWithDefaults instantiates a new RoleGovernanceResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGovernanceResourcesWithDefaults() *RoleGovernanceResources {
	this := RoleGovernanceResources{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *RoleGovernanceResources) GetResources() []RoleGovernanceResource {
	if o == nil || IsNil(o.Resources) {
		var ret []RoleGovernanceResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceResources) GetResourcesOk() ([]RoleGovernanceResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *RoleGovernanceResources) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []RoleGovernanceResource and assigns it to the Resources field.
func (o *RoleGovernanceResources) SetResources(v []RoleGovernanceResource) {
	o.Resources = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoleGovernanceResources) GetLinks() AIAgentOperationListResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret AIAgentOperationListResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceResources) GetLinksOk() (*AIAgentOperationListResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoleGovernanceResources) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AIAgentOperationListResponseLinks and assigns it to the Links field.
func (o *RoleGovernanceResources) SetLinks(v AIAgentOperationListResponseLinks) {
	o.Links = &v
}

func (o RoleGovernanceResources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleGovernanceResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleGovernanceResources) UnmarshalJSON(data []byte) (err error) {
	varRoleGovernanceResources := _RoleGovernanceResources{}

	err = json.Unmarshal(data, &varRoleGovernanceResources)

	if err != nil {
		return err
	}

	*o = RoleGovernanceResources(varRoleGovernanceResources)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resources")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleGovernanceResources struct {
	value *RoleGovernanceResources
	isSet bool
}

func (v NullableRoleGovernanceResources) Get() *RoleGovernanceResources {
	return v.value
}

func (v *NullableRoleGovernanceResources) Set(val *RoleGovernanceResources) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGovernanceResources) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGovernanceResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGovernanceResources(val *RoleGovernanceResources) *NullableRoleGovernanceResources {
	return &NullableRoleGovernanceResources{value: val, isSet: true}
}

func (v NullableRoleGovernanceResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGovernanceResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
