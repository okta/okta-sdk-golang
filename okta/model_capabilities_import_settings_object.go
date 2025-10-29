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

// checks if the CapabilitiesImportSettingsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesImportSettingsObject{}

// CapabilitiesImportSettingsObject Defines import settings
type CapabilitiesImportSettingsObject struct {
	Schedule             *ImportScheduleObject `json:"schedule,omitempty"`
	Username             *ImportUsernameObject `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesImportSettingsObject CapabilitiesImportSettingsObject

// NewCapabilitiesImportSettingsObject instantiates a new CapabilitiesImportSettingsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesImportSettingsObject() *CapabilitiesImportSettingsObject {
	this := CapabilitiesImportSettingsObject{}
	return &this
}

// NewCapabilitiesImportSettingsObjectWithDefaults instantiates a new CapabilitiesImportSettingsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesImportSettingsObjectWithDefaults() *CapabilitiesImportSettingsObject {
	this := CapabilitiesImportSettingsObject{}
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CapabilitiesImportSettingsObject) GetSchedule() ImportScheduleObject {
	if o == nil || IsNil(o.Schedule) {
		var ret ImportScheduleObject
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportSettingsObject) GetScheduleOk() (*ImportScheduleObject, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CapabilitiesImportSettingsObject) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ImportScheduleObject and assigns it to the Schedule field.
func (o *CapabilitiesImportSettingsObject) SetSchedule(v ImportScheduleObject) {
	o.Schedule = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CapabilitiesImportSettingsObject) GetUsername() ImportUsernameObject {
	if o == nil || IsNil(o.Username) {
		var ret ImportUsernameObject
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportSettingsObject) GetUsernameOk() (*ImportUsernameObject, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CapabilitiesImportSettingsObject) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given ImportUsernameObject and assigns it to the Username field.
func (o *CapabilitiesImportSettingsObject) SetUsername(v ImportUsernameObject) {
	o.Username = &v
}

func (o CapabilitiesImportSettingsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesImportSettingsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitiesImportSettingsObject) UnmarshalJSON(data []byte) (err error) {
	varCapabilitiesImportSettingsObject := _CapabilitiesImportSettingsObject{}

	err = json.Unmarshal(data, &varCapabilitiesImportSettingsObject)

	if err != nil {
		return err
	}

	*o = CapabilitiesImportSettingsObject(varCapabilitiesImportSettingsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitiesImportSettingsObject struct {
	value *CapabilitiesImportSettingsObject
	isSet bool
}

func (v NullableCapabilitiesImportSettingsObject) Get() *CapabilitiesImportSettingsObject {
	return v.value
}

func (v *NullableCapabilitiesImportSettingsObject) Set(val *CapabilitiesImportSettingsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesImportSettingsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesImportSettingsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesImportSettingsObject(val *CapabilitiesImportSettingsObject) *NullableCapabilitiesImportSettingsObject {
	return &NullableCapabilitiesImportSettingsObject{value: val, isSet: true}
}

func (v NullableCapabilitiesImportSettingsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesImportSettingsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
