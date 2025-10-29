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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ImageUploadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageUploadResponse{}

// ImageUploadResponse struct for ImageUploadResponse
type ImageUploadResponse struct {
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImageUploadResponse ImageUploadResponse

// NewImageUploadResponse instantiates a new ImageUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageUploadResponse() *ImageUploadResponse {
	this := ImageUploadResponse{}
	return &this
}

// NewImageUploadResponseWithDefaults instantiates a new ImageUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageUploadResponseWithDefaults() *ImageUploadResponse {
	this := ImageUploadResponse{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ImageUploadResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageUploadResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ImageUploadResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ImageUploadResponse) SetUrl(v string) {
	o.Url = &v
}

func (o ImageUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageUploadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageUploadResponse) UnmarshalJSON(data []byte) (err error) {
	varImageUploadResponse := _ImageUploadResponse{}

	err = json.Unmarshal(data, &varImageUploadResponse)

	if err != nil {
		return err
	}

	*o = ImageUploadResponse(varImageUploadResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageUploadResponse struct {
	value *ImageUploadResponse
	isSet bool
}

func (v NullableImageUploadResponse) Get() *ImageUploadResponse {
	return v.value
}

func (v *NullableImageUploadResponse) Set(val *ImageUploadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageUploadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageUploadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageUploadResponse(val *ImageUploadResponse) *NullableImageUploadResponse {
	return &NullableImageUploadResponse{value: val, isSet: true}
}

func (v NullableImageUploadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageUploadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
