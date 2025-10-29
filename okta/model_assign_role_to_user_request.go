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

// AssignRoleToUserRequest - struct for AssignRoleToUserRequest
type AssignRoleToUserRequest struct {
	CustomRoleAssignmentSchema   *CustomRoleAssignmentSchema
	StandardRoleAssignmentSchema *StandardRoleAssignmentSchema
}

// CustomRoleAssignmentSchemaAsAssignRoleToUserRequest is a convenience function that returns CustomRoleAssignmentSchema wrapped in AssignRoleToUserRequest
func CustomRoleAssignmentSchemaAsAssignRoleToUserRequest(v *CustomRoleAssignmentSchema) AssignRoleToUserRequest {
	return AssignRoleToUserRequest{
		CustomRoleAssignmentSchema: v,
	}
}

// StandardRoleAssignmentSchemaAsAssignRoleToUserRequest is a convenience function that returns StandardRoleAssignmentSchema wrapped in AssignRoleToUserRequest
func StandardRoleAssignmentSchemaAsAssignRoleToUserRequest(v *StandardRoleAssignmentSchema) AssignRoleToUserRequest {
	return AssignRoleToUserRequest{
		StandardRoleAssignmentSchema: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssignRoleToUserRequest) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(AssignRoleToUserRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AssignRoleToUserRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssignRoleToUserRequest) MarshalJSON() ([]byte, error) {
	if src.CustomRoleAssignmentSchema != nil {
		return json.Marshal(&src.CustomRoleAssignmentSchema)
	}

	if src.StandardRoleAssignmentSchema != nil {
		return json.Marshal(&src.StandardRoleAssignmentSchema)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssignRoleToUserRequest) GetActualInstance() interface{} {
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
func (obj AssignRoleToUserRequest) GetActualInstanceValue() interface{} {
	if obj.CustomRoleAssignmentSchema != nil {
		return *obj.CustomRoleAssignmentSchema
	}

	if obj.StandardRoleAssignmentSchema != nil {
		return *obj.StandardRoleAssignmentSchema
	}

	// all schemas are nil
	return nil
}

type NullableAssignRoleToUserRequest struct {
	value *AssignRoleToUserRequest
	isSet bool
}

func (v NullableAssignRoleToUserRequest) Get() *AssignRoleToUserRequest {
	return v.value
}

func (v *NullableAssignRoleToUserRequest) Set(val *AssignRoleToUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignRoleToUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignRoleToUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignRoleToUserRequest(val *AssignRoleToUserRequest) *NullableAssignRoleToUserRequest {
	return &NullableAssignRoleToUserRequest{value: val, isSet: true}
}

func (v NullableAssignRoleToUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignRoleToUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
