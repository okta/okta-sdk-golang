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

// EntitlementManagementStatusValidationResponse - struct for EntitlementManagementStatusValidationResponse
type EntitlementManagementStatusValidationResponse struct {
	EntitlementManagementStatusValidationResponseOneOf  *EntitlementManagementStatusValidationResponseOneOf
	EntitlementManagementStatusValidationResponseOneOf1 *EntitlementManagementStatusValidationResponseOneOf1
}

// EntitlementManagementStatusValidationResponseOneOfAsEntitlementManagementStatusValidationResponse is a convenience function that returns EntitlementManagementStatusValidationResponseOneOf wrapped in EntitlementManagementStatusValidationResponse
func EntitlementManagementStatusValidationResponseOneOfAsEntitlementManagementStatusValidationResponse(v *EntitlementManagementStatusValidationResponseOneOf) EntitlementManagementStatusValidationResponse {
	return EntitlementManagementStatusValidationResponse{
		EntitlementManagementStatusValidationResponseOneOf: v,
	}
}

// EntitlementManagementStatusValidationResponseOneOf1AsEntitlementManagementStatusValidationResponse is a convenience function that returns EntitlementManagementStatusValidationResponseOneOf1 wrapped in EntitlementManagementStatusValidationResponse
func EntitlementManagementStatusValidationResponseOneOf1AsEntitlementManagementStatusValidationResponse(v *EntitlementManagementStatusValidationResponseOneOf1) EntitlementManagementStatusValidationResponse {
	return EntitlementManagementStatusValidationResponse{
		EntitlementManagementStatusValidationResponseOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntitlementManagementStatusValidationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntitlementManagementStatusValidationResponseOneOf
	err = json.Unmarshal(data, &dst.EntitlementManagementStatusValidationResponseOneOf)
	if err == nil {
		jsonEntitlementManagementStatusValidationResponseOneOf, _ := json.Marshal(dst.EntitlementManagementStatusValidationResponseOneOf)
		if string(jsonEntitlementManagementStatusValidationResponseOneOf) == "{}" { // empty struct
			dst.EntitlementManagementStatusValidationResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.EntitlementManagementStatusValidationResponseOneOf = nil
	}

	// try to unmarshal data into EntitlementManagementStatusValidationResponseOneOf1
	err = json.Unmarshal(data, &dst.EntitlementManagementStatusValidationResponseOneOf1)
	if err == nil {
		jsonEntitlementManagementStatusValidationResponseOneOf1, _ := json.Marshal(dst.EntitlementManagementStatusValidationResponseOneOf1)
		if string(jsonEntitlementManagementStatusValidationResponseOneOf1) == "{}" { // empty struct
			dst.EntitlementManagementStatusValidationResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.EntitlementManagementStatusValidationResponseOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EntitlementManagementStatusValidationResponseOneOf = nil
		dst.EntitlementManagementStatusValidationResponseOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EntitlementManagementStatusValidationResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EntitlementManagementStatusValidationResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntitlementManagementStatusValidationResponse) MarshalJSON() ([]byte, error) {
	if src.EntitlementManagementStatusValidationResponseOneOf != nil {
		return json.Marshal(&src.EntitlementManagementStatusValidationResponseOneOf)
	}

	if src.EntitlementManagementStatusValidationResponseOneOf1 != nil {
		return json.Marshal(&src.EntitlementManagementStatusValidationResponseOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntitlementManagementStatusValidationResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementManagementStatusValidationResponseOneOf != nil {
		return obj.EntitlementManagementStatusValidationResponseOneOf
	}

	if obj.EntitlementManagementStatusValidationResponseOneOf1 != nil {
		return obj.EntitlementManagementStatusValidationResponseOneOf1
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj EntitlementManagementStatusValidationResponse) GetActualInstanceValue() interface{} {
	if obj.EntitlementManagementStatusValidationResponseOneOf != nil {
		return *obj.EntitlementManagementStatusValidationResponseOneOf
	}

	if obj.EntitlementManagementStatusValidationResponseOneOf1 != nil {
		return *obj.EntitlementManagementStatusValidationResponseOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableEntitlementManagementStatusValidationResponse struct {
	value *EntitlementManagementStatusValidationResponse
	isSet bool
}

func (v NullableEntitlementManagementStatusValidationResponse) Get() *EntitlementManagementStatusValidationResponse {
	return v.value
}

func (v *NullableEntitlementManagementStatusValidationResponse) Set(val *EntitlementManagementStatusValidationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementManagementStatusValidationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementManagementStatusValidationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementManagementStatusValidationResponse(val *EntitlementManagementStatusValidationResponse) *NullableEntitlementManagementStatusValidationResponse {
	return &NullableEntitlementManagementStatusValidationResponse{value: val, isSet: true}
}

func (v NullableEntitlementManagementStatusValidationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementManagementStatusValidationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
