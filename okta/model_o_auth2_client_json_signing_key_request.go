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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// OAuth2ClientJsonSigningKeyRequest - A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses signing keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys for signing tokens.
type OAuth2ClientJsonSigningKeyRequest struct {
	OAuth2ClientJsonWebKeyRequestBase *OAuth2ClientJsonWebKeyRequestBase
}

// OAuth2ClientJsonWebKeyRequestBaseAsOAuth2ClientJsonSigningKeyRequest is a convenience function that returns OAuth2ClientJsonWebKeyRequestBase wrapped in OAuth2ClientJsonSigningKeyRequest
func OAuth2ClientJsonWebKeyRequestBaseAsOAuth2ClientJsonSigningKeyRequest(v *OAuth2ClientJsonWebKeyRequestBase) OAuth2ClientJsonSigningKeyRequest {
	return OAuth2ClientJsonSigningKeyRequest{
		OAuth2ClientJsonWebKeyRequestBase: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OAuth2ClientJsonSigningKeyRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OAuth2ClientJsonWebKeyRequestBase
	err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyRequestBase)
	if err == nil {
		jsonOAuth2ClientJsonWebKeyRequestBase, _ := json.Marshal(dst.OAuth2ClientJsonWebKeyRequestBase)
		if string(jsonOAuth2ClientJsonWebKeyRequestBase) == "{}" { // empty struct
			dst.OAuth2ClientJsonWebKeyRequestBase = nil
		} else {
			match++
		}
	} else {
		dst.OAuth2ClientJsonWebKeyRequestBase = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OAuth2ClientJsonWebKeyRequestBase = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OAuth2ClientJsonSigningKeyRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OAuth2ClientJsonSigningKeyRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OAuth2ClientJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonWebKeyRequestBase != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyRequestBase)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OAuth2ClientJsonSigningKeyRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonWebKeyRequestBase != nil {
		return obj.OAuth2ClientJsonWebKeyRequestBase
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OAuth2ClientJsonSigningKeyRequest) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonWebKeyRequestBase != nil {
		return *obj.OAuth2ClientJsonWebKeyRequestBase
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
