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
)

// checks if the ImportRoadblockEntitlementSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockEntitlementSettings{}

// ImportRoadblockEntitlementSettings The entitlement safeguard settings of the app instance. Only returned for entitlement roadblocks (`ENTITLEMENT_VALUE_CHANGE` and `ENTITLEMENT_SCHEMA_CHANGE`).
type ImportRoadblockEntitlementSettings struct {
	EntitlementRoadblock  *ImportRoadblockEntitlementSettingsEntitlementRoadblock  `json:"entitlementRoadblock,omitempty"`
	EntitlementSchemaLock *ImportRoadblockEntitlementSettingsEntitlementSchemaLock `json:"entitlementSchemaLock,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ImportRoadblockEntitlementSettings ImportRoadblockEntitlementSettings

// NewImportRoadblockEntitlementSettings instantiates a new ImportRoadblockEntitlementSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockEntitlementSettings() *ImportRoadblockEntitlementSettings {
	this := ImportRoadblockEntitlementSettings{}
	return &this
}

// NewImportRoadblockEntitlementSettingsWithDefaults instantiates a new ImportRoadblockEntitlementSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockEntitlementSettingsWithDefaults() *ImportRoadblockEntitlementSettings {
	this := ImportRoadblockEntitlementSettings{}
	return &this
}

// GetEntitlementRoadblock returns the EntitlementRoadblock field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettings) GetEntitlementRoadblock() ImportRoadblockEntitlementSettingsEntitlementRoadblock {
	if o == nil || IsNil(o.EntitlementRoadblock) {
		var ret ImportRoadblockEntitlementSettingsEntitlementRoadblock
		return ret
	}
	return *o.EntitlementRoadblock
}

// GetEntitlementRoadblockOk returns a tuple with the EntitlementRoadblock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettings) GetEntitlementRoadblockOk() (*ImportRoadblockEntitlementSettingsEntitlementRoadblock, bool) {
	if o == nil || IsNil(o.EntitlementRoadblock) {
		return nil, false
	}
	return o.EntitlementRoadblock, true
}

// HasEntitlementRoadblock returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettings) HasEntitlementRoadblock() bool {
	if o != nil && !IsNil(o.EntitlementRoadblock) {
		return true
	}

	return false
}

// SetEntitlementRoadblock gets a reference to the given ImportRoadblockEntitlementSettingsEntitlementRoadblock and assigns it to the EntitlementRoadblock field.
func (o *ImportRoadblockEntitlementSettings) SetEntitlementRoadblock(v ImportRoadblockEntitlementSettingsEntitlementRoadblock) {
	o.EntitlementRoadblock = &v
}

// GetEntitlementSchemaLock returns the EntitlementSchemaLock field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettings) GetEntitlementSchemaLock() ImportRoadblockEntitlementSettingsEntitlementSchemaLock {
	if o == nil || IsNil(o.EntitlementSchemaLock) {
		var ret ImportRoadblockEntitlementSettingsEntitlementSchemaLock
		return ret
	}
	return *o.EntitlementSchemaLock
}

// GetEntitlementSchemaLockOk returns a tuple with the EntitlementSchemaLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettings) GetEntitlementSchemaLockOk() (*ImportRoadblockEntitlementSettingsEntitlementSchemaLock, bool) {
	if o == nil || IsNil(o.EntitlementSchemaLock) {
		return nil, false
	}
	return o.EntitlementSchemaLock, true
}

// HasEntitlementSchemaLock returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettings) HasEntitlementSchemaLock() bool {
	if o != nil && !IsNil(o.EntitlementSchemaLock) {
		return true
	}

	return false
}

// SetEntitlementSchemaLock gets a reference to the given ImportRoadblockEntitlementSettingsEntitlementSchemaLock and assigns it to the EntitlementSchemaLock field.
func (o *ImportRoadblockEntitlementSettings) SetEntitlementSchemaLock(v ImportRoadblockEntitlementSettingsEntitlementSchemaLock) {
	o.EntitlementSchemaLock = &v
}

func (o ImportRoadblockEntitlementSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockEntitlementSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntitlementRoadblock) {
		toSerialize["entitlementRoadblock"] = o.EntitlementRoadblock
	}
	if !IsNil(o.EntitlementSchemaLock) {
		toSerialize["entitlementSchemaLock"] = o.EntitlementSchemaLock
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockEntitlementSettings) UnmarshalJSON(data []byte) (err error) {
	varImportRoadblockEntitlementSettings := _ImportRoadblockEntitlementSettings{}

	err = json.Unmarshal(data, &varImportRoadblockEntitlementSettings)

	if err != nil {
		return err
	}

	*o = ImportRoadblockEntitlementSettings(varImportRoadblockEntitlementSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementRoadblock")
		delete(additionalProperties, "entitlementSchemaLock")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockEntitlementSettings struct {
	value *ImportRoadblockEntitlementSettings
	isSet bool
}

func (v NullableImportRoadblockEntitlementSettings) Get() *ImportRoadblockEntitlementSettings {
	return v.value
}

func (v *NullableImportRoadblockEntitlementSettings) Set(val *ImportRoadblockEntitlementSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockEntitlementSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockEntitlementSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockEntitlementSettings(val *ImportRoadblockEntitlementSettings) *NullableImportRoadblockEntitlementSettings {
	return &NullableImportRoadblockEntitlementSettings{value: val, isSet: true}
}

func (v NullableImportRoadblockEntitlementSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockEntitlementSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
