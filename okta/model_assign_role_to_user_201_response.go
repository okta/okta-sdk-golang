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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// AssignRoleToUser201Response - struct for AssignRoleToUser201Response
type AssignRoleToUser201Response struct {
	CustomRole   *CustomRole
	StandardRole *StandardRole
}

// CustomRoleAsAssignRoleToUser201Response is a convenience function that returns CustomRole wrapped in AssignRoleToUser201Response
func CustomRoleAsAssignRoleToUser201Response(v *CustomRole) AssignRoleToUser201Response {
	return AssignRoleToUser201Response{
		CustomRole: v,
	}
}

// StandardRoleAsAssignRoleToUser201Response is a convenience function that returns StandardRole wrapped in AssignRoleToUser201Response
func StandardRoleAsAssignRoleToUser201Response(v *StandardRole) AssignRoleToUser201Response {
	return AssignRoleToUser201Response{
		StandardRole: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssignRoleToUser201Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["type"].(string)

	// check if the discriminator value is 'ACCESS_CERTIFICATIONS_ADMIN'
	if discriminatorValue == "ACCESS_CERTIFICATIONS_ADMIN" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as CustomRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ACCESS_REQUESTS_ADMIN'
	if discriminatorValue == "ACCESS_REQUESTS_ADMIN" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as CustomRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'API_ACCESS_MANAGEMENT_ADMIN'
	if discriminatorValue == "API_ACCESS_MANAGEMENT_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'APP_ADMIN'
	if discriminatorValue == "APP_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CUSTOM'
	if discriminatorValue == "CUSTOM" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as CustomRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUP_MEMBERSHIP_ADMIN'
	if discriminatorValue == "GROUP_MEMBERSHIP_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HELP_DESK_ADMIN'
	if discriminatorValue == "HELP_DESK_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ORG_ADMIN'
	if discriminatorValue == "ORG_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'READ_ONLY_ADMIN'
	if discriminatorValue == "READ_ONLY_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'REPORT_ADMIN'
	if discriminatorValue == "REPORT_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SUPER_ADMIN'
	if discriminatorValue == "SUPER_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER_ADMIN'
	if discriminatorValue == "USER_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WORKFLOWS_ADMIN'
	if discriminatorValue == "WORKFLOWS_ADMIN" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToUser201Response as CustomRole: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil
		}
		dst.CustomRole = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssignRoleToUser201Response) MarshalJSON() ([]byte, error) {
	if src.CustomRole != nil {
		return json.Marshal(&src.CustomRole)
	}

	if src.StandardRole != nil {
		return json.Marshal(&src.StandardRole)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssignRoleToUser201Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomRole != nil {
		return obj.CustomRole
	}

	if obj.StandardRole != nil {
		return obj.StandardRole
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AssignRoleToUser201Response) GetActualInstanceValue() interface{} {
	if obj.CustomRole != nil {
		return *obj.CustomRole
	}

	if obj.StandardRole != nil {
		return *obj.StandardRole
	}

	// all schemas are nil
	return nil
}

type NullableAssignRoleToUser201Response struct {
	value *AssignRoleToUser201Response
	isSet bool
}

func (v NullableAssignRoleToUser201Response) Get() *AssignRoleToUser201Response {
	return v.value
}

func (v *NullableAssignRoleToUser201Response) Set(val *AssignRoleToUser201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignRoleToUser201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignRoleToUser201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignRoleToUser201Response(val *AssignRoleToUser201Response) *NullableAssignRoleToUser201Response {
	return &NullableAssignRoleToUser201Response{value: val, isSet: true}
}

func (v NullableAssignRoleToUser201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignRoleToUser201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
