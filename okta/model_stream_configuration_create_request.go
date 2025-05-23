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

// StreamConfigurationCreateRequest struct for StreamConfigurationCreateRequest
type StreamConfigurationCreateRequest struct {
	Delivery StreamConfigurationDelivery `json:"delivery"`
	// The events (mapped by the array of event type URIs) that the receiver wants to receive
	EventsRequested []string `json:"events_requested"`
	// The Subject Identifier format expected for any SET transmitted.
	Format *string `json:"format,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StreamConfigurationCreateRequest StreamConfigurationCreateRequest

// NewStreamConfigurationCreateRequest instantiates a new StreamConfigurationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamConfigurationCreateRequest(delivery StreamConfigurationDelivery, eventsRequested []string) *StreamConfigurationCreateRequest {
	this := StreamConfigurationCreateRequest{}
	this.Delivery = delivery
	this.EventsRequested = eventsRequested
	return &this
}

// NewStreamConfigurationCreateRequestWithDefaults instantiates a new StreamConfigurationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamConfigurationCreateRequestWithDefaults() *StreamConfigurationCreateRequest {
	this := StreamConfigurationCreateRequest{}
	return &this
}

// GetDelivery returns the Delivery field value
func (o *StreamConfigurationCreateRequest) GetDelivery() StreamConfigurationDelivery {
	if o == nil {
		var ret StreamConfigurationDelivery
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *StreamConfigurationCreateRequest) GetDeliveryOk() (*StreamConfigurationDelivery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *StreamConfigurationCreateRequest) SetDelivery(v StreamConfigurationDelivery) {
	o.Delivery = v
}

// GetEventsRequested returns the EventsRequested field value
func (o *StreamConfigurationCreateRequest) GetEventsRequested() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventsRequested
}

// GetEventsRequestedOk returns a tuple with the EventsRequested field value
// and a boolean to check if the value has been set.
func (o *StreamConfigurationCreateRequest) GetEventsRequestedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsRequested, true
}

// SetEventsRequested sets field value
func (o *StreamConfigurationCreateRequest) SetEventsRequested(v []string) {
	o.EventsRequested = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *StreamConfigurationCreateRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfigurationCreateRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *StreamConfigurationCreateRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *StreamConfigurationCreateRequest) SetFormat(v string) {
	o.Format = &v
}

func (o StreamConfigurationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["delivery"] = o.Delivery
	}
	if true {
		toSerialize["events_requested"] = o.EventsRequested
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StreamConfigurationCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varStreamConfigurationCreateRequest := _StreamConfigurationCreateRequest{}

	err = json.Unmarshal(bytes, &varStreamConfigurationCreateRequest)
	if err == nil {
		*o = StreamConfigurationCreateRequest(varStreamConfigurationCreateRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "delivery")
		delete(additionalProperties, "events_requested")
		delete(additionalProperties, "format")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableStreamConfigurationCreateRequest struct {
	value *StreamConfigurationCreateRequest
	isSet bool
}

func (v NullableStreamConfigurationCreateRequest) Get() *StreamConfigurationCreateRequest {
	return v.value
}

func (v *NullableStreamConfigurationCreateRequest) Set(val *StreamConfigurationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamConfigurationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamConfigurationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamConfigurationCreateRequest(val *StreamConfigurationCreateRequest) *NullableStreamConfigurationCreateRequest {
	return &NullableStreamConfigurationCreateRequest{value: val, isSet: true}
}

func (v NullableStreamConfigurationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamConfigurationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

