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

// StreamConfiguration struct for StreamConfiguration
type StreamConfiguration struct {
	Aud *StreamConfigurationAud `json:"aud,omitempty"`
	Delivery StreamConfigurationDelivery `json:"delivery"`
	// The events (mapped by the array of event type URIs) that the transmitter actually delivers to the SSF Stream.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter.
	EventsDelivered []string `json:"events_delivered,omitempty"`
	// The events (mapped by the array of event type URIs) that the receiver wants to receive
	EventsRequested []string `json:"events_requested"`
	// An array of event type URIs that the transmitter supports.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter.
	EventsSupported []string `json:"events_supported,omitempty"`
	// The Subject Identifier format expected for any SET transmitted.
	Format *string `json:"format,omitempty"`
	// The issuer used in Security Event Tokens (SETs). This value is set as `iss` in the claim.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter.
	Iss *string `json:"iss,omitempty"`
	// The minimum amount of time, in seconds, between two verification requests.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter.
	MinVerificationInterval NullableInt32 `json:"min_verification_interval,omitempty"`
	// The ID of the SSF Stream configuration
	StreamId *string `json:"stream_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StreamConfiguration StreamConfiguration

// NewStreamConfiguration instantiates a new StreamConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamConfiguration(delivery StreamConfigurationDelivery, eventsRequested []string) *StreamConfiguration {
	this := StreamConfiguration{}
	this.Delivery = delivery
	this.EventsRequested = eventsRequested
	return &this
}

// NewStreamConfigurationWithDefaults instantiates a new StreamConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamConfigurationWithDefaults() *StreamConfiguration {
	this := StreamConfiguration{}
	return &this
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *StreamConfiguration) GetAud() StreamConfigurationAud {
	if o == nil || o.Aud == nil {
		var ret StreamConfigurationAud
		return ret
	}
	return *o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetAudOk() (*StreamConfigurationAud, bool) {
	if o == nil || o.Aud == nil {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *StreamConfiguration) HasAud() bool {
	if o != nil && o.Aud != nil {
		return true
	}

	return false
}

// SetAud gets a reference to the given StreamConfigurationAud and assigns it to the Aud field.
func (o *StreamConfiguration) SetAud(v StreamConfigurationAud) {
	o.Aud = &v
}

// GetDelivery returns the Delivery field value
func (o *StreamConfiguration) GetDelivery() StreamConfigurationDelivery {
	if o == nil {
		var ret StreamConfigurationDelivery
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetDeliveryOk() (*StreamConfigurationDelivery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *StreamConfiguration) SetDelivery(v StreamConfigurationDelivery) {
	o.Delivery = v
}

// GetEventsDelivered returns the EventsDelivered field value if set, zero value otherwise.
func (o *StreamConfiguration) GetEventsDelivered() []string {
	if o == nil || o.EventsDelivered == nil {
		var ret []string
		return ret
	}
	return o.EventsDelivered
}

// GetEventsDeliveredOk returns a tuple with the EventsDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetEventsDeliveredOk() ([]string, bool) {
	if o == nil || o.EventsDelivered == nil {
		return nil, false
	}
	return o.EventsDelivered, true
}

// HasEventsDelivered returns a boolean if a field has been set.
func (o *StreamConfiguration) HasEventsDelivered() bool {
	if o != nil && o.EventsDelivered != nil {
		return true
	}

	return false
}

// SetEventsDelivered gets a reference to the given []string and assigns it to the EventsDelivered field.
func (o *StreamConfiguration) SetEventsDelivered(v []string) {
	o.EventsDelivered = v
}

// GetEventsRequested returns the EventsRequested field value
func (o *StreamConfiguration) GetEventsRequested() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EventsRequested
}

// GetEventsRequestedOk returns a tuple with the EventsRequested field value
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetEventsRequestedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventsRequested, true
}

// SetEventsRequested sets field value
func (o *StreamConfiguration) SetEventsRequested(v []string) {
	o.EventsRequested = v
}

// GetEventsSupported returns the EventsSupported field value if set, zero value otherwise.
func (o *StreamConfiguration) GetEventsSupported() []string {
	if o == nil || o.EventsSupported == nil {
		var ret []string
		return ret
	}
	return o.EventsSupported
}

// GetEventsSupportedOk returns a tuple with the EventsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetEventsSupportedOk() ([]string, bool) {
	if o == nil || o.EventsSupported == nil {
		return nil, false
	}
	return o.EventsSupported, true
}

// HasEventsSupported returns a boolean if a field has been set.
func (o *StreamConfiguration) HasEventsSupported() bool {
	if o != nil && o.EventsSupported != nil {
		return true
	}

	return false
}

// SetEventsSupported gets a reference to the given []string and assigns it to the EventsSupported field.
func (o *StreamConfiguration) SetEventsSupported(v []string) {
	o.EventsSupported = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *StreamConfiguration) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *StreamConfiguration) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *StreamConfiguration) SetFormat(v string) {
	o.Format = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *StreamConfiguration) GetIss() string {
	if o == nil || o.Iss == nil {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetIssOk() (*string, bool) {
	if o == nil || o.Iss == nil {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *StreamConfiguration) HasIss() bool {
	if o != nil && o.Iss != nil {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *StreamConfiguration) SetIss(v string) {
	o.Iss = &v
}

// GetMinVerificationInterval returns the MinVerificationInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConfiguration) GetMinVerificationInterval() int32 {
	if o == nil || o.MinVerificationInterval.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinVerificationInterval.Get()
}

// GetMinVerificationIntervalOk returns a tuple with the MinVerificationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConfiguration) GetMinVerificationIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinVerificationInterval.Get(), o.MinVerificationInterval.IsSet()
}

// HasMinVerificationInterval returns a boolean if a field has been set.
func (o *StreamConfiguration) HasMinVerificationInterval() bool {
	if o != nil && o.MinVerificationInterval.IsSet() {
		return true
	}

	return false
}

// SetMinVerificationInterval gets a reference to the given NullableInt32 and assigns it to the MinVerificationInterval field.
func (o *StreamConfiguration) SetMinVerificationInterval(v int32) {
	o.MinVerificationInterval.Set(&v)
}
// SetMinVerificationIntervalNil sets the value for MinVerificationInterval to be an explicit nil
func (o *StreamConfiguration) SetMinVerificationIntervalNil() {
	o.MinVerificationInterval.Set(nil)
}

// UnsetMinVerificationInterval ensures that no value is present for MinVerificationInterval, not even an explicit nil
func (o *StreamConfiguration) UnsetMinVerificationInterval() {
	o.MinVerificationInterval.Unset()
}

// GetStreamId returns the StreamId field value if set, zero value otherwise.
func (o *StreamConfiguration) GetStreamId() string {
	if o == nil || o.StreamId == nil {
		var ret string
		return ret
	}
	return *o.StreamId
}

// GetStreamIdOk returns a tuple with the StreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamConfiguration) GetStreamIdOk() (*string, bool) {
	if o == nil || o.StreamId == nil {
		return nil, false
	}
	return o.StreamId, true
}

// HasStreamId returns a boolean if a field has been set.
func (o *StreamConfiguration) HasStreamId() bool {
	if o != nil && o.StreamId != nil {
		return true
	}

	return false
}

// SetStreamId gets a reference to the given string and assigns it to the StreamId field.
func (o *StreamConfiguration) SetStreamId(v string) {
	o.StreamId = &v
}

func (o StreamConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aud != nil {
		toSerialize["aud"] = o.Aud
	}
	if true {
		toSerialize["delivery"] = o.Delivery
	}
	if o.EventsDelivered != nil {
		toSerialize["events_delivered"] = o.EventsDelivered
	}
	if true {
		toSerialize["events_requested"] = o.EventsRequested
	}
	if o.EventsSupported != nil {
		toSerialize["events_supported"] = o.EventsSupported
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.MinVerificationInterval.IsSet() {
		toSerialize["min_verification_interval"] = o.MinVerificationInterval.Get()
	}
	if o.StreamId != nil {
		toSerialize["stream_id"] = o.StreamId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StreamConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varStreamConfiguration := _StreamConfiguration{}

	err = json.Unmarshal(bytes, &varStreamConfiguration)
	if err == nil {
		*o = StreamConfiguration(varStreamConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "aud")
		delete(additionalProperties, "delivery")
		delete(additionalProperties, "events_delivered")
		delete(additionalProperties, "events_requested")
		delete(additionalProperties, "events_supported")
		delete(additionalProperties, "format")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "min_verification_interval")
		delete(additionalProperties, "stream_id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableStreamConfiguration struct {
	value *StreamConfiguration
	isSet bool
}

func (v NullableStreamConfiguration) Get() *StreamConfiguration {
	return v.value
}

func (v *NullableStreamConfiguration) Set(val *StreamConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamConfiguration(val *StreamConfiguration) *NullableStreamConfiguration {
	return &NullableStreamConfiguration{value: val, isSet: true}
}

func (v NullableStreamConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

