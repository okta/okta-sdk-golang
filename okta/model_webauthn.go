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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// Webauthn Activates a `webauthn` Factor with the specified attestation and registration information from the WebAuthn authenticator
type Webauthn struct {
	// Base64-encoded attestation from the WebAuthn authenticator
	Attestation *string `json:"attestation,omitempty"`
	// Base64-encoded client data from the WebAuthn authenticator
	ClientData           *string `json:"clientData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Webauthn Webauthn

// NewWebauthn instantiates a new Webauthn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebauthn() *Webauthn {
	this := Webauthn{}
	return &this
}

// NewWebauthnWithDefaults instantiates a new Webauthn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebauthnWithDefaults() *Webauthn {
	this := Webauthn{}
	return &this
}

// GetAttestation returns the Attestation field value if set, zero value otherwise.
func (o *Webauthn) GetAttestation() string {
	if o == nil || o.Attestation == nil {
		var ret string
		return ret
	}
	return *o.Attestation
}

// GetAttestationOk returns a tuple with the Attestation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webauthn) GetAttestationOk() (*string, bool) {
	if o == nil || o.Attestation == nil {
		return nil, false
	}
	return o.Attestation, true
}

// HasAttestation returns a boolean if a field has been set.
func (o *Webauthn) HasAttestation() bool {
	if o != nil && o.Attestation != nil {
		return true
	}

	return false
}

// SetAttestation gets a reference to the given string and assigns it to the Attestation field.
func (o *Webauthn) SetAttestation(v string) {
	o.Attestation = &v
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *Webauthn) GetClientData() string {
	if o == nil || o.ClientData == nil {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webauthn) GetClientDataOk() (*string, bool) {
	if o == nil || o.ClientData == nil {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *Webauthn) HasClientData() bool {
	if o != nil && o.ClientData != nil {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *Webauthn) SetClientData(v string) {
	o.ClientData = &v
}

func (o Webauthn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attestation != nil {
		toSerialize["attestation"] = o.Attestation
	}
	if o.ClientData != nil {
		toSerialize["clientData"] = o.ClientData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Webauthn) UnmarshalJSON(bytes []byte) (err error) {
	varWebauthn := _Webauthn{}

	err = json.Unmarshal(bytes, &varWebauthn)
	if err == nil {
		*o = Webauthn(varWebauthn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "attestation")
		delete(additionalProperties, "clientData")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebauthn struct {
	value *Webauthn
	isSet bool
}

func (v NullableWebauthn) Get() *Webauthn {
	return v.value
}

func (v *NullableWebauthn) Set(val *Webauthn) {
	v.value = val
	v.isSet = true
}

func (v NullableWebauthn) IsSet() bool {
	return v.isSet
}

func (v *NullableWebauthn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebauthn(val *Webauthn) *NullableWebauthn {
	return &NullableWebauthn{value: val, isSet: true}
}

func (v NullableWebauthn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebauthn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
