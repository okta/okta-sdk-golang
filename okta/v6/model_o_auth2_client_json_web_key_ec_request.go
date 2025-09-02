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

// OAuth2ClientJsonWebKeyECRequest An EC signing key
type OAuth2ClientJsonWebKeyECRequest struct {
	// Cryptographic algorithm family for the certificate's key pair
	Kty *string `json:"kty,omitempty"`
	// The public x coordinate for the elliptic curve point
	X *string `json:"x,omitempty"`
	// The public y coordinate for the elliptic curve point
	Y *string `json:"y,omitempty"`
	// Unique identifier of the JSON Web Key in the OAUth 2.0 client's JWKS
	Kid NullableString `json:"kid,omitempty"`
	// Status of the OAuth 2.0 client JSON Web Key
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientJsonWebKeyECRequest OAuth2ClientJsonWebKeyECRequest

// NewOAuth2ClientJsonWebKeyECRequest instantiates a new OAuth2ClientJsonWebKeyECRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientJsonWebKeyECRequest() *OAuth2ClientJsonWebKeyECRequest {
	this := OAuth2ClientJsonWebKeyECRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewOAuth2ClientJsonWebKeyECRequestWithDefaults instantiates a new OAuth2ClientJsonWebKeyECRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientJsonWebKeyECRequestWithDefaults() *OAuth2ClientJsonWebKeyECRequest {
	this := OAuth2ClientJsonWebKeyECRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECRequest) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *OAuth2ClientJsonWebKeyECRequest) SetKty(v string) {
	o.Kty = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECRequest) GetX() string {
	if o == nil || o.X == nil {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetXOk() (*string, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *OAuth2ClientJsonWebKeyECRequest) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECRequest) GetY() string {
	if o == nil || o.Y == nil {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetYOk() (*string, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *OAuth2ClientJsonWebKeyECRequest) SetY(v string) {
	o.Y = &v
}

// GetKid returns the Kid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ClientJsonWebKeyECRequest) GetKid() string {
	if o == nil || o.Kid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Kid.Get()
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ClientJsonWebKeyECRequest) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kid.Get(), o.Kid.IsSet()
}

// HasKid returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) HasKid() bool {
	if o != nil && o.Kid.IsSet() {
		return true
	}

	return false
}

// SetKid gets a reference to the given NullableString and assigns it to the Kid field.
func (o *OAuth2ClientJsonWebKeyECRequest) SetKid(v string) {
	o.Kid.Set(&v)
}
// SetKidNil sets the value for Kid to be an explicit nil
func (o *OAuth2ClientJsonWebKeyECRequest) SetKidNil() {
	o.Kid.Set(nil)
}

// UnsetKid ensures that no value is present for Kid, not even an explicit nil
func (o *OAuth2ClientJsonWebKeyECRequest) UnsetKid() {
	o.Kid.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2ClientJsonWebKeyECRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2ClientJsonWebKeyECRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2ClientJsonWebKeyECRequest) SetStatus(v string) {
	o.Status = &v
}

func (o OAuth2ClientJsonWebKeyECRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
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

func (o *OAuth2ClientJsonWebKeyECRequest) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ClientJsonWebKeyECRequest := _OAuth2ClientJsonWebKeyECRequest{}

	err = json.Unmarshal(bytes, &varOAuth2ClientJsonWebKeyECRequest)
	if err == nil {
		*o = OAuth2ClientJsonWebKeyECRequest(varOAuth2ClientJsonWebKeyECRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "kty")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

