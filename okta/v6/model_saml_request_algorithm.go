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

// SamlRequestAlgorithm Algorithm settings used to secure an `<AuthnRequest>` message
type SamlRequestAlgorithm struct {
	Signature *SamlRequestSignatureAlgorithm `json:"signature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlRequestAlgorithm SamlRequestAlgorithm

// NewSamlRequestAlgorithm instantiates a new SamlRequestAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlRequestAlgorithm() *SamlRequestAlgorithm {
	this := SamlRequestAlgorithm{}
	return &this
}

// NewSamlRequestAlgorithmWithDefaults instantiates a new SamlRequestAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlRequestAlgorithmWithDefaults() *SamlRequestAlgorithm {
	this := SamlRequestAlgorithm{}
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *SamlRequestAlgorithm) GetSignature() SamlRequestSignatureAlgorithm {
	if o == nil || o.Signature == nil {
		var ret SamlRequestSignatureAlgorithm
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlRequestAlgorithm) GetSignatureOk() (*SamlRequestSignatureAlgorithm, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *SamlRequestAlgorithm) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given SamlRequestSignatureAlgorithm and assigns it to the Signature field.
func (o *SamlRequestAlgorithm) SetSignature(v SamlRequestSignatureAlgorithm) {
	o.Signature = &v
}

func (o SamlRequestAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlRequestAlgorithm) UnmarshalJSON(bytes []byte) (err error) {
	varSamlRequestAlgorithm := _SamlRequestAlgorithm{}

	err = json.Unmarshal(bytes, &varSamlRequestAlgorithm)
	if err == nil {
		*o = SamlRequestAlgorithm(varSamlRequestAlgorithm)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signature")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlRequestAlgorithm struct {
	value *SamlRequestAlgorithm
	isSet bool
}

func (v NullableSamlRequestAlgorithm) Get() *SamlRequestAlgorithm {
	return v.value
}

func (v *NullableSamlRequestAlgorithm) Set(val *SamlRequestAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlRequestAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlRequestAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlRequestAlgorithm(val *SamlRequestAlgorithm) *NullableSamlRequestAlgorithm {
	return &NullableSamlRequestAlgorithm{value: val, isSet: true}
}

func (v NullableSamlRequestAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlRequestAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

