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
)

// checks if the WebAuthnCredResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebAuthnCredResponse{}

// WebAuthnCredResponse Credential response object for enrolled credential details, along with enrollment and key identifiers to associate the credential
type WebAuthnCredResponse struct {
	// ID for a WebAuthn preregistration factor in Okta
	AuthenticatorEnrollmentId *string `json:"authenticatorEnrollmentId,omitempty"`
	// Encrypted JSON Web Encryption (JWE) of the credential response from the fulfillment provider
	CredResponseJwe      *string `json:"credResponseJwe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebAuthnCredResponse WebAuthnCredResponse

// NewWebAuthnCredResponse instantiates a new WebAuthnCredResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnCredResponse() *WebAuthnCredResponse {
	this := WebAuthnCredResponse{}
	return &this
}

// NewWebAuthnCredResponseWithDefaults instantiates a new WebAuthnCredResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnCredResponseWithDefaults() *WebAuthnCredResponse {
	this := WebAuthnCredResponse{}
	return &this
}

// GetAuthenticatorEnrollmentId returns the AuthenticatorEnrollmentId field value if set, zero value otherwise.
func (o *WebAuthnCredResponse) GetAuthenticatorEnrollmentId() string {
	if o == nil || IsNil(o.AuthenticatorEnrollmentId) {
		var ret string
		return ret
	}
	return *o.AuthenticatorEnrollmentId
}

// GetAuthenticatorEnrollmentIdOk returns a tuple with the AuthenticatorEnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnCredResponse) GetAuthenticatorEnrollmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorEnrollmentId) {
		return nil, false
	}
	return o.AuthenticatorEnrollmentId, true
}

// HasAuthenticatorEnrollmentId returns a boolean if a field has been set.
func (o *WebAuthnCredResponse) HasAuthenticatorEnrollmentId() bool {
	if o != nil && !IsNil(o.AuthenticatorEnrollmentId) {
		return true
	}

	return false
}

// SetAuthenticatorEnrollmentId gets a reference to the given string and assigns it to the AuthenticatorEnrollmentId field.
func (o *WebAuthnCredResponse) SetAuthenticatorEnrollmentId(v string) {
	o.AuthenticatorEnrollmentId = &v
}

// GetCredResponseJwe returns the CredResponseJwe field value if set, zero value otherwise.
func (o *WebAuthnCredResponse) GetCredResponseJwe() string {
	if o == nil || IsNil(o.CredResponseJwe) {
		var ret string
		return ret
	}
	return *o.CredResponseJwe
}

// GetCredResponseJweOk returns a tuple with the CredResponseJwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnCredResponse) GetCredResponseJweOk() (*string, bool) {
	if o == nil || IsNil(o.CredResponseJwe) {
		return nil, false
	}
	return o.CredResponseJwe, true
}

// HasCredResponseJwe returns a boolean if a field has been set.
func (o *WebAuthnCredResponse) HasCredResponseJwe() bool {
	if o != nil && !IsNil(o.CredResponseJwe) {
		return true
	}

	return false
}

// SetCredResponseJwe gets a reference to the given string and assigns it to the CredResponseJwe field.
func (o *WebAuthnCredResponse) SetCredResponseJwe(v string) {
	o.CredResponseJwe = &v
}

func (o WebAuthnCredResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebAuthnCredResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticatorEnrollmentId) {
		toSerialize["authenticatorEnrollmentId"] = o.AuthenticatorEnrollmentId
	}
	if !IsNil(o.CredResponseJwe) {
		toSerialize["credResponseJwe"] = o.CredResponseJwe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebAuthnCredResponse) UnmarshalJSON(data []byte) (err error) {
	varWebAuthnCredResponse := _WebAuthnCredResponse{}

	err = json.Unmarshal(data, &varWebAuthnCredResponse)

	if err != nil {
		return err
	}

	*o = WebAuthnCredResponse(varWebAuthnCredResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatorEnrollmentId")
		delete(additionalProperties, "credResponseJwe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebAuthnCredResponse struct {
	value *WebAuthnCredResponse
	isSet bool
}

func (v NullableWebAuthnCredResponse) Get() *WebAuthnCredResponse {
	return v.value
}

func (v *NullableWebAuthnCredResponse) Set(val *WebAuthnCredResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnCredResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnCredResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnCredResponse(val *WebAuthnCredResponse) *NullableWebAuthnCredResponse {
	return &NullableWebAuthnCredResponse{value: val, isSet: true}
}

func (v NullableWebAuthnCredResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnCredResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
