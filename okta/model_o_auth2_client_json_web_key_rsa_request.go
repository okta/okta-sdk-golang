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
	"reflect"
	"strings"
)

// checks if the OAuth2ClientJsonWebKeyRsaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeyRsaRequest{}

// OAuth2ClientJsonWebKeyRsaRequest An RSA signing key
type OAuth2ClientJsonWebKeyRsaRequest struct {
	OAuth2ClientJsonSigningKeyRequest
	// RSA key value (exponent) for key binding
	E string `json:"e"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N                    string `json:"n"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyRsaRequest OAuth2ClientJsonWebKeyRsaRequest

// NewOAuth2ClientJsonWebKeyRsaRequest instantiates a new OAuth2ClientJsonWebKeyRsaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyRsaRequest(e string, n string, alg string, use string) *OAuth2ClientJsonWebKeyRsaRequest {
	this := OAuth2ClientJsonWebKeyRsaRequest{}
	this.Alg = alg
	var status string = "ACTIVE"
	this.Status = &status
	this.Use = use
	this.E = e
	this.N = n
	return &this
}

// NewOAuth2ClientJsonWebKeyRsaRequestWithDefaults instantiates a new OAuth2ClientJsonWebKeyRsaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyRsaRequestWithDefaults() *OAuth2ClientJsonWebKeyRsaRequest {
	this := OAuth2ClientJsonWebKeyRsaRequest{}
	return &this
}

// GetE returns the E field value
func (o *OAuth2ClientJsonWebKeyRsaRequest) GetE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRsaRequest) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *OAuth2ClientJsonWebKeyRsaRequest) SetE(v string) {
	o.E = v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRsaRequest) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRsaRequest) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRsaRequest) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyRsaRequest) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value
func (o *OAuth2ClientJsonWebKeyRsaRequest) GetN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRsaRequest) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N, true
}

// SetN sets field value
func (o *OAuth2ClientJsonWebKeyRsaRequest) SetN(v string) {
	o.N = v
}

func (o OAuth2ClientJsonWebKeyRsaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeyRsaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOAuth2ClientJsonSigningKeyRequest, errOAuth2ClientJsonSigningKeyRequest := json.Marshal(o.OAuth2ClientJsonSigningKeyRequest)
	if errOAuth2ClientJsonSigningKeyRequest != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyRequest
	}
	errOAuth2ClientJsonSigningKeyRequest = json.Unmarshal([]byte(serializedOAuth2ClientJsonSigningKeyRequest), &toSerialize)
	if errOAuth2ClientJsonSigningKeyRequest != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyRequest
	}
	toSerialize["e"] = o.E
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	toSerialize["n"] = o.N

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientJsonWebKeyRsaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"e",
		"n",
		"alg",
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

	type OAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct struct {
		// RSA key value (exponent) for key binding
		E string `json:"e"`
		// Cryptographic algorithm family for the certificate's key pair
		Kty *string `json:"kty,omitempty"`
		// RSA key value (modulus) for key binding
		N string `json:"n"`
	}

	varOAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct := OAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct)
	if err == nil {
		varOAuth2ClientJsonWebKeyRsaRequest := _OAuth2ClientJsonWebKeyRsaRequest{}
		varOAuth2ClientJsonWebKeyRsaRequest.E = varOAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct.E
		varOAuth2ClientJsonWebKeyRsaRequest.Kty = varOAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct.Kty
		varOAuth2ClientJsonWebKeyRsaRequest.N = varOAuth2ClientJsonWebKeyRsaRequestWithoutEmbeddedStruct.N
		*o = OAuth2ClientJsonWebKeyRsaRequest(varOAuth2ClientJsonWebKeyRsaRequest)
	} else {
		return err
	}

	varOAuth2ClientJsonWebKeyRsaRequest := _OAuth2ClientJsonWebKeyRsaRequest{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyRsaRequest)
	if err == nil {
		o.OAuth2ClientJsonSigningKeyRequest = varOAuth2ClientJsonWebKeyRsaRequest.OAuth2ClientJsonSigningKeyRequest
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")

		// remove fields from embedded structs
		reflectOAuth2ClientJsonSigningKeyRequest := reflect.ValueOf(o.OAuth2ClientJsonSigningKeyRequest)
		for i := 0; i < reflectOAuth2ClientJsonSigningKeyRequest.Type().NumField(); i++ {
			t := reflectOAuth2ClientJsonSigningKeyRequest.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientJsonWebKeyRsaRequest struct {
	value *OAuth2ClientJsonWebKeyRsaRequest
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyRsaRequest) Get() *OAuth2ClientJsonWebKeyRsaRequest {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyRsaRequest) Set(val *OAuth2ClientJsonWebKeyRsaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyRsaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyRsaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyRsaRequest(val *OAuth2ClientJsonWebKeyRsaRequest) *NullableOAuth2ClientJsonWebKeyRsaRequest {
	return &NullableOAuth2ClientJsonWebKeyRsaRequest{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyRsaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyRsaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
