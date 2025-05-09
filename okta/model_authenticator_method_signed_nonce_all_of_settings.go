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

// AuthenticatorMethodSignedNonceAllOfSettings struct for AuthenticatorMethodSignedNonceAllOfSettings
type AuthenticatorMethodSignedNonceAllOfSettings struct {
	Algorithms []string `json:"algorithms,omitempty"`
	// Indicates whether you must use a hardware key store
	KeyProtection *string `json:"keyProtection,omitempty"`
	// Controls whether to show the Sign in with Okta Verify button on the Sign-In Widget
	ShowSignInWithOV *string `json:"showSignInWithOV,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodSignedNonceAllOfSettings AuthenticatorMethodSignedNonceAllOfSettings

// NewAuthenticatorMethodSignedNonceAllOfSettings instantiates a new AuthenticatorMethodSignedNonceAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodSignedNonceAllOfSettings() *AuthenticatorMethodSignedNonceAllOfSettings {
	this := AuthenticatorMethodSignedNonceAllOfSettings{}
	return &this
}

// NewAuthenticatorMethodSignedNonceAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodSignedNonceAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodSignedNonceAllOfSettingsWithDefaults() *AuthenticatorMethodSignedNonceAllOfSettings {
	this := AuthenticatorMethodSignedNonceAllOfSettings{}
	return &this
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetAlgorithms() []string {
	if o == nil || o.Algorithms == nil {
		var ret []string
		return ret
	}
	return o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetAlgorithmsOk() ([]string, bool) {
	if o == nil || o.Algorithms == nil {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) HasAlgorithms() bool {
	if o != nil && o.Algorithms != nil {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given []string and assigns it to the Algorithms field.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) SetAlgorithms(v []string) {
	o.Algorithms = v
}

// GetKeyProtection returns the KeyProtection field value if set, zero value otherwise.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetKeyProtection() string {
	if o == nil || o.KeyProtection == nil {
		var ret string
		return ret
	}
	return *o.KeyProtection
}

// GetKeyProtectionOk returns a tuple with the KeyProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetKeyProtectionOk() (*string, bool) {
	if o == nil || o.KeyProtection == nil {
		return nil, false
	}
	return o.KeyProtection, true
}

// HasKeyProtection returns a boolean if a field has been set.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) HasKeyProtection() bool {
	if o != nil && o.KeyProtection != nil {
		return true
	}

	return false
}

// SetKeyProtection gets a reference to the given string and assigns it to the KeyProtection field.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) SetKeyProtection(v string) {
	o.KeyProtection = &v
}

// GetShowSignInWithOV returns the ShowSignInWithOV field value if set, zero value otherwise.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetShowSignInWithOV() string {
	if o == nil || o.ShowSignInWithOV == nil {
		var ret string
		return ret
	}
	return *o.ShowSignInWithOV
}

// GetShowSignInWithOVOk returns a tuple with the ShowSignInWithOV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetShowSignInWithOVOk() (*string, bool) {
	if o == nil || o.ShowSignInWithOV == nil {
		return nil, false
	}
	return o.ShowSignInWithOV, true
}

// HasShowSignInWithOV returns a boolean if a field has been set.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) HasShowSignInWithOV() bool {
	if o != nil && o.ShowSignInWithOV != nil {
		return true
	}

	return false
}

// SetShowSignInWithOV gets a reference to the given string and assigns it to the ShowSignInWithOV field.
func (o *AuthenticatorMethodSignedNonceAllOfSettings) SetShowSignInWithOV(v string) {
	o.ShowSignInWithOV = &v
}

func (o AuthenticatorMethodSignedNonceAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithms != nil {
		toSerialize["algorithms"] = o.Algorithms
	}
	if o.KeyProtection != nil {
		toSerialize["keyProtection"] = o.KeyProtection
	}
	if o.ShowSignInWithOV != nil {
		toSerialize["showSignInWithOV"] = o.ShowSignInWithOV
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodSignedNonceAllOfSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorMethodSignedNonceAllOfSettings := _AuthenticatorMethodSignedNonceAllOfSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodSignedNonceAllOfSettings)
	if err == nil {
		*o = AuthenticatorMethodSignedNonceAllOfSettings(varAuthenticatorMethodSignedNonceAllOfSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "keyProtection")
		delete(additionalProperties, "showSignInWithOV")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorMethodSignedNonceAllOfSettings struct {
	value *AuthenticatorMethodSignedNonceAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorMethodSignedNonceAllOfSettings) Get() *AuthenticatorMethodSignedNonceAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorMethodSignedNonceAllOfSettings) Set(val *AuthenticatorMethodSignedNonceAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodSignedNonceAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodSignedNonceAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodSignedNonceAllOfSettings(val *AuthenticatorMethodSignedNonceAllOfSettings) *NullableAuthenticatorMethodSignedNonceAllOfSettings {
	return &NullableAuthenticatorMethodSignedNonceAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodSignedNonceAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodSignedNonceAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

