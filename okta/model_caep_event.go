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

// checks if the CaepEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaepEvent{}

// CaepEvent struct for CaepEvent
type CaepEvent struct {
	// The time of the event (UNIX timestamp)
	EventTimestamp       *int64                                `json:"event_timestamp,omitempty"`
	ReasonAdmin          *CaepCredentialChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser           *CaepCredentialChangeEventReasonUser  `json:"reason_user,omitempty"`
	Subject              *SsfTransmitterSecurityEventSubject   `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaepEvent CaepEvent

// NewCaepEvent instantiates a new CaepEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepEvent() *CaepEvent {
	this := CaepEvent{}
	return &this
}

// NewCaepEventWithDefaults instantiates a new CaepEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepEventWithDefaults() *CaepEvent {
	this := CaepEvent{}
	return &this
}

// GetEventTimestamp returns the EventTimestamp field value if set, zero value otherwise.
func (o *CaepEvent) GetEventTimestamp() int64 {
	if o == nil || IsNil(o.EventTimestamp) {
		var ret int64
		return ret
	}
	return *o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.EventTimestamp) {
		return nil, false
	}
	return o.EventTimestamp, true
}

// HasEventTimestamp returns a boolean if a field has been set.
func (o *CaepEvent) HasEventTimestamp() bool {
	if o != nil && !IsNil(o.EventTimestamp) {
		return true
	}

	return false
}

// SetEventTimestamp gets a reference to the given int64 and assigns it to the EventTimestamp field.
func (o *CaepEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = &v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *CaepEvent) GetReasonAdmin() CaepCredentialChangeEventReasonAdmin {
	if o == nil || IsNil(o.ReasonAdmin) {
		var ret CaepCredentialChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetReasonAdminOk() (*CaepCredentialChangeEventReasonAdmin, bool) {
	if o == nil || IsNil(o.ReasonAdmin) {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *CaepEvent) HasReasonAdmin() bool {
	if o != nil && !IsNil(o.ReasonAdmin) {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepCredentialChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *CaepEvent) SetReasonAdmin(v CaepCredentialChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *CaepEvent) GetReasonUser() CaepCredentialChangeEventReasonUser {
	if o == nil || IsNil(o.ReasonUser) {
		var ret CaepCredentialChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetReasonUserOk() (*CaepCredentialChangeEventReasonUser, bool) {
	if o == nil || IsNil(o.ReasonUser) {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *CaepEvent) HasReasonUser() bool {
	if o != nil && !IsNil(o.ReasonUser) {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepCredentialChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *CaepEvent) SetReasonUser(v CaepCredentialChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CaepEvent) GetSubject() SsfTransmitterSecurityEventSubject {
	if o == nil || IsNil(o.Subject) {
		var ret SsfTransmitterSecurityEventSubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetSubjectOk() (*SsfTransmitterSecurityEventSubject, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CaepEvent) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given SsfTransmitterSecurityEventSubject and assigns it to the Subject field.
func (o *CaepEvent) SetSubject(v SsfTransmitterSecurityEventSubject) {
	o.Subject = &v
}

func (o CaepEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaepEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventTimestamp) {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if !IsNil(o.ReasonAdmin) {
		toSerialize["reason_admin"] = o.ReasonAdmin
	}
	if !IsNil(o.ReasonUser) {
		toSerialize["reason_user"] = o.ReasonUser
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaepEvent) UnmarshalJSON(data []byte) (err error) {
	varCaepEvent := _CaepEvent{}

	err = json.Unmarshal(data, &varCaepEvent)

	if err != nil {
		return err
	}

	*o = CaepEvent(varCaepEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaepEvent struct {
	value *CaepEvent
	isSet bool
}

func (v NullableCaepEvent) Get() *CaepEvent {
	return v.value
}

func (v *NullableCaepEvent) Set(val *CaepEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepEvent(val *CaepEvent) *NullableCaepEvent {
	return &NullableCaepEvent{value: val, isSet: true}
}

func (v NullableCaepEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
