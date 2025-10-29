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

// checks if the CaepDeviceComplianceChangeEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaepDeviceComplianceChangeEvent{}

// CaepDeviceComplianceChangeEvent The subject's device compliance was revoked
type CaepDeviceComplianceChangeEvent struct {
	// Current device compliance status
	CurrentStatus string `json:"current_status"`
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	// The entity that initiated the event
	InitiatingEntity *string `json:"initiating_entity,omitempty"`
	// Previous device compliance status
	PreviousStatus       string                                      `json:"previous_status"`
	ReasonAdmin          *CaepDeviceComplianceChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser           *CaepDeviceComplianceChangeEventReasonUser  `json:"reason_user,omitempty"`
	Subject              SecurityEventSubject                        `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _CaepDeviceComplianceChangeEvent CaepDeviceComplianceChangeEvent

// NewCaepDeviceComplianceChangeEvent instantiates a new CaepDeviceComplianceChangeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepDeviceComplianceChangeEvent(currentStatus string, eventTimestamp int64, previousStatus string, subject SecurityEventSubject) *CaepDeviceComplianceChangeEvent {
	this := CaepDeviceComplianceChangeEvent{}
	this.CurrentStatus = currentStatus
	this.EventTimestamp = eventTimestamp
	this.PreviousStatus = previousStatus
	this.Subject = subject
	return &this
}

// NewCaepDeviceComplianceChangeEventWithDefaults instantiates a new CaepDeviceComplianceChangeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepDeviceComplianceChangeEventWithDefaults() *CaepDeviceComplianceChangeEvent {
	this := CaepDeviceComplianceChangeEvent{}
	return &this
}

// GetCurrentStatus returns the CurrentStatus field value
func (o *CaepDeviceComplianceChangeEvent) GetCurrentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetCurrentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStatus, true
}

// SetCurrentStatus sets field value
func (o *CaepDeviceComplianceChangeEvent) SetCurrentStatus(v string) {
	o.CurrentStatus = v
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *CaepDeviceComplianceChangeEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *CaepDeviceComplianceChangeEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetInitiatingEntity returns the InitiatingEntity field value if set, zero value otherwise.
func (o *CaepDeviceComplianceChangeEvent) GetInitiatingEntity() string {
	if o == nil || IsNil(o.InitiatingEntity) {
		var ret string
		return ret
	}
	return *o.InitiatingEntity
}

// GetInitiatingEntityOk returns a tuple with the InitiatingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetInitiatingEntityOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatingEntity) {
		return nil, false
	}
	return o.InitiatingEntity, true
}

// HasInitiatingEntity returns a boolean if a field has been set.
func (o *CaepDeviceComplianceChangeEvent) HasInitiatingEntity() bool {
	if o != nil && !IsNil(o.InitiatingEntity) {
		return true
	}

	return false
}

// SetInitiatingEntity gets a reference to the given string and assigns it to the InitiatingEntity field.
func (o *CaepDeviceComplianceChangeEvent) SetInitiatingEntity(v string) {
	o.InitiatingEntity = &v
}

// GetPreviousStatus returns the PreviousStatus field value
func (o *CaepDeviceComplianceChangeEvent) GetPreviousStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousStatus
}

// GetPreviousStatusOk returns a tuple with the PreviousStatus field value
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetPreviousStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousStatus, true
}

// SetPreviousStatus sets field value
func (o *CaepDeviceComplianceChangeEvent) SetPreviousStatus(v string) {
	o.PreviousStatus = v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *CaepDeviceComplianceChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin {
	if o == nil || IsNil(o.ReasonAdmin) {
		var ret CaepDeviceComplianceChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool) {
	if o == nil || IsNil(o.ReasonAdmin) {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *CaepDeviceComplianceChangeEvent) HasReasonAdmin() bool {
	if o != nil && !IsNil(o.ReasonAdmin) {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepDeviceComplianceChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *CaepDeviceComplianceChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *CaepDeviceComplianceChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser {
	if o == nil || IsNil(o.ReasonUser) {
		var ret CaepDeviceComplianceChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool) {
	if o == nil || IsNil(o.ReasonUser) {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *CaepDeviceComplianceChangeEvent) HasReasonUser() bool {
	if o != nil && !IsNil(o.ReasonUser) {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepDeviceComplianceChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *CaepDeviceComplianceChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubject returns the Subject field value
func (o *CaepDeviceComplianceChangeEvent) GetSubject() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEvent) GetSubjectOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CaepDeviceComplianceChangeEvent) SetSubject(v SecurityEventSubject) {
	o.Subject = v
}

func (o CaepDeviceComplianceChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaepDeviceComplianceChangeEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_status"] = o.CurrentStatus
	toSerialize["event_timestamp"] = o.EventTimestamp
	if !IsNil(o.InitiatingEntity) {
		toSerialize["initiating_entity"] = o.InitiatingEntity
	}
	toSerialize["previous_status"] = o.PreviousStatus
	if !IsNil(o.ReasonAdmin) {
		toSerialize["reason_admin"] = o.ReasonAdmin
	}
	if !IsNil(o.ReasonUser) {
		toSerialize["reason_user"] = o.ReasonUser
	}
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaepDeviceComplianceChangeEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_status",
		"event_timestamp",
		"previous_status",
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

	varCaepDeviceComplianceChangeEvent := _CaepDeviceComplianceChangeEvent{}

	err = json.Unmarshal(data, &varCaepDeviceComplianceChangeEvent)

	if err != nil {
		return err
	}

	*o = CaepDeviceComplianceChangeEvent(varCaepDeviceComplianceChangeEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current_status")
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "initiating_entity")
		delete(additionalProperties, "previous_status")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaepDeviceComplianceChangeEvent struct {
	value *CaepDeviceComplianceChangeEvent
	isSet bool
}

func (v NullableCaepDeviceComplianceChangeEvent) Get() *CaepDeviceComplianceChangeEvent {
	return v.value
}

func (v *NullableCaepDeviceComplianceChangeEvent) Set(val *CaepDeviceComplianceChangeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepDeviceComplianceChangeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepDeviceComplianceChangeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepDeviceComplianceChangeEvent(val *CaepDeviceComplianceChangeEvent) *NullableCaepDeviceComplianceChangeEvent {
	return &NullableCaepDeviceComplianceChangeEvent{value: val, isSet: true}
}

func (v NullableCaepDeviceComplianceChangeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepDeviceComplianceChangeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
