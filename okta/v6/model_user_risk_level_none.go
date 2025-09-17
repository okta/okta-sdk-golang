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

// checks if the UserRiskLevelNone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskLevelNone{}

// UserRiskLevelNone struct for UserRiskLevelNone
type UserRiskLevelNone struct {
	UserRiskGetResponse
	AdditionalProperties map[string]interface{}
}

type _UserRiskLevelNone UserRiskLevelNone

// NewUserRiskLevelNone instantiates a new UserRiskLevelNone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskLevelNone() *UserRiskLevelNone {
	this := UserRiskLevelNone{}
	return &this
}

// NewUserRiskLevelNoneWithDefaults instantiates a new UserRiskLevelNone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskLevelNoneWithDefaults() *UserRiskLevelNone {
	this := UserRiskLevelNone{}
	return &this
}

func (o UserRiskLevelNone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskLevelNone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUserRiskGetResponse, errUserRiskGetResponse := json.Marshal(o.UserRiskGetResponse)
	if errUserRiskGetResponse != nil {
		return map[string]interface{}{}, errUserRiskGetResponse
	}
	errUserRiskGetResponse = json.Unmarshal([]byte(serializedUserRiskGetResponse), &toSerialize)
	if errUserRiskGetResponse != nil {
		return map[string]interface{}{}, errUserRiskGetResponse
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRiskLevelNone) UnmarshalJSON(data []byte) (err error) {
	type UserRiskLevelNoneWithoutEmbeddedStruct struct {
	}

	varUserRiskLevelNoneWithoutEmbeddedStruct := UserRiskLevelNoneWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserRiskLevelNoneWithoutEmbeddedStruct)
	if err == nil {
		varUserRiskLevelNone := _UserRiskLevelNone{}
		*o = UserRiskLevelNone(varUserRiskLevelNone)
	} else {
		return err
	}

	varUserRiskLevelNone := _UserRiskLevelNone{}

	err = json.Unmarshal(data, &varUserRiskLevelNone)
	if err == nil {
		o.UserRiskGetResponse = varUserRiskLevelNone.UserRiskGetResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectUserRiskGetResponse := reflect.ValueOf(o.UserRiskGetResponse)
		for i := 0; i < reflectUserRiskGetResponse.Type().NumField(); i++ {
			t := reflectUserRiskGetResponse.Type().Field(i)

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

type NullableUserRiskLevelNone struct {
	value *UserRiskLevelNone
	isSet bool
}

func (v NullableUserRiskLevelNone) Get() *UserRiskLevelNone {
	return v.value
}

func (v *NullableUserRiskLevelNone) Set(val *UserRiskLevelNone) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskLevelNone) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskLevelNone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskLevelNone(val *UserRiskLevelNone) *NullableUserRiskLevelNone {
	return &NullableUserRiskLevelNone{value: val, isSet: true}
}

func (v NullableUserRiskLevelNone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskLevelNone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
