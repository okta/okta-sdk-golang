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

// checks if the SelectiveRenewalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectiveRenewalResponse{}

// SelectiveRenewalResponse The certificate the authority renewed onto, and what happened to each SCEP configuration named in the request
type SelectiveRenewalResponse struct {
	// The configurations created against the renewed certificate, each naming the configuration it was cloned from in its own `sourceConfigId`. Empty if the request asked for none.
	NewParallelConfigs []RegistrationAuthority          `json:"newParallelConfigs,omitempty"`
	RenewedCertificate *CertificateAuthorityCertificate `json:"renewedCertificate,omitempty"`
	// The configurations now bound to the renewed certificate. Empty if the request rolled none over.
	RolledOverConfigs    []RegistrationAuthority `json:"rolledOverConfigs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelectiveRenewalResponse SelectiveRenewalResponse

// NewSelectiveRenewalResponse instantiates a new SelectiveRenewalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectiveRenewalResponse() *SelectiveRenewalResponse {
	this := SelectiveRenewalResponse{}
	return &this
}

// NewSelectiveRenewalResponseWithDefaults instantiates a new SelectiveRenewalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectiveRenewalResponseWithDefaults() *SelectiveRenewalResponse {
	this := SelectiveRenewalResponse{}
	return &this
}

// GetNewParallelConfigs returns the NewParallelConfigs field value if set, zero value otherwise.
func (o *SelectiveRenewalResponse) GetNewParallelConfigs() []RegistrationAuthority {
	if o == nil || IsNil(o.NewParallelConfigs) {
		var ret []RegistrationAuthority
		return ret
	}
	return o.NewParallelConfigs
}

// GetNewParallelConfigsOk returns a tuple with the NewParallelConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveRenewalResponse) GetNewParallelConfigsOk() ([]RegistrationAuthority, bool) {
	if o == nil || IsNil(o.NewParallelConfigs) {
		return nil, false
	}
	return o.NewParallelConfigs, true
}

// HasNewParallelConfigs returns a boolean if a field has been set.
func (o *SelectiveRenewalResponse) HasNewParallelConfigs() bool {
	if o != nil && !IsNil(o.NewParallelConfigs) {
		return true
	}

	return false
}

// SetNewParallelConfigs gets a reference to the given []RegistrationAuthority and assigns it to the NewParallelConfigs field.
func (o *SelectiveRenewalResponse) SetNewParallelConfigs(v []RegistrationAuthority) {
	o.NewParallelConfigs = v
}

// GetRenewedCertificate returns the RenewedCertificate field value if set, zero value otherwise.
func (o *SelectiveRenewalResponse) GetRenewedCertificate() CertificateAuthorityCertificate {
	if o == nil || IsNil(o.RenewedCertificate) {
		var ret CertificateAuthorityCertificate
		return ret
	}
	return *o.RenewedCertificate
}

// GetRenewedCertificateOk returns a tuple with the RenewedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveRenewalResponse) GetRenewedCertificateOk() (*CertificateAuthorityCertificate, bool) {
	if o == nil || IsNil(o.RenewedCertificate) {
		return nil, false
	}
	return o.RenewedCertificate, true
}

// HasRenewedCertificate returns a boolean if a field has been set.
func (o *SelectiveRenewalResponse) HasRenewedCertificate() bool {
	if o != nil && !IsNil(o.RenewedCertificate) {
		return true
	}

	return false
}

// SetRenewedCertificate gets a reference to the given CertificateAuthorityCertificate and assigns it to the RenewedCertificate field.
func (o *SelectiveRenewalResponse) SetRenewedCertificate(v CertificateAuthorityCertificate) {
	o.RenewedCertificate = &v
}

// GetRolledOverConfigs returns the RolledOverConfigs field value if set, zero value otherwise.
func (o *SelectiveRenewalResponse) GetRolledOverConfigs() []RegistrationAuthority {
	if o == nil || IsNil(o.RolledOverConfigs) {
		var ret []RegistrationAuthority
		return ret
	}
	return o.RolledOverConfigs
}

// GetRolledOverConfigsOk returns a tuple with the RolledOverConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveRenewalResponse) GetRolledOverConfigsOk() ([]RegistrationAuthority, bool) {
	if o == nil || IsNil(o.RolledOverConfigs) {
		return nil, false
	}
	return o.RolledOverConfigs, true
}

// HasRolledOverConfigs returns a boolean if a field has been set.
func (o *SelectiveRenewalResponse) HasRolledOverConfigs() bool {
	if o != nil && !IsNil(o.RolledOverConfigs) {
		return true
	}

	return false
}

// SetRolledOverConfigs gets a reference to the given []RegistrationAuthority and assigns it to the RolledOverConfigs field.
func (o *SelectiveRenewalResponse) SetRolledOverConfigs(v []RegistrationAuthority) {
	o.RolledOverConfigs = v
}

func (o SelectiveRenewalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectiveRenewalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewParallelConfigs) {
		toSerialize["newParallelConfigs"] = o.NewParallelConfigs
	}
	if !IsNil(o.RenewedCertificate) {
		toSerialize["renewedCertificate"] = o.RenewedCertificate
	}
	if !IsNil(o.RolledOverConfigs) {
		toSerialize["rolledOverConfigs"] = o.RolledOverConfigs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelectiveRenewalResponse) UnmarshalJSON(data []byte) (err error) {
	varSelectiveRenewalResponse := _SelectiveRenewalResponse{}

	err = json.Unmarshal(data, &varSelectiveRenewalResponse)

	if err != nil {
		return err
	}

	*o = SelectiveRenewalResponse(varSelectiveRenewalResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "newParallelConfigs")
		delete(additionalProperties, "renewedCertificate")
		delete(additionalProperties, "rolledOverConfigs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelectiveRenewalResponse struct {
	value *SelectiveRenewalResponse
	isSet bool
}

func (v NullableSelectiveRenewalResponse) Get() *SelectiveRenewalResponse {
	return v.value
}

func (v *NullableSelectiveRenewalResponse) Set(val *SelectiveRenewalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectiveRenewalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectiveRenewalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectiveRenewalResponse(val *SelectiveRenewalResponse) *NullableSelectiveRenewalResponse {
	return &NullableSelectiveRenewalResponse{value: val, isSet: true}
}

func (v NullableSelectiveRenewalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectiveRenewalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
