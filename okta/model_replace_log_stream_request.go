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

// ReplaceLogStreamRequest - struct for ReplaceLogStreamRequest
type ReplaceLogStreamRequest struct {
	LogStreamAwsPutSchema    *LogStreamAwsPutSchema
	LogStreamSplunkPutSchema *LogStreamSplunkPutSchema
}

// LogStreamAwsPutSchemaAsReplaceLogStreamRequest is a convenience function that returns LogStreamAwsPutSchema wrapped in ReplaceLogStreamRequest
func LogStreamAwsPutSchemaAsReplaceLogStreamRequest(v *LogStreamAwsPutSchema) ReplaceLogStreamRequest {
	return ReplaceLogStreamRequest{
		LogStreamAwsPutSchema: v,
	}
}

// LogStreamSplunkPutSchemaAsReplaceLogStreamRequest is a convenience function that returns LogStreamSplunkPutSchema wrapped in ReplaceLogStreamRequest
func LogStreamSplunkPutSchemaAsReplaceLogStreamRequest(v *LogStreamSplunkPutSchema) ReplaceLogStreamRequest {
	return ReplaceLogStreamRequest{
		LogStreamSplunkPutSchema: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReplaceLogStreamRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["type"].(string)

	// check if the discriminator value is 'aws_eventbridge'
	if discriminatorValue == "aws_eventbridge" {
		// try to unmarshal JSON data into LogStreamAwsPutSchema
		err = json.Unmarshal(data, &dst.LogStreamAwsPutSchema)
		if err == nil {
			return nil // data stored in dst.LogStreamAwsPutSchema, return on the first match
		} else {
			dst.LogStreamAwsPutSchema = nil
			return fmt.Errorf("failed to unmarshal ReplaceLogStreamRequest as LogStreamAwsPutSchema: %s", err.Error())
		}
	}

	// check if the discriminator value is 'splunk_cloud_logstreaming'
	if discriminatorValue == "splunk_cloud_logstreaming" {
		// try to unmarshal JSON data into LogStreamSplunkPutSchema
		err = json.Unmarshal(data, &dst.LogStreamSplunkPutSchema)
		if err == nil {
			return nil // data stored in dst.LogStreamSplunkPutSchema, return on the first match
		} else {
			dst.LogStreamSplunkPutSchema = nil
			return fmt.Errorf("failed to unmarshal ReplaceLogStreamRequest as LogStreamSplunkPutSchema: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.LogStreamSplunkPutSchema)
		if err == nil {
			return nil
		}
		dst.LogStreamSplunkPutSchema = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReplaceLogStreamRequest) MarshalJSON() ([]byte, error) {
	if src.LogStreamAwsPutSchema != nil {
		return json.Marshal(&src.LogStreamAwsPutSchema)
	}

	if src.LogStreamSplunkPutSchema != nil {
		return json.Marshal(&src.LogStreamSplunkPutSchema)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReplaceLogStreamRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.LogStreamAwsPutSchema != nil {
		return obj.LogStreamAwsPutSchema
	}

	if obj.LogStreamSplunkPutSchema != nil {
		return obj.LogStreamSplunkPutSchema
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ReplaceLogStreamRequest) GetActualInstanceValue() interface{} {
	if obj.LogStreamAwsPutSchema != nil {
		return *obj.LogStreamAwsPutSchema
	}

	if obj.LogStreamSplunkPutSchema != nil {
		return *obj.LogStreamSplunkPutSchema
	}

	// all schemas are nil
	return nil
}

type NullableReplaceLogStreamRequest struct {
	value *ReplaceLogStreamRequest
	isSet bool
}

func (v NullableReplaceLogStreamRequest) Get() *ReplaceLogStreamRequest {
	return v.value
}

func (v *NullableReplaceLogStreamRequest) Set(val *ReplaceLogStreamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceLogStreamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceLogStreamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceLogStreamRequest(val *ReplaceLogStreamRequest) *NullableReplaceLogStreamRequest {
	return &NullableReplaceLogStreamRequest{value: val, isSet: true}
}

func (v NullableReplaceLogStreamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceLogStreamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
