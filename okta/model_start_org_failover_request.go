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

// StartOrgFailoverRequest - struct for StartOrgFailoverRequest
type StartOrgFailoverRequest struct {
	FailoverRequestSchema *FailoverRequestSchema
	MapmapOfStringAny     *map[string]interface{}
}

// FailoverRequestSchemaAsStartOrgFailoverRequest is a convenience function that returns FailoverRequestSchema wrapped in StartOrgFailoverRequest
func FailoverRequestSchemaAsStartOrgFailoverRequest(v *FailoverRequestSchema) StartOrgFailoverRequest {
	return StartOrgFailoverRequest{
		FailoverRequestSchema: v,
	}
}

// map[string]interface{}AsStartOrgFailoverRequest is a convenience function that returns map[string]interface{} wrapped in StartOrgFailoverRequest
func MapmapOfStringAnyAsStartOrgFailoverRequest(v *map[string]interface{}) StartOrgFailoverRequest {
	return StartOrgFailoverRequest{
		MapmapOfStringAny: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StartOrgFailoverRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FailoverRequestSchema
	err = json.Unmarshal(data, &dst.FailoverRequestSchema)
	if err == nil {
		jsonFailoverRequestSchema, _ := json.Marshal(dst.FailoverRequestSchema)
		if string(jsonFailoverRequestSchema) == "{}" { // empty struct
			dst.FailoverRequestSchema = nil
		} else {
			match++
		}
	} else {
		dst.FailoverRequestSchema = nil
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
		dst.FailoverRequestSchema = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StartOrgFailoverRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StartOrgFailoverRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StartOrgFailoverRequest) MarshalJSON() ([]byte, error) {
	if src.FailoverRequestSchema != nil {
		return json.Marshal(&src.FailoverRequestSchema)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StartOrgFailoverRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FailoverRequestSchema != nil {
		return obj.FailoverRequestSchema
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj StartOrgFailoverRequest) GetActualInstanceValue() interface{} {
	if obj.FailoverRequestSchema != nil {
		return *obj.FailoverRequestSchema
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableStartOrgFailoverRequest struct {
	value *StartOrgFailoverRequest
	isSet bool
}

func (v NullableStartOrgFailoverRequest) Get() *StartOrgFailoverRequest {
	return v.value
}

func (v *NullableStartOrgFailoverRequest) Set(val *StartOrgFailoverRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartOrgFailoverRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartOrgFailoverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartOrgFailoverRequest(val *StartOrgFailoverRequest) *NullableStartOrgFailoverRequest {
	return &NullableStartOrgFailoverRequest{value: val, isSet: true}
}

func (v NullableStartOrgFailoverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartOrgFailoverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
