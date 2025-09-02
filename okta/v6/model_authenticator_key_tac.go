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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AuthenticatorKeyTac struct for AuthenticatorKeyTac
type AuthenticatorKeyTac struct {
	AuthenticatorSimple
	Provider *AuthenticatorKeyTacAllOfProvider `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyTac AuthenticatorKeyTac

// NewAuthenticatorKeyTac instantiates a new AuthenticatorKeyTac object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyTac() *AuthenticatorKeyTac {
	this := AuthenticatorKeyTac{}
	return &this
}

// NewAuthenticatorKeyTacWithDefaults instantiates a new AuthenticatorKeyTac object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyTacWithDefaults() *AuthenticatorKeyTac {
	this := AuthenticatorKeyTac{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AuthenticatorKeyTac) GetProvider() AuthenticatorKeyTacAllOfProvider {
	if o == nil || o.Provider == nil {
		var ret AuthenticatorKeyTacAllOfProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTac) GetProviderOk() (*AuthenticatorKeyTacAllOfProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AuthenticatorKeyTac) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AuthenticatorKeyTacAllOfProvider and assigns it to the Provider field.
func (o *AuthenticatorKeyTac) SetProvider(v AuthenticatorKeyTacAllOfProvider) {
	o.Provider = &v
}

func (o AuthenticatorKeyTac) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorSimple, errAuthenticatorSimple := json.Marshal(o.AuthenticatorSimple)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	errAuthenticatorSimple = json.Unmarshal([]byte(serializedAuthenticatorSimple), &toSerialize)
	if errAuthenticatorSimple != nil {
		return []byte{}, errAuthenticatorSimple
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyTac) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorKeyTacWithoutEmbeddedStruct struct {
		Provider *AuthenticatorKeyTacAllOfProvider `json:"provider,omitempty"`
	}

	varAuthenticatorKeyTacWithoutEmbeddedStruct := AuthenticatorKeyTacWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyTacWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorKeyTac := _AuthenticatorKeyTac{}
		varAuthenticatorKeyTac.Provider = varAuthenticatorKeyTacWithoutEmbeddedStruct.Provider
		*o = AuthenticatorKeyTac(varAuthenticatorKeyTac)
	} else {
		return err
	}

	varAuthenticatorKeyTac := _AuthenticatorKeyTac{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyTac)
	if err == nil {
		o.AuthenticatorSimple = varAuthenticatorKeyTac.AuthenticatorSimple
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyTac struct {
	value *AuthenticatorKeyTac
	isSet bool
}

func (v NullableAuthenticatorKeyTac) Get() *AuthenticatorKeyTac {
	return v.value
}

func (v *NullableAuthenticatorKeyTac) Set(val *AuthenticatorKeyTac) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyTac) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyTac) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyTac(val *AuthenticatorKeyTac) *NullableAuthenticatorKeyTac {
	return &NullableAuthenticatorKeyTac{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyTac) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyTac) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

