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

// AuthenticatorKeyPassword struct for AuthenticatorKeyPassword
type AuthenticatorKeyPassword struct {
	AuthenticatorSimple
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyPassword AuthenticatorKeyPassword

// NewAuthenticatorKeyPassword instantiates a new AuthenticatorKeyPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyPassword() *AuthenticatorKeyPassword {
	this := AuthenticatorKeyPassword{}
	return &this
}

// NewAuthenticatorKeyPasswordWithDefaults instantiates a new AuthenticatorKeyPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyPasswordWithDefaults() *AuthenticatorKeyPassword {
	this := AuthenticatorKeyPassword{}
	return &this
}

func (o AuthenticatorKeyPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorSimple, errAuthenticatorSimple := json.Marshal(o.AuthenticatorSimple)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	errAuthenticatorSimple = json.Unmarshal([]byte(serializedAuthenticatorSimple), &toSerialize)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyPassword) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyPasswordWithoutEmbeddedStruct struct {
	}

	varAuthenticatorKeyPasswordWithoutEmbeddedStruct := AuthenticatorKeyPasswordWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyPasswordWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyPassword := _AuthenticatorKeyPassword{}
		*o = AuthenticatorKeyPassword(varAuthenticatorKeyPassword)
	} else {
		return err
	}

	varAuthenticatorKeyPassword := _AuthenticatorKeyPassword{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyPassword)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyPassword.AuthenticatorSimple
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {

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

type NullableAuthenticatorKeyPassword struct {
	value *AuthenticatorKeyPassword
	isSet bool
}

func (v NullableAuthenticatorKeyPassword) Get() *AuthenticatorKeyPassword {
	return v.value
}

func (v *NullableAuthenticatorKeyPassword) Set(val *AuthenticatorKeyPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyPassword(val *AuthenticatorKeyPassword) *NullableAuthenticatorKeyPassword {
	return &NullableAuthenticatorKeyPassword{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

