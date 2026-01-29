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

// OAuth2ClientJsonSigningKeyRequest - A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses signing keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys for signing tokens.
type OAuth2ClientJsonSigningKeyRequest struct {
	OAuth2ClientJsonWebKeyECRequest  *OAuth2ClientJsonWebKeyECRequest
	OAuth2ClientJsonWebKeyRsaRequest *OAuth2ClientJsonWebKeyRsaRequest
}

// OAuth2ClientJsonWebKeyECRequestAsOAuth2ClientJsonSigningKeyRequest is a convenience function that returns OAuth2ClientJsonWebKeyECRequest wrapped in OAuth2ClientJsonSigningKeyRequest
func OAuth2ClientJsonWebKeyECRequestAsOAuth2ClientJsonSigningKeyRequest(v *OAuth2ClientJsonWebKeyECRequest) OAuth2ClientJsonSigningKeyRequest {
	return OAuth2ClientJsonSigningKeyRequest{
		OAuth2ClientJsonWebKeyECRequest: v,
	}
}

// OAuth2ClientJsonWebKeyRsaRequestAsOAuth2ClientJsonSigningKeyRequest is a convenience function that returns OAuth2ClientJsonWebKeyRsaRequest wrapped in OAuth2ClientJsonSigningKeyRequest
func OAuth2ClientJsonWebKeyRsaRequestAsOAuth2ClientJsonSigningKeyRequest(v *OAuth2ClientJsonWebKeyRsaRequest) OAuth2ClientJsonSigningKeyRequest {
	return OAuth2ClientJsonSigningKeyRequest{
		OAuth2ClientJsonWebKeyRsaRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OAuth2ClientJsonSigningKeyRequest) UnmarshalJSON(data []byte) error {
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
		// try to unmarshal JSON data into OAuth2ClientJsonWebKeyECRequest
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyECRequest)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonWebKeyECRequest, return on the first match
		} else {
			dst.OAuth2ClientJsonWebKeyECRequest = nil
			return fmt.Errorf("failed to unmarshal OAuth2ClientJsonSigningKeyRequest as OAuth2ClientJsonWebKeyECRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RSA'
	if discriminatorValue == "RSA" {
		// try to unmarshal JSON data into OAuth2ClientJsonWebKeyRsaRequest
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyRsaRequest)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonWebKeyRsaRequest, return on the first match
		} else {
			dst.OAuth2ClientJsonWebKeyRsaRequest = nil
			return fmt.Errorf("failed to unmarshal OAuth2ClientJsonSigningKeyRequest as OAuth2ClientJsonWebKeyRsaRequest: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyRsaRequest)
		if err == nil {
			return nil
		}
		dst.OAuth2ClientJsonWebKeyRsaRequest = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OAuth2ClientJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonWebKeyECRequest != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyECRequest)
	}

	if src.OAuth2ClientJsonWebKeyRsaRequest != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyRsaRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OAuth2ClientJsonSigningKeyRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonWebKeyECRequest != nil {
		return obj.OAuth2ClientJsonWebKeyECRequest
	}

	if obj.OAuth2ClientJsonWebKeyRsaRequest != nil {
		return obj.OAuth2ClientJsonWebKeyRsaRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OAuth2ClientJsonSigningKeyRequest) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonWebKeyECRequest != nil {
		return *obj.OAuth2ClientJsonWebKeyECRequest
	}

	if obj.OAuth2ClientJsonWebKeyRsaRequest != nil {
		return *obj.OAuth2ClientJsonWebKeyRsaRequest
	}

	// all schemas are nil
	return nil
}

type NullableOAuth2ClientJsonSigningKeyRequest struct {
	value *OAuth2ClientJsonSigningKeyRequest
	isSet bool
}

func (v NullableOAuth2ClientJsonSigningKeyRequest) Get() *OAuth2ClientJsonSigningKeyRequest {
	return v.value
}

func (v *NullableOAuth2ClientJsonSigningKeyRequest) Set(val *OAuth2ClientJsonSigningKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientJsonSigningKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientJsonSigningKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientJsonSigningKeyRequest(val *OAuth2ClientJsonSigningKeyRequest) *NullableOAuth2ClientJsonSigningKeyRequest {
	return &NullableOAuth2ClientJsonSigningKeyRequest{value: val, isSet: true}
}

func (v NullableOAuth2ClientJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientJsonSigningKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
