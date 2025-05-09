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

// JsonWebKeyEC struct for JsonWebKeyEC
type JsonWebKeyEC struct {
	SchemasJsonWebKey
	// The public x coordinate for the elliptic curve point
	X *string `json:"x,omitempty"`
	// The public y coordinate for the elliptic curve point
	Y *string `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JsonWebKeyEC JsonWebKeyEC

// NewJsonWebKeyEC instantiates a new JsonWebKeyEC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonWebKeyEC() *JsonWebKeyEC {
	this := JsonWebKeyEC{}
	return &this
}

// NewJsonWebKeyECWithDefaults instantiates a new JsonWebKeyEC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonWebKeyECWithDefaults() *JsonWebKeyEC {
	this := JsonWebKeyEC{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *JsonWebKeyEC) GetX() string {
	if o == nil || o.X == nil {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKeyEC) GetXOk() (*string, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *JsonWebKeyEC) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *JsonWebKeyEC) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *JsonWebKeyEC) GetY() string {
	if o == nil || o.Y == nil {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKeyEC) GetYOk() (*string, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *JsonWebKeyEC) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *JsonWebKeyEC) SetY(v string) {
	o.Y = &v
}

func (o JsonWebKeyEC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSchemasJsonWebKey, errSchemasJsonWebKey := json.Marshal(o.SchemasJsonWebKey)
	if errSchemasJsonWebKey != nil {
		return []byte{}, errSchemasJsonWebKey
	}
	errSchemasJsonWebKey = json.Unmarshal([]byte(serializedSchemasJsonWebKey), &toSerialize)
	if errSchemasJsonWebKey != nil {
		return []byte{}, errSchemasJsonWebKey
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JsonWebKeyEC) UnmarshalJSON(bytes []byte) (err error) {
	type JsonWebKeyECWithoutEmbeddedStruct struct {
		// The public x coordinate for the elliptic curve point
		X *string `json:"x,omitempty"`
		// The public y coordinate for the elliptic curve point
		Y *string `json:"y,omitempty"`
	}

	varJsonWebKeyECWithoutEmbeddedStruct := JsonWebKeyECWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varJsonWebKeyECWithoutEmbeddedStruct)
	if err == nil {
		varJsonWebKeyEC := _JsonWebKeyEC{}
		varJsonWebKeyEC.X = varJsonWebKeyECWithoutEmbeddedStruct.X
		varJsonWebKeyEC.Y = varJsonWebKeyECWithoutEmbeddedStruct.Y
		*o = JsonWebKeyEC(varJsonWebKeyEC)
	} else {
		return err
	}

	varJsonWebKeyEC := _JsonWebKeyEC{}

	err = json.Unmarshal(bytes, &varJsonWebKeyEC)
	if err == nil {
		o.SchemasJsonWebKey = varJsonWebKeyEC.SchemasJsonWebKey
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")

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

type NullableJsonWebKeyEC struct {
	value *JsonWebKeyEC
	isSet bool
}

func (v NullableJsonWebKeyEC) Get() *JsonWebKeyEC {
	return v.value
}

func (v *NullableJsonWebKeyEC) Set(val *JsonWebKeyEC) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonWebKeyEC) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonWebKeyEC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonWebKeyEC(val *JsonWebKeyEC) *NullableJsonWebKeyEC {
	return &NullableJsonWebKeyEC{value: val, isSet: true}
}

func (v NullableJsonWebKeyEC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonWebKeyEC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

