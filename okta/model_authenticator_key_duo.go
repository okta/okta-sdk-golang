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

// checks if the AuthenticatorKeyDuo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorKeyDuo{}

// AuthenticatorKeyDuo struct for AuthenticatorKeyDuo
type AuthenticatorKeyDuo struct {
	AuthenticatorSimple
	Provider             *AuthenticatorKeyDuoAllOfProvider `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyDuo AuthenticatorKeyDuo

// NewAuthenticatorKeyDuo instantiates a new AuthenticatorKeyDuo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyDuo() *AuthenticatorKeyDuo {
	this := AuthenticatorKeyDuo{}
	return &this
}

// NewAuthenticatorKeyDuoWithDefaults instantiates a new AuthenticatorKeyDuo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyDuoWithDefaults() *AuthenticatorKeyDuo {
	this := AuthenticatorKeyDuo{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AuthenticatorKeyDuo) GetProvider() AuthenticatorKeyDuoAllOfProvider {
	if o == nil || IsNil(o.Provider) {
		var ret AuthenticatorKeyDuoAllOfProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyDuo) GetProviderOk() (*AuthenticatorKeyDuoAllOfProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthenticatorKeyDuo) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AuthenticatorKeyDuoAllOfProvider and assigns it to the Provider field.
func (o *AuthenticatorKeyDuo) SetProvider(v AuthenticatorKeyDuoAllOfProvider) {
	o.Provider = &v
}

func (o AuthenticatorKeyDuo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorKeyDuo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorSimple, errAuthenticatorSimple := json.Marshal(o.AuthenticatorSimple)
	if errAuthenticatorSimple != nil {
		return map[string]interface{}{}, errAuthenticatorSimple
	}
	errAuthenticatorSimple = json.Unmarshal([]byte(serializedAuthenticatorSimple), &toSerialize)
	if errAuthenticatorSimple != nil {
		return map[string]interface{}{}, errAuthenticatorSimple
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorKeyDuo) UnmarshalJSON(data []byte) (err error) {
	type AuthenticatorKeyDuoWithoutEmbeddedStruct struct {
		Provider *AuthenticatorKeyDuoAllOfProvider `json:"provider,omitempty"`
	}

	varAuthenticatorKeyDuoWithoutEmbeddedStruct := AuthenticatorKeyDuoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAuthenticatorKeyDuoWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyDuo := _AuthenticatorKeyDuo{}
		varAuthenticatorKeyDuo.Provider = varAuthenticatorKeyDuoWithoutEmbeddedStruct.Provider
		*o = AuthenticatorKeyDuo(varAuthenticatorKeyDuo)
	} else {
		return err
	}

	varAuthenticatorKeyDuo := _AuthenticatorKeyDuo{}

	err = json.Unmarshal(data, &varAuthenticatorKeyDuo)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyDuo.AuthenticatorSimple
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider")

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
	}

	return err
}

type NullableAuthenticatorKeyDuo struct {
	value *AuthenticatorKeyDuo
	isSet bool
}

func (v NullableAuthenticatorKeyDuo) Get() *AuthenticatorKeyDuo {
	return v.value
}

func (v *NullableAuthenticatorKeyDuo) Set(val *AuthenticatorKeyDuo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyDuo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyDuo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyDuo(val *AuthenticatorKeyDuo) *NullableAuthenticatorKeyDuo {
	return &NullableAuthenticatorKeyDuo{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyDuo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyDuo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
