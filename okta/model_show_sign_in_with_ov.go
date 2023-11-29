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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// ShowSignInWithOV the model 'ShowSignInWithOV'
type ShowSignInWithOV string

// List of ShowSignInWithOV
const (
	SHOWSIGNINWITHOV_ALWAYS ShowSignInWithOV = "ALWAYS"
	SHOWSIGNINWITHOV_NEVER  ShowSignInWithOV = "NEVER"
)

// All allowed values of ShowSignInWithOV enum
var AllowedShowSignInWithOVEnumValues = []ShowSignInWithOV{
	"ALWAYS",
	"NEVER",
}

func (v *ShowSignInWithOV) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShowSignInWithOV(value)
	for _, existing := range AllowedShowSignInWithOVEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShowSignInWithOV", value)
}

// NewShowSignInWithOVFromValue returns a pointer to a valid ShowSignInWithOV
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShowSignInWithOVFromValue(v string) (*ShowSignInWithOV, error) {
	ev := ShowSignInWithOV(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShowSignInWithOV: valid values are %v", v, AllowedShowSignInWithOVEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShowSignInWithOV) IsValid() bool {
	for _, existing := range AllowedShowSignInWithOVEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShowSignInWithOV value
func (v ShowSignInWithOV) Ptr() *ShowSignInWithOV {
	return &v
}

type NullableShowSignInWithOV struct {
	value *ShowSignInWithOV
	isSet bool
}

func (v NullableShowSignInWithOV) Get() *ShowSignInWithOV {
	return v.value
}

func (v *NullableShowSignInWithOV) Set(val *ShowSignInWithOV) {
	v.value = val
	v.isSet = true
}

func (v NullableShowSignInWithOV) IsSet() bool {
	return v.isSet
}

func (v *NullableShowSignInWithOV) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowSignInWithOV(val *ShowSignInWithOV) *NullableShowSignInWithOV {
	return &NullableShowSignInWithOV{value: val, isSet: true}
}

func (v NullableShowSignInWithOV) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowSignInWithOV) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
