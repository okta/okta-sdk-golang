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

// AuthenticatorKeyYubikey struct for AuthenticatorKeyYubikey
type AuthenticatorKeyYubikey struct {
	AuthenticatorSimple
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyYubikey AuthenticatorKeyYubikey

// NewAuthenticatorKeyYubikey instantiates a new AuthenticatorKeyYubikey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyYubikey() *AuthenticatorKeyYubikey {
	this := AuthenticatorKeyYubikey{}
	return &this
}

// NewAuthenticatorKeyYubikeyWithDefaults instantiates a new AuthenticatorKeyYubikey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyYubikeyWithDefaults() *AuthenticatorKeyYubikey {
	this := AuthenticatorKeyYubikey{}
	return &this
}

func (o AuthenticatorKeyYubikey) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorKeyYubikey) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyYubikeyWithoutEmbeddedStruct struct {
	}

	varAuthenticatorKeyYubikeyWithoutEmbeddedStruct := AuthenticatorKeyYubikeyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyYubikeyWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyYubikey := _AuthenticatorKeyYubikey{}
		*o = AuthenticatorKeyYubikey(varAuthenticatorKeyYubikey)
	} else {
		return err
	}

	varAuthenticatorKeyYubikey := _AuthenticatorKeyYubikey{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyYubikey)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyYubikey.AuthenticatorSimple
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

type NullableAuthenticatorKeyYubikey struct {
	value *AuthenticatorKeyYubikey
	isSet bool
}

func (v NullableAuthenticatorKeyYubikey) Get() *AuthenticatorKeyYubikey {
	return v.value
}

func (v *NullableAuthenticatorKeyYubikey) Set(val *AuthenticatorKeyYubikey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyYubikey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyYubikey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyYubikey(val *AuthenticatorKeyYubikey) *NullableAuthenticatorKeyYubikey {
	return &NullableAuthenticatorKeyYubikey{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyYubikey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyYubikey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

