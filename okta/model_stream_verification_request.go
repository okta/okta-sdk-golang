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

// checks if the StreamVerificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamVerificationRequest{}

// StreamVerificationRequest struct for StreamVerificationRequest
type StreamVerificationRequest struct {
	// An arbitrary string that Okta as a transmitter must echo back to the Event Receiver in the Verification Event's payload
	State *string `json:"state,omitempty"`
	// The ID of the SSF Stream Configuration
	StreamId             string `json:"stream_id"`
	AdditionalProperties map[string]interface{}
}

type _StreamVerificationRequest StreamVerificationRequest

// NewStreamVerificationRequest instantiates a new StreamVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamVerificationRequest(streamId string) *StreamVerificationRequest {
	this := StreamVerificationRequest{}
	this.StreamId = streamId
	return &this
}

// NewStreamVerificationRequestWithDefaults instantiates a new StreamVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamVerificationRequestWithDefaults() *StreamVerificationRequest {
	this := StreamVerificationRequest{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StreamVerificationRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamVerificationRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StreamVerificationRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StreamVerificationRequest) SetState(v string) {
	o.State = &v
}

// GetStreamId returns the StreamId field value
func (o *StreamVerificationRequest) GetStreamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamId
}

// GetStreamIdOk returns a tuple with the StreamId field value
// and a boolean to check if the value has been set.
func (o *StreamVerificationRequest) GetStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamId, true
}

// SetStreamId sets field value
func (o *StreamVerificationRequest) SetStreamId(v string) {
	o.StreamId = v
}

func (o StreamVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["stream_id"] = o.StreamId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StreamVerificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stream_id",
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

	varStreamVerificationRequest := _StreamVerificationRequest{}

	err = json.Unmarshal(data, &varStreamVerificationRequest)

	if err != nil {
		return err
	}

	*o = StreamVerificationRequest(varStreamVerificationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "stream_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStreamVerificationRequest struct {
	value *StreamVerificationRequest
	isSet bool
}

func (v NullableStreamVerificationRequest) Get() *StreamVerificationRequest {
	return v.value
}

func (v *NullableStreamVerificationRequest) Set(val *StreamVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamVerificationRequest(val *StreamVerificationRequest) *NullableStreamVerificationRequest {
	return &NullableStreamVerificationRequest{value: val, isSet: true}
}

func (v NullableStreamVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
