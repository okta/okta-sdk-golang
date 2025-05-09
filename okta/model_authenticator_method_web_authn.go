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

// AuthenticatorMethodWebAuthn struct for AuthenticatorMethodWebAuthn
type AuthenticatorMethodWebAuthn struct {
	AuthenticatorMethodBase
	Settings *AuthenticatorMethodWebAuthnAllOfSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodWebAuthn AuthenticatorMethodWebAuthn

// NewAuthenticatorMethodWebAuthn instantiates a new AuthenticatorMethodWebAuthn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodWebAuthn() *AuthenticatorMethodWebAuthn {
	this := AuthenticatorMethodWebAuthn{}
	return &this
}

// NewAuthenticatorMethodWebAuthnWithDefaults instantiates a new AuthenticatorMethodWebAuthn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodWebAuthnWithDefaults() *AuthenticatorMethodWebAuthn {
	this := AuthenticatorMethodWebAuthn{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthn) GetSettings() AuthenticatorMethodWebAuthnAllOfSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorMethodWebAuthnAllOfSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthn) GetSettingsOk() (*AuthenticatorMethodWebAuthnAllOfSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthn) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorMethodWebAuthnAllOfSettings and assigns it to the Settings field.
func (o *AuthenticatorMethodWebAuthn) SetSettings(v AuthenticatorMethodWebAuthnAllOfSettings) {
	o.Settings = &v
}

func (o AuthenticatorMethodWebAuthn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorMethodBase, errAuthenticatorMethodBase := json.Marshal(o.AuthenticatorMethodBase)
	if errAuthenticatorMethodBase != nil {
		return []byte{}, errAuthenticatorMethodBase
	}
	errAuthenticatorMethodBase = json.Unmarshal([]byte(serializedAuthenticatorMethodBase), &toSerialize)
	if errAuthenticatorMethodBase != nil {
		return []byte{}, errAuthenticatorMethodBase
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodWebAuthn) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorMethodWebAuthnWithoutEmbeddedStruct struct {
		Settings *AuthenticatorMethodWebAuthnAllOfSettings `json:"settings,omitempty"`
	}

	varAuthenticatorMethodWebAuthnWithoutEmbeddedStruct := AuthenticatorMethodWebAuthnWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodWebAuthnWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorMethodWebAuthn := _AuthenticatorMethodWebAuthn{}
		varAuthenticatorMethodWebAuthn.Settings = varAuthenticatorMethodWebAuthnWithoutEmbeddedStruct.Settings
		*o = AuthenticatorMethodWebAuthn(varAuthenticatorMethodWebAuthn)
	} else {
		return err
	}

	varAuthenticatorMethodWebAuthn := _AuthenticatorMethodWebAuthn{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodWebAuthn)
	if err == nil {
		o.AuthenticatorMethodBase = varAuthenticatorMethodWebAuthn.AuthenticatorMethodBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectAuthenticatorMethodBase := reflect.ValueOf(o.AuthenticatorMethodBase)
		for i := 0; i < reflectAuthenticatorMethodBase.Type().NumField(); i++ {
			t := reflectAuthenticatorMethodBase.Type().Field(i)

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

type NullableAuthenticatorMethodWebAuthn struct {
	value *AuthenticatorMethodWebAuthn
	isSet bool
}

func (v NullableAuthenticatorMethodWebAuthn) Get() *AuthenticatorMethodWebAuthn {
	return v.value
}

func (v *NullableAuthenticatorMethodWebAuthn) Set(val *AuthenticatorMethodWebAuthn) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodWebAuthn) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodWebAuthn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodWebAuthn(val *AuthenticatorMethodWebAuthn) *NullableAuthenticatorMethodWebAuthn {
	return &NullableAuthenticatorMethodWebAuthn{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodWebAuthn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodWebAuthn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

