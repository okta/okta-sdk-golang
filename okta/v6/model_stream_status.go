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

// StreamStatus Status corresponding to the `stream_id` of the SSF Stream
type StreamStatus struct {
	// The status of the SSF Stream configuration
	Status *string `json:"status,omitempty"`
	// The ID of the SSF Stream configuration. This corresponds to the value in the query parameter of the request.
	StreamId *string `json:"stream_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StreamStatus StreamStatus

// NewStreamStatus instantiates a new StreamStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamStatus() *StreamStatus {
	this := StreamStatus{}
	return &this
}

// NewStreamStatusWithDefaults instantiates a new StreamStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamStatusWithDefaults() *StreamStatus {
	this := StreamStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StreamStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StreamStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StreamStatus) SetStatus(v string) {
	o.Status = &v
}

// GetStreamId returns the StreamId field value if set, zero value otherwise.
func (o *StreamStatus) GetStreamId() string {
	if o == nil || o.StreamId == nil {
		var ret string
		return ret
	}
	return *o.StreamId
}

// GetStreamIdOk returns a tuple with the StreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamStatus) GetStreamIdOk() (*string, bool) {
	if o == nil || o.StreamId == nil {
		return nil, false
	}
	return o.StreamId, true
}

// HasStreamId returns a boolean if a field has been set.
func (o *StreamStatus) HasStreamId() bool {
	if o != nil && o.StreamId != nil {
		return true
	}

	return false
}

// SetStreamId gets a reference to the given string and assigns it to the StreamId field.
func (o *StreamStatus) SetStreamId(v string) {
	o.StreamId = &v
}

func (o StreamStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StreamId != nil {
		toSerialize["stream_id"] = o.StreamId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StreamStatus) UnmarshalJSON(bytes []byte) (err error) {
	varStreamStatus := _StreamStatus{}

	err = json.Unmarshal(bytes, &varStreamStatus)
	if err == nil {
		*o = StreamStatus(varStreamStatus)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "stream_id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableStreamStatus struct {
	value *StreamStatus
	isSet bool
}

func (v NullableStreamStatus) Get() *StreamStatus {
	return v.value
}

func (v *NullableStreamStatus) Set(val *StreamStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamStatus(val *StreamStatus) *NullableStreamStatus {
	return &NullableStreamStatus{value: val, isSet: true}
}

func (v NullableStreamStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

