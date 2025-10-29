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
	"fmt"
)

// checks if the SecurityEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityEvent{}

// SecurityEvent struct for SecurityEvent
type SecurityEvent struct {
	// The time of the event (UNIX timestamp)
	EventTimestamp       int64                `json:"event_timestamp"`
	Subject              SecurityEventSubject `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEvent SecurityEvent

// NewSecurityEvent instantiates a new SecurityEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEvent(eventTimestamp int64, subject SecurityEventSubject) *SecurityEvent {
	this := SecurityEvent{}
	this.EventTimestamp = eventTimestamp
	this.Subject = subject
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

// GetSubject returns the Subject field value
func (o *SecurityEvent) GetSubject() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *SecurityEvent) GetSubjectOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *SecurityEvent) SetSubject(v SecurityEventSubject) {
	o.Subject = v
}

func (o SecurityEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_timestamp"] = o.EventTimestamp
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_timestamp",
		"subject",
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

	varSecurityEvent := _SecurityEvent{}

	err = json.Unmarshal(data, &varSecurityEvent)

	if err != nil {
		return err
	}

	*o = SecurityEvent(varSecurityEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
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
