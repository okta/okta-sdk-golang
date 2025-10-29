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

// GetUserRisk200Response - struct for GetUserRisk200Response
type GetUserRisk200Response struct {
	UserRiskLevelExists *UserRiskLevelExists
	UserRiskLevelNone   *UserRiskLevelNone
}

// UserRiskLevelExistsAsGetUserRisk200Response is a convenience function that returns UserRiskLevelExists wrapped in GetUserRisk200Response
func UserRiskLevelExistsAsGetUserRisk200Response(v *UserRiskLevelExists) GetUserRisk200Response {
	return GetUserRisk200Response{
		UserRiskLevelExists: v,
	}
}

// UserRiskLevelNoneAsGetUserRisk200Response is a convenience function that returns UserRiskLevelNone wrapped in GetUserRisk200Response
func UserRiskLevelNoneAsGetUserRisk200Response(v *UserRiskLevelNone) GetUserRisk200Response {
	return GetUserRisk200Response{
		UserRiskLevelNone: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetUserRisk200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'HIGH'
	if jsonDict["riskLevel"] == "HIGH" {
		// try to unmarshal JSON data into UserRiskLevelExists
		err = json.Unmarshal(data, &dst.UserRiskLevelExists)
		if err == nil {
			return nil // data stored in dst.UserRiskLevelExists, return on the first match
		} else {
			dst.UserRiskLevelExists = nil
			return fmt.Errorf("failed to unmarshal GetUserRisk200Response as UserRiskLevelExists: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LOW'
	if jsonDict["riskLevel"] == "LOW" {
		// try to unmarshal JSON data into UserRiskLevelExists
		err = json.Unmarshal(data, &dst.UserRiskLevelExists)
		if err == nil {
			return nil // data stored in dst.UserRiskLevelExists, return on the first match
		} else {
			dst.UserRiskLevelExists = nil
			return fmt.Errorf("failed to unmarshal GetUserRisk200Response as UserRiskLevelExists: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MEDIUM'
	if jsonDict["riskLevel"] == "MEDIUM" {
		// try to unmarshal JSON data into UserRiskLevelExists
		err = json.Unmarshal(data, &dst.UserRiskLevelExists)
		if err == nil {
			return nil // data stored in dst.UserRiskLevelExists, return on the first match
		} else {
			dst.UserRiskLevelExists = nil
			return fmt.Errorf("failed to unmarshal GetUserRisk200Response as UserRiskLevelExists: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NONE'
	if jsonDict["riskLevel"] == "NONE" {
		// try to unmarshal JSON data into UserRiskLevelNone
		err = json.Unmarshal(data, &dst.UserRiskLevelNone)
		if err == nil {
			return nil // data stored in dst.UserRiskLevelNone, return on the first match
		} else {
			dst.UserRiskLevelNone = nil
			return fmt.Errorf("failed to unmarshal GetUserRisk200Response as UserRiskLevelNone: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetUserRisk200Response) MarshalJSON() ([]byte, error) {
	if src.UserRiskLevelExists != nil {
		return json.Marshal(&src.UserRiskLevelExists)
	}

	if src.UserRiskLevelNone != nil {
		return json.Marshal(&src.UserRiskLevelNone)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetUserRisk200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserRiskLevelExists != nil {
		return obj.UserRiskLevelExists
	}

	if obj.UserRiskLevelNone != nil {
		return obj.UserRiskLevelNone
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetUserRisk200Response) GetActualInstanceValue() interface{} {
	if obj.UserRiskLevelExists != nil {
		return *obj.UserRiskLevelExists
	}

	if obj.UserRiskLevelNone != nil {
		return *obj.UserRiskLevelNone
	}

	// all schemas are nil
	return nil
}

type NullableGetUserRisk200Response struct {
	value *GetUserRisk200Response
	isSet bool
}

func (v NullableGetUserRisk200Response) Get() *GetUserRisk200Response {
	return v.value
}

func (v *NullableGetUserRisk200Response) Set(val *GetUserRisk200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserRisk200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserRisk200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserRisk200Response(val *GetUserRisk200Response) *NullableGetUserRisk200Response {
	return &NullableGetUserRisk200Response{value: val, isSet: true}
}

func (v NullableGetUserRisk200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserRisk200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
