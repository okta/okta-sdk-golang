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

// OktaIpChangeEvent IP changed for the subject's session
type OktaIpChangeEvent struct {
	// Current IP address of the subject
	CurrentIpAddress string `json:"current_ip_address"`
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	// The entity that initiated the event
	InitiatingEntity *string `json:"initiating_entity,omitempty"`
	// Previous IP address of the subject
	PreviousIpAddress string `json:"previous_ip_address"`
	ReasonAdmin *CaepDeviceComplianceChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser *CaepDeviceComplianceChangeEventReasonUser `json:"reason_user,omitempty"`
	Subjects SecurityEventSubject `json:"subjects"`
	AdditionalProperties map[string]interface{}
}

type _OktaIpChangeEvent OktaIpChangeEvent

// NewOktaIpChangeEvent instantiates a new OktaIpChangeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaIpChangeEvent(currentIpAddress string, eventTimestamp int64, previousIpAddress string, subjects SecurityEventSubject) *OktaIpChangeEvent {
	this := OktaIpChangeEvent{}
	this.CurrentIpAddress = currentIpAddress
	this.EventTimestamp = eventTimestamp
	this.PreviousIpAddress = previousIpAddress
	this.Subjects = subjects
	return &this
}

// NewOktaIpChangeEventWithDefaults instantiates a new OktaIpChangeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaIpChangeEventWithDefaults() *OktaIpChangeEvent {
	this := OktaIpChangeEvent{}
	return &this
}

// GetCurrentIpAddress returns the CurrentIpAddress field value
func (o *OktaIpChangeEvent) GetCurrentIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentIpAddress
}

// GetCurrentIpAddressOk returns a tuple with the CurrentIpAddress field value
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetCurrentIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentIpAddress, true
}

// SetCurrentIpAddress sets field value
func (o *OktaIpChangeEvent) SetCurrentIpAddress(v string) {
	o.CurrentIpAddress = v
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *OktaIpChangeEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *OktaIpChangeEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetInitiatingEntity returns the InitiatingEntity field value if set, zero value otherwise.
func (o *OktaIpChangeEvent) GetInitiatingEntity() string {
	if o == nil || o.InitiatingEntity == nil {
		var ret string
		return ret
	}
	return *o.InitiatingEntity
}

// GetInitiatingEntityOk returns a tuple with the InitiatingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetInitiatingEntityOk() (*string, bool) {
	if o == nil || o.InitiatingEntity == nil {
		return nil, false
	}
	return o.InitiatingEntity, true
}

// HasInitiatingEntity returns a boolean if a field has been set.
func (o *OktaIpChangeEvent) HasInitiatingEntity() bool {
	if o != nil && o.InitiatingEntity != nil {
		return true
	}

	return false
}

// SetInitiatingEntity gets a reference to the given string and assigns it to the InitiatingEntity field.
func (o *OktaIpChangeEvent) SetInitiatingEntity(v string) {
	o.InitiatingEntity = &v
}

// GetPreviousIpAddress returns the PreviousIpAddress field value
func (o *OktaIpChangeEvent) GetPreviousIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousIpAddress
}

// GetPreviousIpAddressOk returns a tuple with the PreviousIpAddress field value
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetPreviousIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousIpAddress, true
}

// SetPreviousIpAddress sets field value
func (o *OktaIpChangeEvent) SetPreviousIpAddress(v string) {
	o.PreviousIpAddress = v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *OktaIpChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin {
	if o == nil || o.ReasonAdmin == nil {
		var ret CaepDeviceComplianceChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool) {
	if o == nil || o.ReasonAdmin == nil {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *OktaIpChangeEvent) HasReasonAdmin() bool {
	if o != nil && o.ReasonAdmin != nil {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepDeviceComplianceChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *OktaIpChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *OktaIpChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser {
	if o == nil || o.ReasonUser == nil {
		var ret CaepDeviceComplianceChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool) {
	if o == nil || o.ReasonUser == nil {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *OktaIpChangeEvent) HasReasonUser() bool {
	if o != nil && o.ReasonUser != nil {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepDeviceComplianceChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *OktaIpChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubjects returns the Subjects field value
func (o *OktaIpChangeEvent) GetSubjects() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *OktaIpChangeEvent) GetSubjectsOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value
func (o *OktaIpChangeEvent) SetSubjects(v SecurityEventSubject) {
	o.Subjects = v
}

func (o OktaIpChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_ip_address"] = o.CurrentIpAddress
	}
	if true {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.InitiatingEntity != nil {
		toSerialize["initiating_entity"] = o.InitiatingEntity
	}
	if true {
		toSerialize["previous_ip_address"] = o.PreviousIpAddress
	}
	if o.ReasonAdmin != nil {
		toSerialize["reason_admin"] = o.ReasonAdmin
	}
	if o.ReasonUser != nil {
		toSerialize["reason_user"] = o.ReasonUser
	}
	if true {
		toSerialize["subjects"] = o.Subjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaIpChangeEvent) UnmarshalJSON(bytes []byte) (err error) {
	varOktaIpChangeEvent := _OktaIpChangeEvent{}

	err = json.Unmarshal(bytes, &varOktaIpChangeEvent)
	if err == nil {
		*o = OktaIpChangeEvent(varOktaIpChangeEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "current_ip_address")
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "initiating_entity")
		delete(additionalProperties, "previous_ip_address")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subjects")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaIpChangeEvent struct {
	value *OktaIpChangeEvent
	isSet bool
}

func (v NullableOktaIpChangeEvent) Get() *OktaIpChangeEvent {
	return v.value
}

func (v *NullableOktaIpChangeEvent) Set(val *OktaIpChangeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaIpChangeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaIpChangeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaIpChangeEvent(val *OktaIpChangeEvent) *NullableOktaIpChangeEvent {
	return &NullableOktaIpChangeEvent{value: val, isSet: true}
}

func (v NullableOktaIpChangeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaIpChangeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

