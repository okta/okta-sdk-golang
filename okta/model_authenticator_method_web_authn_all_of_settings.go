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

// AuthenticatorMethodWebAuthnAllOfSettings struct for AuthenticatorMethodWebAuthnAllOfSettings
type AuthenticatorMethodWebAuthnAllOfSettings struct {
	// <x-lifecycle class=\"ea\"></x-lifecycle> The FIDO2 AAGUID groups available to the WebAuthn authenticator
	AaguidGroups []AAGUIDGroupObject `json:"aaguidGroups,omitempty"`
	// User verification setting. Possible values `DISCOURAGED` (the authenticator isn't asked to perform user verification, but may do so at its discretion), `PREFERRED` (the client uses an authenticator capable of user verification if possible), or `REQUIRED`(the client uses only an authenticator capable of user verification)
	UserVerification *string `json:"userVerification,omitempty"`
	// Method attachment
	Attachment *string `json:"attachment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodWebAuthnAllOfSettings AuthenticatorMethodWebAuthnAllOfSettings

// NewAuthenticatorMethodWebAuthnAllOfSettings instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodWebAuthnAllOfSettings() *AuthenticatorMethodWebAuthnAllOfSettings {
	this := AuthenticatorMethodWebAuthnAllOfSettings{}
	return &this
}

// NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults() *AuthenticatorMethodWebAuthnAllOfSettings {
	this := AuthenticatorMethodWebAuthnAllOfSettings{}
	return &this
}

// GetAaguidGroups returns the AaguidGroups field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroups() []AAGUIDGroupObject {
	if o == nil || o.AaguidGroups == nil {
		var ret []AAGUIDGroupObject
		return ret
	}
	return o.AaguidGroups
}

// GetAaguidGroupsOk returns a tuple with the AaguidGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroupsOk() ([]AAGUIDGroupObject, bool) {
	if o == nil || o.AaguidGroups == nil {
		return nil, false
	}
	return o.AaguidGroups, true
}

// HasAaguidGroups returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAaguidGroups() bool {
	if o != nil && o.AaguidGroups != nil {
		return true
	}

	return false
}

// SetAaguidGroups gets a reference to the given []AAGUIDGroupObject and assigns it to the AaguidGroups field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAaguidGroups(v []AAGUIDGroupObject) {
	o.AaguidGroups = v
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerification() string {
	if o == nil || o.UserVerification == nil {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationOk() (*string, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetUserVerification(v string) {
	o.UserVerification = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachment() string {
	if o == nil || o.Attachment == nil {
		var ret string
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachmentOk() (*string, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAttachment() bool {
	if o != nil && o.Attachment != nil {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given string and assigns it to the Attachment field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAttachment(v string) {
	o.Attachment = &v
}

func (o AuthenticatorMethodWebAuthnAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AaguidGroups != nil {
		toSerialize["aaguidGroups"] = o.AaguidGroups
	}
	if o.UserVerification != nil {
		toSerialize["userVerification"] = o.UserVerification
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodWebAuthnAllOfSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorMethodWebAuthnAllOfSettings := _AuthenticatorMethodWebAuthnAllOfSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodWebAuthnAllOfSettings)
	if err == nil {
		*o = AuthenticatorMethodWebAuthnAllOfSettings(varAuthenticatorMethodWebAuthnAllOfSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "aaguidGroups")
		delete(additionalProperties, "userVerification")
		delete(additionalProperties, "attachment")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorMethodWebAuthnAllOfSettings struct {
	value *AuthenticatorMethodWebAuthnAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorMethodWebAuthnAllOfSettings) Get() *AuthenticatorMethodWebAuthnAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorMethodWebAuthnAllOfSettings) Set(val *AuthenticatorMethodWebAuthnAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodWebAuthnAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodWebAuthnAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodWebAuthnAllOfSettings(val *AuthenticatorMethodWebAuthnAllOfSettings) *NullableAuthenticatorMethodWebAuthnAllOfSettings {
	return &NullableAuthenticatorMethodWebAuthnAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodWebAuthnAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodWebAuthnAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

