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
)

// checks if the AuthenticatorKeyCustomAppAllOfProviderConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorKeyCustomAppAllOfProviderConfiguration{}

// AuthenticatorKeyCustomAppAllOfProviderConfiguration The configuration of the provider
type AuthenticatorKeyCustomAppAllOfProviderConfiguration struct {
	Apns                 *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns `json:"apns,omitempty"`
	Fcm                  *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm  `json:"fcm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyCustomAppAllOfProviderConfiguration AuthenticatorKeyCustomAppAllOfProviderConfiguration

// NewAuthenticatorKeyCustomAppAllOfProviderConfiguration instantiates a new AuthenticatorKeyCustomAppAllOfProviderConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyCustomAppAllOfProviderConfiguration() *AuthenticatorKeyCustomAppAllOfProviderConfiguration {
	this := AuthenticatorKeyCustomAppAllOfProviderConfiguration{}
	return &this
}

// NewAuthenticatorKeyCustomAppAllOfProviderConfigurationWithDefaults instantiates a new AuthenticatorKeyCustomAppAllOfProviderConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyCustomAppAllOfProviderConfigurationWithDefaults() *AuthenticatorKeyCustomAppAllOfProviderConfiguration {
	this := AuthenticatorKeyCustomAppAllOfProviderConfiguration{}
	return &this
}

// GetApns returns the Apns field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) GetApns() AuthenticatorKeyCustomAppAllOfProviderConfigurationApns {
	if o == nil || IsNil(o.Apns) {
		var ret AuthenticatorKeyCustomAppAllOfProviderConfigurationApns
		return ret
	}
	return *o.Apns
}

// GetApnsOk returns a tuple with the Apns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) GetApnsOk() (*AuthenticatorKeyCustomAppAllOfProviderConfigurationApns, bool) {
	if o == nil || IsNil(o.Apns) {
		return nil, false
	}
	return o.Apns, true
}

// HasApns returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) HasApns() bool {
	if o != nil && !IsNil(o.Apns) {
		return true
	}

	return false
}

// SetApns gets a reference to the given AuthenticatorKeyCustomAppAllOfProviderConfigurationApns and assigns it to the Apns field.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) SetApns(v AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) {
	o.Apns = &v
}

// GetFcm returns the Fcm field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) GetFcm() AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm {
	if o == nil || IsNil(o.Fcm) {
		var ret AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm
		return ret
	}
	return *o.Fcm
}

// GetFcmOk returns a tuple with the Fcm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) GetFcmOk() (*AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm, bool) {
	if o == nil || IsNil(o.Fcm) {
		return nil, false
	}
	return o.Fcm, true
}

// HasFcm returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) HasFcm() bool {
	if o != nil && !IsNil(o.Fcm) {
		return true
	}

	return false
}

// SetFcm gets a reference to the given AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm and assigns it to the Fcm field.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) SetFcm(v AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) {
	o.Fcm = &v
}

func (o AuthenticatorKeyCustomAppAllOfProviderConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorKeyCustomAppAllOfProviderConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apns) {
		toSerialize["apns"] = o.Apns
	}
	if !IsNil(o.Fcm) {
		toSerialize["fcm"] = o.Fcm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorKeyCustomAppAllOfProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorKeyCustomAppAllOfProviderConfiguration := _AuthenticatorKeyCustomAppAllOfProviderConfiguration{}

	err = json.Unmarshal(data, &varAuthenticatorKeyCustomAppAllOfProviderConfiguration)

	if err != nil {
		return err
	}

	*o = AuthenticatorKeyCustomAppAllOfProviderConfiguration(varAuthenticatorKeyCustomAppAllOfProviderConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apns")
		delete(additionalProperties, "fcm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration struct {
	value *AuthenticatorKeyCustomAppAllOfProviderConfiguration
	isSet bool
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration) Get() *AuthenticatorKeyCustomAppAllOfProviderConfiguration {
	return v.value
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration) Set(val *AuthenticatorKeyCustomAppAllOfProviderConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyCustomAppAllOfProviderConfiguration(val *AuthenticatorKeyCustomAppAllOfProviderConfiguration) *NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration {
	return &NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
