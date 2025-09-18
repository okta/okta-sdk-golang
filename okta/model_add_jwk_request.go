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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// AddJwkRequest - struct for AddJwkRequest
type AddJwkRequest struct {
	OAuth2ClientJsonEncryptionKeyRequest *OAuth2ClientJsonEncryptionKeyRequest
	OAuth2ClientJsonSigningKeyRequest    *OAuth2ClientJsonSigningKeyRequest
}

// OAuth2ClientJsonEncryptionKeyRequestAsAddJwkRequest is a convenience function that returns OAuth2ClientJsonEncryptionKeyRequest wrapped in AddJwkRequest
func OAuth2ClientJsonEncryptionKeyRequestAsAddJwkRequest(v *OAuth2ClientJsonEncryptionKeyRequest) AddJwkRequest {
	return AddJwkRequest{
		OAuth2ClientJsonEncryptionKeyRequest: v,
	}
}

// OAuth2ClientJsonSigningKeyRequestAsAddJwkRequest is a convenience function that returns OAuth2ClientJsonSigningKeyRequest wrapped in AddJwkRequest
func OAuth2ClientJsonSigningKeyRequestAsAddJwkRequest(v *OAuth2ClientJsonSigningKeyRequest) AddJwkRequest {
	return AddJwkRequest{
		OAuth2ClientJsonSigningKeyRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddJwkRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OAuth2ClientJsonEncryptionKeyRequest
	err = json.Unmarshal(data, &dst.OAuth2ClientJsonEncryptionKeyRequest)
	if err == nil {
		jsonOAuth2ClientJsonEncryptionKeyRequest, _ := json.Marshal(dst.OAuth2ClientJsonEncryptionKeyRequest)
		if string(jsonOAuth2ClientJsonEncryptionKeyRequest) == "{}" { // empty struct
			dst.OAuth2ClientJsonEncryptionKeyRequest = nil
		} else {
			match++
		}
	} else {
		dst.OAuth2ClientJsonEncryptionKeyRequest = nil
	}

	// try to unmarshal data into OAuth2ClientJsonSigningKeyRequest
	err = json.Unmarshal(data, &dst.OAuth2ClientJsonSigningKeyRequest)
	if err == nil {
		jsonOAuth2ClientJsonSigningKeyRequest, _ := json.Marshal(dst.OAuth2ClientJsonSigningKeyRequest)
		if string(jsonOAuth2ClientJsonSigningKeyRequest) == "{}" { // empty struct
			dst.OAuth2ClientJsonSigningKeyRequest = nil
		} else {
			match++
		}
	} else {
		dst.OAuth2ClientJsonSigningKeyRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OAuth2ClientJsonEncryptionKeyRequest = nil
		dst.OAuth2ClientJsonSigningKeyRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddJwkRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddJwkRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddJwkRequest) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonEncryptionKeyRequest != nil {
		return json.Marshal(&src.OAuth2ClientJsonEncryptionKeyRequest)
	}

	if src.OAuth2ClientJsonSigningKeyRequest != nil {
		return json.Marshal(&src.OAuth2ClientJsonSigningKeyRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddJwkRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonEncryptionKeyRequest != nil {
		return obj.OAuth2ClientJsonEncryptionKeyRequest
	}

	if obj.OAuth2ClientJsonSigningKeyRequest != nil {
		return obj.OAuth2ClientJsonSigningKeyRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddJwkRequest) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonEncryptionKeyRequest != nil {
		return *obj.OAuth2ClientJsonEncryptionKeyRequest
	}

	if obj.OAuth2ClientJsonSigningKeyRequest != nil {
		return *obj.OAuth2ClientJsonSigningKeyRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddJwkRequest struct {
	value *AddJwkRequest
	isSet bool
}

func (v NullableAddJwkRequest) Get() *AddJwkRequest {
	return v.value
}

func (v *NullableAddJwkRequest) Set(val *AddJwkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJwkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJwkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJwkRequest(val *AddJwkRequest) *NullableAddJwkRequest {
	return &NullableAddJwkRequest{value: val, isSet: true}
}

func (v NullableAddJwkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJwkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
