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

// WebAuthnCredRequest Credential request object for the initialized credential, along with the enrollment and key identifiers to associate with the credential
type WebAuthnCredRequest struct {
	// ID for a WebAuthn Preregistration Factor in Okta
	AuthenticatorEnrollmentId *string `json:"authenticatorEnrollmentId,omitempty"`
	// Encrypted JWE of credential request for the fulfillment provider
	CredRequestJwe *string `json:"credRequestJwe,omitempty"`
	// ID for the Okta response key-pair used to encrypt and decrypt credential requests and responses
	KeyId *string `json:"keyId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebAuthnCredRequest WebAuthnCredRequest

// NewWebAuthnCredRequest instantiates a new WebAuthnCredRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnCredRequest() *WebAuthnCredRequest {
	this := WebAuthnCredRequest{}
	return &this
}

// NewWebAuthnCredRequestWithDefaults instantiates a new WebAuthnCredRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnCredRequestWithDefaults() *WebAuthnCredRequest {
	this := WebAuthnCredRequest{}
	return &this
}

// GetAuthenticatorEnrollmentId returns the AuthenticatorEnrollmentId field value if set, zero value otherwise.
func (o *WebAuthnCredRequest) GetAuthenticatorEnrollmentId() string {
	if o == nil || o.AuthenticatorEnrollmentId == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorEnrollmentId
}

// GetAuthenticatorEnrollmentIdOk returns a tuple with the AuthenticatorEnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnCredRequest) GetAuthenticatorEnrollmentIdOk() (*string, bool) {
	if o == nil || o.AuthenticatorEnrollmentId == nil {
		return nil, false
	}
	return o.AuthenticatorEnrollmentId, true
}

// HasAuthenticatorEnrollmentId returns a boolean if a field has been set.
func (o *WebAuthnCredRequest) HasAuthenticatorEnrollmentId() bool {
	if o != nil && o.AuthenticatorEnrollmentId != nil {
		return true
	}

	return false
}

// SetAuthenticatorEnrollmentId gets a reference to the given string and assigns it to the AuthenticatorEnrollmentId field.
func (o *WebAuthnCredRequest) SetAuthenticatorEnrollmentId(v string) {
	o.AuthenticatorEnrollmentId = &v
}

// GetCredRequestJwe returns the CredRequestJwe field value if set, zero value otherwise.
func (o *WebAuthnCredRequest) GetCredRequestJwe() string {
	if o == nil || o.CredRequestJwe == nil {
		var ret string
		return ret
	}
	return *o.CredRequestJwe
}

// GetCredRequestJweOk returns a tuple with the CredRequestJwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnCredRequest) GetCredRequestJweOk() (*string, bool) {
	if o == nil || o.CredRequestJwe == nil {
		return nil, false
	}
	return o.CredRequestJwe, true
}

// HasCredRequestJwe returns a boolean if a field has been set.
func (o *WebAuthnCredRequest) HasCredRequestJwe() bool {
	if o != nil && o.CredRequestJwe != nil {
		return true
	}

	return false
}

// SetCredRequestJwe gets a reference to the given string and assigns it to the CredRequestJwe field.
func (o *WebAuthnCredRequest) SetCredRequestJwe(v string) {
	o.CredRequestJwe = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *WebAuthnCredRequest) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnCredRequest) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *WebAuthnCredRequest) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *WebAuthnCredRequest) SetKeyId(v string) {
	o.KeyId = &v
}

func (o WebAuthnCredRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatorEnrollmentId != nil {
		toSerialize["authenticatorEnrollmentId"] = o.AuthenticatorEnrollmentId
	}
	if o.CredRequestJwe != nil {
		toSerialize["credRequestJwe"] = o.CredRequestJwe
	}
	if o.KeyId != nil {
		toSerialize["keyId"] = o.KeyId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebAuthnCredRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWebAuthnCredRequest := _WebAuthnCredRequest{}

	err = json.Unmarshal(bytes, &varWebAuthnCredRequest)
	if err == nil {
		*o = WebAuthnCredRequest(varWebAuthnCredRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticatorEnrollmentId")
		delete(additionalProperties, "credRequestJwe")
		delete(additionalProperties, "keyId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebAuthnCredRequest struct {
	value *WebAuthnCredRequest
	isSet bool
}

func (v NullableWebAuthnCredRequest) Get() *WebAuthnCredRequest {
	return v.value
}

func (v *NullableWebAuthnCredRequest) Set(val *WebAuthnCredRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnCredRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnCredRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnCredRequest(val *WebAuthnCredRequest) *NullableWebAuthnCredRequest {
	return &NullableWebAuthnCredRequest{value: val, isSet: true}
}

func (v NullableWebAuthnCredRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnCredRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

