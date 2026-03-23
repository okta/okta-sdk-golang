/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// CustomTelephonyProviderSettingsCall - Method for making voice calls. Choose one method for making voice calls based on your telephony provider setup.
type CustomTelephonyProviderSettingsCall struct {
	CustomTelephonyProviderSettingsTelesignServiceCall *CustomTelephonyProviderSettingsTelesignServiceCall
	CustomTelephonyProviderSettingsTwilioCallerId      *CustomTelephonyProviderSettingsTwilioCallerId
	CustomTelephonyProviderSettingsTwilioPhoneNumber   *CustomTelephonyProviderSettingsTwilioPhoneNumber
	CustomTelephonyProviderSettingsTwilioVerify        *CustomTelephonyProviderSettingsTwilioVerify
}

// CustomTelephonyProviderSettingsTelesignServiceCallAsCustomTelephonyProviderSettingsCall is a convenience function that returns CustomTelephonyProviderSettingsTelesignServiceCall wrapped in CustomTelephonyProviderSettingsCall
func CustomTelephonyProviderSettingsTelesignServiceCallAsCustomTelephonyProviderSettingsCall(v *CustomTelephonyProviderSettingsTelesignServiceCall) CustomTelephonyProviderSettingsCall {
	return CustomTelephonyProviderSettingsCall{
		CustomTelephonyProviderSettingsTelesignServiceCall: v,
	}
}

// CustomTelephonyProviderSettingsTwilioCallerIdAsCustomTelephonyProviderSettingsCall is a convenience function that returns CustomTelephonyProviderSettingsTwilioCallerId wrapped in CustomTelephonyProviderSettingsCall
func CustomTelephonyProviderSettingsTwilioCallerIdAsCustomTelephonyProviderSettingsCall(v *CustomTelephonyProviderSettingsTwilioCallerId) CustomTelephonyProviderSettingsCall {
	return CustomTelephonyProviderSettingsCall{
		CustomTelephonyProviderSettingsTwilioCallerId: v,
	}
}

// CustomTelephonyProviderSettingsTwilioPhoneNumberAsCustomTelephonyProviderSettingsCall is a convenience function that returns CustomTelephonyProviderSettingsTwilioPhoneNumber wrapped in CustomTelephonyProviderSettingsCall
func CustomTelephonyProviderSettingsTwilioPhoneNumberAsCustomTelephonyProviderSettingsCall(v *CustomTelephonyProviderSettingsTwilioPhoneNumber) CustomTelephonyProviderSettingsCall {
	return CustomTelephonyProviderSettingsCall{
		CustomTelephonyProviderSettingsTwilioPhoneNumber: v,
	}
}

// CustomTelephonyProviderSettingsTwilioVerifyAsCustomTelephonyProviderSettingsCall is a convenience function that returns CustomTelephonyProviderSettingsTwilioVerify wrapped in CustomTelephonyProviderSettingsCall
func CustomTelephonyProviderSettingsTwilioVerifyAsCustomTelephonyProviderSettingsCall(v *CustomTelephonyProviderSettingsTwilioVerify) CustomTelephonyProviderSettingsCall {
	return CustomTelephonyProviderSettingsCall{
		CustomTelephonyProviderSettingsTwilioVerify: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomTelephonyProviderSettingsCall) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomTelephonyProviderSettingsTelesignServiceCall
	err = json.Unmarshal(data, &dst.CustomTelephonyProviderSettingsTelesignServiceCall)
	if err == nil {
		jsonCustomTelephonyProviderSettingsTelesignServiceCall, _ := json.Marshal(dst.CustomTelephonyProviderSettingsTelesignServiceCall)
		if string(jsonCustomTelephonyProviderSettingsTelesignServiceCall) == "{}" { // empty struct
			dst.CustomTelephonyProviderSettingsTelesignServiceCall = nil
		} else {
			match++
		}
	} else {
		dst.CustomTelephonyProviderSettingsTelesignServiceCall = nil
	}

	// try to unmarshal data into CustomTelephonyProviderSettingsTwilioCallerId
	err = json.Unmarshal(data, &dst.CustomTelephonyProviderSettingsTwilioCallerId)
	if err == nil {
		jsonCustomTelephonyProviderSettingsTwilioCallerId, _ := json.Marshal(dst.CustomTelephonyProviderSettingsTwilioCallerId)
		if string(jsonCustomTelephonyProviderSettingsTwilioCallerId) == "{}" { // empty struct
			dst.CustomTelephonyProviderSettingsTwilioCallerId = nil
		} else {
			match++
		}
	} else {
		dst.CustomTelephonyProviderSettingsTwilioCallerId = nil
	}

	// try to unmarshal data into CustomTelephonyProviderSettingsTwilioPhoneNumber
	err = json.Unmarshal(data, &dst.CustomTelephonyProviderSettingsTwilioPhoneNumber)
	if err == nil {
		jsonCustomTelephonyProviderSettingsTwilioPhoneNumber, _ := json.Marshal(dst.CustomTelephonyProviderSettingsTwilioPhoneNumber)
		if string(jsonCustomTelephonyProviderSettingsTwilioPhoneNumber) == "{}" { // empty struct
			dst.CustomTelephonyProviderSettingsTwilioPhoneNumber = nil
		} else {
			match++
		}
	} else {
		dst.CustomTelephonyProviderSettingsTwilioPhoneNumber = nil
	}

	// try to unmarshal data into CustomTelephonyProviderSettingsTwilioVerify
	err = json.Unmarshal(data, &dst.CustomTelephonyProviderSettingsTwilioVerify)
	if err == nil {
		jsonCustomTelephonyProviderSettingsTwilioVerify, _ := json.Marshal(dst.CustomTelephonyProviderSettingsTwilioVerify)
		if string(jsonCustomTelephonyProviderSettingsTwilioVerify) == "{}" { // empty struct
			dst.CustomTelephonyProviderSettingsTwilioVerify = nil
		} else {
			match++
		}
	} else {
		dst.CustomTelephonyProviderSettingsTwilioVerify = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomTelephonyProviderSettingsTelesignServiceCall = nil
		dst.CustomTelephonyProviderSettingsTwilioCallerId = nil
		dst.CustomTelephonyProviderSettingsTwilioPhoneNumber = nil
		dst.CustomTelephonyProviderSettingsTwilioVerify = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CustomTelephonyProviderSettingsCall)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CustomTelephonyProviderSettingsCall)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomTelephonyProviderSettingsCall) MarshalJSON() ([]byte, error) {
	if src.CustomTelephonyProviderSettingsTelesignServiceCall != nil {
		return json.Marshal(&src.CustomTelephonyProviderSettingsTelesignServiceCall)
	}

	if src.CustomTelephonyProviderSettingsTwilioCallerId != nil {
		return json.Marshal(&src.CustomTelephonyProviderSettingsTwilioCallerId)
	}

	if src.CustomTelephonyProviderSettingsTwilioPhoneNumber != nil {
		return json.Marshal(&src.CustomTelephonyProviderSettingsTwilioPhoneNumber)
	}

	if src.CustomTelephonyProviderSettingsTwilioVerify != nil {
		return json.Marshal(&src.CustomTelephonyProviderSettingsTwilioVerify)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CustomTelephonyProviderSettingsCall) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomTelephonyProviderSettingsTelesignServiceCall != nil {
		return obj.CustomTelephonyProviderSettingsTelesignServiceCall
	}

	if obj.CustomTelephonyProviderSettingsTwilioCallerId != nil {
		return obj.CustomTelephonyProviderSettingsTwilioCallerId
	}

	if obj.CustomTelephonyProviderSettingsTwilioPhoneNumber != nil {
		return obj.CustomTelephonyProviderSettingsTwilioPhoneNumber
	}

	if obj.CustomTelephonyProviderSettingsTwilioVerify != nil {
		return obj.CustomTelephonyProviderSettingsTwilioVerify
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CustomTelephonyProviderSettingsCall) GetActualInstanceValue() interface{} {
	if obj.CustomTelephonyProviderSettingsTelesignServiceCall != nil {
		return *obj.CustomTelephonyProviderSettingsTelesignServiceCall
	}

	if obj.CustomTelephonyProviderSettingsTwilioCallerId != nil {
		return *obj.CustomTelephonyProviderSettingsTwilioCallerId
	}

	if obj.CustomTelephonyProviderSettingsTwilioPhoneNumber != nil {
		return *obj.CustomTelephonyProviderSettingsTwilioPhoneNumber
	}

	if obj.CustomTelephonyProviderSettingsTwilioVerify != nil {
		return *obj.CustomTelephonyProviderSettingsTwilioVerify
	}

	// all schemas are nil
	return nil
}

type NullableCustomTelephonyProviderSettingsCall struct {
	value *CustomTelephonyProviderSettingsCall
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsCall) Get() *CustomTelephonyProviderSettingsCall {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsCall) Set(val *CustomTelephonyProviderSettingsCall) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsCall) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsCall(val *CustomTelephonyProviderSettingsCall) *NullableCustomTelephonyProviderSettingsCall {
	return &NullableCustomTelephonyProviderSettingsCall{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
