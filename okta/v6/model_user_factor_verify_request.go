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

// UserFactorVerifyRequest - struct for UserFactorVerifyRequest
type UserFactorVerifyRequest struct {
	Call1              *Call1
	Email1             *Email1
	Push1              *Push1
	Question           *Question
	Sms1               *Sms1
	Token              *Token
	TokenHardware      *TokenHardware
	TokenHotp          *TokenHotp
	TokenSoftwareTotp1 *TokenSoftwareTotp1
	U2f1               *U2f1
	Webauthn1          *Webauthn1
}

// Call1AsUserFactorVerifyRequest is a convenience function that returns Call1 wrapped in UserFactorVerifyRequest
func Call1AsUserFactorVerifyRequest(v *Call1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Call1: v,
	}
}

// Email1AsUserFactorVerifyRequest is a convenience function that returns Email1 wrapped in UserFactorVerifyRequest
func Email1AsUserFactorVerifyRequest(v *Email1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Email1: v,
	}
}

// Push1AsUserFactorVerifyRequest is a convenience function that returns Push1 wrapped in UserFactorVerifyRequest
func Push1AsUserFactorVerifyRequest(v *Push1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Push1: v,
	}
}

// QuestionAsUserFactorVerifyRequest is a convenience function that returns Question wrapped in UserFactorVerifyRequest
func QuestionAsUserFactorVerifyRequest(v *Question) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Question: v,
	}
}

// Sms1AsUserFactorVerifyRequest is a convenience function that returns Sms1 wrapped in UserFactorVerifyRequest
func Sms1AsUserFactorVerifyRequest(v *Sms1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Sms1: v,
	}
}

// TokenAsUserFactorVerifyRequest is a convenience function that returns Token wrapped in UserFactorVerifyRequest
func TokenAsUserFactorVerifyRequest(v *Token) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Token: v,
	}
}

// TokenHardwareAsUserFactorVerifyRequest is a convenience function that returns TokenHardware wrapped in UserFactorVerifyRequest
func TokenHardwareAsUserFactorVerifyRequest(v *TokenHardware) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		TokenHardware: v,
	}
}

// TokenHotpAsUserFactorVerifyRequest is a convenience function that returns TokenHotp wrapped in UserFactorVerifyRequest
func TokenHotpAsUserFactorVerifyRequest(v *TokenHotp) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		TokenHotp: v,
	}
}

// TokenSoftwareTotp1AsUserFactorVerifyRequest is a convenience function that returns TokenSoftwareTotp1 wrapped in UserFactorVerifyRequest
func TokenSoftwareTotp1AsUserFactorVerifyRequest(v *TokenSoftwareTotp1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		TokenSoftwareTotp1: v,
	}
}

// U2f1AsUserFactorVerifyRequest is a convenience function that returns U2f1 wrapped in UserFactorVerifyRequest
func U2f1AsUserFactorVerifyRequest(v *U2f1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		U2f1: v,
	}
}

// Webauthn1AsUserFactorVerifyRequest is a convenience function that returns Webauthn1 wrapped in UserFactorVerifyRequest
func Webauthn1AsUserFactorVerifyRequest(v *Webauthn1) UserFactorVerifyRequest {
	return UserFactorVerifyRequest{
		Webauthn1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UserFactorVerifyRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Call1
	err = json.Unmarshal(data, &dst.Call1)
	if err == nil {
		jsonCall1, _ := json.Marshal(dst.Call1)
		if string(jsonCall1) == "{}" { // empty struct
			dst.Call1 = nil
		} else {
			match++
		}
	} else {
		dst.Call1 = nil
	}

	// try to unmarshal data into Email1
	err = json.Unmarshal(data, &dst.Email1)
	if err == nil {
		jsonEmail1, _ := json.Marshal(dst.Email1)
		if string(jsonEmail1) == "{}" { // empty struct
			dst.Email1 = nil
		} else {
			match++
		}
	} else {
		dst.Email1 = nil
	}

	// try to unmarshal data into Push1
	err = json.Unmarshal(data, &dst.Push1)
	if err == nil {
		jsonPush1, _ := json.Marshal(dst.Push1)
		if string(jsonPush1) == "{}" { // empty struct
			dst.Push1 = nil
		} else {
			match++
		}
	} else {
		dst.Push1 = nil
	}

	// try to unmarshal data into Question
	err = json.Unmarshal(data, &dst.Question)
	if err == nil {
		jsonQuestion, _ := json.Marshal(dst.Question)
		if string(jsonQuestion) == "{}" { // empty struct
			dst.Question = nil
		} else {
			match++
		}
	} else {
		dst.Question = nil
	}

	// try to unmarshal data into Sms1
	err = json.Unmarshal(data, &dst.Sms1)
	if err == nil {
		jsonSms1, _ := json.Marshal(dst.Sms1)
		if string(jsonSms1) == "{}" { // empty struct
			dst.Sms1 = nil
		} else {
			match++
		}
	} else {
		dst.Sms1 = nil
	}

	// try to unmarshal data into Token
	err = json.Unmarshal(data, &dst.Token)
	if err == nil {
		jsonToken, _ := json.Marshal(dst.Token)
		if string(jsonToken) == "{}" { // empty struct
			dst.Token = nil
		} else {
			match++
		}
	} else {
		dst.Token = nil
	}

	// try to unmarshal data into TokenHardware
	err = json.Unmarshal(data, &dst.TokenHardware)
	if err == nil {
		jsonTokenHardware, _ := json.Marshal(dst.TokenHardware)
		if string(jsonTokenHardware) == "{}" { // empty struct
			dst.TokenHardware = nil
		} else {
			match++
		}
	} else {
		dst.TokenHardware = nil
	}

	// try to unmarshal data into TokenHotp
	err = json.Unmarshal(data, &dst.TokenHotp)
	if err == nil {
		jsonTokenHotp, _ := json.Marshal(dst.TokenHotp)
		if string(jsonTokenHotp) == "{}" { // empty struct
			dst.TokenHotp = nil
		} else {
			match++
		}
	} else {
		dst.TokenHotp = nil
	}

	// try to unmarshal data into TokenSoftwareTotp1
	err = json.Unmarshal(data, &dst.TokenSoftwareTotp1)
	if err == nil {
		jsonTokenSoftwareTotp1, _ := json.Marshal(dst.TokenSoftwareTotp1)
		if string(jsonTokenSoftwareTotp1) == "{}" { // empty struct
			dst.TokenSoftwareTotp1 = nil
		} else {
			match++
		}
	} else {
		dst.TokenSoftwareTotp1 = nil
	}

	// try to unmarshal data into U2f1
	err = json.Unmarshal(data, &dst.U2f1)
	if err == nil {
		jsonU2f1, _ := json.Marshal(dst.U2f1)
		if string(jsonU2f1) == "{}" { // empty struct
			dst.U2f1 = nil
		} else {
			match++
		}
	} else {
		dst.U2f1 = nil
	}

	// try to unmarshal data into Webauthn1
	err = json.Unmarshal(data, &dst.Webauthn1)
	if err == nil {
		jsonWebauthn1, _ := json.Marshal(dst.Webauthn1)
		if string(jsonWebauthn1) == "{}" { // empty struct
			dst.Webauthn1 = nil
		} else {
			match++
		}
	} else {
		dst.Webauthn1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Call1 = nil
		dst.Email1 = nil
		dst.Push1 = nil
		dst.Question = nil
		dst.Sms1 = nil
		dst.Token = nil
		dst.TokenHardware = nil
		dst.TokenHotp = nil
		dst.TokenSoftwareTotp1 = nil
		dst.U2f1 = nil
		dst.Webauthn1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UserFactorVerifyRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UserFactorVerifyRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserFactorVerifyRequest) MarshalJSON() ([]byte, error) {
	if src.Call1 != nil {
		return json.Marshal(&src.Call1)
	}

	if src.Email1 != nil {
		return json.Marshal(&src.Email1)
	}

	if src.Push1 != nil {
		return json.Marshal(&src.Push1)
	}

	if src.Question != nil {
		return json.Marshal(&src.Question)
	}

	if src.Sms1 != nil {
		return json.Marshal(&src.Sms1)
	}

	if src.Token != nil {
		return json.Marshal(&src.Token)
	}

	if src.TokenHardware != nil {
		return json.Marshal(&src.TokenHardware)
	}

	if src.TokenHotp != nil {
		return json.Marshal(&src.TokenHotp)
	}

	if src.TokenSoftwareTotp1 != nil {
		return json.Marshal(&src.TokenSoftwareTotp1)
	}

	if src.U2f1 != nil {
		return json.Marshal(&src.U2f1)
	}

	if src.Webauthn1 != nil {
		return json.Marshal(&src.Webauthn1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserFactorVerifyRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Call1 != nil {
		return obj.Call1
	}

	if obj.Email1 != nil {
		return obj.Email1
	}

	if obj.Push1 != nil {
		return obj.Push1
	}

	if obj.Question != nil {
		return obj.Question
	}

	if obj.Sms1 != nil {
		return obj.Sms1
	}

	if obj.Token != nil {
		return obj.Token
	}

	if obj.TokenHardware != nil {
		return obj.TokenHardware
	}

	if obj.TokenHotp != nil {
		return obj.TokenHotp
	}

	if obj.TokenSoftwareTotp1 != nil {
		return obj.TokenSoftwareTotp1
	}

	if obj.U2f1 != nil {
		return obj.U2f1
	}

	if obj.Webauthn1 != nil {
		return obj.Webauthn1
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj UserFactorVerifyRequest) GetActualInstanceValue() interface{} {
	if obj.Call1 != nil {
		return *obj.Call1
	}

	if obj.Email1 != nil {
		return *obj.Email1
	}

	if obj.Push1 != nil {
		return *obj.Push1
	}

	if obj.Question != nil {
		return *obj.Question
	}

	if obj.Sms1 != nil {
		return *obj.Sms1
	}

	if obj.Token != nil {
		return *obj.Token
	}

	if obj.TokenHardware != nil {
		return *obj.TokenHardware
	}

	if obj.TokenHotp != nil {
		return *obj.TokenHotp
	}

	if obj.TokenSoftwareTotp1 != nil {
		return *obj.TokenSoftwareTotp1
	}

	if obj.U2f1 != nil {
		return *obj.U2f1
	}

	if obj.Webauthn1 != nil {
		return *obj.Webauthn1
	}

	// all schemas are nil
	return nil
}

type NullableUserFactorVerifyRequest struct {
	value *UserFactorVerifyRequest
	isSet bool
}

func (v NullableUserFactorVerifyRequest) Get() *UserFactorVerifyRequest {
	return v.value
}

func (v *NullableUserFactorVerifyRequest) Set(val *UserFactorVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorVerifyRequest(val *UserFactorVerifyRequest) *NullableUserFactorVerifyRequest {
	return &NullableUserFactorVerifyRequest{value: val, isSet: true}
}

func (v NullableUserFactorVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
