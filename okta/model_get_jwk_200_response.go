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

// GetJwk200Response - struct for GetJwk200Response
type GetJwk200Response struct {
	OAuth2ClientJsonEncryptionKeyResponse *OAuth2ClientJsonEncryptionKeyResponse
	OAuth2ClientJsonSigningKeyResponse    *OAuth2ClientJsonSigningKeyResponse
}

// OAuth2ClientJsonEncryptionKeyResponseAsGetJwk200Response is a convenience function that returns OAuth2ClientJsonEncryptionKeyResponse wrapped in GetJwk200Response
func OAuth2ClientJsonEncryptionKeyResponseAsGetJwk200Response(v *OAuth2ClientJsonEncryptionKeyResponse) GetJwk200Response {
	return GetJwk200Response{
		OAuth2ClientJsonEncryptionKeyResponse: v,
	}
}

// OAuth2ClientJsonSigningKeyResponseAsGetJwk200Response is a convenience function that returns OAuth2ClientJsonSigningKeyResponse wrapped in GetJwk200Response
func OAuth2ClientJsonSigningKeyResponseAsGetJwk200Response(v *OAuth2ClientJsonSigningKeyResponse) GetJwk200Response {
	return GetJwk200Response{
		OAuth2ClientJsonSigningKeyResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetJwk200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'enc'
	if jsonDict["use"] == "enc" {
		// try to unmarshal JSON data into OAuth2ClientJsonEncryptionKeyResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonEncryptionKeyResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonEncryptionKeyResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonEncryptionKeyResponse = nil
			return fmt.Errorf("failed to unmarshal GetJwk200Response as OAuth2ClientJsonEncryptionKeyResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sig'
	if jsonDict["use"] == "sig" {
		// try to unmarshal JSON data into OAuth2ClientJsonSigningKeyResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonSigningKeyResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonSigningKeyResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonSigningKeyResponse = nil
			return fmt.Errorf("failed to unmarshal GetJwk200Response as OAuth2ClientJsonSigningKeyResponse: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetJwk200Response) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonEncryptionKeyResponse)
	}

	if src.OAuth2ClientJsonSigningKeyResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonSigningKeyResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetJwk200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return obj.OAuth2ClientJsonEncryptionKeyResponse
	}

	if obj.OAuth2ClientJsonSigningKeyResponse != nil {
		return obj.OAuth2ClientJsonSigningKeyResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetJwk200Response) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return *obj.OAuth2ClientJsonEncryptionKeyResponse
	}

	if obj.OAuth2ClientJsonSigningKeyResponse != nil {
		return *obj.OAuth2ClientJsonSigningKeyResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetJwk200Response struct {
	value *GetJwk200Response
	isSet bool
}

func (v NullableGetJwk200Response) Get() *GetJwk200Response {
	return v.value
}

func (v *NullableGetJwk200Response) Set(val *GetJwk200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJwk200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJwk200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJwk200Response(val *GetJwk200Response) *NullableGetJwk200Response {
	return &NullableGetJwk200Response{value: val, isSet: true}
}

func (v NullableGetJwk200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJwk200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
