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
// UserFactorTokenAllOfVerify - struct for UserFactorTokenAllOfVerify
type UserFactorTokenAllOfVerify struct {
	UserFactorTokenVerifyRSA *UserFactorTokenVerifyRSA
	UserFactorTokenVerifySymantec *UserFactorTokenVerifySymantec
}

// UserFactorTokenVerifyRSAAsUserFactorTokenAllOfVerify is a convenience function that returns UserFactorTokenVerifyRSA wrapped in UserFactorTokenAllOfVerify
func UserFactorTokenVerifyRSAAsUserFactorTokenAllOfVerify(v *UserFactorTokenVerifyRSA) UserFactorTokenAllOfVerify {
	return UserFactorTokenAllOfVerify{
		UserFactorTokenVerifyRSA: v,
	}
}

// UserFactorTokenVerifySymantecAsUserFactorTokenAllOfVerify is a convenience function that returns UserFactorTokenVerifySymantec wrapped in UserFactorTokenAllOfVerify
func UserFactorTokenVerifySymantecAsUserFactorTokenAllOfVerify(v *UserFactorTokenVerifySymantec) UserFactorTokenAllOfVerify {
	return UserFactorTokenAllOfVerify{
		UserFactorTokenVerifySymantec: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *UserFactorTokenAllOfVerify) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into UserFactorTokenVerifyRSA
        err = json.Unmarshal(data, &dst.UserFactorTokenVerifyRSA)
        if err == nil {
                jsonUserFactorTokenVerifyRSA, _ := json.Marshal(dst.UserFactorTokenVerifyRSA)
                if string(jsonUserFactorTokenVerifyRSA) == "{}" { // empty struct
                        dst.UserFactorTokenVerifyRSA = nil
                } else {
                        match++
                }
        } else {
                dst.UserFactorTokenVerifyRSA = nil
        }

        // try to unmarshal data into UserFactorTokenVerifySymantec
        err = json.Unmarshal(data, &dst.UserFactorTokenVerifySymantec)
        if err == nil {
                jsonUserFactorTokenVerifySymantec, _ := json.Marshal(dst.UserFactorTokenVerifySymantec)
                if string(jsonUserFactorTokenVerifySymantec) == "{}" { // empty struct
                        dst.UserFactorTokenVerifySymantec = nil
                } else {
                        match++
                }
        } else {
                dst.UserFactorTokenVerifySymantec = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.UserFactorTokenVerifyRSA = nil
                dst.UserFactorTokenVerifySymantec = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(UserFactorTokenAllOfVerify)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(UserFactorTokenAllOfVerify)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserFactorTokenAllOfVerify) MarshalJSON() ([]byte, error) {
	if src.UserFactorTokenVerifyRSA != nil {
		return json.Marshal(&src.UserFactorTokenVerifyRSA)
	}

	if src.UserFactorTokenVerifySymantec != nil {
		return json.Marshal(&src.UserFactorTokenVerifySymantec)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserFactorTokenAllOfVerify) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UserFactorTokenVerifyRSA != nil {
		return obj.UserFactorTokenVerifyRSA
	}

	if obj.UserFactorTokenVerifySymantec != nil {
		return obj.UserFactorTokenVerifySymantec
	}

	// all schemas are nil
	return nil
}

type NullableUserFactorTokenAllOfVerify struct {
	value *UserFactorTokenAllOfVerify
	isSet bool
}

func (v NullableUserFactorTokenAllOfVerify) Get() *UserFactorTokenAllOfVerify {
	return v.value
}

func (v *NullableUserFactorTokenAllOfVerify) Set(val *UserFactorTokenAllOfVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenAllOfVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenAllOfVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenAllOfVerify(val *UserFactorTokenAllOfVerify) *NullableUserFactorTokenAllOfVerify {
	return &NullableUserFactorTokenAllOfVerify{value: val, isSet: true}
}

func (v NullableUserFactorTokenAllOfVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenAllOfVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


