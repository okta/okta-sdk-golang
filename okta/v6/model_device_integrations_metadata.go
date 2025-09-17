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
	"fmt"
)

// DeviceIntegrationsMetadata - The metadata of the device integration
type DeviceIntegrationsMetadata struct {
	DeviceIntegrationsMetadataOneOf  *DeviceIntegrationsMetadataOneOf
	DeviceIntegrationsMetadataOneOf1 *DeviceIntegrationsMetadataOneOf1
	DeviceIntegrationsMetadataOneOf2 *DeviceIntegrationsMetadataOneOf2
}

// DeviceIntegrationsMetadataOneOfAsDeviceIntegrationsMetadata is a convenience function that returns DeviceIntegrationsMetadataOneOf wrapped in DeviceIntegrationsMetadata
func DeviceIntegrationsMetadataOneOfAsDeviceIntegrationsMetadata(v *DeviceIntegrationsMetadataOneOf) DeviceIntegrationsMetadata {
	return DeviceIntegrationsMetadata{
		DeviceIntegrationsMetadataOneOf: v,
	}
}

// DeviceIntegrationsMetadataOneOf1AsDeviceIntegrationsMetadata is a convenience function that returns DeviceIntegrationsMetadataOneOf1 wrapped in DeviceIntegrationsMetadata
func DeviceIntegrationsMetadataOneOf1AsDeviceIntegrationsMetadata(v *DeviceIntegrationsMetadataOneOf1) DeviceIntegrationsMetadata {
	return DeviceIntegrationsMetadata{
		DeviceIntegrationsMetadataOneOf1: v,
	}
}

// DeviceIntegrationsMetadataOneOf2AsDeviceIntegrationsMetadata is a convenience function that returns DeviceIntegrationsMetadataOneOf2 wrapped in DeviceIntegrationsMetadata
func DeviceIntegrationsMetadataOneOf2AsDeviceIntegrationsMetadata(v *DeviceIntegrationsMetadataOneOf2) DeviceIntegrationsMetadata {
	return DeviceIntegrationsMetadata{
		DeviceIntegrationsMetadataOneOf2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeviceIntegrationsMetadata) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeviceIntegrationsMetadataOneOf
	err = json.Unmarshal(data, &dst.DeviceIntegrationsMetadataOneOf)
	if err == nil {
		jsonDeviceIntegrationsMetadataOneOf, _ := json.Marshal(dst.DeviceIntegrationsMetadataOneOf)
		if string(jsonDeviceIntegrationsMetadataOneOf) == "{}" { // empty struct
			dst.DeviceIntegrationsMetadataOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DeviceIntegrationsMetadataOneOf = nil
	}

	// try to unmarshal data into DeviceIntegrationsMetadataOneOf1
	err = json.Unmarshal(data, &dst.DeviceIntegrationsMetadataOneOf1)
	if err == nil {
		jsonDeviceIntegrationsMetadataOneOf1, _ := json.Marshal(dst.DeviceIntegrationsMetadataOneOf1)
		if string(jsonDeviceIntegrationsMetadataOneOf1) == "{}" { // empty struct
			dst.DeviceIntegrationsMetadataOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.DeviceIntegrationsMetadataOneOf1 = nil
	}

	// try to unmarshal data into DeviceIntegrationsMetadataOneOf2
	err = json.Unmarshal(data, &dst.DeviceIntegrationsMetadataOneOf2)
	if err == nil {
		jsonDeviceIntegrationsMetadataOneOf2, _ := json.Marshal(dst.DeviceIntegrationsMetadataOneOf2)
		if string(jsonDeviceIntegrationsMetadataOneOf2) == "{}" { // empty struct
			dst.DeviceIntegrationsMetadataOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.DeviceIntegrationsMetadataOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeviceIntegrationsMetadataOneOf = nil
		dst.DeviceIntegrationsMetadataOneOf1 = nil
		dst.DeviceIntegrationsMetadataOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeviceIntegrationsMetadata)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeviceIntegrationsMetadata)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeviceIntegrationsMetadata) MarshalJSON() ([]byte, error) {
	if src.DeviceIntegrationsMetadataOneOf != nil {
		return json.Marshal(&src.DeviceIntegrationsMetadataOneOf)
	}

	if src.DeviceIntegrationsMetadataOneOf1 != nil {
		return json.Marshal(&src.DeviceIntegrationsMetadataOneOf1)
	}

	if src.DeviceIntegrationsMetadataOneOf2 != nil {
		return json.Marshal(&src.DeviceIntegrationsMetadataOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeviceIntegrationsMetadata) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DeviceIntegrationsMetadataOneOf != nil {
		return obj.DeviceIntegrationsMetadataOneOf
	}

	if obj.DeviceIntegrationsMetadataOneOf1 != nil {
		return obj.DeviceIntegrationsMetadataOneOf1
	}

	if obj.DeviceIntegrationsMetadataOneOf2 != nil {
		return obj.DeviceIntegrationsMetadataOneOf2
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj DeviceIntegrationsMetadata) GetActualInstanceValue() interface{} {
	if obj.DeviceIntegrationsMetadataOneOf != nil {
		return *obj.DeviceIntegrationsMetadataOneOf
	}

	if obj.DeviceIntegrationsMetadataOneOf1 != nil {
		return *obj.DeviceIntegrationsMetadataOneOf1
	}

	if obj.DeviceIntegrationsMetadataOneOf2 != nil {
		return *obj.DeviceIntegrationsMetadataOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableDeviceIntegrationsMetadata struct {
	value *DeviceIntegrationsMetadata
	isSet bool
}

func (v NullableDeviceIntegrationsMetadata) Get() *DeviceIntegrationsMetadata {
	return v.value
}

func (v *NullableDeviceIntegrationsMetadata) Set(val *DeviceIntegrationsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntegrationsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntegrationsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntegrationsMetadata(val *DeviceIntegrationsMetadata) *NullableDeviceIntegrationsMetadata {
	return &NullableDeviceIntegrationsMetadata{value: val, isSet: true}
}

func (v NullableDeviceIntegrationsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntegrationsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
