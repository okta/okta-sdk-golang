/*
Okta Admin Management API

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

// checks if the UserIdentificationPasskeySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationPasskeySettings{}

// UserIdentificationPasskeySettings <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the passkey user identification settings
type UserIdentificationPasskeySettings struct {
	// Controls whether to show the **Sign in with a passkey** button on the Sign-In Widget.  > **Note:** You can only set `showSignInButton` to `ALWAYS` when the Passkey (FIDO2 WebAuthn) authenticator is active.
	ShowSignInButton     *string `json:"showSignInButton,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationPasskeySettings UserIdentificationPasskeySettings

// NewUserIdentificationPasskeySettings instantiates a new UserIdentificationPasskeySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationPasskeySettings() *UserIdentificationPasskeySettings {
	this := UserIdentificationPasskeySettings{}
	return &this
}

// NewUserIdentificationPasskeySettingsWithDefaults instantiates a new UserIdentificationPasskeySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationPasskeySettingsWithDefaults() *UserIdentificationPasskeySettings {
	this := UserIdentificationPasskeySettings{}
	return &this
}

// GetShowSignInButton returns the ShowSignInButton field value if set, zero value otherwise.
func (o *UserIdentificationPasskeySettings) GetShowSignInButton() string {
	if o == nil || IsNil(o.ShowSignInButton) {
		var ret string
		return ret
	}
	return *o.ShowSignInButton
}

// GetShowSignInButtonOk returns a tuple with the ShowSignInButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationPasskeySettings) GetShowSignInButtonOk() (*string, bool) {
	if o == nil || IsNil(o.ShowSignInButton) {
		return nil, false
	}
	return o.ShowSignInButton, true
}

// HasShowSignInButton returns a boolean if a field has been set.
func (o *UserIdentificationPasskeySettings) HasShowSignInButton() bool {
	if o != nil && !IsNil(o.ShowSignInButton) {
		return true
	}

	return false
}

// SetShowSignInButton gets a reference to the given string and assigns it to the ShowSignInButton field.
func (o *UserIdentificationPasskeySettings) SetShowSignInButton(v string) {
	o.ShowSignInButton = &v
}

func (o UserIdentificationPasskeySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationPasskeySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShowSignInButton) {
		toSerialize["showSignInButton"] = o.ShowSignInButton
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationPasskeySettings) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationPasskeySettings := _UserIdentificationPasskeySettings{}

	err = json.Unmarshal(data, &varUserIdentificationPasskeySettings)

	if err != nil {
		return err
	}

	*o = UserIdentificationPasskeySettings(varUserIdentificationPasskeySettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "showSignInButton")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationPasskeySettings struct {
	value *UserIdentificationPasskeySettings
	isSet bool
}

func (v NullableUserIdentificationPasskeySettings) Get() *UserIdentificationPasskeySettings {
	return v.value
}

func (v *NullableUserIdentificationPasskeySettings) Set(val *UserIdentificationPasskeySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationPasskeySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationPasskeySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationPasskeySettings(val *UserIdentificationPasskeySettings) *NullableUserIdentificationPasskeySettings {
	return &NullableUserIdentificationPasskeySettings{value: val, isSet: true}
}

func (v NullableUserIdentificationPasskeySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationPasskeySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
