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

// checks if the RegistrationInlineHookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookRequest{}

// RegistrationInlineHookRequest Registration inline hook request
type RegistrationInlineHookRequest struct {
	// The type of inline hook. The registration inline hook type is `com.okta.user.pre-registration`.
	EventType *string `json:"eventType,omitempty"`
	// The type of registration hook. Use either `self.service.registration` or `progressive.profile`.
	RequestType *string `json:"requestType,omitempty"`
	// The ID of the registration inline hook
	Source               *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookRequest RegistrationInlineHookRequest

// NewRegistrationInlineHookRequest instantiates a new RegistrationInlineHookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookRequest() *RegistrationInlineHookRequest {
	this := RegistrationInlineHookRequest{}
	return &this
}

// NewRegistrationInlineHookRequestWithDefaults instantiates a new RegistrationInlineHookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookRequestWithDefaults() *RegistrationInlineHookRequest {
	this := RegistrationInlineHookRequest{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RegistrationInlineHookRequest) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RegistrationInlineHookRequest) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *RegistrationInlineHookRequest) SetEventType(v string) {
	o.EventType = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *RegistrationInlineHookRequest) GetRequestType() string {
	if o == nil || IsNil(o.RequestType) {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookRequest) GetRequestTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *RegistrationInlineHookRequest) HasRequestType() bool {
	if o != nil && !IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *RegistrationInlineHookRequest) SetRequestType(v string) {
	o.RequestType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RegistrationInlineHookRequest) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookRequest) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RegistrationInlineHookRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RegistrationInlineHookRequest) SetSource(v string) {
	o.Source = &v
}

func (o RegistrationInlineHookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.RequestType) {
		toSerialize["requestType"] = o.RequestType
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookRequest) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookRequest := _RegistrationInlineHookRequest{}

	err = json.Unmarshal(data, &varRegistrationInlineHookRequest)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookRequest(varRegistrationInlineHookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookRequest struct {
	value *RegistrationInlineHookRequest
	isSet bool
}

func (v NullableRegistrationInlineHookRequest) Get() *RegistrationInlineHookRequest {
	return v.value
}

func (v *NullableRegistrationInlineHookRequest) Set(val *RegistrationInlineHookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookRequest(val *RegistrationInlineHookRequest) *NullableRegistrationInlineHookRequest {
	return &NullableRegistrationInlineHookRequest{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
