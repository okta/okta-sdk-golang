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

// checks if the CertificateAuthorityEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityEmbedded{}

// CertificateAuthorityEmbedded Resources embedded in the authority. Present only when the request sets `expand=certificates`.
type CertificateAuthorityEmbedded struct {
	// Every valid certificate the authority currently holds
	Certificates         []CertificateAuthorityCertificate `json:"certificates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityEmbedded CertificateAuthorityEmbedded

// NewCertificateAuthorityEmbedded instantiates a new CertificateAuthorityEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityEmbedded() *CertificateAuthorityEmbedded {
	this := CertificateAuthorityEmbedded{}
	return &this
}

// NewCertificateAuthorityEmbeddedWithDefaults instantiates a new CertificateAuthorityEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityEmbeddedWithDefaults() *CertificateAuthorityEmbedded {
	this := CertificateAuthorityEmbedded{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *CertificateAuthorityEmbedded) GetCertificates() []CertificateAuthorityCertificate {
	if o == nil || IsNil(o.Certificates) {
		var ret []CertificateAuthorityCertificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityEmbedded) GetCertificatesOk() ([]CertificateAuthorityCertificate, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *CertificateAuthorityEmbedded) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificateAuthorityCertificate and assigns it to the Certificates field.
func (o *CertificateAuthorityEmbedded) SetCertificates(v []CertificateAuthorityCertificate) {
	o.Certificates = v
}

func (o CertificateAuthorityEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityEmbedded) UnmarshalJSON(data []byte) (err error) {
	varCertificateAuthorityEmbedded := _CertificateAuthorityEmbedded{}

	err = json.Unmarshal(data, &varCertificateAuthorityEmbedded)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityEmbedded(varCertificateAuthorityEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityEmbedded struct {
	value *CertificateAuthorityEmbedded
	isSet bool
}

func (v NullableCertificateAuthorityEmbedded) Get() *CertificateAuthorityEmbedded {
	return v.value
}

func (v *NullableCertificateAuthorityEmbedded) Set(val *CertificateAuthorityEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityEmbedded(val *CertificateAuthorityEmbedded) *NullableCertificateAuthorityEmbedded {
	return &NullableCertificateAuthorityEmbedded{value: val, isSet: true}
}

func (v NullableCertificateAuthorityEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
