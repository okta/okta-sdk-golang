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
	"fmt"
)

// checks if the DomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainRequest{}

// DomainRequest struct for DomainRequest
type DomainRequest struct {
	// Certificate source type that indicates whether the certificate is provided by the user or Okta.
	CertificateSourceType string `json:"certificateSourceType"`
	// Custom domain name  > **Note:** You can't use the reserved `drapp.{yourOrgSubDomain}.okta.com` domain.
	Domain               string `json:"domain"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificateSourceType"] = o.CertificateSourceType
	toSerialize["domain"] = o.Domain

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DomainRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificateSourceType",
		"domain",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDomainRequest := _DomainRequest{}

	err = json.Unmarshal(data, &varDomainRequest)

	if err != nil {
		return err
	}

	*o = DomainRequest(varDomainRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificateSourceType")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
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
