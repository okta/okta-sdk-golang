/*
Okta Admin Management API

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
	"time"
)

// checks if the ImportRoadblockEntitlementSettingsEntitlementSchemaLock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockEntitlementSettingsEntitlementSchemaLock{}

// ImportRoadblockEntitlementSettingsEntitlementSchemaLock Lock on the entitlement schema of the app instance that prevents schema changes
type ImportRoadblockEntitlementSettingsEntitlementSchemaLock struct {
	// Indicates whether the entitlement schema is locked
	Locked *bool `json:"locked,omitempty"`
	// Timestamp when the entitlement schema was locked
	LockedAt *time.Time `json:"lockedAt,omitempty"`
	// ID of the user who locked the entitlement schema
	LockedBy             *string `json:"lockedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblockEntitlementSettingsEntitlementSchemaLock ImportRoadblockEntitlementSettingsEntitlementSchemaLock

// NewImportRoadblockEntitlementSettingsEntitlementSchemaLock instantiates a new ImportRoadblockEntitlementSettingsEntitlementSchemaLock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockEntitlementSettingsEntitlementSchemaLock() *ImportRoadblockEntitlementSettingsEntitlementSchemaLock {
	this := ImportRoadblockEntitlementSettingsEntitlementSchemaLock{}
	return &this
}

// NewImportRoadblockEntitlementSettingsEntitlementSchemaLockWithDefaults instantiates a new ImportRoadblockEntitlementSettingsEntitlementSchemaLock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockEntitlementSettingsEntitlementSchemaLockWithDefaults() *ImportRoadblockEntitlementSettingsEntitlementSchemaLock {
	this := ImportRoadblockEntitlementSettingsEntitlementSchemaLock{}
	return &this
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) SetLocked(v bool) {
	o.Locked = &v
}

// GetLockedAt returns the LockedAt field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedAt() time.Time {
	if o == nil || IsNil(o.LockedAt) {
		var ret time.Time
		return ret
	}
	return *o.LockedAt
}

// GetLockedAtOk returns a tuple with the LockedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LockedAt) {
		return nil, false
	}
	return o.LockedAt, true
}

// HasLockedAt returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) HasLockedAt() bool {
	if o != nil && !IsNil(o.LockedAt) {
		return true
	}

	return false
}

// SetLockedAt gets a reference to the given time.Time and assigns it to the LockedAt field.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) SetLockedAt(v time.Time) {
	o.LockedAt = &v
}

// GetLockedBy returns the LockedBy field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedBy() string {
	if o == nil || IsNil(o.LockedBy) {
		var ret string
		return ret
	}
	return *o.LockedBy
}

// GetLockedByOk returns a tuple with the LockedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedByOk() (*string, bool) {
	if o == nil || IsNil(o.LockedBy) {
		return nil, false
	}
	return o.LockedBy, true
}

// HasLockedBy returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) HasLockedBy() bool {
	if o != nil && !IsNil(o.LockedBy) {
		return true
	}

	return false
}

// SetLockedBy gets a reference to the given string and assigns it to the LockedBy field.
func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) SetLockedBy(v string) {
	o.LockedBy = &v
}

func (o ImportRoadblockEntitlementSettingsEntitlementSchemaLock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockEntitlementSettingsEntitlementSchemaLock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.LockedAt) {
		toSerialize["lockedAt"] = o.LockedAt
	}
	if !IsNil(o.LockedBy) {
		toSerialize["lockedBy"] = o.LockedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) UnmarshalJSON(data []byte) (err error) {
	varImportRoadblockEntitlementSettingsEntitlementSchemaLock := _ImportRoadblockEntitlementSettingsEntitlementSchemaLock{}

	err = json.Unmarshal(data, &varImportRoadblockEntitlementSettingsEntitlementSchemaLock)

	if err != nil {
		return err
	}

	*o = ImportRoadblockEntitlementSettingsEntitlementSchemaLock(varImportRoadblockEntitlementSettingsEntitlementSchemaLock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "locked")
		delete(additionalProperties, "lockedAt")
		delete(additionalProperties, "lockedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock struct {
	value *ImportRoadblockEntitlementSettingsEntitlementSchemaLock
	isSet bool
}

func (v NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock) Get() *ImportRoadblockEntitlementSettingsEntitlementSchemaLock {
	return v.value
}

func (v *NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock) Set(val *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockEntitlementSettingsEntitlementSchemaLock(val *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) *NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock {
	return &NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock{value: val, isSet: true}
}

func (v NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockEntitlementSettingsEntitlementSchemaLock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
