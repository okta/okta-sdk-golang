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
// ResendEnrollFactorRequest - struct for ResendEnrollFactorRequest
type ResendEnrollFactorRequest struct {
	UserFactorCall *UserFactorCall
	UserFactorEmail *UserFactorEmail
	UserFactorSMS *UserFactorSMS
}

// UserFactorCallAsResendEnrollFactorRequest is a convenience function that returns UserFactorCall wrapped in ResendEnrollFactorRequest
func UserFactorCallAsResendEnrollFactorRequest(v *UserFactorCall) ResendEnrollFactorRequest {
	return ResendEnrollFactorRequest{
		UserFactorCall: v,
	}
}

// UserFactorEmailAsResendEnrollFactorRequest is a convenience function that returns UserFactorEmail wrapped in ResendEnrollFactorRequest
func UserFactorEmailAsResendEnrollFactorRequest(v *UserFactorEmail) ResendEnrollFactorRequest {
	return ResendEnrollFactorRequest{
		UserFactorEmail: v,
	}
}

// UserFactorSMSAsResendEnrollFactorRequest is a convenience function that returns UserFactorSMS wrapped in ResendEnrollFactorRequest
func UserFactorSMSAsResendEnrollFactorRequest(v *UserFactorSMS) ResendEnrollFactorRequest {
	return ResendEnrollFactorRequest{
		UserFactorSMS: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ResendEnrollFactorRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'UserFactorCall'
	if jsonDict["factorType"] == "UserFactorCall" {
		// try to unmarshal JSON data into UserFactorCall
		err = json.Unmarshal(data, &dst.UserFactorCall)
		if err == nil {
			return nil // data stored in dst.UserFactorCall, return on the first match
		} else {
			dst.UserFactorCall = nil
			return fmt.Errorf("Failed to unmarshal ResendEnrollFactorRequest as UserFactorCall: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorEmail'
	if jsonDict["factorType"] == "UserFactorEmail" {
		// try to unmarshal JSON data into UserFactorEmail
		err = json.Unmarshal(data, &dst.UserFactorEmail)
		if err == nil {
			return nil // data stored in dst.UserFactorEmail, return on the first match
		} else {
			dst.UserFactorEmail = nil
			return fmt.Errorf("Failed to unmarshal ResendEnrollFactorRequest as UserFactorEmail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UserFactorSMS'
	if jsonDict["factorType"] == "UserFactorSMS" {
		// try to unmarshal JSON data into UserFactorSMS
		err = json.Unmarshal(data, &dst.UserFactorSMS)
		if err == nil {
			return nil // data stored in dst.UserFactorSMS, return on the first match
		} else {
			dst.UserFactorSMS = nil
			return fmt.Errorf("Failed to unmarshal ResendEnrollFactorRequest as UserFactorSMS: %s", err.Error())
		}
	}

	// check if the discriminator value is 'call'
	if jsonDict["factorType"] == "call" {
		// try to unmarshal JSON data into UserFactorCall
		err = json.Unmarshal(data, &dst.UserFactorCall)
		if err == nil {
			return nil // data stored in dst.UserFactorCall, return on the first match
		} else {
			dst.UserFactorCall = nil
			return fmt.Errorf("Failed to unmarshal ResendEnrollFactorRequest as UserFactorCall: %s", err.Error())
		}
	}

	// check if the discriminator value is 'email'
	if jsonDict["factorType"] == "email" {
		// try to unmarshal JSON data into UserFactorEmail
		err = json.Unmarshal(data, &dst.UserFactorEmail)
		if err == nil {
			return nil // data stored in dst.UserFactorEmail, return on the first match
		} else {
			dst.UserFactorEmail = nil
			return fmt.Errorf("Failed to unmarshal ResendEnrollFactorRequest as UserFactorEmail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sms'
	if jsonDict["factorType"] == "sms" {
		// try to unmarshal JSON data into UserFactorSMS
		err = json.Unmarshal(data, &dst.UserFactorSMS)
		if err == nil {
			return nil // data stored in dst.UserFactorSMS, return on the first match
		} else {
			dst.UserFactorSMS = nil
			return fmt.Errorf("Failed to unmarshal ResendEnrollFactorRequest as UserFactorSMS: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResendEnrollFactorRequest) MarshalJSON() ([]byte, error) {
	if src.UserFactorCall != nil {
		return json.Marshal(&src.UserFactorCall)
	}

	if src.UserFactorEmail != nil {
		return json.Marshal(&src.UserFactorEmail)
	}

	if src.UserFactorSMS != nil {
		return json.Marshal(&src.UserFactorSMS)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResendEnrollFactorRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UserFactorCall != nil {
		return obj.UserFactorCall
	}

	if obj.UserFactorEmail != nil {
		return obj.UserFactorEmail
	}

	if obj.UserFactorSMS != nil {
		return obj.UserFactorSMS
	}

	// all schemas are nil
	return nil
}

type NullableResendEnrollFactorRequest struct {
	value *ResendEnrollFactorRequest
	isSet bool
}

func (v NullableResendEnrollFactorRequest) Get() *ResendEnrollFactorRequest {
	return v.value
}

func (v *NullableResendEnrollFactorRequest) Set(val *ResendEnrollFactorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResendEnrollFactorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResendEnrollFactorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResendEnrollFactorRequest(val *ResendEnrollFactorRequest) *NullableResendEnrollFactorRequest {
	return &NullableResendEnrollFactorRequest{value: val, isSet: true}
}

func (v NullableResendEnrollFactorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResendEnrollFactorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


