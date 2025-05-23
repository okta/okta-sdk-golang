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
	"fmt"
)


//model_oneof.mustache
// ListDeviceAssurancePolicies200ResponseInner - struct for ListDeviceAssurancePolicies200ResponseInner
type ListDeviceAssurancePolicies200ResponseInner struct {
	DeviceAssuranceAndroidPlatform *DeviceAssuranceAndroidPlatform
	DeviceAssuranceChromeOSPlatform *DeviceAssuranceChromeOSPlatform
	DeviceAssuranceIOSPlatform *DeviceAssuranceIOSPlatform
	DeviceAssuranceMacOSPlatform *DeviceAssuranceMacOSPlatform
	DeviceAssuranceWindowsPlatform *DeviceAssuranceWindowsPlatform
}

// DeviceAssuranceAndroidPlatformAsListDeviceAssurancePolicies200ResponseInner is a convenience function that returns DeviceAssuranceAndroidPlatform wrapped in ListDeviceAssurancePolicies200ResponseInner
func DeviceAssuranceAndroidPlatformAsListDeviceAssurancePolicies200ResponseInner(v *DeviceAssuranceAndroidPlatform) ListDeviceAssurancePolicies200ResponseInner {
	return ListDeviceAssurancePolicies200ResponseInner{
		DeviceAssuranceAndroidPlatform: v,
	}
}

// DeviceAssuranceChromeOSPlatformAsListDeviceAssurancePolicies200ResponseInner is a convenience function that returns DeviceAssuranceChromeOSPlatform wrapped in ListDeviceAssurancePolicies200ResponseInner
func DeviceAssuranceChromeOSPlatformAsListDeviceAssurancePolicies200ResponseInner(v *DeviceAssuranceChromeOSPlatform) ListDeviceAssurancePolicies200ResponseInner {
	return ListDeviceAssurancePolicies200ResponseInner{
		DeviceAssuranceChromeOSPlatform: v,
	}
}

// DeviceAssuranceIOSPlatformAsListDeviceAssurancePolicies200ResponseInner is a convenience function that returns DeviceAssuranceIOSPlatform wrapped in ListDeviceAssurancePolicies200ResponseInner
func DeviceAssuranceIOSPlatformAsListDeviceAssurancePolicies200ResponseInner(v *DeviceAssuranceIOSPlatform) ListDeviceAssurancePolicies200ResponseInner {
	return ListDeviceAssurancePolicies200ResponseInner{
		DeviceAssuranceIOSPlatform: v,
	}
}

// DeviceAssuranceMacOSPlatformAsListDeviceAssurancePolicies200ResponseInner is a convenience function that returns DeviceAssuranceMacOSPlatform wrapped in ListDeviceAssurancePolicies200ResponseInner
func DeviceAssuranceMacOSPlatformAsListDeviceAssurancePolicies200ResponseInner(v *DeviceAssuranceMacOSPlatform) ListDeviceAssurancePolicies200ResponseInner {
	return ListDeviceAssurancePolicies200ResponseInner{
		DeviceAssuranceMacOSPlatform: v,
	}
}

// DeviceAssuranceWindowsPlatformAsListDeviceAssurancePolicies200ResponseInner is a convenience function that returns DeviceAssuranceWindowsPlatform wrapped in ListDeviceAssurancePolicies200ResponseInner
func DeviceAssuranceWindowsPlatformAsListDeviceAssurancePolicies200ResponseInner(v *DeviceAssuranceWindowsPlatform) ListDeviceAssurancePolicies200ResponseInner {
	return ListDeviceAssurancePolicies200ResponseInner{
		DeviceAssuranceWindowsPlatform: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListDeviceAssurancePolicies200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ANDROID'
	if jsonDict["platform"] == "ANDROID" {
		// try to unmarshal JSON data into DeviceAssuranceAndroidPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceAndroidPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceAndroidPlatform, return on the first match
		} else {
			dst.DeviceAssuranceAndroidPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceAndroidPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CHROMEOS'
	if jsonDict["platform"] == "CHROMEOS" {
		// try to unmarshal JSON data into DeviceAssuranceChromeOSPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceChromeOSPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceChromeOSPlatform, return on the first match
		} else {
			dst.DeviceAssuranceChromeOSPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceChromeOSPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DeviceAssuranceAndroidPlatform'
	if jsonDict["platform"] == "DeviceAssuranceAndroidPlatform" {
		// try to unmarshal JSON data into DeviceAssuranceAndroidPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceAndroidPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceAndroidPlatform, return on the first match
		} else {
			dst.DeviceAssuranceAndroidPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceAndroidPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DeviceAssuranceChromeOSPlatform'
	if jsonDict["platform"] == "DeviceAssuranceChromeOSPlatform" {
		// try to unmarshal JSON data into DeviceAssuranceChromeOSPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceChromeOSPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceChromeOSPlatform, return on the first match
		} else {
			dst.DeviceAssuranceChromeOSPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceChromeOSPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DeviceAssuranceIOSPlatform'
	if jsonDict["platform"] == "DeviceAssuranceIOSPlatform" {
		// try to unmarshal JSON data into DeviceAssuranceIOSPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceIOSPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceIOSPlatform, return on the first match
		} else {
			dst.DeviceAssuranceIOSPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceIOSPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DeviceAssuranceMacOSPlatform'
	if jsonDict["platform"] == "DeviceAssuranceMacOSPlatform" {
		// try to unmarshal JSON data into DeviceAssuranceMacOSPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceMacOSPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceMacOSPlatform, return on the first match
		} else {
			dst.DeviceAssuranceMacOSPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceMacOSPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DeviceAssuranceWindowsPlatform'
	if jsonDict["platform"] == "DeviceAssuranceWindowsPlatform" {
		// try to unmarshal JSON data into DeviceAssuranceWindowsPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceWindowsPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceWindowsPlatform, return on the first match
		} else {
			dst.DeviceAssuranceWindowsPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceWindowsPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IOS'
	if jsonDict["platform"] == "IOS" {
		// try to unmarshal JSON data into DeviceAssuranceIOSPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceIOSPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceIOSPlatform, return on the first match
		} else {
			dst.DeviceAssuranceIOSPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceIOSPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MACOS'
	if jsonDict["platform"] == "MACOS" {
		// try to unmarshal JSON data into DeviceAssuranceMacOSPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceMacOSPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceMacOSPlatform, return on the first match
		} else {
			dst.DeviceAssuranceMacOSPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceMacOSPlatform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WINDOWS'
	if jsonDict["platform"] == "WINDOWS" {
		// try to unmarshal JSON data into DeviceAssuranceWindowsPlatform
		err = json.Unmarshal(data, &dst.DeviceAssuranceWindowsPlatform)
		if err == nil {
			return nil // data stored in dst.DeviceAssuranceWindowsPlatform, return on the first match
		} else {
			dst.DeviceAssuranceWindowsPlatform = nil
			return fmt.Errorf("Failed to unmarshal ListDeviceAssurancePolicies200ResponseInner as DeviceAssuranceWindowsPlatform: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListDeviceAssurancePolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.DeviceAssuranceAndroidPlatform != nil {
		return json.Marshal(&src.DeviceAssuranceAndroidPlatform)
	}

	if src.DeviceAssuranceChromeOSPlatform != nil {
		return json.Marshal(&src.DeviceAssuranceChromeOSPlatform)
	}

	if src.DeviceAssuranceIOSPlatform != nil {
		return json.Marshal(&src.DeviceAssuranceIOSPlatform)
	}

	if src.DeviceAssuranceMacOSPlatform != nil {
		return json.Marshal(&src.DeviceAssuranceMacOSPlatform)
	}

	if src.DeviceAssuranceWindowsPlatform != nil {
		return json.Marshal(&src.DeviceAssuranceWindowsPlatform)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListDeviceAssurancePolicies200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DeviceAssuranceAndroidPlatform != nil {
		return obj.DeviceAssuranceAndroidPlatform
	}

	if obj.DeviceAssuranceChromeOSPlatform != nil {
		return obj.DeviceAssuranceChromeOSPlatform
	}

	if obj.DeviceAssuranceIOSPlatform != nil {
		return obj.DeviceAssuranceIOSPlatform
	}

	if obj.DeviceAssuranceMacOSPlatform != nil {
		return obj.DeviceAssuranceMacOSPlatform
	}

	if obj.DeviceAssuranceWindowsPlatform != nil {
		return obj.DeviceAssuranceWindowsPlatform
	}

	// all schemas are nil
	return nil
}

type NullableListDeviceAssurancePolicies200ResponseInner struct {
	value *ListDeviceAssurancePolicies200ResponseInner
	isSet bool
}

func (v NullableListDeviceAssurancePolicies200ResponseInner) Get() *ListDeviceAssurancePolicies200ResponseInner {
	return v.value
}

func (v *NullableListDeviceAssurancePolicies200ResponseInner) Set(val *ListDeviceAssurancePolicies200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeviceAssurancePolicies200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeviceAssurancePolicies200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeviceAssurancePolicies200ResponseInner(val *ListDeviceAssurancePolicies200ResponseInner) *NullableListDeviceAssurancePolicies200ResponseInner {
	return &NullableListDeviceAssurancePolicies200ResponseInner{value: val, isSet: true}
}

func (v NullableListDeviceAssurancePolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeviceAssurancePolicies200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


