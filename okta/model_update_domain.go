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

// UpdateDomain struct for UpdateDomain
type UpdateDomain struct {
	// The `id` of the brand used to replace the existing brand.
	BrandId string `json:"brandId"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDomain UpdateDomain

// NewUpdateDomain instantiates a new UpdateDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDomain(brandId string) *UpdateDomain {
	this := UpdateDomain{}
	this.BrandId = brandId
	return &this
}

// NewUpdateDomainWithDefaults instantiates a new UpdateDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDomainWithDefaults() *UpdateDomain {
	this := UpdateDomain{}
	return &this
}

// GetBrandId returns the BrandId field value
func (o *UpdateDomain) GetBrandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value
// and a boolean to check if the value has been set.
func (o *UpdateDomain) GetBrandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandId, true
}

// SetBrandId sets field value
func (o *UpdateDomain) SetBrandId(v string) {
	o.BrandId = v
}

func (o UpdateDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brandId"] = o.BrandId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateDomain) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateDomain := _UpdateDomain{}

	err = json.Unmarshal(bytes, &varUpdateDomain)
	if err == nil {
		*o = UpdateDomain(varUpdateDomain)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "brandId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUpdateDomain struct {
	value *UpdateDomain
	isSet bool
}

func (v NullableUpdateDomain) Get() *UpdateDomain {
	return v.value
}

func (v *NullableUpdateDomain) Set(val *UpdateDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDomain(val *UpdateDomain) *NullableUpdateDomain {
	return &NullableUpdateDomain{value: val, isSet: true}
}

func (v NullableUpdateDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

