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
	"reflect"
	"strings"
)

// checks if the RegistrationInlineHookPPData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookPPData{}

// RegistrationInlineHookPPData struct for RegistrationInlineHookPPData
type RegistrationInlineHookPPData struct {
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
	EventTypeVersion     *string                                `json:"eventTypeVersion,omitempty"`
	Data                 *RegistrationInlineHookPPDataAllOfData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookPPData RegistrationInlineHookPPData

// NewRegistrationInlineHookPPData instantiates a new RegistrationInlineHookPPData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookPPData() *RegistrationInlineHookPPData {
	this := RegistrationInlineHookPPData{}
	return &this
}

// NewRegistrationInlineHookPPDataWithDefaults instantiates a new RegistrationInlineHookPPData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookPPDataWithDefaults() *RegistrationInlineHookPPData {
	this := RegistrationInlineHookPPData{}
	return &this
}

// GetCloudEventVersion returns the CloudEventVersion field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPData) GetCloudEventVersion() string {
	if o == nil || IsNil(o.CloudEventVersion) {
		var ret string
		return ret
	}
	return *o.CloudEventVersion
}

// GetCloudEventVersionOk returns a tuple with the CloudEventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPData) GetCloudEventVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CloudEventVersion) {
		return nil, false
	}
	return o.CloudEventVersion, true
}

// HasCloudEventVersion returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPData) HasCloudEventVersion() bool {
	if o != nil && !IsNil(o.CloudEventVersion) {
		return true
	}

	return false
}

// SetCloudEventVersion gets a reference to the given string and assigns it to the CloudEventVersion field.
func (o *RegistrationInlineHookPPData) SetCloudEventVersion(v string) {
	o.CloudEventVersion = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPData) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPData) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPData) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *RegistrationInlineHookPPData) SetContentType(v string) {
	o.ContentType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPData) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPData) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPData) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *RegistrationInlineHookPPData) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPData) GetEventTime() string {
	if o == nil || IsNil(o.EventTime) {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPData) GetEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPData) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *RegistrationInlineHookPPData) SetEventTime(v string) {
	o.EventTime = &v
}

// GetEventTypeVersion returns the EventTypeVersion field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPData) GetEventTypeVersion() string {
	if o == nil || IsNil(o.EventTypeVersion) {
		var ret string
		return ret
	}
	return *o.EventTypeVersion
}

// GetEventTypeVersionOk returns a tuple with the EventTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPData) GetEventTypeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeVersion) {
		return nil, false
	}
	return o.EventTypeVersion, true
}

// HasEventTypeVersion returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPData) HasEventTypeVersion() bool {
	if o != nil && !IsNil(o.EventTypeVersion) {
		return true
	}

	return false
}

// SetEventTypeVersion gets a reference to the given string and assigns it to the EventTypeVersion field.
func (o *RegistrationInlineHookPPData) SetEventTypeVersion(v string) {
	o.EventTypeVersion = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPData) GetData() RegistrationInlineHookPPDataAllOfData {
	if o == nil || IsNil(o.Data) {
		var ret RegistrationInlineHookPPDataAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPData) GetDataOk() (*RegistrationInlineHookPPDataAllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RegistrationInlineHookPPDataAllOfData and assigns it to the Data field.
func (o *RegistrationInlineHookPPData) SetData(v RegistrationInlineHookPPDataAllOfData) {
	o.Data = &v
}

func (o RegistrationInlineHookPPData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookPPData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedRegistrationInlineHookRequest, errRegistrationInlineHookRequest := json.Marshal(o.RegistrationInlineHookRequest)
	if errRegistrationInlineHookRequest != nil {
		return map[string]interface{}{}, errRegistrationInlineHookRequest
	}
	errRegistrationInlineHookRequest = json.Unmarshal([]byte(serializedRegistrationInlineHookRequest), &toSerialize)
	if errRegistrationInlineHookRequest != nil {
		return map[string]interface{}{}, errRegistrationInlineHookRequest
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookPPData) UnmarshalJSON(data []byte) (err error) {
	type RegistrationInlineHookPPDataWithoutEmbeddedStruct struct {
		// The inline hook cloud version
		CloudEventVersion *string `json:"cloudEventVersion,omitempty"`
		// The inline hook request header content
		ContentType *string `json:"contentType,omitempty"`
		// The individual inline hook request ID
		EventId *string `json:"eventId,omitempty"`
		// The time the inline hook request was sent
		EventTime *string `json:"eventTime,omitempty"`
		// The inline hook version
		EventTypeVersion *string                                `json:"eventTypeVersion,omitempty"`
		Data             *RegistrationInlineHookPPDataAllOfData `json:"data,omitempty"`
	}

	varRegistrationInlineHookPPDataWithoutEmbeddedStruct := RegistrationInlineHookPPDataWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRegistrationInlineHookPPDataWithoutEmbeddedStruct)
	if err == nil {
		varRegistrationInlineHookPPData := _RegistrationInlineHookPPData{}
		varRegistrationInlineHookPPData.CloudEventVersion = varRegistrationInlineHookPPDataWithoutEmbeddedStruct.CloudEventVersion
		varRegistrationInlineHookPPData.ContentType = varRegistrationInlineHookPPDataWithoutEmbeddedStruct.ContentType
		varRegistrationInlineHookPPData.EventId = varRegistrationInlineHookPPDataWithoutEmbeddedStruct.EventId
		varRegistrationInlineHookPPData.EventTime = varRegistrationInlineHookPPDataWithoutEmbeddedStruct.EventTime
		varRegistrationInlineHookPPData.EventTypeVersion = varRegistrationInlineHookPPDataWithoutEmbeddedStruct.EventTypeVersion
		varRegistrationInlineHookPPData.Data = varRegistrationInlineHookPPDataWithoutEmbeddedStruct.Data
		*o = RegistrationInlineHookPPData(varRegistrationInlineHookPPData)
	} else {
		return err
	}

	varRegistrationInlineHookPPData := _RegistrationInlineHookPPData{}

	err = json.Unmarshal(data, &varRegistrationInlineHookPPData)
	if err == nil {
		o.RegistrationInlineHookRequest = varRegistrationInlineHookPPData.RegistrationInlineHookRequest
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableRegistrationInlineHookPPData struct {
	value *RegistrationInlineHookPPData
	isSet bool
}

func (v NullableRegistrationInlineHookPPData) Get() *RegistrationInlineHookPPData {
	return v.value
}

func (v *NullableRegistrationInlineHookPPData) Set(val *RegistrationInlineHookPPData) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookPPData) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookPPData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookPPData(val *RegistrationInlineHookPPData) *NullableRegistrationInlineHookPPData {
	return &NullableRegistrationInlineHookPPData{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookPPData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookPPData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
