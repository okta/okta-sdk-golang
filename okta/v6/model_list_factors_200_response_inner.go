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

// ListFactors200ResponseInner - struct for ListFactors200ResponseInner
type ListFactors200ResponseInner struct {
	UserFactorCall              *UserFactorCall
	UserFactorEmail             *UserFactorEmail
	UserFactorPush              *UserFactorPush
	UserFactorSMS               *UserFactorSMS
	UserFactorSecurityQuestion  *UserFactorSecurityQuestion
	UserFactorToken             *UserFactorToken
	UserFactorTokenHOTP         *UserFactorTokenHOTP
	UserFactorTokenHardware     *UserFactorTokenHardware
	UserFactorTokenSoftwareTOTP *UserFactorTokenSoftwareTOTP
	UserFactorU2F               *UserFactorU2F
	UserFactorWeb               *UserFactorWeb
	UserFactorWebAuthn          *UserFactorWebAuthn
}

// UserFactorCallAsListFactors200ResponseInner is a convenience function that returns UserFactorCall wrapped in ListFactors200ResponseInner
func UserFactorCallAsListFactors200ResponseInner(v *UserFactorCall) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorCall: v,
	}
}

// UserFactorEmailAsListFactors200ResponseInner is a convenience function that returns UserFactorEmail wrapped in ListFactors200ResponseInner
func UserFactorEmailAsListFactors200ResponseInner(v *UserFactorEmail) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorEmail: v,
	}
}

// UserFactorPushAsListFactors200ResponseInner is a convenience function that returns UserFactorPush wrapped in ListFactors200ResponseInner
func UserFactorPushAsListFactors200ResponseInner(v *UserFactorPush) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorPush: v,
	}
}

// UserFactorSMSAsListFactors200ResponseInner is a convenience function that returns UserFactorSMS wrapped in ListFactors200ResponseInner
func UserFactorSMSAsListFactors200ResponseInner(v *UserFactorSMS) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorSMS: v,
	}
}

// UserFactorSecurityQuestionAsListFactors200ResponseInner is a convenience function that returns UserFactorSecurityQuestion wrapped in ListFactors200ResponseInner
func UserFactorSecurityQuestionAsListFactors200ResponseInner(v *UserFactorSecurityQuestion) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorSecurityQuestion: v,
	}
}

// UserFactorTokenAsListFactors200ResponseInner is a convenience function that returns UserFactorToken wrapped in ListFactors200ResponseInner
func UserFactorTokenAsListFactors200ResponseInner(v *UserFactorToken) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorToken: v,
	}
}

// UserFactorTokenHOTPAsListFactors200ResponseInner is a convenience function that returns UserFactorTokenHOTP wrapped in ListFactors200ResponseInner
func UserFactorTokenHOTPAsListFactors200ResponseInner(v *UserFactorTokenHOTP) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorTokenHOTP: v,
	}
}

// UserFactorTokenHardwareAsListFactors200ResponseInner is a convenience function that returns UserFactorTokenHardware wrapped in ListFactors200ResponseInner
func UserFactorTokenHardwareAsListFactors200ResponseInner(v *UserFactorTokenHardware) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorTokenHardware: v,
	}
}

// UserFactorTokenSoftwareTOTPAsListFactors200ResponseInner is a convenience function that returns UserFactorTokenSoftwareTOTP wrapped in ListFactors200ResponseInner
func UserFactorTokenSoftwareTOTPAsListFactors200ResponseInner(v *UserFactorTokenSoftwareTOTP) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorTokenSoftwareTOTP: v,
	}
}

// UserFactorU2FAsListFactors200ResponseInner is a convenience function that returns UserFactorU2F wrapped in ListFactors200ResponseInner
func UserFactorU2FAsListFactors200ResponseInner(v *UserFactorU2F) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorU2F: v,
	}
}

// UserFactorWebAsListFactors200ResponseInner is a convenience function that returns UserFactorWeb wrapped in ListFactors200ResponseInner
func UserFactorWebAsListFactors200ResponseInner(v *UserFactorWeb) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorWeb: v,
	}
}

// UserFactorWebAuthnAsListFactors200ResponseInner is a convenience function that returns UserFactorWebAuthn wrapped in ListFactors200ResponseInner
func UserFactorWebAuthnAsListFactors200ResponseInner(v *UserFactorWebAuthn) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		UserFactorWebAuthn: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListFactors200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'call'
	if jsonDict["factorType"] == "call" {
		// try to unmarshal JSON data into UserFactorCall
		err = json.Unmarshal(data, &dst.UserFactorCall)
		if err == nil {
			return nil // data stored in dst.UserFactorCall, return on the first match
		} else {
			dst.UserFactorCall = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorCall: %s", err.Error())
		}
	}

	// check if the discriminator value is 'email'
	if jsonDict["factorType"] == "email" {
		// try to unmarshal JSON data into UserFactorEmail
		err = json.Unmarshal(data, &dst.UserFactorEmail)
		if err == nil {
			return nil // data stored in dst.UserFactorEmail, return on the first match
		} else {
			dst.UserFactorEmail = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorEmail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'push'
	if jsonDict["factorType"] == "push" {
		// try to unmarshal JSON data into UserFactorPush
		err = json.Unmarshal(data, &dst.UserFactorPush)
		if err == nil {
			return nil // data stored in dst.UserFactorPush, return on the first match
		} else {
			dst.UserFactorPush = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorPush: %s", err.Error())
		}
	}

	// check if the discriminator value is 'question'
	if jsonDict["factorType"] == "question" {
		// try to unmarshal JSON data into UserFactorSecurityQuestion
		err = json.Unmarshal(data, &dst.UserFactorSecurityQuestion)
		if err == nil {
			return nil // data stored in dst.UserFactorSecurityQuestion, return on the first match
		} else {
			dst.UserFactorSecurityQuestion = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorSecurityQuestion: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sms'
	if jsonDict["factorType"] == "sms" {
		// try to unmarshal JSON data into UserFactorSMS
		err = json.Unmarshal(data, &dst.UserFactorSMS)
		if err == nil {
			return nil // data stored in dst.UserFactorSMS, return on the first match
		} else {
			dst.UserFactorSMS = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorSMS: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token'
	if jsonDict["factorType"] == "token" {
		// try to unmarshal JSON data into UserFactorToken
		err = json.Unmarshal(data, &dst.UserFactorToken)
		if err == nil {
			return nil // data stored in dst.UserFactorToken, return on the first match
		} else {
			dst.UserFactorToken = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorToken: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token:hardware'
	if jsonDict["factorType"] == "token:hardware" {
		// try to unmarshal JSON data into UserFactorTokenHardware
		err = json.Unmarshal(data, &dst.UserFactorTokenHardware)
		if err == nil {
			return nil // data stored in dst.UserFactorTokenHardware, return on the first match
		} else {
			dst.UserFactorTokenHardware = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorTokenHardware: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token:hotp'
	if jsonDict["factorType"] == "token:hotp" {
		// try to unmarshal JSON data into UserFactorTokenHOTP
		err = json.Unmarshal(data, &dst.UserFactorTokenHOTP)
		if err == nil {
			return nil // data stored in dst.UserFactorTokenHOTP, return on the first match
		} else {
			dst.UserFactorTokenHOTP = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorTokenHOTP: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token:software:totp'
	if jsonDict["factorType"] == "token:software:totp" {
		// try to unmarshal JSON data into UserFactorTokenSoftwareTOTP
		err = json.Unmarshal(data, &dst.UserFactorTokenSoftwareTOTP)
		if err == nil {
			return nil // data stored in dst.UserFactorTokenSoftwareTOTP, return on the first match
		} else {
			dst.UserFactorTokenSoftwareTOTP = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorTokenSoftwareTOTP: %s", err.Error())
		}
	}

	// check if the discriminator value is 'u2f'
	if jsonDict["factorType"] == "u2f" {
		// try to unmarshal JSON data into UserFactorU2F
		err = json.Unmarshal(data, &dst.UserFactorU2F)
		if err == nil {
			return nil // data stored in dst.UserFactorU2F, return on the first match
		} else {
			dst.UserFactorU2F = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorU2F: %s", err.Error())
		}
	}

	// check if the discriminator value is 'web'
	if jsonDict["factorType"] == "web" {
		// try to unmarshal JSON data into UserFactorWeb
		err = json.Unmarshal(data, &dst.UserFactorWeb)
		if err == nil {
			return nil // data stored in dst.UserFactorWeb, return on the first match
		} else {
			dst.UserFactorWeb = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorWeb: %s", err.Error())
		}
	}

	// check if the discriminator value is 'webauthn'
	if jsonDict["factorType"] == "webauthn" {
		// try to unmarshal JSON data into UserFactorWebAuthn
		err = json.Unmarshal(data, &dst.UserFactorWebAuthn)
		if err == nil {
			return nil // data stored in dst.UserFactorWebAuthn, return on the first match
		} else {
			dst.UserFactorWebAuthn = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorWebAuthn: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorCall'
	if jsonDict["factorType"] == "UserFactorCall" {
		// try to unmarshal JSON data into UserFactorCall
		err = json.Unmarshal(data, &dst.UserFactorCall)
		if err == nil {
			return nil // data stored in dst.UserFactorCall, return on the first match
		} else {
			dst.UserFactorCall = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorCall: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorEmail'
	if jsonDict["factorType"] == "UserFactorEmail" {
		// try to unmarshal JSON data into UserFactorEmail
		err = json.Unmarshal(data, &dst.UserFactorEmail)
		if err == nil {
			return nil // data stored in dst.UserFactorEmail, return on the first match
		} else {
			dst.UserFactorEmail = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorEmail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorPush'
	if jsonDict["factorType"] == "UserFactorPush" {
		// try to unmarshal JSON data into UserFactorPush
		err = json.Unmarshal(data, &dst.UserFactorPush)
		if err == nil {
			return nil // data stored in dst.UserFactorPush, return on the first match
		} else {
			dst.UserFactorPush = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorPush: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorSMS'
	if jsonDict["factorType"] == "UserFactorSMS" {
		// try to unmarshal JSON data into UserFactorSMS
		err = json.Unmarshal(data, &dst.UserFactorSMS)
		if err == nil {
			return nil // data stored in dst.UserFactorSMS, return on the first match
		} else {
			dst.UserFactorSMS = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorSMS: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorSecurityQuestion'
	if jsonDict["factorType"] == "UserFactorSecurityQuestion" {
		// try to unmarshal JSON data into UserFactorSecurityQuestion
		err = json.Unmarshal(data, &dst.UserFactorSecurityQuestion)
		if err == nil {
			return nil // data stored in dst.UserFactorSecurityQuestion, return on the first match
		} else {
			dst.UserFactorSecurityQuestion = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorSecurityQuestion: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorToken'
	if jsonDict["factorType"] == "UserFactorToken" {
		// try to unmarshal JSON data into UserFactorToken
		err = json.Unmarshal(data, &dst.UserFactorToken)
		if err == nil {
			return nil // data stored in dst.UserFactorToken, return on the first match
		} else {
			dst.UserFactorToken = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorToken: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorTokenHOTP'
	if jsonDict["factorType"] == "UserFactorTokenHOTP" {
		// try to unmarshal JSON data into UserFactorTokenHOTP
		err = json.Unmarshal(data, &dst.UserFactorTokenHOTP)
		if err == nil {
			return nil // data stored in dst.UserFactorTokenHOTP, return on the first match
		} else {
			dst.UserFactorTokenHOTP = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorTokenHOTP: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorTokenHardware'
	if jsonDict["factorType"] == "UserFactorTokenHardware" {
		// try to unmarshal JSON data into UserFactorTokenHardware
		err = json.Unmarshal(data, &dst.UserFactorTokenHardware)
		if err == nil {
			return nil // data stored in dst.UserFactorTokenHardware, return on the first match
		} else {
			dst.UserFactorTokenHardware = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorTokenHardware: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorTokenSoftwareTOTP'
	if jsonDict["factorType"] == "UserFactorTokenSoftwareTOTP" {
		// try to unmarshal JSON data into UserFactorTokenSoftwareTOTP
		err = json.Unmarshal(data, &dst.UserFactorTokenSoftwareTOTP)
		if err == nil {
			return nil // data stored in dst.UserFactorTokenSoftwareTOTP, return on the first match
		} else {
			dst.UserFactorTokenSoftwareTOTP = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorTokenSoftwareTOTP: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorU2F'
	if jsonDict["factorType"] == "UserFactorU2F" {
		// try to unmarshal JSON data into UserFactorU2F
		err = json.Unmarshal(data, &dst.UserFactorU2F)
		if err == nil {
			return nil // data stored in dst.UserFactorU2F, return on the first match
		} else {
			dst.UserFactorU2F = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorU2F: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorWeb'
	if jsonDict["factorType"] == "UserFactorWeb" {
		// try to unmarshal JSON data into UserFactorWeb
		err = json.Unmarshal(data, &dst.UserFactorWeb)
		if err == nil {
			return nil // data stored in dst.UserFactorWeb, return on the first match
		} else {
			dst.UserFactorWeb = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorWeb: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorWebAuthn'
	if jsonDict["factorType"] == "UserFactorWebAuthn" {
		// try to unmarshal JSON data into UserFactorWebAuthn
		err = json.Unmarshal(data, &dst.UserFactorWebAuthn)
		if err == nil {
			return nil // data stored in dst.UserFactorWebAuthn, return on the first match
		} else {
			dst.UserFactorWebAuthn = nil
			return fmt.Errorf("failed to unmarshal ListFactors200ResponseInner as UserFactorWebAuthn: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListFactors200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.UserFactorCall != nil {
		return json.Marshal(&src.UserFactorCall)
	}

	if src.UserFactorEmail != nil {
		return json.Marshal(&src.UserFactorEmail)
	}

	if src.UserFactorPush != nil {
		return json.Marshal(&src.UserFactorPush)
	}

	if src.UserFactorSMS != nil {
		return json.Marshal(&src.UserFactorSMS)
	}

	if src.UserFactorSecurityQuestion != nil {
		return json.Marshal(&src.UserFactorSecurityQuestion)
	}

	if src.UserFactorToken != nil {
		return json.Marshal(&src.UserFactorToken)
	}

	if src.UserFactorTokenHOTP != nil {
		return json.Marshal(&src.UserFactorTokenHOTP)
	}

	if src.UserFactorTokenHardware != nil {
		return json.Marshal(&src.UserFactorTokenHardware)
	}

	if src.UserFactorTokenSoftwareTOTP != nil {
		return json.Marshal(&src.UserFactorTokenSoftwareTOTP)
	}

	if src.UserFactorU2F != nil {
		return json.Marshal(&src.UserFactorU2F)
	}

	if src.UserFactorWeb != nil {
		return json.Marshal(&src.UserFactorWeb)
	}

	if src.UserFactorWebAuthn != nil {
		return json.Marshal(&src.UserFactorWebAuthn)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListFactors200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserFactorCall != nil {
		return obj.UserFactorCall
	}

	if obj.UserFactorEmail != nil {
		return obj.UserFactorEmail
	}

	if obj.UserFactorPush != nil {
		return obj.UserFactorPush
	}

	if obj.UserFactorSMS != nil {
		return obj.UserFactorSMS
	}

	if obj.UserFactorSecurityQuestion != nil {
		return obj.UserFactorSecurityQuestion
	}

	if obj.UserFactorToken != nil {
		return obj.UserFactorToken
	}

	if obj.UserFactorTokenHOTP != nil {
		return obj.UserFactorTokenHOTP
	}

	if obj.UserFactorTokenHardware != nil {
		return obj.UserFactorTokenHardware
	}

	if obj.UserFactorTokenSoftwareTOTP != nil {
		return obj.UserFactorTokenSoftwareTOTP
	}

	if obj.UserFactorU2F != nil {
		return obj.UserFactorU2F
	}

	if obj.UserFactorWeb != nil {
		return obj.UserFactorWeb
	}

	if obj.UserFactorWebAuthn != nil {
		return obj.UserFactorWebAuthn
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListFactors200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.UserFactorCall != nil {
		return *obj.UserFactorCall
	}

	if obj.UserFactorEmail != nil {
		return *obj.UserFactorEmail
	}

	if obj.UserFactorPush != nil {
		return *obj.UserFactorPush
	}

	if obj.UserFactorSMS != nil {
		return *obj.UserFactorSMS
	}

	if obj.UserFactorSecurityQuestion != nil {
		return *obj.UserFactorSecurityQuestion
	}

	if obj.UserFactorToken != nil {
		return *obj.UserFactorToken
	}

	if obj.UserFactorTokenHOTP != nil {
		return *obj.UserFactorTokenHOTP
	}

	if obj.UserFactorTokenHardware != nil {
		return *obj.UserFactorTokenHardware
	}

	if obj.UserFactorTokenSoftwareTOTP != nil {
		return *obj.UserFactorTokenSoftwareTOTP
	}

	if obj.UserFactorU2F != nil {
		return *obj.UserFactorU2F
	}

	if obj.UserFactorWeb != nil {
		return *obj.UserFactorWeb
	}

	if obj.UserFactorWebAuthn != nil {
		return *obj.UserFactorWebAuthn
	}

	// all schemas are nil
	return nil
}

type NullableListFactors200ResponseInner struct {
	value *ListFactors200ResponseInner
	isSet bool
}

func (v NullableListFactors200ResponseInner) Get() *ListFactors200ResponseInner {
	return v.value
}

func (v *NullableListFactors200ResponseInner) Set(val *ListFactors200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListFactors200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListFactors200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFactors200ResponseInner(val *ListFactors200ResponseInner) *NullableListFactors200ResponseInner {
	return &NullableListFactors200ResponseInner{value: val, isSet: true}
}

func (v NullableListFactors200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFactors200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
