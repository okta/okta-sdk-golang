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

// checks if the CertificateAuthority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthority{}

// CertificateAuthority A certificate authority hosted by Okta for your org, identified by its `authorityInstanceId` and stable across renewals
type CertificateAuthority struct {
	// The stable identifier of the authority. For the authority's original certificate, this is also that certificate's `ski`, because a CA's own subject key identifier becomes the authority's identifier when the CA is created.
	AuthorityInstanceId *string `json:"authorityInstanceId,omitempty"`
	// The family of the certificate authority. `OKTA_AS_CA` is the Okta CA used for SCEP enrollment and device management attestation; `DEVICE_ACCESS_OKTA_AS_CA` is the Device Access CA. The `caType` query parameter filters on the same values.
	CaType *string `json:"caType,omitempty"`
	// The status of the certificate. `VALID` is in service and able to issue certificates, which is how Okta generates a replacement CA certificate. `INACTIVE` is a replacement that Okta generated but hasn't brought into service, which occurs only for an authority renewed before selective renewal was available; renewing such an authority activates it.
	Status *string `json:"status,omitempty"`
	// The distinguished name in the subject of the authority's certificate
	SubjectDN *string `json:"subjectDN,omitempty"`
	// The position of the certificate authority in the hierarchy. A root is self-signed; an intermediate is signed by the root and signs the certificates Okta issues to devices. The `type` query parameter filters on the same values.
	Type                 *string                       `json:"type,omitempty"`
	Embedded             *CertificateAuthorityEmbedded `json:"_embedded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthority CertificateAuthority

// NewCertificateAuthority instantiates a new CertificateAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthority() *CertificateAuthority {
	this := CertificateAuthority{}
	return &this
}

// NewCertificateAuthorityWithDefaults instantiates a new CertificateAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityWithDefaults() *CertificateAuthority {
	this := CertificateAuthority{}
	return &this
}

// GetAuthorityInstanceId returns the AuthorityInstanceId field value if set, zero value otherwise.
func (o *CertificateAuthority) GetAuthorityInstanceId() string {
	if o == nil || IsNil(o.AuthorityInstanceId) {
		var ret string
		return ret
	}
	return *o.AuthorityInstanceId
}

// GetAuthorityInstanceIdOk returns a tuple with the AuthorityInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthority) GetAuthorityInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityInstanceId) {
		return nil, false
	}
	return o.AuthorityInstanceId, true
}

// HasAuthorityInstanceId returns a boolean if a field has been set.
func (o *CertificateAuthority) HasAuthorityInstanceId() bool {
	if o != nil && !IsNil(o.AuthorityInstanceId) {
		return true
	}

	return false
}

// SetAuthorityInstanceId gets a reference to the given string and assigns it to the AuthorityInstanceId field.
func (o *CertificateAuthority) SetAuthorityInstanceId(v string) {
	o.AuthorityInstanceId = &v
}

// GetCaType returns the CaType field value if set, zero value otherwise.
func (o *CertificateAuthority) GetCaType() string {
	if o == nil || IsNil(o.CaType) {
		var ret string
		return ret
	}
	return *o.CaType
}

// GetCaTypeOk returns a tuple with the CaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthority) GetCaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CaType) {
		return nil, false
	}
	return o.CaType, true
}

// HasCaType returns a boolean if a field has been set.
func (o *CertificateAuthority) HasCaType() bool {
	if o != nil && !IsNil(o.CaType) {
		return true
	}

	return false
}

// SetCaType gets a reference to the given string and assigns it to the CaType field.
func (o *CertificateAuthority) SetCaType(v string) {
	o.CaType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateAuthority) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthority) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateAuthority) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertificateAuthority) SetStatus(v string) {
	o.Status = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *CertificateAuthority) GetSubjectDN() string {
	if o == nil || IsNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthority) GetSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *CertificateAuthority) HasSubjectDN() bool {
	if o != nil && !IsNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *CertificateAuthority) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CertificateAuthority) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthority) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CertificateAuthority) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CertificateAuthority) SetType(v string) {
	o.Type = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *CertificateAuthority) GetEmbedded() CertificateAuthorityEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret CertificateAuthorityEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthority) GetEmbeddedOk() (*CertificateAuthorityEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *CertificateAuthority) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given CertificateAuthorityEmbedded and assigns it to the Embedded field.
func (o *CertificateAuthority) SetEmbedded(v CertificateAuthorityEmbedded) {
	o.Embedded = &v
}

func (o CertificateAuthority) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthority) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorityInstanceId) {
		toSerialize["authorityInstanceId"] = o.AuthorityInstanceId
	}
	if !IsNil(o.CaType) {
		toSerialize["caType"] = o.CaType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubjectDN) {
		toSerialize["subjectDN"] = o.SubjectDN
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	varCertificateAuthority := _CertificateAuthority{}

	err = json.Unmarshal(data, &varCertificateAuthority)

	if err != nil {
		return err
	}

	*o = CertificateAuthority(varCertificateAuthority)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorityInstanceId")
		delete(additionalProperties, "caType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subjectDN")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_embedded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthority struct {
	value *CertificateAuthority
	isSet bool
}

func (v NullableCertificateAuthority) Get() *CertificateAuthority {
	return v.value
}

func (v *NullableCertificateAuthority) Set(val *CertificateAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthority(val *CertificateAuthority) *NullableCertificateAuthority {
	return &NullableCertificateAuthority{value: val, isSet: true}
}

func (v NullableCertificateAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
