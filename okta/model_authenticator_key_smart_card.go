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

// AuthenticatorKeySmartCard struct for AuthenticatorKeySmartCard
type AuthenticatorKeySmartCard struct {
	AuthenticatorSimple
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeySmartCard AuthenticatorKeySmartCard

// NewAuthenticatorKeySmartCard instantiates a new AuthenticatorKeySmartCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeySmartCard() *AuthenticatorKeySmartCard {
	this := AuthenticatorKeySmartCard{}
	return &this
}

// NewAuthenticatorKeySmartCardWithDefaults instantiates a new AuthenticatorKeySmartCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeySmartCardWithDefaults() *AuthenticatorKeySmartCard {
	this := AuthenticatorKeySmartCard{}
	return &this
}

func (o AuthenticatorKeySmartCard) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorKeySmartCard) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeySmartCardWithoutEmbeddedStruct struct {
	}

	varAuthenticatorKeySmartCardWithoutEmbeddedStruct := AuthenticatorKeySmartCardWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeySmartCardWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeySmartCard := _AuthenticatorKeySmartCard{}
		*o = AuthenticatorKeySmartCard(varAuthenticatorKeySmartCard)
	} else {
		return err
	}

	varAuthenticatorKeySmartCard := _AuthenticatorKeySmartCard{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeySmartCard)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeySmartCard.AuthenticatorSimple
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

type NullableAuthenticatorKeySmartCard struct {
	value *AuthenticatorKeySmartCard
	isSet bool
}

func (v NullableAuthenticatorKeySmartCard) Get() *AuthenticatorKeySmartCard {
	return v.value
}

func (v *NullableAuthenticatorKeySmartCard) Set(val *AuthenticatorKeySmartCard) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeySmartCard) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeySmartCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeySmartCard(val *AuthenticatorKeySmartCard) *NullableAuthenticatorKeySmartCard {
	return &NullableAuthenticatorKeySmartCard{value: val, isSet: true}
}

func (v NullableAuthenticatorKeySmartCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeySmartCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

