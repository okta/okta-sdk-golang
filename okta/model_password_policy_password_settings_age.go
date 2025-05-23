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

// PasswordPolicyPasswordSettingsAge struct for PasswordPolicyPasswordSettingsAge
type PasswordPolicyPasswordSettingsAge struct {
	ExpireWarnDays *int32 `json:"expireWarnDays,omitempty"`
	HistoryCount *int32 `json:"historyCount,omitempty"`
	MaxAgeDays *int32 `json:"maxAgeDays,omitempty"`
	MinAgeMinutes *int32 `json:"minAgeMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyPasswordSettingsAge PasswordPolicyPasswordSettingsAge

// NewPasswordPolicyPasswordSettingsAge instantiates a new PasswordPolicyPasswordSettingsAge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyPasswordSettingsAge() *PasswordPolicyPasswordSettingsAge {
	this := PasswordPolicyPasswordSettingsAge{}
	return &this
}

// NewPasswordPolicyPasswordSettingsAgeWithDefaults instantiates a new PasswordPolicyPasswordSettingsAge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsAgeWithDefaults() *PasswordPolicyPasswordSettingsAge {
	this := PasswordPolicyPasswordSettingsAge{}
	return &this
}

// GetExpireWarnDays returns the ExpireWarnDays field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsAge) GetExpireWarnDays() int32 {
	if o == nil || o.ExpireWarnDays == nil {
		var ret int32
		return ret
	}
	return *o.ExpireWarnDays
}

// GetExpireWarnDaysOk returns a tuple with the ExpireWarnDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetExpireWarnDaysOk() (*int32, bool) {
	if o == nil || o.ExpireWarnDays == nil {
		return nil, false
	}
	return o.ExpireWarnDays, true
}

// HasExpireWarnDays returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasExpireWarnDays() bool {
	if o != nil && o.ExpireWarnDays != nil {
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
	if o == nil || o.HistoryCount == nil {
		var ret int32
		return ret
	}
	return *o.HistoryCount
}

// GetHistoryCountOk returns a tuple with the HistoryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetHistoryCountOk() (*int32, bool) {
	if o == nil || o.HistoryCount == nil {
		return nil, false
	}
	return o.HistoryCount, true
}

// HasHistoryCount returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasHistoryCount() bool {
	if o != nil && o.HistoryCount != nil {
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
	if o == nil || o.MaxAgeDays == nil {
		var ret int32
		return ret
	}
	return *o.MaxAgeDays
}

// GetMaxAgeDaysOk returns a tuple with the MaxAgeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetMaxAgeDaysOk() (*int32, bool) {
	if o == nil || o.MaxAgeDays == nil {
		return nil, false
	}
	return o.MaxAgeDays, true
}

// HasMaxAgeDays returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasMaxAgeDays() bool {
	if o != nil && o.MaxAgeDays != nil {
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
	if o == nil || o.MinAgeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.MinAgeMinutes
}

// GetMinAgeMinutesOk returns a tuple with the MinAgeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsAge) GetMinAgeMinutesOk() (*int32, bool) {
	if o == nil || o.MinAgeMinutes == nil {
		return nil, false
	}
	return o.MinAgeMinutes, true
}

// HasMinAgeMinutes returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsAge) HasMinAgeMinutes() bool {
	if o != nil && o.MinAgeMinutes != nil {
		return true
	}

	return false
}

// SetMinAgeMinutes gets a reference to the given int32 and assigns it to the MinAgeMinutes field.
func (o *PasswordPolicyPasswordSettingsAge) SetMinAgeMinutes(v int32) {
	o.MinAgeMinutes = &v
}

func (o PasswordPolicyPasswordSettingsAge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpireWarnDays != nil {
		toSerialize["expireWarnDays"] = o.ExpireWarnDays
	}
	if o.HistoryCount != nil {
		toSerialize["historyCount"] = o.HistoryCount
	}
	if o.MaxAgeDays != nil {
		toSerialize["maxAgeDays"] = o.MaxAgeDays
	}
	if o.MinAgeMinutes != nil {
		toSerialize["minAgeMinutes"] = o.MinAgeMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyPasswordSettingsAge) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyPasswordSettingsAge := _PasswordPolicyPasswordSettingsAge{}

	err = json.Unmarshal(bytes, &varPasswordPolicyPasswordSettingsAge)
	if err == nil {
		*o = PasswordPolicyPasswordSettingsAge(varPasswordPolicyPasswordSettingsAge)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expireWarnDays")
		delete(additionalProperties, "historyCount")
		delete(additionalProperties, "maxAgeDays")
		delete(additionalProperties, "minAgeMinutes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

