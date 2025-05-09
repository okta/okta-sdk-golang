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

// RiscIdentifierChangedEvent The subject's identifier has changed, which is either an email address or a phone number change
type RiscIdentifierChangedEvent struct {
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	// The new identifier value
	NewValue *string `json:"new-value,omitempty"`
	Subjects SecurityEventSubject `json:"subjects"`
	AdditionalProperties map[string]interface{}
}

type _RiscIdentifierChangedEvent RiscIdentifierChangedEvent

// NewRiscIdentifierChangedEvent instantiates a new RiscIdentifierChangedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiscIdentifierChangedEvent(eventTimestamp int64, subjects SecurityEventSubject) *RiscIdentifierChangedEvent {
	this := RiscIdentifierChangedEvent{}
	this.EventTimestamp = eventTimestamp
	this.Subjects = subjects
	return &this
}

// NewRiscIdentifierChangedEventWithDefaults instantiates a new RiscIdentifierChangedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiscIdentifierChangedEventWithDefaults() *RiscIdentifierChangedEvent {
	this := RiscIdentifierChangedEvent{}
	return &this
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *RiscIdentifierChangedEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *RiscIdentifierChangedEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *RiscIdentifierChangedEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *RiscIdentifierChangedEvent) GetNewValue() string {
	if o == nil || o.NewValue == nil {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiscIdentifierChangedEvent) GetNewValueOk() (*string, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *RiscIdentifierChangedEvent) HasNewValue() bool {
	if o != nil && o.NewValue != nil {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *RiscIdentifierChangedEvent) SetNewValue(v string) {
	o.NewValue = &v
}

// GetSubjects returns the Subjects field value
func (o *RiscIdentifierChangedEvent) GetSubjects() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *RiscIdentifierChangedEvent) GetSubjectsOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value
func (o *RiscIdentifierChangedEvent) SetSubjects(v SecurityEventSubject) {
	o.Subjects = v
}

func (o RiscIdentifierChangedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.NewValue != nil {
		toSerialize["new-value"] = o.NewValue
	}
	if true {
		toSerialize["subjects"] = o.Subjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiscIdentifierChangedEvent) UnmarshalJSON(bytes []byte) (err error) {
	varRiscIdentifierChangedEvent := _RiscIdentifierChangedEvent{}

	err = json.Unmarshal(bytes, &varRiscIdentifierChangedEvent)
	if err == nil {
		*o = RiscIdentifierChangedEvent(varRiscIdentifierChangedEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "new-value")
		delete(additionalProperties, "subjects")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiscIdentifierChangedEvent struct {
	value *RiscIdentifierChangedEvent
	isSet bool
}

func (v NullableRiscIdentifierChangedEvent) Get() *RiscIdentifierChangedEvent {
	return v.value
}

func (v *NullableRiscIdentifierChangedEvent) Set(val *RiscIdentifierChangedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRiscIdentifierChangedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRiscIdentifierChangedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiscIdentifierChangedEvent(val *RiscIdentifierChangedEvent) *NullableRiscIdentifierChangedEvent {
	return &NullableRiscIdentifierChangedEvent{value: val, isSet: true}
}

func (v NullableRiscIdentifierChangedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiscIdentifierChangedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

