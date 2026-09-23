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

// checks if the UserIdentificationFastpassSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationFastpassSettings{}

// UserIdentificationFastpassSettings <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the Okta FastPass user identification settings
type UserIdentificationFastpassSettings struct {
	// Controls whether to show the Sign in with Okta Verify button on the Sign-In Widget
	ShowSignInButton     *string `json:"showSignInButton,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationFastpassSettings UserIdentificationFastpassSettings

// NewUserIdentificationFastpassSettings instantiates a new UserIdentificationFastpassSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationFastpassSettings() *UserIdentificationFastpassSettings {
	this := UserIdentificationFastpassSettings{}
	return &this
}

// NewUserIdentificationFastpassSettingsWithDefaults instantiates a new UserIdentificationFastpassSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationFastpassSettingsWithDefaults() *UserIdentificationFastpassSettings {
	this := UserIdentificationFastpassSettings{}
	return &this
}

// GetShowSignInButton returns the ShowSignInButton field value if set, zero value otherwise.
func (o *UserIdentificationFastpassSettings) GetShowSignInButton() string {
	if o == nil || IsNil(o.ShowSignInButton) {
		var ret string
		return ret
	}
	return *o.ShowSignInButton
}

// GetShowSignInButtonOk returns a tuple with the ShowSignInButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationFastpassSettings) GetShowSignInButtonOk() (*string, bool) {
	if o == nil || IsNil(o.ShowSignInButton) {
		return nil, false
	}
	return o.ShowSignInButton, true
}

// HasShowSignInButton returns a boolean if a field has been set.
func (o *UserIdentificationFastpassSettings) HasShowSignInButton() bool {
	if o != nil && !IsNil(o.ShowSignInButton) {
		return true
	}

	return false
}

// SetShowSignInButton gets a reference to the given string and assigns it to the ShowSignInButton field.
func (o *UserIdentificationFastpassSettings) SetShowSignInButton(v string) {
	o.ShowSignInButton = &v
}

func (o UserIdentificationFastpassSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationFastpassSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShowSignInButton) {
		toSerialize["showSignInButton"] = o.ShowSignInButton
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationFastpassSettings) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationFastpassSettings := _UserIdentificationFastpassSettings{}

	err = json.Unmarshal(data, &varUserIdentificationFastpassSettings)

	if err != nil {
		return err
	}

	*o = UserIdentificationFastpassSettings(varUserIdentificationFastpassSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "showSignInButton")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationFastpassSettings struct {
	value *UserIdentificationFastpassSettings
	isSet bool
}

func (v NullableUserIdentificationFastpassSettings) Get() *UserIdentificationFastpassSettings {
	return v.value
}

func (v *NullableUserIdentificationFastpassSettings) Set(val *UserIdentificationFastpassSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationFastpassSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationFastpassSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationFastpassSettings(val *UserIdentificationFastpassSettings) *NullableUserIdentificationFastpassSettings {
	return &NullableUserIdentificationFastpassSettings{value: val, isSet: true}
}

func (v NullableUserIdentificationFastpassSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationFastpassSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
