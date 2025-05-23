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

// JsonWebKeyRsa struct for JsonWebKeyRsa
type JsonWebKeyRsa struct {
	SchemasJsonWebKey
	// The key exponent of a RSA key
	E *string `json:"e,omitempty"`
	// The modulus of the RSA key
	N *string `json:"n,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JsonWebKeyRsa JsonWebKeyRsa

// NewJsonWebKeyRsa instantiates a new JsonWebKeyRsa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonWebKeyRsa() *JsonWebKeyRsa {
	this := JsonWebKeyRsa{}
	return &this
}

// NewJsonWebKeyRsaWithDefaults instantiates a new JsonWebKeyRsa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonWebKeyRsaWithDefaults() *JsonWebKeyRsa {
	this := JsonWebKeyRsa{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *JsonWebKeyRsa) GetE() string {
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKeyRsa) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *JsonWebKeyRsa) HasE() bool {
	if o != nil && o.E != nil {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *JsonWebKeyRsa) SetE(v string) {
	o.E = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *JsonWebKeyRsa) GetN() string {
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKeyRsa) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *JsonWebKeyRsa) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *JsonWebKeyRsa) SetN(v string) {
	o.N = &v
}

func (o JsonWebKeyRsa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSchemasJsonWebKey, errSchemasJsonWebKey := json.Marshal(o.SchemasJsonWebKey)
	if errSchemasJsonWebKey != nil {
		return []byte{}, errSchemasJsonWebKey
	}
	errSchemasJsonWebKey = json.Unmarshal([]byte(serializedSchemasJsonWebKey), &toSerialize)
	if errSchemasJsonWebKey != nil {
		return []byte{}, errSchemasJsonWebKey
	}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JsonWebKeyRsa) UnmarshalJSON(bytes []byte) (err error) {
	type JsonWebKeyRsaWithoutEmbeddedStruct struct {
		// The key exponent of a RSA key
		E *string `json:"e,omitempty"`
		// The modulus of the RSA key
		N *string `json:"n,omitempty"`
	}

	varJsonWebKeyRsaWithoutEmbeddedStruct := JsonWebKeyRsaWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varJsonWebKeyRsaWithoutEmbeddedStruct)
	if err == nil {
		varJsonWebKeyRsa := _JsonWebKeyRsa{}
		varJsonWebKeyRsa.E = varJsonWebKeyRsaWithoutEmbeddedStruct.E
		varJsonWebKeyRsa.N = varJsonWebKeyRsaWithoutEmbeddedStruct.N
		*o = JsonWebKeyRsa(varJsonWebKeyRsa)
	} else {
		return err
	}

	varJsonWebKeyRsa := _JsonWebKeyRsa{}

	err = json.Unmarshal(bytes, &varJsonWebKeyRsa)
	if err == nil {
		o.SchemasJsonWebKey = varJsonWebKeyRsa.SchemasJsonWebKey
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "n")

		// remove fields from embedded structs
		reflectSchemasJsonWebKey := reflect.ValueOf(o.SchemasJsonWebKey)
		for i := 0; i < reflectSchemasJsonWebKey.Type().NumField(); i++ {
			t := reflectSchemasJsonWebKey.Type().Field(i)

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

type NullableJsonWebKeyRsa struct {
	value *JsonWebKeyRsa
	isSet bool
}

func (v NullableJsonWebKeyRsa) Get() *JsonWebKeyRsa {
	return v.value
}

func (v *NullableJsonWebKeyRsa) Set(val *JsonWebKeyRsa) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonWebKeyRsa) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonWebKeyRsa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonWebKeyRsa(val *JsonWebKeyRsa) *NullableJsonWebKeyRsa {
	return &NullableJsonWebKeyRsa{value: val, isSet: true}
}

func (v NullableJsonWebKeyRsa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonWebKeyRsa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

