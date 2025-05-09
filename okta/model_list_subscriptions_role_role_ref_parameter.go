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
// ListSubscriptionsRoleRoleRefParameter - struct for ListSubscriptionsRoleRoleRefParameter
type ListSubscriptionsRoleRoleRefParameter struct {
	String *string
}

// stringAsListSubscriptionsRoleRoleRefParameter is a convenience function that returns string wrapped in ListSubscriptionsRoleRoleRefParameter
func StringAsListSubscriptionsRoleRoleRefParameter(v *string) ListSubscriptionsRoleRoleRefParameter {
	return ListSubscriptionsRoleRoleRefParameter{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListSubscriptionsRoleRoleRefParameter) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
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
                dst.String = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(ListSubscriptionsRoleRoleRefParameter)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(ListSubscriptionsRoleRoleRefParameter)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListSubscriptionsRoleRoleRefParameter) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListSubscriptionsRoleRoleRefParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableListSubscriptionsRoleRoleRefParameter struct {
	value *ListSubscriptionsRoleRoleRefParameter
	isSet bool
}

func (v NullableListSubscriptionsRoleRoleRefParameter) Get() *ListSubscriptionsRoleRoleRefParameter {
	return v.value
}

func (v *NullableListSubscriptionsRoleRoleRefParameter) Set(val *ListSubscriptionsRoleRoleRefParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubscriptionsRoleRoleRefParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubscriptionsRoleRoleRefParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubscriptionsRoleRoleRefParameter(val *ListSubscriptionsRoleRoleRefParameter) *NullableListSubscriptionsRoleRoleRefParameter {
	return &NullableListSubscriptionsRoleRoleRefParameter{value: val, isSet: true}
}

func (v NullableListSubscriptionsRoleRoleRefParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubscriptionsRoleRoleRefParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


