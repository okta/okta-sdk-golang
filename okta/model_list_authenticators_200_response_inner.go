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
// ListAuthenticators200ResponseInner - struct for ListAuthenticators200ResponseInner
type ListAuthenticators200ResponseInner struct {
	AuthenticatorKeyCustomApp *AuthenticatorKeyCustomApp
	AuthenticatorKeyDuo *AuthenticatorKeyDuo
	AuthenticatorKeyEmail *AuthenticatorKeyEmail
	AuthenticatorKeyExternalIdp *AuthenticatorKeyExternalIdp
	AuthenticatorKeyGoogleOtp *AuthenticatorKeyGoogleOtp
	AuthenticatorKeyOktaVerify *AuthenticatorKeyOktaVerify
	AuthenticatorKeyOnprem *AuthenticatorKeyOnprem
	AuthenticatorKeyPassword *AuthenticatorKeyPassword
	AuthenticatorKeyPhone *AuthenticatorKeyPhone
	AuthenticatorKeySecurityKey *AuthenticatorKeySecurityKey
	AuthenticatorKeySecurityQuestion *AuthenticatorKeySecurityQuestion
	AuthenticatorKeySmartCard *AuthenticatorKeySmartCard
	AuthenticatorKeySymantecVip *AuthenticatorKeySymantecVip
	AuthenticatorKeyWebauthn *AuthenticatorKeyWebauthn
	AuthenticatorKeyYubikey *AuthenticatorKeyYubikey
}

// AuthenticatorKeyCustomAppAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyCustomApp wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyCustomAppAsListAuthenticators200ResponseInner(v *AuthenticatorKeyCustomApp) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyCustomApp: v,
	}
}

// AuthenticatorKeyDuoAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyDuo wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyDuoAsListAuthenticators200ResponseInner(v *AuthenticatorKeyDuo) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyDuo: v,
	}
}

// AuthenticatorKeyEmailAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyEmail wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyEmailAsListAuthenticators200ResponseInner(v *AuthenticatorKeyEmail) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyEmail: v,
	}
}

// AuthenticatorKeyExternalIdpAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyExternalIdp wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyExternalIdpAsListAuthenticators200ResponseInner(v *AuthenticatorKeyExternalIdp) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyExternalIdp: v,
	}
}

// AuthenticatorKeyGoogleOtpAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyGoogleOtp wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyGoogleOtpAsListAuthenticators200ResponseInner(v *AuthenticatorKeyGoogleOtp) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyGoogleOtp: v,
	}
}

// AuthenticatorKeyOktaVerifyAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyOktaVerify wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyOktaVerifyAsListAuthenticators200ResponseInner(v *AuthenticatorKeyOktaVerify) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyOktaVerify: v,
	}
}

// AuthenticatorKeyOnpremAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyOnprem wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyOnpremAsListAuthenticators200ResponseInner(v *AuthenticatorKeyOnprem) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyOnprem: v,
	}
}

// AuthenticatorKeyPasswordAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyPassword wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyPasswordAsListAuthenticators200ResponseInner(v *AuthenticatorKeyPassword) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyPassword: v,
	}
}

// AuthenticatorKeyPhoneAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyPhone wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyPhoneAsListAuthenticators200ResponseInner(v *AuthenticatorKeyPhone) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyPhone: v,
	}
}

// AuthenticatorKeySecurityKeyAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeySecurityKey wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeySecurityKeyAsListAuthenticators200ResponseInner(v *AuthenticatorKeySecurityKey) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeySecurityKey: v,
	}
}

// AuthenticatorKeySecurityQuestionAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeySecurityQuestion wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeySecurityQuestionAsListAuthenticators200ResponseInner(v *AuthenticatorKeySecurityQuestion) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeySecurityQuestion: v,
	}
}

// AuthenticatorKeySmartCardAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeySmartCard wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeySmartCardAsListAuthenticators200ResponseInner(v *AuthenticatorKeySmartCard) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeySmartCard: v,
	}
}

// AuthenticatorKeySymantecVipAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeySymantecVip wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeySymantecVipAsListAuthenticators200ResponseInner(v *AuthenticatorKeySymantecVip) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeySymantecVip: v,
	}
}

// AuthenticatorKeyWebauthnAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyWebauthn wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyWebauthnAsListAuthenticators200ResponseInner(v *AuthenticatorKeyWebauthn) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyWebauthn: v,
	}
}

// AuthenticatorKeyYubikeyAsListAuthenticators200ResponseInner is a convenience function that returns AuthenticatorKeyYubikey wrapped in ListAuthenticators200ResponseInner
func AuthenticatorKeyYubikeyAsListAuthenticators200ResponseInner(v *AuthenticatorKeyYubikey) ListAuthenticators200ResponseInner {
	return ListAuthenticators200ResponseInner{
		AuthenticatorKeyYubikey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListAuthenticators200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AuthenticatorKeyCustomApp'
	if jsonDict["key"] == "AuthenticatorKeyCustomApp" {
		// try to unmarshal JSON data into AuthenticatorKeyCustomApp
		err = json.Unmarshal(data, &dst.AuthenticatorKeyCustomApp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyCustomApp, return on the first match
		} else {
			dst.AuthenticatorKeyCustomApp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyCustomApp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyDuo'
	if jsonDict["key"] == "AuthenticatorKeyDuo" {
		// try to unmarshal JSON data into AuthenticatorKeyDuo
		err = json.Unmarshal(data, &dst.AuthenticatorKeyDuo)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyDuo, return on the first match
		} else {
			dst.AuthenticatorKeyDuo = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyDuo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyEmail'
	if jsonDict["key"] == "AuthenticatorKeyEmail" {
		// try to unmarshal JSON data into AuthenticatorKeyEmail
		err = json.Unmarshal(data, &dst.AuthenticatorKeyEmail)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyEmail, return on the first match
		} else {
			dst.AuthenticatorKeyEmail = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyEmail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyExternalIdp'
	if jsonDict["key"] == "AuthenticatorKeyExternalIdp" {
		// try to unmarshal JSON data into AuthenticatorKeyExternalIdp
		err = json.Unmarshal(data, &dst.AuthenticatorKeyExternalIdp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyExternalIdp, return on the first match
		} else {
			dst.AuthenticatorKeyExternalIdp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyExternalIdp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyGoogleOtp'
	if jsonDict["key"] == "AuthenticatorKeyGoogleOtp" {
		// try to unmarshal JSON data into AuthenticatorKeyGoogleOtp
		err = json.Unmarshal(data, &dst.AuthenticatorKeyGoogleOtp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyGoogleOtp, return on the first match
		} else {
			dst.AuthenticatorKeyGoogleOtp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyGoogleOtp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyOktaVerify'
	if jsonDict["key"] == "AuthenticatorKeyOktaVerify" {
		// try to unmarshal JSON data into AuthenticatorKeyOktaVerify
		err = json.Unmarshal(data, &dst.AuthenticatorKeyOktaVerify)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyOktaVerify, return on the first match
		} else {
			dst.AuthenticatorKeyOktaVerify = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyOktaVerify: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyOnprem'
	if jsonDict["key"] == "AuthenticatorKeyOnprem" {
		// try to unmarshal JSON data into AuthenticatorKeyOnprem
		err = json.Unmarshal(data, &dst.AuthenticatorKeyOnprem)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyOnprem, return on the first match
		} else {
			dst.AuthenticatorKeyOnprem = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyOnprem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyPassword'
	if jsonDict["key"] == "AuthenticatorKeyPassword" {
		// try to unmarshal JSON data into AuthenticatorKeyPassword
		err = json.Unmarshal(data, &dst.AuthenticatorKeyPassword)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyPassword, return on the first match
		} else {
			dst.AuthenticatorKeyPassword = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyPassword: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyPhone'
	if jsonDict["key"] == "AuthenticatorKeyPhone" {
		// try to unmarshal JSON data into AuthenticatorKeyPhone
		err = json.Unmarshal(data, &dst.AuthenticatorKeyPhone)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyPhone, return on the first match
		} else {
			dst.AuthenticatorKeyPhone = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyPhone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeySecurityKey'
	if jsonDict["key"] == "AuthenticatorKeySecurityKey" {
		// try to unmarshal JSON data into AuthenticatorKeySecurityKey
		err = json.Unmarshal(data, &dst.AuthenticatorKeySecurityKey)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySecurityKey, return on the first match
		} else {
			dst.AuthenticatorKeySecurityKey = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySecurityKey: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeySecurityQuestion'
	if jsonDict["key"] == "AuthenticatorKeySecurityQuestion" {
		// try to unmarshal JSON data into AuthenticatorKeySecurityQuestion
		err = json.Unmarshal(data, &dst.AuthenticatorKeySecurityQuestion)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySecurityQuestion, return on the first match
		} else {
			dst.AuthenticatorKeySecurityQuestion = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySecurityQuestion: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeySmartCard'
	if jsonDict["key"] == "AuthenticatorKeySmartCard" {
		// try to unmarshal JSON data into AuthenticatorKeySmartCard
		err = json.Unmarshal(data, &dst.AuthenticatorKeySmartCard)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySmartCard, return on the first match
		} else {
			dst.AuthenticatorKeySmartCard = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySmartCard: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeySymantecVip'
	if jsonDict["key"] == "AuthenticatorKeySymantecVip" {
		// try to unmarshal JSON data into AuthenticatorKeySymantecVip
		err = json.Unmarshal(data, &dst.AuthenticatorKeySymantecVip)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySymantecVip, return on the first match
		} else {
			dst.AuthenticatorKeySymantecVip = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySymantecVip: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyWebauthn'
	if jsonDict["key"] == "AuthenticatorKeyWebauthn" {
		// try to unmarshal JSON data into AuthenticatorKeyWebauthn
		err = json.Unmarshal(data, &dst.AuthenticatorKeyWebauthn)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyWebauthn, return on the first match
		} else {
			dst.AuthenticatorKeyWebauthn = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyWebauthn: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorKeyYubikey'
	if jsonDict["key"] == "AuthenticatorKeyYubikey" {
		// try to unmarshal JSON data into AuthenticatorKeyYubikey
		err = json.Unmarshal(data, &dst.AuthenticatorKeyYubikey)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyYubikey, return on the first match
		} else {
			dst.AuthenticatorKeyYubikey = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyYubikey: %s", err.Error())
		}
	}

	// check if the discriminator value is 'custom_app'
	if jsonDict["key"] == "custom_app" {
		// try to unmarshal JSON data into AuthenticatorKeyCustomApp
		err = json.Unmarshal(data, &dst.AuthenticatorKeyCustomApp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyCustomApp, return on the first match
		} else {
			dst.AuthenticatorKeyCustomApp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyCustomApp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'duo'
	if jsonDict["key"] == "duo" {
		// try to unmarshal JSON data into AuthenticatorKeyDuo
		err = json.Unmarshal(data, &dst.AuthenticatorKeyDuo)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyDuo, return on the first match
		} else {
			dst.AuthenticatorKeyDuo = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyDuo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'external_idp'
	if jsonDict["key"] == "external_idp" {
		// try to unmarshal JSON data into AuthenticatorKeyExternalIdp
		err = json.Unmarshal(data, &dst.AuthenticatorKeyExternalIdp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyExternalIdp, return on the first match
		} else {
			dst.AuthenticatorKeyExternalIdp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyExternalIdp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'google_otp'
	if jsonDict["key"] == "google_otp" {
		// try to unmarshal JSON data into AuthenticatorKeyGoogleOtp
		err = json.Unmarshal(data, &dst.AuthenticatorKeyGoogleOtp)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyGoogleOtp, return on the first match
		} else {
			dst.AuthenticatorKeyGoogleOtp = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyGoogleOtp: %s", err.Error())
		}
	}

	// check if the discriminator value is 'okta_email'
	if jsonDict["key"] == "okta_email" {
		// try to unmarshal JSON data into AuthenticatorKeyEmail
		err = json.Unmarshal(data, &dst.AuthenticatorKeyEmail)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyEmail, return on the first match
		} else {
			dst.AuthenticatorKeyEmail = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyEmail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'okta_password'
	if jsonDict["key"] == "okta_password" {
		// try to unmarshal JSON data into AuthenticatorKeyPassword
		err = json.Unmarshal(data, &dst.AuthenticatorKeyPassword)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyPassword, return on the first match
		} else {
			dst.AuthenticatorKeyPassword = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyPassword: %s", err.Error())
		}
	}

	// check if the discriminator value is 'okta_verify'
	if jsonDict["key"] == "okta_verify" {
		// try to unmarshal JSON data into AuthenticatorKeyOktaVerify
		err = json.Unmarshal(data, &dst.AuthenticatorKeyOktaVerify)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyOktaVerify, return on the first match
		} else {
			dst.AuthenticatorKeyOktaVerify = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyOktaVerify: %s", err.Error())
		}
	}

	// check if the discriminator value is 'onprem_mfa'
	if jsonDict["key"] == "onprem_mfa" {
		// try to unmarshal JSON data into AuthenticatorKeyOnprem
		err = json.Unmarshal(data, &dst.AuthenticatorKeyOnprem)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyOnprem, return on the first match
		} else {
			dst.AuthenticatorKeyOnprem = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyOnprem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'phone_number'
	if jsonDict["key"] == "phone_number" {
		// try to unmarshal JSON data into AuthenticatorKeyPhone
		err = json.Unmarshal(data, &dst.AuthenticatorKeyPhone)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyPhone, return on the first match
		} else {
			dst.AuthenticatorKeyPhone = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyPhone: %s", err.Error())
		}
	}

	// check if the discriminator value is 'security_key'
	if jsonDict["key"] == "security_key" {
		// try to unmarshal JSON data into AuthenticatorKeySecurityKey
		err = json.Unmarshal(data, &dst.AuthenticatorKeySecurityKey)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySecurityKey, return on the first match
		} else {
			dst.AuthenticatorKeySecurityKey = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySecurityKey: %s", err.Error())
		}
	}

	// check if the discriminator value is 'security_question'
	if jsonDict["key"] == "security_question" {
		// try to unmarshal JSON data into AuthenticatorKeySecurityQuestion
		err = json.Unmarshal(data, &dst.AuthenticatorKeySecurityQuestion)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySecurityQuestion, return on the first match
		} else {
			dst.AuthenticatorKeySecurityQuestion = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySecurityQuestion: %s", err.Error())
		}
	}

	// check if the discriminator value is 'smart_card_idp'
	if jsonDict["key"] == "smart_card_idp" {
		// try to unmarshal JSON data into AuthenticatorKeySmartCard
		err = json.Unmarshal(data, &dst.AuthenticatorKeySmartCard)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySmartCard, return on the first match
		} else {
			dst.AuthenticatorKeySmartCard = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySmartCard: %s", err.Error())
		}
	}

	// check if the discriminator value is 'symantec_vip'
	if jsonDict["key"] == "symantec_vip" {
		// try to unmarshal JSON data into AuthenticatorKeySymantecVip
		err = json.Unmarshal(data, &dst.AuthenticatorKeySymantecVip)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeySymantecVip, return on the first match
		} else {
			dst.AuthenticatorKeySymantecVip = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeySymantecVip: %s", err.Error())
		}
	}

	// check if the discriminator value is 'webauthn'
	if jsonDict["key"] == "webauthn" {
		// try to unmarshal JSON data into AuthenticatorKeyWebauthn
		err = json.Unmarshal(data, &dst.AuthenticatorKeyWebauthn)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyWebauthn, return on the first match
		} else {
			dst.AuthenticatorKeyWebauthn = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyWebauthn: %s", err.Error())
		}
	}

	// check if the discriminator value is 'yubikey_token'
	if jsonDict["key"] == "yubikey_token" {
		// try to unmarshal JSON data into AuthenticatorKeyYubikey
		err = json.Unmarshal(data, &dst.AuthenticatorKeyYubikey)
		if err == nil {
			return nil // data stored in dst.AuthenticatorKeyYubikey, return on the first match
		} else {
			dst.AuthenticatorKeyYubikey = nil
			return fmt.Errorf("Failed to unmarshal ListAuthenticators200ResponseInner as AuthenticatorKeyYubikey: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAuthenticators200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AuthenticatorKeyCustomApp != nil {
		return json.Marshal(&src.AuthenticatorKeyCustomApp)
	}

	if src.AuthenticatorKeyDuo != nil {
		return json.Marshal(&src.AuthenticatorKeyDuo)
	}

	if src.AuthenticatorKeyEmail != nil {
		return json.Marshal(&src.AuthenticatorKeyEmail)
	}

	if src.AuthenticatorKeyExternalIdp != nil {
		return json.Marshal(&src.AuthenticatorKeyExternalIdp)
	}

	if src.AuthenticatorKeyGoogleOtp != nil {
		return json.Marshal(&src.AuthenticatorKeyGoogleOtp)
	}

	if src.AuthenticatorKeyOktaVerify != nil {
		return json.Marshal(&src.AuthenticatorKeyOktaVerify)
	}

	if src.AuthenticatorKeyOnprem != nil {
		return json.Marshal(&src.AuthenticatorKeyOnprem)
	}

	if src.AuthenticatorKeyPassword != nil {
		return json.Marshal(&src.AuthenticatorKeyPassword)
	}

	if src.AuthenticatorKeyPhone != nil {
		return json.Marshal(&src.AuthenticatorKeyPhone)
	}

	if src.AuthenticatorKeySecurityKey != nil {
		return json.Marshal(&src.AuthenticatorKeySecurityKey)
	}

	if src.AuthenticatorKeySecurityQuestion != nil {
		return json.Marshal(&src.AuthenticatorKeySecurityQuestion)
	}

	if src.AuthenticatorKeySmartCard != nil {
		return json.Marshal(&src.AuthenticatorKeySmartCard)
	}

	if src.AuthenticatorKeySymantecVip != nil {
		return json.Marshal(&src.AuthenticatorKeySymantecVip)
	}

	if src.AuthenticatorKeyWebauthn != nil {
		return json.Marshal(&src.AuthenticatorKeyWebauthn)
	}

	if src.AuthenticatorKeyYubikey != nil {
		return json.Marshal(&src.AuthenticatorKeyYubikey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAuthenticators200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuthenticatorKeyCustomApp != nil {
		return obj.AuthenticatorKeyCustomApp
	}

	if obj.AuthenticatorKeyDuo != nil {
		return obj.AuthenticatorKeyDuo
	}

	if obj.AuthenticatorKeyEmail != nil {
		return obj.AuthenticatorKeyEmail
	}

	if obj.AuthenticatorKeyExternalIdp != nil {
		return obj.AuthenticatorKeyExternalIdp
	}

	if obj.AuthenticatorKeyGoogleOtp != nil {
		return obj.AuthenticatorKeyGoogleOtp
	}

	if obj.AuthenticatorKeyOktaVerify != nil {
		return obj.AuthenticatorKeyOktaVerify
	}

	if obj.AuthenticatorKeyOnprem != nil {
		return obj.AuthenticatorKeyOnprem
	}

	if obj.AuthenticatorKeyPassword != nil {
		return obj.AuthenticatorKeyPassword
	}

	if obj.AuthenticatorKeyPhone != nil {
		return obj.AuthenticatorKeyPhone
	}

	if obj.AuthenticatorKeySecurityKey != nil {
		return obj.AuthenticatorKeySecurityKey
	}

	if obj.AuthenticatorKeySecurityQuestion != nil {
		return obj.AuthenticatorKeySecurityQuestion
	}

	if obj.AuthenticatorKeySmartCard != nil {
		return obj.AuthenticatorKeySmartCard
	}

	if obj.AuthenticatorKeySymantecVip != nil {
		return obj.AuthenticatorKeySymantecVip
	}

	if obj.AuthenticatorKeyWebauthn != nil {
		return obj.AuthenticatorKeyWebauthn
	}

	if obj.AuthenticatorKeyYubikey != nil {
		return obj.AuthenticatorKeyYubikey
	}

	// all schemas are nil
	return nil
}

type NullableListAuthenticators200ResponseInner struct {
	value *ListAuthenticators200ResponseInner
	isSet bool
}

func (v NullableListAuthenticators200ResponseInner) Get() *ListAuthenticators200ResponseInner {
	return v.value
}

func (v *NullableListAuthenticators200ResponseInner) Set(val *ListAuthenticators200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuthenticators200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuthenticators200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuthenticators200ResponseInner(val *ListAuthenticators200ResponseInner) *NullableListAuthenticators200ResponseInner {
	return &NullableListAuthenticators200ResponseInner{value: val, isSet: true}
}

func (v NullableListAuthenticators200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuthenticators200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


