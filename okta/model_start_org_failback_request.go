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

// StartOrgFailbackRequest - struct for StartOrgFailbackRequest
type StartOrgFailbackRequest struct {
	FailbackRequestSchema *FailbackRequestSchema
	MapmapOfStringAny     *map[string]interface{}
}

// FailbackRequestSchemaAsStartOrgFailbackRequest is a convenience function that returns FailbackRequestSchema wrapped in StartOrgFailbackRequest
func FailbackRequestSchemaAsStartOrgFailbackRequest(v *FailbackRequestSchema) StartOrgFailbackRequest {
	return StartOrgFailbackRequest{
		FailbackRequestSchema: v,
	}
}

// map[string]interface{}AsStartOrgFailbackRequest is a convenience function that returns map[string]interface{} wrapped in StartOrgFailbackRequest
func MapmapOfStringAnyAsStartOrgFailbackRequest(v *map[string]interface{}) StartOrgFailbackRequest {
	return StartOrgFailbackRequest{
		MapmapOfStringAny: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StartOrgFailbackRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FailbackRequestSchema
	err = json.Unmarshal(data, &dst.FailbackRequestSchema)
	if err == nil {
		jsonFailbackRequestSchema, _ := json.Marshal(dst.FailbackRequestSchema)
		if string(jsonFailbackRequestSchema) == "{}" { // empty struct
			dst.FailbackRequestSchema = nil
		} else {
			match++
		}
	} else {
		dst.FailbackRequestSchema = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FailbackRequestSchema = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StartOrgFailbackRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StartOrgFailbackRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StartOrgFailbackRequest) MarshalJSON() ([]byte, error) {
	if src.FailbackRequestSchema != nil {
		return json.Marshal(&src.FailbackRequestSchema)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StartOrgFailbackRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FailbackRequestSchema != nil {
		return obj.FailbackRequestSchema
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj StartOrgFailbackRequest) GetActualInstanceValue() interface{} {
	if obj.FailbackRequestSchema != nil {
		return *obj.FailbackRequestSchema
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableStartOrgFailbackRequest struct {
	value *StartOrgFailbackRequest
	isSet bool
}

func (v NullableStartOrgFailbackRequest) Get() *StartOrgFailbackRequest {
	return v.value
}

func (v *NullableStartOrgFailbackRequest) Set(val *StartOrgFailbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartOrgFailbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartOrgFailbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartOrgFailbackRequest(val *StartOrgFailbackRequest) *NullableStartOrgFailbackRequest {
	return &NullableStartOrgFailbackRequest{value: val, isSet: true}
}

func (v NullableStartOrgFailbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartOrgFailbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
