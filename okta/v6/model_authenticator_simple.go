/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the AuthenticatorSimple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorSimple{}

// AuthenticatorSimple struct for AuthenticatorSimple
type AuthenticatorSimple struct {
	AuthenticatorBase
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorSimple AuthenticatorSimple

// NewAuthenticatorSimple instantiates a new AuthenticatorSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorSimple() *AuthenticatorSimple {
	this := AuthenticatorSimple{}
	return &this
}

// NewAuthenticatorSimpleWithDefaults instantiates a new AuthenticatorSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorSimpleWithDefaults() *AuthenticatorSimple {
	this := AuthenticatorSimple{}
	return &this
}

func (o AuthenticatorSimple) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorSimple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorBase, errAuthenticatorBase := json.Marshal(o.AuthenticatorBase)
	if errAuthenticatorBase != nil {
		return map[string]interface{}{}, errAuthenticatorBase
	}
	errAuthenticatorBase = json.Unmarshal([]byte(serializedAuthenticatorBase), &toSerialize)
	if errAuthenticatorBase != nil {
		return map[string]interface{}{}, errAuthenticatorBase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorSimple) UnmarshalJSON(data []byte) (err error) {
	type AuthenticatorSimpleWithoutEmbeddedStruct struct {
	}

	varAuthenticatorSimpleWithoutEmbeddedStruct := AuthenticatorSimpleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAuthenticatorSimpleWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorSimple := _AuthenticatorSimple{}
		*o = AuthenticatorSimple(varAuthenticatorSimple)
	} else {
		return err
	}

	varAuthenticatorSimple := _AuthenticatorSimple{}

	err = json.Unmarshal(data, &varAuthenticatorSimple)
	if err == nil {
		o.AuthenticatorBase = varAuthenticatorSimple.AuthenticatorBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectAuthenticatorBase := reflect.ValueOf(o.AuthenticatorBase)
		for i := 0; i < reflectAuthenticatorBase.Type().NumField(); i++ {
			t := reflectAuthenticatorBase.Type().Field(i)

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
	}

	return err
}

type NullableAuthenticatorSimple struct {
	value *AuthenticatorSimple
	isSet bool
}

func (v NullableAuthenticatorSimple) Get() *AuthenticatorSimple {
	return v.value
}

func (v *NullableAuthenticatorSimple) Set(val *AuthenticatorSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorSimple(val *AuthenticatorSimple) *NullableAuthenticatorSimple {
	return &NullableAuthenticatorSimple{value: val, isSet: true}
}

func (v NullableAuthenticatorSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
