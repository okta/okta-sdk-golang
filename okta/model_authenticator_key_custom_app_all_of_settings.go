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

// checks if the AuthenticatorKeyCustomAppAllOfSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorKeyCustomAppAllOfSettings{}

// AuthenticatorKeyCustomAppAllOfSettings struct for AuthenticatorKeyCustomAppAllOfSettings
type AuthenticatorKeyCustomAppAllOfSettings struct {
	// User verification setting
	UserVerification *string `json:"userVerification,omitempty"`
	// The application instance ID. For custom_app, you need to create an OIDC native app using the [Apps API](https://developer.okta.com/docs/reference/api/apps/) with `Authorization Code` and `Refresh Token` grant types. You can leave both `Sign-in redirect URIs` and `Sign-out redirect URIs` as the default values.
	AppInstanceId        *string `json:"appInstanceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyCustomAppAllOfSettings AuthenticatorKeyCustomAppAllOfSettings

// NewAuthenticatorKeyCustomAppAllOfSettings instantiates a new AuthenticatorKeyCustomAppAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyCustomAppAllOfSettings() *AuthenticatorKeyCustomAppAllOfSettings {
	this := AuthenticatorKeyCustomAppAllOfSettings{}
	return &this
}

// NewAuthenticatorKeyCustomAppAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyCustomAppAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyCustomAppAllOfSettingsWithDefaults() *AuthenticatorKeyCustomAppAllOfSettings {
	this := AuthenticatorKeyCustomAppAllOfSettings{}
	return &this
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfSettings) GetUserVerification() string {
	if o == nil || IsNil(o.UserVerification) {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfSettings) GetUserVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.UserVerification) {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfSettings) HasUserVerification() bool {
	if o != nil && !IsNil(o.UserVerification) {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *AuthenticatorKeyCustomAppAllOfSettings) SetUserVerification(v string) {
	o.UserVerification = &v
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfSettings) GetAppInstanceId() string {
	if o == nil || IsNil(o.AppInstanceId) {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfSettings) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceId) {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfSettings) HasAppInstanceId() bool {
	if o != nil && !IsNil(o.AppInstanceId) {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *AuthenticatorKeyCustomAppAllOfSettings) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

func (o AuthenticatorKeyCustomAppAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorKeyCustomAppAllOfSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserVerification) {
		toSerialize["userVerification"] = o.UserVerification
	}
	if !IsNil(o.AppInstanceId) {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorKeyCustomAppAllOfSettings) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorKeyCustomAppAllOfSettings := _AuthenticatorKeyCustomAppAllOfSettings{}

	err = json.Unmarshal(data, &varAuthenticatorKeyCustomAppAllOfSettings)

	if err != nil {
		return err
	}

	*o = AuthenticatorKeyCustomAppAllOfSettings(varAuthenticatorKeyCustomAppAllOfSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userVerification")
		delete(additionalProperties, "appInstanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorKeyCustomAppAllOfSettings struct {
	value *AuthenticatorKeyCustomAppAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorKeyCustomAppAllOfSettings) Get() *AuthenticatorKeyCustomAppAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorKeyCustomAppAllOfSettings) Set(val *AuthenticatorKeyCustomAppAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyCustomAppAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyCustomAppAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyCustomAppAllOfSettings(val *AuthenticatorKeyCustomAppAllOfSettings) *NullableAuthenticatorKeyCustomAppAllOfSettings {
	return &NullableAuthenticatorKeyCustomAppAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyCustomAppAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyCustomAppAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
