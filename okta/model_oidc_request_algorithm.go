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

// checks if the OidcRequestAlgorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcRequestAlgorithm{}

// OidcRequestAlgorithm Algorithm settings used to sign an authorization request
type OidcRequestAlgorithm struct {
	Signature            *OidcRequestSignatureAlgorithm `json:"signature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcRequestAlgorithm OidcRequestAlgorithm

// NewOidcRequestAlgorithm instantiates a new OidcRequestAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcRequestAlgorithm() *OidcRequestAlgorithm {
	this := OidcRequestAlgorithm{}
	return &this
}

// NewOidcRequestAlgorithmWithDefaults instantiates a new OidcRequestAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcRequestAlgorithmWithDefaults() *OidcRequestAlgorithm {
	this := OidcRequestAlgorithm{}
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *OidcRequestAlgorithm) GetSignature() OidcRequestSignatureAlgorithm {
	if o == nil || IsNil(o.Signature) {
		var ret OidcRequestSignatureAlgorithm
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcRequestAlgorithm) GetSignatureOk() (*OidcRequestSignatureAlgorithm, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *OidcRequestAlgorithm) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given OidcRequestSignatureAlgorithm and assigns it to the Signature field.
func (o *OidcRequestAlgorithm) SetSignature(v OidcRequestSignatureAlgorithm) {
	o.Signature = &v
}

func (o OidcRequestAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcRequestAlgorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OidcRequestAlgorithm) UnmarshalJSON(data []byte) (err error) {
	varOidcRequestAlgorithm := _OidcRequestAlgorithm{}

	err = json.Unmarshal(data, &varOidcRequestAlgorithm)

	if err != nil {
		return err
	}

	*o = OidcRequestAlgorithm(varOidcRequestAlgorithm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidcRequestAlgorithm struct {
	value *OidcRequestAlgorithm
	isSet bool
}

func (v NullableOidcRequestAlgorithm) Get() *OidcRequestAlgorithm {
	return v.value
}

func (v *NullableOidcRequestAlgorithm) Set(val *OidcRequestAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcRequestAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcRequestAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcRequestAlgorithm(val *OidcRequestAlgorithm) *NullableOidcRequestAlgorithm {
	return &NullableOidcRequestAlgorithm{value: val, isSet: true}
}

func (v NullableOidcRequestAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcRequestAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
