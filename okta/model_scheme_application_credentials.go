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

// checks if the SchemeApplicationCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemeApplicationCredentials{}

// SchemeApplicationCredentials struct for SchemeApplicationCredentials
type SchemeApplicationCredentials struct {
	Signing          *ApplicationCredentialsSigning          `json:"signing,omitempty"`
	UserNameTemplate *ApplicationCredentialsUsernameTemplate `json:"userNameTemplate,omitempty"`
	Password         *PasswordCredential                     `json:"password,omitempty"`
	// Allow users to securely see their password
	RevealPassword *bool `json:"revealPassword,omitempty"`
	// Apps with `BASIC_AUTH`, `BROWSER_PLUGIN`, or `SECURE_PASSWORD_STORE` sign-on modes have credentials vaulted by Okta and can be configured with the following schemes.
	Scheme *string `json:"scheme,omitempty"`
	// Shared username for the app
	UserName             *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchemeApplicationCredentials SchemeApplicationCredentials

// NewSchemeApplicationCredentials instantiates a new SchemeApplicationCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemeApplicationCredentials() *SchemeApplicationCredentials {
	this := SchemeApplicationCredentials{}
	return &this
}

// NewSchemeApplicationCredentialsWithDefaults instantiates a new SchemeApplicationCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemeApplicationCredentialsWithDefaults() *SchemeApplicationCredentials {
	this := SchemeApplicationCredentials{}
	return &this
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *SchemeApplicationCredentials) GetSigning() ApplicationCredentialsSigning {
	if o == nil || IsNil(o.Signing) {
		var ret ApplicationCredentialsSigning
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentials) GetSigningOk() (*ApplicationCredentialsSigning, bool) {
	if o == nil || IsNil(o.Signing) {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *SchemeApplicationCredentials) HasSigning() bool {
	if o != nil && !IsNil(o.Signing) {
		return true
	}

	return false
}

// SetSigning gets a reference to the given ApplicationCredentialsSigning and assigns it to the Signing field.
func (o *SchemeApplicationCredentials) SetSigning(v ApplicationCredentialsSigning) {
	o.Signing = &v
}

// GetUserNameTemplate returns the UserNameTemplate field value if set, zero value otherwise.
func (o *SchemeApplicationCredentials) GetUserNameTemplate() ApplicationCredentialsUsernameTemplate {
	if o == nil || IsNil(o.UserNameTemplate) {
		var ret ApplicationCredentialsUsernameTemplate
		return ret
	}
	return *o.UserNameTemplate
}

// GetUserNameTemplateOk returns a tuple with the UserNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentials) GetUserNameTemplateOk() (*ApplicationCredentialsUsernameTemplate, bool) {
	if o == nil || IsNil(o.UserNameTemplate) {
		return nil, false
	}
	return o.UserNameTemplate, true
}

// HasUserNameTemplate returns a boolean if a field has been set.
func (o *SchemeApplicationCredentials) HasUserNameTemplate() bool {
	if o != nil && !IsNil(o.UserNameTemplate) {
		return true
	}

	return false
}

// SetUserNameTemplate gets a reference to the given ApplicationCredentialsUsernameTemplate and assigns it to the UserNameTemplate field.
func (o *SchemeApplicationCredentials) SetUserNameTemplate(v ApplicationCredentialsUsernameTemplate) {
	o.UserNameTemplate = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SchemeApplicationCredentials) GetPassword() PasswordCredential {
	if o == nil || IsNil(o.Password) {
		var ret PasswordCredential
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentials) GetPasswordOk() (*PasswordCredential, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SchemeApplicationCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given PasswordCredential and assigns it to the Password field.
func (o *SchemeApplicationCredentials) SetPassword(v PasswordCredential) {
	o.Password = &v
}

// GetRevealPassword returns the RevealPassword field value if set, zero value otherwise.
func (o *SchemeApplicationCredentials) GetRevealPassword() bool {
	if o == nil || IsNil(o.RevealPassword) {
		var ret bool
		return ret
	}
	return *o.RevealPassword
}

// GetRevealPasswordOk returns a tuple with the RevealPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentials) GetRevealPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.RevealPassword) {
		return nil, false
	}
	return o.RevealPassword, true
}

// HasRevealPassword returns a boolean if a field has been set.
func (o *SchemeApplicationCredentials) HasRevealPassword() bool {
	if o != nil && !IsNil(o.RevealPassword) {
		return true
	}

	return false
}

// SetRevealPassword gets a reference to the given bool and assigns it to the RevealPassword field.
func (o *SchemeApplicationCredentials) SetRevealPassword(v bool) {
	o.RevealPassword = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *SchemeApplicationCredentials) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentials) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *SchemeApplicationCredentials) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *SchemeApplicationCredentials) SetScheme(v string) {
	o.Scheme = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SchemeApplicationCredentials) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentials) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SchemeApplicationCredentials) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SchemeApplicationCredentials) SetUserName(v string) {
	o.UserName = &v
}

func (o SchemeApplicationCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemeApplicationCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signing) {
		toSerialize["signing"] = o.Signing
	}
	if !IsNil(o.UserNameTemplate) {
		toSerialize["userNameTemplate"] = o.UserNameTemplate
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.RevealPassword) {
		toSerialize["revealPassword"] = o.RevealPassword
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchemeApplicationCredentials) UnmarshalJSON(data []byte) (err error) {
	varSchemeApplicationCredentials := _SchemeApplicationCredentials{}

	err = json.Unmarshal(data, &varSchemeApplicationCredentials)

	if err != nil {
		return err
	}

	*o = SchemeApplicationCredentials(varSchemeApplicationCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signing")
		delete(additionalProperties, "userNameTemplate")
		delete(additionalProperties, "password")
		delete(additionalProperties, "revealPassword")
		delete(additionalProperties, "scheme")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchemeApplicationCredentials struct {
	value *SchemeApplicationCredentials
	isSet bool
}

func (v NullableSchemeApplicationCredentials) Get() *SchemeApplicationCredentials {
	return v.value
}

func (v *NullableSchemeApplicationCredentials) Set(val *SchemeApplicationCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemeApplicationCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemeApplicationCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemeApplicationCredentials(val *SchemeApplicationCredentials) *NullableSchemeApplicationCredentials {
	return &NullableSchemeApplicationCredentials{value: val, isSet: true}
}

func (v NullableSchemeApplicationCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemeApplicationCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
