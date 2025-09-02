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
	"fmt"
)

// IdPCertificateCredential struct for IdPCertificateCredential
type IdPCertificateCredential struct {
	// Base64-encoded X.509 certificate chain with DER encoding
	X5c []string `json:"x5c"`
	AdditionalProperties map[string]interface{}
}

type _IdPCertificateCredential IdPCertificateCredential

// NewIdPCertificateCredential instantiates a new IdPCertificateCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdPCertificateCredential(x5c []string) *IdPCertificateCredential {
	this := IdPCertificateCredential{}
	this.X5c = x5c
	return &this
}

// NewIdPCertificateCredentialWithDefaults instantiates a new IdPCertificateCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdPCertificateCredentialWithDefaults() *IdPCertificateCredential {
	this := IdPCertificateCredential{}
	return &this
}

// GetX5c returns the X5c field value
func (o *IdPCertificateCredential) GetX5c() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value
// and a boolean to check if the value has been set.
func (o *IdPCertificateCredential) GetX5cOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.X5c, true
}

// SetX5c sets field value
func (o *IdPCertificateCredential) SetX5c(v []string) {
	o.X5c = v
}

func (o IdPCertificateCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x5c"] = o.X5c
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdPCertificateCredential) UnmarshalJSON(bytes []byte) (err error) {
	varIdPCertificateCredential := _IdPCertificateCredential{}

	err = json.Unmarshal(bytes, &varIdPCertificateCredential)
	if err == nil {
		*o = IdPCertificateCredential(varIdPCertificateCredential)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "x5c")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdPCertificateCredential struct {
	value *IdPCertificateCredential
	isSet bool
}

func (v NullableIdPCertificateCredential) Get() *IdPCertificateCredential {
	return v.value
}

func (v *NullableIdPCertificateCredential) Set(val *IdPCertificateCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableIdPCertificateCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableIdPCertificateCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdPCertificateCredential(val *IdPCertificateCredential) *NullableIdPCertificateCredential {
	return &NullableIdPCertificateCredential{value: val, isSet: true}
}

func (v NullableIdPCertificateCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdPCertificateCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

