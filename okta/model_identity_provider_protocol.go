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

// IdentityProviderProtocol - IdP-specific protocol settings for endpoints, bindings, and algorithms used to connect with the IdP and validate messages
type IdentityProviderProtocol struct {
	ProtocolIdVerification *ProtocolIdVerification
	ProtocolMtls           *ProtocolMtls
	ProtocolOAuth          *ProtocolOAuth
	ProtocolOidc           *ProtocolOidc
	ProtocolSaml           *ProtocolSaml
}

// ProtocolIdVerificationAsIdentityProviderProtocol is a convenience function that returns ProtocolIdVerification wrapped in IdentityProviderProtocol
func ProtocolIdVerificationAsIdentityProviderProtocol(v *ProtocolIdVerification) IdentityProviderProtocol {
	return IdentityProviderProtocol{
		ProtocolIdVerification: v,
	}
}

// ProtocolMtlsAsIdentityProviderProtocol is a convenience function that returns ProtocolMtls wrapped in IdentityProviderProtocol
func ProtocolMtlsAsIdentityProviderProtocol(v *ProtocolMtls) IdentityProviderProtocol {
	return IdentityProviderProtocol{
		ProtocolMtls: v,
	}
}

// ProtocolOAuthAsIdentityProviderProtocol is a convenience function that returns ProtocolOAuth wrapped in IdentityProviderProtocol
func ProtocolOAuthAsIdentityProviderProtocol(v *ProtocolOAuth) IdentityProviderProtocol {
	return IdentityProviderProtocol{
		ProtocolOAuth: v,
	}
}

// ProtocolOidcAsIdentityProviderProtocol is a convenience function that returns ProtocolOidc wrapped in IdentityProviderProtocol
func ProtocolOidcAsIdentityProviderProtocol(v *ProtocolOidc) IdentityProviderProtocol {
	return IdentityProviderProtocol{
		ProtocolOidc: v,
	}
}

// ProtocolSamlAsIdentityProviderProtocol is a convenience function that returns ProtocolSaml wrapped in IdentityProviderProtocol
func ProtocolSamlAsIdentityProviderProtocol(v *ProtocolSaml) IdentityProviderProtocol {
	return IdentityProviderProtocol{
		ProtocolSaml: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityProviderProtocol) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ID_PROOFING'
	if jsonDict["type"] == "ID_PROOFING" {
		// try to unmarshal JSON data into ProtocolIdVerification
		err = json.Unmarshal(data, &dst.ProtocolIdVerification)
		if err == nil {
			return nil // data stored in dst.ProtocolIdVerification, return on the first match
		} else {
			dst.ProtocolIdVerification = nil
			return fmt.Errorf("failed to unmarshal IdentityProviderProtocol as ProtocolIdVerification: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MTLS'
	if jsonDict["type"] == "MTLS" {
		// try to unmarshal JSON data into ProtocolMtls
		err = json.Unmarshal(data, &dst.ProtocolMtls)
		if err == nil {
			return nil // data stored in dst.ProtocolMtls, return on the first match
		} else {
			dst.ProtocolMtls = nil
			return fmt.Errorf("failed to unmarshal IdentityProviderProtocol as ProtocolMtls: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OAUTH2'
	if jsonDict["type"] == "OAUTH2" {
		// try to unmarshal JSON data into ProtocolOAuth
		err = json.Unmarshal(data, &dst.ProtocolOAuth)
		if err == nil {
			return nil // data stored in dst.ProtocolOAuth, return on the first match
		} else {
			dst.ProtocolOAuth = nil
			return fmt.Errorf("failed to unmarshal IdentityProviderProtocol as ProtocolOAuth: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OIDC'
	if jsonDict["type"] == "OIDC" {
		// try to unmarshal JSON data into ProtocolOidc
		err = json.Unmarshal(data, &dst.ProtocolOidc)
		if err == nil {
			return nil // data stored in dst.ProtocolOidc, return on the first match
		} else {
			dst.ProtocolOidc = nil
			return fmt.Errorf("failed to unmarshal IdentityProviderProtocol as ProtocolOidc: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SAML2'
	if jsonDict["type"] == "SAML2" {
		// try to unmarshal JSON data into ProtocolSaml
		err = json.Unmarshal(data, &dst.ProtocolSaml)
		if err == nil {
			return nil // data stored in dst.ProtocolSaml, return on the first match
		} else {
			dst.ProtocolSaml = nil
			return fmt.Errorf("failed to unmarshal IdentityProviderProtocol as ProtocolSaml: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityProviderProtocol) MarshalJSON() ([]byte, error) {
	if src.ProtocolIdVerification != nil {
		return json.Marshal(&src.ProtocolIdVerification)
	}

	if src.ProtocolMtls != nil {
		return json.Marshal(&src.ProtocolMtls)
	}

	if src.ProtocolOAuth != nil {
		return json.Marshal(&src.ProtocolOAuth)
	}

	if src.ProtocolOidc != nil {
		return json.Marshal(&src.ProtocolOidc)
	}

	if src.ProtocolSaml != nil {
		return json.Marshal(&src.ProtocolSaml)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityProviderProtocol) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ProtocolIdVerification != nil {
		return obj.ProtocolIdVerification
	}

	if obj.ProtocolMtls != nil {
		return obj.ProtocolMtls
	}

	if obj.ProtocolOAuth != nil {
		return obj.ProtocolOAuth
	}

	if obj.ProtocolOidc != nil {
		return obj.ProtocolOidc
	}

	if obj.ProtocolSaml != nil {
		return obj.ProtocolSaml
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj IdentityProviderProtocol) GetActualInstanceValue() interface{} {
	if obj.ProtocolIdVerification != nil {
		return *obj.ProtocolIdVerification
	}

	if obj.ProtocolMtls != nil {
		return *obj.ProtocolMtls
	}

	if obj.ProtocolOAuth != nil {
		return *obj.ProtocolOAuth
	}

	if obj.ProtocolOidc != nil {
		return *obj.ProtocolOidc
	}

	if obj.ProtocolSaml != nil {
		return *obj.ProtocolSaml
	}

	// all schemas are nil
	return nil
}

type NullableIdentityProviderProtocol struct {
	value *IdentityProviderProtocol
	isSet bool
}

func (v NullableIdentityProviderProtocol) Get() *IdentityProviderProtocol {
	return v.value
}

func (v *NullableIdentityProviderProtocol) Set(val *IdentityProviderProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderProtocol(val *IdentityProviderProtocol) *NullableIdentityProviderProtocol {
	return &NullableIdentityProviderProtocol{value: val, isSet: true}
}

func (v NullableIdentityProviderProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
