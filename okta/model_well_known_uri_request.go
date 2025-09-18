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

// checks if the WellKnownURIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownURIRequest{}

// WellKnownURIRequest struct for WellKnownURIRequest
type WellKnownURIRequest struct {
	// The well-known URI content in JSON object format
	Representation       map[string]interface{} `json:"representation"`
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
		return map[string]interface{}{}, false
	}
	return o.Representation, true
}

// SetRepresentation sets field value
func (o *WellKnownURIRequest) SetRepresentation(v map[string]interface{}) {
	o.Representation = v
}

func (o WellKnownURIRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownURIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["representation"] = o.Representation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownURIRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"representation",
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

	varWellKnownURIRequest := _WellKnownURIRequest{}

	err = json.Unmarshal(data, &varWellKnownURIRequest)

	if err != nil {
		return err
	}

	*o = WellKnownURIRequest(varWellKnownURIRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "representation")
		o.AdditionalProperties = additionalProperties
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
