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

// DomainRequest struct for DomainRequest
type DomainRequest struct {
	// Certificate source type that indicates whether the certificate is provided by the user or Okta.
	CertificateSourceType string `json:"certificateSourceType"`
	// Custom domain name
	Domain string `json:"domain"`
	AdditionalProperties map[string]interface{}
}

type _DomainRequest DomainRequest

// NewDomainRequest instantiates a new DomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRequest(certificateSourceType string, domain string) *DomainRequest {
	this := DomainRequest{}
	this.CertificateSourceType = certificateSourceType
	this.Domain = domain
	return &this
}

// NewDomainRequestWithDefaults instantiates a new DomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRequestWithDefaults() *DomainRequest {
	this := DomainRequest{}
	return &this
}

// GetCertificateSourceType returns the CertificateSourceType field value
func (o *DomainRequest) GetCertificateSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateSourceType
}

// GetCertificateSourceTypeOk returns a tuple with the CertificateSourceType field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetCertificateSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateSourceType, true
}

// SetCertificateSourceType sets field value
func (o *DomainRequest) SetCertificateSourceType(v string) {
	o.CertificateSourceType = v
}

// GetDomain returns the Domain field value
func (o *DomainRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainRequest) SetDomain(v string) {
	o.Domain = v
}

func (o DomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certificateSourceType"] = o.CertificateSourceType
	}
	if true {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDomainRequest := _DomainRequest{}

	err = json.Unmarshal(bytes, &varDomainRequest)
	if err == nil {
		*o = DomainRequest(varDomainRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "certificateSourceType")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomainRequest struct {
	value *DomainRequest
	isSet bool
}

func (v NullableDomainRequest) Get() *DomainRequest {
	return v.value
}

func (v *NullableDomainRequest) Set(val *DomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRequest(val *DomainRequest) *NullableDomainRequest {
	return &NullableDomainRequest{value: val, isSet: true}
}

func (v NullableDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

