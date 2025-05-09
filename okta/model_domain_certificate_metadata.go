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

// DomainCertificateMetadata Certificate metadata for the domain
type DomainCertificateMetadata struct {
	// Certificate expiration
	Expiration *string `json:"expiration,omitempty"`
	// Certificate fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Certificate subject
	Subject *string `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DomainCertificateMetadata DomainCertificateMetadata

// NewDomainCertificateMetadata instantiates a new DomainCertificateMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainCertificateMetadata() *DomainCertificateMetadata {
	this := DomainCertificateMetadata{}
	return &this
}

// NewDomainCertificateMetadataWithDefaults instantiates a new DomainCertificateMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainCertificateMetadataWithDefaults() *DomainCertificateMetadata {
	this := DomainCertificateMetadata{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DomainCertificateMetadata) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificateMetadata) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DomainCertificateMetadata) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *DomainCertificateMetadata) SetExpiration(v string) {
	o.Expiration = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *DomainCertificateMetadata) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificateMetadata) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *DomainCertificateMetadata) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *DomainCertificateMetadata) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *DomainCertificateMetadata) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificateMetadata) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *DomainCertificateMetadata) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *DomainCertificateMetadata) SetSubject(v string) {
	o.Subject = &v
}

func (o DomainCertificateMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainCertificateMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varDomainCertificateMetadata := _DomainCertificateMetadata{}

	err = json.Unmarshal(bytes, &varDomainCertificateMetadata)
	if err == nil {
		*o = DomainCertificateMetadata(varDomainCertificateMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomainCertificateMetadata struct {
	value *DomainCertificateMetadata
	isSet bool
}

func (v NullableDomainCertificateMetadata) Get() *DomainCertificateMetadata {
	return v.value
}

func (v *NullableDomainCertificateMetadata) Set(val *DomainCertificateMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainCertificateMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainCertificateMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainCertificateMetadata(val *DomainCertificateMetadata) *NullableDomainCertificateMetadata {
	return &NullableDomainCertificateMetadata{value: val, isSet: true}
}

func (v NullableDomainCertificateMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainCertificateMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

