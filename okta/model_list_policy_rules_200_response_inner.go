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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// ListPolicyRules200ResponseInner - struct for ListPolicyRules200ResponseInner
type ListPolicyRules200ResponseInner struct {
	AccessPolicyRule                  *AccessPolicyRule
	AuthenticatorEnrollmentPolicyRule *AuthenticatorEnrollmentPolicyRule
	DeviceSignalCollectionPolicyRule  *DeviceSignalCollectionPolicyRule
	EntityRiskPolicyRule              *EntityRiskPolicyRule
	IdpDiscoveryPolicyRule            *IdpDiscoveryPolicyRule
	OktaSignOnPolicyRule              *OktaSignOnPolicyRule
	PasswordPolicyRule                *PasswordPolicyRule
	PostAuthSessionPolicyRule         *PostAuthSessionPolicyRule
	ProfileEnrollmentPolicyRule       *ProfileEnrollmentPolicyRule
}

// AccessPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns AccessPolicyRule wrapped in ListPolicyRules200ResponseInner
func AccessPolicyRuleAsListPolicyRules200ResponseInner(v *AccessPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		AccessPolicyRule: v,
	}
}

// AuthenticatorEnrollmentPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns AuthenticatorEnrollmentPolicyRule wrapped in ListPolicyRules200ResponseInner
func AuthenticatorEnrollmentPolicyRuleAsListPolicyRules200ResponseInner(v *AuthenticatorEnrollmentPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		AuthenticatorEnrollmentPolicyRule: v,
	}
}

// DeviceSignalCollectionPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns DeviceSignalCollectionPolicyRule wrapped in ListPolicyRules200ResponseInner
func DeviceSignalCollectionPolicyRuleAsListPolicyRules200ResponseInner(v *DeviceSignalCollectionPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		DeviceSignalCollectionPolicyRule: v,
	}
}

// EntityRiskPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns EntityRiskPolicyRule wrapped in ListPolicyRules200ResponseInner
func EntityRiskPolicyRuleAsListPolicyRules200ResponseInner(v *EntityRiskPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		EntityRiskPolicyRule: v,
	}
}

// IdpDiscoveryPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns IdpDiscoveryPolicyRule wrapped in ListPolicyRules200ResponseInner
func IdpDiscoveryPolicyRuleAsListPolicyRules200ResponseInner(v *IdpDiscoveryPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		IdpDiscoveryPolicyRule: v,
	}
}

// OktaSignOnPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns OktaSignOnPolicyRule wrapped in ListPolicyRules200ResponseInner
func OktaSignOnPolicyRuleAsListPolicyRules200ResponseInner(v *OktaSignOnPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		OktaSignOnPolicyRule: v,
	}
}

// PasswordPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns PasswordPolicyRule wrapped in ListPolicyRules200ResponseInner
func PasswordPolicyRuleAsListPolicyRules200ResponseInner(v *PasswordPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		PasswordPolicyRule: v,
	}
}

// PostAuthSessionPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns PostAuthSessionPolicyRule wrapped in ListPolicyRules200ResponseInner
func PostAuthSessionPolicyRuleAsListPolicyRules200ResponseInner(v *PostAuthSessionPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		PostAuthSessionPolicyRule: v,
	}
}

// ProfileEnrollmentPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns ProfileEnrollmentPolicyRule wrapped in ListPolicyRules200ResponseInner
func ProfileEnrollmentPolicyRuleAsListPolicyRules200ResponseInner(v *ProfileEnrollmentPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		ProfileEnrollmentPolicyRule: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListPolicyRules200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ACCESS_POLICY'
	if jsonDict["type"] == "ACCESS_POLICY" {
		// try to unmarshal JSON data into AccessPolicyRule
		err = json.Unmarshal(data, &dst.AccessPolicyRule)
		if err == nil {
			return nil // data stored in dst.AccessPolicyRule, return on the first match
		} else {
			dst.AccessPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as AccessPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DEVICE_SIGNAL_COLLECTION'
	if jsonDict["type"] == "DEVICE_SIGNAL_COLLECTION" {
		// try to unmarshal JSON data into DeviceSignalCollectionPolicyRule
		err = json.Unmarshal(data, &dst.DeviceSignalCollectionPolicyRule)
		if err == nil {
			return nil // data stored in dst.DeviceSignalCollectionPolicyRule, return on the first match
		} else {
			dst.DeviceSignalCollectionPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as DeviceSignalCollectionPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITY_RISK'
	if jsonDict["type"] == "ENTITY_RISK" {
		// try to unmarshal JSON data into EntityRiskPolicyRule
		err = json.Unmarshal(data, &dst.EntityRiskPolicyRule)
		if err == nil {
			return nil // data stored in dst.EntityRiskPolicyRule, return on the first match
		} else {
			dst.EntityRiskPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as EntityRiskPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IDP_DISCOVERY'
	if jsonDict["type"] == "IDP_DISCOVERY" {
		// try to unmarshal JSON data into IdpDiscoveryPolicyRule
		err = json.Unmarshal(data, &dst.IdpDiscoveryPolicyRule)
		if err == nil {
			return nil // data stored in dst.IdpDiscoveryPolicyRule, return on the first match
		} else {
			dst.IdpDiscoveryPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as IdpDiscoveryPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MFA_ENROLL'
	if jsonDict["type"] == "MFA_ENROLL" {
		// try to unmarshal JSON data into AuthenticatorEnrollmentPolicyRule
		err = json.Unmarshal(data, &dst.AuthenticatorEnrollmentPolicyRule)
		if err == nil {
			return nil // data stored in dst.AuthenticatorEnrollmentPolicyRule, return on the first match
		} else {
			dst.AuthenticatorEnrollmentPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as AuthenticatorEnrollmentPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PASSWORD'
	if jsonDict["type"] == "PASSWORD" {
		// try to unmarshal JSON data into PasswordPolicyRule
		err = json.Unmarshal(data, &dst.PasswordPolicyRule)
		if err == nil {
			return nil // data stored in dst.PasswordPolicyRule, return on the first match
		} else {
			dst.PasswordPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as PasswordPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'POST_AUTH_SESSION'
	if jsonDict["type"] == "POST_AUTH_SESSION" {
		// try to unmarshal JSON data into PostAuthSessionPolicyRule
		err = json.Unmarshal(data, &dst.PostAuthSessionPolicyRule)
		if err == nil {
			return nil // data stored in dst.PostAuthSessionPolicyRule, return on the first match
		} else {
			dst.PostAuthSessionPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as PostAuthSessionPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PROFILE_ENROLLMENT'
	if jsonDict["type"] == "PROFILE_ENROLLMENT" {
		// try to unmarshal JSON data into ProfileEnrollmentPolicyRule
		err = json.Unmarshal(data, &dst.ProfileEnrollmentPolicyRule)
		if err == nil {
			return nil // data stored in dst.ProfileEnrollmentPolicyRule, return on the first match
		} else {
			dst.ProfileEnrollmentPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as ProfileEnrollmentPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SIGN_ON'
	if jsonDict["type"] == "SIGN_ON" {
		// try to unmarshal JSON data into OktaSignOnPolicyRule
		err = json.Unmarshal(data, &dst.OktaSignOnPolicyRule)
		if err == nil {
			return nil // data stored in dst.OktaSignOnPolicyRule, return on the first match
		} else {
			dst.OktaSignOnPolicyRule = nil
			return fmt.Errorf("failed to unmarshal ListPolicyRules200ResponseInner as OktaSignOnPolicyRule: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListPolicyRules200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AccessPolicyRule != nil {
		return json.Marshal(&src.AccessPolicyRule)
	}

	if src.AuthenticatorEnrollmentPolicyRule != nil {
		return json.Marshal(&src.AuthenticatorEnrollmentPolicyRule)
	}

	if src.DeviceSignalCollectionPolicyRule != nil {
		return json.Marshal(&src.DeviceSignalCollectionPolicyRule)
	}

	if src.EntityRiskPolicyRule != nil {
		return json.Marshal(&src.EntityRiskPolicyRule)
	}

	if src.IdpDiscoveryPolicyRule != nil {
		return json.Marshal(&src.IdpDiscoveryPolicyRule)
	}

	if src.OktaSignOnPolicyRule != nil {
		return json.Marshal(&src.OktaSignOnPolicyRule)
	}

	if src.PasswordPolicyRule != nil {
		return json.Marshal(&src.PasswordPolicyRule)
	}

	if src.PostAuthSessionPolicyRule != nil {
		return json.Marshal(&src.PostAuthSessionPolicyRule)
	}

	if src.ProfileEnrollmentPolicyRule != nil {
		return json.Marshal(&src.ProfileEnrollmentPolicyRule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListPolicyRules200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessPolicyRule != nil {
		return obj.AccessPolicyRule
	}

	if obj.AuthenticatorEnrollmentPolicyRule != nil {
		return obj.AuthenticatorEnrollmentPolicyRule
	}

	if obj.DeviceSignalCollectionPolicyRule != nil {
		return obj.DeviceSignalCollectionPolicyRule
	}

	if obj.EntityRiskPolicyRule != nil {
		return obj.EntityRiskPolicyRule
	}

	if obj.IdpDiscoveryPolicyRule != nil {
		return obj.IdpDiscoveryPolicyRule
	}

	if obj.OktaSignOnPolicyRule != nil {
		return obj.OktaSignOnPolicyRule
	}

	if obj.PasswordPolicyRule != nil {
		return obj.PasswordPolicyRule
	}

	if obj.PostAuthSessionPolicyRule != nil {
		return obj.PostAuthSessionPolicyRule
	}

	if obj.ProfileEnrollmentPolicyRule != nil {
		return obj.ProfileEnrollmentPolicyRule
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListPolicyRules200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.AccessPolicyRule != nil {
		return *obj.AccessPolicyRule
	}

	if obj.AuthenticatorEnrollmentPolicyRule != nil {
		return *obj.AuthenticatorEnrollmentPolicyRule
	}

	if obj.DeviceSignalCollectionPolicyRule != nil {
		return *obj.DeviceSignalCollectionPolicyRule
	}

	if obj.EntityRiskPolicyRule != nil {
		return *obj.EntityRiskPolicyRule
	}

	if obj.IdpDiscoveryPolicyRule != nil {
		return *obj.IdpDiscoveryPolicyRule
	}

	if obj.OktaSignOnPolicyRule != nil {
		return *obj.OktaSignOnPolicyRule
	}

	if obj.PasswordPolicyRule != nil {
		return *obj.PasswordPolicyRule
	}

	if obj.PostAuthSessionPolicyRule != nil {
		return *obj.PostAuthSessionPolicyRule
	}

	if obj.ProfileEnrollmentPolicyRule != nil {
		return *obj.ProfileEnrollmentPolicyRule
	}

	// all schemas are nil
	return nil
}

type NullableListPolicyRules200ResponseInner struct {
	value *ListPolicyRules200ResponseInner
	isSet bool
}

func (v NullableListPolicyRules200ResponseInner) Get() *ListPolicyRules200ResponseInner {
	return v.value
}

func (v *NullableListPolicyRules200ResponseInner) Set(val *ListPolicyRules200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPolicyRules200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPolicyRules200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPolicyRules200ResponseInner(val *ListPolicyRules200ResponseInner) *NullableListPolicyRules200ResponseInner {
	return &NullableListPolicyRules200ResponseInner{value: val, isSet: true}
}

func (v NullableListPolicyRules200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPolicyRules200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
