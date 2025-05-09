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

// SecurityEventsProviderSettingsNonSSFCompliant Security Events Provider with issuer and JWKS settings for signal ingestion
type SecurityEventsProviderSettingsNonSSFCompliant struct {
	// Issuer URL
	Issuer string `json:"issuer"`
	// The public URL where the JWKS public key is uploaded
	JwksUrl string `json:"jwks_url"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventsProviderSettingsNonSSFCompliant SecurityEventsProviderSettingsNonSSFCompliant

// NewSecurityEventsProviderSettingsNonSSFCompliant instantiates a new SecurityEventsProviderSettingsNonSSFCompliant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventsProviderSettingsNonSSFCompliant(issuer string, jwksUrl string) *SecurityEventsProviderSettingsNonSSFCompliant {
	this := SecurityEventsProviderSettingsNonSSFCompliant{}
	this.Issuer = issuer
	this.JwksUrl = jwksUrl
	return &this
}

// NewSecurityEventsProviderSettingsNonSSFCompliantWithDefaults instantiates a new SecurityEventsProviderSettingsNonSSFCompliant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventsProviderSettingsNonSSFCompliantWithDefaults() *SecurityEventsProviderSettingsNonSSFCompliant {
	this := SecurityEventsProviderSettingsNonSSFCompliant{}
	return &this
}

// GetIssuer returns the Issuer field value
func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *SecurityEventsProviderSettingsNonSSFCompliant) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwksUrl returns the JwksUrl field value
func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetJwksUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetJwksUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwksUrl, true
}

// SetJwksUrl sets field value
func (o *SecurityEventsProviderSettingsNonSSFCompliant) SetJwksUrl(v string) {
	o.JwksUrl = v
}

func (o SecurityEventsProviderSettingsNonSSFCompliant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["jwks_url"] = o.JwksUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventsProviderSettingsNonSSFCompliant) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventsProviderSettingsNonSSFCompliant := _SecurityEventsProviderSettingsNonSSFCompliant{}

	err = json.Unmarshal(bytes, &varSecurityEventsProviderSettingsNonSSFCompliant)
	if err == nil {
		*o = SecurityEventsProviderSettingsNonSSFCompliant(varSecurityEventsProviderSettingsNonSSFCompliant)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "jwks_url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventsProviderSettingsNonSSFCompliant struct {
	value *SecurityEventsProviderSettingsNonSSFCompliant
	isSet bool
}

func (v NullableSecurityEventsProviderSettingsNonSSFCompliant) Get() *SecurityEventsProviderSettingsNonSSFCompliant {
	return v.value
}

func (v *NullableSecurityEventsProviderSettingsNonSSFCompliant) Set(val *SecurityEventsProviderSettingsNonSSFCompliant) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventsProviderSettingsNonSSFCompliant) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventsProviderSettingsNonSSFCompliant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventsProviderSettingsNonSSFCompliant(val *SecurityEventsProviderSettingsNonSSFCompliant) *NullableSecurityEventsProviderSettingsNonSSFCompliant {
	return &NullableSecurityEventsProviderSettingsNonSSFCompliant{value: val, isSet: true}
}

func (v NullableSecurityEventsProviderSettingsNonSSFCompliant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventsProviderSettingsNonSSFCompliant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

