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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// DomainCertificate Defines the properties of the certificate
type DomainCertificate struct {
	// Certificate content
	Certificate string `json:"certificate"`
	// Certificate chain
	CertificateChain string `json:"certificateChain"`
	// Certificate private key
	PrivateKey string `json:"privateKey"`
	// Certificate type
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _DomainCertificate DomainCertificate

// NewDomainCertificate instantiates a new DomainCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainCertificate(certificate string, certificateChain string, privateKey string, type_ string) *DomainCertificate {
	this := DomainCertificate{}
	this.Certificate = certificate
	this.CertificateChain = certificateChain
	this.PrivateKey = privateKey
	this.Type = type_
	return &this
}

// NewDomainCertificateWithDefaults instantiates a new DomainCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainCertificateWithDefaults() *DomainCertificate {
	this := DomainCertificate{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *DomainCertificate) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *DomainCertificate) SetCertificate(v string) {
	o.Certificate = v
}

// GetCertificateChain returns the CertificateChain field value
func (o *DomainCertificate) GetCertificateChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateChain
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetCertificateChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateChain, true
}

// SetCertificateChain sets field value
func (o *DomainCertificate) SetCertificateChain(v string) {
	o.CertificateChain = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *DomainCertificate) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *DomainCertificate) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetType returns the Type field value
func (o *DomainCertificate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DomainCertificate) SetType(v string) {
	o.Type = v
}

func (o DomainCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certificate"] = o.Certificate
	}
	if true {
		toSerialize["certificateChain"] = o.CertificateChain
	}
	if true {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varDomainCertificate := _DomainCertificate{}

	err = json.Unmarshal(bytes, &varDomainCertificate)
	if err == nil {
		*o = DomainCertificate(varDomainCertificate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "certificateChain")
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomainCertificate struct {
	value *DomainCertificate
	isSet bool
}

func (v NullableDomainCertificate) Get() *DomainCertificate {
	return v.value
}

func (v *NullableDomainCertificate) Set(val *DomainCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainCertificate(val *DomainCertificate) *NullableDomainCertificate {
	return &NullableDomainCertificate{value: val, isSet: true}
}

func (v NullableDomainCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

