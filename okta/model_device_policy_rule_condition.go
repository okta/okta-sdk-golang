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

// checks if the DevicePolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePolicyRuleCondition{}

// DevicePolicyRuleCondition struct for DevicePolicyRuleCondition
type DevicePolicyRuleCondition struct {
	Migrated             *bool                              `json:"migrated,omitempty"`
	Platform             *DevicePolicyRuleConditionPlatform `json:"platform,omitempty"`
	Rooted               *bool                              `json:"rooted,omitempty"`
	TrustLevel           *string                            `json:"trustLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePolicyRuleCondition DevicePolicyRuleCondition

// NewDevicePolicyRuleCondition instantiates a new DevicePolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePolicyRuleCondition() *DevicePolicyRuleCondition {
	this := DevicePolicyRuleCondition{}
	return &this
}

// NewDevicePolicyRuleConditionWithDefaults instantiates a new DevicePolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePolicyRuleConditionWithDefaults() *DevicePolicyRuleCondition {
	this := DevicePolicyRuleCondition{}
	return &this
}

// GetMigrated returns the Migrated field value if set, zero value otherwise.
func (o *DevicePolicyRuleCondition) GetMigrated() bool {
	if o == nil || IsNil(o.Migrated) {
		var ret bool
		return ret
	}
	return *o.Migrated
}

// GetMigratedOk returns a tuple with the Migrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleCondition) GetMigratedOk() (*bool, bool) {
	if o == nil || IsNil(o.Migrated) {
		return nil, false
	}
	return o.Migrated, true
}

// HasMigrated returns a boolean if a field has been set.
func (o *DevicePolicyRuleCondition) HasMigrated() bool {
	if o != nil && !IsNil(o.Migrated) {
		return true
	}

	return false
}

// SetMigrated gets a reference to the given bool and assigns it to the Migrated field.
func (o *DevicePolicyRuleCondition) SetMigrated(v bool) {
	o.Migrated = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DevicePolicyRuleCondition) GetPlatform() DevicePolicyRuleConditionPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret DevicePolicyRuleConditionPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleCondition) GetPlatformOk() (*DevicePolicyRuleConditionPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DevicePolicyRuleCondition) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given DevicePolicyRuleConditionPlatform and assigns it to the Platform field.
func (o *DevicePolicyRuleCondition) SetPlatform(v DevicePolicyRuleConditionPlatform) {
	o.Platform = &v
}

// GetRooted returns the Rooted field value if set, zero value otherwise.
func (o *DevicePolicyRuleCondition) GetRooted() bool {
	if o == nil || IsNil(o.Rooted) {
		var ret bool
		return ret
	}
	return *o.Rooted
}

// GetRootedOk returns a tuple with the Rooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleCondition) GetRootedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rooted) {
		return nil, false
	}
	return o.Rooted, true
}

// HasRooted returns a boolean if a field has been set.
func (o *DevicePolicyRuleCondition) HasRooted() bool {
	if o != nil && !IsNil(o.Rooted) {
		return true
	}

	return false
}

// SetRooted gets a reference to the given bool and assigns it to the Rooted field.
func (o *DevicePolicyRuleCondition) SetRooted(v bool) {
	o.Rooted = &v
}

// GetTrustLevel returns the TrustLevel field value if set, zero value otherwise.
func (o *DevicePolicyRuleCondition) GetTrustLevel() string {
	if o == nil || IsNil(o.TrustLevel) {
		var ret string
		return ret
	}
	return *o.TrustLevel
}

// GetTrustLevelOk returns a tuple with the TrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleCondition) GetTrustLevelOk() (*string, bool) {
	if o == nil || IsNil(o.TrustLevel) {
		return nil, false
	}
	return o.TrustLevel, true
}

// HasTrustLevel returns a boolean if a field has been set.
func (o *DevicePolicyRuleCondition) HasTrustLevel() bool {
	if o != nil && !IsNil(o.TrustLevel) {
		return true
	}

	return false
}

// SetTrustLevel gets a reference to the given string and assigns it to the TrustLevel field.
func (o *DevicePolicyRuleCondition) SetTrustLevel(v string) {
	o.TrustLevel = &v
}

func (o DevicePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Migrated) {
		toSerialize["migrated"] = o.Migrated
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Rooted) {
		toSerialize["rooted"] = o.Rooted
	}
	if !IsNil(o.TrustLevel) {
		toSerialize["trustLevel"] = o.TrustLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varDevicePolicyRuleCondition := _DevicePolicyRuleCondition{}

	err = json.Unmarshal(data, &varDevicePolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = DevicePolicyRuleCondition(varDevicePolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "migrated")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "rooted")
		delete(additionalProperties, "trustLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePolicyRuleCondition struct {
	value *DevicePolicyRuleCondition
	isSet bool
}

func (v NullableDevicePolicyRuleCondition) Get() *DevicePolicyRuleCondition {
	return v.value
}

func (v *NullableDevicePolicyRuleCondition) Set(val *DevicePolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePolicyRuleCondition(val *DevicePolicyRuleCondition) *NullableDevicePolicyRuleCondition {
	return &NullableDevicePolicyRuleCondition{value: val, isSet: true}
}

func (v NullableDevicePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
