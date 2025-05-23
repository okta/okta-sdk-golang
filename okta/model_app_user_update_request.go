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
// AppUserUpdateRequest - struct for AppUserUpdateRequest
type AppUserUpdateRequest struct {
	AppUserCredentialsRequestPayload *AppUserCredentialsRequestPayload
	AppUserProfileRequestPayload *AppUserProfileRequestPayload
}

// AppUserCredentialsRequestPayloadAsAppUserUpdateRequest is a convenience function that returns AppUserCredentialsRequestPayload wrapped in AppUserUpdateRequest
func AppUserCredentialsRequestPayloadAsAppUserUpdateRequest(v *AppUserCredentialsRequestPayload) AppUserUpdateRequest {
	return AppUserUpdateRequest{
		AppUserCredentialsRequestPayload: v,
	}
}

// AppUserProfileRequestPayloadAsAppUserUpdateRequest is a convenience function that returns AppUserProfileRequestPayload wrapped in AppUserUpdateRequest
func AppUserProfileRequestPayloadAsAppUserUpdateRequest(v *AppUserProfileRequestPayload) AppUserUpdateRequest {
	return AppUserUpdateRequest{
		AppUserProfileRequestPayload: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *AppUserUpdateRequest) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into AppUserCredentialsRequestPayload
        err = json.Unmarshal(data, &dst.AppUserCredentialsRequestPayload)
        if err == nil {
                jsonAppUserCredentialsRequestPayload, _ := json.Marshal(dst.AppUserCredentialsRequestPayload)
                if string(jsonAppUserCredentialsRequestPayload) == "{}" { // empty struct
                        dst.AppUserCredentialsRequestPayload = nil
                } else {
                        match++
                }
        } else {
                dst.AppUserCredentialsRequestPayload = nil
        }

        // try to unmarshal data into AppUserProfileRequestPayload
        err = json.Unmarshal(data, &dst.AppUserProfileRequestPayload)
        if err == nil {
                jsonAppUserProfileRequestPayload, _ := json.Marshal(dst.AppUserProfileRequestPayload)
                if string(jsonAppUserProfileRequestPayload) == "{}" { // empty struct
                        dst.AppUserProfileRequestPayload = nil
                } else {
                        match++
                }
        } else {
                dst.AppUserProfileRequestPayload = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.AppUserCredentialsRequestPayload = nil
                dst.AppUserProfileRequestPayload = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(AppUserUpdateRequest)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(AppUserUpdateRequest)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppUserUpdateRequest) MarshalJSON() ([]byte, error) {
	if src.AppUserCredentialsRequestPayload != nil {
		return json.Marshal(&src.AppUserCredentialsRequestPayload)
	}

	if src.AppUserProfileRequestPayload != nil {
		return json.Marshal(&src.AppUserProfileRequestPayload)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppUserUpdateRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppUserCredentialsRequestPayload != nil {
		return obj.AppUserCredentialsRequestPayload
	}

	if obj.AppUserProfileRequestPayload != nil {
		return obj.AppUserProfileRequestPayload
	}

	// all schemas are nil
	return nil
}

type NullableAppUserUpdateRequest struct {
	value *AppUserUpdateRequest
	isSet bool
}

func (v NullableAppUserUpdateRequest) Get() *AppUserUpdateRequest {
	return v.value
}

func (v *NullableAppUserUpdateRequest) Set(val *AppUserUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserUpdateRequest(val *AppUserUpdateRequest) *NullableAppUserUpdateRequest {
	return &NullableAppUserUpdateRequest{value: val, isSet: true}
}

func (v NullableAppUserUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


