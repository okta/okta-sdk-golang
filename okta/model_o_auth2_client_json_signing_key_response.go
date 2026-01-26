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

// OAuth2ClientJsonSigningKeyResponse - A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses signing keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys for signing tokens.
type OAuth2ClientJsonSigningKeyResponse struct {
	OAuth2ClientJsonWebKeyECResponse  *OAuth2ClientJsonWebKeyECResponse
	OAuth2ClientJsonWebKeyRsaResponse *OAuth2ClientJsonWebKeyRsaResponse
}

// OAuth2ClientJsonWebKeyECResponseAsOAuth2ClientJsonSigningKeyResponse is a convenience function that returns OAuth2ClientJsonWebKeyECResponse wrapped in OAuth2ClientJsonSigningKeyResponse
func OAuth2ClientJsonWebKeyECResponseAsOAuth2ClientJsonSigningKeyResponse(v *OAuth2ClientJsonWebKeyECResponse) OAuth2ClientJsonSigningKeyResponse {
	return OAuth2ClientJsonSigningKeyResponse{
		OAuth2ClientJsonWebKeyECResponse: v,
	}
}

// OAuth2ClientJsonWebKeyRsaResponseAsOAuth2ClientJsonSigningKeyResponse is a convenience function that returns OAuth2ClientJsonWebKeyRsaResponse wrapped in OAuth2ClientJsonSigningKeyResponse
func OAuth2ClientJsonWebKeyRsaResponseAsOAuth2ClientJsonSigningKeyResponse(v *OAuth2ClientJsonWebKeyRsaResponse) OAuth2ClientJsonSigningKeyResponse {
	return OAuth2ClientJsonSigningKeyResponse{
		OAuth2ClientJsonWebKeyRsaResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OAuth2ClientJsonSigningKeyResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["kty"].(string)

	// check if the discriminator value is 'EC'
	if discriminatorValue == "EC" {
		// try to unmarshal JSON data into OAuth2ClientJsonWebKeyECResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyECResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonWebKeyECResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonWebKeyECResponse = nil
			return fmt.Errorf("failed to unmarshal OAuth2ClientJsonSigningKeyResponse as OAuth2ClientJsonWebKeyECResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RSA'
	if discriminatorValue == "RSA" {
		// try to unmarshal JSON data into OAuth2ClientJsonWebKeyRsaResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyRsaResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonWebKeyRsaResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonWebKeyRsaResponse = nil
			return fmt.Errorf("failed to unmarshal OAuth2ClientJsonSigningKeyResponse as OAuth2ClientJsonWebKeyRsaResponse: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyRsaResponse)
		if err == nil {
			return nil
		}
		dst.OAuth2ClientJsonWebKeyRsaResponse = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OAuth2ClientJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonWebKeyECResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyECResponse)
	}

	if src.OAuth2ClientJsonWebKeyRsaResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyRsaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OAuth2ClientJsonSigningKeyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonWebKeyECResponse != nil {
		return obj.OAuth2ClientJsonWebKeyECResponse
	}

	if obj.OAuth2ClientJsonWebKeyRsaResponse != nil {
		return obj.OAuth2ClientJsonWebKeyRsaResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OAuth2ClientJsonSigningKeyResponse) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonWebKeyECResponse != nil {
		return *obj.OAuth2ClientJsonWebKeyECResponse
	}

	if obj.OAuth2ClientJsonWebKeyRsaResponse != nil {
		return *obj.OAuth2ClientJsonWebKeyRsaResponse
	}

	// all schemas are nil
	return nil
}

type NullableOAuth2ClientJsonSigningKeyResponse struct {
	value *OAuth2ClientJsonSigningKeyResponse
	isSet bool
}

func (v NullableOAuth2ClientJsonSigningKeyResponse) Get() *OAuth2ClientJsonSigningKeyResponse {
	return v.value
}

func (v *NullableOAuth2ClientJsonSigningKeyResponse) Set(val *OAuth2ClientJsonSigningKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonSigningKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonSigningKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonSigningKeyResponse(val *OAuth2ClientJsonSigningKeyResponse) *NullableOAuth2ClientJsonSigningKeyResponse {
	return &NullableOAuth2ClientJsonSigningKeyResponse{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonSigningKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
