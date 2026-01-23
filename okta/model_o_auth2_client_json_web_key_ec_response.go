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

// checks if the OAuth2ClientJsonWebKeyECResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientJsonWebKeyECResponse{}

// OAuth2ClientJsonWebKeyECResponse An EC signing key
type OAuth2ClientJsonWebKeyECResponse struct {
	OAuth2ClientJsonSigningKeyResponse
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

type _OAuth2ClientJsonWebKeyECResponse OAuth2ClientJsonWebKeyECResponse

// NewOAuth2ClientJsonWebKeyECResponse instantiates a new OAuth2ClientJsonWebKeyECResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyECResponse(crv string, x string, y string, alg string, created string, id string, lastUpdated string, use string) *OAuth2ClientJsonWebKeyECResponse {
	this := OAuth2ClientJsonWebKeyECResponse{}
	this.Alg = alg
	this.Created = created
	this.Id = id
	this.X = x
	this.Y = y
	this.LastUpdated = lastUpdated
	this.Use = use
	this.Crv = crv
	return &this
}

// NewOAuth2ClientJsonWebKeyECResponseWithDefaults instantiates a new OAuth2ClientJsonWebKeyECResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyECResponseWithDefaults() *OAuth2ClientJsonWebKeyECResponse {
	this := OAuth2ClientJsonWebKeyECResponse{}
	return &this
}

// GetCrv returns the Crv field value
func (o *OAuth2ClientJsonWebKeyECResponse) GetCrv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Crv
}

// GetCrvOk returns a tuple with the Crv field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetCrvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crv, true
}

// SetCrv sets field value
func (o *OAuth2ClientJsonWebKeyECResponse) SetCrv(v string) {
	o.Crv = v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECResponse) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyECResponse) SetKty(v string) {
	o.Kty = &v
}

// GetX returns the X field value
func (o *OAuth2ClientJsonWebKeyECResponse) GetX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *OAuth2ClientJsonWebKeyECResponse) SetX(v string) {
	o.X = v
}

// GetY returns the Y field value
func (o *OAuth2ClientJsonWebKeyECResponse) GetY() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECResponse) GetYOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *OAuth2ClientJsonWebKeyECResponse) SetY(v string) {
	o.Y = v
}

func (o OAuth2ClientJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientJsonWebKeyECResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOAuth2ClientJsonSigningKeyResponse, errOAuth2ClientJsonSigningKeyResponse := json.Marshal(o.OAuth2ClientJsonSigningKeyResponse)
	if errOAuth2ClientJsonSigningKeyResponse != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyResponse
	}
	errOAuth2ClientJsonSigningKeyResponse = json.Unmarshal([]byte(serializedOAuth2ClientJsonSigningKeyResponse), &toSerialize)
	if errOAuth2ClientJsonSigningKeyResponse != nil {
		return map[string]interface{}{}, errOAuth2ClientJsonSigningKeyResponse
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

func (o *OAuth2ClientJsonWebKeyECResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"crv",
		"x",
		"y",
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

	type OAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct struct {
		// The cryptographic curve used with the key
		Crv string `json:"crv"`
		// Cryptographic algorithm family for the certificate's key pair
		Kty *string `json:"kty,omitempty"`
		// The public x coordinate for the elliptic curve point
		X string `json:"x"`
		// The public y coordinate for the elliptic curve point
		Y string `json:"y"`
	}

	varOAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct := OAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct)
	if err == nil {
		varOAuth2ClientJsonWebKeyECResponse := _OAuth2ClientJsonWebKeyECResponse{}
		varOAuth2ClientJsonWebKeyECResponse.Crv = varOAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct.Crv
		varOAuth2ClientJsonWebKeyECResponse.Kty = varOAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct.Kty
		varOAuth2ClientJsonWebKeyECResponse.X = varOAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct.X
		varOAuth2ClientJsonWebKeyECResponse.Y = varOAuth2ClientJsonWebKeyECResponseWithoutEmbeddedStruct.Y
		*o = OAuth2ClientJsonWebKeyECResponse(varOAuth2ClientJsonWebKeyECResponse)
	} else {
		return err
	}

	varOAuth2ClientJsonWebKeyECResponse := _OAuth2ClientJsonWebKeyECResponse{}

	err = json.Unmarshal(data, &varOAuth2ClientJsonWebKeyECResponse)
	if err == nil {
		o.OAuth2ClientJsonSigningKeyResponse = varOAuth2ClientJsonWebKeyECResponse.OAuth2ClientJsonSigningKeyResponse
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

type NullableOAuth2ClientJsonWebKeyECResponse struct {
	value *OAuth2ClientJsonWebKeyECResponse
	isSet bool
}

func (v NullableOAuth2ClientJsonWebKeyECResponse) Get() *OAuth2ClientJsonWebKeyECResponse {
	return v.value
}

func (v *NullableOAuth2ClientJsonWebKeyECResponse) Set(val *OAuth2ClientJsonWebKeyECResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonWebKeyECResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonWebKeyECResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonWebKeyECResponse(val *OAuth2ClientJsonWebKeyECResponse) *NullableOAuth2ClientJsonWebKeyECResponse {
	return &NullableOAuth2ClientJsonWebKeyECResponse{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonWebKeyECResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
