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

// checks if the OAuth2SettingsPublicKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2SettingsPublicKey{}

// OAuth2SettingsPublicKey The public key in JWK format. Returned when the OAuth authentication method is `PRIVATE_KEY_JWT`.
type OAuth2SettingsPublicKey struct {
	// Key type (e.g., `RSA`, `EC`)
	Kty *string `json:"kty,omitempty"`
	// Key ID
	Kid *string `json:"kid,omitempty"`
	// RSA public exponent
	E *string `json:"e,omitempty"`
	// RSA modulus
	N                    *string `json:"n,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2SettingsPublicKey OAuth2SettingsPublicKey

// NewOAuth2SettingsPublicKey instantiates a new OAuth2SettingsPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2SettingsPublicKey() *OAuth2SettingsPublicKey {
	this := OAuth2SettingsPublicKey{}
	return &this
}

// NewOAuth2SettingsPublicKeyWithDefaults instantiates a new OAuth2SettingsPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2SettingsPublicKeyWithDefaults() *OAuth2SettingsPublicKey {
	this := OAuth2SettingsPublicKey{}
	return &this
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2SettingsPublicKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2SettingsPublicKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2SettingsPublicKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2SettingsPublicKey) SetKty(v string) {
	o.Kty = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *OAuth2SettingsPublicKey) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2SettingsPublicKey) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2SettingsPublicKey) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *OAuth2SettingsPublicKey) SetKid(v string) {
	o.Kid = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OAuth2SettingsPublicKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2SettingsPublicKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OAuth2SettingsPublicKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OAuth2SettingsPublicKey) SetE(v string) {
	o.E = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OAuth2SettingsPublicKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2SettingsPublicKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OAuth2SettingsPublicKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OAuth2SettingsPublicKey) SetN(v string) {
	o.N = &v
}

func (o OAuth2SettingsPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2SettingsPublicKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2SettingsPublicKey) UnmarshalJSON(data []byte) (err error) {
	varOAuth2SettingsPublicKey := _OAuth2SettingsPublicKey{}

	err = json.Unmarshal(data, &varOAuth2SettingsPublicKey)

	if err != nil {
		return err
	}

	*o = OAuth2SettingsPublicKey(varOAuth2SettingsPublicKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kty")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "e")
		delete(additionalProperties, "n")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2SettingsPublicKey struct {
	value *OAuth2SettingsPublicKey
	isSet bool
}

func (v NullableOAuth2SettingsPublicKey) Get() *OAuth2SettingsPublicKey {
	return v.value
}

func (v *NullableOAuth2SettingsPublicKey) Set(val *OAuth2SettingsPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2SettingsPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2SettingsPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2SettingsPublicKey(val *OAuth2SettingsPublicKey) *NullableOAuth2SettingsPublicKey {
	return &NullableOAuth2SettingsPublicKey{value: val, isSet: true}
}

func (v NullableOAuth2SettingsPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2SettingsPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
