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

// AuthenticatorKeyWebauthn struct for AuthenticatorKeyWebauthn
type AuthenticatorKeyWebauthn struct {
	AuthenticatorSimple
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyWebauthn AuthenticatorKeyWebauthn

// NewAuthenticatorKeyWebauthn instantiates a new AuthenticatorKeyWebauthn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyWebauthn() *AuthenticatorKeyWebauthn {
	this := AuthenticatorKeyWebauthn{}
	return &this
}

// NewAuthenticatorKeyWebauthnWithDefaults instantiates a new AuthenticatorKeyWebauthn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyWebauthnWithDefaults() *AuthenticatorKeyWebauthn {
	this := AuthenticatorKeyWebauthn{}
	return &this
}

func (o AuthenticatorKeyWebauthn) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorKeyWebauthn) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyWebauthnWithoutEmbeddedStruct struct {
	}

	varAuthenticatorKeyWebauthnWithoutEmbeddedStruct := AuthenticatorKeyWebauthnWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyWebauthnWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyWebauthn := _AuthenticatorKeyWebauthn{}
		*o = AuthenticatorKeyWebauthn(varAuthenticatorKeyWebauthn)
	} else {
		return err
	}

	varAuthenticatorKeyWebauthn := _AuthenticatorKeyWebauthn{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyWebauthn)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyWebauthn.AuthenticatorSimple
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

type NullableAuthenticatorKeyWebauthn struct {
	value *AuthenticatorKeyWebauthn
	isSet bool
}

func (v NullableAuthenticatorKeyWebauthn) Get() *AuthenticatorKeyWebauthn {
	return v.value
}

func (v *NullableAuthenticatorKeyWebauthn) Set(val *AuthenticatorKeyWebauthn) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyWebauthn) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyWebauthn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyWebauthn(val *AuthenticatorKeyWebauthn) *NullableAuthenticatorKeyWebauthn {
	return &NullableAuthenticatorKeyWebauthn{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyWebauthn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyWebauthn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

