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

// AuthenticatorKeyPhoneAllOfSettings struct for AuthenticatorKeyPhoneAllOfSettings
type AuthenticatorKeyPhoneAllOfSettings struct {
	// The allowed types of uses for the Authenticator
	AllowedFor *string `json:"allowedFor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyPhoneAllOfSettings AuthenticatorKeyPhoneAllOfSettings

// NewAuthenticatorKeyPhoneAllOfSettings instantiates a new AuthenticatorKeyPhoneAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyPhoneAllOfSettings() *AuthenticatorKeyPhoneAllOfSettings {
	this := AuthenticatorKeyPhoneAllOfSettings{}
	return &this
}

// NewAuthenticatorKeyPhoneAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyPhoneAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyPhoneAllOfSettingsWithDefaults() *AuthenticatorKeyPhoneAllOfSettings {
	this := AuthenticatorKeyPhoneAllOfSettings{}
	return &this
}

// GetAllowedFor returns the AllowedFor field value if set, zero value otherwise.
func (o *AuthenticatorKeyPhoneAllOfSettings) GetAllowedFor() string {
	if o == nil || o.AllowedFor == nil {
		var ret string
		return ret
	}
	return *o.AllowedFor
}

// GetAllowedForOk returns a tuple with the AllowedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyPhoneAllOfSettings) GetAllowedForOk() (*string, bool) {
	if o == nil || o.AllowedFor == nil {
		return nil, false
	}
	return o.AllowedFor, true
}

// HasAllowedFor returns a boolean if a field has been set.
func (o *AuthenticatorKeyPhoneAllOfSettings) HasAllowedFor() bool {
	if o != nil && o.AllowedFor != nil {
		return true
	}

	return false
}

// SetAllowedFor gets a reference to the given string and assigns it to the AllowedFor field.
func (o *AuthenticatorKeyPhoneAllOfSettings) SetAllowedFor(v string) {
	o.AllowedFor = &v
}

func (o AuthenticatorKeyPhoneAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedFor != nil {
		toSerialize["allowedFor"] = o.AllowedFor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyPhoneAllOfSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyPhoneAllOfSettings := _AuthenticatorKeyPhoneAllOfSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyPhoneAllOfSettings)
	if err == nil {
		*o = AuthenticatorKeyPhoneAllOfSettings(varAuthenticatorKeyPhoneAllOfSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allowedFor")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyPhoneAllOfSettings struct {
	value *AuthenticatorKeyPhoneAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorKeyPhoneAllOfSettings) Get() *AuthenticatorKeyPhoneAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorKeyPhoneAllOfSettings) Set(val *AuthenticatorKeyPhoneAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyPhoneAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyPhoneAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyPhoneAllOfSettings(val *AuthenticatorKeyPhoneAllOfSettings) *NullableAuthenticatorKeyPhoneAllOfSettings {
	return &NullableAuthenticatorKeyPhoneAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyPhoneAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyPhoneAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

