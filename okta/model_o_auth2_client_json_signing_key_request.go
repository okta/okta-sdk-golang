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
	"fmt"
)

// checks if the OAuth2ClientJsonSigningKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonSigningKeyRequest{}

// OAuth2ClientJsonSigningKeyRequest A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses signing keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys for signing tokens.
type OAuth2ClientJsonSigningKeyRequest struct {
	// Algorithm used in the key
	Alg string `json:"alg"`
	// Unique identifier of the JSON Web Key in the OAuth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty string `json:"kty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	// Acceptable use of the JSON Web Key
	Use                  string `json:"use"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonSigningKeyRequest OAuth2ClientJsonSigningKeyRequest

// NewOAuth2ClientJsonSigningKeyRequest instantiates a new OAuth2ClientJsonSigningKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonSigningKeyRequest(alg string, kty string, use string) *OAuth2ClientJsonSigningKeyRequest {
	this := OAuth2ClientJsonSigningKeyRequest{}
	this.Alg = alg
	this.Kty = kty
	var status string = "ACTIVE"
	this.Status = &status
	this.Use = use
	return &this
}

// NewOAuth2ClientJsonSigningKeyRequestWithDefaults instantiates a new OAuth2ClientJsonSigningKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonSigningKeyRequestWithDefaults() *OAuth2ClientJsonSigningKeyRequest {
	this := OAuth2ClientJsonSigningKeyRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetAlg returns the Alg field value
func (o *OAuth2ClientJsonSigningKeyRequest) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyRequest) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *OAuth2ClientJsonSigningKeyRequest) SetAlg(v string) {
	o.Alg = v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonSigningKeyRequest) GetKid() string {
	if o == nil || IsNil(o.Kid.Get()) {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonSigningKeyRequest) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonSigningKeyRequest) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonSigningKeyRequest) SetKid(v string) {
	o.Kid.Set(&v)
}

// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonSigningKeyRequest) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonSigningKeyRequest) UnsetKid() {
	o.Kid.Unset()
}

// GetKty returns the Kty field value
func (o *OAuth2ClientJsonSigningKeyRequest) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyRequest) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *OAuth2ClientJsonSigningKeyRequest) SetKty(v string) {
	o.Kty = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonSigningKeyRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonSigningKeyRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonSigningKeyRequest) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value
func (o *OAuth2ClientJsonSigningKeyRequest) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonSigningKeyRequest) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *OAuth2ClientJsonSigningKeyRequest) SetUse(v string) {
	o.Use = v
}

func (o OAuth2ClientJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonSigningKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alg"] = o.Alg
	if o.Kid.IsSet() {
		toSerialize["kid"] = o.Kid.Get()
	}
	toSerialize["kty"] = o.Kty
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["use"] = o.Use

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientJsonSigningKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alg",
		"kty",
		"use",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOAuth2ClientJsonSigningKeyRequest := _OAuth2ClientJsonSigningKeyRequest{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonSigningKeyRequest)

	if err != nil {
		return err
	}

	*o = OAuth2ClientJsonSigningKeyRequest(varOAuth2ClientJsonSigningKeyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonSigningKeyRequest struct {
	value *OAuth2ClientJsonSigningKeyRequest
	isSet bool
}

func (v NullableOAuth2ClientJsonSigningKeyRequest) Get() *OAuth2ClientJsonSigningKeyRequest {
	return v.value
}

func (v *NullableOAuth2ClientJsonSigningKeyRequest) Set(val *OAuth2ClientJsonSigningKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonSigningKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonSigningKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonSigningKeyRequest(val *OAuth2ClientJsonSigningKeyRequest) *NullableOAuth2ClientJsonSigningKeyRequest {
	return &NullableOAuth2ClientJsonSigningKeyRequest{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonSigningKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
