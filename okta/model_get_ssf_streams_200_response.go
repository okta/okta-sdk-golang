/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)


//model_oneof.mustache
// GetSsfStreams200Response - struct for GetSsfStreams200Response
type GetSsfStreams200Response struct {
	StreamConfiguration *StreamConfiguration
	ArrayOfStreamConfiguration *[]StreamConfiguration
}

// StreamConfigurationAsGetSsfStreams200Response is a convenience function that returns StreamConfiguration wrapped in GetSsfStreams200Response
func StreamConfigurationAsGetSsfStreams200Response(v *StreamConfiguration) GetSsfStreams200Response {
	return GetSsfStreams200Response{
		StreamConfiguration: v,
	}
}

// []StreamConfigurationAsGetSsfStreams200Response is a convenience function that returns []StreamConfiguration wrapped in GetSsfStreams200Response
func ArrayOfStreamConfigurationAsGetSsfStreams200Response(v *[]StreamConfiguration) GetSsfStreams200Response {
	return GetSsfStreams200Response{
		ArrayOfStreamConfiguration: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *GetSsfStreams200Response) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into StreamConfiguration
        err = json.Unmarshal(data, &dst.StreamConfiguration)
        if err == nil {
                jsonStreamConfiguration, _ := json.Marshal(dst.StreamConfiguration)
                if string(jsonStreamConfiguration) == "{}" { // empty struct
                        dst.StreamConfiguration = nil
                } else {
                        match++
                }
        } else {
                dst.StreamConfiguration = nil
        }

        // try to unmarshal data into ArrayOfStreamConfiguration
        err = json.Unmarshal(data, &dst.ArrayOfStreamConfiguration)
        if err == nil {
                jsonArrayOfStreamConfiguration, _ := json.Marshal(dst.ArrayOfStreamConfiguration)
                if string(jsonArrayOfStreamConfiguration) == "{}" { // empty struct
                        dst.ArrayOfStreamConfiguration = nil
                } else {
                        match++
                }
        } else {
                dst.ArrayOfStreamConfiguration = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.StreamConfiguration = nil
                dst.ArrayOfStreamConfiguration = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(GetSsfStreams200Response)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(GetSsfStreams200Response)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSsfStreams200Response) MarshalJSON() ([]byte, error) {
	if src.StreamConfiguration != nil {
		return json.Marshal(&src.StreamConfiguration)
	}

	if src.ArrayOfStreamConfiguration != nil {
		return json.Marshal(&src.ArrayOfStreamConfiguration)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSsfStreams200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.StreamConfiguration != nil {
		return obj.StreamConfiguration
	}

	if obj.ArrayOfStreamConfiguration != nil {
		return obj.ArrayOfStreamConfiguration
	}

	// all schemas are nil
	return nil
}

type NullableGetSsfStreams200Response struct {
	value *GetSsfStreams200Response
	isSet bool
}

func (v NullableGetSsfStreams200Response) Get() *GetSsfStreams200Response {
	return v.value
}

func (v *NullableGetSsfStreams200Response) Set(val *GetSsfStreams200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSsfStreams200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSsfStreams200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSsfStreams200Response(val *GetSsfStreams200Response) *NullableGetSsfStreams200Response {
	return &NullableGetSsfStreams200Response{value: val, isSet: true}
}

func (v NullableGetSsfStreams200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSsfStreams200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


