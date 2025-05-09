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

// EmailDomainResponseWithEmbeddedEmbedded struct for EmailDomainResponseWithEmbeddedEmbedded
type EmailDomainResponseWithEmbeddedEmbedded struct {
	Brands []Brand `json:"brands,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailDomainResponseWithEmbeddedEmbedded EmailDomainResponseWithEmbeddedEmbedded

// NewEmailDomainResponseWithEmbeddedEmbedded instantiates a new EmailDomainResponseWithEmbeddedEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainResponseWithEmbeddedEmbedded() *EmailDomainResponseWithEmbeddedEmbedded {
	this := EmailDomainResponseWithEmbeddedEmbedded{}
	return &this
}

// NewEmailDomainResponseWithEmbeddedEmbeddedWithDefaults instantiates a new EmailDomainResponseWithEmbeddedEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainResponseWithEmbeddedEmbeddedWithDefaults() *EmailDomainResponseWithEmbeddedEmbedded {
	this := EmailDomainResponseWithEmbeddedEmbedded{}
	return &this
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbeddedEmbedded) GetBrands() []Brand {
	if o == nil || o.Brands == nil {
		var ret []Brand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbeddedEmbedded) GetBrandsOk() ([]Brand, bool) {
	if o == nil || o.Brands == nil {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbeddedEmbedded) HasBrands() bool {
	if o != nil && o.Brands != nil {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []Brand and assigns it to the Brands field.
func (o *EmailDomainResponseWithEmbeddedEmbedded) SetBrands(v []Brand) {
	o.Brands = v
}

func (o EmailDomainResponseWithEmbeddedEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Brands != nil {
		toSerialize["brands"] = o.Brands
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailDomainResponseWithEmbeddedEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDomainResponseWithEmbeddedEmbedded := _EmailDomainResponseWithEmbeddedEmbedded{}

	err = json.Unmarshal(bytes, &varEmailDomainResponseWithEmbeddedEmbedded)
	if err == nil {
		*o = EmailDomainResponseWithEmbeddedEmbedded(varEmailDomainResponseWithEmbeddedEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "brands")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDomainResponseWithEmbeddedEmbedded struct {
	value *EmailDomainResponseWithEmbeddedEmbedded
	isSet bool
}

func (v NullableEmailDomainResponseWithEmbeddedEmbedded) Get() *EmailDomainResponseWithEmbeddedEmbedded {
	return v.value
}

func (v *NullableEmailDomainResponseWithEmbeddedEmbedded) Set(val *EmailDomainResponseWithEmbeddedEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainResponseWithEmbeddedEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainResponseWithEmbeddedEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainResponseWithEmbeddedEmbedded(val *EmailDomainResponseWithEmbeddedEmbedded) *NullableEmailDomainResponseWithEmbeddedEmbedded {
	return &NullableEmailDomainResponseWithEmbeddedEmbedded{value: val, isSet: true}
}

func (v NullableEmailDomainResponseWithEmbeddedEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainResponseWithEmbeddedEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

