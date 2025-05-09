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
// ListPolicies200ResponseInner - struct for ListPolicies200ResponseInner
type ListPolicies200ResponseInner struct {
	AccessPolicy *AccessPolicy
	ContinuousAccessPolicy *ContinuousAccessPolicy
	EntityRiskPolicy *EntityRiskPolicy
	IdpDiscoveryPolicy *IdpDiscoveryPolicy
	MultifactorEnrollmentPolicy *MultifactorEnrollmentPolicy
	OktaSignOnPolicy *OktaSignOnPolicy
	PasswordPolicy *PasswordPolicy
	ProfileEnrollmentPolicy *ProfileEnrollmentPolicy
}

// AccessPolicyAsListPolicies200ResponseInner is a convenience function that returns AccessPolicy wrapped in ListPolicies200ResponseInner
func AccessPolicyAsListPolicies200ResponseInner(v *AccessPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		AccessPolicy: v,
	}
}

// ContinuousAccessPolicyAsListPolicies200ResponseInner is a convenience function that returns ContinuousAccessPolicy wrapped in ListPolicies200ResponseInner
func ContinuousAccessPolicyAsListPolicies200ResponseInner(v *ContinuousAccessPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		ContinuousAccessPolicy: v,
	}
}

// EntityRiskPolicyAsListPolicies200ResponseInner is a convenience function that returns EntityRiskPolicy wrapped in ListPolicies200ResponseInner
func EntityRiskPolicyAsListPolicies200ResponseInner(v *EntityRiskPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		EntityRiskPolicy: v,
	}
}

// IdpDiscoveryPolicyAsListPolicies200ResponseInner is a convenience function that returns IdpDiscoveryPolicy wrapped in ListPolicies200ResponseInner
func IdpDiscoveryPolicyAsListPolicies200ResponseInner(v *IdpDiscoveryPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		IdpDiscoveryPolicy: v,
	}
}

// MultifactorEnrollmentPolicyAsListPolicies200ResponseInner is a convenience function that returns MultifactorEnrollmentPolicy wrapped in ListPolicies200ResponseInner
func MultifactorEnrollmentPolicyAsListPolicies200ResponseInner(v *MultifactorEnrollmentPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		MultifactorEnrollmentPolicy: v,
	}
}

// OktaSignOnPolicyAsListPolicies200ResponseInner is a convenience function that returns OktaSignOnPolicy wrapped in ListPolicies200ResponseInner
func OktaSignOnPolicyAsListPolicies200ResponseInner(v *OktaSignOnPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		OktaSignOnPolicy: v,
	}
}

// PasswordPolicyAsListPolicies200ResponseInner is a convenience function that returns PasswordPolicy wrapped in ListPolicies200ResponseInner
func PasswordPolicyAsListPolicies200ResponseInner(v *PasswordPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		PasswordPolicy: v,
	}
}

// ProfileEnrollmentPolicyAsListPolicies200ResponseInner is a convenience function that returns ProfileEnrollmentPolicy wrapped in ListPolicies200ResponseInner
func ProfileEnrollmentPolicyAsListPolicies200ResponseInner(v *ProfileEnrollmentPolicy) ListPolicies200ResponseInner {
	return ListPolicies200ResponseInner{
		ProfileEnrollmentPolicy: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListPolicies200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ACCESS_POLICY'
	if jsonDict["type"] == "ACCESS_POLICY" {
		// try to unmarshal JSON data into AccessPolicy
		err = json.Unmarshal(data, &dst.AccessPolicy)
		if err == nil {
			return nil // data stored in dst.AccessPolicy, return on the first match
		} else {
			dst.AccessPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as AccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AccessPolicy'
	if jsonDict["type"] == "AccessPolicy" {
		// try to unmarshal JSON data into AccessPolicy
		err = json.Unmarshal(data, &dst.AccessPolicy)
		if err == nil {
			return nil // data stored in dst.AccessPolicy, return on the first match
		} else {
			dst.AccessPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as AccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CONTINUOUS_ACCESS'
	if jsonDict["type"] == "CONTINUOUS_ACCESS" {
		// try to unmarshal JSON data into ContinuousAccessPolicy
		err = json.Unmarshal(data, &dst.ContinuousAccessPolicy)
		if err == nil {
			return nil // data stored in dst.ContinuousAccessPolicy, return on the first match
		} else {
			dst.ContinuousAccessPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as ContinuousAccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ContinuousAccessPolicy'
	if jsonDict["type"] == "ContinuousAccessPolicy" {
		// try to unmarshal JSON data into ContinuousAccessPolicy
		err = json.Unmarshal(data, &dst.ContinuousAccessPolicy)
		if err == nil {
			return nil // data stored in dst.ContinuousAccessPolicy, return on the first match
		} else {
			dst.ContinuousAccessPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as ContinuousAccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITY_RISK'
	if jsonDict["type"] == "ENTITY_RISK" {
		// try to unmarshal JSON data into EntityRiskPolicy
		err = json.Unmarshal(data, &dst.EntityRiskPolicy)
		if err == nil {
			return nil // data stored in dst.EntityRiskPolicy, return on the first match
		} else {
			dst.EntityRiskPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as EntityRiskPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EntityRiskPolicy'
	if jsonDict["type"] == "EntityRiskPolicy" {
		// try to unmarshal JSON data into EntityRiskPolicy
		err = json.Unmarshal(data, &dst.EntityRiskPolicy)
		if err == nil {
			return nil // data stored in dst.EntityRiskPolicy, return on the first match
		} else {
			dst.EntityRiskPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as EntityRiskPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IDP_DISCOVERY'
	if jsonDict["type"] == "IDP_DISCOVERY" {
		// try to unmarshal JSON data into IdpDiscoveryPolicy
		err = json.Unmarshal(data, &dst.IdpDiscoveryPolicy)
		if err == nil {
			return nil // data stored in dst.IdpDiscoveryPolicy, return on the first match
		} else {
			dst.IdpDiscoveryPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as IdpDiscoveryPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IdpDiscoveryPolicy'
	if jsonDict["type"] == "IdpDiscoveryPolicy" {
		// try to unmarshal JSON data into IdpDiscoveryPolicy
		err = json.Unmarshal(data, &dst.IdpDiscoveryPolicy)
		if err == nil {
			return nil // data stored in dst.IdpDiscoveryPolicy, return on the first match
		} else {
			dst.IdpDiscoveryPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as IdpDiscoveryPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MFA_ENROLL'
	if jsonDict["type"] == "MFA_ENROLL" {
		// try to unmarshal JSON data into MultifactorEnrollmentPolicy
		err = json.Unmarshal(data, &dst.MultifactorEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.MultifactorEnrollmentPolicy, return on the first match
		} else {
			dst.MultifactorEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as MultifactorEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MultifactorEnrollmentPolicy'
	if jsonDict["type"] == "MultifactorEnrollmentPolicy" {
		// try to unmarshal JSON data into MultifactorEnrollmentPolicy
		err = json.Unmarshal(data, &dst.MultifactorEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.MultifactorEnrollmentPolicy, return on the first match
		} else {
			dst.MultifactorEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as MultifactorEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OKTA_SIGN_ON'
	if jsonDict["type"] == "OKTA_SIGN_ON" {
		// try to unmarshal JSON data into OktaSignOnPolicy
		err = json.Unmarshal(data, &dst.OktaSignOnPolicy)
		if err == nil {
			return nil // data stored in dst.OktaSignOnPolicy, return on the first match
		} else {
			dst.OktaSignOnPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as OktaSignOnPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OktaSignOnPolicy'
	if jsonDict["type"] == "OktaSignOnPolicy" {
		// try to unmarshal JSON data into OktaSignOnPolicy
		err = json.Unmarshal(data, &dst.OktaSignOnPolicy)
		if err == nil {
			return nil // data stored in dst.OktaSignOnPolicy, return on the first match
		} else {
			dst.OktaSignOnPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as OktaSignOnPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PASSWORD'
	if jsonDict["type"] == "PASSWORD" {
		// try to unmarshal JSON data into PasswordPolicy
		err = json.Unmarshal(data, &dst.PasswordPolicy)
		if err == nil {
			return nil // data stored in dst.PasswordPolicy, return on the first match
		} else {
			dst.PasswordPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as PasswordPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PROFILE_ENROLLMENT'
	if jsonDict["type"] == "PROFILE_ENROLLMENT" {
		// try to unmarshal JSON data into ProfileEnrollmentPolicy
		err = json.Unmarshal(data, &dst.ProfileEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.ProfileEnrollmentPolicy, return on the first match
		} else {
			dst.ProfileEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as ProfileEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PasswordPolicy'
	if jsonDict["type"] == "PasswordPolicy" {
		// try to unmarshal JSON data into PasswordPolicy
		err = json.Unmarshal(data, &dst.PasswordPolicy)
		if err == nil {
			return nil // data stored in dst.PasswordPolicy, return on the first match
		} else {
			dst.PasswordPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as PasswordPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ProfileEnrollmentPolicy'
	if jsonDict["type"] == "ProfileEnrollmentPolicy" {
		// try to unmarshal JSON data into ProfileEnrollmentPolicy
		err = json.Unmarshal(data, &dst.ProfileEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.ProfileEnrollmentPolicy, return on the first match
		} else {
			dst.ProfileEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal ListPolicies200ResponseInner as ProfileEnrollmentPolicy: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AccessPolicy != nil {
		return json.Marshal(&src.AccessPolicy)
	}

	if src.ContinuousAccessPolicy != nil {
		return json.Marshal(&src.ContinuousAccessPolicy)
	}

	if src.EntityRiskPolicy != nil {
		return json.Marshal(&src.EntityRiskPolicy)
	}

	if src.IdpDiscoveryPolicy != nil {
		return json.Marshal(&src.IdpDiscoveryPolicy)
	}

	if src.MultifactorEnrollmentPolicy != nil {
		return json.Marshal(&src.MultifactorEnrollmentPolicy)
	}

	if src.OktaSignOnPolicy != nil {
		return json.Marshal(&src.OktaSignOnPolicy)
	}

	if src.PasswordPolicy != nil {
		return json.Marshal(&src.PasswordPolicy)
	}

	if src.ProfileEnrollmentPolicy != nil {
		return json.Marshal(&src.ProfileEnrollmentPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListPolicies200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccessPolicy != nil {
		return obj.AccessPolicy
	}

	if obj.ContinuousAccessPolicy != nil {
		return obj.ContinuousAccessPolicy
	}

	if obj.EntityRiskPolicy != nil {
		return obj.EntityRiskPolicy
	}

	if obj.IdpDiscoveryPolicy != nil {
		return obj.IdpDiscoveryPolicy
	}

	if obj.MultifactorEnrollmentPolicy != nil {
		return obj.MultifactorEnrollmentPolicy
	}

	if obj.OktaSignOnPolicy != nil {
		return obj.OktaSignOnPolicy
	}

	if obj.PasswordPolicy != nil {
		return obj.PasswordPolicy
	}

	if obj.ProfileEnrollmentPolicy != nil {
		return obj.ProfileEnrollmentPolicy
	}

	// all schemas are nil
	return nil
}

type NullableListPolicies200ResponseInner struct {
	value *ListPolicies200ResponseInner
	isSet bool
}

func (v NullableListPolicies200ResponseInner) Get() *ListPolicies200ResponseInner {
	return v.value
}

func (v *NullableListPolicies200ResponseInner) Set(val *ListPolicies200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPolicies200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPolicies200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPolicies200ResponseInner(val *ListPolicies200ResponseInner) *NullableListPolicies200ResponseInner {
	return &NullableListPolicies200ResponseInner{value: val, isSet: true}
}

func (v NullableListPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPolicies200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


