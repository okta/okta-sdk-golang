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

// checks if the IAMBundleEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IAMBundleEntitlement{}

// IAMBundleEntitlement struct for IAMBundleEntitlement
type IAMBundleEntitlement struct {
	ResourceSets         []string `json:"resourceSets,omitempty"`
	Role                 *string  `json:"role,omitempty"`
	Targets              []string `json:"targets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IAMBundleEntitlement IAMBundleEntitlement

// NewIAMBundleEntitlement instantiates a new IAMBundleEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIAMBundleEntitlement() *IAMBundleEntitlement {
	this := IAMBundleEntitlement{}
	return &this
}

// NewIAMBundleEntitlementWithDefaults instantiates a new IAMBundleEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIAMBundleEntitlementWithDefaults() *IAMBundleEntitlement {
	this := IAMBundleEntitlement{}
	return &this
}

// GetResourceSets returns the ResourceSets field value if set, zero value otherwise.
func (o *IAMBundleEntitlement) GetResourceSets() []string {
	if o == nil || IsNil(o.ResourceSets) {
		var ret []string
		return ret
	}
	return o.ResourceSets
}

// GetResourceSetsOk returns a tuple with the ResourceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IAMBundleEntitlement) GetResourceSetsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourceSets) {
		return nil, false
	}
	return o.ResourceSets, true
}

// HasResourceSets returns a boolean if a field has been set.
func (o *IAMBundleEntitlement) HasResourceSets() bool {
	if o != nil && !IsNil(o.ResourceSets) {
		return true
	}

	return false
}

// SetResourceSets gets a reference to the given []string and assigns it to the ResourceSets field.
func (o *IAMBundleEntitlement) SetResourceSets(v []string) {
	o.ResourceSets = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *IAMBundleEntitlement) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IAMBundleEntitlement) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *IAMBundleEntitlement) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *IAMBundleEntitlement) SetRole(v string) {
	o.Role = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *IAMBundleEntitlement) GetTargets() []string {
	if o == nil || IsNil(o.Targets) {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IAMBundleEntitlement) GetTargetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *IAMBundleEntitlement) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *IAMBundleEntitlement) SetTargets(v []string) {
	o.Targets = v
}

func (o IAMBundleEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IAMBundleEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceSets) {
		toSerialize["resourceSets"] = o.ResourceSets
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IAMBundleEntitlement) UnmarshalJSON(data []byte) (err error) {
	varIAMBundleEntitlement := _IAMBundleEntitlement{}

	err = json.Unmarshal(data, &varIAMBundleEntitlement)

	if err != nil {
		return err
	}

	*o = IAMBundleEntitlement(varIAMBundleEntitlement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceSets")
		delete(additionalProperties, "role")
		delete(additionalProperties, "targets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIAMBundleEntitlement struct {
	value *IAMBundleEntitlement
	isSet bool
}

func (v NullableIAMBundleEntitlement) Get() *IAMBundleEntitlement {
	return v.value
}

func (v *NullableIAMBundleEntitlement) Set(val *IAMBundleEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableIAMBundleEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableIAMBundleEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIAMBundleEntitlement(val *IAMBundleEntitlement) *NullableIAMBundleEntitlement {
	return &NullableIAMBundleEntitlement{value: val, isSet: true}
}

func (v NullableIAMBundleEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIAMBundleEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
