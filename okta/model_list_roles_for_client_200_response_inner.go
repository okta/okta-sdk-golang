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

// ListRolesForClient200ResponseInner - struct for ListRolesForClient200ResponseInner
type ListRolesForClient200ResponseInner struct {
	CustomRole   *CustomRole
	StandardRole *StandardRole
}

// CustomRoleAsListRolesForClient200ResponseInner is a convenience function that returns CustomRole wrapped in ListRolesForClient200ResponseInner
func CustomRoleAsListRolesForClient200ResponseInner(v *CustomRole) ListRolesForClient200ResponseInner {
	return ListRolesForClient200ResponseInner{
		CustomRole: v,
	}
}

// StandardRoleAsListRolesForClient200ResponseInner is a convenience function that returns StandardRole wrapped in ListRolesForClient200ResponseInner
func StandardRoleAsListRolesForClient200ResponseInner(v *StandardRole) ListRolesForClient200ResponseInner {
	return ListRolesForClient200ResponseInner{
		StandardRole: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRolesForClient200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ACCESS_CERTIFICATIONS_ADMIN'
	if jsonDict["type"] == "ACCESS_CERTIFICATIONS_ADMIN" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as CustomRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ACCESS_REQUESTS_ADMIN'
	if jsonDict["type"] == "ACCESS_REQUESTS_ADMIN" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as CustomRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'API_ACCESS_MANAGEMENT_ADMIN'
	if jsonDict["type"] == "API_ACCESS_MANAGEMENT_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'APP_ADMIN'
	if jsonDict["type"] == "APP_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as CustomRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUP_MEMBERSHIP_ADMIN'
	if jsonDict["type"] == "GROUP_MEMBERSHIP_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HELP_DESK_ADMIN'
	if jsonDict["type"] == "HELP_DESK_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ORG_ADMIN'
	if jsonDict["type"] == "ORG_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'READ_ONLY_ADMIN'
	if jsonDict["type"] == "READ_ONLY_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'REPORT_ADMIN'
	if jsonDict["type"] == "REPORT_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SUPER_ADMIN'
	if jsonDict["type"] == "SUPER_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER_ADMIN'
	if jsonDict["type"] == "USER_ADMIN" {
		// try to unmarshal JSON data into StandardRole
		err = json.Unmarshal(data, &dst.StandardRole)
		if err == nil {
			return nil // data stored in dst.StandardRole, return on the first match
		} else {
			dst.StandardRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as StandardRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WORKFLOWS_ADMIN'
	if jsonDict["type"] == "WORKFLOWS_ADMIN" {
		// try to unmarshal JSON data into CustomRole
		err = json.Unmarshal(data, &dst.CustomRole)
		if err == nil {
			return nil // data stored in dst.CustomRole, return on the first match
		} else {
			dst.CustomRole = nil
			return fmt.Errorf("failed to unmarshal ListRolesForClient200ResponseInner as CustomRole: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRolesForClient200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.CustomRole != nil {
		return json.Marshal(&src.CustomRole)
	}

	if src.StandardRole != nil {
		return json.Marshal(&src.StandardRole)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRolesForClient200ResponseInner) GetActualInstance() interface{} {
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
func (obj ListRolesForClient200ResponseInner) GetActualInstanceValue() interface{} {
	if obj.CustomRole != nil {
		return *obj.CustomRole
	}

	if obj.StandardRole != nil {
		return *obj.StandardRole
	}

	// all schemas are nil
	return nil
}

type NullableListRolesForClient200ResponseInner struct {
	value *ListRolesForClient200ResponseInner
	isSet bool
}

func (v NullableListRolesForClient200ResponseInner) Get() *ListRolesForClient200ResponseInner {
	return v.value
}

func (v *NullableListRolesForClient200ResponseInner) Set(val *ListRolesForClient200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesForClient200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesForClient200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesForClient200ResponseInner(val *ListRolesForClient200ResponseInner) *NullableListRolesForClient200ResponseInner {
	return &NullableListRolesForClient200ResponseInner{value: val, isSet: true}
}

func (v NullableListRolesForClient200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesForClient200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
