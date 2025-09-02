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

// Webauthn1 Verifies a `webauthn` factor challenge by posting a signed assertion using the challenge `nonce`
type Webauthn1 struct {
	// Base64-encoded authenticator data from the WebAuthn authenticator
	AuthenticatorData *string `json:"authenticatorData,omitempty"`
	// Base64-encoded client data from the WebAuthn authenticator
	ClientData *string `json:"clientData,omitempty"`
	// Base64-encoded signature data from the WebAuthn authenticator
	SignatureData *string `json:"signatureData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Webauthn1 Webauthn1

// NewWebauthn1 instantiates a new Webauthn1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebauthn1() *Webauthn1 {
	this := Webauthn1{}
	return &this
}

// NewWebauthn1WithDefaults instantiates a new Webauthn1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebauthn1WithDefaults() *Webauthn1 {
	this := Webauthn1{}
	return &this
}

// GetAuthenticatorData returns the AuthenticatorData field value if set, zero value otherwise.
func (o *Webauthn1) GetAuthenticatorData() string {
	if o == nil || o.AuthenticatorData == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorData
}

// GetAuthenticatorDataOk returns a tuple with the AuthenticatorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webauthn1) GetAuthenticatorDataOk() (*string, bool) {
	if o == nil || o.AuthenticatorData == nil {
		return nil, false
	}
	return o.AuthenticatorData, true
}

// HasAuthenticatorData returns a boolean if a field has been set.
func (o *Webauthn1) HasAuthenticatorData() bool {
	if o != nil && o.AuthenticatorData != nil {
		return true
	}

	return false
}

// SetAuthenticatorData gets a reference to the given string and assigns it to the AuthenticatorData field.
func (o *Webauthn1) SetAuthenticatorData(v string) {
	o.AuthenticatorData = &v
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *Webauthn1) GetClientData() string {
	if o == nil || o.ClientData == nil {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webauthn1) GetClientDataOk() (*string, bool) {
	if o == nil || o.ClientData == nil {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *Webauthn1) HasClientData() bool {
	if o != nil && o.ClientData != nil {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *Webauthn1) SetClientData(v string) {
	o.ClientData = &v
}

// GetSignatureData returns the SignatureData field value if set, zero value otherwise.
func (o *Webauthn1) GetSignatureData() string {
	if o == nil || o.SignatureData == nil {
		var ret string
		return ret
	}
	return *o.SignatureData
}

// GetSignatureDataOk returns a tuple with the SignatureData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webauthn1) GetSignatureDataOk() (*string, bool) {
	if o == nil || o.SignatureData == nil {
		return nil, false
	}
	return o.SignatureData, true
}

// HasSignatureData returns a boolean if a field has been set.
func (o *Webauthn1) HasSignatureData() bool {
	if o != nil && o.SignatureData != nil {
		return true
	}

	return false
}

// SetSignatureData gets a reference to the given string and assigns it to the SignatureData field.
func (o *Webauthn1) SetSignatureData(v string) {
	o.SignatureData = &v
}

func (o Webauthn1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatorData != nil {
		toSerialize["authenticatorData"] = o.AuthenticatorData
	}
	if o.ClientData != nil {
		toSerialize["clientData"] = o.ClientData
	}
	if o.SignatureData != nil {
		toSerialize["signatureData"] = o.SignatureData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Webauthn1) UnmarshalJSON(bytes []byte) (err error) {
	varWebauthn1 := _Webauthn1{}

	err = json.Unmarshal(bytes, &varWebauthn1)
	if err == nil {
		*o = Webauthn1(varWebauthn1)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticatorData")
		delete(additionalProperties, "clientData")
		delete(additionalProperties, "signatureData")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebauthn1 struct {
	value *Webauthn1
	isSet bool
}

func (v NullableWebauthn1) Get() *Webauthn1 {
	return v.value
}

func (v *NullableWebauthn1) Set(val *Webauthn1) {
	v.value = val
	v.isSet = true
}

func (v NullableWebauthn1) IsSet() bool {
	return v.isSet
}

func (v *NullableWebauthn1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebauthn1(val *Webauthn1) *NullableWebauthn1 {
	return &NullableWebauthn1{value: val, isSet: true}
}

func (v NullableWebauthn1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebauthn1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

