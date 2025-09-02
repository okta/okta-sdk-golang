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

// OAuth2ClientJsonEncryptionKeyRequest <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses an encryption key to encrypt an ID token JWT minted by the org authorization server or custom authorization server. Okta supports only RSA keys for encrypting tokens.
type OAuth2ClientJsonEncryptionKeyRequest struct {
	// RSA key value (exponent) for key binding
	E *string `json:"e,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N *string `json:"n,omitempty"`
	// Acceptable use of the JSON Web Key
	Use *string `json:"use,omitempty"`
	// Unique identifier of the JSON Web Key in the OAUth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonEncryptionKeyRequest OAuth2ClientJsonEncryptionKeyRequest

// NewOAuth2ClientJsonEncryptionKeyRequest instantiates a new OAuth2ClientJsonEncryptionKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonEncryptionKeyRequest() *OAuth2ClientJsonEncryptionKeyRequest {
	this := OAuth2ClientJsonEncryptionKeyRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ClientJsonEncryptionKeyRequestWithDefaults instantiates a new OAuth2ClientJsonEncryptionKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonEncryptionKeyRequestWithDefaults() *OAuth2ClientJsonEncryptionKeyRequest {
	this := OAuth2ClientJsonEncryptionKeyRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetE() string {
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) HasE() bool {
	if o != nil && o.E != nil {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetE(v string) {
	o.E = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetN() string {
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetUse(v string) {
	o.Use = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetKid() string {
	if o == nil || o.Kid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetKid(v string) {
	o.Kid.Set(&v)
}
// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonEncryptionKeyRequest) UnsetKid() {
	o.Kid.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonEncryptionKeyRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonEncryptionKeyRequest) SetStatus(v string) {
	o.Status = &v
}

func (o OAuth2ClientJsonEncryptionKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClientJsonEncryptionKeyRequest) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientJsonEncryptionKeyRequest := _OAuth2ClientJsonEncryptionKeyRequest{}

	err = json.Unmarshal(bytes, &varOAuth2ClientJsonEncryptionKeyRequest)
	if err == nil {
		*o = OAuth2ClientJsonEncryptionKeyRequest(varOAuth2ClientJsonEncryptionKeyRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ClientJsonEncryptionKeyRequest struct {
	value *OAuth2ClientJsonEncryptionKeyRequest
	isSet bool
}

func (v NullableOAuth2ClientJsonEncryptionKeyRequest) Get() *OAuth2ClientJsonEncryptionKeyRequest {
	return v.value
}

func (v *NullableOAuth2ClientJsonEncryptionKeyRequest) Set(val *OAuth2ClientJsonEncryptionKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonEncryptionKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonEncryptionKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonEncryptionKeyRequest(val *OAuth2ClientJsonEncryptionKeyRequest) *NullableOAuth2ClientJsonEncryptionKeyRequest {
	return &NullableOAuth2ClientJsonEncryptionKeyRequest{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonEncryptionKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonEncryptionKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

