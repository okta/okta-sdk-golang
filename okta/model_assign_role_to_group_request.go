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
	match := 0
	// try to unmarshal data into CustomRoleAssignmentSchema
	err = json.Unmarshal(data, &dst.CustomRoleAssignmentSchema)
	if err == nil {
		jsonCustomRoleAssignmentSchema, _ := json.Marshal(dst.CustomRoleAssignmentSchema)
		if string(jsonCustomRoleAssignmentSchema) == "{}" { // empty struct
			dst.CustomRoleAssignmentSchema = nil
		} else {
			match++
		}
	} else {
		dst.CustomRoleAssignmentSchema = nil
	}

	// try to unmarshal data into StandardRoleAssignmentSchema
	err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
	if err == nil {
		jsonStandardRoleAssignmentSchema, _ := json.Marshal(dst.StandardRoleAssignmentSchema)
		if string(jsonStandardRoleAssignmentSchema) == "{}" { // empty struct
			dst.StandardRoleAssignmentSchema = nil
		} else {
			match++
		}
	} else {
		dst.StandardRoleAssignmentSchema = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomRoleAssignmentSchema = nil
		dst.StandardRoleAssignmentSchema = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AssignRoleToGroupRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AssignRoleToGroupRequest)")
	}
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
