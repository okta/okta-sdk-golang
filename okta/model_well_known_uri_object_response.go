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
)

// checks if the WellKnownURIObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownURIObjectResponse{}

// WellKnownURIObjectResponse struct for WellKnownURIObjectResponse
type WellKnownURIObjectResponse struct {
	// The well-known URI content in JSON format
	Representation       map[string]interface{}          `json:"representation,omitempty"`
	Links                *WellKnownURIArrayResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownURIObjectResponse WellKnownURIObjectResponse

// NewWellKnownURIObjectResponse instantiates a new WellKnownURIObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIObjectResponse() *WellKnownURIObjectResponse {
	this := WellKnownURIObjectResponse{}
	return &this
}

// NewWellKnownURIObjectResponseWithDefaults instantiates a new WellKnownURIObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIObjectResponseWithDefaults() *WellKnownURIObjectResponse {
	this := WellKnownURIObjectResponse{}
	return &this
}

// GetRepresentation returns the Representation field value if set, zero value otherwise.
func (o *WellKnownURIObjectResponse) GetRepresentation() map[string]interface{} {
	if o == nil || IsNil(o.Representation) {
		var ret map[string]interface{}
		return ret
	}
	return o.Representation
}

// GetRepresentationOk returns a tuple with the Representation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIObjectResponse) GetRepresentationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Representation) {
		return map[string]interface{}{}, false
	}
	return o.Representation, true
}

// HasRepresentation returns a boolean if a field has been set.
func (o *WellKnownURIObjectResponse) HasRepresentation() bool {
	if o != nil && !IsNil(o.Representation) {
		return true
	}

	return false
}

// SetRepresentation gets a reference to the given map[string]interface{} and assigns it to the Representation field.
func (o *WellKnownURIObjectResponse) SetRepresentation(v map[string]interface{}) {
	o.Representation = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WellKnownURIObjectResponse) GetLinks() WellKnownURIArrayResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret WellKnownURIArrayResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIObjectResponse) GetLinksOk() (*WellKnownURIArrayResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WellKnownURIObjectResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WellKnownURIArrayResponseLinks and assigns it to the Links field.
func (o *WellKnownURIObjectResponse) SetLinks(v WellKnownURIArrayResponseLinks) {
	o.Links = &v
}

func (o WellKnownURIObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownURIObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Representation) {
		toSerialize["representation"] = o.Representation
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownURIObjectResponse) UnmarshalJSON(data []byte) (err error) {
	varWellKnownURIObjectResponse := _WellKnownURIObjectResponse{}

	err = json.Unmarshal(data, &varWellKnownURIObjectResponse)

	if err != nil {
		return err
	}

	*o = WellKnownURIObjectResponse(varWellKnownURIObjectResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "representation")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownURIObjectResponse struct {
	value *WellKnownURIObjectResponse
	isSet bool
}

func (v NullableWellKnownURIObjectResponse) Get() *WellKnownURIObjectResponse {
	return v.value
}

func (v *NullableWellKnownURIObjectResponse) Set(val *WellKnownURIObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIObjectResponse(val *WellKnownURIObjectResponse) *NullableWellKnownURIObjectResponse {
	return &NullableWellKnownURIObjectResponse{value: val, isSet: true}
}

func (v NullableWellKnownURIObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
