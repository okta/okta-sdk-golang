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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RiskEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvent{}

// RiskEvent struct for RiskEvent
type RiskEvent struct {
	// Timestamp at which the event expires (expressed as a UTC time zone using ISO 8601 format: yyyy-MM-dd`T`HH:mm:ss.SSS`Z`). If this optional field isn't included, Okta automatically expires the event 24 hours after the event is consumed.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// List of risk event subjects
	Subjects []RiskEventSubject `json:"subjects"`
	// Timestamp of when the event is produced (expressed as a UTC time zone using ISO 8601 format: yyyy-MM-dd`T`HH:mm:ss.SSS`Z`)
	Timestamp            *time.Time `json:"timestamp,omitempty"`
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
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvent) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *RiskEvent) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
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
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RiskEvent) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *RiskEvent) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o RiskEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["subjects"] = o.Subjects
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subjects",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRiskEvent := _RiskEvent{}

	err = json.Unmarshal(data, &varRiskEvent)

	if err != nil {
		return err
	}

	*o = RiskEvent(varRiskEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "subjects")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
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
