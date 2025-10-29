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

// checks if the SamlRequestSignatureAlgorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlRequestSignatureAlgorithm{}

// SamlRequestSignatureAlgorithm XML digital Signature Algorithm settings for signing `<AuthnRequest>` messages sent to the IdP > **Note:**  The `algorithm` property is ignored when you disable request signatures (`scope` set as `NONE`).
type SamlRequestSignatureAlgorithm struct {
	Algorithm *string `json:"algorithm,omitempty"`
	// Specifies whether to digitally sign authorization requests to the IdP
	Scope                *string `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlRequestSignatureAlgorithm SamlRequestSignatureAlgorithm

// NewSamlRequestSignatureAlgorithm instantiates a new SamlRequestSignatureAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlRequestSignatureAlgorithm() *SamlRequestSignatureAlgorithm {
	this := SamlRequestSignatureAlgorithm{}
	return &this
}

// NewSamlRequestSignatureAlgorithmWithDefaults instantiates a new SamlRequestSignatureAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlRequestSignatureAlgorithmWithDefaults() *SamlRequestSignatureAlgorithm {
	this := SamlRequestSignatureAlgorithm{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *SamlRequestSignatureAlgorithm) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlRequestSignatureAlgorithm) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *SamlRequestSignatureAlgorithm) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *SamlRequestSignatureAlgorithm) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SamlRequestSignatureAlgorithm) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlRequestSignatureAlgorithm) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SamlRequestSignatureAlgorithm) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SamlRequestSignatureAlgorithm) SetScope(v string) {
	o.Scope = &v
}

func (o SamlRequestSignatureAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlRequestSignatureAlgorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlRequestSignatureAlgorithm) UnmarshalJSON(data []byte) (err error) {
	varSamlRequestSignatureAlgorithm := _SamlRequestSignatureAlgorithm{}

	err = json.Unmarshal(data, &varSamlRequestSignatureAlgorithm)

	if err != nil {
		return err
	}

	*o = SamlRequestSignatureAlgorithm(varSamlRequestSignatureAlgorithm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlRequestSignatureAlgorithm struct {
	value *SamlRequestSignatureAlgorithm
	isSet bool
}

func (v NullableSamlRequestSignatureAlgorithm) Get() *SamlRequestSignatureAlgorithm {
	return v.value
}

func (v *NullableSamlRequestSignatureAlgorithm) Set(val *SamlRequestSignatureAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlRequestSignatureAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlRequestSignatureAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlRequestSignatureAlgorithm(val *SamlRequestSignatureAlgorithm) *NullableSamlRequestSignatureAlgorithm {
	return &NullableSamlRequestSignatureAlgorithm{value: val, isSet: true}
}

func (v NullableSamlRequestSignatureAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlRequestSignatureAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
