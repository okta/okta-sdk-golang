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

// checks if the OAuth2ClientJsonWebKeyECRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeyECRequest{}

// OAuth2ClientJsonWebKeyECRequest An EC signing key
type OAuth2ClientJsonWebKeyECRequest struct {
	OAuth2ClientJsonSigningKeyRequest
	// The cryptographic curve used with the key
	Crv string `json:"crv"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// The public x coordinate for the elliptic curve point
	X string `json:"x"`
	// The public y coordinate for the elliptic curve point
	Y                    string `json:"y"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyECRequest OAuth2ClientJsonWebKeyECRequest

// NewOAuth2ClientJsonWebKeyECRequest instantiates a new OAuth2ClientJsonWebKeyECRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyECRequest(crv string, x string, y string, alg string, use string) *OAuth2ClientJsonWebKeyECRequest {
	this := OAuth2ClientJsonWebKeyECRequest{}
	this.Alg = alg
	var status string = "ACTIVE"
	this.Status = &status
	this.Use = use
	this.Crv = crv
	this.X = x
	this.Y = y
	return &this
}

// NewOAuth2ClientJsonWebKeyECRequestWithDefaults instantiates a new OAuth2ClientJsonWebKeyECRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyECRequestWithDefaults() *OAuth2ClientJsonWebKeyECRequest {
	this := OAuth2ClientJsonWebKeyECRequest{}
	return &this
}

// GetCrv returns the Crv field value
func (o *OAuth2ClientJsonWebKeyECRequest) GetCrv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Crv
}

// GetCrvOk returns a tuple with the Crv field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetCrvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crv, true
}

// SetCrv sets field value
func (o *OAuth2ClientJsonWebKeyECRequest) SetCrv(v string) {
	o.Crv = v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECRequest) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyECRequest) SetKty(v string) {
	o.Kty = &v
}

// GetX returns the X field value
func (o *OAuth2ClientJsonWebKeyECRequest) GetX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *OAuth2ClientJsonWebKeyECRequest) SetX(v string) {
	o.X = v
}

// GetY returns the Y field value
func (o *OAuth2ClientJsonWebKeyECRequest) GetY() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetYOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *OAuth2ClientJsonWebKeyECRequest) SetY(v string) {
	o.Y = v
}

func (o OAuth2ClientJsonWebKeyECRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeyECRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOAuth2ClientJsonSigningKeyRequest, errOAuth2ClientJsonSigningKeyRequest := json.Marshal(o.OAuth2ClientJsonSigningKeyRequest)
	if errOAuth2ClientJsonSigningKeyRequest != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyRequest
	}
	errOAuth2ClientJsonSigningKeyRequest = json.Unmarshal([]byte(serializedOAuth2ClientJsonSigningKeyRequest), &toSerialize)
	if errOAuth2ClientJsonSigningKeyRequest != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyRequest
	}
	toSerialize["crv"] = o.Crv
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientJsonWebKeyECRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"crv",
		"x",
		"y",
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

	type OAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct struct {
		// The cryptographic curve used with the key
		Crv string `json:"crv"`
		// Cryptographic algorithm family for the certificate's key pair
		Kty *string `json:"kty,omitempty"`
		// The public x coordinate for the elliptic curve point
		X string `json:"x"`
		// The public y coordinate for the elliptic curve point
		Y string `json:"y"`
	}

	varOAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct := OAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct)
	if err == nil {
		varOAuth2ClientJsonWebKeyECRequest := _OAuth2ClientJsonWebKeyECRequest{}
		varOAuth2ClientJsonWebKeyECRequest.Crv = varOAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct.Crv
		varOAuth2ClientJsonWebKeyECRequest.Kty = varOAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct.Kty
		varOAuth2ClientJsonWebKeyECRequest.X = varOAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct.X
		varOAuth2ClientJsonWebKeyECRequest.Y = varOAuth2ClientJsonWebKeyECRequestWithoutEmbeddedStruct.Y
		*o = OAuth2ClientJsonWebKeyECRequest(varOAuth2ClientJsonWebKeyECRequest)
	} else {
		return err
	}

	varOAuth2ClientJsonWebKeyECRequest := _OAuth2ClientJsonWebKeyECRequest{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyECRequest)
	if err == nil {
		o.OAuth2ClientJsonSigningKeyRequest = varOAuth2ClientJsonWebKeyECRequest.OAuth2ClientJsonSigningKeyRequest
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "crv")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")

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

type NullableOAuth2ClientJsonWebKeyECRequest struct {
	value *OAuth2ClientJsonWebKeyECRequest
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyECRequest) Get() *OAuth2ClientJsonWebKeyECRequest {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyECRequest) Set(val *OAuth2ClientJsonWebKeyECRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyECRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyECRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyECRequest(val *OAuth2ClientJsonWebKeyECRequest) *NullableOAuth2ClientJsonWebKeyECRequest {
	return &NullableOAuth2ClientJsonWebKeyECRequest{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyECRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyECRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
