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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)


//model_oneof.mustache
// SAMLHookResponseCommandsInnerValueInnerValue - The value of the claim that you add or replace, and can also include other attributes. If adding to a claim, add another `value` attribute residing within an array called `attributeValues`.  See the following examples:  #### Simple value (integer or string)  `\"value\": 300` or `\"value\": \"replacementString\"`  #### Attribute value (object)  ` \"value\": {     \"authContextClassRef\": \"replacementValue\"   }`  #### AttributeValues array value (object)  ` \"value\": {     \"attributes\": {       \"NameFormat\": \"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\"     },     \"attributeValues\": [       {\"attributes\": {         \"xsi:type\": \"xs:string\"       },       \"value\": \"4321\"}       ]     }`
type SAMLHookResponseCommandsInnerValueInnerValue struct {
	Int32 *int32
	MapmapOfStringAny *map[string]interface{}
	String *string
}

// int32AsSAMLHookResponseCommandsInnerValueInnerValue is a convenience function that returns int32 wrapped in SAMLHookResponseCommandsInnerValueInnerValue
func Int32AsSAMLHookResponseCommandsInnerValueInnerValue(v *int32) SAMLHookResponseCommandsInnerValueInnerValue {
	return SAMLHookResponseCommandsInnerValueInnerValue{
		Int32: v,
	}
}

// map[string]interface{}AsSAMLHookResponseCommandsInnerValueInnerValue is a convenience function that returns map[string]interface{} wrapped in SAMLHookResponseCommandsInnerValueInnerValue
func MapmapOfStringAnyAsSAMLHookResponseCommandsInnerValueInnerValue(v *map[string]interface{}) SAMLHookResponseCommandsInnerValueInnerValue {
	return SAMLHookResponseCommandsInnerValueInnerValue{
		MapmapOfStringAny: v,
	}
}

// stringAsSAMLHookResponseCommandsInnerValueInnerValue is a convenience function that returns string wrapped in SAMLHookResponseCommandsInnerValueInnerValue
func StringAsSAMLHookResponseCommandsInnerValueInnerValue(v *string) SAMLHookResponseCommandsInnerValueInnerValue {
	return SAMLHookResponseCommandsInnerValueInnerValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *SAMLHookResponseCommandsInnerValueInnerValue) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into Int32
        err = json.Unmarshal(data, &dst.Int32)
        if err == nil {
                jsonInt32, _ := json.Marshal(dst.Int32)
                if string(jsonInt32) == "{}" { // empty struct
                        dst.Int32 = nil
                } else {
                        match++
                }
        } else {
                dst.Int32 = nil
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
                dst.Int32 = nil
                dst.MapmapOfStringAny = nil
                dst.String = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(SAMLHookResponseCommandsInnerValueInnerValue)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(SAMLHookResponseCommandsInnerValueInnerValue)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SAMLHookResponseCommandsInnerValueInnerValue) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SAMLHookResponseCommandsInnerValueInnerValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSAMLHookResponseCommandsInnerValueInnerValue struct {
	value *SAMLHookResponseCommandsInnerValueInnerValue
	isSet bool
}

func (v NullableSAMLHookResponseCommandsInnerValueInnerValue) Get() *SAMLHookResponseCommandsInnerValueInnerValue {
	return v.value
}

func (v *NullableSAMLHookResponseCommandsInnerValueInnerValue) Set(val *SAMLHookResponseCommandsInnerValueInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLHookResponseCommandsInnerValueInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLHookResponseCommandsInnerValueInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLHookResponseCommandsInnerValueInnerValue(val *SAMLHookResponseCommandsInnerValueInnerValue) *NullableSAMLHookResponseCommandsInnerValueInnerValue {
	return &NullableSAMLHookResponseCommandsInnerValueInnerValue{value: val, isSet: true}
}

func (v NullableSAMLHookResponseCommandsInnerValueInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLHookResponseCommandsInnerValueInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


