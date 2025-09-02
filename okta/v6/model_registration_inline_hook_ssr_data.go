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
	"reflect"
	"strings"
)

// RegistrationInlineHookSSRData struct for RegistrationInlineHookSSRData
type RegistrationInlineHookSSRData struct {
	RegistrationInlineHookRequest
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
	Data *RegistrationInlineHookSSRDataAllOfData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookSSRData RegistrationInlineHookSSRData

// NewRegistrationInlineHookSSRData instantiates a new RegistrationInlineHookSSRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookSSRData() *RegistrationInlineHookSSRData {
	this := RegistrationInlineHookSSRData{}
	return &this
}

// NewRegistrationInlineHookSSRDataWithDefaults instantiates a new RegistrationInlineHookSSRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookSSRDataWithDefaults() *RegistrationInlineHookSSRData {
	this := RegistrationInlineHookSSRData{}
	return &this
}

// GetCloudEventVersion returns the CloudEventVersion field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRData) GetCloudEventVersion() string {
	if o == nil || o.CloudEventVersion == nil {
		var ret string
		return ret
	}
	return *o.CloudEventVersion
}

// GetCloudEventVersionOk returns a tuple with the CloudEventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRData) GetCloudEventVersionOk() (*string, bool) {
	if o == nil || o.CloudEventVersion == nil {
		return nil, false
	}
	return o.CloudEventVersion, true
}

// HasCloudEventVersion returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRData) HasCloudEventVersion() bool {
	if o != nil && o.CloudEventVersion != nil {
		return true
	}

	return false
}

// SetCloudEventVersion gets a reference to the given string and assigns it to the CloudEventVersion field.
func (o *RegistrationInlineHookSSRData) SetCloudEventVersion(v string) {
	o.CloudEventVersion = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRData) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRData) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRData) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *RegistrationInlineHookSSRData) SetContentType(v string) {
	o.ContentType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRData) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRData) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRData) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *RegistrationInlineHookSSRData) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRData) GetEventTime() string {
	if o == nil || o.EventTime == nil {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRData) GetEventTimeOk() (*string, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRData) HasEventTime() bool {
	if o != nil && o.EventTime != nil {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *RegistrationInlineHookSSRData) SetEventTime(v string) {
	o.EventTime = &v
}

// GetEventTypeVersion returns the EventTypeVersion field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRData) GetEventTypeVersion() string {
	if o == nil || o.EventTypeVersion == nil {
		var ret string
		return ret
	}
	return *o.EventTypeVersion
}

// GetEventTypeVersionOk returns a tuple with the EventTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRData) GetEventTypeVersionOk() (*string, bool) {
	if o == nil || o.EventTypeVersion == nil {
		return nil, false
	}
	return o.EventTypeVersion, true
}

// HasEventTypeVersion returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRData) HasEventTypeVersion() bool {
	if o != nil && o.EventTypeVersion != nil {
		return true
	}

	return false
}

// SetEventTypeVersion gets a reference to the given string and assigns it to the EventTypeVersion field.
func (o *RegistrationInlineHookSSRData) SetEventTypeVersion(v string) {
	o.EventTypeVersion = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRData) GetData() RegistrationInlineHookSSRDataAllOfData {
	if o == nil || o.Data == nil {
		var ret RegistrationInlineHookSSRDataAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRData) GetDataOk() (*RegistrationInlineHookSSRDataAllOfData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RegistrationInlineHookSSRDataAllOfData and assigns it to the Data field.
func (o *RegistrationInlineHookSSRData) SetData(v RegistrationInlineHookSSRDataAllOfData) {
	o.Data = &v
}

func (o RegistrationInlineHookSSRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRegistrationInlineHookRequest, errRegistrationInlineHookRequest := json.Marshal(o.RegistrationInlineHookRequest)
	if errRegistrationInlineHookRequest != nil {
		return []byte{}, errRegistrationInlineHookRequest
	}
	errRegistrationInlineHookRequest = json.Unmarshal([]byte(serializedRegistrationInlineHookRequest), &toSerialize)
	if errRegistrationInlineHookRequest != nil {
		return []byte{}, errRegistrationInlineHookRequest
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegistrationInlineHookSSRData) UnmarshalJSON(bytes []byte) (err error) {
	type RegistrationInlineHookSSRDataWithoutEmbeddedStruct struct {
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
		Data *RegistrationInlineHookSSRDataAllOfData `json:"data,omitempty"`
	}

	varRegistrationInlineHookSSRDataWithoutEmbeddedStruct := RegistrationInlineHookSSRDataWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRegistrationInlineHookSSRDataWithoutEmbeddedStruct)
	if err == nil {
		varRegistrationInlineHookSSRData := _RegistrationInlineHookSSRData{}
		varRegistrationInlineHookSSRData.CloudEventVersion = varRegistrationInlineHookSSRDataWithoutEmbeddedStruct.CloudEventVersion
		varRegistrationInlineHookSSRData.ContentType = varRegistrationInlineHookSSRDataWithoutEmbeddedStruct.ContentType
		varRegistrationInlineHookSSRData.EventId = varRegistrationInlineHookSSRDataWithoutEmbeddedStruct.EventId
		varRegistrationInlineHookSSRData.EventTime = varRegistrationInlineHookSSRDataWithoutEmbeddedStruct.EventTime
		varRegistrationInlineHookSSRData.EventTypeVersion = varRegistrationInlineHookSSRDataWithoutEmbeddedStruct.EventTypeVersion
		varRegistrationInlineHookSSRData.Data = varRegistrationInlineHookSSRDataWithoutEmbeddedStruct.Data
		*o = RegistrationInlineHookSSRData(varRegistrationInlineHookSSRData)
	} else {
		return err
	}

	varRegistrationInlineHookSSRData := _RegistrationInlineHookSSRData{}

	err = json.Unmarshal(bytes, &varRegistrationInlineHookSSRData)
	if err == nil {
		o.RegistrationInlineHookRequest = varRegistrationInlineHookSSRData.RegistrationInlineHookRequest
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

		// remove fields from embedded structs
		reflectRegistrationInlineHookRequest := reflect.ValueOf(o.RegistrationInlineHookRequest)
		for i := 0; i < reflectRegistrationInlineHookRequest.Type().NumField(); i++ {
			t := reflectRegistrationInlineHookRequest.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRegistrationInlineHookSSRData struct {
	value *RegistrationInlineHookSSRData
	isSet bool
}

func (v NullableRegistrationInlineHookSSRData) Get() *RegistrationInlineHookSSRData {
	return v.value
}

func (v *NullableRegistrationInlineHookSSRData) Set(val *RegistrationInlineHookSSRData) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookSSRData) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookSSRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookSSRData(val *RegistrationInlineHookSSRData) *NullableRegistrationInlineHookSSRData {
	return &NullableRegistrationInlineHookSSRData{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookSSRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookSSRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

