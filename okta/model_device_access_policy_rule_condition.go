/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// DeviceAccessPolicyRuleCondition struct for DeviceAccessPolicyRuleCondition
type DeviceAccessPolicyRuleCondition struct {
	Migrated *bool `json:"migrated,omitempty"`
	Platform *DevicePolicyRuleConditionPlatform `json:"platform,omitempty"`
	Rooted *bool `json:"rooted,omitempty"`
	TrustLevel *string `json:"trustLevel,omitempty"`
	Assurance *DevicePolicyRuleConditionAssurance `json:"assurance,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	Registered *bool `json:"registered,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAccessPolicyRuleCondition DeviceAccessPolicyRuleCondition

// NewDeviceAccessPolicyRuleCondition instantiates a new DeviceAccessPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAccessPolicyRuleCondition() *DeviceAccessPolicyRuleCondition {
	this := DeviceAccessPolicyRuleCondition{}
	return &this
}

// NewDeviceAccessPolicyRuleConditionWithDefaults instantiates a new DeviceAccessPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAccessPolicyRuleConditionWithDefaults() *DeviceAccessPolicyRuleCondition {
	this := DeviceAccessPolicyRuleCondition{}
	return &this
}

// GetMigrated returns the Migrated field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetMigrated() bool {
	if o == nil || o.Migrated == nil {
		var ret bool
		return ret
	}
	return *o.Migrated
}

// GetMigratedOk returns a tuple with the Migrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetMigratedOk() (*bool, bool) {
	if o == nil || o.Migrated == nil {
		return nil, false
	}
	return o.Migrated, true
}

// HasMigrated returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasMigrated() bool {
	if o != nil && o.Migrated != nil {
		return true
	}

	return false
}

// SetMigrated gets a reference to the given bool and assigns it to the Migrated field.
func (o *DeviceAccessPolicyRuleCondition) SetMigrated(v bool) {
	o.Migrated = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetPlatform() DevicePolicyRuleConditionPlatform {
	if o == nil || o.Platform == nil {
		var ret DevicePolicyRuleConditionPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetPlatformOk() (*DevicePolicyRuleConditionPlatform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given DevicePolicyRuleConditionPlatform and assigns it to the Platform field.
func (o *DeviceAccessPolicyRuleCondition) SetPlatform(v DevicePolicyRuleConditionPlatform) {
	o.Platform = &v
}

// GetRooted returns the Rooted field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetRooted() bool {
	if o == nil || o.Rooted == nil {
		var ret bool
		return ret
	}
	return *o.Rooted
}

// GetRootedOk returns a tuple with the Rooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetRootedOk() (*bool, bool) {
	if o == nil || o.Rooted == nil {
		return nil, false
	}
	return o.Rooted, true
}

// HasRooted returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasRooted() bool {
	if o != nil && o.Rooted != nil {
		return true
	}

	return false
}

// SetRooted gets a reference to the given bool and assigns it to the Rooted field.
func (o *DeviceAccessPolicyRuleCondition) SetRooted(v bool) {
	o.Rooted = &v
}

// GetTrustLevel returns the TrustLevel field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetTrustLevel() string {
	if o == nil || o.TrustLevel == nil {
		var ret string
		return ret
	}
	return *o.TrustLevel
}

// GetTrustLevelOk returns a tuple with the TrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetTrustLevelOk() (*string, bool) {
	if o == nil || o.TrustLevel == nil {
		return nil, false
	}
	return o.TrustLevel, true
}

// HasTrustLevel returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasTrustLevel() bool {
	if o != nil && o.TrustLevel != nil {
		return true
	}

	return false
}

// SetTrustLevel gets a reference to the given string and assigns it to the TrustLevel field.
func (o *DeviceAccessPolicyRuleCondition) SetTrustLevel(v string) {
	o.TrustLevel = &v
}

// GetAssurance returns the Assurance field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetAssurance() DevicePolicyRuleConditionAssurance {
	if o == nil || o.Assurance == nil {
		var ret DevicePolicyRuleConditionAssurance
		return ret
	}
	return *o.Assurance
}

// GetAssuranceOk returns a tuple with the Assurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetAssuranceOk() (*DevicePolicyRuleConditionAssurance, bool) {
	if o == nil || o.Assurance == nil {
		return nil, false
	}
	return o.Assurance, true
}

// HasAssurance returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasAssurance() bool {
	if o != nil && o.Assurance != nil {
		return true
	}

	return false
}

// SetAssurance gets a reference to the given DevicePolicyRuleConditionAssurance and assigns it to the Assurance field.
func (o *DeviceAccessPolicyRuleCondition) SetAssurance(v DevicePolicyRuleConditionAssurance) {
	o.Assurance = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *DeviceAccessPolicyRuleCondition) SetManaged(v bool) {
	o.Managed = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetRegistered() bool {
	if o == nil || o.Registered == nil {
		var ret bool
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetRegisteredOk() (*bool, bool) {
	if o == nil || o.Registered == nil {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasRegistered() bool {
	if o != nil && o.Registered != nil {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given bool and assigns it to the Registered field.
func (o *DeviceAccessPolicyRuleCondition) SetRegistered(v bool) {
	o.Registered = &v
}

func (o DeviceAccessPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Migrated != nil {
		toSerialize["migrated"] = o.Migrated
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Rooted != nil {
		toSerialize["rooted"] = o.Rooted
	}
	if o.TrustLevel != nil {
		toSerialize["trustLevel"] = o.TrustLevel
	}
	if o.Assurance != nil {
		toSerialize["assurance"] = o.Assurance
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAccessPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAccessPolicyRuleCondition := _DeviceAccessPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varDeviceAccessPolicyRuleCondition)
	if err == nil {
		*o = DeviceAccessPolicyRuleCondition(varDeviceAccessPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "migrated")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "rooted")
		delete(additionalProperties, "trustLevel")
		delete(additionalProperties, "assurance")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "registered")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceAccessPolicyRuleCondition struct {
	value *DeviceAccessPolicyRuleCondition
	isSet bool
}

func (v NullableDeviceAccessPolicyRuleCondition) Get() *DeviceAccessPolicyRuleCondition {
	return v.value
}

func (v *NullableDeviceAccessPolicyRuleCondition) Set(val *DeviceAccessPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAccessPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAccessPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAccessPolicyRuleCondition(val *DeviceAccessPolicyRuleCondition) *NullableDeviceAccessPolicyRuleCondition {
	return &NullableDeviceAccessPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableDeviceAccessPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAccessPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

