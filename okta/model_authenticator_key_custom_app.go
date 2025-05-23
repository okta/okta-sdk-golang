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
	"reflect"
	"strings"
)

// AuthenticatorKeyCustomApp struct for AuthenticatorKeyCustomApp
type AuthenticatorKeyCustomApp struct {
	AuthenticatorSimple
	// A value of `true` indicates that the administrator accepts the [terms](https://www.okta.com/privacy-policy/)for creating a new authenticator. Okta requires that you accept the terms when creating a new `custom_app` authenticator. Other authenticators don't require this field.
	AgreeToTerms *bool `json:"agreeToTerms,omitempty"`
	Provider *AuthenticatorKeyCustomAppAllOfProvider `json:"provider,omitempty"`
	Settings *AuthenticatorKeyCustomAppAllOfSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyCustomApp AuthenticatorKeyCustomApp

// NewAuthenticatorKeyCustomApp instantiates a new AuthenticatorKeyCustomApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyCustomApp() *AuthenticatorKeyCustomApp {
	this := AuthenticatorKeyCustomApp{}
	return &this
}

// NewAuthenticatorKeyCustomAppWithDefaults instantiates a new AuthenticatorKeyCustomApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyCustomAppWithDefaults() *AuthenticatorKeyCustomApp {
	this := AuthenticatorKeyCustomApp{}
	return &this
}

// GetAgreeToTerms returns the AgreeToTerms field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomApp) GetAgreeToTerms() bool {
	if o == nil || o.AgreeToTerms == nil {
		var ret bool
		return ret
	}
	return *o.AgreeToTerms
}

// GetAgreeToTermsOk returns a tuple with the AgreeToTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomApp) GetAgreeToTermsOk() (*bool, bool) {
	if o == nil || o.AgreeToTerms == nil {
		return nil, false
	}
	return o.AgreeToTerms, true
}

// HasAgreeToTerms returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomApp) HasAgreeToTerms() bool {
	if o != nil && o.AgreeToTerms != nil {
		return true
	}

	return false
}

// SetAgreeToTerms gets a reference to the given bool and assigns it to the AgreeToTerms field.
func (o *AuthenticatorKeyCustomApp) SetAgreeToTerms(v bool) {
	o.AgreeToTerms = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomApp) GetProvider() AuthenticatorKeyCustomAppAllOfProvider {
	if o == nil || o.Provider == nil {
		var ret AuthenticatorKeyCustomAppAllOfProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomApp) GetProviderOk() (*AuthenticatorKeyCustomAppAllOfProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomApp) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AuthenticatorKeyCustomAppAllOfProvider and assigns it to the Provider field.
func (o *AuthenticatorKeyCustomApp) SetProvider(v AuthenticatorKeyCustomAppAllOfProvider) {
	o.Provider = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomApp) GetSettings() AuthenticatorKeyCustomAppAllOfSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorKeyCustomAppAllOfSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomApp) GetSettingsOk() (*AuthenticatorKeyCustomAppAllOfSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomApp) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorKeyCustomAppAllOfSettings and assigns it to the Settings field.
func (o *AuthenticatorKeyCustomApp) SetSettings(v AuthenticatorKeyCustomAppAllOfSettings) {
	o.Settings = &v
}

func (o AuthenticatorKeyCustomApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorSimple, errAuthenticatorSimple := json.Marshal(o.AuthenticatorSimple)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	errAuthenticatorSimple = json.Unmarshal([]byte(serializedAuthenticatorSimple), &toSerialize)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	if o.AgreeToTerms != nil {
		toSerialize["agreeToTerms"] = o.AgreeToTerms
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyCustomApp) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyCustomAppWithoutEmbeddedStruct struct {
		// A value of `true` indicates that the administrator accepts the [terms](https://www.okta.com/privacy-policy/)for creating a new authenticator. Okta requires that you accept the terms when creating a new `custom_app` authenticator. Other authenticators don't require this field.
		AgreeToTerms *bool `json:"agreeToTerms,omitempty"`
		Provider *AuthenticatorKeyCustomAppAllOfProvider `json:"provider,omitempty"`
		Settings *AuthenticatorKeyCustomAppAllOfSettings `json:"settings,omitempty"`
	}

	varAuthenticatorKeyCustomAppWithoutEmbeddedStruct := AuthenticatorKeyCustomAppWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyCustomAppWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyCustomApp := _AuthenticatorKeyCustomApp{}
		varAuthenticatorKeyCustomApp.AgreeToTerms = varAuthenticatorKeyCustomAppWithoutEmbeddedStruct.AgreeToTerms
		varAuthenticatorKeyCustomApp.Provider = varAuthenticatorKeyCustomAppWithoutEmbeddedStruct.Provider
		varAuthenticatorKeyCustomApp.Settings = varAuthenticatorKeyCustomAppWithoutEmbeddedStruct.Settings
		*o = AuthenticatorKeyCustomApp(varAuthenticatorKeyCustomApp)
	} else {
		return err
	}

	varAuthenticatorKeyCustomApp := _AuthenticatorKeyCustomApp{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyCustomApp)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyCustomApp.AuthenticatorSimple
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "agreeToTerms")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectAuthenticatorSimple := reflect.ValueOf(o.AuthenticatorSimple)
		for i := 0; i < reflectAuthenticatorSimple.Type().NumField(); i++ {
			t := reflectAuthenticatorSimple.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyCustomApp struct {
	value *AuthenticatorKeyCustomApp
	isSet bool
}

func (v NullableAuthenticatorKeyCustomApp) Get() *AuthenticatorKeyCustomApp {
	return v.value
}

func (v *NullableAuthenticatorKeyCustomApp) Set(val *AuthenticatorKeyCustomApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyCustomApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyCustomApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyCustomApp(val *AuthenticatorKeyCustomApp) *NullableAuthenticatorKeyCustomApp {
	return &NullableAuthenticatorKeyCustomApp{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyCustomApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyCustomApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

