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

// checks if the DomainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainResponse{}

// DomainResponse The properties that define an individual domain.
type DomainResponse struct {
	// The ID number of the brand
	BrandId *string `json:"brandId,omitempty"`
	// Certificate source type that indicates whether the certificate is provided by the user or Okta.
	CertificateSourceType *string            `json:"certificateSourceType,omitempty"`
	DnsRecords            []DNSRecordDomains `json:"dnsRecords,omitempty"`
	// Custom domain name
	Domain *string `json:"domain,omitempty"`
	// Unique ID of the domain
	Id                *string                    `json:"id,omitempty"`
	PublicCertificate *DomainCertificateMetadata `json:"publicCertificate,omitempty"`
	// Status of the domain
	ValidationStatus     *string      `json:"validationStatus,omitempty"`
	Links                *DomainLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DomainResponse DomainResponse

// NewDomainResponse instantiates a new DomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainResponse() *DomainResponse {
	this := DomainResponse{}
	return &this
}

// NewDomainResponseWithDefaults instantiates a new DomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainResponseWithDefaults() *DomainResponse {
	this := DomainResponse{}
	return &this
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *DomainResponse) GetBrandId() string {
	if o == nil || IsNil(o.BrandId) {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *DomainResponse) HasBrandId() bool {
	if o != nil && !IsNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *DomainResponse) SetBrandId(v string) {
	o.BrandId = &v
}

// GetCertificateSourceType returns the CertificateSourceType field value if set, zero value otherwise.
func (o *DomainResponse) GetCertificateSourceType() string {
	if o == nil || IsNil(o.CertificateSourceType) {
		var ret string
		return ret
	}
	return *o.CertificateSourceType
}

// GetCertificateSourceTypeOk returns a tuple with the CertificateSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetCertificateSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateSourceType) {
		return nil, false
	}
	return o.CertificateSourceType, true
}

// HasCertificateSourceType returns a boolean if a field has been set.
func (o *DomainResponse) HasCertificateSourceType() bool {
	if o != nil && !IsNil(o.CertificateSourceType) {
		return true
	}

	return false
}

// SetCertificateSourceType gets a reference to the given string and assigns it to the CertificateSourceType field.
func (o *DomainResponse) SetCertificateSourceType(v string) {
	o.CertificateSourceType = &v
}

// GetDnsRecords returns the DnsRecords field value if set, zero value otherwise.
func (o *DomainResponse) GetDnsRecords() []DNSRecordDomains {
	if o == nil || IsNil(o.DnsRecords) {
		var ret []DNSRecordDomains
		return ret
	}
	return o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetDnsRecordsOk() ([]DNSRecordDomains, bool) {
	if o == nil || IsNil(o.DnsRecords) {
		return nil, false
	}
	return o.DnsRecords, true
}

// HasDnsRecords returns a boolean if a field has been set.
func (o *DomainResponse) HasDnsRecords() bool {
	if o != nil && !IsNil(o.DnsRecords) {
		return true
	}

	return false
}

// SetDnsRecords gets a reference to the given []DNSRecordDomains and assigns it to the DnsRecords field.
func (o *DomainResponse) SetDnsRecords(v []DNSRecordDomains) {
	o.DnsRecords = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DomainResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DomainResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DomainResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DomainResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DomainResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DomainResponse) SetId(v string) {
	o.Id = &v
}

// GetPublicCertificate returns the PublicCertificate field value if set, zero value otherwise.
func (o *DomainResponse) GetPublicCertificate() DomainCertificateMetadata {
	if o == nil || IsNil(o.PublicCertificate) {
		var ret DomainCertificateMetadata
		return ret
	}
	return *o.PublicCertificate
}

// GetPublicCertificateOk returns a tuple with the PublicCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetPublicCertificateOk() (*DomainCertificateMetadata, bool) {
	if o == nil || IsNil(o.PublicCertificate) {
		return nil, false
	}
	return o.PublicCertificate, true
}

// HasPublicCertificate returns a boolean if a field has been set.
func (o *DomainResponse) HasPublicCertificate() bool {
	if o != nil && !IsNil(o.PublicCertificate) {
		return true
	}

	return false
}

// SetPublicCertificate gets a reference to the given DomainCertificateMetadata and assigns it to the PublicCertificate field.
func (o *DomainResponse) SetPublicCertificate(v DomainCertificateMetadata) {
	o.PublicCertificate = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *DomainResponse) GetValidationStatus() string {
	if o == nil || IsNil(o.ValidationStatus) {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetValidationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationStatus) {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *DomainResponse) HasValidationStatus() bool {
	if o != nil && !IsNil(o.ValidationStatus) {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *DomainResponse) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DomainResponse) GetLinks() DomainLinks {
	if o == nil || IsNil(o.Links) {
		var ret DomainLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetLinksOk() (*DomainLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DomainResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DomainLinks and assigns it to the Links field.
func (o *DomainResponse) SetLinks(v DomainLinks) {
	o.Links = &v
}

func (o DomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandId) {
		toSerialize["brandId"] = o.BrandId
	}
	if !IsNil(o.CertificateSourceType) {
		toSerialize["certificateSourceType"] = o.CertificateSourceType
	}
	if !IsNil(o.DnsRecords) {
		toSerialize["dnsRecords"] = o.DnsRecords
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PublicCertificate) {
		toSerialize["publicCertificate"] = o.PublicCertificate
	}
	if !IsNil(o.ValidationStatus) {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DomainResponse) UnmarshalJSON(data []byte) (err error) {
	varDomainResponse := _DomainResponse{}

	err = json.Unmarshal(data, &varDomainResponse)

	if err != nil {
		return err
	}

	*o = DomainResponse(varDomainResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "brandId")
		delete(additionalProperties, "certificateSourceType")
		delete(additionalProperties, "dnsRecords")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "id")
		delete(additionalProperties, "publicCertificate")
		delete(additionalProperties, "validationStatus")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDomainResponse struct {
	value *DomainResponse
	isSet bool
}

func (v NullableDomainResponse) Get() *DomainResponse {
	return v.value
}

func (v *NullableDomainResponse) Set(val *DomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainResponse(val *DomainResponse) *NullableDomainResponse {
	return &NullableDomainResponse{value: val, isSet: true}
}

func (v NullableDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
