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

// checks if the CaepCredentialChangeEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaepCredentialChangeEvent{}

// CaepCredentialChangeEvent The credential was created, changed, revoked or deleted
type CaepCredentialChangeEvent struct {
	// The type of action done towards the credential
	ChangeType string `json:"change_type"`
	// The credential type of the changed credential. It will one of the supported enum values or any other credential type supported mutually by the Transmitter and the Receiver.
	CredentialType string `json:"credential_type"`
	// The time of the event (UNIX timestamp)
	EventTimestamp *int64 `json:"event_timestamp,omitempty"`
	// FIDO2 Authenticator Attestation GUID
	Fido2Aaguid *string `json:"fido2_aaguid,omitempty"`
	// Credential friendly name
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The entity that initiated the event
	InitiatingEntity     *string                               `json:"initiating_entity,omitempty"`
	ReasonAdmin          *CaepCredentialChangeEventReasonAdmin `json:"reason_admin,omitempty"`
	ReasonUser           *CaepCredentialChangeEventReasonUser  `json:"reason_user,omitempty"`
	Subject              *SecurityEventSubject                 `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaepCredentialChangeEvent CaepCredentialChangeEvent

// NewCaepCredentialChangeEvent instantiates a new CaepCredentialChangeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepCredentialChangeEvent(changeType string, credentialType string) *CaepCredentialChangeEvent {
	this := CaepCredentialChangeEvent{}
	this.ChangeType = changeType
	this.CredentialType = credentialType
	return &this
}

// NewCaepCredentialChangeEventWithDefaults instantiates a new CaepCredentialChangeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepCredentialChangeEventWithDefaults() *CaepCredentialChangeEvent {
	this := CaepCredentialChangeEvent{}
	return &this
}

// GetChangeType returns the ChangeType field value
func (o *CaepCredentialChangeEvent) GetChangeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetChangeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeType, true
}

// SetChangeType sets field value
func (o *CaepCredentialChangeEvent) SetChangeType(v string) {
	o.ChangeType = v
}

// GetCredentialType returns the CredentialType field value
func (o *CaepCredentialChangeEvent) GetCredentialType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetCredentialTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialType, true
}

// SetCredentialType sets field value
func (o *CaepCredentialChangeEvent) SetCredentialType(v string) {
	o.CredentialType = v
}

// GetEventTimestamp returns the EventTimestamp field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetEventTimestamp() int64 {
	if o == nil || IsNil(o.EventTimestamp) {
		var ret int64
		return ret
	}
	return *o.EventTimestamp
}

// GetEventTimestampOk returns a tuple with the EventTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetEventTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.EventTimestamp) {
		return nil, false
	}
	return o.EventTimestamp, true
}

// HasEventTimestamp returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasEventTimestamp() bool {
	if o != nil && !IsNil(o.EventTimestamp) {
		return true
	}

	return false
}

// SetEventTimestamp gets a reference to the given int64 and assigns it to the EventTimestamp field.
func (o *CaepCredentialChangeEvent) SetEventTimestamp(v int64) {
	o.EventTimestamp = &v
}

// GetFido2Aaguid returns the Fido2Aaguid field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetFido2Aaguid() string {
	if o == nil || IsNil(o.Fido2Aaguid) {
		var ret string
		return ret
	}
	return *o.Fido2Aaguid
}

// GetFido2AaguidOk returns a tuple with the Fido2Aaguid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetFido2AaguidOk() (*string, bool) {
	if o == nil || IsNil(o.Fido2Aaguid) {
		return nil, false
	}
	return o.Fido2Aaguid, true
}

// HasFido2Aaguid returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasFido2Aaguid() bool {
	if o != nil && !IsNil(o.Fido2Aaguid) {
		return true
	}

	return false
}

// SetFido2Aaguid gets a reference to the given string and assigns it to the Fido2Aaguid field.
func (o *CaepCredentialChangeEvent) SetFido2Aaguid(v string) {
	o.Fido2Aaguid = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *CaepCredentialChangeEvent) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetInitiatingEntity returns the InitiatingEntity field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetInitiatingEntity() string {
	if o == nil || IsNil(o.InitiatingEntity) {
		var ret string
		return ret
	}
	return *o.InitiatingEntity
}

// GetInitiatingEntityOk returns a tuple with the InitiatingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetInitiatingEntityOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatingEntity) {
		return nil, false
	}
	return o.InitiatingEntity, true
}

// HasInitiatingEntity returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasInitiatingEntity() bool {
	if o != nil && !IsNil(o.InitiatingEntity) {
		return true
	}

	return false
}

// SetInitiatingEntity gets a reference to the given string and assigns it to the InitiatingEntity field.
func (o *CaepCredentialChangeEvent) SetInitiatingEntity(v string) {
	o.InitiatingEntity = &v
}

// GetReasonAdmin returns the ReasonAdmin field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetReasonAdmin() CaepCredentialChangeEventReasonAdmin {
	if o == nil || IsNil(o.ReasonAdmin) {
		var ret CaepCredentialChangeEventReasonAdmin
		return ret
	}
	return *o.ReasonAdmin
}

// GetReasonAdminOk returns a tuple with the ReasonAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetReasonAdminOk() (*CaepCredentialChangeEventReasonAdmin, bool) {
	if o == nil || IsNil(o.ReasonAdmin) {
		return nil, false
	}
	return o.ReasonAdmin, true
}

// HasReasonAdmin returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasReasonAdmin() bool {
	if o != nil && !IsNil(o.ReasonAdmin) {
		return true
	}

	return false
}

// SetReasonAdmin gets a reference to the given CaepCredentialChangeEventReasonAdmin and assigns it to the ReasonAdmin field.
func (o *CaepCredentialChangeEvent) SetReasonAdmin(v CaepCredentialChangeEventReasonAdmin) {
	o.ReasonAdmin = &v
}

// GetReasonUser returns the ReasonUser field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetReasonUser() CaepCredentialChangeEventReasonUser {
	if o == nil || IsNil(o.ReasonUser) {
		var ret CaepCredentialChangeEventReasonUser
		return ret
	}
	return *o.ReasonUser
}

// GetReasonUserOk returns a tuple with the ReasonUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetReasonUserOk() (*CaepCredentialChangeEventReasonUser, bool) {
	if o == nil || IsNil(o.ReasonUser) {
		return nil, false
	}
	return o.ReasonUser, true
}

// HasReasonUser returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasReasonUser() bool {
	if o != nil && !IsNil(o.ReasonUser) {
		return true
	}

	return false
}

// SetReasonUser gets a reference to the given CaepCredentialChangeEventReasonUser and assigns it to the ReasonUser field.
func (o *CaepCredentialChangeEvent) SetReasonUser(v CaepCredentialChangeEventReasonUser) {
	o.ReasonUser = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CaepCredentialChangeEvent) GetSubject() SecurityEventSubject {
	if o == nil || IsNil(o.Subject) {
		var ret SecurityEventSubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEvent) GetSubjectOk() (*SecurityEventSubject, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CaepCredentialChangeEvent) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given SecurityEventSubject and assigns it to the Subject field.
func (o *CaepCredentialChangeEvent) SetSubject(v SecurityEventSubject) {
	o.Subject = &v
}

func (o CaepCredentialChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaepCredentialChangeEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["change_type"] = o.ChangeType
	toSerialize["credential_type"] = o.CredentialType
	if !IsNil(o.EventTimestamp) {
		toSerialize["event_timestamp"] = o.EventTimestamp
	}
	if !IsNil(o.Fido2Aaguid) {
		toSerialize["fido2_aaguid"] = o.Fido2Aaguid
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendly_name"] = o.FriendlyName
	}
	if !IsNil(o.InitiatingEntity) {
		toSerialize["initiating_entity"] = o.InitiatingEntity
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

func (o *CaepCredentialChangeEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"change_type",
		"credential_type",
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

	varCaepCredentialChangeEvent := _CaepCredentialChangeEvent{}

	err = json.Unmarshal(data, &varCaepCredentialChangeEvent)

	if err != nil {
		return err
	}

	*o = CaepCredentialChangeEvent(varCaepCredentialChangeEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "change_type")
		delete(additionalProperties, "credential_type")
		delete(additionalProperties, "event_timestamp")
		delete(additionalProperties, "fido2_aaguid")
		delete(additionalProperties, "friendly_name")
		delete(additionalProperties, "initiating_entity")
		delete(additionalProperties, "reason_admin")
		delete(additionalProperties, "reason_user")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaepCredentialChangeEvent struct {
	value *CaepCredentialChangeEvent
	isSet bool
}

func (v NullableCaepCredentialChangeEvent) Get() *CaepCredentialChangeEvent {
	return v.value
}

func (v *NullableCaepCredentialChangeEvent) Set(val *CaepCredentialChangeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepCredentialChangeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepCredentialChangeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepCredentialChangeEvent(val *CaepCredentialChangeEvent) *NullableCaepCredentialChangeEvent {
	return &NullableCaepCredentialChangeEvent{value: val, isSet: true}
}

func (v NullableCaepCredentialChangeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepCredentialChangeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
