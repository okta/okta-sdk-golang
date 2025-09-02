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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// CaepEvent struct for CaepEvent
type CaepEvent struct {
	// The time of the event (UNIX timestamp)
	EventTimestamp *int64 `json:"event_timestamp,omitempty"`
	ReasonAdmin *CaepCredentialChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser *CaepCredentialChangeEventReasonUser `json:"reason_user,omitempty"`
	Subject *SecurityEventSubject `json:"subject,omitempty"`
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
	if o == nil || o.EventTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil || o.EventTimestamp == nil {
		return nil, false
	}
	return o.EventTimestamp, true
}

// HasEventTimestamp returns a boolean if a field has been set.
func (o *CaepEvent) HasEventTimestamp() bool {
	if o != nil && o.EventTimestamp != nil {
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
	if o == nil || o.ReasonAdmin == nil {
		var ret CaepCredentialChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetReasonAdminOk() (*CaepCredentialChangeEventReasonAdmin, bool) {
	if o == nil || o.ReasonAdmin == nil {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *CaepEvent) HasReasonAdmin() bool {
	if o != nil && o.ReasonAdmin != nil {
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
	if o == nil || o.ReasonUser == nil {
		var ret CaepCredentialChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetReasonUserOk() (*CaepCredentialChangeEventReasonUser, bool) {
	if o == nil || o.ReasonUser == nil {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *CaepEvent) HasReasonUser() bool {
	if o != nil && o.ReasonUser != nil {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepCredentialChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *CaepEvent) SetReasonUser(v CaepCredentialChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CaepEvent) GetSubject() SecurityEventSubject {
	if o == nil || o.Subject == nil {
		var ret SecurityEventSubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepEvent) GetSubjectOk() (*SecurityEventSubject, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CaepEvent) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given SecurityEventSubject and assigns it to the Subject field.
func (o *CaepEvent) SetSubject(v SecurityEventSubject) {
	o.Subject = &v
}

func (o CaepEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventTimestamp != nil {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.ReasonAdmin != nil {
		toSerialize["reason_admin"] = o.ReasonAdmin
	}
	if o.ReasonUser != nil {
		toSerialize["reason_user"] = o.ReasonUser
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CaepEvent) UnmarshalJSON(bytes []byte) (err error) {
	varCaepEvent := _CaepEvent{}

	err = json.Unmarshal(bytes, &varCaepEvent)
	if err == nil {
		*o = CaepEvent(varCaepEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

