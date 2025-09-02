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

// PasswordImportRequestExecute Password import inline hook request
type PasswordImportRequestExecute struct {
	// The inline hook cloud version
	CloudEventVersion *string `json:"cloudEventVersion,omitempty"`
	// The inline hook request header content
	ContentType *string `json:"contentType,omitempty"`
	// The individual inline hook request ID
	EventId *string `json:"eventId,omitempty"`
	// The time the inline hook request was sent
	EventTime *string `json:"eventTime,omitempty"`
	// The inline hook version
	EventTypeVersion *string `json:"eventTypeVersion,omitempty"`
	Data *PasswordImportRequestData `json:"data,omitempty"`
	// The type of inline hook. The password import inline hook type is `com.okta.user.credential.password.import`.
	EventType *string `json:"eventType,omitempty"`
	// The ID and URL of the password import inline hook
	Source *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportRequestExecute PasswordImportRequestExecute

// NewPasswordImportRequestExecute instantiates a new PasswordImportRequestExecute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportRequestExecute() *PasswordImportRequestExecute {
	this := PasswordImportRequestExecute{}
	return &this
}

// NewPasswordImportRequestExecuteWithDefaults instantiates a new PasswordImportRequestExecute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportRequestExecuteWithDefaults() *PasswordImportRequestExecute {
	this := PasswordImportRequestExecute{}
	return &this
}

// GetCloudEventVersion returns the CloudEventVersion field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetCloudEventVersion() string {
	if o == nil || o.CloudEventVersion == nil {
		var ret string
		return ret
	}
	return *o.CloudEventVersion
}

// GetCloudEventVersionOk returns a tuple with the CloudEventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetCloudEventVersionOk() (*string, bool) {
	if o == nil || o.CloudEventVersion == nil {
		return nil, false
	}
	return o.CloudEventVersion, true
}

// HasCloudEventVersion returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasCloudEventVersion() bool {
	if o != nil && o.CloudEventVersion != nil {
		return true
	}

	return false
}

// SetCloudEventVersion gets a reference to the given string and assigns it to the CloudEventVersion field.
func (o *PasswordImportRequestExecute) SetCloudEventVersion(v string) {
	o.CloudEventVersion = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PasswordImportRequestExecute) SetContentType(v string) {
	o.ContentType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *PasswordImportRequestExecute) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetEventTime() string {
	if o == nil || o.EventTime == nil {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetEventTimeOk() (*string, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasEventTime() bool {
	if o != nil && o.EventTime != nil {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *PasswordImportRequestExecute) SetEventTime(v string) {
	o.EventTime = &v
}

// GetEventTypeVersion returns the EventTypeVersion field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetEventTypeVersion() string {
	if o == nil || o.EventTypeVersion == nil {
		var ret string
		return ret
	}
	return *o.EventTypeVersion
}

// GetEventTypeVersionOk returns a tuple with the EventTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetEventTypeVersionOk() (*string, bool) {
	if o == nil || o.EventTypeVersion == nil {
		return nil, false
	}
	return o.EventTypeVersion, true
}

// HasEventTypeVersion returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasEventTypeVersion() bool {
	if o != nil && o.EventTypeVersion != nil {
		return true
	}

	return false
}

// SetEventTypeVersion gets a reference to the given string and assigns it to the EventTypeVersion field.
func (o *PasswordImportRequestExecute) SetEventTypeVersion(v string) {
	o.EventTypeVersion = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetData() PasswordImportRequestData {
	if o == nil || o.Data == nil {
		var ret PasswordImportRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetDataOk() (*PasswordImportRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PasswordImportRequestData and assigns it to the Data field.
func (o *PasswordImportRequestExecute) SetData(v PasswordImportRequestData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *PasswordImportRequestExecute) SetEventType(v string) {
	o.EventType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PasswordImportRequestExecute) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestExecute) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PasswordImportRequestExecute) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PasswordImportRequestExecute) SetSource(v string) {
	o.Source = &v
}

func (o PasswordImportRequestExecute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudEventVersion != nil {
		toSerialize["cloudEventVersion"] = o.CloudEventVersion
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if o.EventTime != nil {
		toSerialize["eventTime"] = o.EventTime
	}
	if o.EventTypeVersion != nil {
		toSerialize["eventTypeVersion"] = o.EventTypeVersion
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordImportRequestExecute) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordImportRequestExecute := _PasswordImportRequestExecute{}

	err = json.Unmarshal(bytes, &varPasswordImportRequestExecute)
	if err == nil {
		*o = PasswordImportRequestExecute(varPasswordImportRequestExecute)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "cloudEventVersion")
		delete(additionalProperties, "contentType")
		delete(additionalProperties, "eventId")
		delete(additionalProperties, "eventTime")
		delete(additionalProperties, "eventTypeVersion")
		delete(additionalProperties, "data")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordImportRequestExecute struct {
	value *PasswordImportRequestExecute
	isSet bool
}

func (v NullablePasswordImportRequestExecute) Get() *PasswordImportRequestExecute {
	return v.value
}

func (v *NullablePasswordImportRequestExecute) Set(val *PasswordImportRequestExecute) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportRequestExecute) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportRequestExecute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportRequestExecute(val *PasswordImportRequestExecute) *NullablePasswordImportRequestExecute {
	return &NullablePasswordImportRequestExecute{value: val, isSet: true}
}

func (v NullablePasswordImportRequestExecute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportRequestExecute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

