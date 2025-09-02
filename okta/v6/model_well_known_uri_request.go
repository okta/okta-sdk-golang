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
	"fmt"
)

// WellKnownURIRequest struct for WellKnownURIRequest
type WellKnownURIRequest struct {
	// The well-known URI content in JSON object format
	Representation map[string]interface{} `json:"representation"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownURIRequest WellKnownURIRequest

// NewWellKnownURIRequest instantiates a new WellKnownURIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIRequest(representation map[string]interface{}) *WellKnownURIRequest {
	this := WellKnownURIRequest{}
	this.Representation = representation
	return &this
}

// NewWellKnownURIRequestWithDefaults instantiates a new WellKnownURIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIRequestWithDefaults() *WellKnownURIRequest {
	this := WellKnownURIRequest{}
	return &this
}

// GetRepresentation returns the Representation field value
func (o *WellKnownURIRequest) GetRepresentation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Representation
}

// GetRepresentationOk returns a tuple with the Representation field value
// and a boolean to check if the value has been set.
func (o *WellKnownURIRequest) GetRepresentationOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Representation, true
}

// SetRepresentation sets field value
func (o *WellKnownURIRequest) SetRepresentation(v map[string]interface{}) {
	o.Representation = v
}

func (o WellKnownURIRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["representation"] = o.Representation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownURIRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownURIRequest := _WellKnownURIRequest{}

	err = json.Unmarshal(bytes, &varWellKnownURIRequest)
	if err == nil {
		*o = WellKnownURIRequest(varWellKnownURIRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "representation")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWellKnownURIRequest struct {
	value *WellKnownURIRequest
	isSet bool
}

func (v NullableWellKnownURIRequest) Get() *WellKnownURIRequest {
	return v.value
}

func (v *NullableWellKnownURIRequest) Set(val *WellKnownURIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIRequest(val *WellKnownURIRequest) *NullableWellKnownURIRequest {
	return &NullableWellKnownURIRequest{value: val, isSet: true}
}

func (v NullableWellKnownURIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

