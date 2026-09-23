/*
Okta Admin Management API

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

// checks if the CertificateAuthorityJsonWebKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityJsonWebKey{}

// CertificateAuthorityJsonWebKey The public key of the certificate, as a JSON Web Key. Only public RSA parameters are returned.
type CertificateAuthorityJsonWebKey struct {
	// The RSA public exponent, base64url-encoded
	E *string `json:"e,omitempty"`
	// The key type. Always `RSA`, because Okta generates RSA key pairs for these authorities.
	Kty *string `json:"kty,omitempty"`
	// The RSA modulus, base64url-encoded
	N *string `json:"n,omitempty"`
	// What the key is used for. Always `sig`, because a CA key signs certificates.
	Use                  *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityJsonWebKey CertificateAuthorityJsonWebKey

// NewCertificateAuthorityJsonWebKey instantiates a new CertificateAuthorityJsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityJsonWebKey() *CertificateAuthorityJsonWebKey {
	this := CertificateAuthorityJsonWebKey{}
	return &this
}

// NewCertificateAuthorityJsonWebKeyWithDefaults instantiates a new CertificateAuthorityJsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityJsonWebKeyWithDefaults() *CertificateAuthorityJsonWebKey {
	this := CertificateAuthorityJsonWebKey{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *CertificateAuthorityJsonWebKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityJsonWebKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *CertificateAuthorityJsonWebKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *CertificateAuthorityJsonWebKey) SetE(v string) {
	o.E = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *CertificateAuthorityJsonWebKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityJsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *CertificateAuthorityJsonWebKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *CertificateAuthorityJsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *CertificateAuthorityJsonWebKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityJsonWebKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *CertificateAuthorityJsonWebKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *CertificateAuthorityJsonWebKey) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *CertificateAuthorityJsonWebKey) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityJsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *CertificateAuthorityJsonWebKey) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *CertificateAuthorityJsonWebKey) SetUse(v string) {
	o.Use = &v
}

func (o CertificateAuthorityJsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityJsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityJsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varCertificateAuthorityJsonWebKey := _CertificateAuthorityJsonWebKey{}

	err = json.Unmarshal(data, &varCertificateAuthorityJsonWebKey)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityJsonWebKey(varCertificateAuthorityJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityJsonWebKey struct {
	value *CertificateAuthorityJsonWebKey
	isSet bool
}

func (v NullableCertificateAuthorityJsonWebKey) Get() *CertificateAuthorityJsonWebKey {
	return v.value
}

func (v *NullableCertificateAuthorityJsonWebKey) Set(val *CertificateAuthorityJsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityJsonWebKey(val *CertificateAuthorityJsonWebKey) *NullableCertificateAuthorityJsonWebKey {
	return &NullableCertificateAuthorityJsonWebKey{value: val, isSet: true}
}

func (v NullableCertificateAuthorityJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
