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

// AuthenticatorMethodWithVerifiableProperties struct for AuthenticatorMethodWithVerifiableProperties
type AuthenticatorMethodWithVerifiableProperties struct {
	AuthenticatorMethodBase
	VerifiableProperties []string `json:"verifiableProperties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodWithVerifiableProperties AuthenticatorMethodWithVerifiableProperties

// NewAuthenticatorMethodWithVerifiableProperties instantiates a new AuthenticatorMethodWithVerifiableProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodWithVerifiableProperties() *AuthenticatorMethodWithVerifiableProperties {
	this := AuthenticatorMethodWithVerifiableProperties{}
	return &this
}

// NewAuthenticatorMethodWithVerifiablePropertiesWithDefaults instantiates a new AuthenticatorMethodWithVerifiableProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodWithVerifiablePropertiesWithDefaults() *AuthenticatorMethodWithVerifiableProperties {
	this := AuthenticatorMethodWithVerifiableProperties{}
	return &this
}

// GetVerifiableProperties returns the VerifiableProperties field value if set, zero value otherwise.
func (o *AuthenticatorMethodWithVerifiableProperties) GetVerifiableProperties() []string {
	if o == nil || o.VerifiableProperties == nil {
		var ret []string
		return ret
	}
	return o.VerifiableProperties
}

// GetVerifiablePropertiesOk returns a tuple with the VerifiableProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWithVerifiableProperties) GetVerifiablePropertiesOk() ([]string, bool) {
	if o == nil || o.VerifiableProperties == nil {
		return nil, false
	}
	return o.VerifiableProperties, true
}

// HasVerifiableProperties returns a boolean if a field has been set.
func (o *AuthenticatorMethodWithVerifiableProperties) HasVerifiableProperties() bool {
	if o != nil && o.VerifiableProperties != nil {
		return true
	}

	return false
}

// SetVerifiableProperties gets a reference to the given []string and assigns it to the VerifiableProperties field.
func (o *AuthenticatorMethodWithVerifiableProperties) SetVerifiableProperties(v []string) {
	o.VerifiableProperties = v
}

func (o AuthenticatorMethodWithVerifiableProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorMethodBase, errAuthenticatorMethodBase := json.Marshal(o.AuthenticatorMethodBase)
	if errAuthenticatorMethodBase != nil {
		return []byte{}, errAuthenticatorMethodBase
	}
	errAuthenticatorMethodBase = json.Unmarshal([]byte(serializedAuthenticatorMethodBase), &toSerialize)
	if errAuthenticatorMethodBase != nil {
		return []byte{}, errAuthenticatorMethodBase
	}
	if o.VerifiableProperties != nil {
		toSerialize["verifiableProperties"] = o.VerifiableProperties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodWithVerifiableProperties) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorMethodWithVerifiablePropertiesWithoutEmbeddedStruct struct {
		VerifiableProperties []string `json:"verifiableProperties,omitempty"`
	}

	varAuthenticatorMethodWithVerifiablePropertiesWithoutEmbeddedStruct := AuthenticatorMethodWithVerifiablePropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodWithVerifiablePropertiesWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorMethodWithVerifiableProperties := _AuthenticatorMethodWithVerifiableProperties{}
		varAuthenticatorMethodWithVerifiableProperties.VerifiableProperties = varAuthenticatorMethodWithVerifiablePropertiesWithoutEmbeddedStruct.VerifiableProperties
		*o = AuthenticatorMethodWithVerifiableProperties(varAuthenticatorMethodWithVerifiableProperties)
	} else {
		return err
	}

	varAuthenticatorMethodWithVerifiableProperties := _AuthenticatorMethodWithVerifiableProperties{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodWithVerifiableProperties)
	if err == nil {
		o.AuthenticatorMethodBase = varAuthenticatorMethodWithVerifiableProperties.AuthenticatorMethodBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "verifiableProperties")

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

type NullableAuthenticatorMethodWithVerifiableProperties struct {
	value *AuthenticatorMethodWithVerifiableProperties
	isSet bool
}

func (v NullableAuthenticatorMethodWithVerifiableProperties) Get() *AuthenticatorMethodWithVerifiableProperties {
	return v.value
}

func (v *NullableAuthenticatorMethodWithVerifiableProperties) Set(val *AuthenticatorMethodWithVerifiableProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodWithVerifiableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodWithVerifiableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodWithVerifiableProperties(val *AuthenticatorMethodWithVerifiableProperties) *NullableAuthenticatorMethodWithVerifiableProperties {
	return &NullableAuthenticatorMethodWithVerifiableProperties{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodWithVerifiableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodWithVerifiableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

