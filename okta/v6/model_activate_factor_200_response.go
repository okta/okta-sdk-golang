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

package okta

import (
	"encoding/json"
	"fmt"
)


//model_oneof.mustache
// ActivateFactor200Response - struct for ActivateFactor200Response
type ActivateFactor200Response struct {
	UserFactorCall *UserFactorCall
	UserFactorEmail *UserFactorEmail
	UserFactorPush *UserFactorPush
	UserFactorSMS *UserFactorSMS
	UserFactorTokenSoftwareTOTP *UserFactorTokenSoftwareTOTP
	UserFactorU2F *UserFactorU2F
	UserFactorWebAuthn *UserFactorWebAuthn
}

// UserFactorCallAsActivateFactor200Response is a convenience function that returns UserFactorCall wrapped in ActivateFactor200Response
func UserFactorCallAsActivateFactor200Response(v *UserFactorCall) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorCall: v,
	}
}

// UserFactorEmailAsActivateFactor200Response is a convenience function that returns UserFactorEmail wrapped in ActivateFactor200Response
func UserFactorEmailAsActivateFactor200Response(v *UserFactorEmail) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorEmail: v,
	}
}

// UserFactorPushAsActivateFactor200Response is a convenience function that returns UserFactorPush wrapped in ActivateFactor200Response
func UserFactorPushAsActivateFactor200Response(v *UserFactorPush) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorPush: v,
	}
}

// UserFactorSMSAsActivateFactor200Response is a convenience function that returns UserFactorSMS wrapped in ActivateFactor200Response
func UserFactorSMSAsActivateFactor200Response(v *UserFactorSMS) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorSMS: v,
	}
}

// UserFactorTokenSoftwareTOTPAsActivateFactor200Response is a convenience function that returns UserFactorTokenSoftwareTOTP wrapped in ActivateFactor200Response
func UserFactorTokenSoftwareTOTPAsActivateFactor200Response(v *UserFactorTokenSoftwareTOTP) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorTokenSoftwareTOTP: v,
	}
}

// UserFactorU2FAsActivateFactor200Response is a convenience function that returns UserFactorU2F wrapped in ActivateFactor200Response
func UserFactorU2FAsActivateFactor200Response(v *UserFactorU2F) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorU2F: v,
	}
}

// UserFactorWebAuthnAsActivateFactor200Response is a convenience function that returns UserFactorWebAuthn wrapped in ActivateFactor200Response
func UserFactorWebAuthnAsActivateFactor200Response(v *UserFactorWebAuthn) ActivateFactor200Response {
	return ActivateFactor200Response{
		UserFactorWebAuthn: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ActivateFactor200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'call'
	if jsonDict["factorType"] == "call" {
		// try to unmarshal JSON data into UserFactorCall
		err = json.Unmarshal(data, &dst.UserFactorCall)
		if err == nil {
			return nil // data stored in dst.UserFactorCall, return on the first match
		} else {
			dst.UserFactorCall = nil
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorCall: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorEmail: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorPush: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorSMS: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorTokenSoftwareTOTP: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorU2F: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorWebAuthn: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorCall: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorEmail: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorPush: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorSMS: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorTokenSoftwareTOTP: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorU2F: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ActivateFactor200Response as UserFactorWebAuthn: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ActivateFactor200Response) MarshalJSON() ([]byte, error) {
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

	if src.UserFactorTokenSoftwareTOTP != nil {
		return json.Marshal(&src.UserFactorTokenSoftwareTOTP)
	}

	if src.UserFactorU2F != nil {
		return json.Marshal(&src.UserFactorU2F)
	}

	if src.UserFactorWebAuthn != nil {
		return json.Marshal(&src.UserFactorWebAuthn)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ActivateFactor200Response) GetActualInstance() (interface{}) {
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

	if obj.UserFactorTokenSoftwareTOTP != nil {
		return obj.UserFactorTokenSoftwareTOTP
	}

	if obj.UserFactorU2F != nil {
		return obj.UserFactorU2F
	}

	if obj.UserFactorWebAuthn != nil {
		return obj.UserFactorWebAuthn
	}

	// all schemas are nil
	return nil
}

type NullableActivateFactor200Response struct {
	value *ActivateFactor200Response
	isSet bool
}

func (v NullableActivateFactor200Response) Get() *ActivateFactor200Response {
	return v.value
}

func (v *NullableActivateFactor200Response) Set(val *ActivateFactor200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateFactor200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateFactor200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateFactor200Response(val *ActivateFactor200Response) *NullableActivateFactor200Response {
	return &NullableActivateFactor200Response{value: val, isSet: true}
}

func (v NullableActivateFactor200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateFactor200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


