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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SAMLPayloadExecute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayloadExecute{}

// SAMLPayloadExecute SAML assertion inline hook request
type SAMLPayloadExecute struct {
	// The inline hook cloud version
	CloudEventVersion *string `json:"cloudEventVersion,omitempty"`
	// The inline hook request header content
	ContentType *string `json:"contentType,omitempty"`
	// The individual inline hook request ID
	EventId *string `json:"eventId,omitempty"`
	// The time the inline hook request was sent
	EventTime *string `json:"eventTime,omitempty"`
	// The inline hook version
	EventTypeVersion *string          `json:"eventTypeVersion,omitempty"`
	Data             *SAMLPayLoadData `json:"data,omitempty"`
	// The type of inline hook. The SAML assertion inline hook type is `com.okta.saml.tokens.transform`.
	EventType *string `json:"eventType,omitempty"`
	// The ID and URL of the SAML assertion inline hook
	Source               *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayloadExecute SAMLPayloadExecute

// NewSAMLPayloadExecute instantiates a new SAMLPayloadExecute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayloadExecute() *SAMLPayloadExecute {
	this := SAMLPayloadExecute{}
	return &this
}

// NewSAMLPayloadExecuteWithDefaults instantiates a new SAMLPayloadExecute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayloadExecuteWithDefaults() *SAMLPayloadExecute {
	this := SAMLPayloadExecute{}
	return &this
}

// GetCloudEventVersion returns the CloudEventVersion field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetCloudEventVersion() string {
	if o == nil || IsNil(o.CloudEventVersion) {
		var ret string
		return ret
	}
	return *o.CloudEventVersion
}

// GetCloudEventVersionOk returns a tuple with the CloudEventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetCloudEventVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CloudEventVersion) {
		return nil, false
	}
	return o.CloudEventVersion, true
}

// HasCloudEventVersion returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasCloudEventVersion() bool {
	if o != nil && !IsNil(o.CloudEventVersion) {
		return true
	}

	return false
}

// SetCloudEventVersion gets a reference to the given string and assigns it to the CloudEventVersion field.
func (o *SAMLPayloadExecute) SetCloudEventVersion(v string) {
	o.CloudEventVersion = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SAMLPayloadExecute) SetContentType(v string) {
	o.ContentType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *SAMLPayloadExecute) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetEventTime() string {
	if o == nil || IsNil(o.EventTime) {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *SAMLPayloadExecute) SetEventTime(v string) {
	o.EventTime = &v
}

// GetEventTypeVersion returns the EventTypeVersion field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetEventTypeVersion() string {
	if o == nil || IsNil(o.EventTypeVersion) {
		var ret string
		return ret
	}
	return *o.EventTypeVersion
}

// GetEventTypeVersionOk returns a tuple with the EventTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetEventTypeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeVersion) {
		return nil, false
	}
	return o.EventTypeVersion, true
}

// HasEventTypeVersion returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasEventTypeVersion() bool {
	if o != nil && !IsNil(o.EventTypeVersion) {
		return true
	}

	return false
}

// SetEventTypeVersion gets a reference to the given string and assigns it to the EventTypeVersion field.
func (o *SAMLPayloadExecute) SetEventTypeVersion(v string) {
	o.EventTypeVersion = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetData() SAMLPayLoadData {
	if o == nil || IsNil(o.Data) {
		var ret SAMLPayLoadData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetDataOk() (*SAMLPayLoadData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SAMLPayLoadData and assigns it to the Data field.
func (o *SAMLPayloadExecute) SetData(v SAMLPayLoadData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *SAMLPayloadExecute) SetEventType(v string) {
	o.EventType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SAMLPayloadExecute) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayloadExecute) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SAMLPayloadExecute) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SAMLPayloadExecute) SetSource(v string) {
	o.Source = &v
}

func (o SAMLPayloadExecute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayloadExecute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudEventVersion) {
		toSerialize["cloudEventVersion"] = o.CloudEventVersion
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.EventTypeVersion) {
		toSerialize["eventTypeVersion"] = o.EventTypeVersion
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayloadExecute) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayloadExecute := _SAMLPayloadExecute{}

	err = json.Unmarshal(data, &varSAMLPayloadExecute)

	if err != nil {
		return err
	}

	*o = SAMLPayloadExecute(varSAMLPayloadExecute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudEventVersion")
		delete(additionalProperties, "contentType")
		delete(additionalProperties, "eventId")
		delete(additionalProperties, "eventTime")
		delete(additionalProperties, "eventTypeVersion")
		delete(additionalProperties, "data")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayloadExecute struct {
	value *SAMLPayloadExecute
	isSet bool
}

func (v NullableSAMLPayloadExecute) Get() *SAMLPayloadExecute {
	return v.value
}

func (v *NullableSAMLPayloadExecute) Set(val *SAMLPayloadExecute) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayloadExecute) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayloadExecute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayloadExecute(val *SAMLPayloadExecute) *NullableSAMLPayloadExecute {
	return &NullableSAMLPayloadExecute{value: val, isSet: true}
}

func (v NullableSAMLPayloadExecute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayloadExecute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
