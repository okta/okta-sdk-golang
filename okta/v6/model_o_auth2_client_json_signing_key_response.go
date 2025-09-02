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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)


//model_oneof.mustache
// OAuth2ClientJsonSigningKeyResponse - A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta uses signing keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys for signing tokens.
type OAuth2ClientJsonSigningKeyResponse struct {
	OAuth2ClientJsonWebKeyResponseBase *OAuth2ClientJsonWebKeyResponseBase
}

// OAuth2ClientJsonWebKeyResponseBaseAsOAuth2ClientJsonSigningKeyResponse is a convenience function that returns OAuth2ClientJsonWebKeyResponseBase wrapped in OAuth2ClientJsonSigningKeyResponse
func OAuth2ClientJsonWebKeyResponseBaseAsOAuth2ClientJsonSigningKeyResponse(v *OAuth2ClientJsonWebKeyResponseBase) OAuth2ClientJsonSigningKeyResponse {
	return OAuth2ClientJsonSigningKeyResponse{
		OAuth2ClientJsonWebKeyResponseBase: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *OAuth2ClientJsonSigningKeyResponse) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into OAuth2ClientJsonWebKeyResponseBase
        err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyResponseBase)
        if err == nil {
                jsonOAuth2ClientJsonWebKeyResponseBase, _ := json.Marshal(dst.OAuth2ClientJsonWebKeyResponseBase)
                if string(jsonOAuth2ClientJsonWebKeyResponseBase) == "{}" { // empty struct
                        dst.OAuth2ClientJsonWebKeyResponseBase = nil
                } else {
                        match++
                }
        } else {
                dst.OAuth2ClientJsonWebKeyResponseBase = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.OAuth2ClientJsonWebKeyResponseBase = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(OAuth2ClientJsonSigningKeyResponse)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(OAuth2ClientJsonSigningKeyResponse)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OAuth2ClientJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonWebKeyResponseBase != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyResponseBase)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OAuth2ClientJsonSigningKeyResponse) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.OAuth2ClientJsonWebKeyResponseBase != nil {
		return obj.OAuth2ClientJsonWebKeyResponseBase
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


