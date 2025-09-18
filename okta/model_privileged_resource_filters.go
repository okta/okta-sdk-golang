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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PrivilegedResourceFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedResourceFilters{}

// PrivilegedResourceFilters struct for PrivilegedResourceFilters
type PrivilegedResourceFilters struct {
	// Array of app groups whose members might be privileged app users
	AppGroups []AppGroup `json:"appGroups,omitempty"`
	// Array of organizational units where privileged app users are present
	OrganizationalUnits  []OrganizationalUnit `json:"organizationalUnits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceFilters PrivilegedResourceFilters

// NewPrivilegedResourceFilters instantiates a new PrivilegedResourceFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceFilters() *PrivilegedResourceFilters {
	this := PrivilegedResourceFilters{}
	return &this
}

// NewPrivilegedResourceFiltersWithDefaults instantiates a new PrivilegedResourceFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceFiltersWithDefaults() *PrivilegedResourceFilters {
	this := PrivilegedResourceFilters{}
	return &this
}

// GetAppGroups returns the AppGroups field value if set, zero value otherwise.
func (o *PrivilegedResourceFilters) GetAppGroups() []AppGroup {
	if o == nil || IsNil(o.AppGroups) {
		var ret []AppGroup
		return ret
	}
	return o.AppGroups
}

// GetAppGroupsOk returns a tuple with the AppGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceFilters) GetAppGroupsOk() ([]AppGroup, bool) {
	if o == nil || IsNil(o.AppGroups) {
		return nil, false
	}
	return o.AppGroups, true
}

// HasAppGroups returns a boolean if a field has been set.
func (o *PrivilegedResourceFilters) HasAppGroups() bool {
	if o != nil && !IsNil(o.AppGroups) {
		return true
	}

	return false
}

// SetAppGroups gets a reference to the given []AppGroup and assigns it to the AppGroups field.
func (o *PrivilegedResourceFilters) SetAppGroups(v []AppGroup) {
	o.AppGroups = v
}

// GetOrganizationalUnits returns the OrganizationalUnits field value if set, zero value otherwise.
func (o *PrivilegedResourceFilters) GetOrganizationalUnits() []OrganizationalUnit {
	if o == nil || IsNil(o.OrganizationalUnits) {
		var ret []OrganizationalUnit
		return ret
	}
	return o.OrganizationalUnits
}

// GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceFilters) GetOrganizationalUnitsOk() ([]OrganizationalUnit, bool) {
	if o == nil || IsNil(o.OrganizationalUnits) {
		return nil, false
	}
	return o.OrganizationalUnits, true
}

// HasOrganizationalUnits returns a boolean if a field has been set.
func (o *PrivilegedResourceFilters) HasOrganizationalUnits() bool {
	if o != nil && !IsNil(o.OrganizationalUnits) {
		return true
	}

	return false
}

// SetOrganizationalUnits gets a reference to the given []OrganizationalUnit and assigns it to the OrganizationalUnits field.
func (o *PrivilegedResourceFilters) SetOrganizationalUnits(v []OrganizationalUnit) {
	o.OrganizationalUnits = v
}

func (o PrivilegedResourceFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedResourceFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppGroups) {
		toSerialize["appGroups"] = o.AppGroups
	}
	if !IsNil(o.OrganizationalUnits) {
		toSerialize["organizationalUnits"] = o.OrganizationalUnits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivilegedResourceFilters) UnmarshalJSON(data []byte) (err error) {
	varPrivilegedResourceFilters := _PrivilegedResourceFilters{}

	err = json.Unmarshal(data, &varPrivilegedResourceFilters)

	if err != nil {
		return err
	}

	*o = PrivilegedResourceFilters(varPrivilegedResourceFilters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appGroups")
		delete(additionalProperties, "organizationalUnits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrivilegedResourceFilters struct {
	value *PrivilegedResourceFilters
	isSet bool
}

func (v NullablePrivilegedResourceFilters) Get() *PrivilegedResourceFilters {
	return v.value
}

func (v *NullablePrivilegedResourceFilters) Set(val *PrivilegedResourceFilters) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceFilters) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceFilters(val *PrivilegedResourceFilters) *NullablePrivilegedResourceFilters {
	return &NullablePrivilegedResourceFilters{value: val, isSet: true}
}

func (v NullablePrivilegedResourceFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
