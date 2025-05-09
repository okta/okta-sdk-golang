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

// AuthenticatorKeyOktaVerify struct for AuthenticatorKeyOktaVerify
type AuthenticatorKeyOktaVerify struct {
	AuthenticatorSimple
	Settings *AuthenticatorKeyOktaVerifyAllOfSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyOktaVerify AuthenticatorKeyOktaVerify

// NewAuthenticatorKeyOktaVerify instantiates a new AuthenticatorKeyOktaVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyOktaVerify() *AuthenticatorKeyOktaVerify {
	this := AuthenticatorKeyOktaVerify{}
	return &this
}

// NewAuthenticatorKeyOktaVerifyWithDefaults instantiates a new AuthenticatorKeyOktaVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyOktaVerifyWithDefaults() *AuthenticatorKeyOktaVerify {
	this := AuthenticatorKeyOktaVerify{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorKeyOktaVerify) GetSettings() AuthenticatorKeyOktaVerifyAllOfSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorKeyOktaVerifyAllOfSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyOktaVerify) GetSettingsOk() (*AuthenticatorKeyOktaVerifyAllOfSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorKeyOktaVerify) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorKeyOktaVerifyAllOfSettings and assigns it to the Settings field.
func (o *AuthenticatorKeyOktaVerify) SetSettings(v AuthenticatorKeyOktaVerifyAllOfSettings) {
	o.Settings = &v
}

func (o AuthenticatorKeyOktaVerify) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorKeyOktaVerify) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyOktaVerifyWithoutEmbeddedStruct struct {
		Settings *AuthenticatorKeyOktaVerifyAllOfSettings `json:"settings,omitempty"`
	}

	varAuthenticatorKeyOktaVerifyWithoutEmbeddedStruct := AuthenticatorKeyOktaVerifyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyOktaVerifyWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyOktaVerify := _AuthenticatorKeyOktaVerify{}
		varAuthenticatorKeyOktaVerify.Settings = varAuthenticatorKeyOktaVerifyWithoutEmbeddedStruct.Settings
		*o = AuthenticatorKeyOktaVerify(varAuthenticatorKeyOktaVerify)
	} else {
		return err
	}

	varAuthenticatorKeyOktaVerify := _AuthenticatorKeyOktaVerify{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyOktaVerify)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyOktaVerify.AuthenticatorSimple
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

type NullableAuthenticatorKeyOktaVerify struct {
	value *AuthenticatorKeyOktaVerify
	isSet bool
}

func (v NullableAuthenticatorKeyOktaVerify) Get() *AuthenticatorKeyOktaVerify {
	return v.value
}

func (v *NullableAuthenticatorKeyOktaVerify) Set(val *AuthenticatorKeyOktaVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyOktaVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyOktaVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyOktaVerify(val *AuthenticatorKeyOktaVerify) *NullableAuthenticatorKeyOktaVerify {
	return &NullableAuthenticatorKeyOktaVerify{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyOktaVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyOktaVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

