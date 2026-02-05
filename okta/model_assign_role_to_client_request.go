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

// AssignRoleToClientRequest - struct for AssignRoleToClientRequest
type AssignRoleToClientRequest struct {
	CustomRoleAssignmentSchema   *CustomRoleAssignmentSchema
	StandardRoleAssignmentSchema *StandardRoleAssignmentSchema
}

// CustomRoleAssignmentSchemaAsAssignRoleToClientRequest is a convenience function that returns CustomRoleAssignmentSchema wrapped in AssignRoleToClientRequest
func CustomRoleAssignmentSchemaAsAssignRoleToClientRequest(v *CustomRoleAssignmentSchema) AssignRoleToClientRequest {
	return AssignRoleToClientRequest{
		CustomRoleAssignmentSchema: v,
	}
}

// StandardRoleAssignmentSchemaAsAssignRoleToClientRequest is a convenience function that returns StandardRoleAssignmentSchema wrapped in AssignRoleToClientRequest
func StandardRoleAssignmentSchemaAsAssignRoleToClientRequest(v *StandardRoleAssignmentSchema) AssignRoleToClientRequest {
	return AssignRoleToClientRequest{
		StandardRoleAssignmentSchema: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AssignRoleToClientRequest) UnmarshalJSON(data []byte) error {
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
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ACCESS_REQUESTS_ADMIN'
	if discriminatorValue == "ACCESS_REQUESTS_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'API_ACCESS_MANAGEMENT_ADMIN'
	if discriminatorValue == "API_ACCESS_MANAGEMENT_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'APP_ADMIN'
	if discriminatorValue == "APP_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CUSTOM'
	if discriminatorValue == "CUSTOM" {
		// try to unmarshal JSON data into CustomRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.CustomRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.CustomRoleAssignmentSchema, return on the first match
		} else {
			dst.CustomRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as CustomRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUP_MEMBERSHIP_ADMIN'
	if discriminatorValue == "GROUP_MEMBERSHIP_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HELP_DESK_ADMIN'
	if discriminatorValue == "HELP_DESK_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ORG_ADMIN'
	if discriminatorValue == "ORG_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'READ_ONLY_ADMIN'
	if discriminatorValue == "READ_ONLY_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'REPORT_ADMIN'
	if discriminatorValue == "REPORT_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SUPER_ADMIN'
	if discriminatorValue == "SUPER_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER_ADMIN'
	if discriminatorValue == "USER_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WORKFLOWS_ADMIN'
	if discriminatorValue == "WORKFLOWS_ADMIN" {
		// try to unmarshal JSON data into StandardRoleAssignmentSchema
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil // data stored in dst.StandardRoleAssignmentSchema, return on the first match
		} else {
			dst.StandardRoleAssignmentSchema = nil
			return fmt.Errorf("failed to unmarshal AssignRoleToClientRequest as StandardRoleAssignmentSchema: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.StandardRoleAssignmentSchema)
		if err == nil {
			return nil
		}
		dst.StandardRoleAssignmentSchema = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AssignRoleToClientRequest) MarshalJSON() ([]byte, error) {
	if src.CustomRoleAssignmentSchema != nil {
		return json.Marshal(&src.CustomRoleAssignmentSchema)
	}

	if src.StandardRoleAssignmentSchema != nil {
		return json.Marshal(&src.StandardRoleAssignmentSchema)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AssignRoleToClientRequest) GetActualInstance() interface{} {
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
func (obj AssignRoleToClientRequest) GetActualInstanceValue() interface{} {
	if obj.CustomRoleAssignmentSchema != nil {
		return *obj.CustomRoleAssignmentSchema
	}

	if obj.StandardRoleAssignmentSchema != nil {
		return *obj.StandardRoleAssignmentSchema
	}

	// all schemas are nil
	return nil
}

type NullableAssignRoleToClientRequest struct {
	value *AssignRoleToClientRequest
	isSet bool
}

func (v NullableAssignRoleToClientRequest) Get() *AssignRoleToClientRequest {
	return v.value
}

func (v *NullableAssignRoleToClientRequest) Set(val *AssignRoleToClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignRoleToClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignRoleToClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignRoleToClientRequest(val *AssignRoleToClientRequest) *NullableAssignRoleToClientRequest {
	return &NullableAssignRoleToClientRequest{value: val, isSet: true}
}

func (v NullableAssignRoleToClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignRoleToClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
