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

// checks if the SamlResponseAlgorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlResponseAlgorithm{}

// SamlResponseAlgorithm Algorithm settings for verifying `<SAMLResponse>` messages and `<Assertion>` elements from the IdP
type SamlResponseAlgorithm struct {
	Signature            *SamlResponseSignatureAlgorithm `json:"signature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlResponseAlgorithm SamlResponseAlgorithm

// NewSamlResponseAlgorithm instantiates a new SamlResponseAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlResponseAlgorithm() *SamlResponseAlgorithm {
	this := SamlResponseAlgorithm{}
	return &this
}

// NewSamlResponseAlgorithmWithDefaults instantiates a new SamlResponseAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlResponseAlgorithmWithDefaults() *SamlResponseAlgorithm {
	this := SamlResponseAlgorithm{}
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *SamlResponseAlgorithm) GetSignature() SamlResponseSignatureAlgorithm {
	if o == nil || IsNil(o.Signature) {
		var ret SamlResponseSignatureAlgorithm
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlResponseAlgorithm) GetSignatureOk() (*SamlResponseSignatureAlgorithm, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *SamlResponseAlgorithm) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given SamlResponseSignatureAlgorithm and assigns it to the Signature field.
func (o *SamlResponseAlgorithm) SetSignature(v SamlResponseSignatureAlgorithm) {
	o.Signature = &v
}

func (o SamlResponseAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlResponseAlgorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlResponseAlgorithm) UnmarshalJSON(data []byte) (err error) {
	varSamlResponseAlgorithm := _SamlResponseAlgorithm{}

	err = json.Unmarshal(data, &varSamlResponseAlgorithm)

	if err != nil {
		return err
	}

	*o = SamlResponseAlgorithm(varSamlResponseAlgorithm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlResponseAlgorithm struct {
	value *SamlResponseAlgorithm
	isSet bool
}

func (v NullableSamlResponseAlgorithm) Get() *SamlResponseAlgorithm {
	return v.value
}

func (v *NullableSamlResponseAlgorithm) Set(val *SamlResponseAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlResponseAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlResponseAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlResponseAlgorithm(val *SamlResponseAlgorithm) *NullableSamlResponseAlgorithm {
	return &NullableSamlResponseAlgorithm{value: val, isSet: true}
}

func (v NullableSamlResponseAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlResponseAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
