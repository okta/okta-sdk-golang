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
// ListAuthenticatorMethods200ResponseInner - struct for ListAuthenticatorMethods200ResponseInner
type ListAuthenticatorMethods200ResponseInner struct {
	AuthenticatorMethodOtp *AuthenticatorMethodOtp
	AuthenticatorMethodPush *AuthenticatorMethodPush
	AuthenticatorMethodSignedNonce *AuthenticatorMethodSignedNonce
	AuthenticatorMethodSimple *AuthenticatorMethodSimple
	AuthenticatorMethodTotp *AuthenticatorMethodTotp
	AuthenticatorMethodWebAuthn *AuthenticatorMethodWebAuthn
	AuthenticatorMethodWithVerifiableProperties *AuthenticatorMethodWithVerifiableProperties
}

// AuthenticatorMethodOtpAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodOtp wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodOtpAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodOtp) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodOtp: v,
	}
}

// AuthenticatorMethodPushAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodPush wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodPushAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodPush) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodPush: v,
	}
}

// AuthenticatorMethodSignedNonceAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodSignedNonce wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodSignedNonceAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodSignedNonce) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodSignedNonce: v,
	}
}

// AuthenticatorMethodSimpleAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodSimple wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodSimpleAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodSimple) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodSimple: v,
	}
}

// AuthenticatorMethodTotpAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodTotp wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodTotpAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodTotp) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodTotp: v,
	}
}

// AuthenticatorMethodWebAuthnAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodWebAuthn wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodWebAuthnAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodWebAuthn) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodWebAuthn: v,
	}
}

// AuthenticatorMethodWithVerifiablePropertiesAsListAuthenticatorMethods200ResponseInner is a convenience function that returns AuthenticatorMethodWithVerifiableProperties wrapped in ListAuthenticatorMethods200ResponseInner
func AuthenticatorMethodWithVerifiablePropertiesAsListAuthenticatorMethods200ResponseInner(v *AuthenticatorMethodWithVerifiableProperties) ListAuthenticatorMethods200ResponseInner {
	return ListAuthenticatorMethods200ResponseInner{
		AuthenticatorMethodWithVerifiableProperties: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListAuthenticatorMethods200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AuthenticatorMethodOtp'
	if jsonDict["type"] == "AuthenticatorMethodOtp" {
		// try to unmarshal JSON data into AuthenticatorMethodOtp
		err = json.Unmarshal(data, &dst.AuthenticatorMethodOtp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodOtp, return on the first match
		} else {
			dst.AuthenticatorMethodOtp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodOtp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorMethodPush'
	if jsonDict["type"] == "AuthenticatorMethodPush" {
		// try to unmarshal JSON data into AuthenticatorMethodPush
		err = json.Unmarshal(data, &dst.AuthenticatorMethodPush)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodPush, return on the first match
		} else {
			dst.AuthenticatorMethodPush = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodPush: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorMethodSignedNonce'
	if jsonDict["type"] == "AuthenticatorMethodSignedNonce" {
		// try to unmarshal JSON data into AuthenticatorMethodSignedNonce
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSignedNonce)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSignedNonce, return on the first match
		} else {
			dst.AuthenticatorMethodSignedNonce = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSignedNonce: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorMethodSimple'
	if jsonDict["type"] == "AuthenticatorMethodSimple" {
		// try to unmarshal JSON data into AuthenticatorMethodSimple
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSimple)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSimple, return on the first match
		} else {
			dst.AuthenticatorMethodSimple = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSimple: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorMethodTotp'
	if jsonDict["type"] == "AuthenticatorMethodTotp" {
		// try to unmarshal JSON data into AuthenticatorMethodTotp
		err = json.Unmarshal(data, &dst.AuthenticatorMethodTotp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodTotp, return on the first match
		} else {
			dst.AuthenticatorMethodTotp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodTotp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorMethodWebAuthn'
	if jsonDict["type"] == "AuthenticatorMethodWebAuthn" {
		// try to unmarshal JSON data into AuthenticatorMethodWebAuthn
		err = json.Unmarshal(data, &dst.AuthenticatorMethodWebAuthn)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodWebAuthn, return on the first match
		} else {
			dst.AuthenticatorMethodWebAuthn = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodWebAuthn: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorMethodWithVerifiableProperties'
	if jsonDict["type"] == "AuthenticatorMethodWithVerifiableProperties" {
		// try to unmarshal JSON data into AuthenticatorMethodWithVerifiableProperties
		err = json.Unmarshal(data, &dst.AuthenticatorMethodWithVerifiableProperties)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodWithVerifiableProperties, return on the first match
		} else {
			dst.AuthenticatorMethodWithVerifiableProperties = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodWithVerifiableProperties: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cert'
	if jsonDict["type"] == "cert" {
		// try to unmarshal JSON data into AuthenticatorMethodWithVerifiableProperties
		err = json.Unmarshal(data, &dst.AuthenticatorMethodWithVerifiableProperties)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodWithVerifiableProperties, return on the first match
		} else {
			dst.AuthenticatorMethodWithVerifiableProperties = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodWithVerifiableProperties: %s", err.Error())
		}
	}

	// check if the discriminator value is 'duo'
	if jsonDict["type"] == "duo" {
		// try to unmarshal JSON data into AuthenticatorMethodWithVerifiableProperties
		err = json.Unmarshal(data, &dst.AuthenticatorMethodWithVerifiableProperties)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodWithVerifiableProperties, return on the first match
		} else {
			dst.AuthenticatorMethodWithVerifiableProperties = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodWithVerifiableProperties: %s", err.Error())
		}
	}

	// check if the discriminator value is 'email'
	if jsonDict["type"] == "email" {
		// try to unmarshal JSON data into AuthenticatorMethodSimple
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSimple)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSimple, return on the first match
		} else {
			dst.AuthenticatorMethodSimple = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSimple: %s", err.Error())
		}
	}

	// check if the discriminator value is 'idp'
	if jsonDict["type"] == "idp" {
		// try to unmarshal JSON data into AuthenticatorMethodWithVerifiableProperties
		err = json.Unmarshal(data, &dst.AuthenticatorMethodWithVerifiableProperties)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodWithVerifiableProperties, return on the first match
		} else {
			dst.AuthenticatorMethodWithVerifiableProperties = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodWithVerifiableProperties: %s", err.Error())
		}
	}

	// check if the discriminator value is 'otp'
	if jsonDict["type"] == "otp" {
		// try to unmarshal JSON data into AuthenticatorMethodOtp
		err = json.Unmarshal(data, &dst.AuthenticatorMethodOtp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodOtp, return on the first match
		} else {
			dst.AuthenticatorMethodOtp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodOtp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'password'
	if jsonDict["type"] == "password" {
		// try to unmarshal JSON data into AuthenticatorMethodSimple
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSimple)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSimple, return on the first match
		} else {
			dst.AuthenticatorMethodSimple = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSimple: %s", err.Error())
		}
	}

	// check if the discriminator value is 'push'
	if jsonDict["type"] == "push" {
		// try to unmarshal JSON data into AuthenticatorMethodPush
		err = json.Unmarshal(data, &dst.AuthenticatorMethodPush)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodPush, return on the first match
		} else {
			dst.AuthenticatorMethodPush = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodPush: %s", err.Error())
		}
	}

	// check if the discriminator value is 'security_question'
	if jsonDict["type"] == "security_question" {
		// try to unmarshal JSON data into AuthenticatorMethodSimple
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSimple)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSimple, return on the first match
		} else {
			dst.AuthenticatorMethodSimple = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSimple: %s", err.Error())
		}
	}

	// check if the discriminator value is 'signed_nonce'
	if jsonDict["type"] == "signed_nonce" {
		// try to unmarshal JSON data into AuthenticatorMethodSignedNonce
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSignedNonce)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSignedNonce, return on the first match
		} else {
			dst.AuthenticatorMethodSignedNonce = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSignedNonce: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sms'
	if jsonDict["type"] == "sms" {
		// try to unmarshal JSON data into AuthenticatorMethodSimple
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSimple)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSimple, return on the first match
		} else {
			dst.AuthenticatorMethodSimple = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSimple: %s", err.Error())
		}
	}

	// check if the discriminator value is 'totp'
	if jsonDict["type"] == "totp" {
		// try to unmarshal JSON data into AuthenticatorMethodTotp
		err = json.Unmarshal(data, &dst.AuthenticatorMethodTotp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodTotp, return on the first match
		} else {
			dst.AuthenticatorMethodTotp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodTotp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'voice'
	if jsonDict["type"] == "voice" {
		// try to unmarshal JSON data into AuthenticatorMethodSimple
		err = json.Unmarshal(data, &dst.AuthenticatorMethodSimple)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodSimple, return on the first match
		} else {
			dst.AuthenticatorMethodSimple = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodSimple: %s", err.Error())
		}
	}

	// check if the discriminator value is 'webauthn'
	if jsonDict["type"] == "webauthn" {
		// try to unmarshal JSON data into AuthenticatorMethodWebAuthn
		err = json.Unmarshal(data, &dst.AuthenticatorMethodWebAuthn)
		if err == nil {
			return nil // data stored in dst.AuthenticatorMethodWebAuthn, return on the first match
		} else {
			dst.AuthenticatorMethodWebAuthn = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticatorMethods200ResponseInner as AuthenticatorMethodWebAuthn: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAuthenticatorMethods200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AuthenticatorMethodOtp != nil {
		return json.Marshal(&src.AuthenticatorMethodOtp)
	}

	if src.AuthenticatorMethodPush != nil {
		return json.Marshal(&src.AuthenticatorMethodPush)
	}

	if src.AuthenticatorMethodSignedNonce != nil {
		return json.Marshal(&src.AuthenticatorMethodSignedNonce)
	}

	if src.AuthenticatorMethodSimple != nil {
		return json.Marshal(&src.AuthenticatorMethodSimple)
	}

	if src.AuthenticatorMethodTotp != nil {
		return json.Marshal(&src.AuthenticatorMethodTotp)
	}

	if src.AuthenticatorMethodWebAuthn != nil {
		return json.Marshal(&src.AuthenticatorMethodWebAuthn)
	}

	if src.AuthenticatorMethodWithVerifiableProperties != nil {
		return json.Marshal(&src.AuthenticatorMethodWithVerifiableProperties)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAuthenticatorMethods200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuthenticatorMethodOtp != nil {
		return obj.AuthenticatorMethodOtp
	}

	if obj.AuthenticatorMethodPush != nil {
		return obj.AuthenticatorMethodPush
	}

	if obj.AuthenticatorMethodSignedNonce != nil {
		return obj.AuthenticatorMethodSignedNonce
	}

	if obj.AuthenticatorMethodSimple != nil {
		return obj.AuthenticatorMethodSimple
	}

	if obj.AuthenticatorMethodTotp != nil {
		return obj.AuthenticatorMethodTotp
	}

	if obj.AuthenticatorMethodWebAuthn != nil {
		return obj.AuthenticatorMethodWebAuthn
	}

	if obj.AuthenticatorMethodWithVerifiableProperties != nil {
		return obj.AuthenticatorMethodWithVerifiableProperties
	}

	// all schemas are nil
	return nil
}

type NullableListAuthenticatorMethods200ResponseInner struct {
	value *ListAuthenticatorMethods200ResponseInner
	isSet bool
}

func (v NullableListAuthenticatorMethods200ResponseInner) Get() *ListAuthenticatorMethods200ResponseInner {
	return v.value
}

func (v *NullableListAuthenticatorMethods200ResponseInner) Set(val *ListAuthenticatorMethods200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuthenticatorMethods200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuthenticatorMethods200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuthenticatorMethods200ResponseInner(val *ListAuthenticatorMethods200ResponseInner) *NullableListAuthenticatorMethods200ResponseInner {
	return &NullableListAuthenticatorMethods200ResponseInner{value: val, isSet: true}
}

func (v NullableListAuthenticatorMethods200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuthenticatorMethods200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


