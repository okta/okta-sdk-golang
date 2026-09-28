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
	"time"
)

// checks if the CertificateAuthorityCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityCertificate{}

// CertificateAuthorityCertificate One generation of a certificate authority — a distinct key pair, X.509 certificate, and expiry. Renewal adds a new one to the same authority.
type CertificateAuthorityCertificate struct {
	// The authority key identifier of this certificate — the `ski` of the CA certificate that signed it. This is how the trust chain is walked. A self-signed root has no parent.
	Aki *string `json:"aki,omitempty"`
	// The stable identifier of the authority this certificate belongs to
	AuthorityInstanceId *string `json:"authorityInstanceId,omitempty"`
	// The family of the certificate authority. `OKTA_AS_CA` is the Okta CA used for SCEP enrollment and device management attestation; `DEVICE_ACCESS_OKTA_AS_CA` is the Device Access CA. The `caType` query parameter filters on the same values.
	CaType *string `json:"caType,omitempty"`
	// Timestamp when the object was created
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Timestamp when this certificate expires
	ExpirationDate *time.Time                      `json:"expirationDate,omitempty"`
	Jwk            *CertificateAuthorityJsonWebKey `json:"jwk,omitempty"`
	// Timestamp when the object was last updated
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// The rotation stage of the CA certificate. `ACTIVE` indicates the authority's newest certificate, which is the one that a renewal migrates SCEP configurations onto. `RETIRING` indicates a certificate that has been superseded by a renewal but remains valid, allowing enrolled devices to continue functioning during migration.
	RotationState *string `json:"rotationState,omitempty"`
	// The subject key identifier of this certificate, in uppercase hexadecimal. It identifies the certificate within its authority.
	Ski *string `json:"ski,omitempty"`
	// The status of the certificate. `VALID` is in service and able to issue certificates, which is how Okta generates a replacement CA certificate. `INACTIVE` is a replacement that Okta generated but hasn't brought into service, which occurs only for an authority renewed before selective renewal was available; renewing such an authority activates it.
	Status *string `json:"status,omitempty"`
	// The distinguished name in the subject of this certificate
	SubjectDN *string `json:"subjectDN,omitempty"`
	// The position of the certificate authority in the hierarchy. A root is self-signed; an intermediate is signed by the root and signs the certificates Okta issues to devices. The `type` query parameter filters on the same values.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityCertificate CertificateAuthorityCertificate

// NewCertificateAuthorityCertificate instantiates a new CertificateAuthorityCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityCertificate() *CertificateAuthorityCertificate {
	this := CertificateAuthorityCertificate{}
	return &this
}

// NewCertificateAuthorityCertificateWithDefaults instantiates a new CertificateAuthorityCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityCertificateWithDefaults() *CertificateAuthorityCertificate {
	this := CertificateAuthorityCertificate{}
	return &this
}

// GetAki returns the Aki field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetAki() string {
	if o == nil || IsNil(o.Aki) {
		var ret string
		return ret
	}
	return *o.Aki
}

// GetAkiOk returns a tuple with the Aki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetAkiOk() (*string, bool) {
	if o == nil || IsNil(o.Aki) {
		return nil, false
	}
	return o.Aki, true
}

// HasAki returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasAki() bool {
	if o != nil && !IsNil(o.Aki) {
		return true
	}

	return false
}

// SetAki gets a reference to the given string and assigns it to the Aki field.
func (o *CertificateAuthorityCertificate) SetAki(v string) {
	o.Aki = &v
}

// GetAuthorityInstanceId returns the AuthorityInstanceId field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetAuthorityInstanceId() string {
	if o == nil || IsNil(o.AuthorityInstanceId) {
		var ret string
		return ret
	}
	return *o.AuthorityInstanceId
}

// GetAuthorityInstanceIdOk returns a tuple with the AuthorityInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetAuthorityInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityInstanceId) {
		return nil, false
	}
	return o.AuthorityInstanceId, true
}

// HasAuthorityInstanceId returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasAuthorityInstanceId() bool {
	if o != nil && !IsNil(o.AuthorityInstanceId) {
		return true
	}

	return false
}

// SetAuthorityInstanceId gets a reference to the given string and assigns it to the AuthorityInstanceId field.
func (o *CertificateAuthorityCertificate) SetAuthorityInstanceId(v string) {
	o.AuthorityInstanceId = &v
}

// GetCaType returns the CaType field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetCaType() string {
	if o == nil || IsNil(o.CaType) {
		var ret string
		return ret
	}
	return *o.CaType
}

// GetCaTypeOk returns a tuple with the CaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetCaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CaType) {
		return nil, false
	}
	return o.CaType, true
}

// HasCaType returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasCaType() bool {
	if o != nil && !IsNil(o.CaType) {
		return true
	}

	return false
}

// SetCaType gets a reference to the given string and assigns it to the CaType field.
func (o *CertificateAuthorityCertificate) SetCaType(v string) {
	o.CaType = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *CertificateAuthorityCertificate) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *CertificateAuthorityCertificate) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetJwk returns the Jwk field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetJwk() CertificateAuthorityJsonWebKey {
	if o == nil || IsNil(o.Jwk) {
		var ret CertificateAuthorityJsonWebKey
		return ret
	}
	return *o.Jwk
}

// GetJwkOk returns a tuple with the Jwk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetJwkOk() (*CertificateAuthorityJsonWebKey, bool) {
	if o == nil || IsNil(o.Jwk) {
		return nil, false
	}
	return o.Jwk, true
}

// HasJwk returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasJwk() bool {
	if o != nil && !IsNil(o.Jwk) {
		return true
	}

	return false
}

// SetJwk gets a reference to the given CertificateAuthorityJsonWebKey and assigns it to the Jwk field.
func (o *CertificateAuthorityCertificate) SetJwk(v CertificateAuthorityJsonWebKey) {
	o.Jwk = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *CertificateAuthorityCertificate) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetRotationState returns the RotationState field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetRotationState() string {
	if o == nil || IsNil(o.RotationState) {
		var ret string
		return ret
	}
	return *o.RotationState
}

// GetRotationStateOk returns a tuple with the RotationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetRotationStateOk() (*string, bool) {
	if o == nil || IsNil(o.RotationState) {
		return nil, false
	}
	return o.RotationState, true
}

// HasRotationState returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasRotationState() bool {
	if o != nil && !IsNil(o.RotationState) {
		return true
	}

	return false
}

// SetRotationState gets a reference to the given string and assigns it to the RotationState field.
func (o *CertificateAuthorityCertificate) SetRotationState(v string) {
	o.RotationState = &v
}

// GetSki returns the Ski field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetSki() string {
	if o == nil || IsNil(o.Ski) {
		var ret string
		return ret
	}
	return *o.Ski
}

// GetSkiOk returns a tuple with the Ski field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetSkiOk() (*string, bool) {
	if o == nil || IsNil(o.Ski) {
		return nil, false
	}
	return o.Ski, true
}

// HasSki returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasSki() bool {
	if o != nil && !IsNil(o.Ski) {
		return true
	}

	return false
}

// SetSki gets a reference to the given string and assigns it to the Ski field.
func (o *CertificateAuthorityCertificate) SetSki(v string) {
	o.Ski = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertificateAuthorityCertificate) SetStatus(v string) {
	o.Status = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetSubjectDN() string {
	if o == nil || IsNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasSubjectDN() bool {
	if o != nil && !IsNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *CertificateAuthorityCertificate) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CertificateAuthorityCertificate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityCertificate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CertificateAuthorityCertificate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CertificateAuthorityCertificate) SetType(v string) {
	o.Type = &v
}

func (o CertificateAuthorityCertificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aki) {
		toSerialize["aki"] = o.Aki
	}
	if !IsNil(o.AuthorityInstanceId) {
		toSerialize["authorityInstanceId"] = o.AuthorityInstanceId
	}
	if !IsNil(o.CaType) {
		toSerialize["caType"] = o.CaType
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Jwk) {
		toSerialize["jwk"] = o.Jwk
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if !IsNil(o.RotationState) {
		toSerialize["rotationState"] = o.RotationState
	}
	if !IsNil(o.Ski) {
		toSerialize["ski"] = o.Ski
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityCertificate) UnmarshalJSON(data []byte) (err error) {
	varCertificateAuthorityCertificate := _CertificateAuthorityCertificate{}

	err = json.Unmarshal(data, &varCertificateAuthorityCertificate)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityCertificate(varCertificateAuthorityCertificate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aki")
		delete(additionalProperties, "authorityInstanceId")
		delete(additionalProperties, "caType")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "jwk")
		delete(additionalProperties, "lastUpdatedDate")
		delete(additionalProperties, "rotationState")
		delete(additionalProperties, "ski")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subjectDN")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityCertificate struct {
	value *CertificateAuthorityCertificate
	isSet bool
}

func (v NullableCertificateAuthorityCertificate) Get() *CertificateAuthorityCertificate {
	return v.value
}

func (v *NullableCertificateAuthorityCertificate) Set(val *CertificateAuthorityCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityCertificate(val *CertificateAuthorityCertificate) *NullableCertificateAuthorityCertificate {
	return &NullableCertificateAuthorityCertificate{value: val, isSet: true}
}

func (v NullableCertificateAuthorityCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
