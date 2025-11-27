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

// AssignRoleToGroupRequest - struct for AssignRoleToGroupRequest
type AssignRoleToGroupRequest struct {
	CustomRoleAssignmentSchema   *CustomRoleAssignmentSchema
	StandardRoleAssignmentSchema *StandardRoleAssignmentSchema
}

// CustomRoleAssignmentSchemaAsAssignRoleToGroupRequest is a convenience function that returns CustomRoleAssignmentSchema wrapped in AssignRoleToGroupRequest
func CustomRoleAssignmentSchemaAsAssignRoleToGroupRequest(v *CustomRoleAssignmentSchema) AssignRoleToGroupRequest {
	return AssignRoleToGroupRequest{
		CustomRoleAssignmentSchema: v,
	}
}

// StandardRoleAssignmentSchemaAsAssignRoleToGroupRequest is a convenience function that returns StandardRoleAssignmentSchema wrapped in AssignRoleToGroupRequest
func StandardRoleAssignmentSchemaAsAssignRoleToGroupRequest(v *StandardRoleAssignmentSchema) AssignRoleToGroupRequest {
	return AssignRoleToGroupRequest{
		StandardRoleAssignmentSchema: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssignRoleToGroupRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ACCESS_CERTIFICATIONS_ADMIN'
	if jsonDict["type"] == "ACCESS_CERTIFICATIONS_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ACCESS_REQUESTS_ADMIN'
	if jsonDict["type"] == "ACCESS_REQUESTS_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'API_ACCESS_MANAGEMENT_ADMIN'
	if jsonDict["type"] == "API_ACCESS_MANAGEMENT_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'APP_ADMIN'
	if jsonDict["type"] == "APP_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into CustomRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.CustomRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.CustomRoleAssignmentSchema, return on the first match
		} else {
			dst.CustomRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as CustomRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUP_MEMBERSHIP_ADMIN'
	if jsonDict["type"] == "GROUP_MEMBERSHIP_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HELP_DESK_ADMIN'
	if jsonDict["type"] == "HELP_DESK_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ORG_ADMIN'
	if jsonDict["type"] == "ORG_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'READ_ONLY_ADMIN'
	if jsonDict["type"] == "READ_ONLY_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'REPORT_ADMIN'
	if jsonDict["type"] == "REPORT_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SUPER_ADMIN'
	if jsonDict["type"] == "SUPER_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER_ADMIN'
	if jsonDict["type"] == "USER_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WORKFLOWS_ADMIN'
	if jsonDict["type"] == "WORKFLOWS_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToGroupRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssignRoleToGroupRequest) MarshalJSON() ([]byte, error) {
	if src.CustomRoleAssignmentSchema != nil {
		return json.Marshal(&src.CustomRoleAssignmentSchema)
	}

	if src.StandardRoleAssignmentSchema != nil {
		return json.Marshal(&src.StandardRoleAssignmentSchema)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssignRoleToGroupRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomRoleAssignmentSchema != nil {
		return obj.CustomRoleAssignmentSchema
	}

	if obj.StandardRoleAssignmentSchema != nil {
		return obj.StandardRoleAssignmentSchema
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AssignRoleToGroupRequest) GetActualInstanceValue() interface{} {
	if obj.CustomRoleAssignmentSchema != nil {
		return *obj.CustomRoleAssignmentSchema
	}

	if obj.StandardRoleAssignmentSchema != nil {
		return *obj.StandardRoleAssignmentSchema
	}

	// all schemas are nil
	return nil
}

type NullableAssignRoleToGroupRequest struct {
	value *AssignRoleToGroupRequest
	isSet bool
}

func (v NullableAssignRoleToGroupRequest) Get() *AssignRoleToGroupRequest {
	return v.value
}

func (v *NullableAssignRoleToGroupRequest) Set(val *AssignRoleToGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignRoleToGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignRoleToGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignRoleToGroupRequest(val *AssignRoleToGroupRequest) *NullableAssignRoleToGroupRequest {
	return &NullableAssignRoleToGroupRequest{value: val, isSet: true}
}

func (v NullableAssignRoleToGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignRoleToGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
