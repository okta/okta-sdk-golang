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

// AuthenticatorMethodTotp struct for AuthenticatorMethodTotp
type AuthenticatorMethodTotp struct {
	AuthenticatorMethodBase
	Settings *AuthenticatorMethodTotpAllOfSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodTotp AuthenticatorMethodTotp

// NewAuthenticatorMethodTotp instantiates a new AuthenticatorMethodTotp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodTotp() *AuthenticatorMethodTotp {
	this := AuthenticatorMethodTotp{}
	return &this
}

// NewAuthenticatorMethodTotpWithDefaults instantiates a new AuthenticatorMethodTotp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodTotpWithDefaults() *AuthenticatorMethodTotp {
	this := AuthenticatorMethodTotp{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorMethodTotp) GetSettings() AuthenticatorMethodTotpAllOfSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorMethodTotpAllOfSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTotp) GetSettingsOk() (*AuthenticatorMethodTotpAllOfSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorMethodTotp) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorMethodTotpAllOfSettings and assigns it to the Settings field.
func (o *AuthenticatorMethodTotp) SetSettings(v AuthenticatorMethodTotpAllOfSettings) {
	o.Settings = &v
}

func (o AuthenticatorMethodTotp) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorMethodTotp) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorMethodTotpWithoutEmbeddedStruct struct {
		Settings *AuthenticatorMethodTotpAllOfSettings `json:"settings,omitempty"`
	}

	varAuthenticatorMethodTotpWithoutEmbeddedStruct := AuthenticatorMethodTotpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodTotpWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorMethodTotp := _AuthenticatorMethodTotp{}
		varAuthenticatorMethodTotp.Settings = varAuthenticatorMethodTotpWithoutEmbeddedStruct.Settings
		*o = AuthenticatorMethodTotp(varAuthenticatorMethodTotp)
	} else {
		return err
	}

	varAuthenticatorMethodTotp := _AuthenticatorMethodTotp{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodTotp)
	if err == nil {
		o.AuthenticatorMethodBase = varAuthenticatorMethodTotp.AuthenticatorMethodBase
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

type NullableAuthenticatorMethodTotp struct {
	value *AuthenticatorMethodTotp
	isSet bool
}

func (v NullableAuthenticatorMethodTotp) Get() *AuthenticatorMethodTotp {
	return v.value
}

func (v *NullableAuthenticatorMethodTotp) Set(val *AuthenticatorMethodTotp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodTotp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodTotp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodTotp(val *AuthenticatorMethodTotp) *NullableAuthenticatorMethodTotp {
	return &NullableAuthenticatorMethodTotp{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodTotp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodTotp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

