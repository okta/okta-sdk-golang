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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// ExecuteInlineHookRequest - struct for ExecuteInlineHookRequest
type ExecuteInlineHookRequest struct {
	PasswordImportRequestExecute  *PasswordImportRequestExecute
	RegistrationInlineHookRequest *RegistrationInlineHookRequest
	SAMLPayloadExecute            *SAMLPayloadExecute
	TelephonyRequestExecute       *TelephonyRequestExecute
	TokenRequest                  *TokenRequest
	UserImportRequestExecute      *UserImportRequestExecute
}

// PasswordImportRequestExecuteAsExecuteInlineHookRequest is a convenience function that returns PasswordImportRequestExecute wrapped in ExecuteInlineHookRequest
func PasswordImportRequestExecuteAsExecuteInlineHookRequest(v *PasswordImportRequestExecute) ExecuteInlineHookRequest {
	return ExecuteInlineHookRequest{
		PasswordImportRequestExecute: v,
	}
}

// RegistrationInlineHookRequestAsExecuteInlineHookRequest is a convenience function that returns RegistrationInlineHookRequest wrapped in ExecuteInlineHookRequest
func RegistrationInlineHookRequestAsExecuteInlineHookRequest(v *RegistrationInlineHookRequest) ExecuteInlineHookRequest {
	return ExecuteInlineHookRequest{
		RegistrationInlineHookRequest: v,
	}
}

// SAMLPayloadExecuteAsExecuteInlineHookRequest is a convenience function that returns SAMLPayloadExecute wrapped in ExecuteInlineHookRequest
func SAMLPayloadExecuteAsExecuteInlineHookRequest(v *SAMLPayloadExecute) ExecuteInlineHookRequest {
	return ExecuteInlineHookRequest{
		SAMLPayloadExecute: v,
	}
}

// TelephonyRequestExecuteAsExecuteInlineHookRequest is a convenience function that returns TelephonyRequestExecute wrapped in ExecuteInlineHookRequest
func TelephonyRequestExecuteAsExecuteInlineHookRequest(v *TelephonyRequestExecute) ExecuteInlineHookRequest {
	return ExecuteInlineHookRequest{
		TelephonyRequestExecute: v,
	}
}

// TokenRequestAsExecuteInlineHookRequest is a convenience function that returns TokenRequest wrapped in ExecuteInlineHookRequest
func TokenRequestAsExecuteInlineHookRequest(v *TokenRequest) ExecuteInlineHookRequest {
	return ExecuteInlineHookRequest{
		TokenRequest: v,
	}
}

// UserImportRequestExecuteAsExecuteInlineHookRequest is a convenience function that returns UserImportRequestExecute wrapped in ExecuteInlineHookRequest
func UserImportRequestExecuteAsExecuteInlineHookRequest(v *UserImportRequestExecute) ExecuteInlineHookRequest {
	return ExecuteInlineHookRequest{
		UserImportRequestExecute: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExecuteInlineHookRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PasswordImportRequestExecute
	err = json.Unmarshal(data, &dst.PasswordImportRequestExecute)
	if err == nil {
		jsonPasswordImportRequestExecute, _ := json.Marshal(dst.PasswordImportRequestExecute)
		if string(jsonPasswordImportRequestExecute) == "{}" { // empty struct
			dst.PasswordImportRequestExecute = nil
		} else {
			match++
		}
	} else {
		dst.PasswordImportRequestExecute = nil
	}

	// try to unmarshal data into RegistrationInlineHookRequest
	err = json.Unmarshal(data, &dst.RegistrationInlineHookRequest)
	if err == nil {
		jsonRegistrationInlineHookRequest, _ := json.Marshal(dst.RegistrationInlineHookRequest)
		if string(jsonRegistrationInlineHookRequest) == "{}" { // empty struct
			dst.RegistrationInlineHookRequest = nil
		} else {
			match++
		}
	} else {
		dst.RegistrationInlineHookRequest = nil
	}

	// try to unmarshal data into SAMLPayloadExecute
	err = json.Unmarshal(data, &dst.SAMLPayloadExecute)
	if err == nil {
		jsonSAMLPayloadExecute, _ := json.Marshal(dst.SAMLPayloadExecute)
		if string(jsonSAMLPayloadExecute) == "{}" { // empty struct
			dst.SAMLPayloadExecute = nil
		} else {
			match++
		}
	} else {
		dst.SAMLPayloadExecute = nil
	}

	// try to unmarshal data into TelephonyRequestExecute
	err = json.Unmarshal(data, &dst.TelephonyRequestExecute)
	if err == nil {
		jsonTelephonyRequestExecute, _ := json.Marshal(dst.TelephonyRequestExecute)
		if string(jsonTelephonyRequestExecute) == "{}" { // empty struct
			dst.TelephonyRequestExecute = nil
		} else {
			match++
		}
	} else {
		dst.TelephonyRequestExecute = nil
	}

	// try to unmarshal data into TokenRequest
	err = json.Unmarshal(data, &dst.TokenRequest)
	if err == nil {
		jsonTokenRequest, _ := json.Marshal(dst.TokenRequest)
		if string(jsonTokenRequest) == "{}" { // empty struct
			dst.TokenRequest = nil
		} else {
			match++
		}
	} else {
		dst.TokenRequest = nil
	}

	// try to unmarshal data into UserImportRequestExecute
	err = json.Unmarshal(data, &dst.UserImportRequestExecute)
	if err == nil {
		jsonUserImportRequestExecute, _ := json.Marshal(dst.UserImportRequestExecute)
		if string(jsonUserImportRequestExecute) == "{}" { // empty struct
			dst.UserImportRequestExecute = nil
		} else {
			match++
		}
	} else {
		dst.UserImportRequestExecute = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PasswordImportRequestExecute = nil
		dst.RegistrationInlineHookRequest = nil
		dst.SAMLPayloadExecute = nil
		dst.TelephonyRequestExecute = nil
		dst.TokenRequest = nil
		dst.UserImportRequestExecute = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ExecuteInlineHookRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExecuteInlineHookRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExecuteInlineHookRequest) MarshalJSON() ([]byte, error) {
	if src.PasswordImportRequestExecute != nil {
		return json.Marshal(&src.PasswordImportRequestExecute)
	}

	if src.RegistrationInlineHookRequest != nil {
		return json.Marshal(&src.RegistrationInlineHookRequest)
	}

	if src.SAMLPayloadExecute != nil {
		return json.Marshal(&src.SAMLPayloadExecute)
	}

	if src.TelephonyRequestExecute != nil {
		return json.Marshal(&src.TelephonyRequestExecute)
	}

	if src.TokenRequest != nil {
		return json.Marshal(&src.TokenRequest)
	}

	if src.UserImportRequestExecute != nil {
		return json.Marshal(&src.UserImportRequestExecute)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExecuteInlineHookRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PasswordImportRequestExecute != nil {
		return obj.PasswordImportRequestExecute
	}

	if obj.RegistrationInlineHookRequest != nil {
		return obj.RegistrationInlineHookRequest
	}

	if obj.SAMLPayloadExecute != nil {
		return obj.SAMLPayloadExecute
	}

	if obj.TelephonyRequestExecute != nil {
		return obj.TelephonyRequestExecute
	}

	if obj.TokenRequest != nil {
		return obj.TokenRequest
	}

	if obj.UserImportRequestExecute != nil {
		return obj.UserImportRequestExecute
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ExecuteInlineHookRequest) GetActualInstanceValue() interface{} {
	if obj.PasswordImportRequestExecute != nil {
		return *obj.PasswordImportRequestExecute
	}

	if obj.RegistrationInlineHookRequest != nil {
		return *obj.RegistrationInlineHookRequest
	}

	if obj.SAMLPayloadExecute != nil {
		return *obj.SAMLPayloadExecute
	}

	if obj.TelephonyRequestExecute != nil {
		return *obj.TelephonyRequestExecute
	}

	if obj.TokenRequest != nil {
		return *obj.TokenRequest
	}

	if obj.UserImportRequestExecute != nil {
		return *obj.UserImportRequestExecute
	}

	// all schemas are nil
	return nil
}

type NullableExecuteInlineHookRequest struct {
	value *ExecuteInlineHookRequest
	isSet bool
}

func (v NullableExecuteInlineHookRequest) Get() *ExecuteInlineHookRequest {
	return v.value
}

func (v *NullableExecuteInlineHookRequest) Set(val *ExecuteInlineHookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteInlineHookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteInlineHookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteInlineHookRequest(val *ExecuteInlineHookRequest) *NullableExecuteInlineHookRequest {
	return &NullableExecuteInlineHookRequest{value: val, isSet: true}
}

func (v NullableExecuteInlineHookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteInlineHookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
