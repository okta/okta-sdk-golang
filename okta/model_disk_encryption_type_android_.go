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

// DiskEncryptionTypeAndroid the model 'DiskEncryptionTypeAndroid'
type DiskEncryptionTypeAndroid string

// List of DiskEncryptionTypeAndroid
const (
	DISKENCRYPTIONTYPEANDROID_FULL DiskEncryptionTypeAndroid = "FULL"
	DISKENCRYPTIONTYPEANDROID_USER DiskEncryptionTypeAndroid = "USER"
)

// All allowed values of DiskEncryptionTypeAndroid enum
var AllowedDiskEncryptionTypeAndroidEnumValues = []DiskEncryptionTypeAndroid{
	"FULL",
	"USER",
}

func (v *DiskEncryptionTypeAndroid) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiskEncryptionTypeAndroid(value)
	for _, existing := range AllowedDiskEncryptionTypeAndroidEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiskEncryptionTypeAndroid", value)
}

// NewDiskEncryptionTypeAndroidFromValue returns a pointer to a valid DiskEncryptionTypeAndroid
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiskEncryptionTypeAndroidFromValue(v string) (*DiskEncryptionTypeAndroid, error) {
	ev := DiskEncryptionTypeAndroid(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiskEncryptionTypeAndroid: valid values are %v", v, AllowedDiskEncryptionTypeAndroidEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiskEncryptionTypeAndroid) IsValid() bool {
	for _, existing := range AllowedDiskEncryptionTypeAndroidEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiskEncryptionTypeAndroid value
func (v DiskEncryptionTypeAndroid) Ptr() *DiskEncryptionTypeAndroid {
	return &v
}

type NullableDiskEncryptionTypeAndroid struct {
	value *DiskEncryptionTypeAndroid
	isSet bool
}

func (v NullableDiskEncryptionTypeAndroid) Get() *DiskEncryptionTypeAndroid {
	return v.value
}

func (v *NullableDiskEncryptionTypeAndroid) Set(val *DiskEncryptionTypeAndroid) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskEncryptionTypeAndroid) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskEncryptionTypeAndroid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskEncryptionTypeAndroid(val *DiskEncryptionTypeAndroid) *NullableDiskEncryptionTypeAndroid {
	return &NullableDiskEncryptionTypeAndroid{value: val, isSet: true}
}

func (v NullableDiskEncryptionTypeAndroid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskEncryptionTypeAndroid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
