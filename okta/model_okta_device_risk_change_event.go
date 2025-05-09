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

// OktaDeviceRiskChangeEvent The device risk level changed
type OktaDeviceRiskChangeEvent struct {
	// Current risk level of the device
	CurrentLevel string `json:"current_level"`
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	// The entity that initiated the event
	InitiatingEntity *string `json:"initiating_entity,omitempty"`
	// Previous risk level of the device
	PreviousLevel string `json:"previous_level"`
	ReasonAdmin *CaepDeviceComplianceChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser *CaepDeviceComplianceChangeEventReasonUser `json:"reason_user,omitempty"`
	Subjects SecurityEventSubject `json:"subjects"`
	AdditionalProperties map[string]interface{}
}

type _OktaDeviceRiskChangeEvent OktaDeviceRiskChangeEvent

// NewOktaDeviceRiskChangeEvent instantiates a new OktaDeviceRiskChangeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaDeviceRiskChangeEvent(currentLevel string, eventTimestamp int64, previousLevel string, subjects SecurityEventSubject) *OktaDeviceRiskChangeEvent {
	this := OktaDeviceRiskChangeEvent{}
	this.CurrentLevel = currentLevel
	this.EventTimestamp = eventTimestamp
	this.PreviousLevel = previousLevel
	this.Subjects = subjects
	return &this
}

// NewOktaDeviceRiskChangeEventWithDefaults instantiates a new OktaDeviceRiskChangeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaDeviceRiskChangeEventWithDefaults() *OktaDeviceRiskChangeEvent {
	this := OktaDeviceRiskChangeEvent{}
	return &this
}

// GetCurrentLevel returns the CurrentLevel field value
func (o *OktaDeviceRiskChangeEvent) GetCurrentLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentLevel
}

// GetCurrentLevelOk returns a tuple with the CurrentLevel field value
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetCurrentLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLevel, true
}

// SetCurrentLevel sets field value
func (o *OktaDeviceRiskChangeEvent) SetCurrentLevel(v string) {
	o.CurrentLevel = v
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *OktaDeviceRiskChangeEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *OktaDeviceRiskChangeEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetInitiatingEntity returns the InitiatingEntity field value if set, zero value otherwise.
func (o *OktaDeviceRiskChangeEvent) GetInitiatingEntity() string {
	if o == nil || o.InitiatingEntity == nil {
		var ret string
		return ret
	}
	return *o.InitiatingEntity
}

// GetInitiatingEntityOk returns a tuple with the InitiatingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetInitiatingEntityOk() (*string, bool) {
	if o == nil || o.InitiatingEntity == nil {
		return nil, false
	}
	return o.InitiatingEntity, true
}

// HasInitiatingEntity returns a boolean if a field has been set.
func (o *OktaDeviceRiskChangeEvent) HasInitiatingEntity() bool {
	if o != nil && o.InitiatingEntity != nil {
		return true
	}

	return false
}

// SetInitiatingEntity gets a reference to the given string and assigns it to the InitiatingEntity field.
func (o *OktaDeviceRiskChangeEvent) SetInitiatingEntity(v string) {
	o.InitiatingEntity = &v
}

// GetPreviousLevel returns the PreviousLevel field value
func (o *OktaDeviceRiskChangeEvent) GetPreviousLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousLevel
}

// GetPreviousLevelOk returns a tuple with the PreviousLevel field value
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetPreviousLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousLevel, true
}

// SetPreviousLevel sets field value
func (o *OktaDeviceRiskChangeEvent) SetPreviousLevel(v string) {
	o.PreviousLevel = v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *OktaDeviceRiskChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin {
	if o == nil || o.ReasonAdmin == nil {
		var ret CaepDeviceComplianceChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool) {
	if o == nil || o.ReasonAdmin == nil {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *OktaDeviceRiskChangeEvent) HasReasonAdmin() bool {
	if o != nil && o.ReasonAdmin != nil {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepDeviceComplianceChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *OktaDeviceRiskChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *OktaDeviceRiskChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser {
	if o == nil || o.ReasonUser == nil {
		var ret CaepDeviceComplianceChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool) {
	if o == nil || o.ReasonUser == nil {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *OktaDeviceRiskChangeEvent) HasReasonUser() bool {
	if o != nil && o.ReasonUser != nil {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepDeviceComplianceChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *OktaDeviceRiskChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubjects returns the Subjects field value
func (o *OktaDeviceRiskChangeEvent) GetSubjects() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *OktaDeviceRiskChangeEvent) GetSubjectsOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value
func (o *OktaDeviceRiskChangeEvent) SetSubjects(v SecurityEventSubject) {
	o.Subjects = v
}

func (o OktaDeviceRiskChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_level"] = o.CurrentLevel
	}
	if true {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.InitiatingEntity != nil {
		toSerialize["initiating_entity"] = o.InitiatingEntity
	}
	if true {
		toSerialize["previous_level"] = o.PreviousLevel
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

func (o *OktaDeviceRiskChangeEvent) UnmarshalJSON(bytes []byte) (err error) {
	varOktaDeviceRiskChangeEvent := _OktaDeviceRiskChangeEvent{}

	err = json.Unmarshal(bytes, &varOktaDeviceRiskChangeEvent)
	if err == nil {
		*o = OktaDeviceRiskChangeEvent(varOktaDeviceRiskChangeEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "current_level")
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "initiating_entity")
		delete(additionalProperties, "previous_level")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subjects")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaDeviceRiskChangeEvent struct {
	value *OktaDeviceRiskChangeEvent
	isSet bool
}

func (v NullableOktaDeviceRiskChangeEvent) Get() *OktaDeviceRiskChangeEvent {
	return v.value
}

func (v *NullableOktaDeviceRiskChangeEvent) Set(val *OktaDeviceRiskChangeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaDeviceRiskChangeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaDeviceRiskChangeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaDeviceRiskChangeEvent(val *OktaDeviceRiskChangeEvent) *NullableOktaDeviceRiskChangeEvent {
	return &NullableOktaDeviceRiskChangeEvent{value: val, isSet: true}
}

func (v NullableOktaDeviceRiskChangeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaDeviceRiskChangeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

