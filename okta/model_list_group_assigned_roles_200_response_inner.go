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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// ListGroupAssignedRoles200ResponseInner - struct for ListGroupAssignedRoles200ResponseInner
type ListGroupAssignedRoles200ResponseInner struct {
	CustomRole   *CustomRole
	StandardRole *StandardRole
}

// CustomRoleAsListGroupAssignedRoles200ResponseInner is a convenience function that returns CustomRole wrapped in ListGroupAssignedRoles200ResponseInner
func CustomRoleAsListGroupAssignedRoles200ResponseInner(v *CustomRole) ListGroupAssignedRoles200ResponseInner {
	return ListGroupAssignedRoles200ResponseInner{
		CustomRole: v,
	}
}

// StandardRoleAsListGroupAssignedRoles200ResponseInner is a convenience function that returns StandardRole wrapped in ListGroupAssignedRoles200ResponseInner
func StandardRoleAsListGroupAssignedRoles200ResponseInner(v *StandardRole) ListGroupAssignedRoles200ResponseInner {
	return ListGroupAssignedRoles200ResponseInner{
		StandardRole: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGroupAssignedRoles200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomRole
	err = json.Unmarshal(data, &dst.CustomRole)
	if err == nil {
		jsonCustomRole, _ := json.Marshal(dst.CustomRole)
		if string(jsonCustomRole) == "{}" { // empty struct
			dst.CustomRole = nil
		} else {
			match++
		}
	} else {
		dst.CustomRole = nil
	}

	// try to unmarshal data into StandardRole
	err = json.Unmarshal(data, &dst.StandardRole)
	if err == nil {
		jsonStandardRole, _ := json.Marshal(dst.StandardRole)
		if string(jsonStandardRole) == "{}" { // empty struct
			dst.StandardRole = nil
		} else {
			match++
		}
	} else {
		dst.StandardRole = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomRole = nil
		dst.StandardRole = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGroupAssignedRoles200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGroupAssignedRoles200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGroupAssignedRoles200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.CustomRole != nil {
		return json.Marshal(&src.CustomRole)
	}

	if src.StandardRole != nil {
		return json.Marshal(&src.StandardRole)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGroupAssignedRoles200ResponseInner) GetActualInstance() interface{} {
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
func (obj ListGroupAssignedRoles200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.CustomRole != nil {
		return *obj.CustomRole
	}

	if obj.StandardRole != nil {
		return *obj.StandardRole
	}

	// all schemas are nil
	return nil
}

type NullableListGroupAssignedRoles200ResponseInner struct {
	value *ListGroupAssignedRoles200ResponseInner
	isSet bool
}

func (v NullableListGroupAssignedRoles200ResponseInner) Get() *ListGroupAssignedRoles200ResponseInner {
	return v.value
}

func (v *NullableListGroupAssignedRoles200ResponseInner) Set(val *ListGroupAssignedRoles200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListGroupAssignedRoles200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListGroupAssignedRoles200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGroupAssignedRoles200ResponseInner(val *ListGroupAssignedRoles200ResponseInner) *NullableListGroupAssignedRoles200ResponseInner {
	return &NullableListGroupAssignedRoles200ResponseInner{value: val, isSet: true}
}

func (v NullableListGroupAssignedRoles200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGroupAssignedRoles200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
