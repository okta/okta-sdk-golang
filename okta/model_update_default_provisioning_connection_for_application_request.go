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
// UpdateDefaultProvisioningConnectionForApplicationRequest - struct for UpdateDefaultProvisioningConnectionForApplicationRequest
type UpdateDefaultProvisioningConnectionForApplicationRequest struct {
	ProvisioningConnectionOauthRequest *ProvisioningConnectionOauthRequest
	ProvisioningConnectionTokenRequest *ProvisioningConnectionTokenRequest
}

// ProvisioningConnectionOauthRequestAsUpdateDefaultProvisioningConnectionForApplicationRequest is a convenience function that returns ProvisioningConnectionOauthRequest wrapped in UpdateDefaultProvisioningConnectionForApplicationRequest
func ProvisioningConnectionOauthRequestAsUpdateDefaultProvisioningConnectionForApplicationRequest(v *ProvisioningConnectionOauthRequest) UpdateDefaultProvisioningConnectionForApplicationRequest {
	return UpdateDefaultProvisioningConnectionForApplicationRequest{
		ProvisioningConnectionOauthRequest: v,
	}
}

// ProvisioningConnectionTokenRequestAsUpdateDefaultProvisioningConnectionForApplicationRequest is a convenience function that returns ProvisioningConnectionTokenRequest wrapped in UpdateDefaultProvisioningConnectionForApplicationRequest
func ProvisioningConnectionTokenRequestAsUpdateDefaultProvisioningConnectionForApplicationRequest(v *ProvisioningConnectionTokenRequest) UpdateDefaultProvisioningConnectionForApplicationRequest {
	return UpdateDefaultProvisioningConnectionForApplicationRequest{
		ProvisioningConnectionTokenRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *UpdateDefaultProvisioningConnectionForApplicationRequest) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into ProvisioningConnectionOauthRequest
        err = json.Unmarshal(data, &dst.ProvisioningConnectionOauthRequest)
        if err == nil {
                jsonProvisioningConnectionOauthRequest, _ := json.Marshal(dst.ProvisioningConnectionOauthRequest)
                if string(jsonProvisioningConnectionOauthRequest) == "{}" { // empty struct
                        dst.ProvisioningConnectionOauthRequest = nil
                } else {
                        match++
                }
        } else {
                dst.ProvisioningConnectionOauthRequest = nil
        }

        // try to unmarshal data into ProvisioningConnectionTokenRequest
        err = json.Unmarshal(data, &dst.ProvisioningConnectionTokenRequest)
        if err == nil {
                jsonProvisioningConnectionTokenRequest, _ := json.Marshal(dst.ProvisioningConnectionTokenRequest)
                if string(jsonProvisioningConnectionTokenRequest) == "{}" { // empty struct
                        dst.ProvisioningConnectionTokenRequest = nil
                } else {
                        match++
                }
        } else {
                dst.ProvisioningConnectionTokenRequest = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.ProvisioningConnectionOauthRequest = nil
                dst.ProvisioningConnectionTokenRequest = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(UpdateDefaultProvisioningConnectionForApplicationRequest)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(UpdateDefaultProvisioningConnectionForApplicationRequest)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateDefaultProvisioningConnectionForApplicationRequest) MarshalJSON() ([]byte, error) {
	if src.ProvisioningConnectionOauthRequest != nil {
		return json.Marshal(&src.ProvisioningConnectionOauthRequest)
	}

	if src.ProvisioningConnectionTokenRequest != nil {
		return json.Marshal(&src.ProvisioningConnectionTokenRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateDefaultProvisioningConnectionForApplicationRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ProvisioningConnectionOauthRequest != nil {
		return obj.ProvisioningConnectionOauthRequest
	}

	if obj.ProvisioningConnectionTokenRequest != nil {
		return obj.ProvisioningConnectionTokenRequest
	}

	// all schemas are nil
	return nil
}

type NullableUpdateDefaultProvisioningConnectionForApplicationRequest struct {
	value *UpdateDefaultProvisioningConnectionForApplicationRequest
	isSet bool
}

func (v NullableUpdateDefaultProvisioningConnectionForApplicationRequest) Get() *UpdateDefaultProvisioningConnectionForApplicationRequest {
	return v.value
}

func (v *NullableUpdateDefaultProvisioningConnectionForApplicationRequest) Set(val *UpdateDefaultProvisioningConnectionForApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDefaultProvisioningConnectionForApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDefaultProvisioningConnectionForApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDefaultProvisioningConnectionForApplicationRequest(val *UpdateDefaultProvisioningConnectionForApplicationRequest) *NullableUpdateDefaultProvisioningConnectionForApplicationRequest {
	return &NullableUpdateDefaultProvisioningConnectionForApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateDefaultProvisioningConnectionForApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDefaultProvisioningConnectionForApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


