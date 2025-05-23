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
// StreamConfigurationAud - The audience used in the SET. This value is set as `aud` in the claim.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter.
type StreamConfigurationAud struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsStreamConfigurationAud is a convenience function that returns []string wrapped in StreamConfigurationAud
func ArrayOfStringAsStreamConfigurationAud(v *[]string) StreamConfigurationAud {
	return StreamConfigurationAud{
		ArrayOfString: v,
	}
}

// stringAsStreamConfigurationAud is a convenience function that returns string wrapped in StreamConfigurationAud
func StringAsStreamConfigurationAud(v *string) StreamConfigurationAud {
	return StreamConfigurationAud{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *StreamConfigurationAud) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into ArrayOfString
        err = json.Unmarshal(data, &dst.ArrayOfString)
        if err == nil {
                jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
                if string(jsonArrayOfString) == "{}" { // empty struct
                        dst.ArrayOfString = nil
                } else {
                        match++
                }
        } else {
                dst.ArrayOfString = nil
        }

        // try to unmarshal data into String
        err = json.Unmarshal(data, &dst.String)
        if err == nil {
                jsonString, _ := json.Marshal(dst.String)
                if string(jsonString) == "{}" { // empty struct
                        dst.String = nil
                } else {
                        match++
                }
        } else {
                dst.String = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.ArrayOfString = nil
                dst.String = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(StreamConfigurationAud)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(StreamConfigurationAud)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StreamConfigurationAud) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StreamConfigurationAud) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableStreamConfigurationAud struct {
	value *StreamConfigurationAud
	isSet bool
}

func (v NullableStreamConfigurationAud) Get() *StreamConfigurationAud {
	return v.value
}

func (v *NullableStreamConfigurationAud) Set(val *StreamConfigurationAud) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamConfigurationAud) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamConfigurationAud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamConfigurationAud(val *StreamConfigurationAud) *NullableStreamConfigurationAud {
	return &NullableStreamConfigurationAud{value: val, isSet: true}
}

func (v NullableStreamConfigurationAud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamConfigurationAud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


