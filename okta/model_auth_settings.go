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
	"fmt"
)

// checks if the AuthSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthSettings{}

// AuthSettings struct for AuthSettings
type AuthSettings struct {
	AuthType             string              `json:"authType"`
	CustomSettings       *CustomAuthSettings `json:"customSettings,omitempty"`
	OAuth2Settings       *OAuth2Settings     `json:"oAuth2Settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthSettings AuthSettings

// NewAuthSettings instantiates a new AuthSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthSettings(authType string) *AuthSettings {
	this := AuthSettings{}
	this.AuthType = authType
	return &this
}

// NewAuthSettingsWithDefaults instantiates a new AuthSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthSettingsWithDefaults() *AuthSettings {
	this := AuthSettings{}
	return &this
}

// GetAuthType returns the AuthType field value
func (o *AuthSettings) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *AuthSettings) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *AuthSettings) SetAuthType(v string) {
	o.AuthType = v
}

// GetCustomSettings returns the CustomSettings field value if set, zero value otherwise.
func (o *AuthSettings) GetCustomSettings() CustomAuthSettings {
	if o == nil || IsNil(o.CustomSettings) {
		var ret CustomAuthSettings
		return ret
	}
	return *o.CustomSettings
}

// GetCustomSettingsOk returns a tuple with the CustomSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSettings) GetCustomSettingsOk() (*CustomAuthSettings, bool) {
	if o == nil || IsNil(o.CustomSettings) {
		return nil, false
	}
	return o.CustomSettings, true
}

// HasCustomSettings returns a boolean if a field has been set.
func (o *AuthSettings) HasCustomSettings() bool {
	if o != nil && !IsNil(o.CustomSettings) {
		return true
	}

	return false
}

// SetCustomSettings gets a reference to the given CustomAuthSettings and assigns it to the CustomSettings field.
func (o *AuthSettings) SetCustomSettings(v CustomAuthSettings) {
	o.CustomSettings = &v
}

// GetOAuth2Settings returns the OAuth2Settings field value if set, zero value otherwise.
func (o *AuthSettings) GetOAuth2Settings() OAuth2Settings {
	if o == nil || IsNil(o.OAuth2Settings) {
		var ret OAuth2Settings
		return ret
	}
	return *o.OAuth2Settings
}

// GetOAuth2SettingsOk returns a tuple with the OAuth2Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSettings) GetOAuth2SettingsOk() (*OAuth2Settings, bool) {
	if o == nil || IsNil(o.OAuth2Settings) {
		return nil, false
	}
	return o.OAuth2Settings, true
}

// HasOAuth2Settings returns a boolean if a field has been set.
func (o *AuthSettings) HasOAuth2Settings() bool {
	if o != nil && !IsNil(o.OAuth2Settings) {
		return true
	}

	return false
}

// SetOAuth2Settings gets a reference to the given OAuth2Settings and assigns it to the OAuth2Settings field.
func (o *AuthSettings) SetOAuth2Settings(v OAuth2Settings) {
	o.OAuth2Settings = &v
}

func (o AuthSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authType"] = o.AuthType
	if !IsNil(o.CustomSettings) {
		toSerialize["customSettings"] = o.CustomSettings
	}
	if !IsNil(o.OAuth2Settings) {
		toSerialize["oAuth2Settings"] = o.OAuth2Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthSettings := _AuthSettings{}

	err = json.Unmarshal(data, &varAuthSettings)

	if err != nil {
		return err
	}

	*o = AuthSettings(varAuthSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authType")
		delete(additionalProperties, "customSettings")
		delete(additionalProperties, "oAuth2Settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthSettings struct {
	value *AuthSettings
	isSet bool
}

func (v NullableAuthSettings) Get() *AuthSettings {
	return v.value
}

func (v *NullableAuthSettings) Set(val *AuthSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthSettings(val *AuthSettings) *NullableAuthSettings {
	return &NullableAuthSettings{value: val, isSet: true}
}

func (v NullableAuthSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
