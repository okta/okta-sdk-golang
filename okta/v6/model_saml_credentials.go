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
)

// SamlCredentials Federation Trust Credentials for verifying assertions from the IdP and signing requests to the IdP
type SamlCredentials struct {
	Signing *SamlSigningCredentials `json:"signing,omitempty"`
	Trust *SamlTrustCredentials `json:"trust,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlCredentials SamlCredentials

// NewSamlCredentials instantiates a new SamlCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlCredentials() *SamlCredentials {
	this := SamlCredentials{}
	return &this
}

// NewSamlCredentialsWithDefaults instantiates a new SamlCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlCredentialsWithDefaults() *SamlCredentials {
	this := SamlCredentials{}
	return &this
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *SamlCredentials) GetSigning() SamlSigningCredentials {
	if o == nil || o.Signing == nil {
		var ret SamlSigningCredentials
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlCredentials) GetSigningOk() (*SamlSigningCredentials, bool) {
	if o == nil || o.Signing == nil {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *SamlCredentials) HasSigning() bool {
	if o != nil && o.Signing != nil {
		return true
	}

	return false
}

// SetSigning gets a reference to the given SamlSigningCredentials and assigns it to the Signing field.
func (o *SamlCredentials) SetSigning(v SamlSigningCredentials) {
	o.Signing = &v
}

// GetTrust returns the Trust field value if set, zero value otherwise.
func (o *SamlCredentials) GetTrust() SamlTrustCredentials {
	if o == nil || o.Trust == nil {
		var ret SamlTrustCredentials
		return ret
	}
	return *o.Trust
}

// GetTrustOk returns a tuple with the Trust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlCredentials) GetTrustOk() (*SamlTrustCredentials, bool) {
	if o == nil || o.Trust == nil {
		return nil, false
	}
	return o.Trust, true
}

// HasTrust returns a boolean if a field has been set.
func (o *SamlCredentials) HasTrust() bool {
	if o != nil && o.Trust != nil {
		return true
	}

	return false
}

// SetTrust gets a reference to the given SamlTrustCredentials and assigns it to the Trust field.
func (o *SamlCredentials) SetTrust(v SamlTrustCredentials) {
	o.Trust = &v
}

func (o SamlCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Signing != nil {
		toSerialize["signing"] = o.Signing
	}
	if o.Trust != nil {
		toSerialize["trust"] = o.Trust
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varSamlCredentials := _SamlCredentials{}

	err = json.Unmarshal(bytes, &varSamlCredentials)
	if err == nil {
		*o = SamlCredentials(varSamlCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signing")
		delete(additionalProperties, "trust")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlCredentials struct {
	value *SamlCredentials
	isSet bool
}

func (v NullableSamlCredentials) Get() *SamlCredentials {
	return v.value
}

func (v *NullableSamlCredentials) Set(val *SamlCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlCredentials(val *SamlCredentials) *NullableSamlCredentials {
	return &NullableSamlCredentials{value: val, isSet: true}
}

func (v NullableSamlCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

