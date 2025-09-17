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

// ExecuteInlineHook200Response - struct for ExecuteInlineHook200Response
type ExecuteInlineHook200Response struct {
	PasswordImportResponse         *PasswordImportResponse
	RegistrationInlineHookResponse *RegistrationInlineHookResponse
	SAMLHookResponse               *SAMLHookResponse
	TelephonyResponse              *TelephonyResponse
	TokenHookResponse              *TokenHookResponse
	UserImportResponse             *UserImportResponse
}

// PasswordImportResponseAsExecuteInlineHook200Response is a convenience function that returns PasswordImportResponse wrapped in ExecuteInlineHook200Response
func PasswordImportResponseAsExecuteInlineHook200Response(v *PasswordImportResponse) ExecuteInlineHook200Response {
	return ExecuteInlineHook200Response{
		PasswordImportResponse: v,
	}
}

// RegistrationInlineHookResponseAsExecuteInlineHook200Response is a convenience function that returns RegistrationInlineHookResponse wrapped in ExecuteInlineHook200Response
func RegistrationInlineHookResponseAsExecuteInlineHook200Response(v *RegistrationInlineHookResponse) ExecuteInlineHook200Response {
	return ExecuteInlineHook200Response{
		RegistrationInlineHookResponse: v,
	}
}

// SAMLHookResponseAsExecuteInlineHook200Response is a convenience function that returns SAMLHookResponse wrapped in ExecuteInlineHook200Response
func SAMLHookResponseAsExecuteInlineHook200Response(v *SAMLHookResponse) ExecuteInlineHook200Response {
	return ExecuteInlineHook200Response{
		SAMLHookResponse: v,
	}
}

// TelephonyResponseAsExecuteInlineHook200Response is a convenience function that returns TelephonyResponse wrapped in ExecuteInlineHook200Response
func TelephonyResponseAsExecuteInlineHook200Response(v *TelephonyResponse) ExecuteInlineHook200Response {
	return ExecuteInlineHook200Response{
		TelephonyResponse: v,
	}
}

// TokenHookResponseAsExecuteInlineHook200Response is a convenience function that returns TokenHookResponse wrapped in ExecuteInlineHook200Response
func TokenHookResponseAsExecuteInlineHook200Response(v *TokenHookResponse) ExecuteInlineHook200Response {
	return ExecuteInlineHook200Response{
		TokenHookResponse: v,
	}
}

// UserImportResponseAsExecuteInlineHook200Response is a convenience function that returns UserImportResponse wrapped in ExecuteInlineHook200Response
func UserImportResponseAsExecuteInlineHook200Response(v *UserImportResponse) ExecuteInlineHook200Response {
	return ExecuteInlineHook200Response{
		UserImportResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExecuteInlineHook200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PasswordImportResponse
	err = json.Unmarshal(data, &dst.PasswordImportResponse)
	if err == nil {
		jsonPasswordImportResponse, _ := json.Marshal(dst.PasswordImportResponse)
		if string(jsonPasswordImportResponse) == "{}" { // empty struct
			dst.PasswordImportResponse = nil
		} else {
			match++
		}
	} else {
		dst.PasswordImportResponse = nil
	}

	// try to unmarshal data into RegistrationInlineHookResponse
	err = json.Unmarshal(data, &dst.RegistrationInlineHookResponse)
	if err == nil {
		jsonRegistrationInlineHookResponse, _ := json.Marshal(dst.RegistrationInlineHookResponse)
		if string(jsonRegistrationInlineHookResponse) == "{}" { // empty struct
			dst.RegistrationInlineHookResponse = nil
		} else {
			match++
		}
	} else {
		dst.RegistrationInlineHookResponse = nil
	}

	// try to unmarshal data into SAMLHookResponse
	err = json.Unmarshal(data, &dst.SAMLHookResponse)
	if err == nil {
		jsonSAMLHookResponse, _ := json.Marshal(dst.SAMLHookResponse)
		if string(jsonSAMLHookResponse) == "{}" { // empty struct
			dst.SAMLHookResponse = nil
		} else {
			match++
		}
	} else {
		dst.SAMLHookResponse = nil
	}

	// try to unmarshal data into TelephonyResponse
	err = json.Unmarshal(data, &dst.TelephonyResponse)
	if err == nil {
		jsonTelephonyResponse, _ := json.Marshal(dst.TelephonyResponse)
		if string(jsonTelephonyResponse) == "{}" { // empty struct
			dst.TelephonyResponse = nil
		} else {
			match++
		}
	} else {
		dst.TelephonyResponse = nil
	}

	// try to unmarshal data into TokenHookResponse
	err = json.Unmarshal(data, &dst.TokenHookResponse)
	if err == nil {
		jsonTokenHookResponse, _ := json.Marshal(dst.TokenHookResponse)
		if string(jsonTokenHookResponse) == "{}" { // empty struct
			dst.TokenHookResponse = nil
		} else {
			match++
		}
	} else {
		dst.TokenHookResponse = nil
	}

	// try to unmarshal data into UserImportResponse
	err = json.Unmarshal(data, &dst.UserImportResponse)
	if err == nil {
		jsonUserImportResponse, _ := json.Marshal(dst.UserImportResponse)
		if string(jsonUserImportResponse) == "{}" { // empty struct
			dst.UserImportResponse = nil
		} else {
			match++
		}
	} else {
		dst.UserImportResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PasswordImportResponse = nil
		dst.RegistrationInlineHookResponse = nil
		dst.SAMLHookResponse = nil
		dst.TelephonyResponse = nil
		dst.TokenHookResponse = nil
		dst.UserImportResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ExecuteInlineHook200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExecuteInlineHook200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExecuteInlineHook200Response) MarshalJSON() ([]byte, error) {
	if src.PasswordImportResponse != nil {
		return json.Marshal(&src.PasswordImportResponse)
	}

	if src.RegistrationInlineHookResponse != nil {
		return json.Marshal(&src.RegistrationInlineHookResponse)
	}

	if src.SAMLHookResponse != nil {
		return json.Marshal(&src.SAMLHookResponse)
	}

	if src.TelephonyResponse != nil {
		return json.Marshal(&src.TelephonyResponse)
	}

	if src.TokenHookResponse != nil {
		return json.Marshal(&src.TokenHookResponse)
	}

	if src.UserImportResponse != nil {
		return json.Marshal(&src.UserImportResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExecuteInlineHook200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PasswordImportResponse != nil {
		return obj.PasswordImportResponse
	}

	if obj.RegistrationInlineHookResponse != nil {
		return obj.RegistrationInlineHookResponse
	}

	if obj.SAMLHookResponse != nil {
		return obj.SAMLHookResponse
	}

	if obj.TelephonyResponse != nil {
		return obj.TelephonyResponse
	}

	if obj.TokenHookResponse != nil {
		return obj.TokenHookResponse
	}

	if obj.UserImportResponse != nil {
		return obj.UserImportResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ExecuteInlineHook200Response) GetActualInstanceValue() interface{} {
	if obj.PasswordImportResponse != nil {
		return *obj.PasswordImportResponse
	}

	if obj.RegistrationInlineHookResponse != nil {
		return *obj.RegistrationInlineHookResponse
	}

	if obj.SAMLHookResponse != nil {
		return *obj.SAMLHookResponse
	}

	if obj.TelephonyResponse != nil {
		return *obj.TelephonyResponse
	}

	if obj.TokenHookResponse != nil {
		return *obj.TokenHookResponse
	}

	if obj.UserImportResponse != nil {
		return *obj.UserImportResponse
	}

	// all schemas are nil
	return nil
}

type NullableExecuteInlineHook200Response struct {
	value *ExecuteInlineHook200Response
	isSet bool
}

func (v NullableExecuteInlineHook200Response) Get() *ExecuteInlineHook200Response {
	return v.value
}

func (v *NullableExecuteInlineHook200Response) Set(val *ExecuteInlineHook200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteInlineHook200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteInlineHook200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteInlineHook200Response(val *ExecuteInlineHook200Response) *NullableExecuteInlineHook200Response {
	return &NullableExecuteInlineHook200Response{value: val, isSet: true}
}

func (v NullableExecuteInlineHook200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteInlineHook200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
