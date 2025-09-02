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

// OidcRequestSignatureAlgorithm Signature Algorithm settings for signing authorization requests sent to the IdP > **Note:**  The `algorithm` property is ignored when you disable request signatures (`scope` set as `NONE`).
type OidcRequestSignatureAlgorithm struct {
	Algorithm *string `json:"algorithm,omitempty"`
	// Specifies whether to digitally sign authorization requests to the IdP
	Scope *string `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcRequestSignatureAlgorithm OidcRequestSignatureAlgorithm

// NewOidcRequestSignatureAlgorithm instantiates a new OidcRequestSignatureAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcRequestSignatureAlgorithm() *OidcRequestSignatureAlgorithm {
	this := OidcRequestSignatureAlgorithm{}
	return &this
}

// NewOidcRequestSignatureAlgorithmWithDefaults instantiates a new OidcRequestSignatureAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcRequestSignatureAlgorithmWithDefaults() *OidcRequestSignatureAlgorithm {
	this := OidcRequestSignatureAlgorithm{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *OidcRequestSignatureAlgorithm) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcRequestSignatureAlgorithm) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *OidcRequestSignatureAlgorithm) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *OidcRequestSignatureAlgorithm) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OidcRequestSignatureAlgorithm) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcRequestSignatureAlgorithm) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OidcRequestSignatureAlgorithm) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OidcRequestSignatureAlgorithm) SetScope(v string) {
	o.Scope = &v
}

func (o OidcRequestSignatureAlgorithm) MarshalJSON() ([]byte, error) {
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

func (o *OidcRequestSignatureAlgorithm) UnmarshalJSON(bytes []byte) (err error) {
	varOidcRequestSignatureAlgorithm := _OidcRequestSignatureAlgorithm{}

	err = json.Unmarshal(bytes, &varOidcRequestSignatureAlgorithm)
	if err == nil {
		*o = OidcRequestSignatureAlgorithm(varOidcRequestSignatureAlgorithm)
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

type NullableOidcRequestSignatureAlgorithm struct {
	value *OidcRequestSignatureAlgorithm
	isSet bool
}

func (v NullableOidcRequestSignatureAlgorithm) Get() *OidcRequestSignatureAlgorithm {
	return v.value
}

func (v *NullableOidcRequestSignatureAlgorithm) Set(val *OidcRequestSignatureAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcRequestSignatureAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcRequestSignatureAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcRequestSignatureAlgorithm(val *OidcRequestSignatureAlgorithm) *NullableOidcRequestSignatureAlgorithm {
	return &NullableOidcRequestSignatureAlgorithm{value: val, isSet: true}
}

func (v NullableOidcRequestSignatureAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcRequestSignatureAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

