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
	"fmt"
	"time"
)

// checks if the CertificateAuthorityToScopeMappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityToScopeMappingResponse{}

// CertificateAuthorityToScopeMappingResponse A Certificate Authority eligible for the requested scope, annotated with its current mapping (if any).
type CertificateAuthorityToScopeMappingResponse struct {
	// The type of the Certificate Authority.
	CaType *string `json:"caType,omitempty"`
	// Timestamp when the object was created
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Timestamp when the object expires
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	// The `id` of the Certificate Authority.
	Id string `json:"id"`
	// The issuer distinguished name parsed from the Certificate Authority's trust anchor.
	IssuerDN *string `json:"issuerDN,omitempty"`
	// Timestamp when the object was last updated
	LastUpdatedDate *time.Time                                         `json:"lastUpdatedDate,omitempty"`
	Mapping         *CertificateAuthorityToScopeMappingResponseMapping `json:"mapping,omitempty"`
	// The validity status of the Certificate Authority's trust anchor.
	Status *string `json:"status,omitempty"`
	// The subject distinguished name parsed from the Certificate Authority's trust anchor.
	SubjectDN            *string `json:"subjectDN,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityToScopeMappingResponse CertificateAuthorityToScopeMappingResponse

// NewCertificateAuthorityToScopeMappingResponse instantiates a new CertificateAuthorityToScopeMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityToScopeMappingResponse(id string) *CertificateAuthorityToScopeMappingResponse {
	this := CertificateAuthorityToScopeMappingResponse{}
	this.Id = id
	return &this
}

// NewCertificateAuthorityToScopeMappingResponseWithDefaults instantiates a new CertificateAuthorityToScopeMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityToScopeMappingResponseWithDefaults() *CertificateAuthorityToScopeMappingResponse {
	this := CertificateAuthorityToScopeMappingResponse{}
	return &this
}

// GetCaType returns the CaType field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetCaType() string {
	if o == nil || IsNil(o.CaType) {
		var ret string
		return ret
	}
	return *o.CaType
}

// GetCaTypeOk returns a tuple with the CaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetCaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CaType) {
		return nil, false
	}
	return o.CaType, true
}

// HasCaType returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasCaType() bool {
	if o != nil && !IsNil(o.CaType) {
		return true
	}

	return false
}

// SetCaType gets a reference to the given string and assigns it to the CaType field.
func (o *CertificateAuthorityToScopeMappingResponse) SetCaType(v string) {
	o.CaType = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *CertificateAuthorityToScopeMappingResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *CertificateAuthorityToScopeMappingResponse) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetId returns the Id field value
func (o *CertificateAuthorityToScopeMappingResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertificateAuthorityToScopeMappingResponse) SetId(v string) {
	o.Id = v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetIssuerDN() string {
	if o == nil || IsNil(o.IssuerDN) {
		var ret string
		return ret
	}
	return *o.IssuerDN
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetIssuerDNOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerDN) {
		return nil, false
	}
	return o.IssuerDN, true
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasIssuerDN() bool {
	if o != nil && !IsNil(o.IssuerDN) {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given string and assigns it to the IssuerDN field.
func (o *CertificateAuthorityToScopeMappingResponse) SetIssuerDN(v string) {
	o.IssuerDN = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *CertificateAuthorityToScopeMappingResponse) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetMapping() CertificateAuthorityToScopeMappingResponseMapping {
	if o == nil || IsNil(o.Mapping) {
		var ret CertificateAuthorityToScopeMappingResponseMapping
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetMappingOk() (*CertificateAuthorityToScopeMappingResponseMapping, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given CertificateAuthorityToScopeMappingResponseMapping and assigns it to the Mapping field.
func (o *CertificateAuthorityToScopeMappingResponse) SetMapping(v CertificateAuthorityToScopeMappingResponseMapping) {
	o.Mapping = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertificateAuthorityToScopeMappingResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *CertificateAuthorityToScopeMappingResponse) GetSubjectDN() string {
	if o == nil || IsNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponse) GetSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *CertificateAuthorityToScopeMappingResponse) HasSubjectDN() bool {
	if o != nil && !IsNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *CertificateAuthorityToScopeMappingResponse) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

func (o CertificateAuthorityToScopeMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityToScopeMappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaType) {
		toSerialize["caType"] = o.CaType
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IssuerDN) {
		toSerialize["issuerDN"] = o.IssuerDN
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubjectDN) {
		toSerialize["subjectDN"] = o.SubjectDN
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityToScopeMappingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varCertificateAuthorityToScopeMappingResponse := _CertificateAuthorityToScopeMappingResponse{}

	err = json.Unmarshal(data, &varCertificateAuthorityToScopeMappingResponse)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityToScopeMappingResponse(varCertificateAuthorityToScopeMappingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "caType")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issuerDN")
		delete(additionalProperties, "lastUpdatedDate")
		delete(additionalProperties, "mapping")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subjectDN")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityToScopeMappingResponse struct {
	value *CertificateAuthorityToScopeMappingResponse
	isSet bool
}

func (v NullableCertificateAuthorityToScopeMappingResponse) Get() *CertificateAuthorityToScopeMappingResponse {
	return v.value
}

func (v *NullableCertificateAuthorityToScopeMappingResponse) Set(val *CertificateAuthorityToScopeMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityToScopeMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityToScopeMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityToScopeMappingResponse(val *CertificateAuthorityToScopeMappingResponse) *NullableCertificateAuthorityToScopeMappingResponse {
	return &NullableCertificateAuthorityToScopeMappingResponse{value: val, isSet: true}
}

func (v NullableCertificateAuthorityToScopeMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityToScopeMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
