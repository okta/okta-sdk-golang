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

// AuthenticatorKeySecurityQuestion struct for AuthenticatorKeySecurityQuestion
type AuthenticatorKeySecurityQuestion struct {
	AuthenticatorSimple
	Settings *AuthenticatorKeyPhoneAllOfSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeySecurityQuestion AuthenticatorKeySecurityQuestion

// NewAuthenticatorKeySecurityQuestion instantiates a new AuthenticatorKeySecurityQuestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeySecurityQuestion() *AuthenticatorKeySecurityQuestion {
	this := AuthenticatorKeySecurityQuestion{}
	return &this
}

// NewAuthenticatorKeySecurityQuestionWithDefaults instantiates a new AuthenticatorKeySecurityQuestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeySecurityQuestionWithDefaults() *AuthenticatorKeySecurityQuestion {
	this := AuthenticatorKeySecurityQuestion{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorKeySecurityQuestion) GetSettings() AuthenticatorKeyPhoneAllOfSettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorKeyPhoneAllOfSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeySecurityQuestion) GetSettingsOk() (*AuthenticatorKeyPhoneAllOfSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorKeySecurityQuestion) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorKeyPhoneAllOfSettings and assigns it to the Settings field.
func (o *AuthenticatorKeySecurityQuestion) SetSettings(v AuthenticatorKeyPhoneAllOfSettings) {
	o.Settings = &v
}

func (o AuthenticatorKeySecurityQuestion) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorKeySecurityQuestion) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeySecurityQuestionWithoutEmbeddedStruct struct {
		Settings *AuthenticatorKeyPhoneAllOfSettings `json:"settings,omitempty"`
	}

	varAuthenticatorKeySecurityQuestionWithoutEmbeddedStruct := AuthenticatorKeySecurityQuestionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeySecurityQuestionWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeySecurityQuestion := _AuthenticatorKeySecurityQuestion{}
		varAuthenticatorKeySecurityQuestion.Settings = varAuthenticatorKeySecurityQuestionWithoutEmbeddedStruct.Settings
		*o = AuthenticatorKeySecurityQuestion(varAuthenticatorKeySecurityQuestion)
	} else {
		return err
	}

	varAuthenticatorKeySecurityQuestion := _AuthenticatorKeySecurityQuestion{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeySecurityQuestion)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeySecurityQuestion.AuthenticatorSimple
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

type NullableAuthenticatorKeySecurityQuestion struct {
	value *AuthenticatorKeySecurityQuestion
	isSet bool
}

func (v NullableAuthenticatorKeySecurityQuestion) Get() *AuthenticatorKeySecurityQuestion {
	return v.value
}

func (v *NullableAuthenticatorKeySecurityQuestion) Set(val *AuthenticatorKeySecurityQuestion) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeySecurityQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeySecurityQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeySecurityQuestion(val *AuthenticatorKeySecurityQuestion) *NullableAuthenticatorKeySecurityQuestion {
	return &NullableAuthenticatorKeySecurityQuestion{value: val, isSet: true}
}

func (v NullableAuthenticatorKeySecurityQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeySecurityQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

