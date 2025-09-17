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
)

// checks if the OktaUserRiskChangeEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaUserRiskChangeEvent{}

// OktaUserRiskChangeEvent The user risk level changed
type OktaUserRiskChangeEvent struct {
	// Current risk level of the user
	CurrentLevel string `json:"current_level"`
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	// The entity that initiated the event
	InitiatingEntity *string `json:"initiating_entity,omitempty"`
	// Previous risk level of the user
	PreviousLevel        string                                      `json:"previous_level"`
	ReasonAdmin          *CaepDeviceComplianceChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser           *CaepDeviceComplianceChangeEventReasonUser  `json:"reason_user,omitempty"`
	Subject              SecurityEventSubject                        `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _OktaUserRiskChangeEvent OktaUserRiskChangeEvent

// NewOktaUserRiskChangeEvent instantiates a new OktaUserRiskChangeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaUserRiskChangeEvent(currentLevel string, eventTimestamp int64, previousLevel string, subject SecurityEventSubject) *OktaUserRiskChangeEvent {
	this := OktaUserRiskChangeEvent{}
	this.CurrentLevel = currentLevel
	this.EventTimestamp = eventTimestamp
	this.PreviousLevel = previousLevel
	this.Subject = subject
	return &this
}

// NewOktaUserRiskChangeEventWithDefaults instantiates a new OktaUserRiskChangeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaUserRiskChangeEventWithDefaults() *OktaUserRiskChangeEvent {
	this := OktaUserRiskChangeEvent{}
	return &this
}

// GetCurrentLevel returns the CurrentLevel field value
func (o *OktaUserRiskChangeEvent) GetCurrentLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentLevel
}

// GetCurrentLevelOk returns a tuple with the CurrentLevel field value
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetCurrentLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLevel, true
}

// SetCurrentLevel sets field value
func (o *OktaUserRiskChangeEvent) SetCurrentLevel(v string) {
	o.CurrentLevel = v
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *OktaUserRiskChangeEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *OktaUserRiskChangeEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetInitiatingEntity returns the InitiatingEntity field value if set, zero value otherwise.
func (o *OktaUserRiskChangeEvent) GetInitiatingEntity() string {
	if o == nil || IsNil(o.InitiatingEntity) {
		var ret string
		return ret
	}
	return *o.InitiatingEntity
}

// GetInitiatingEntityOk returns a tuple with the InitiatingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetInitiatingEntityOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatingEntity) {
		return nil, false
	}
	return o.InitiatingEntity, true
}

// HasInitiatingEntity returns a boolean if a field has been set.
func (o *OktaUserRiskChangeEvent) HasInitiatingEntity() bool {
	if o != nil && !IsNil(o.InitiatingEntity) {
		return true
	}

	return false
}

// SetInitiatingEntity gets a reference to the given string and assigns it to the InitiatingEntity field.
func (o *OktaUserRiskChangeEvent) SetInitiatingEntity(v string) {
	o.InitiatingEntity = &v
}

// GetPreviousLevel returns the PreviousLevel field value
func (o *OktaUserRiskChangeEvent) GetPreviousLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousLevel
}

// GetPreviousLevelOk returns a tuple with the PreviousLevel field value
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetPreviousLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousLevel, true
}

// SetPreviousLevel sets field value
func (o *OktaUserRiskChangeEvent) SetPreviousLevel(v string) {
	o.PreviousLevel = v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *OktaUserRiskChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin {
	if o == nil || IsNil(o.ReasonAdmin) {
		var ret CaepDeviceComplianceChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool) {
	if o == nil || IsNil(o.ReasonAdmin) {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *OktaUserRiskChangeEvent) HasReasonAdmin() bool {
	if o != nil && !IsNil(o.ReasonAdmin) {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepDeviceComplianceChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *OktaUserRiskChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *OktaUserRiskChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser {
	if o == nil || IsNil(o.ReasonUser) {
		var ret CaepDeviceComplianceChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool) {
	if o == nil || IsNil(o.ReasonUser) {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *OktaUserRiskChangeEvent) HasReasonUser() bool {
	if o != nil && !IsNil(o.ReasonUser) {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepDeviceComplianceChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *OktaUserRiskChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubject returns the Subject field value
func (o *OktaUserRiskChangeEvent) GetSubject() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *OktaUserRiskChangeEvent) GetSubjectOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *OktaUserRiskChangeEvent) SetSubject(v SecurityEventSubject) {
	o.Subject = v
}

func (o OktaUserRiskChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaUserRiskChangeEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_level"] = o.CurrentLevel
	toSerialize["event_timestamp"] = o.EventTimestamp
	if !IsNil(o.InitiatingEntity) {
		toSerialize["initiating_entity"] = o.InitiatingEntity
	}
	toSerialize["previous_level"] = o.PreviousLevel
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

func (o *OktaUserRiskChangeEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_level",
		"event_timestamp",
		"previous_level",
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

	varOktaUserRiskChangeEvent := _OktaUserRiskChangeEvent{}

	err = json.Unmarshal(data, &varOktaUserRiskChangeEvent)

	if err != nil {
		return err
	}

	*o = OktaUserRiskChangeEvent(varOktaUserRiskChangeEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current_level")
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "initiating_entity")
		delete(additionalProperties, "previous_level")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaUserRiskChangeEvent struct {
	value *OktaUserRiskChangeEvent
	isSet bool
}

func (v NullableOktaUserRiskChangeEvent) Get() *OktaUserRiskChangeEvent {
	return v.value
}

func (v *NullableOktaUserRiskChangeEvent) Set(val *OktaUserRiskChangeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaUserRiskChangeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaUserRiskChangeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaUserRiskChangeEvent(val *OktaUserRiskChangeEvent) *NullableOktaUserRiskChangeEvent {
	return &NullableOktaUserRiskChangeEvent{value: val, isSet: true}
}

func (v NullableOktaUserRiskChangeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaUserRiskChangeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
