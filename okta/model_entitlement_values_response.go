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

// checks if the EntitlementValuesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementValuesResponse{}

// EntitlementValuesResponse struct for EntitlementValuesResponse
type EntitlementValuesResponse struct {
	EntitlementValues    []EntitlementValue              `json:"entitlementValues,omitempty"`
	Links                *EntitlementValuesResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValuesResponse EntitlementValuesResponse

// NewEntitlementValuesResponse instantiates a new EntitlementValuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValuesResponse() *EntitlementValuesResponse {
	this := EntitlementValuesResponse{}
	return &this
}

// NewEntitlementValuesResponseWithDefaults instantiates a new EntitlementValuesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValuesResponseWithDefaults() *EntitlementValuesResponse {
	this := EntitlementValuesResponse{}
	return &this
}

// GetEntitlementValues returns the EntitlementValues field value if set, zero value otherwise.
func (o *EntitlementValuesResponse) GetEntitlementValues() []EntitlementValue {
	if o == nil || IsNil(o.EntitlementValues) {
		var ret []EntitlementValue
		return ret
	}
	return o.EntitlementValues
}

// GetEntitlementValuesOk returns a tuple with the EntitlementValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponse) GetEntitlementValuesOk() ([]EntitlementValue, bool) {
	if o == nil || IsNil(o.EntitlementValues) {
		return nil, false
	}
	return o.EntitlementValues, true
}

// HasEntitlementValues returns a boolean if a field has been set.
func (o *EntitlementValuesResponse) HasEntitlementValues() bool {
	if o != nil && !IsNil(o.EntitlementValues) {
		return true
	}

	return false
}

// SetEntitlementValues gets a reference to the given []EntitlementValue and assigns it to the EntitlementValues field.
func (o *EntitlementValuesResponse) SetEntitlementValues(v []EntitlementValue) {
	o.EntitlementValues = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EntitlementValuesResponse) GetLinks() EntitlementValuesResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret EntitlementValuesResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponse) GetLinksOk() (*EntitlementValuesResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EntitlementValuesResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EntitlementValuesResponseLinks and assigns it to the Links field.
func (o *EntitlementValuesResponse) SetLinks(v EntitlementValuesResponseLinks) {
	o.Links = &v
}

func (o EntitlementValuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementValuesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntitlementValues) {
		toSerialize["entitlementValues"] = o.EntitlementValues
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementValuesResponse) UnmarshalJSON(data []byte) (err error) {
	varEntitlementValuesResponse := _EntitlementValuesResponse{}

	err = json.Unmarshal(data, &varEntitlementValuesResponse)

	if err != nil {
		return err
	}

	*o = EntitlementValuesResponse(varEntitlementValuesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementValues")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementValuesResponse struct {
	value *EntitlementValuesResponse
	isSet bool
}

func (v NullableEntitlementValuesResponse) Get() *EntitlementValuesResponse {
	return v.value
}

func (v *NullableEntitlementValuesResponse) Set(val *EntitlementValuesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuesResponse(val *EntitlementValuesResponse) *NullableEntitlementValuesResponse {
	return &NullableEntitlementValuesResponse{value: val, isSet: true}
}

func (v NullableEntitlementValuesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
