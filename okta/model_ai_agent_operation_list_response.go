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

// checks if the AIAgentOperationListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIAgentOperationListResponse{}

// AIAgentOperationListResponse struct for AIAgentOperationListResponse
type AIAgentOperationListResponse struct {
	Data                 []AIAgentOperationResponse         `json:"data,omitempty"`
	Links                *AIAgentOperationListResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIAgentOperationListResponse AIAgentOperationListResponse

// NewAIAgentOperationListResponse instantiates a new AIAgentOperationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIAgentOperationListResponse() *AIAgentOperationListResponse {
	this := AIAgentOperationListResponse{}
	return &this
}

// NewAIAgentOperationListResponseWithDefaults instantiates a new AIAgentOperationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIAgentOperationListResponseWithDefaults() *AIAgentOperationListResponse {
	this := AIAgentOperationListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AIAgentOperationListResponse) GetData() []AIAgentOperationResponse {
	if o == nil || IsNil(o.Data) {
		var ret []AIAgentOperationResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationListResponse) GetDataOk() ([]AIAgentOperationResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AIAgentOperationListResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AIAgentOperationResponse and assigns it to the Data field.
func (o *AIAgentOperationListResponse) SetData(v []AIAgentOperationResponse) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AIAgentOperationListResponse) GetLinks() AIAgentOperationListResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret AIAgentOperationListResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationListResponse) GetLinksOk() (*AIAgentOperationListResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AIAgentOperationListResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AIAgentOperationListResponseLinks and assigns it to the Links field.
func (o *AIAgentOperationListResponse) SetLinks(v AIAgentOperationListResponseLinks) {
	o.Links = &v
}

func (o AIAgentOperationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIAgentOperationListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIAgentOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	varAIAgentOperationListResponse := _AIAgentOperationListResponse{}

	err = json.Unmarshal(data, &varAIAgentOperationListResponse)

	if err != nil {
		return err
	}

	*o = AIAgentOperationListResponse(varAIAgentOperationListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIAgentOperationListResponse struct {
	value *AIAgentOperationListResponse
	isSet bool
}

func (v NullableAIAgentOperationListResponse) Get() *AIAgentOperationListResponse {
	return v.value
}

func (v *NullableAIAgentOperationListResponse) Set(val *AIAgentOperationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAIAgentOperationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAIAgentOperationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIAgentOperationListResponse(val *AIAgentOperationListResponse) *NullableAIAgentOperationListResponse {
	return &NullableAIAgentOperationListResponse{value: val, isSet: true}
}

func (v NullableAIAgentOperationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIAgentOperationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
