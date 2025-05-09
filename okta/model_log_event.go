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
	"time"
)

// LogEvent struct for LogEvent
type LogEvent struct {
	Actor *LogActor `json:"actor,omitempty"`
	AuthenticationContext *LogAuthenticationContext `json:"authenticationContext,omitempty"`
	Client *LogClient `json:"client,omitempty"`
	DebugContext *LogDebugContext `json:"debugContext,omitempty"`
	DisplayMessage *string `json:"displayMessage,omitempty"`
	EventType *string `json:"eventType,omitempty"`
	LegacyEventType *string `json:"legacyEventType,omitempty"`
	Outcome *LogOutcome `json:"outcome,omitempty"`
	Published *time.Time `json:"published,omitempty"`
	Request *LogRequest `json:"request,omitempty"`
	SecurityContext *LogSecurityContext `json:"securityContext,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Target []LogTarget `json:"target,omitempty"`
	Transaction *LogTransaction `json:"transaction,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogEvent LogEvent

// NewLogEvent instantiates a new LogEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogEvent() *LogEvent {
	this := LogEvent{}
	return &this
}

// NewLogEventWithDefaults instantiates a new LogEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogEventWithDefaults() *LogEvent {
	this := LogEvent{}
	return &this
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *LogEvent) GetActor() LogActor {
	if o == nil || o.Actor == nil {
		var ret LogActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetActorOk() (*LogActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *LogEvent) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given LogActor and assigns it to the Actor field.
func (o *LogEvent) SetActor(v LogActor) {
	o.Actor = &v
}

// GetAuthenticationContext returns the AuthenticationContext field value if set, zero value otherwise.
func (o *LogEvent) GetAuthenticationContext() LogAuthenticationContext {
	if o == nil || o.AuthenticationContext == nil {
		var ret LogAuthenticationContext
		return ret
	}
	return *o.AuthenticationContext
}

// GetAuthenticationContextOk returns a tuple with the AuthenticationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetAuthenticationContextOk() (*LogAuthenticationContext, bool) {
	if o == nil || o.AuthenticationContext == nil {
		return nil, false
	}
	return o.AuthenticationContext, true
}

// HasAuthenticationContext returns a boolean if a field has been set.
func (o *LogEvent) HasAuthenticationContext() bool {
	if o != nil && o.AuthenticationContext != nil {
		return true
	}

	return false
}

// SetAuthenticationContext gets a reference to the given LogAuthenticationContext and assigns it to the AuthenticationContext field.
func (o *LogEvent) SetAuthenticationContext(v LogAuthenticationContext) {
	o.AuthenticationContext = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *LogEvent) GetClient() LogClient {
	if o == nil || o.Client == nil {
		var ret LogClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetClientOk() (*LogClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *LogEvent) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given LogClient and assigns it to the Client field.
func (o *LogEvent) SetClient(v LogClient) {
	o.Client = &v
}

// GetDebugContext returns the DebugContext field value if set, zero value otherwise.
func (o *LogEvent) GetDebugContext() LogDebugContext {
	if o == nil || o.DebugContext == nil {
		var ret LogDebugContext
		return ret
	}
	return *o.DebugContext
}

// GetDebugContextOk returns a tuple with the DebugContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetDebugContextOk() (*LogDebugContext, bool) {
	if o == nil || o.DebugContext == nil {
		return nil, false
	}
	return o.DebugContext, true
}

// HasDebugContext returns a boolean if a field has been set.
func (o *LogEvent) HasDebugContext() bool {
	if o != nil && o.DebugContext != nil {
		return true
	}

	return false
}

// SetDebugContext gets a reference to the given LogDebugContext and assigns it to the DebugContext field.
func (o *LogEvent) SetDebugContext(v LogDebugContext) {
	o.DebugContext = &v
}

// GetDisplayMessage returns the DisplayMessage field value if set, zero value otherwise.
func (o *LogEvent) GetDisplayMessage() string {
	if o == nil || o.DisplayMessage == nil {
		var ret string
		return ret
	}
	return *o.DisplayMessage
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetDisplayMessageOk() (*string, bool) {
	if o == nil || o.DisplayMessage == nil {
		return nil, false
	}
	return o.DisplayMessage, true
}

// HasDisplayMessage returns a boolean if a field has been set.
func (o *LogEvent) HasDisplayMessage() bool {
	if o != nil && o.DisplayMessage != nil {
		return true
	}

	return false
}

// SetDisplayMessage gets a reference to the given string and assigns it to the DisplayMessage field.
func (o *LogEvent) SetDisplayMessage(v string) {
	o.DisplayMessage = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *LogEvent) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *LogEvent) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *LogEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetLegacyEventType returns the LegacyEventType field value if set, zero value otherwise.
func (o *LogEvent) GetLegacyEventType() string {
	if o == nil || o.LegacyEventType == nil {
		var ret string
		return ret
	}
	return *o.LegacyEventType
}

// GetLegacyEventTypeOk returns a tuple with the LegacyEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetLegacyEventTypeOk() (*string, bool) {
	if o == nil || o.LegacyEventType == nil {
		return nil, false
	}
	return o.LegacyEventType, true
}

// HasLegacyEventType returns a boolean if a field has been set.
func (o *LogEvent) HasLegacyEventType() bool {
	if o != nil && o.LegacyEventType != nil {
		return true
	}

	return false
}

// SetLegacyEventType gets a reference to the given string and assigns it to the LegacyEventType field.
func (o *LogEvent) SetLegacyEventType(v string) {
	o.LegacyEventType = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *LogEvent) GetOutcome() LogOutcome {
	if o == nil || o.Outcome == nil {
		var ret LogOutcome
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetOutcomeOk() (*LogOutcome, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *LogEvent) HasOutcome() bool {
	if o != nil && o.Outcome != nil {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given LogOutcome and assigns it to the Outcome field.
func (o *LogEvent) SetOutcome(v LogOutcome) {
	o.Outcome = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *LogEvent) GetPublished() time.Time {
	if o == nil || o.Published == nil {
		var ret time.Time
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetPublishedOk() (*time.Time, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *LogEvent) HasPublished() bool {
	if o != nil && o.Published != nil {
		return true
	}

	return false
}

// SetPublished gets a reference to the given time.Time and assigns it to the Published field.
func (o *LogEvent) SetPublished(v time.Time) {
	o.Published = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *LogEvent) GetRequest() LogRequest {
	if o == nil || o.Request == nil {
		var ret LogRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetRequestOk() (*LogRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *LogEvent) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given LogRequest and assigns it to the Request field.
func (o *LogEvent) SetRequest(v LogRequest) {
	o.Request = &v
}

// GetSecurityContext returns the SecurityContext field value if set, zero value otherwise.
func (o *LogEvent) GetSecurityContext() LogSecurityContext {
	if o == nil || o.SecurityContext == nil {
		var ret LogSecurityContext
		return ret
	}
	return *o.SecurityContext
}

// GetSecurityContextOk returns a tuple with the SecurityContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetSecurityContextOk() (*LogSecurityContext, bool) {
	if o == nil || o.SecurityContext == nil {
		return nil, false
	}
	return o.SecurityContext, true
}

// HasSecurityContext returns a boolean if a field has been set.
func (o *LogEvent) HasSecurityContext() bool {
	if o != nil && o.SecurityContext != nil {
		return true
	}

	return false
}

// SetSecurityContext gets a reference to the given LogSecurityContext and assigns it to the SecurityContext field.
func (o *LogEvent) SetSecurityContext(v LogSecurityContext) {
	o.SecurityContext = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *LogEvent) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *LogEvent) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *LogEvent) SetSeverity(v string) {
	o.Severity = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *LogEvent) GetTarget() []LogTarget {
	if o == nil || o.Target == nil {
		var ret []LogTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetTargetOk() ([]LogTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *LogEvent) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given []LogTarget and assigns it to the Target field.
func (o *LogEvent) SetTarget(v []LogTarget) {
	o.Target = v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *LogEvent) GetTransaction() LogTransaction {
	if o == nil || o.Transaction == nil {
		var ret LogTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetTransactionOk() (*LogTransaction, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *LogEvent) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given LogTransaction and assigns it to the Transaction field.
func (o *LogEvent) SetTransaction(v LogTransaction) {
	o.Transaction = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *LogEvent) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *LogEvent) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *LogEvent) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LogEvent) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEvent) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LogEvent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LogEvent) SetVersion(v string) {
	o.Version = &v
}

func (o LogEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.AuthenticationContext != nil {
		toSerialize["authenticationContext"] = o.AuthenticationContext
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.DebugContext != nil {
		toSerialize["debugContext"] = o.DebugContext
	}
	if o.DisplayMessage != nil {
		toSerialize["displayMessage"] = o.DisplayMessage
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.LegacyEventType != nil {
		toSerialize["legacyEventType"] = o.LegacyEventType
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.SecurityContext != nil {
		toSerialize["securityContext"] = o.SecurityContext
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogEvent) UnmarshalJSON(bytes []byte) (err error) {
	varLogEvent := _LogEvent{}

	err = json.Unmarshal(bytes, &varLogEvent)
	if err == nil {
		*o = LogEvent(varLogEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actor")
		delete(additionalProperties, "authenticationContext")
		delete(additionalProperties, "client")
		delete(additionalProperties, "debugContext")
		delete(additionalProperties, "displayMessage")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "legacyEventType")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "published")
		delete(additionalProperties, "request")
		delete(additionalProperties, "securityContext")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "target")
		delete(additionalProperties, "transaction")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogEvent struct {
	value *LogEvent
	isSet bool
}

func (v NullableLogEvent) Get() *LogEvent {
	return v.value
}

func (v *NullableLogEvent) Set(val *LogEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLogEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLogEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogEvent(val *LogEvent) *NullableLogEvent {
	return &NullableLogEvent{value: val, isSet: true}
}

func (v NullableLogEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

