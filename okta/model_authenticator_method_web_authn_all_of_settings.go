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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AuthenticatorMethodWebAuthnAllOfSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorMethodWebAuthnAllOfSettings{}

// AuthenticatorMethodWebAuthnAllOfSettings The settings for the WebAuthn authenticator method
type AuthenticatorMethodWebAuthnAllOfSettings struct {
	// The FIDO2 Authenticator Attestation Global Unique Identifiers (AAGUID) groups available to the WebAuthn authenticator
	AaguidGroups []AAGUIDGroupObject `json:"aaguidGroups,omitempty"`
	// User verification setting. Possible values `DISCOURAGED` (the authenticator isn't asked to perform user verification, but may do so at its discretion), `PREFERRED` (the client uses an authenticator capable of user verification if possible), or `REQUIRED`(the client uses only an authenticator capable of user verification)
	UserVerification *string `json:"userVerification,omitempty"`
	// Method attachment
	Attachment *string `json:"attachment,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Enables the passkeys autofill UI to display available WebAuthn discoverable credentials (\"resident key\") from the Sign-In Widget username field
	EnableAutofillUI     *bool `json:"enableAutofillUI,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodWebAuthnAllOfSettings AuthenticatorMethodWebAuthnAllOfSettings

// NewAuthenticatorMethodWebAuthnAllOfSettings instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodWebAuthnAllOfSettings() *AuthenticatorMethodWebAuthnAllOfSettings {
	this := AuthenticatorMethodWebAuthnAllOfSettings{}
	var enableAutofillUI bool = false
	this.EnableAutofillUI = &enableAutofillUI
	return &this
}

// NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults() *AuthenticatorMethodWebAuthnAllOfSettings {
	this := AuthenticatorMethodWebAuthnAllOfSettings{}
	var enableAutofillUI bool = false
	this.EnableAutofillUI = &enableAutofillUI
	return &this
}

// GetAaguidGroups returns the AaguidGroups field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroups() []AAGUIDGroupObject {
	if o == nil || IsNil(o.AaguidGroups) {
		var ret []AAGUIDGroupObject
		return ret
	}
	return o.AaguidGroups
}

// GetAaguidGroupsOk returns a tuple with the AaguidGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroupsOk() ([]AAGUIDGroupObject, bool) {
	if o == nil || IsNil(o.AaguidGroups) {
		return nil, false
	}
	return o.AaguidGroups, true
}

// HasAaguidGroups returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAaguidGroups() bool {
	if o != nil && !IsNil(o.AaguidGroups) {
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
	if o == nil || IsNil(o.UserVerification) {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.UserVerification) {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerification() bool {
	if o != nil && !IsNil(o.UserVerification) {
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
	if o == nil || IsNil(o.Attachment) {
		var ret string
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given string and assigns it to the Attachment field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAttachment(v string) {
	o.Attachment = &v
}

// GetEnableAutofillUI returns the EnableAutofillUI field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetEnableAutofillUI() bool {
	if o == nil || IsNil(o.EnableAutofillUI) {
		var ret bool
		return ret
	}
	return *o.EnableAutofillUI
}

// GetEnableAutofillUIOk returns a tuple with the EnableAutofillUI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetEnableAutofillUIOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutofillUI) {
		return nil, false
	}
	return o.EnableAutofillUI, true
}

// HasEnableAutofillUI returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasEnableAutofillUI() bool {
	if o != nil && !IsNil(o.EnableAutofillUI) {
		return true
	}

	return false
}

// SetEnableAutofillUI gets a reference to the given bool and assigns it to the EnableAutofillUI field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetEnableAutofillUI(v bool) {
	o.EnableAutofillUI = &v
}

func (o AuthenticatorMethodWebAuthnAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorMethodWebAuthnAllOfSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AaguidGroups) {
		toSerialize["aaguidGroups"] = o.AaguidGroups
	}
	if !IsNil(o.UserVerification) {
		toSerialize["userVerification"] = o.UserVerification
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	if !IsNil(o.EnableAutofillUI) {
		toSerialize["enableAutofillUI"] = o.EnableAutofillUI
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorMethodWebAuthnAllOfSettings) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorMethodWebAuthnAllOfSettings := _AuthenticatorMethodWebAuthnAllOfSettings{}

	err = json.Unmarshal(data, &varAuthenticatorMethodWebAuthnAllOfSettings)

	if err != nil {
		return err
	}

	*o = AuthenticatorMethodWebAuthnAllOfSettings(varAuthenticatorMethodWebAuthnAllOfSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aaguidGroups")
		delete(additionalProperties, "userVerification")
		delete(additionalProperties, "attachment")
		delete(additionalProperties, "enableAutofillUI")
		o.AdditionalProperties = additionalProperties
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
