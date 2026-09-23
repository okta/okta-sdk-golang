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

// checks if the CertificateAuthorityCertificateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityCertificateList{}

// CertificateAuthorityCertificateList A collection of certificate authority certificates
type CertificateAuthorityCertificateList struct {
	// The certificates, newest first
	Data                 []CertificateAuthorityCertificate `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityCertificateList CertificateAuthorityCertificateList

// NewCertificateAuthorityCertificateList instantiates a new CertificateAuthorityCertificateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityCertificateList() *CertificateAuthorityCertificateList {
	this := CertificateAuthorityCertificateList{}
	return &this
}

// NewCertificateAuthorityCertificateListWithDefaults instantiates a new CertificateAuthorityCertificateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityCertificateListWithDefaults() *CertificateAuthorityCertificateList {
	this := CertificateAuthorityCertificateList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificateList) GetData() []CertificateAuthorityCertificate {
	if o == nil || IsNil(o.Data) {
		var ret []CertificateAuthorityCertificate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificateList) GetDataOk() ([]CertificateAuthorityCertificate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificateList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CertificateAuthorityCertificate and assigns it to the Data field.
func (o *CertificateAuthorityCertificateList) SetData(v []CertificateAuthorityCertificate) {
	o.Data = v
}

func (o CertificateAuthorityCertificateList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityCertificateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityCertificateList) UnmarshalJSON(data []byte) (err error) {
	varCertificateAuthorityCertificateList := _CertificateAuthorityCertificateList{}

	err = json.Unmarshal(data, &varCertificateAuthorityCertificateList)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityCertificateList(varCertificateAuthorityCertificateList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityCertificateList struct {
	value *CertificateAuthorityCertificateList
	isSet bool
}

func (v NullableCertificateAuthorityCertificateList) Get() *CertificateAuthorityCertificateList {
	return v.value
}

func (v *NullableCertificateAuthorityCertificateList) Set(val *CertificateAuthorityCertificateList) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityCertificateList) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityCertificateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityCertificateList(val *CertificateAuthorityCertificateList) *NullableCertificateAuthorityCertificateList {
	return &NullableCertificateAuthorityCertificateList{value: val, isSet: true}
}

func (v NullableCertificateAuthorityCertificateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityCertificateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
