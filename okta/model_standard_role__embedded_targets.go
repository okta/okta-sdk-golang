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

// checks if the StandardRoleEmbeddedTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardRoleEmbeddedTargets{}

// StandardRoleEmbeddedTargets Targets configured for the role assignment
type StandardRoleEmbeddedTargets struct {
	// Group targets
	Groups               []Group                             `json:"groups,omitempty"`
	Catalog              *StandardRoleEmbeddedTargetsCatalog `json:"catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StandardRoleEmbeddedTargets StandardRoleEmbeddedTargets

// NewStandardRoleEmbeddedTargets instantiates a new StandardRoleEmbeddedTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardRoleEmbeddedTargets() *StandardRoleEmbeddedTargets {
	this := StandardRoleEmbeddedTargets{}
	return &this
}

// NewStandardRoleEmbeddedTargetsWithDefaults instantiates a new StandardRoleEmbeddedTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardRoleEmbeddedTargetsWithDefaults() *StandardRoleEmbeddedTargets {
	this := StandardRoleEmbeddedTargets{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *StandardRoleEmbeddedTargets) GetGroups() []Group {
	if o == nil || IsNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRoleEmbeddedTargets) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *StandardRoleEmbeddedTargets) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *StandardRoleEmbeddedTargets) SetGroups(v []Group) {
	o.Groups = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *StandardRoleEmbeddedTargets) GetCatalog() StandardRoleEmbeddedTargetsCatalog {
	if o == nil || IsNil(o.Catalog) {
		var ret StandardRoleEmbeddedTargetsCatalog
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRoleEmbeddedTargets) GetCatalogOk() (*StandardRoleEmbeddedTargetsCatalog, bool) {
	if o == nil || IsNil(o.Catalog) {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *StandardRoleEmbeddedTargets) HasCatalog() bool {
	if o != nil && !IsNil(o.Catalog) {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given StandardRoleEmbeddedTargetsCatalog and assigns it to the Catalog field.
func (o *StandardRoleEmbeddedTargets) SetCatalog(v StandardRoleEmbeddedTargetsCatalog) {
	o.Catalog = &v
}

func (o StandardRoleEmbeddedTargets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardRoleEmbeddedTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Catalog) {
		toSerialize["catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StandardRoleEmbeddedTargets) UnmarshalJSON(data []byte) (err error) {
	varStandardRoleEmbeddedTargets := _StandardRoleEmbeddedTargets{}

	err = json.Unmarshal(data, &varStandardRoleEmbeddedTargets)

	if err != nil {
		return err
	}

	*o = StandardRoleEmbeddedTargets(varStandardRoleEmbeddedTargets)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groups")
		delete(additionalProperties, "catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStandardRoleEmbeddedTargets struct {
	value *StandardRoleEmbeddedTargets
	isSet bool
}

func (v NullableStandardRoleEmbeddedTargets) Get() *StandardRoleEmbeddedTargets {
	return v.value
}

func (v *NullableStandardRoleEmbeddedTargets) Set(val *StandardRoleEmbeddedTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardRoleEmbeddedTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardRoleEmbeddedTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardRoleEmbeddedTargets(val *StandardRoleEmbeddedTargets) *NullableStandardRoleEmbeddedTargets {
	return &NullableStandardRoleEmbeddedTargets{value: val, isSet: true}
}

func (v NullableStandardRoleEmbeddedTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardRoleEmbeddedTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
