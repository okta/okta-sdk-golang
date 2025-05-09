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

// UserFactorWebAuthnProfile struct for UserFactorWebAuthnProfile
type UserFactorWebAuthnProfile struct {
	// Human-readable name of the authenticator
	AuthenticatorName *string `json:"authenticatorName,omitempty"`
	// ID for the Factor credential
	CredentialId *string `json:"credentialId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorWebAuthnProfile UserFactorWebAuthnProfile

// NewUserFactorWebAuthnProfile instantiates a new UserFactorWebAuthnProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorWebAuthnProfile() *UserFactorWebAuthnProfile {
	this := UserFactorWebAuthnProfile{}
	return &this
}

// NewUserFactorWebAuthnProfileWithDefaults instantiates a new UserFactorWebAuthnProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorWebAuthnProfileWithDefaults() *UserFactorWebAuthnProfile {
	this := UserFactorWebAuthnProfile{}
	return &this
}

// GetAuthenticatorName returns the AuthenticatorName field value if set, zero value otherwise.
func (o *UserFactorWebAuthnProfile) GetAuthenticatorName() string {
	if o == nil || o.AuthenticatorName == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorName
}

// GetAuthenticatorNameOk returns a tuple with the AuthenticatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorWebAuthnProfile) GetAuthenticatorNameOk() (*string, bool) {
	if o == nil || o.AuthenticatorName == nil {
		return nil, false
	}
	return o.AuthenticatorName, true
}

// HasAuthenticatorName returns a boolean if a field has been set.
func (o *UserFactorWebAuthnProfile) HasAuthenticatorName() bool {
	if o != nil && o.AuthenticatorName != nil {
		return true
	}

	return false
}

// SetAuthenticatorName gets a reference to the given string and assigns it to the AuthenticatorName field.
func (o *UserFactorWebAuthnProfile) SetAuthenticatorName(v string) {
	o.AuthenticatorName = &v
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *UserFactorWebAuthnProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorWebAuthnProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *UserFactorWebAuthnProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *UserFactorWebAuthnProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o UserFactorWebAuthnProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatorName != nil {
		toSerialize["authenticatorName"] = o.AuthenticatorName
	}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorWebAuthnProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorWebAuthnProfile := _UserFactorWebAuthnProfile{}

	err = json.Unmarshal(bytes, &varUserFactorWebAuthnProfile)
	if err == nil {
		*o = UserFactorWebAuthnProfile(varUserFactorWebAuthnProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticatorName")
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorWebAuthnProfile struct {
	value *UserFactorWebAuthnProfile
	isSet bool
}

func (v NullableUserFactorWebAuthnProfile) Get() *UserFactorWebAuthnProfile {
	return v.value
}

func (v *NullableUserFactorWebAuthnProfile) Set(val *UserFactorWebAuthnProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorWebAuthnProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorWebAuthnProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorWebAuthnProfile(val *UserFactorWebAuthnProfile) *NullableUserFactorWebAuthnProfile {
	return &NullableUserFactorWebAuthnProfile{value: val, isSet: true}
}

func (v NullableUserFactorWebAuthnProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorWebAuthnProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

