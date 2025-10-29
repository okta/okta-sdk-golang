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
)

// checks if the UserFactorWebAuthnProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorWebAuthnProfile{}

// UserFactorWebAuthnProfile struct for UserFactorWebAuthnProfile
type UserFactorWebAuthnProfile struct {
	// Human-readable name of the authenticator
	AuthenticatorName *string `json:"authenticatorName,omitempty"`
	// ID for the factor credential
	CredentialId         *string `json:"credentialId,omitempty"`
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
	if o == nil || IsNil(o.AuthenticatorName) {
		var ret string
		return ret
	}
	return *o.AuthenticatorName
}

// GetAuthenticatorNameOk returns a tuple with the AuthenticatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorWebAuthnProfile) GetAuthenticatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorName) {
		return nil, false
	}
	return o.AuthenticatorName, true
}

// HasAuthenticatorName returns a boolean if a field has been set.
func (o *UserFactorWebAuthnProfile) HasAuthenticatorName() bool {
	if o != nil && !IsNil(o.AuthenticatorName) {
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
	if o == nil || IsNil(o.CredentialId) {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorWebAuthnProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialId) {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *UserFactorWebAuthnProfile) HasCredentialId() bool {
	if o != nil && !IsNil(o.CredentialId) {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *UserFactorWebAuthnProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o UserFactorWebAuthnProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorWebAuthnProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticatorName) {
		toSerialize["authenticatorName"] = o.AuthenticatorName
	}
	if !IsNil(o.CredentialId) {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorWebAuthnProfile) UnmarshalJSON(data []byte) (err error) {
	varUserFactorWebAuthnProfile := _UserFactorWebAuthnProfile{}

	err = json.Unmarshal(data, &varUserFactorWebAuthnProfile)

	if err != nil {
		return err
	}

	*o = UserFactorWebAuthnProfile(varUserFactorWebAuthnProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatorName")
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
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
