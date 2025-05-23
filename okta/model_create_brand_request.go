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

// CreateBrandRequest struct for CreateBrandRequest
type CreateBrandRequest struct {
	// The name of the Brand
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreateBrandRequest CreateBrandRequest

// NewCreateBrandRequest instantiates a new CreateBrandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBrandRequest(name string) *CreateBrandRequest {
	this := CreateBrandRequest{}
	this.Name = name
	return &this
}

// NewCreateBrandRequestWithDefaults instantiates a new CreateBrandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBrandRequestWithDefaults() *CreateBrandRequest {
	this := CreateBrandRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateBrandRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBrandRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBrandRequest) SetName(v string) {
	o.Name = v
}

func (o CreateBrandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateBrandRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateBrandRequest := _CreateBrandRequest{}

	err = json.Unmarshal(bytes, &varCreateBrandRequest)
	if err == nil {
		*o = CreateBrandRequest(varCreateBrandRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateBrandRequest struct {
	value *CreateBrandRequest
	isSet bool
}

func (v NullableCreateBrandRequest) Get() *CreateBrandRequest {
	return v.value
}

func (v *NullableCreateBrandRequest) Set(val *CreateBrandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBrandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBrandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBrandRequest(val *CreateBrandRequest) *NullableCreateBrandRequest {
	return &NullableCreateBrandRequest{value: val, isSet: true}
}

func (v NullableCreateBrandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBrandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

