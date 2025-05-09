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

// TelephonyRequest struct for TelephonyRequest
type TelephonyRequest struct {
	Data *TelephonyRequestData `json:"data,omitempty"`
	// The type of inline hook. The Telephony inline hook type is `com.okta.telephony.provider`.
	EventType *string `json:"eventType,omitempty"`
	// The type of inline hook request. For example, `com.okta.user.telephony.pre-enrollment`.
	RequestType *string `json:"requestType,omitempty"`
	// The ID and URL of the Telephony inline hook
	Source *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyRequest TelephonyRequest

// NewTelephonyRequest instantiates a new TelephonyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyRequest() *TelephonyRequest {
	this := TelephonyRequest{}
	return &this
}

// NewTelephonyRequestWithDefaults instantiates a new TelephonyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyRequestWithDefaults() *TelephonyRequest {
	this := TelephonyRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TelephonyRequest) GetData() TelephonyRequestData {
	if o == nil || o.Data == nil {
		var ret TelephonyRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequest) GetDataOk() (*TelephonyRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TelephonyRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TelephonyRequestData and assigns it to the Data field.
func (o *TelephonyRequest) SetData(v TelephonyRequestData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *TelephonyRequest) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *TelephonyRequest) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *TelephonyRequest) SetEventType(v string) {
	o.EventType = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *TelephonyRequest) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequest) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *TelephonyRequest) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *TelephonyRequest) SetRequestType(v string) {
	o.RequestType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TelephonyRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TelephonyRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TelephonyRequest) SetSource(v string) {
	o.Source = &v
}

func (o TelephonyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *TelephonyRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyRequest := _TelephonyRequest{}

	err = json.Unmarshal(bytes, &varTelephonyRequest)
	if err == nil {
		*o = TelephonyRequest(varTelephonyRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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

type NullableTelephonyRequest struct {
	value *TelephonyRequest
	isSet bool
}

func (v NullableTelephonyRequest) Get() *TelephonyRequest {
	return v.value
}

func (v *NullableTelephonyRequest) Set(val *TelephonyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyRequest(val *TelephonyRequest) *NullableTelephonyRequest {
	return &NullableTelephonyRequest{value: val, isSet: true}
}

func (v NullableTelephonyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

