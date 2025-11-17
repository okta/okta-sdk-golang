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

// DeactivateOAuth2ClientJsonWebKey200Response - struct for DeactivateOAuth2ClientJsonWebKey200Response
type DeactivateOAuth2ClientJsonWebKey200Response struct {
	OAuth2ClientJsonWebKeyECResponse  *OAuth2ClientJsonWebKeyECResponse
	OAuth2ClientJsonWebKeyRsaResponse *OAuth2ClientJsonWebKeyRsaResponse
}

// OAuth2ClientJsonWebKeyECResponseAsDeactivateOAuth2ClientJsonWebKey200Response is a convenience function that returns OAuth2ClientJsonWebKeyECResponse wrapped in DeactivateOAuth2ClientJsonWebKey200Response
func OAuth2ClientJsonWebKeyECResponseAsDeactivateOAuth2ClientJsonWebKey200Response(v *OAuth2ClientJsonWebKeyECResponse) DeactivateOAuth2ClientJsonWebKey200Response {
	return DeactivateOAuth2ClientJsonWebKey200Response{
		OAuth2ClientJsonWebKeyECResponse: v,
	}
}

// OAuth2ClientJsonWebKeyRsaResponseAsDeactivateOAuth2ClientJsonWebKey200Response is a convenience function that returns OAuth2ClientJsonWebKeyRsaResponse wrapped in DeactivateOAuth2ClientJsonWebKey200Response
func OAuth2ClientJsonWebKeyRsaResponseAsDeactivateOAuth2ClientJsonWebKey200Response(v *OAuth2ClientJsonWebKeyRsaResponse) DeactivateOAuth2ClientJsonWebKey200Response {
	return DeactivateOAuth2ClientJsonWebKey200Response{
		OAuth2ClientJsonWebKeyRsaResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeactivateOAuth2ClientJsonWebKey200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EC'
	if jsonDict["kty"] == "EC" {
		// try to unmarshal JSON data into OAuth2ClientJsonWebKeyECResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyECResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonWebKeyECResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonWebKeyECResponse = nil
			return fmt.Errorf("failed to unmarshal DeactivateOAuth2ClientJsonWebKey200Response as OAuth2ClientJsonWebKeyECResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RSA'
	if jsonDict["kty"] == "RSA" {
		// try to unmarshal JSON data into OAuth2ClientJsonWebKeyRsaResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonWebKeyRsaResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonWebKeyRsaResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonWebKeyRsaResponse = nil
			return fmt.Errorf("failed to unmarshal DeactivateOAuth2ClientJsonWebKey200Response as OAuth2ClientJsonWebKeyRsaResponse: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeactivateOAuth2ClientJsonWebKey200Response) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonWebKeyECResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyECResponse)
	}

	if src.OAuth2ClientJsonWebKeyRsaResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonWebKeyRsaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeactivateOAuth2ClientJsonWebKey200Response) GetActualInstance() interface{} {
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
func (obj DeactivateOAuth2ClientJsonWebKey200Response) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonWebKeyECResponse != nil {
		return *obj.OAuth2ClientJsonWebKeyECResponse
	}

	if obj.OAuth2ClientJsonWebKeyRsaResponse != nil {
		return *obj.OAuth2ClientJsonWebKeyRsaResponse
	}

	// all schemas are nil
	return nil
}

type NullableDeactivateOAuth2ClientJsonWebKey200Response struct {
	value *DeactivateOAuth2ClientJsonWebKey200Response
	isSet bool
}

func (v NullableDeactivateOAuth2ClientJsonWebKey200Response) Get() *DeactivateOAuth2ClientJsonWebKey200Response {
	return v.value
}

func (v *NullableDeactivateOAuth2ClientJsonWebKey200Response) Set(val *DeactivateOAuth2ClientJsonWebKey200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeactivateOAuth2ClientJsonWebKey200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeactivateOAuth2ClientJsonWebKey200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeactivateOAuth2ClientJsonWebKey200Response(val *DeactivateOAuth2ClientJsonWebKey200Response) *NullableDeactivateOAuth2ClientJsonWebKey200Response {
	return &NullableDeactivateOAuth2ClientJsonWebKey200Response{value: val, isSet: true}
}

func (v NullableDeactivateOAuth2ClientJsonWebKey200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeactivateOAuth2ClientJsonWebKey200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
