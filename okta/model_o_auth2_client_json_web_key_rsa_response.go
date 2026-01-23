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

// checks if the OAuth2ClientJsonWebKeyRsaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeyRsaResponse{}

// OAuth2ClientJsonWebKeyRsaResponse An RSA signing key
type OAuth2ClientJsonWebKeyRsaResponse struct {
	OAuth2ClientJsonSigningKeyResponse
	// RSA key value (exponent) for key binding
	E string `json:"e"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N                    string `json:"n"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyRsaResponse OAuth2ClientJsonWebKeyRsaResponse

// NewOAuth2ClientJsonWebKeyRsaResponse instantiates a new OAuth2ClientJsonWebKeyRsaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyRsaResponse(e string, n string, alg string, created string, id string, lastUpdated string, use string) *OAuth2ClientJsonWebKeyRsaResponse {
	this := OAuth2ClientJsonWebKeyRsaResponse{}
	this.Alg = alg
	this.Created = created
	this.Id = id
	this.LastUpdated = lastUpdated
	this.Use = use
	this.E = e
	this.N = n
	return &this
}

// NewOAuth2ClientJsonWebKeyRsaResponseWithDefaults instantiates a new OAuth2ClientJsonWebKeyRsaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyRsaResponseWithDefaults() *OAuth2ClientJsonWebKeyRsaResponse {
	this := OAuth2ClientJsonWebKeyRsaResponse{}
	return &this
}

// GetE returns the E field value
func (o *OAuth2ClientJsonWebKeyRsaResponse) GetE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRsaResponse) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *OAuth2ClientJsonWebKeyRsaResponse) SetE(v string) {
	o.E = v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyRsaResponse) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRsaResponse) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyRsaResponse) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyRsaResponse) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value
func (o *OAuth2ClientJsonWebKeyRsaResponse) GetN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyRsaResponse) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N, true
}

// SetN sets field value
func (o *OAuth2ClientJsonWebKeyRsaResponse) SetN(v string) {
	o.N = v
}

func (o OAuth2ClientJsonWebKeyRsaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeyRsaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOAuth2ClientJsonSigningKeyResponse, errOAuth2ClientJsonSigningKeyResponse := json.Marshal(o.OAuth2ClientJsonSigningKeyResponse)
	if errOAuth2ClientJsonSigningKeyResponse != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyResponse
	}
	errOAuth2ClientJsonSigningKeyResponse = json.Unmarshal([]byte(serializedOAuth2ClientJsonSigningKeyResponse), &toSerialize)
	if errOAuth2ClientJsonSigningKeyResponse != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyResponse
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

func (o *OAuth2ClientJsonWebKeyRsaResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"e",
		"n",
		"alg",
		"created",
		"id",
		"lastUpdated",
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

	type OAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct struct {
		// RSA key value (exponent) for key binding
		E string `json:"e"`
		// Cryptographic algorithm family for the certificate's key pair
		Kty *string `json:"kty,omitempty"`
		// RSA key value (modulus) for key binding
		N string `json:"n"`
	}

	varOAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct := OAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct)
	if err == nil {
		varOAuth2ClientJsonWebKeyRsaResponse := _OAuth2ClientJsonWebKeyRsaResponse{}
		varOAuth2ClientJsonWebKeyRsaResponse.E = varOAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct.E
		varOAuth2ClientJsonWebKeyRsaResponse.Kty = varOAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct.Kty
		varOAuth2ClientJsonWebKeyRsaResponse.N = varOAuth2ClientJsonWebKeyRsaResponseWithoutEmbeddedStruct.N
		*o = OAuth2ClientJsonWebKeyRsaResponse(varOAuth2ClientJsonWebKeyRsaResponse)
	} else {
		return err
	}

	varOAuth2ClientJsonWebKeyRsaResponse := _OAuth2ClientJsonWebKeyRsaResponse{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyRsaResponse)
	if err == nil {
		o.OAuth2ClientJsonSigningKeyResponse = varOAuth2ClientJsonWebKeyRsaResponse.OAuth2ClientJsonSigningKeyResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")

		// remove fields from embedded structs
		reflectOAuth2ClientJsonSigningKeyResponse := reflect.ValueOf(o.OAuth2ClientJsonSigningKeyResponse)
		for i := 0; i < reflectOAuth2ClientJsonSigningKeyResponse.Type().NumField(); i++ {
			t := reflectOAuth2ClientJsonSigningKeyResponse.Type().Field(i)

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

type NullableOAuth2ClientJsonWebKeyRsaResponse struct {
	value *OAuth2ClientJsonWebKeyRsaResponse
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyRsaResponse) Get() *OAuth2ClientJsonWebKeyRsaResponse {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyRsaResponse) Set(val *OAuth2ClientJsonWebKeyRsaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyRsaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyRsaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyRsaResponse(val *OAuth2ClientJsonWebKeyRsaResponse) *NullableOAuth2ClientJsonWebKeyRsaResponse {
	return &NullableOAuth2ClientJsonWebKeyRsaResponse{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyRsaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyRsaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
