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

// AuthenticatorKeyEmailAllOfSettings struct for AuthenticatorKeyEmailAllOfSettings
type AuthenticatorKeyEmailAllOfSettings struct {
	// The allowed types of uses for the Authenticator
	AllowedFor *string `json:"allowedFor,omitempty"`
	// Specifies the lifetime of an email token. Default value is 5 minutes.
	TokenLifetimeInMinutes *float32 `json:"tokenLifetimeInMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyEmailAllOfSettings AuthenticatorKeyEmailAllOfSettings

// NewAuthenticatorKeyEmailAllOfSettings instantiates a new AuthenticatorKeyEmailAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyEmailAllOfSettings() *AuthenticatorKeyEmailAllOfSettings {
	this := AuthenticatorKeyEmailAllOfSettings{}
	var tokenLifetimeInMinutes float32 = 5
	this.TokenLifetimeInMinutes = &tokenLifetimeInMinutes
	return &this
}

// NewAuthenticatorKeyEmailAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyEmailAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyEmailAllOfSettingsWithDefaults() *AuthenticatorKeyEmailAllOfSettings {
	this := AuthenticatorKeyEmailAllOfSettings{}
	var tokenLifetimeInMinutes float32 = 5
	this.TokenLifetimeInMinutes = &tokenLifetimeInMinutes
	return &this
}

// GetAllowedFor returns the AllowedFor field value if set, zero value otherwise.
func (o *AuthenticatorKeyEmailAllOfSettings) GetAllowedFor() string {
	if o == nil || o.AllowedFor == nil {
		var ret string
		return ret
	}
	return *o.AllowedFor
}

// GetAllowedForOk returns a tuple with the AllowedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyEmailAllOfSettings) GetAllowedForOk() (*string, bool) {
	if o == nil || o.AllowedFor == nil {
		return nil, false
	}
	return o.AllowedFor, true
}

// HasAllowedFor returns a boolean if a field has been set.
func (o *AuthenticatorKeyEmailAllOfSettings) HasAllowedFor() bool {
	if o != nil && o.AllowedFor != nil {
		return true
	}

	return false
}

// SetAllowedFor gets a reference to the given string and assigns it to the AllowedFor field.
func (o *AuthenticatorKeyEmailAllOfSettings) SetAllowedFor(v string) {
	o.AllowedFor = &v
}

// GetTokenLifetimeInMinutes returns the TokenLifetimeInMinutes field value if set, zero value otherwise.
func (o *AuthenticatorKeyEmailAllOfSettings) GetTokenLifetimeInMinutes() float32 {
	if o == nil || o.TokenLifetimeInMinutes == nil {
		var ret float32
		return ret
	}
	return *o.TokenLifetimeInMinutes
}

// GetTokenLifetimeInMinutesOk returns a tuple with the TokenLifetimeInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyEmailAllOfSettings) GetTokenLifetimeInMinutesOk() (*float32, bool) {
	if o == nil || o.TokenLifetimeInMinutes == nil {
		return nil, false
	}
	return o.TokenLifetimeInMinutes, true
}

// HasTokenLifetimeInMinutes returns a boolean if a field has been set.
func (o *AuthenticatorKeyEmailAllOfSettings) HasTokenLifetimeInMinutes() bool {
	if o != nil && o.TokenLifetimeInMinutes != nil {
		return true
	}

	return false
}

// SetTokenLifetimeInMinutes gets a reference to the given float32 and assigns it to the TokenLifetimeInMinutes field.
func (o *AuthenticatorKeyEmailAllOfSettings) SetTokenLifetimeInMinutes(v float32) {
	o.TokenLifetimeInMinutes = &v
}

func (o AuthenticatorKeyEmailAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedFor != nil {
		toSerialize["allowedFor"] = o.AllowedFor
	}
	if o.TokenLifetimeInMinutes != nil {
		toSerialize["tokenLifetimeInMinutes"] = o.TokenLifetimeInMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyEmailAllOfSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyEmailAllOfSettings := _AuthenticatorKeyEmailAllOfSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyEmailAllOfSettings)
	if err == nil {
		*o = AuthenticatorKeyEmailAllOfSettings(varAuthenticatorKeyEmailAllOfSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allowedFor")
		delete(additionalProperties, "tokenLifetimeInMinutes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyEmailAllOfSettings struct {
	value *AuthenticatorKeyEmailAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorKeyEmailAllOfSettings) Get() *AuthenticatorKeyEmailAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorKeyEmailAllOfSettings) Set(val *AuthenticatorKeyEmailAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyEmailAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyEmailAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyEmailAllOfSettings(val *AuthenticatorKeyEmailAllOfSettings) *NullableAuthenticatorKeyEmailAllOfSettings {
	return &NullableAuthenticatorKeyEmailAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyEmailAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyEmailAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

