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

// SecurityEvent struct for SecurityEvent
type SecurityEvent struct {
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	Subjects SecurityEventSubject `json:"subjects"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEvent SecurityEvent

// NewSecurityEvent instantiates a new SecurityEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEvent(eventTimestamp int64, subjects SecurityEventSubject) *SecurityEvent {
	this := SecurityEvent{}
	this.EventTimestamp = eventTimestamp
	this.Subjects = subjects
	return &this
}

// NewSecurityEventWithDefaults instantiates a new SecurityEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventWithDefaults() *SecurityEvent {
	this := SecurityEvent{}
	return &this
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *SecurityEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *SecurityEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *SecurityEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetSubjects returns the Subjects field value
func (o *SecurityEvent) GetSubjects() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *SecurityEvent) GetSubjectsOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value
func (o *SecurityEvent) SetSubjects(v SecurityEventSubject) {
	o.Subjects = v
}

func (o SecurityEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if true {
		toSerialize["subjects"] = o.Subjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEvent) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEvent := _SecurityEvent{}

	err = json.Unmarshal(bytes, &varSecurityEvent)
	if err == nil {
		*o = SecurityEvent(varSecurityEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "subjects")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEvent struct {
	value *SecurityEvent
	isSet bool
}

func (v NullableSecurityEvent) Get() *SecurityEvent {
	return v.value
}

func (v *NullableSecurityEvent) Set(val *SecurityEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEvent(val *SecurityEvent) *NullableSecurityEvent {
	return &NullableSecurityEvent{value: val, isSet: true}
}

func (v NullableSecurityEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

