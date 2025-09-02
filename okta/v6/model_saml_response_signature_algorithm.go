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

// SamlResponseSignatureAlgorithm XML digital Signature Algorithm settings for verifying `<SAMLResponse>` messages and `<Assertion>` elements from the IdP
type SamlResponseSignatureAlgorithm struct {
	Algorithm *string `json:"algorithm,omitempty"`
	// Specifies whether to verify responses from the IdP
	Scope *string `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlResponseSignatureAlgorithm SamlResponseSignatureAlgorithm

// NewSamlResponseSignatureAlgorithm instantiates a new SamlResponseSignatureAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlResponseSignatureAlgorithm() *SamlResponseSignatureAlgorithm {
	this := SamlResponseSignatureAlgorithm{}
	return &this
}

// NewSamlResponseSignatureAlgorithmWithDefaults instantiates a new SamlResponseSignatureAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlResponseSignatureAlgorithmWithDefaults() *SamlResponseSignatureAlgorithm {
	this := SamlResponseSignatureAlgorithm{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *SamlResponseSignatureAlgorithm) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlResponseSignatureAlgorithm) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *SamlResponseSignatureAlgorithm) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *SamlResponseSignatureAlgorithm) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SamlResponseSignatureAlgorithm) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlResponseSignatureAlgorithm) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SamlResponseSignatureAlgorithm) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SamlResponseSignatureAlgorithm) SetScope(v string) {
	o.Scope = &v
}

func (o SamlResponseSignatureAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlResponseSignatureAlgorithm) UnmarshalJSON(bytes []byte) (err error) {
	varSamlResponseSignatureAlgorithm := _SamlResponseSignatureAlgorithm{}

	err = json.Unmarshal(bytes, &varSamlResponseSignatureAlgorithm)
	if err == nil {
		*o = SamlResponseSignatureAlgorithm(varSamlResponseSignatureAlgorithm)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlResponseSignatureAlgorithm struct {
	value *SamlResponseSignatureAlgorithm
	isSet bool
}

func (v NullableSamlResponseSignatureAlgorithm) Get() *SamlResponseSignatureAlgorithm {
	return v.value
}

func (v *NullableSamlResponseSignatureAlgorithm) Set(val *SamlResponseSignatureAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlResponseSignatureAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlResponseSignatureAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlResponseSignatureAlgorithm(val *SamlResponseSignatureAlgorithm) *NullableSamlResponseSignatureAlgorithm {
	return &NullableSamlResponseSignatureAlgorithm{value: val, isSet: true}
}

func (v NullableSamlResponseSignatureAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlResponseSignatureAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

