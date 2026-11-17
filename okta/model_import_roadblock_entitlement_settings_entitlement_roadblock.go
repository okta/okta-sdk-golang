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

// checks if the ImportRoadblockEntitlementSettingsEntitlementRoadblock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockEntitlementSettingsEntitlementRoadblock{}

// ImportRoadblockEntitlementSettingsEntitlementRoadblock Thresholds that create a roadblock when an import changes the entitlements of too many users. A roadblock is created when any enabled threshold is crossed.
type ImportRoadblockEntitlementSettingsEntitlementRoadblock struct {
	// Indicates whether the entitlement roadblock is enabled for the app instance
	Enabled             *bool                                `json:"enabled,omitempty"`
	PercentageThreshold *ImportRoadblockEntitlementThreshold `json:"percentageThreshold,omitempty"`
	AbsoluteThreshold   *ImportRoadblockEntitlementThreshold `json:"absoluteThreshold,omitempty"`
	// Indicates whether an entitlement baseline has been captured for the app instance. Changes are only measured against a baseline.
	BaselineCaptured     *bool `json:"baselineCaptured,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblockEntitlementSettingsEntitlementRoadblock ImportRoadblockEntitlementSettingsEntitlementRoadblock

// NewImportRoadblockEntitlementSettingsEntitlementRoadblock instantiates a new ImportRoadblockEntitlementSettingsEntitlementRoadblock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockEntitlementSettingsEntitlementRoadblock() *ImportRoadblockEntitlementSettingsEntitlementRoadblock {
	this := ImportRoadblockEntitlementSettingsEntitlementRoadblock{}
	return &this
}

// NewImportRoadblockEntitlementSettingsEntitlementRoadblockWithDefaults instantiates a new ImportRoadblockEntitlementSettingsEntitlementRoadblock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockEntitlementSettingsEntitlementRoadblockWithDefaults() *ImportRoadblockEntitlementSettingsEntitlementRoadblock {
	this := ImportRoadblockEntitlementSettingsEntitlementRoadblock{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPercentageThreshold returns the PercentageThreshold field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetPercentageThreshold() ImportRoadblockEntitlementThreshold {
	if o == nil || IsNil(o.PercentageThreshold) {
		var ret ImportRoadblockEntitlementThreshold
		return ret
	}
	return *o.PercentageThreshold
}

// GetPercentageThresholdOk returns a tuple with the PercentageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetPercentageThresholdOk() (*ImportRoadblockEntitlementThreshold, bool) {
	if o == nil || IsNil(o.PercentageThreshold) {
		return nil, false
	}
	return o.PercentageThreshold, true
}

// HasPercentageThreshold returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasPercentageThreshold() bool {
	if o != nil && !IsNil(o.PercentageThreshold) {
		return true
	}

	return false
}

// SetPercentageThreshold gets a reference to the given ImportRoadblockEntitlementThreshold and assigns it to the PercentageThreshold field.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetPercentageThreshold(v ImportRoadblockEntitlementThreshold) {
	o.PercentageThreshold = &v
}

// GetAbsoluteThreshold returns the AbsoluteThreshold field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetAbsoluteThreshold() ImportRoadblockEntitlementThreshold {
	if o == nil || IsNil(o.AbsoluteThreshold) {
		var ret ImportRoadblockEntitlementThreshold
		return ret
	}
	return *o.AbsoluteThreshold
}

// GetAbsoluteThresholdOk returns a tuple with the AbsoluteThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetAbsoluteThresholdOk() (*ImportRoadblockEntitlementThreshold, bool) {
	if o == nil || IsNil(o.AbsoluteThreshold) {
		return nil, false
	}
	return o.AbsoluteThreshold, true
}

// HasAbsoluteThreshold returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasAbsoluteThreshold() bool {
	if o != nil && !IsNil(o.AbsoluteThreshold) {
		return true
	}

	return false
}

// SetAbsoluteThreshold gets a reference to the given ImportRoadblockEntitlementThreshold and assigns it to the AbsoluteThreshold field.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetAbsoluteThreshold(v ImportRoadblockEntitlementThreshold) {
	o.AbsoluteThreshold = &v
}

// GetBaselineCaptured returns the BaselineCaptured field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetBaselineCaptured() bool {
	if o == nil || IsNil(o.BaselineCaptured) {
		var ret bool
		return ret
	}
	return *o.BaselineCaptured
}

// GetBaselineCapturedOk returns a tuple with the BaselineCaptured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetBaselineCapturedOk() (*bool, bool) {
	if o == nil || IsNil(o.BaselineCaptured) {
		return nil, false
	}
	return o.BaselineCaptured, true
}

// HasBaselineCaptured returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasBaselineCaptured() bool {
	if o != nil && !IsNil(o.BaselineCaptured) {
		return true
	}

	return false
}

// SetBaselineCaptured gets a reference to the given bool and assigns it to the BaselineCaptured field.
func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetBaselineCaptured(v bool) {
	o.BaselineCaptured = &v
}

func (o ImportRoadblockEntitlementSettingsEntitlementRoadblock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockEntitlementSettingsEntitlementRoadblock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.PercentageThreshold) {
		toSerialize["percentageThreshold"] = o.PercentageThreshold
	}
	if !IsNil(o.AbsoluteThreshold) {
		toSerialize["absoluteThreshold"] = o.AbsoluteThreshold
	}
	if !IsNil(o.BaselineCaptured) {
		toSerialize["baselineCaptured"] = o.BaselineCaptured
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) UnmarshalJSON(data []byte) (err error) {
	varImportRoadblockEntitlementSettingsEntitlementRoadblock := _ImportRoadblockEntitlementSettingsEntitlementRoadblock{}

	err = json.Unmarshal(data, &varImportRoadblockEntitlementSettingsEntitlementRoadblock)

	if err != nil {
		return err
	}

	*o = ImportRoadblockEntitlementSettingsEntitlementRoadblock(varImportRoadblockEntitlementSettingsEntitlementRoadblock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "percentageThreshold")
		delete(additionalProperties, "absoluteThreshold")
		delete(additionalProperties, "baselineCaptured")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockEntitlementSettingsEntitlementRoadblock struct {
	value *ImportRoadblockEntitlementSettingsEntitlementRoadblock
	isSet bool
}

func (v NullableImportRoadblockEntitlementSettingsEntitlementRoadblock) Get() *ImportRoadblockEntitlementSettingsEntitlementRoadblock {
	return v.value
}

func (v *NullableImportRoadblockEntitlementSettingsEntitlementRoadblock) Set(val *ImportRoadblockEntitlementSettingsEntitlementRoadblock) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockEntitlementSettingsEntitlementRoadblock) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockEntitlementSettingsEntitlementRoadblock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockEntitlementSettingsEntitlementRoadblock(val *ImportRoadblockEntitlementSettingsEntitlementRoadblock) *NullableImportRoadblockEntitlementSettingsEntitlementRoadblock {
	return &NullableImportRoadblockEntitlementSettingsEntitlementRoadblock{value: val, isSet: true}
}

func (v NullableImportRoadblockEntitlementSettingsEntitlementRoadblock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockEntitlementSettingsEntitlementRoadblock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
