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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the AuthenticatorMethodSimple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorMethodSimple{}

// AuthenticatorMethodSimple struct for AuthenticatorMethodSimple
type AuthenticatorMethodSimple struct {
	AuthenticatorMethodBase
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodSimple AuthenticatorMethodSimple

// NewAuthenticatorMethodSimple instantiates a new AuthenticatorMethodSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodSimple() *AuthenticatorMethodSimple {
	this := AuthenticatorMethodSimple{}
	return &this
}

// NewAuthenticatorMethodSimpleWithDefaults instantiates a new AuthenticatorMethodSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodSimpleWithDefaults() *AuthenticatorMethodSimple {
	this := AuthenticatorMethodSimple{}
	return &this
}

func (o AuthenticatorMethodSimple) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorMethodSimple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorMethodBase, errAuthenticatorMethodBase := json.Marshal(o.AuthenticatorMethodBase)
	if errAuthenticatorMethodBase != nil {
		return map[string]interface{}{}, errAuthenticatorMethodBase
	}
	errAuthenticatorMethodBase = json.Unmarshal([]byte(serializedAuthenticatorMethodBase), &toSerialize)
	if errAuthenticatorMethodBase != nil {
		return map[string]interface{}{}, errAuthenticatorMethodBase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorMethodSimple) UnmarshalJSON(data []byte) (err error) {
	type AuthenticatorMethodSimpleWithoutEmbeddedStruct struct {
	}

	varAuthenticatorMethodSimpleWithoutEmbeddedStruct := AuthenticatorMethodSimpleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAuthenticatorMethodSimpleWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorMethodSimple := _AuthenticatorMethodSimple{}
		*o = AuthenticatorMethodSimple(varAuthenticatorMethodSimple)
	} else {
		return err
	}

	varAuthenticatorMethodSimple := _AuthenticatorMethodSimple{}

	err = json.Unmarshal(data, &varAuthenticatorMethodSimple)
	if err == nil {
		o.AuthenticatorMethodBase = varAuthenticatorMethodSimple.AuthenticatorMethodBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

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
	}

	return err
}

type NullableAuthenticatorMethodSimple struct {
	value *AuthenticatorMethodSimple
	isSet bool
}

func (v NullableAuthenticatorMethodSimple) Get() *AuthenticatorMethodSimple {
	return v.value
}

func (v *NullableAuthenticatorMethodSimple) Set(val *AuthenticatorMethodSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodSimple(val *AuthenticatorMethodSimple) *NullableAuthenticatorMethodSimple {
	return &NullableAuthenticatorMethodSimple{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
