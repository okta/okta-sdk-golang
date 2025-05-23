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

// ResourceSetBindingResponse struct for ResourceSetBindingResponse
type ResourceSetBindingResponse struct {
	// `id` of the role
	Id *string `json:"id,omitempty"`
	Links *ResourceSetBindingResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingResponse ResourceSetBindingResponse

// NewResourceSetBindingResponse instantiates a new ResourceSetBindingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingResponse() *ResourceSetBindingResponse {
	this := ResourceSetBindingResponse{}
	return &this
}

// NewResourceSetBindingResponseWithDefaults instantiates a new ResourceSetBindingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingResponseWithDefaults() *ResourceSetBindingResponse {
	this := ResourceSetBindingResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceSetBindingResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceSetBindingResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceSetBindingResponse) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetBindingResponse) GetLinks() ResourceSetBindingResponseLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSetBindingResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponse) GetLinksOk() (*ResourceSetBindingResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetBindingResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetBindingResponseLinks and assigns it to the Links field.
func (o *ResourceSetBindingResponse) SetLinks(v ResourceSetBindingResponseLinks) {
	o.Links = &v
}

func (o ResourceSetBindingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingResponse) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingResponse := _ResourceSetBindingResponse{}

	err = json.Unmarshal(bytes, &varResourceSetBindingResponse)
	if err == nil {
		*o = ResourceSetBindingResponse(varResourceSetBindingResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingResponse struct {
	value *ResourceSetBindingResponse
	isSet bool
}

func (v NullableResourceSetBindingResponse) Get() *ResourceSetBindingResponse {
	return v.value
}

func (v *NullableResourceSetBindingResponse) Set(val *ResourceSetBindingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingResponse(val *ResourceSetBindingResponse) *NullableResourceSetBindingResponse {
	return &NullableResourceSetBindingResponse{value: val, isSet: true}
}

func (v NullableResourceSetBindingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

