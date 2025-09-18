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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OAuth2ResourceServerJsonWebKeyRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ResourceServerJsonWebKeyRequestBody{}

// OAuth2ResourceServerJsonWebKeyRequestBody struct for OAuth2ResourceServerJsonWebKeyRequestBody
type OAuth2ResourceServerJsonWebKeyRequestBody struct {
	// RSA key value (exponent) for key binding
	E *string `json:"e,omitempty"`
	// Unique identifier of the JSON web key in the custom authorization server's public JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N *string `json:"n,omitempty"`
	// Status of the JSON Web Key
	Status *string `json:"status,omitempty"`
	// Acceptable use of the JSON Web Key
	Use                  *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ResourceServerJsonWebKeyRequestBody OAuth2ResourceServerJsonWebKeyRequestBody

// NewOAuth2ResourceServerJsonWebKeyRequestBody instantiates a new OAuth2ResourceServerJsonWebKeyRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ResourceServerJsonWebKeyRequestBody() *OAuth2ResourceServerJsonWebKeyRequestBody {
	this := OAuth2ResourceServerJsonWebKeyRequestBody{}
	return &this
}

// NewOAuth2ResourceServerJsonWebKeyRequestBodyWithDefaults instantiates a new OAuth2ResourceServerJsonWebKeyRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ResourceServerJsonWebKeyRequestBodyWithDefaults() *OAuth2ResourceServerJsonWebKeyRequestBody {
	this := OAuth2ResourceServerJsonWebKeyRequestBody{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetE(v string) {
	o.E = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKid() string {
	if o == nil || IsNil(o.Kid.Get()) {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetKid(v string) {
	o.Kid.Set(&v)
}

// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) UnsetKid() {
	o.Kid.Unset()
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetN(v string) {
	o.N = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetUse(v string) {
	o.Use = &v
}

func (o OAuth2ResourceServerJsonWebKeyRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ResourceServerJsonWebKeyRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ResourceServerJsonWebKeyRequestBody) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ResourceServerJsonWebKeyRequestBody := _OAuth2ResourceServerJsonWebKeyRequestBody{}

	err = json.Unmarshal(data, &varOAuth2ResourceServerJsonWebKeyRequestBody)

	if err != nil {
		return err
	}

	*o = OAuth2ResourceServerJsonWebKeyRequestBody(varOAuth2ResourceServerJsonWebKeyRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ResourceServerJsonWebKeyRequestBody struct {
	value *OAuth2ResourceServerJsonWebKeyRequestBody
	isSet bool
}

func (v NullableOAuth2ResourceServerJsonWebKeyRequestBody) Get() *OAuth2ResourceServerJsonWebKeyRequestBody {
	return v.value
}

func (v *NullableOAuth2ResourceServerJsonWebKeyRequestBody) Set(val *OAuth2ResourceServerJsonWebKeyRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ResourceServerJsonWebKeyRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ResourceServerJsonWebKeyRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ResourceServerJsonWebKeyRequestBody(val *OAuth2ResourceServerJsonWebKeyRequestBody) *NullableOAuth2ResourceServerJsonWebKeyRequestBody {
	return &NullableOAuth2ResourceServerJsonWebKeyRequestBody{value: val, isSet: true}
}

func (v NullableOAuth2ResourceServerJsonWebKeyRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ResourceServerJsonWebKeyRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
