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

// OpenIdConnectApplicationSettingsClientKeysKeysInner - A JSON Web Key for an OpenID Connect application
type OpenIdConnectApplicationSettingsClientKeysKeysInner struct {
	OAuth2ClientJsonEncryptionKeyResponse *OAuth2ClientJsonEncryptionKeyResponse
	OAuth2ClientJsonSigningKeyResponse    *OAuth2ClientJsonSigningKeyResponse
}

// OAuth2ClientJsonEncryptionKeyResponseAsOpenIdConnectApplicationSettingsClientKeysKeysInner is a convenience function that returns OAuth2ClientJsonEncryptionKeyResponse wrapped in OpenIdConnectApplicationSettingsClientKeysKeysInner
func OAuth2ClientJsonEncryptionKeyResponseAsOpenIdConnectApplicationSettingsClientKeysKeysInner(v *OAuth2ClientJsonEncryptionKeyResponse) OpenIdConnectApplicationSettingsClientKeysKeysInner {
	return OpenIdConnectApplicationSettingsClientKeysKeysInner{
		OAuth2ClientJsonEncryptionKeyResponse: v,
	}
}

// OAuth2ClientJsonSigningKeyResponseAsOpenIdConnectApplicationSettingsClientKeysKeysInner is a convenience function that returns OAuth2ClientJsonSigningKeyResponse wrapped in OpenIdConnectApplicationSettingsClientKeysKeysInner
func OAuth2ClientJsonSigningKeyResponseAsOpenIdConnectApplicationSettingsClientKeysKeysInner(v *OAuth2ClientJsonSigningKeyResponse) OpenIdConnectApplicationSettingsClientKeysKeysInner {
	return OpenIdConnectApplicationSettingsClientKeysKeysInner{
		OAuth2ClientJsonSigningKeyResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OpenIdConnectApplicationSettingsClientKeysKeysInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// Get discriminator value, treating nil/missing as empty string for comparison
	discriminatorValue, _ := jsonDict["use"].(string)

	// check if the discriminator value is 'enc'
	if discriminatorValue == "enc" {
		// try to unmarshal JSON data into OAuth2ClientJsonEncryptionKeyResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonEncryptionKeyResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonEncryptionKeyResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonEncryptionKeyResponse = nil
			return fmt.Errorf("failed to unmarshal OpenIdConnectApplicationSettingsClientKeysKeysInner as OAuth2ClientJsonEncryptionKeyResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sig'
	if discriminatorValue == "sig" {
		// try to unmarshal JSON data into OAuth2ClientJsonSigningKeyResponse
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonSigningKeyResponse)
		if err == nil {
			return nil // data stored in dst.OAuth2ClientJsonSigningKeyResponse, return on the first match
		} else {
			dst.OAuth2ClientJsonSigningKeyResponse = nil
			return fmt.Errorf("failed to unmarshal OpenIdConnectApplicationSettingsClientKeysKeysInner as OAuth2ClientJsonSigningKeyResponse: %s", err.Error())
		}
	}

	// If discriminator value is empty/missing, default to the last mapped model (typically the most common type)
	if discriminatorValue == "" {
		err = json.Unmarshal(data, &dst.OAuth2ClientJsonSigningKeyResponse)
		if err == nil {
			return nil
		}
		dst.OAuth2ClientJsonSigningKeyResponse = nil
	}

	// No match found or unmarshal failed - return nil to allow partial unmarshalling
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OpenIdConnectApplicationSettingsClientKeysKeysInner) MarshalJSON() ([]byte, error) {
	if src.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonEncryptionKeyResponse)
	}

	if src.OAuth2ClientJsonSigningKeyResponse != nil {
		return json.Marshal(&src.OAuth2ClientJsonSigningKeyResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetActualInstance() interface{} {
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
func (obj OpenIdConnectApplicationSettingsClientKeysKeysInner) GetActualInstanceValue() interface{} {
	if obj.OAuth2ClientJsonEncryptionKeyResponse != nil {
		return *obj.OAuth2ClientJsonEncryptionKeyResponse
	}

	if obj.OAuth2ClientJsonSigningKeyResponse != nil {
		return *obj.OAuth2ClientJsonSigningKeyResponse
	}

	// all schemas are nil
	return nil
}

type NullableOpenIdConnectApplicationSettingsClientKeysKeysInner struct {
	value *OpenIdConnectApplicationSettingsClientKeysKeysInner
	isSet bool
}

func (v NullableOpenIdConnectApplicationSettingsClientKeysKeysInner) Get() *OpenIdConnectApplicationSettingsClientKeysKeysInner {
	return v.value
}

func (v *NullableOpenIdConnectApplicationSettingsClientKeysKeysInner) Set(val *OpenIdConnectApplicationSettingsClientKeysKeysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationSettingsClientKeysKeysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationSettingsClientKeysKeysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationSettingsClientKeysKeysInner(val *OpenIdConnectApplicationSettingsClientKeysKeysInner) *NullableOpenIdConnectApplicationSettingsClientKeysKeysInner {
	return &NullableOpenIdConnectApplicationSettingsClientKeysKeysInner{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationSettingsClientKeysKeysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationSettingsClientKeysKeysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
