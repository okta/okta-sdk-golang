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

// TelephonyRequestExecute Telephony inline hook request body
type TelephonyRequestExecute struct {
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
	Data *TelephonyRequestData `json:"data,omitempty"`
	// The type of inline hook. The telephony inline hook type is `com.okta.telephony.provider`.
	EventType *string `json:"eventType,omitempty"`
	// The type of inline hook request. For example, `com.okta.user.telephony.pre-enrollment`.
	RequestType *string `json:"requestType,omitempty"`
	// The ID and URL of the telephony inline hook
	Source *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyRequestExecute TelephonyRequestExecute

// NewTelephonyRequestExecute instantiates a new TelephonyRequestExecute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyRequestExecute() *TelephonyRequestExecute {
	this := TelephonyRequestExecute{}
	return &this
}

// NewTelephonyRequestExecuteWithDefaults instantiates a new TelephonyRequestExecute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyRequestExecuteWithDefaults() *TelephonyRequestExecute {
	this := TelephonyRequestExecute{}
	return &this
}

// GetCloudEventVersion returns the CloudEventVersion field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetCloudEventVersion() string {
	if o == nil || o.CloudEventVersion == nil {
		var ret string
		return ret
	}
	return *o.CloudEventVersion
}

// GetCloudEventVersionOk returns a tuple with the CloudEventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetCloudEventVersionOk() (*string, bool) {
	if o == nil || o.CloudEventVersion == nil {
		return nil, false
	}
	return o.CloudEventVersion, true
}

// HasCloudEventVersion returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasCloudEventVersion() bool {
	if o != nil && o.CloudEventVersion != nil {
		return true
	}

	return false
}

// SetCloudEventVersion gets a reference to the given string and assigns it to the CloudEventVersion field.
func (o *TelephonyRequestExecute) SetCloudEventVersion(v string) {
	o.CloudEventVersion = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *TelephonyRequestExecute) SetContentType(v string) {
	o.ContentType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *TelephonyRequestExecute) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetEventTime() string {
	if o == nil || o.EventTime == nil {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetEventTimeOk() (*string, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasEventTime() bool {
	if o != nil && o.EventTime != nil {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *TelephonyRequestExecute) SetEventTime(v string) {
	o.EventTime = &v
}

// GetEventTypeVersion returns the EventTypeVersion field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetEventTypeVersion() string {
	if o == nil || o.EventTypeVersion == nil {
		var ret string
		return ret
	}
	return *o.EventTypeVersion
}

// GetEventTypeVersionOk returns a tuple with the EventTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetEventTypeVersionOk() (*string, bool) {
	if o == nil || o.EventTypeVersion == nil {
		return nil, false
	}
	return o.EventTypeVersion, true
}

// HasEventTypeVersion returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasEventTypeVersion() bool {
	if o != nil && o.EventTypeVersion != nil {
		return true
	}

	return false
}

// SetEventTypeVersion gets a reference to the given string and assigns it to the EventTypeVersion field.
func (o *TelephonyRequestExecute) SetEventTypeVersion(v string) {
	o.EventTypeVersion = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetData() TelephonyRequestData {
	if o == nil || o.Data == nil {
		var ret TelephonyRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetDataOk() (*TelephonyRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TelephonyRequestData and assigns it to the Data field.
func (o *TelephonyRequestExecute) SetData(v TelephonyRequestData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *TelephonyRequestExecute) SetEventType(v string) {
	o.EventType = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *TelephonyRequestExecute) SetRequestType(v string) {
	o.RequestType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TelephonyRequestExecute) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestExecute) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TelephonyRequestExecute) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TelephonyRequestExecute) SetSource(v string) {
	o.Source = &v
}

func (o TelephonyRequestExecute) MarshalJSON() ([]byte, error) {
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
	if o.RequestType != nil {
		toSerialize["requestType"] = o.RequestType
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyRequestExecute) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyRequestExecute := _TelephonyRequestExecute{}

	err = json.Unmarshal(bytes, &varTelephonyRequestExecute)
	if err == nil {
		*o = TelephonyRequestExecute(varTelephonyRequestExecute)
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
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTelephonyRequestExecute struct {
	value *TelephonyRequestExecute
	isSet bool
}

func (v NullableTelephonyRequestExecute) Get() *TelephonyRequestExecute {
	return v.value
}

func (v *NullableTelephonyRequestExecute) Set(val *TelephonyRequestExecute) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyRequestExecute) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyRequestExecute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyRequestExecute(val *TelephonyRequestExecute) *NullableTelephonyRequestExecute {
	return &NullableTelephonyRequestExecute{value: val, isSet: true}
}

func (v NullableTelephonyRequestExecute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyRequestExecute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

