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

// AuthenticatorKeyEmail struct for AuthenticatorKeyEmail
type AuthenticatorKeyEmail struct {
	AuthenticatorSimple
	Settings *AuthenticatorKeyEmailAllOfSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyEmail AuthenticatorKeyEmail

// NewAuthenticatorKeyEmail instantiates a new AuthenticatorKeyEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyEmail() *AuthenticatorKeyEmail {
	this := AuthenticatorKeyEmail{}
	return &this
}

// NewAuthenticatorKeyEmailWithDefaults instantiates a new AuthenticatorKeyEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyEmailWithDefaults() *AuthenticatorKeyEmail {
	this := AuthenticatorKeyEmail{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorKeyEmail) GetSettings() AuthenticatorKeyEmailAllOfSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorKeyEmailAllOfSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyEmail) GetSettingsOk() (*AuthenticatorKeyEmailAllOfSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorKeyEmail) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorKeyEmailAllOfSettings and assigns it to the Settings field.
func (o *AuthenticatorKeyEmail) SetSettings(v AuthenticatorKeyEmailAllOfSettings) {
	o.Settings = &v
}

func (o AuthenticatorKeyEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorSimple, errAuthenticatorSimple := json.Marshal(o.AuthenticatorSimple)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	errAuthenticatorSimple = json.Unmarshal([]byte(serializedAuthenticatorSimple), &toSerialize)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyEmail) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyEmailWithoutEmbeddedStruct struct {
		Settings *AuthenticatorKeyEmailAllOfSettings `json:"settings,omitempty"`
	}

	varAuthenticatorKeyEmailWithoutEmbeddedStruct := AuthenticatorKeyEmailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyEmailWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyEmail := _AuthenticatorKeyEmail{}
		varAuthenticatorKeyEmail.Settings = varAuthenticatorKeyEmailWithoutEmbeddedStruct.Settings
		*o = AuthenticatorKeyEmail(varAuthenticatorKeyEmail)
	} else {
		return err
	}

	varAuthenticatorKeyEmail := _AuthenticatorKeyEmail{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyEmail)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyEmail.AuthenticatorSimple
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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

type NullableAuthenticatorKeyEmail struct {
	value *AuthenticatorKeyEmail
	isSet bool
}

func (v NullableAuthenticatorKeyEmail) Get() *AuthenticatorKeyEmail {
	return v.value
}

func (v *NullableAuthenticatorKeyEmail) Set(val *AuthenticatorKeyEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyEmail(val *AuthenticatorKeyEmail) *NullableAuthenticatorKeyEmail {
	return &NullableAuthenticatorKeyEmail{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

