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
// ListPolicyRules200ResponseInner - struct for ListPolicyRules200ResponseInner
type ListPolicyRules200ResponseInner struct {
	AccessPolicyRule *AccessPolicyRule
	AuthorizationServerPolicyRule *AuthorizationServerPolicyRule
	ContinuousAccessPolicyRule *ContinuousAccessPolicyRule
	EntityRiskPolicyRule *EntityRiskPolicyRule
	IdpDiscoveryPolicyRule *IdpDiscoveryPolicyRule
	OktaSignOnPolicyRule *OktaSignOnPolicyRule
	PasswordPolicyRule *PasswordPolicyRule
	ProfileEnrollmentPolicyRule *ProfileEnrollmentPolicyRule
}

// AccessPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns AccessPolicyRule wrapped in ListPolicyRules200ResponseInner
func AccessPolicyRuleAsListPolicyRules200ResponseInner(v *AccessPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		AccessPolicyRule: v,
	}
}

// AuthorizationServerPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns AuthorizationServerPolicyRule wrapped in ListPolicyRules200ResponseInner
func AuthorizationServerPolicyRuleAsListPolicyRules200ResponseInner(v *AuthorizationServerPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		AuthorizationServerPolicyRule: v,
	}
}

// ContinuousAccessPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns ContinuousAccessPolicyRule wrapped in ListPolicyRules200ResponseInner
func ContinuousAccessPolicyRuleAsListPolicyRules200ResponseInner(v *ContinuousAccessPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		ContinuousAccessPolicyRule: v,
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

// ProfileEnrollmentPolicyRuleAsListPolicyRules200ResponseInner is a convenience function that returns ProfileEnrollmentPolicyRule wrapped in ListPolicyRules200ResponseInner
func ProfileEnrollmentPolicyRuleAsListPolicyRules200ResponseInner(v *ProfileEnrollmentPolicyRule) ListPolicyRules200ResponseInner {
	return ListPolicyRules200ResponseInner{
		ProfileEnrollmentPolicyRule: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListPolicyRules200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ACCESS_POLICY'
	if jsonDict["type"] == "ACCESS_POLICY" {
		// try to unmarshal JSON data into AccessPolicyRule
		err = json.Unmarshal(data, &dst.AccessPolicyRule)
		if err == nil {
			return nil // data stored in dst.AccessPolicyRule, return on the first match
		} else {
			dst.AccessPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as AccessPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AccessPolicyRule'
	if jsonDict["type"] == "AccessPolicyRule" {
		// try to unmarshal JSON data into AccessPolicyRule
		err = json.Unmarshal(data, &dst.AccessPolicyRule)
		if err == nil {
			return nil // data stored in dst.AccessPolicyRule, return on the first match
		} else {
			dst.AccessPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as AccessPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthorizationServerPolicyRule'
	if jsonDict["type"] == "AuthorizationServerPolicyRule" {
		// try to unmarshal JSON data into AuthorizationServerPolicyRule
		err = json.Unmarshal(data, &dst.AuthorizationServerPolicyRule)
		if err == nil {
			return nil // data stored in dst.AuthorizationServerPolicyRule, return on the first match
		} else {
			dst.AuthorizationServerPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as AuthorizationServerPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CONTINUOUS_ACCESS'
	if jsonDict["type"] == "CONTINUOUS_ACCESS" {
		// try to unmarshal JSON data into ContinuousAccessPolicyRule
		err = json.Unmarshal(data, &dst.ContinuousAccessPolicyRule)
		if err == nil {
			return nil // data stored in dst.ContinuousAccessPolicyRule, return on the first match
		} else {
			dst.ContinuousAccessPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as ContinuousAccessPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ContinuousAccessPolicyRule'
	if jsonDict["type"] == "ContinuousAccessPolicyRule" {
		// try to unmarshal JSON data into ContinuousAccessPolicyRule
		err = json.Unmarshal(data, &dst.ContinuousAccessPolicyRule)
		if err == nil {
			return nil // data stored in dst.ContinuousAccessPolicyRule, return on the first match
		} else {
			dst.ContinuousAccessPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as ContinuousAccessPolicyRule: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as EntityRiskPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EntityRiskPolicyRule'
	if jsonDict["type"] == "EntityRiskPolicyRule" {
		// try to unmarshal JSON data into EntityRiskPolicyRule
		err = json.Unmarshal(data, &dst.EntityRiskPolicyRule)
		if err == nil {
			return nil // data stored in dst.EntityRiskPolicyRule, return on the first match
		} else {
			dst.EntityRiskPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as EntityRiskPolicyRule: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as IdpDiscoveryPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IdpDiscoveryPolicyRule'
	if jsonDict["type"] == "IdpDiscoveryPolicyRule" {
		// try to unmarshal JSON data into IdpDiscoveryPolicyRule
		err = json.Unmarshal(data, &dst.IdpDiscoveryPolicyRule)
		if err == nil {
			return nil // data stored in dst.IdpDiscoveryPolicyRule, return on the first match
		} else {
			dst.IdpDiscoveryPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as IdpDiscoveryPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OktaSignOnPolicyRule'
	if jsonDict["type"] == "OktaSignOnPolicyRule" {
		// try to unmarshal JSON data into OktaSignOnPolicyRule
		err = json.Unmarshal(data, &dst.OktaSignOnPolicyRule)
		if err == nil {
			return nil // data stored in dst.OktaSignOnPolicyRule, return on the first match
		} else {
			dst.OktaSignOnPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as OktaSignOnPolicyRule: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as PasswordPolicyRule: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as ProfileEnrollmentPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PasswordPolicyRule'
	if jsonDict["type"] == "PasswordPolicyRule" {
		// try to unmarshal JSON data into PasswordPolicyRule
		err = json.Unmarshal(data, &dst.PasswordPolicyRule)
		if err == nil {
			return nil // data stored in dst.PasswordPolicyRule, return on the first match
		} else {
			dst.PasswordPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as PasswordPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ProfileEnrollmentPolicyRule'
	if jsonDict["type"] == "ProfileEnrollmentPolicyRule" {
		// try to unmarshal JSON data into ProfileEnrollmentPolicyRule
		err = json.Unmarshal(data, &dst.ProfileEnrollmentPolicyRule)
		if err == nil {
			return nil // data stored in dst.ProfileEnrollmentPolicyRule, return on the first match
		} else {
			dst.ProfileEnrollmentPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as ProfileEnrollmentPolicyRule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESOURCE_ACCESS'
	if jsonDict["type"] == "RESOURCE_ACCESS" {
		// try to unmarshal JSON data into AuthorizationServerPolicyRule
		err = json.Unmarshal(data, &dst.AuthorizationServerPolicyRule)
		if err == nil {
			return nil // data stored in dst.AuthorizationServerPolicyRule, return on the first match
		} else {
			dst.AuthorizationServerPolicyRule = nil
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as AuthorizationServerPolicyRule: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ListPolicyRules200ResponseInner as OktaSignOnPolicyRule: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListPolicyRules200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AccessPolicyRule != nil {
		return json.Marshal(&src.AccessPolicyRule)
	}

	if src.AuthorizationServerPolicyRule != nil {
		return json.Marshal(&src.AuthorizationServerPolicyRule)
	}

	if src.ContinuousAccessPolicyRule != nil {
		return json.Marshal(&src.ContinuousAccessPolicyRule)
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

	if src.ProfileEnrollmentPolicyRule != nil {
		return json.Marshal(&src.ProfileEnrollmentPolicyRule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListPolicyRules200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessPolicyRule != nil {
		return obj.AccessPolicyRule
	}

	if obj.AuthorizationServerPolicyRule != nil {
		return obj.AuthorizationServerPolicyRule
	}

	if obj.ContinuousAccessPolicyRule != nil {
		return obj.ContinuousAccessPolicyRule
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

	if obj.ProfileEnrollmentPolicyRule != nil {
		return obj.ProfileEnrollmentPolicyRule
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


