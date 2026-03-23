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

// CustomTelephonyProviderSettingsSms - Method for sending SMS messages. Choose one method for sending SMS messages based on your telephony provider setup.
type CustomTelephonyProviderSettingsSms struct {
	CustomTelephonyProviderSettingsTelesignServiceSms     *CustomTelephonyProviderSettingsTelesignServiceSms
	CustomTelephonyProviderSettingsTwilioMessagingService *CustomTelephonyProviderSettingsTwilioMessagingService
	CustomTelephonyProviderSettingsTwilioPhoneNumber      *CustomTelephonyProviderSettingsTwilioPhoneNumber
	CustomTelephonyProviderSettingsTwilioVerify           *CustomTelephonyProviderSettingsTwilioVerify
}

// CustomTelephonyProviderSettingsTelesignServiceSmsAsCustomTelephonyProviderSettingsSms is a convenience function that returns CustomTelephonyProviderSettingsTelesignServiceSms wrapped in CustomTelephonyProviderSettingsSms
func CustomTelephonyProviderSettingsTelesignServiceSmsAsCustomTelephonyProviderSettingsSms(v *CustomTelephonyProviderSettingsTelesignServiceSms) CustomTelephonyProviderSettingsSms {
	return CustomTelephonyProviderSettingsSms{
		CustomTelephonyProviderSettingsTelesignServiceSms: v,
	}
}

// CustomTelephonyProviderSettingsTwilioMessagingServiceAsCustomTelephonyProviderSettingsSms is a convenience function that returns CustomTelephonyProviderSettingsTwilioMessagingService wrapped in CustomTelephonyProviderSettingsSms
func CustomTelephonyProviderSettingsTwilioMessagingServiceAsCustomTelephonyProviderSettingsSms(v *CustomTelephonyProviderSettingsTwilioMessagingService) CustomTelephonyProviderSettingsSms {
	return CustomTelephonyProviderSettingsSms{
		CustomTelephonyProviderSettingsTwilioMessagingService: v,
	}
}

// CustomTelephonyProviderSettingsTwilioPhoneNumberAsCustomTelephonyProviderSettingsSms is a convenience function that returns CustomTelephonyProviderSettingsTwilioPhoneNumber wrapped in CustomTelephonyProviderSettingsSms
func CustomTelephonyProviderSettingsTwilioPhoneNumberAsCustomTelephonyProviderSettingsSms(v *CustomTelephonyProviderSettingsTwilioPhoneNumber) CustomTelephonyProviderSettingsSms {
	return CustomTelephonyProviderSettingsSms{
		CustomTelephonyProviderSettingsTwilioPhoneNumber: v,
	}
}

// CustomTelephonyProviderSettingsTwilioVerifyAsCustomTelephonyProviderSettingsSms is a convenience function that returns CustomTelephonyProviderSettingsTwilioVerify wrapped in CustomTelephonyProviderSettingsSms
func CustomTelephonyProviderSettingsTwilioVerifyAsCustomTelephonyProviderSettingsSms(v *CustomTelephonyProviderSettingsTwilioVerify) CustomTelephonyProviderSettingsSms {
	return CustomTelephonyProviderSettingsSms{
		CustomTelephonyProviderSettingsTwilioVerify: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomTelephonyProviderSettingsSms) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomTelephonyProviderSettingsTelesignServiceSms
	err = json.Unmarshal(data, &dst.CustomTelephonyProviderSettingsTelesignServiceSms)
	if err == nil {
		jsonCustomTelephonyProviderSettingsTelesignServiceSms, _ := json.Marshal(dst.CustomTelephonyProviderSettingsTelesignServiceSms)
		if string(jsonCustomTelephonyProviderSettingsTelesignServiceSms) == "{}" { // empty struct
			dst.CustomTelephonyProviderSettingsTelesignServiceSms = nil
		} else {
			match++
		}
	} else {
		dst.CustomTelephonyProviderSettingsTelesignServiceSms = nil
	}

	// try to unmarshal data into CustomTelephonyProviderSettingsTwilioMessagingService
	err = json.Unmarshal(data, &dst.CustomTelephonyProviderSettingsTwilioMessagingService)
	if err == nil {
		jsonCustomTelephonyProviderSettingsTwilioMessagingService, _ := json.Marshal(dst.CustomTelephonyProviderSettingsTwilioMessagingService)
		if string(jsonCustomTelephonyProviderSettingsTwilioMessagingService) == "{}" { // empty struct
			dst.CustomTelephonyProviderSettingsTwilioMessagingService = nil
		} else {
			match++
		}
	} else {
		dst.CustomTelephonyProviderSettingsTwilioMessagingService = nil
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
		dst.CustomTelephonyProviderSettingsTelesignServiceSms = nil
		dst.CustomTelephonyProviderSettingsTwilioMessagingService = nil
		dst.CustomTelephonyProviderSettingsTwilioPhoneNumber = nil
		dst.CustomTelephonyProviderSettingsTwilioVerify = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CustomTelephonyProviderSettingsSms)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CustomTelephonyProviderSettingsSms)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomTelephonyProviderSettingsSms) MarshalJSON() ([]byte, error) {
	if src.CustomTelephonyProviderSettingsTelesignServiceSms != nil {
		return json.Marshal(&src.CustomTelephonyProviderSettingsTelesignServiceSms)
	}

	if src.CustomTelephonyProviderSettingsTwilioMessagingService != nil {
		return json.Marshal(&src.CustomTelephonyProviderSettingsTwilioMessagingService)
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
func (obj *CustomTelephonyProviderSettingsSms) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomTelephonyProviderSettingsTelesignServiceSms != nil {
		return obj.CustomTelephonyProviderSettingsTelesignServiceSms
	}

	if obj.CustomTelephonyProviderSettingsTwilioMessagingService != nil {
		return obj.CustomTelephonyProviderSettingsTwilioMessagingService
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
func (obj CustomTelephonyProviderSettingsSms) GetActualInstanceValue() interface{} {
	if obj.CustomTelephonyProviderSettingsTelesignServiceSms != nil {
		return *obj.CustomTelephonyProviderSettingsTelesignServiceSms
	}

	if obj.CustomTelephonyProviderSettingsTwilioMessagingService != nil {
		return *obj.CustomTelephonyProviderSettingsTwilioMessagingService
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

type NullableCustomTelephonyProviderSettingsSms struct {
	value *CustomTelephonyProviderSettingsSms
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsSms) Get() *CustomTelephonyProviderSettingsSms {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsSms) Set(val *CustomTelephonyProviderSettingsSms) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsSms) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsSms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsSms(val *CustomTelephonyProviderSettingsSms) *NullableCustomTelephonyProviderSettingsSms {
	return &NullableCustomTelephonyProviderSettingsSms{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsSms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsSms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
