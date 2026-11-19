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

// checks if the ImportRoadblockEntitlementThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockEntitlementThreshold{}

// ImportRoadblockEntitlementThreshold An entitlement roadblock threshold. For `percentageThreshold` the values are percentages of users; for `absoluteThreshold` they're numbers of users.
type ImportRoadblockEntitlementThreshold struct {
	// Indicates whether this threshold is enforced
	Enabled *bool `json:"enabled,omitempty"`
	// The org-level default for this threshold
	OrgDefault *float32 `json:"orgDefault,omitempty"`
	// The threshold set for the app instance. `-1` means the app instance uses `orgDefault`.
	Value                *float32 `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblockEntitlementThreshold ImportRoadblockEntitlementThreshold

// NewImportRoadblockEntitlementThreshold instantiates a new ImportRoadblockEntitlementThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockEntitlementThreshold() *ImportRoadblockEntitlementThreshold {
	this := ImportRoadblockEntitlementThreshold{}
	return &this
}

// NewImportRoadblockEntitlementThresholdWithDefaults instantiates a new ImportRoadblockEntitlementThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockEntitlementThresholdWithDefaults() *ImportRoadblockEntitlementThreshold {
	this := ImportRoadblockEntitlementThreshold{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementThreshold) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementThreshold) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementThreshold) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ImportRoadblockEntitlementThreshold) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrgDefault returns the OrgDefault field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementThreshold) GetOrgDefault() float32 {
	if o == nil || IsNil(o.OrgDefault) {
		var ret float32
		return ret
	}
	return *o.OrgDefault
}

// GetOrgDefaultOk returns a tuple with the OrgDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementThreshold) GetOrgDefaultOk() (*float32, bool) {
	if o == nil || IsNil(o.OrgDefault) {
		return nil, false
	}
	return o.OrgDefault, true
}

// HasOrgDefault returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementThreshold) HasOrgDefault() bool {
	if o != nil && !IsNil(o.OrgDefault) {
		return true
	}

	return false
}

// SetOrgDefault gets a reference to the given float32 and assigns it to the OrgDefault field.
func (o *ImportRoadblockEntitlementThreshold) SetOrgDefault(v float32) {
	o.OrgDefault = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementThreshold) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementThreshold) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementThreshold) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ImportRoadblockEntitlementThreshold) SetValue(v float32) {
	o.Value = &v
}

func (o ImportRoadblockEntitlementThreshold) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockEntitlementThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.OrgDefault) {
		toSerialize["orgDefault"] = o.OrgDefault
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockEntitlementThreshold) UnmarshalJSON(data []byte) (err error) {
	varImportRoadblockEntitlementThreshold := _ImportRoadblockEntitlementThreshold{}

	err = json.Unmarshal(data, &varImportRoadblockEntitlementThreshold)

	if err != nil {
		return err
	}

	*o = ImportRoadblockEntitlementThreshold(varImportRoadblockEntitlementThreshold)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "orgDefault")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockEntitlementThreshold struct {
	value *ImportRoadblockEntitlementThreshold
	isSet bool
}

func (v NullableImportRoadblockEntitlementThreshold) Get() *ImportRoadblockEntitlementThreshold {
	return v.value
}

func (v *NullableImportRoadblockEntitlementThreshold) Set(val *ImportRoadblockEntitlementThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockEntitlementThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockEntitlementThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockEntitlementThreshold(val *ImportRoadblockEntitlementThreshold) *NullableImportRoadblockEntitlementThreshold {
	return &NullableImportRoadblockEntitlementThreshold{value: val, isSet: true}
}

func (v NullableImportRoadblockEntitlementThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockEntitlementThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
