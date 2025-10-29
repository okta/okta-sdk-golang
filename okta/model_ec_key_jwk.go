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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the ECKeyJWK type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECKeyJWK{}

// ECKeyJWK Elliptic curve key in JSON Web Key (JWK) format. It's used during enrollment to encrypt fulfillment requests to Yubico, or during activation to verify Yubico's JWS (JSON Web Signature) objects in fulfillment responses. The currently agreed protocol uses P-384.
type ECKeyJWK struct {
	// The elliptic curve protocol
	Crv string `json:"crv"`
	// The unique identifier of the key
	Kid string `json:"kid"`
	// The type of public key
	Kty string `json:"kty"`
	// The intended use for the key. This value is either `enc` (encryption) during enrollment, when Okta uses the ECKeyJWK to encrypt requests to Yubico. Or it's `sig` (signature) during activation, when Okta uses the ECKeyJWK to verify the responses from Yubico.
	Use string `json:"use"`
	// The public x coordinate for the elliptic curve point
	X string `json:"x"`
	// The public y coordinate for the elliptic curve point
	Y                    string `json:"y"`
	AdditionalProperties map[string]interface{}
}

type _ECKeyJWK ECKeyJWK

// NewECKeyJWK instantiates a new ECKeyJWK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECKeyJWK(crv string, kid string, kty string, use string, x string, y string) *ECKeyJWK {
	this := ECKeyJWK{}
	this.Crv = crv
	this.Kid = kid
	this.Kty = kty
	this.Use = use
	this.X = x
	this.Y = y
	return &this
}

// NewECKeyJWKWithDefaults instantiates a new ECKeyJWK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECKeyJWKWithDefaults() *ECKeyJWK {
	this := ECKeyJWK{}
	return &this
}

// GetCrv returns the Crv field value
func (o *ECKeyJWK) GetCrv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Crv
}

// GetCrvOk returns a tuple with the Crv field value
// and a boolean to check if the value has been set.
func (o *ECKeyJWK) GetCrvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crv, true
}

// SetCrv sets field value
func (o *ECKeyJWK) SetCrv(v string) {
	o.Crv = v
}

// GetKid returns the Kid field value
func (o *ECKeyJWK) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *ECKeyJWK) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *ECKeyJWK) SetKid(v string) {
	o.Kid = v
}

// GetKty returns the Kty field value
func (o *ECKeyJWK) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *ECKeyJWK) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *ECKeyJWK) SetKty(v string) {
	o.Kty = v
}

// GetUse returns the Use field value
func (o *ECKeyJWK) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *ECKeyJWK) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *ECKeyJWK) SetUse(v string) {
	o.Use = v
}

// GetX returns the X field value
func (o *ECKeyJWK) GetX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ECKeyJWK) GetXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ECKeyJWK) SetX(v string) {
	o.X = v
}

// GetY returns the Y field value
func (o *ECKeyJWK) GetY() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ECKeyJWK) GetYOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ECKeyJWK) SetY(v string) {
	o.Y = v
}

func (o ECKeyJWK) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECKeyJWK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["crv"] = o.Crv
	toSerialize["kid"] = o.Kid
	toSerialize["kty"] = o.Kty
	toSerialize["use"] = o.Use
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ECKeyJWK) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"crv",
		"kid",
		"kty",
		"use",
		"x",
		"y",
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

	varECKeyJWK := _ECKeyJWK{}

	err = json.Unmarshal(data, &varECKeyJWK)

	if err != nil {
		return err
	}

	*o = ECKeyJWK(varECKeyJWK)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "crv")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableECKeyJWK struct {
	value *ECKeyJWK
	isSet bool
}

func (v NullableECKeyJWK) Get() *ECKeyJWK {
	return v.value
}

func (v *NullableECKeyJWK) Set(val *ECKeyJWK) {
	v.value = val
	v.isSet = true
}

func (v NullableECKeyJWK) IsSet() bool {
	return v.isSet
}

func (v *NullableECKeyJWK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECKeyJWK(val *ECKeyJWK) *NullableECKeyJWK {
	return &NullableECKeyJWK{value: val, isSet: true}
}

func (v NullableECKeyJWK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECKeyJWK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
