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
)

// ResourceSetBindingEditResponse struct for ResourceSetBindingEditResponse
type ResourceSetBindingEditResponse struct {
	Links *ResourceSetBindingEditResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingEditResponse ResourceSetBindingEditResponse

// NewResourceSetBindingEditResponse instantiates a new ResourceSetBindingEditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingEditResponse() *ResourceSetBindingEditResponse {
	this := ResourceSetBindingEditResponse{}
	return &this
}

// NewResourceSetBindingEditResponseWithDefaults instantiates a new ResourceSetBindingEditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingEditResponseWithDefaults() *ResourceSetBindingEditResponse {
	this := ResourceSetBindingEditResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetBindingEditResponse) GetLinks() ResourceSetBindingEditResponseLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSetBindingEditResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingEditResponse) GetLinksOk() (*ResourceSetBindingEditResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetBindingEditResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetBindingEditResponseLinks and assigns it to the Links field.
func (o *ResourceSetBindingEditResponse) SetLinks(v ResourceSetBindingEditResponseLinks) {
	o.Links = &v
}

func (o ResourceSetBindingEditResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingEditResponse) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingEditResponse := _ResourceSetBindingEditResponse{}

	err = json.Unmarshal(bytes, &varResourceSetBindingEditResponse)
	if err == nil {
		*o = ResourceSetBindingEditResponse(varResourceSetBindingEditResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingEditResponse struct {
	value *ResourceSetBindingEditResponse
	isSet bool
}

func (v NullableResourceSetBindingEditResponse) Get() *ResourceSetBindingEditResponse {
	return v.value
}

func (v *NullableResourceSetBindingEditResponse) Set(val *ResourceSetBindingEditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingEditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingEditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingEditResponse(val *ResourceSetBindingEditResponse) *NullableResourceSetBindingEditResponse {
	return &NullableResourceSetBindingEditResponse{value: val, isSet: true}
}

func (v NullableResourceSetBindingEditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingEditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

