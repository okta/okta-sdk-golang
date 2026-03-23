/*
Okta Admin Management API

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

// checks if the GroupQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupQueryResponse{}

// GroupQueryResponse struct for GroupQueryResponse
type GroupQueryResponse struct {
	// Identifier used to poll for the query result
	ResultId             *string `json:"resultId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupQueryResponse GroupQueryResponse

// NewGroupQueryResponse instantiates a new GroupQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupQueryResponse() *GroupQueryResponse {
	this := GroupQueryResponse{}
	return &this
}

// NewGroupQueryResponseWithDefaults instantiates a new GroupQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupQueryResponseWithDefaults() *GroupQueryResponse {
	this := GroupQueryResponse{}
	return &this
}

// GetResultId returns the ResultId field value if set, zero value otherwise.
func (o *GroupQueryResponse) GetResultId() string {
	if o == nil || IsNil(o.ResultId) {
		var ret string
		return ret
	}
	return *o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupQueryResponse) GetResultIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResultId) {
		return nil, false
	}
	return o.ResultId, true
}

// HasResultId returns a boolean if a field has been set.
func (o *GroupQueryResponse) HasResultId() bool {
	if o != nil && !IsNil(o.ResultId) {
		return true
	}

	return false
}

// SetResultId gets a reference to the given string and assigns it to the ResultId field.
func (o *GroupQueryResponse) SetResultId(v string) {
	o.ResultId = &v
}

func (o GroupQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultId) {
		toSerialize["resultId"] = o.ResultId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupQueryResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupQueryResponse := _GroupQueryResponse{}

	err = json.Unmarshal(data, &varGroupQueryResponse)

	if err != nil {
		return err
	}

	*o = GroupQueryResponse(varGroupQueryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resultId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupQueryResponse struct {
	value *GroupQueryResponse
	isSet bool
}

func (v NullableGroupQueryResponse) Get() *GroupQueryResponse {
	return v.value
}

func (v *NullableGroupQueryResponse) Set(val *GroupQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupQueryResponse(val *GroupQueryResponse) *NullableGroupQueryResponse {
	return &NullableGroupQueryResponse{value: val, isSet: true}
}

func (v NullableGroupQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
