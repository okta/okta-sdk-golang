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

// CaepSessionRevokedEvent The session of the subject was revoked
type CaepSessionRevokedEvent struct {
	// Current IP of the session
	CurrentIp *string `json:"current_ip,omitempty"`
	// Current User Agent of the session
	CurrentUserAgent *string `json:"current_user_agent,omitempty"`
	// The time of the event (UNIX timestamp)
	EventTimestamp int64 `json:"event_timestamp"`
	// The entity that initiated the event
	InitiatingEntity *string `json:"initiating_entity,omitempty"`
	// Last known IP of the session
	LastKnownIp *string `json:"last_known_ip,omitempty"`
	// Last known User Agent of the session
	LastKnownUserAgent *string `json:"last_known_user_agent,omitempty"`
	ReasonAdmin *CaepDeviceComplianceChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser *CaepDeviceComplianceChangeEventReasonUser `json:"reason_user,omitempty"`
	Subjects SecurityEventSubject `json:"subjects"`
	AdditionalProperties map[string]interface{}
}

type _CaepSessionRevokedEvent CaepSessionRevokedEvent

// NewCaepSessionRevokedEvent instantiates a new CaepSessionRevokedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepSessionRevokedEvent(eventTimestamp int64, subjects SecurityEventSubject) *CaepSessionRevokedEvent {
	this := CaepSessionRevokedEvent{}
	this.EventTimestamp = eventTimestamp
	this.Subjects = subjects
	return &this
}

// NewCaepSessionRevokedEventWithDefaults instantiates a new CaepSessionRevokedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepSessionRevokedEventWithDefaults() *CaepSessionRevokedEvent {
	this := CaepSessionRevokedEvent{}
	return &this
}

// GetCurrentIp returns the CurrentIp field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetCurrentIp() string {
	if o == nil || o.CurrentIp == nil {
		var ret string
		return ret
	}
	return *o.CurrentIp
}

// GetCurrentIpOk returns a tuple with the CurrentIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetCurrentIpOk() (*string, bool) {
	if o == nil || o.CurrentIp == nil {
		return nil, false
	}
	return o.CurrentIp, true
}

// HasCurrentIp returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasCurrentIp() bool {
	if o != nil && o.CurrentIp != nil {
		return true
	}

	return false
}

// SetCurrentIp gets a reference to the given string and assigns it to the CurrentIp field.
func (o *CaepSessionRevokedEvent) SetCurrentIp(v string) {
	o.CurrentIp = &v
}

// GetCurrentUserAgent returns the CurrentUserAgent field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetCurrentUserAgent() string {
	if o == nil || o.CurrentUserAgent == nil {
		var ret string
		return ret
	}
	return *o.CurrentUserAgent
}

// GetCurrentUserAgentOk returns a tuple with the CurrentUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetCurrentUserAgentOk() (*string, bool) {
	if o == nil || o.CurrentUserAgent == nil {
		return nil, false
	}
	return o.CurrentUserAgent, true
}

// HasCurrentUserAgent returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasCurrentUserAgent() bool {
	if o != nil && o.CurrentUserAgent != nil {
		return true
	}

	return false
}

// SetCurrentUserAgent gets a reference to the given string and assigns it to the CurrentUserAgent field.
func (o *CaepSessionRevokedEvent) SetCurrentUserAgent(v string) {
	o.CurrentUserAgent = &v
}

// GetEventTimestamp returns the EventTimestamp field value
func (o *CaepSessionRevokedEvent) GetEventTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTimestamp, true
}

// SetEventTimestamp sets field value
func (o *CaepSessionRevokedEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = v
}

// GetInitiatingEntity returns the InitiatingEntity field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetInitiatingEntity() string {
	if o == nil || o.InitiatingEntity == nil {
		var ret string
		return ret
	}
	return *o.InitiatingEntity
}

// GetInitiatingEntityOk returns a tuple with the InitiatingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetInitiatingEntityOk() (*string, bool) {
	if o == nil || o.InitiatingEntity == nil {
		return nil, false
	}
	return o.InitiatingEntity, true
}

// HasInitiatingEntity returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasInitiatingEntity() bool {
	if o != nil && o.InitiatingEntity != nil {
		return true
	}

	return false
}

// SetInitiatingEntity gets a reference to the given string and assigns it to the InitiatingEntity field.
func (o *CaepSessionRevokedEvent) SetInitiatingEntity(v string) {
	o.InitiatingEntity = &v
}

// GetLastKnownIp returns the LastKnownIp field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetLastKnownIp() string {
	if o == nil || o.LastKnownIp == nil {
		var ret string
		return ret
	}
	return *o.LastKnownIp
}

// GetLastKnownIpOk returns a tuple with the LastKnownIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetLastKnownIpOk() (*string, bool) {
	if o == nil || o.LastKnownIp == nil {
		return nil, false
	}
	return o.LastKnownIp, true
}

// HasLastKnownIp returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasLastKnownIp() bool {
	if o != nil && o.LastKnownIp != nil {
		return true
	}

	return false
}

// SetLastKnownIp gets a reference to the given string and assigns it to the LastKnownIp field.
func (o *CaepSessionRevokedEvent) SetLastKnownIp(v string) {
	o.LastKnownIp = &v
}

// GetLastKnownUserAgent returns the LastKnownUserAgent field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetLastKnownUserAgent() string {
	if o == nil || o.LastKnownUserAgent == nil {
		var ret string
		return ret
	}
	return *o.LastKnownUserAgent
}

// GetLastKnownUserAgentOk returns a tuple with the LastKnownUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetLastKnownUserAgentOk() (*string, bool) {
	if o == nil || o.LastKnownUserAgent == nil {
		return nil, false
	}
	return o.LastKnownUserAgent, true
}

// HasLastKnownUserAgent returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasLastKnownUserAgent() bool {
	if o != nil && o.LastKnownUserAgent != nil {
		return true
	}

	return false
}

// SetLastKnownUserAgent gets a reference to the given string and assigns it to the LastKnownUserAgent field.
func (o *CaepSessionRevokedEvent) SetLastKnownUserAgent(v string) {
	o.LastKnownUserAgent = &v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin {
	if o == nil || o.ReasonAdmin == nil {
		var ret CaepDeviceComplianceChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool) {
	if o == nil || o.ReasonAdmin == nil {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasReasonAdmin() bool {
	if o != nil && o.ReasonAdmin != nil {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepDeviceComplianceChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *CaepSessionRevokedEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *CaepSessionRevokedEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser {
	if o == nil || o.ReasonUser == nil {
		var ret CaepDeviceComplianceChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool) {
	if o == nil || o.ReasonUser == nil {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *CaepSessionRevokedEvent) HasReasonUser() bool {
	if o != nil && o.ReasonUser != nil {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepDeviceComplianceChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *CaepSessionRevokedEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubjects returns the Subjects field value
func (o *CaepSessionRevokedEvent) GetSubjects() SecurityEventSubject {
	if o == nil {
		var ret SecurityEventSubject
		return ret
	}

	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *CaepSessionRevokedEvent) GetSubjectsOk() (*SecurityEventSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value
func (o *CaepSessionRevokedEvent) SetSubjects(v SecurityEventSubject) {
	o.Subjects = v
}

func (o CaepSessionRevokedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentIp != nil {
		toSerialize["current_ip"] = o.CurrentIp
	}
	if o.CurrentUserAgent != nil {
		toSerialize["current_user_agent"] = o.CurrentUserAgent
	}
	if true {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if o.InitiatingEntity != nil {
		toSerialize["initiating_entity"] = o.InitiatingEntity
	}
	if o.LastKnownIp != nil {
		toSerialize["last_known_ip"] = o.LastKnownIp
	}
	if o.LastKnownUserAgent != nil {
		toSerialize["last_known_user_agent"] = o.LastKnownUserAgent
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

func (o *CaepSessionRevokedEvent) UnmarshalJSON(bytes []byte) (err error) {
	varCaepSessionRevokedEvent := _CaepSessionRevokedEvent{}

	err = json.Unmarshal(bytes, &varCaepSessionRevokedEvent)
	if err == nil {
		*o = CaepSessionRevokedEvent(varCaepSessionRevokedEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "current_ip")
		delete(additionalProperties, "current_user_agent")
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "initiating_entity")
		delete(additionalProperties, "last_known_ip")
		delete(additionalProperties, "last_known_user_agent")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subjects")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCaepSessionRevokedEvent struct {
	value *CaepSessionRevokedEvent
	isSet bool
}

func (v NullableCaepSessionRevokedEvent) Get() *CaepSessionRevokedEvent {
	return v.value
}

func (v *NullableCaepSessionRevokedEvent) Set(val *CaepSessionRevokedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepSessionRevokedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepSessionRevokedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepSessionRevokedEvent(val *CaepSessionRevokedEvent) *NullableCaepSessionRevokedEvent {
	return &NullableCaepSessionRevokedEvent{value: val, isSet: true}
}

func (v NullableCaepSessionRevokedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepSessionRevokedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

