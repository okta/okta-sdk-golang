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

// WellKnownAppAuthenticatorConfigurationSettings struct for WellKnownAppAuthenticatorConfigurationSettings
type WellKnownAppAuthenticatorConfigurationSettings struct {
	// User verification setting
	UserVerification *string `json:"userVerification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownAppAuthenticatorConfigurationSettings WellKnownAppAuthenticatorConfigurationSettings

// NewWellKnownAppAuthenticatorConfigurationSettings instantiates a new WellKnownAppAuthenticatorConfigurationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownAppAuthenticatorConfigurationSettings() *WellKnownAppAuthenticatorConfigurationSettings {
	this := WellKnownAppAuthenticatorConfigurationSettings{}
	return &this
}

// NewWellKnownAppAuthenticatorConfigurationSettingsWithDefaults instantiates a new WellKnownAppAuthenticatorConfigurationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownAppAuthenticatorConfigurationSettingsWithDefaults() *WellKnownAppAuthenticatorConfigurationSettings {
	this := WellKnownAppAuthenticatorConfigurationSettings{}
	return &this
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfigurationSettings) GetUserVerification() string {
	if o == nil || o.UserVerification == nil {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfigurationSettings) GetUserVerificationOk() (*string, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfigurationSettings) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *WellKnownAppAuthenticatorConfigurationSettings) SetUserVerification(v string) {
	o.UserVerification = &v
}

func (o WellKnownAppAuthenticatorConfigurationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserVerification != nil {
		toSerialize["userVerification"] = o.UserVerification
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownAppAuthenticatorConfigurationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownAppAuthenticatorConfigurationSettings := _WellKnownAppAuthenticatorConfigurationSettings{}

	err = json.Unmarshal(bytes, &varWellKnownAppAuthenticatorConfigurationSettings)
	if err == nil {
		*o = WellKnownAppAuthenticatorConfigurationSettings(varWellKnownAppAuthenticatorConfigurationSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "userVerification")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWellKnownAppAuthenticatorConfigurationSettings struct {
	value *WellKnownAppAuthenticatorConfigurationSettings
	isSet bool
}

func (v NullableWellKnownAppAuthenticatorConfigurationSettings) Get() *WellKnownAppAuthenticatorConfigurationSettings {
	return v.value
}

func (v *NullableWellKnownAppAuthenticatorConfigurationSettings) Set(val *WellKnownAppAuthenticatorConfigurationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownAppAuthenticatorConfigurationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownAppAuthenticatorConfigurationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownAppAuthenticatorConfigurationSettings(val *WellKnownAppAuthenticatorConfigurationSettings) *NullableWellKnownAppAuthenticatorConfigurationSettings {
	return &NullableWellKnownAppAuthenticatorConfigurationSettings{value: val, isSet: true}
}

func (v NullableWellKnownAppAuthenticatorConfigurationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownAppAuthenticatorConfigurationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

