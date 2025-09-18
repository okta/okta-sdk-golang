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
)

// checks if the TokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRequest{}

// TokenRequest Token inline hook request
type TokenRequest struct {
	// The inline hook cloud version
	CloudEventVersion *string `json:"cloudEventVersion,omitempty"`
	// The inline hook request header content
	ContentType *string `json:"contentType,omitempty"`
	// The individual inline hook request ID
	EventId *string `json:"eventId,omitempty"`
	// The time the inline hook request was sent
	EventTime *string `json:"eventTime,omitempty"`
	// The inline hook version
	EventTypeVersion *string           `json:"eventTypeVersion,omitempty"`
	Data             *TokenPayLoadData `json:"data,omitempty"`
	// The type of inline hook. The token inline hook type is `com.okta.oauth2.tokens.transform`.
	EventType *string `json:"eventType,omitempty"`
	// The URL of the token inline hook
	Source               *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenRequest TokenRequest

// NewTokenRequest instantiates a new TokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequest() *TokenRequest {
	this := TokenRequest{}
	return &this
}

// NewTokenRequestWithDefaults instantiates a new TokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestWithDefaults() *TokenRequest {
	this := TokenRequest{}
	return &this
}

// GetCloudEventVersion returns the CloudEventVersion field value if set, zero value otherwise.
func (o *TokenRequest) GetCloudEventVersion() string {
	if o == nil || IsNil(o.CloudEventVersion) {
		var ret string
		return ret
	}
	return *o.CloudEventVersion
}

// GetCloudEventVersionOk returns a tuple with the CloudEventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetCloudEventVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CloudEventVersion) {
		return nil, false
	}
	return o.CloudEventVersion, true
}

// HasCloudEventVersion returns a boolean if a field has been set.
func (o *TokenRequest) HasCloudEventVersion() bool {
	if o != nil && !IsNil(o.CloudEventVersion) {
		return true
	}

	return false
}

// SetCloudEventVersion gets a reference to the given string and assigns it to the CloudEventVersion field.
func (o *TokenRequest) SetCloudEventVersion(v string) {
	o.CloudEventVersion = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *TokenRequest) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *TokenRequest) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *TokenRequest) SetContentType(v string) {
	o.ContentType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *TokenRequest) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *TokenRequest) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *TokenRequest) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *TokenRequest) GetEventTime() string {
	if o == nil || IsNil(o.EventTime) {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *TokenRequest) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *TokenRequest) SetEventTime(v string) {
	o.EventTime = &v
}

// GetEventTypeVersion returns the EventTypeVersion field value if set, zero value otherwise.
func (o *TokenRequest) GetEventTypeVersion() string {
	if o == nil || IsNil(o.EventTypeVersion) {
		var ret string
		return ret
	}
	return *o.EventTypeVersion
}

// GetEventTypeVersionOk returns a tuple with the EventTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetEventTypeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeVersion) {
		return nil, false
	}
	return o.EventTypeVersion, true
}

// HasEventTypeVersion returns a boolean if a field has been set.
func (o *TokenRequest) HasEventTypeVersion() bool {
	if o != nil && !IsNil(o.EventTypeVersion) {
		return true
	}

	return false
}

// SetEventTypeVersion gets a reference to the given string and assigns it to the EventTypeVersion field.
func (o *TokenRequest) SetEventTypeVersion(v string) {
	o.EventTypeVersion = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TokenRequest) GetData() TokenPayLoadData {
	if o == nil || IsNil(o.Data) {
		var ret TokenPayLoadData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetDataOk() (*TokenPayLoadData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TokenRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TokenPayLoadData and assigns it to the Data field.
func (o *TokenRequest) SetData(v TokenPayLoadData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *TokenRequest) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *TokenRequest) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *TokenRequest) SetEventType(v string) {
	o.EventType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TokenRequest) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TokenRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TokenRequest) SetSource(v string) {
	o.Source = &v
}

func (o TokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRequest) ToMap() (map[string]interface{}, error) {
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

func (o *TokenRequest) UnmarshalJSON(data []byte) (err error) {
	varTokenRequest := _TokenRequest{}

	err = json.Unmarshal(data, &varTokenRequest)

	if err != nil {
		return err
	}

	*o = TokenRequest(varTokenRequest)

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

type NullableTokenRequest struct {
	value *TokenRequest
	isSet bool
}

func (v NullableTokenRequest) Get() *TokenRequest {
	return v.value
}

func (v *NullableTokenRequest) Set(val *TokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequest(val *TokenRequest) *NullableTokenRequest {
	return &NullableTokenRequest{value: val, isSet: true}
}

func (v NullableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
