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
	"time"
)

// RiskEvent struct for RiskEvent
type RiskEvent struct {
	// Timestamp at which the event expires (expressed as a UTC time zone using ISO 8601 format: yyyy-MM-dd`T`HH:mm:ss.SSS`Z`). If this optional field is not included, Okta automatically expires the event 24 hours after the event is consumed.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// List of Risk Event Subjects
	Subjects []RiskEventSubject `json:"subjects"`
	// Timestamp of when the event is produced (expressed as a UTC time zone using ISO 8601 format: yyyy-MM-dd`T`HH:mm:ss.SSS`Z`)
	Timestamp *time.Time `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskEvent RiskEvent

// NewRiskEvent instantiates a new RiskEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvent(subjects []RiskEventSubject) *RiskEvent {
	this := RiskEvent{}
	this.Subjects = subjects
	return &this
}

// NewRiskEventWithDefaults instantiates a new RiskEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEventWithDefaults() *RiskEvent {
	this := RiskEvent{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *RiskEvent) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvent) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *RiskEvent) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *RiskEvent) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetSubjects returns the Subjects field value
func (o *RiskEvent) GetSubjects() []RiskEventSubject {
	if o == nil {
		var ret []RiskEventSubject
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *RiskEvent) GetSubjectsOk() ([]RiskEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subjects, true
}

// SetSubjects sets field value
func (o *RiskEvent) SetSubjects(v []RiskEventSubject) {
	o.Subjects = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *RiskEvent) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RiskEvent) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *RiskEvent) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o RiskEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if true {
		toSerialize["subjects"] = o.Subjects
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskEvent) UnmarshalJSON(bytes []byte) (err error) {
	varRiskEvent := _RiskEvent{}

	err = json.Unmarshal(bytes, &varRiskEvent)
	if err == nil {
		*o = RiskEvent(varRiskEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "subjects")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskEvent struct {
	value *RiskEvent
	isSet bool
}

func (v NullableRiskEvent) Get() *RiskEvent {
	return v.value
}

func (v *NullableRiskEvent) Set(val *RiskEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvent(val *RiskEvent) *NullableRiskEvent {
	return &NullableRiskEvent{value: val, isSet: true}
}

func (v NullableRiskEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

