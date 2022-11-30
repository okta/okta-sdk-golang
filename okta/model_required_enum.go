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

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// RequiredEnum the model 'RequiredEnum'
type RequiredEnum string

// List of RequiredEnum
const (
	REQUIREDENUM_ALWAYS         RequiredEnum = "ALWAYS"
	REQUIREDENUM_HIGH_RISK_ONLY RequiredEnum = "HIGH_RISK_ONLY"
	REQUIREDENUM_NEVER          RequiredEnum = "NEVER"
)

// All allowed values of RequiredEnum enum
var AllowedRequiredEnumEnumValues = []RequiredEnum{
	"ALWAYS",
	"HIGH_RISK_ONLY",
	"NEVER",
}

func (v *RequiredEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequiredEnum(value)
	for _, existing := range AllowedRequiredEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequiredEnum", value)
}

// NewRequiredEnumFromValue returns a pointer to a valid RequiredEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequiredEnumFromValue(v string) (*RequiredEnum, error) {
	ev := RequiredEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequiredEnum: valid values are %v", v, AllowedRequiredEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequiredEnum) IsValid() bool {
	for _, existing := range AllowedRequiredEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequiredEnum value
func (v RequiredEnum) Ptr() *RequiredEnum {
	return &v
}

type NullableRequiredEnum struct {
	value *RequiredEnum
	isSet bool
}

func (v NullableRequiredEnum) Get() *RequiredEnum {
	return v.value
}

func (v *NullableRequiredEnum) Set(val *RequiredEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredEnum(val *RequiredEnum) *NullableRequiredEnum {
	return &NullableRequiredEnum{value: val, isSet: true}
}

func (v NullableRequiredEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequiredEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}