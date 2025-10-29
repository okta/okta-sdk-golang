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

// checks if the PasswordPolicyPasswordSettingsAge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyPasswordSettingsAge{}

// PasswordPolicyPasswordSettingsAge Age settings
type PasswordPolicyPasswordSettingsAge struct {
	// Specifies the number of days prior to password expiration when a User is warned to reset their password: `0` indicates no warning
	ExpireWarnDays *int32 `json:"expireWarnDays,omitempty"`
	// Specifies the number of distinct passwords that a User must create before they can reuse a previous password: `0` indicates none
	HistoryCount *int32 `json:"historyCount,omitempty"`
	// Specifies how long (in days) a password remains valid before it expires: `0` indicates no limit
	MaxAgeDays *int32 `json:"maxAgeDays,omitempty"`
	// Specifies the minimum time interval (in minutes) between password changes: `0` indicates no limit
	MinAgeMinutes        *int32 `json:"minAgeMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyPasswordSettingsAge PasswordPolicyPasswordSettingsAge

// NewPasswordPolicyPasswordSettingsAge instantiates a new PasswordPolicyPasswordSettingsAge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyPasswordSettingsAge() *PasswordPolicyPasswordSettingsAge {
	this := PasswordPolicyPasswordSettingsAge{}
	var expireWarnDays int32 = 0
	this.ExpireWarnDays = &expireWarnDays
	var historyCount int32 = 0
	this.HistoryCount = &historyCount
	var maxAgeDays int32 = 0
	this.MaxAgeDays = &maxAgeDays
	var minAgeMinutes int32 = 0
	this.MinAgeMinutes = &minAgeMinutes
	return &this
}

// NewPasswordPolicyPasswordSettingsAgeWithDefaults instantiates a new PasswordPolicyPasswordSettingsAge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsAgeWithDefaults() *PasswordPolicyPasswordSettingsAge {
	this := PasswordPolicyPasswordSettingsAge{}
	var expireWarnDays int32 = 0
	this.ExpireWarnDays = &expireWarnDays
	var historyCount int32 = 0
	this.HistoryCount = &historyCount
	var maxAgeDays int32 = 0
	this.MaxAgeDays = &maxAgeDays
	var minAgeMinutes int32 = 0
	this.MinAgeMinutes = &minAgeMinutes
	return &this
}

// GetExpireWarnDays returns the ExpireWarnDays field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsAge) GetExpireWarnDays() int32 {
	if o == nil || IsNil(o.ExpireWarnDays) {
		var ret int32
		return ret
	}
	return *o.ExpireWarnDays
}

// GetExpireWarnDaysOk returns a tuple with the ExpireWarnDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetExpireWarnDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireWarnDays) {
		return nil, false
	}
	return o.ExpireWarnDays, true
}

// HasExpireWarnDays returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasExpireWarnDays() bool {
	if o != nil && !IsNil(o.ExpireWarnDays) {
		return true
	}

	return false
}

// SetExpireWarnDays gets a reference to the given int32 and assigns it to the ExpireWarnDays field.
func (o *PasswordPolicyPasswordSettingsAge) SetExpireWarnDays(v int32) {
	o.ExpireWarnDays = &v
}

// GetHistoryCount returns the HistoryCount field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsAge) GetHistoryCount() int32 {
	if o == nil || IsNil(o.HistoryCount) {
		var ret int32
		return ret
	}
	return *o.HistoryCount
}

// GetHistoryCountOk returns a tuple with the HistoryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetHistoryCountOk() (*int32, bool) {
	if o == nil || IsNil(o.HistoryCount) {
		return nil, false
	}
	return o.HistoryCount, true
}

// HasHistoryCount returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasHistoryCount() bool {
	if o != nil && !IsNil(o.HistoryCount) {
		return true
	}

	return false
}

// SetHistoryCount gets a reference to the given int32 and assigns it to the HistoryCount field.
func (o *PasswordPolicyPasswordSettingsAge) SetHistoryCount(v int32) {
	o.HistoryCount = &v
}

// GetMaxAgeDays returns the MaxAgeDays field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsAge) GetMaxAgeDays() int32 {
	if o == nil || IsNil(o.MaxAgeDays) {
		var ret int32
		return ret
	}
	return *o.MaxAgeDays
}

// GetMaxAgeDaysOk returns a tuple with the MaxAgeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetMaxAgeDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAgeDays) {
		return nil, false
	}
	return o.MaxAgeDays, true
}

// HasMaxAgeDays returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasMaxAgeDays() bool {
	if o != nil && !IsNil(o.MaxAgeDays) {
		return true
	}

	return false
}

// SetMaxAgeDays gets a reference to the given int32 and assigns it to the MaxAgeDays field.
func (o *PasswordPolicyPasswordSettingsAge) SetMaxAgeDays(v int32) {
	o.MaxAgeDays = &v
}

// GetMinAgeMinutes returns the MinAgeMinutes field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsAge) GetMinAgeMinutes() int32 {
	if o == nil || IsNil(o.MinAgeMinutes) {
		var ret int32
		return ret
	}
	return *o.MinAgeMinutes
}

// GetMinAgeMinutesOk returns a tuple with the MinAgeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetMinAgeMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinAgeMinutes) {
		return nil, false
	}
	return o.MinAgeMinutes, true
}

// HasMinAgeMinutes returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasMinAgeMinutes() bool {
	if o != nil && !IsNil(o.MinAgeMinutes) {
		return true
	}

	return false
}

// SetMinAgeMinutes gets a reference to the given int32 and assigns it to the MinAgeMinutes field.
func (o *PasswordPolicyPasswordSettingsAge) SetMinAgeMinutes(v int32) {
	o.MinAgeMinutes = &v
}

func (o PasswordPolicyPasswordSettingsAge) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyPasswordSettingsAge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpireWarnDays) {
		toSerialize["expireWarnDays"] = o.ExpireWarnDays
	}
	if !IsNil(o.HistoryCount) {
		toSerialize["historyCount"] = o.HistoryCount
	}
	if !IsNil(o.MaxAgeDays) {
		toSerialize["maxAgeDays"] = o.MaxAgeDays
	}
	if !IsNil(o.MinAgeMinutes) {
		toSerialize["minAgeMinutes"] = o.MinAgeMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyPasswordSettingsAge) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyPasswordSettingsAge := _PasswordPolicyPasswordSettingsAge{}

	err = json.Unmarshal(data, &varPasswordPolicyPasswordSettingsAge)

	if err != nil {
		return err
	}

	*o = PasswordPolicyPasswordSettingsAge(varPasswordPolicyPasswordSettingsAge)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expireWarnDays")
		delete(additionalProperties, "historyCount")
		delete(additionalProperties, "maxAgeDays")
		delete(additionalProperties, "minAgeMinutes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyPasswordSettingsAge struct {
	value *PasswordPolicyPasswordSettingsAge
	isSet bool
}

func (v NullablePasswordPolicyPasswordSettingsAge) Get() *PasswordPolicyPasswordSettingsAge {
	return v.value
}

func (v *NullablePasswordPolicyPasswordSettingsAge) Set(val *PasswordPolicyPasswordSettingsAge) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyPasswordSettingsAge) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyPasswordSettingsAge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyPasswordSettingsAge(val *PasswordPolicyPasswordSettingsAge) *NullablePasswordPolicyPasswordSettingsAge {
	return &NullablePasswordPolicyPasswordSettingsAge{value: val, isSet: true}
}

func (v NullablePasswordPolicyPasswordSettingsAge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyPasswordSettingsAge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
