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
	"time"
)

// checks if the AutoUpdateSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoUpdateSchedule{}

// AutoUpdateSchedule The schedule of auto-update configured by the admin
type AutoUpdateSchedule struct {
	// The schedule of the update in cron format. The cron settings are limited to only the day of the month or the nth-day-of-the-week configurations. For example, `0 8 ? * 6#3` indicates every third Saturday at 8:00 AM.
	Cron *string `json:"cron,omitempty"`
	// Delay in days
	Delay *int32 `json:"delay,omitempty"`
	// Duration in minutes
	Duration *int32 `json:"duration,omitempty"`
	// Timestamp when the update finished (only for a successful or failed update, not for a cancelled update). Null is returned if the job hasn't finished once yet.
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Timezone of where the scheduled job takes place
	Timezone             *string `json:"timezone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoUpdateSchedule AutoUpdateSchedule

// NewAutoUpdateSchedule instantiates a new AutoUpdateSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoUpdateSchedule() *AutoUpdateSchedule {
	this := AutoUpdateSchedule{}
	return &this
}

// NewAutoUpdateScheduleWithDefaults instantiates a new AutoUpdateSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoUpdateScheduleWithDefaults() *AutoUpdateSchedule {
	this := AutoUpdateSchedule{}
	return &this
}

// GetCron returns the Cron field value if set, zero value otherwise.
func (o *AutoUpdateSchedule) GetCron() string {
	if o == nil || IsNil(o.Cron) {
		var ret string
		return ret
	}
	return *o.Cron
}

// GetCronOk returns a tuple with the Cron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSchedule) GetCronOk() (*string, bool) {
	if o == nil || IsNil(o.Cron) {
		return nil, false
	}
	return o.Cron, true
}

// HasCron returns a boolean if a field has been set.
func (o *AutoUpdateSchedule) HasCron() bool {
	if o != nil && !IsNil(o.Cron) {
		return true
	}

	return false
}

// SetCron gets a reference to the given string and assigns it to the Cron field.
func (o *AutoUpdateSchedule) SetCron(v string) {
	o.Cron = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *AutoUpdateSchedule) GetDelay() int32 {
	if o == nil || IsNil(o.Delay) {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSchedule) GetDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *AutoUpdateSchedule) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *AutoUpdateSchedule) SetDelay(v int32) {
	o.Delay = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AutoUpdateSchedule) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSchedule) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AutoUpdateSchedule) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AutoUpdateSchedule) SetDuration(v int32) {
	o.Duration = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AutoUpdateSchedule) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSchedule) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AutoUpdateSchedule) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AutoUpdateSchedule) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AutoUpdateSchedule) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoUpdateSchedule) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AutoUpdateSchedule) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AutoUpdateSchedule) SetTimezone(v string) {
	o.Timezone = &v
}

func (o AutoUpdateSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoUpdateSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cron) {
		toSerialize["cron"] = o.Cron
	}
	if !IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoUpdateSchedule) UnmarshalJSON(data []byte) (err error) {
	varAutoUpdateSchedule := _AutoUpdateSchedule{}

	err = json.Unmarshal(data, &varAutoUpdateSchedule)

	if err != nil {
		return err
	}

	*o = AutoUpdateSchedule(varAutoUpdateSchedule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cron")
		delete(additionalProperties, "delay")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoUpdateSchedule struct {
	value *AutoUpdateSchedule
	isSet bool
}

func (v NullableAutoUpdateSchedule) Get() *AutoUpdateSchedule {
	return v.value
}

func (v *NullableAutoUpdateSchedule) Set(val *AutoUpdateSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoUpdateSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoUpdateSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoUpdateSchedule(val *AutoUpdateSchedule) *NullableAutoUpdateSchedule {
	return &NullableAutoUpdateSchedule{value: val, isSet: true}
}

func (v NullableAutoUpdateSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoUpdateSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
