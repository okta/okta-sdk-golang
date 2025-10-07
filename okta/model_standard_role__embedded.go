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

// checks if the StandardRoleEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardRoleEmbedded{}

// StandardRoleEmbedded Optional embedded resources for the role assignment
type StandardRoleEmbedded struct {
	Targets              *StandardRoleEmbeddedTargets `json:"targets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StandardRoleEmbedded StandardRoleEmbedded

// NewStandardRoleEmbedded instantiates a new StandardRoleEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardRoleEmbedded() *StandardRoleEmbedded {
	this := StandardRoleEmbedded{}
	return &this
}

// NewStandardRoleEmbeddedWithDefaults instantiates a new StandardRoleEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardRoleEmbeddedWithDefaults() *StandardRoleEmbedded {
	this := StandardRoleEmbedded{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *StandardRoleEmbedded) GetTargets() StandardRoleEmbeddedTargets {
	if o == nil || IsNil(o.Targets) {
		var ret StandardRoleEmbeddedTargets
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRoleEmbedded) GetTargetsOk() (*StandardRoleEmbeddedTargets, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *StandardRoleEmbedded) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given StandardRoleEmbeddedTargets and assigns it to the Targets field.
func (o *StandardRoleEmbedded) SetTargets(v StandardRoleEmbeddedTargets) {
	o.Targets = &v
}

func (o StandardRoleEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardRoleEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StandardRoleEmbedded) UnmarshalJSON(data []byte) (err error) {
	varStandardRoleEmbedded := _StandardRoleEmbedded{}

	err = json.Unmarshal(data, &varStandardRoleEmbedded)

	if err != nil {
		return err
	}

	*o = StandardRoleEmbedded(varStandardRoleEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "targets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStandardRoleEmbedded struct {
	value *StandardRoleEmbedded
	isSet bool
}

func (v NullableStandardRoleEmbedded) Get() *StandardRoleEmbedded {
	return v.value
}

func (v *NullableStandardRoleEmbedded) Set(val *StandardRoleEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardRoleEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardRoleEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardRoleEmbedded(val *StandardRoleEmbedded) *NullableStandardRoleEmbedded {
	return &NullableStandardRoleEmbedded{value: val, isSet: true}
}

func (v NullableStandardRoleEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardRoleEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
