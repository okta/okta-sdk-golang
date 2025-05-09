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

// SecurityEventsProviderSettingsResponse Security Events Provider settings
type SecurityEventsProviderSettingsResponse struct {
	// Issuer URL
	Issuer *string `json:"issuer,omitempty"`
	// The public URL where the JWKS public key is uploaded
	JwksUrl *string `json:"jwks_url,omitempty"`
	// The well-known URL of the Security Events Provider (the SSF transmitter)
	WellKnownUrl NullableString `json:"well_known_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventsProviderSettingsResponse SecurityEventsProviderSettingsResponse

// NewSecurityEventsProviderSettingsResponse instantiates a new SecurityEventsProviderSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventsProviderSettingsResponse() *SecurityEventsProviderSettingsResponse {
	this := SecurityEventsProviderSettingsResponse{}
	return &this
}

// NewSecurityEventsProviderSettingsResponseWithDefaults instantiates a new SecurityEventsProviderSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventsProviderSettingsResponseWithDefaults() *SecurityEventsProviderSettingsResponse {
	this := SecurityEventsProviderSettingsResponse{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SecurityEventsProviderSettingsResponse) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderSettingsResponse) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SecurityEventsProviderSettingsResponse) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SecurityEventsProviderSettingsResponse) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *SecurityEventsProviderSettingsResponse) GetJwksUrl() string {
	if o == nil || o.JwksUrl == nil {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderSettingsResponse) GetJwksUrlOk() (*string, bool) {
	if o == nil || o.JwksUrl == nil {
		return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *SecurityEventsProviderSettingsResponse) HasJwksUrl() bool {
	if o != nil && o.JwksUrl != nil {
		return true
	}

	return false
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *SecurityEventsProviderSettingsResponse) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

// GetWellKnownUrl returns the WellKnownUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityEventsProviderSettingsResponse) GetWellKnownUrl() string {
	if o == nil || o.WellKnownUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WellKnownUrl.Get()
}

// GetWellKnownUrlOk returns a tuple with the WellKnownUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityEventsProviderSettingsResponse) GetWellKnownUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WellKnownUrl.Get(), o.WellKnownUrl.IsSet()
}

// HasWellKnownUrl returns a boolean if a field has been set.
func (o *SecurityEventsProviderSettingsResponse) HasWellKnownUrl() bool {
	if o != nil && o.WellKnownUrl.IsSet() {
		return true
	}

	return false
}

// SetWellKnownUrl gets a reference to the given NullableString and assigns it to the WellKnownUrl field.
func (o *SecurityEventsProviderSettingsResponse) SetWellKnownUrl(v string) {
	o.WellKnownUrl.Set(&v)
}
// SetWellKnownUrlNil sets the value for WellKnownUrl to be an explicit nil
func (o *SecurityEventsProviderSettingsResponse) SetWellKnownUrlNil() {
	o.WellKnownUrl.Set(nil)
}

// UnsetWellKnownUrl ensures that no value is present for WellKnownUrl, not even an explicit nil
func (o *SecurityEventsProviderSettingsResponse) UnsetWellKnownUrl() {
	o.WellKnownUrl.Unset()
}

func (o SecurityEventsProviderSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwksUrl != nil {
		toSerialize["jwks_url"] = o.JwksUrl
	}
	if o.WellKnownUrl.IsSet() {
		toSerialize["well_known_url"] = o.WellKnownUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventsProviderSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventsProviderSettingsResponse := _SecurityEventsProviderSettingsResponse{}

	err = json.Unmarshal(bytes, &varSecurityEventsProviderSettingsResponse)
	if err == nil {
		*o = SecurityEventsProviderSettingsResponse(varSecurityEventsProviderSettingsResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "jwks_url")
		delete(additionalProperties, "well_known_url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventsProviderSettingsResponse struct {
	value *SecurityEventsProviderSettingsResponse
	isSet bool
}

func (v NullableSecurityEventsProviderSettingsResponse) Get() *SecurityEventsProviderSettingsResponse {
	return v.value
}

func (v *NullableSecurityEventsProviderSettingsResponse) Set(val *SecurityEventsProviderSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventsProviderSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventsProviderSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventsProviderSettingsResponse(val *SecurityEventsProviderSettingsResponse) *NullableSecurityEventsProviderSettingsResponse {
	return &NullableSecurityEventsProviderSettingsResponse{value: val, isSet: true}
}

func (v NullableSecurityEventsProviderSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventsProviderSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

