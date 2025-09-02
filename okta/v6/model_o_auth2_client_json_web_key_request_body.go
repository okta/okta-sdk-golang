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

// OAuth2ClientJsonWebKeyRequestBody struct for OAuth2ClientJsonWebKeyRequestBody
type OAuth2ClientJsonWebKeyRequestBody struct {
	// Algorithm used in the key
	Alg *string `json:"alg,omitempty"`
	// RSA key value (exponent) for key binding
	E *string `json:"e,omitempty"`
	// Unique identifier of the JSON Web Key in the OAUth 2.0 Client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N *string `json:"n,omitempty"`
	// Status of the OAuth 2.0 Client JSON Web Key
	Status *string `json:"status,omitempty"`
	// Acceptable use of the JSON Web Key
	Use *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyRequestBody OAuth2ClientJsonWebKeyRequestBody

// NewOAuth2ClientJsonWebKeyRequestBody instantiates a new OAuth2ClientJsonWebKeyRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyRequestBody() *OAuth2ClientJsonWebKeyRequestBody {
	this := OAuth2ClientJsonWebKeyRequestBody{}
	return &this
}

// NewOAuth2ClientJsonWebKeyRequestBodyWithDefaults instantiates a new OAuth2ClientJsonWebKeyRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyRequestBodyWithDefaults() *OAuth2ClientJsonWebKeyRequestBody {
	this := OAuth2ClientJsonWebKeyRequestBody{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetAlg(v string) {
	o.Alg = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetE() string {
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasE() bool {
	if o != nil && o.E != nil {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetE(v string) {
	o.E = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonWebKeyRequestBody) GetKid() string {
	if o == nil || o.Kid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonWebKeyRequestBody) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetKid(v string) {
	o.Kid.Set(&v)
}
// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonWebKeyRequestBody) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonWebKeyRequestBody) UnsetKid() {
	o.Kid.Unset()
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetN() string {
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetN(v string) {
	o.N = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRequestBody) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *OAuth2ClientJsonWebKeyRequestBody) SetUse(v string) {
	o.Use = &v
}

func (o OAuth2ClientJsonWebKeyRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ClientJsonWebKeyRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientJsonWebKeyRequestBody := _OAuth2ClientJsonWebKeyRequestBody{}

	err = json.Unmarshal(bytes, &varOAuth2ClientJsonWebKeyRequestBody)
	if err == nil {
		*o = OAuth2ClientJsonWebKeyRequestBody(varOAuth2ClientJsonWebKeyRequestBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "e")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ClientJsonWebKeyRequestBody struct {
	value *OAuth2ClientJsonWebKeyRequestBody
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyRequestBody) Get() *OAuth2ClientJsonWebKeyRequestBody {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyRequestBody) Set(val *OAuth2ClientJsonWebKeyRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyRequestBody(val *OAuth2ClientJsonWebKeyRequestBody) *NullableOAuth2ClientJsonWebKeyRequestBody {
	return &NullableOAuth2ClientJsonWebKeyRequestBody{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

