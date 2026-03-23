/*
Okta Admin Management API

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

// SubmissionCapabilityEnhanced - Enhanced capability structure with embedded protocol configurations for submission GET endpoint
type SubmissionCapabilityEnhanced struct {
	ApiServiceCapability      *ApiServiceCapability
	EntitlementsCapability    *EntitlementsCapability
	ProvisioningCapability    *ProvisioningCapability
	SsoCapability             *SsoCapability
	UniversalLogoutCapability *UniversalLogoutCapability
}

// ApiServiceCapabilityAsSubmissionCapabilityEnhanced is a convenience function that returns ApiServiceCapability wrapped in SubmissionCapabilityEnhanced
func ApiServiceCapabilityAsSubmissionCapabilityEnhanced(v *ApiServiceCapability) SubmissionCapabilityEnhanced {
	return SubmissionCapabilityEnhanced{
		ApiServiceCapability: v,
	}
}

// EntitlementsCapabilityAsSubmissionCapabilityEnhanced is a convenience function that returns EntitlementsCapability wrapped in SubmissionCapabilityEnhanced
func EntitlementsCapabilityAsSubmissionCapabilityEnhanced(v *EntitlementsCapability) SubmissionCapabilityEnhanced {
	return SubmissionCapabilityEnhanced{
		EntitlementsCapability: v,
	}
}

// ProvisioningCapabilityAsSubmissionCapabilityEnhanced is a convenience function that returns ProvisioningCapability wrapped in SubmissionCapabilityEnhanced
func ProvisioningCapabilityAsSubmissionCapabilityEnhanced(v *ProvisioningCapability) SubmissionCapabilityEnhanced {
	return SubmissionCapabilityEnhanced{
		ProvisioningCapability: v,
	}
}

// SsoCapabilityAsSubmissionCapabilityEnhanced is a convenience function that returns SsoCapability wrapped in SubmissionCapabilityEnhanced
func SsoCapabilityAsSubmissionCapabilityEnhanced(v *SsoCapability) SubmissionCapabilityEnhanced {
	return SubmissionCapabilityEnhanced{
		SsoCapability: v,
	}
}

// UniversalLogoutCapabilityAsSubmissionCapabilityEnhanced is a convenience function that returns UniversalLogoutCapability wrapped in SubmissionCapabilityEnhanced
func UniversalLogoutCapabilityAsSubmissionCapabilityEnhanced(v *UniversalLogoutCapability) SubmissionCapabilityEnhanced {
	return SubmissionCapabilityEnhanced{
		UniversalLogoutCapability: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmissionCapabilityEnhanced) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["capability"].(string)

	// check if the discriminator value is 'API_SERVICE'
	if discriminatorValue == "API_SERVICE" {
		// try to unmarshal JSON data into ApiServiceCapability
		err = json.Unmarshal(data, &dst.ApiServiceCapability)
		if err == nil {
			return nil // data stored in dst.ApiServiceCapability, return on the first match
		} else {
			dst.ApiServiceCapability = nil
			return fmt.Errorf("failed to unmarshal SubmissionCapabilityEnhanced as ApiServiceCapability: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITLEMENTS'
	if discriminatorValue == "ENTITLEMENTS" {
		// try to unmarshal JSON data into EntitlementsCapability
		err = json.Unmarshal(data, &dst.EntitlementsCapability)
		if err == nil {
			return nil // data stored in dst.EntitlementsCapability, return on the first match
		} else {
			dst.EntitlementsCapability = nil
			return fmt.Errorf("failed to unmarshal SubmissionCapabilityEnhanced as EntitlementsCapability: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PROVISIONING'
	if discriminatorValue == "PROVISIONING" {
		// try to unmarshal JSON data into ProvisioningCapability
		err = json.Unmarshal(data, &dst.ProvisioningCapability)
		if err == nil {
			return nil // data stored in dst.ProvisioningCapability, return on the first match
		} else {
			dst.ProvisioningCapability = nil
			return fmt.Errorf("failed to unmarshal SubmissionCapabilityEnhanced as ProvisioningCapability: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SSO'
	if discriminatorValue == "SSO" {
		// try to unmarshal JSON data into SsoCapability
		err = json.Unmarshal(data, &dst.SsoCapability)
		if err == nil {
			return nil // data stored in dst.SsoCapability, return on the first match
		} else {
			dst.SsoCapability = nil
			return fmt.Errorf("failed to unmarshal SubmissionCapabilityEnhanced as SsoCapability: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UNIVERSAL_LOGOUT'
	if discriminatorValue == "UNIVERSAL_LOGOUT" {
		// try to unmarshal JSON data into UniversalLogoutCapability
		err = json.Unmarshal(data, &dst.UniversalLogoutCapability)
		if err == nil {
			return nil // data stored in dst.UniversalLogoutCapability, return on the first match
		} else {
			dst.UniversalLogoutCapability = nil
			return fmt.Errorf("failed to unmarshal SubmissionCapabilityEnhanced as UniversalLogoutCapability: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.UniversalLogoutCapability)
		if err == nil {
			return nil
		}
		dst.UniversalLogoutCapability = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmissionCapabilityEnhanced) MarshalJSON() ([]byte, error) {
	if src.ApiServiceCapability != nil {
		return json.Marshal(&src.ApiServiceCapability)
	}

	if src.EntitlementsCapability != nil {
		return json.Marshal(&src.EntitlementsCapability)
	}

	if src.ProvisioningCapability != nil {
		return json.Marshal(&src.ProvisioningCapability)
	}

	if src.SsoCapability != nil {
		return json.Marshal(&src.SsoCapability)
	}

	if src.UniversalLogoutCapability != nil {
		return json.Marshal(&src.UniversalLogoutCapability)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmissionCapabilityEnhanced) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApiServiceCapability != nil {
		return obj.ApiServiceCapability
	}

	if obj.EntitlementsCapability != nil {
		return obj.EntitlementsCapability
	}

	if obj.ProvisioningCapability != nil {
		return obj.ProvisioningCapability
	}

	if obj.SsoCapability != nil {
		return obj.SsoCapability
	}

	if obj.UniversalLogoutCapability != nil {
		return obj.UniversalLogoutCapability
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SubmissionCapabilityEnhanced) GetActualInstanceValue() interface{} {
	if obj.ApiServiceCapability != nil {
		return *obj.ApiServiceCapability
	}

	if obj.EntitlementsCapability != nil {
		return *obj.EntitlementsCapability
	}

	if obj.ProvisioningCapability != nil {
		return *obj.ProvisioningCapability
	}

	if obj.SsoCapability != nil {
		return *obj.SsoCapability
	}

	if obj.UniversalLogoutCapability != nil {
		return *obj.UniversalLogoutCapability
	}

	// all schemas are nil
	return nil
}

type NullableSubmissionCapabilityEnhanced struct {
	value *SubmissionCapabilityEnhanced
	isSet bool
}

func (v NullableSubmissionCapabilityEnhanced) Get() *SubmissionCapabilityEnhanced {
	return v.value
}

func (v *NullableSubmissionCapabilityEnhanced) Set(val *SubmissionCapabilityEnhanced) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionCapabilityEnhanced) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionCapabilityEnhanced) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionCapabilityEnhanced(val *SubmissionCapabilityEnhanced) *NullableSubmissionCapabilityEnhanced {
	return &NullableSubmissionCapabilityEnhanced{value: val, isSet: true}
}

func (v NullableSubmissionCapabilityEnhanced) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionCapabilityEnhanced) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
