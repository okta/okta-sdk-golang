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

// ListPolicies200Response - struct for ListPolicies200Response
type ListPolicies200Response struct {
	AccessPolicy                  *AccessPolicy
	AuthenticatorEnrollmentPolicy *AuthenticatorEnrollmentPolicy
	DeviceSignalCollectionPolicy  *DeviceSignalCollectionPolicy
	EntityRiskPolicy              *EntityRiskPolicy
	IdpDiscoveryPolicy            *IdpDiscoveryPolicy
	OktaSignOnPolicy              *OktaSignOnPolicy
	PasswordPolicy                *PasswordPolicy
	PostAuthSessionPolicy         *PostAuthSessionPolicy
	ProfileEnrollmentPolicy       *ProfileEnrollmentPolicy
}

// AccessPolicyAsListPolicies200Response is a convenience function that returns AccessPolicy wrapped in ListPolicies200Response
func AccessPolicyAsListPolicies200Response(v *AccessPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		AccessPolicy: v,
	}
}

// AuthenticatorEnrollmentPolicyAsListPolicies200Response is a convenience function that returns AuthenticatorEnrollmentPolicy wrapped in ListPolicies200Response
func AuthenticatorEnrollmentPolicyAsListPolicies200Response(v *AuthenticatorEnrollmentPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		AuthenticatorEnrollmentPolicy: v,
	}
}

// DeviceSignalCollectionPolicyAsListPolicies200Response is a convenience function that returns DeviceSignalCollectionPolicy wrapped in ListPolicies200Response
func DeviceSignalCollectionPolicyAsListPolicies200Response(v *DeviceSignalCollectionPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		DeviceSignalCollectionPolicy: v,
	}
}

// EntityRiskPolicyAsListPolicies200Response is a convenience function that returns EntityRiskPolicy wrapped in ListPolicies200Response
func EntityRiskPolicyAsListPolicies200Response(v *EntityRiskPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		EntityRiskPolicy: v,
	}
}

// IdpDiscoveryPolicyAsListPolicies200Response is a convenience function that returns IdpDiscoveryPolicy wrapped in ListPolicies200Response
func IdpDiscoveryPolicyAsListPolicies200Response(v *IdpDiscoveryPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		IdpDiscoveryPolicy: v,
	}
}

// OktaSignOnPolicyAsListPolicies200Response is a convenience function that returns OktaSignOnPolicy wrapped in ListPolicies200Response
func OktaSignOnPolicyAsListPolicies200Response(v *OktaSignOnPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		OktaSignOnPolicy: v,
	}
}

// PasswordPolicyAsListPolicies200Response is a convenience function that returns PasswordPolicy wrapped in ListPolicies200Response
func PasswordPolicyAsListPolicies200Response(v *PasswordPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		PasswordPolicy: v,
	}
}

// PostAuthSessionPolicyAsListPolicies200Response is a convenience function that returns PostAuthSessionPolicy wrapped in ListPolicies200Response
func PostAuthSessionPolicyAsListPolicies200Response(v *PostAuthSessionPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		PostAuthSessionPolicy: v,
	}
}

// ProfileEnrollmentPolicyAsListPolicies200Response is a convenience function that returns ProfileEnrollmentPolicy wrapped in ListPolicies200Response
func ProfileEnrollmentPolicyAsListPolicies200Response(v *ProfileEnrollmentPolicy) ListPolicies200Response {
	return ListPolicies200Response{
		ProfileEnrollmentPolicy: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListPolicies200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ACCESS_POLICY'
	if jsonDict["type"] == "ACCESS_POLICY" {
		// try to unmarshal JSON data into AccessPolicy
		err = json.Unmarshal(data, &dst.AccessPolicy)
		if err == nil {
			return nil // data stored in dst.AccessPolicy, return on the first match
		} else {
			dst.AccessPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as AccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DEVICE_SIGNAL_COLLECTION'
	if jsonDict["type"] == "DEVICE_SIGNAL_COLLECTION" {
		// try to unmarshal JSON data into DeviceSignalCollectionPolicy
		err = json.Unmarshal(data, &dst.DeviceSignalCollectionPolicy)
		if err == nil {
			return nil // data stored in dst.DeviceSignalCollectionPolicy, return on the first match
		} else {
			dst.DeviceSignalCollectionPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as DeviceSignalCollectionPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as EntityRiskPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as IdpDiscoveryPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MFA_ENROLL'
	if jsonDict["type"] == "MFA_ENROLL" {
		// try to unmarshal JSON data into AuthenticatorEnrollmentPolicy
		err = json.Unmarshal(data, &dst.AuthenticatorEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.AuthenticatorEnrollmentPolicy, return on the first match
		} else {
			dst.AuthenticatorEnrollmentPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as AuthenticatorEnrollmentPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as OktaSignOnPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as PasswordPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'POST_AUTH_SESSION'
	if jsonDict["type"] == "POST_AUTH_SESSION" {
		// try to unmarshal JSON data into PostAuthSessionPolicy
		err = json.Unmarshal(data, &dst.PostAuthSessionPolicy)
		if err == nil {
			return nil // data stored in dst.PostAuthSessionPolicy, return on the first match
		} else {
			dst.PostAuthSessionPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as PostAuthSessionPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as ProfileEnrollmentPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as AccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorEnrollmentPolicy'
	if jsonDict["type"] == "AuthenticatorEnrollmentPolicy" {
		// try to unmarshal JSON data into AuthenticatorEnrollmentPolicy
		err = json.Unmarshal(data, &dst.AuthenticatorEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.AuthenticatorEnrollmentPolicy, return on the first match
		} else {
			dst.AuthenticatorEnrollmentPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as AuthenticatorEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DeviceSignalCollectionPolicy'
	if jsonDict["type"] == "DeviceSignalCollectionPolicy" {
		// try to unmarshal JSON data into DeviceSignalCollectionPolicy
		err = json.Unmarshal(data, &dst.DeviceSignalCollectionPolicy)
		if err == nil {
			return nil // data stored in dst.DeviceSignalCollectionPolicy, return on the first match
		} else {
			dst.DeviceSignalCollectionPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as DeviceSignalCollectionPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as EntityRiskPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as IdpDiscoveryPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as OktaSignOnPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as PasswordPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PostAuthSessionPolicy'
	if jsonDict["type"] == "PostAuthSessionPolicy" {
		// try to unmarshal JSON data into PostAuthSessionPolicy
		err = json.Unmarshal(data, &dst.PostAuthSessionPolicy)
		if err == nil {
			return nil // data stored in dst.PostAuthSessionPolicy, return on the first match
		} else {
			dst.PostAuthSessionPolicy = nil
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as PostAuthSessionPolicy: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal ListPolicies200Response as ProfileEnrollmentPolicy: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListPolicies200Response) MarshalJSON() ([]byte, error) {
	if src.AccessPolicy != nil {
		return json.Marshal(&src.AccessPolicy)
	}

	if src.AuthenticatorEnrollmentPolicy != nil {
		return json.Marshal(&src.AuthenticatorEnrollmentPolicy)
	}

	if src.DeviceSignalCollectionPolicy != nil {
		return json.Marshal(&src.DeviceSignalCollectionPolicy)
	}

	if src.EntityRiskPolicy != nil {
		return json.Marshal(&src.EntityRiskPolicy)
	}

	if src.IdpDiscoveryPolicy != nil {
		return json.Marshal(&src.IdpDiscoveryPolicy)
	}

	if src.OktaSignOnPolicy != nil {
		return json.Marshal(&src.OktaSignOnPolicy)
	}

	if src.PasswordPolicy != nil {
		return json.Marshal(&src.PasswordPolicy)
	}

	if src.PostAuthSessionPolicy != nil {
		return json.Marshal(&src.PostAuthSessionPolicy)
	}

	if src.ProfileEnrollmentPolicy != nil {
		return json.Marshal(&src.ProfileEnrollmentPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListPolicies200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessPolicy != nil {
		return obj.AccessPolicy
	}

	if obj.AuthenticatorEnrollmentPolicy != nil {
		return obj.AuthenticatorEnrollmentPolicy
	}

	if obj.DeviceSignalCollectionPolicy != nil {
		return obj.DeviceSignalCollectionPolicy
	}

	if obj.EntityRiskPolicy != nil {
		return obj.EntityRiskPolicy
	}

	if obj.IdpDiscoveryPolicy != nil {
		return obj.IdpDiscoveryPolicy
	}

	if obj.OktaSignOnPolicy != nil {
		return obj.OktaSignOnPolicy
	}

	if obj.PasswordPolicy != nil {
		return obj.PasswordPolicy
	}

	if obj.PostAuthSessionPolicy != nil {
		return obj.PostAuthSessionPolicy
	}

	if obj.ProfileEnrollmentPolicy != nil {
		return obj.ProfileEnrollmentPolicy
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListPolicies200Response) GetActualInstanceValue() interface{} {
	if obj.AccessPolicy != nil {
		return *obj.AccessPolicy
	}

	if obj.AuthenticatorEnrollmentPolicy != nil {
		return *obj.AuthenticatorEnrollmentPolicy
	}

	if obj.DeviceSignalCollectionPolicy != nil {
		return *obj.DeviceSignalCollectionPolicy
	}

	if obj.EntityRiskPolicy != nil {
		return *obj.EntityRiskPolicy
	}

	if obj.IdpDiscoveryPolicy != nil {
		return *obj.IdpDiscoveryPolicy
	}

	if obj.OktaSignOnPolicy != nil {
		return *obj.OktaSignOnPolicy
	}

	if obj.PasswordPolicy != nil {
		return *obj.PasswordPolicy
	}

	if obj.PostAuthSessionPolicy != nil {
		return *obj.PostAuthSessionPolicy
	}

	if obj.ProfileEnrollmentPolicy != nil {
		return *obj.ProfileEnrollmentPolicy
	}

	// all schemas are nil
	return nil
}

type NullableListPolicies200Response struct {
	value *ListPolicies200Response
	isSet bool
}

func (v NullableListPolicies200Response) Get() *ListPolicies200Response {
	return v.value
}

func (v *NullableListPolicies200Response) Set(val *ListPolicies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListPolicies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListPolicies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPolicies200Response(val *ListPolicies200Response) *NullableListPolicies200Response {
	return &NullableListPolicies200Response{value: val, isSet: true}
}

func (v NullableListPolicies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPolicies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
