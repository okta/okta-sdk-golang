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
	match := 0
	// try to unmarshal data into ProtocolIdVerification
	err = json.Unmarshal(data, &dst.ProtocolIdVerification)
	if err == nil {
		jsonProtocolIdVerification, _ := json.Marshal(dst.ProtocolIdVerification)
		if string(jsonProtocolIdVerification) == "{}" { // empty struct
			dst.ProtocolIdVerification = nil
		} else {
			match++
		}
	} else {
		dst.ProtocolIdVerification = nil
	}

	// try to unmarshal data into ProtocolMtls
	err = json.Unmarshal(data, &dst.ProtocolMtls)
	if err == nil {
		jsonProtocolMtls, _ := json.Marshal(dst.ProtocolMtls)
		if string(jsonProtocolMtls) == "{}" { // empty struct
			dst.ProtocolMtls = nil
		} else {
			match++
		}
	} else {
		dst.ProtocolMtls = nil
	}

	// try to unmarshal data into ProtocolOAuth
	err = json.Unmarshal(data, &dst.ProtocolOAuth)
	if err == nil {
		jsonProtocolOAuth, _ := json.Marshal(dst.ProtocolOAuth)
		if string(jsonProtocolOAuth) == "{}" { // empty struct
			dst.ProtocolOAuth = nil
		} else {
			match++
		}
	} else {
		dst.ProtocolOAuth = nil
	}

	// try to unmarshal data into ProtocolOidc
	err = json.Unmarshal(data, &dst.ProtocolOidc)
	if err == nil {
		jsonProtocolOidc, _ := json.Marshal(dst.ProtocolOidc)
		if string(jsonProtocolOidc) == "{}" { // empty struct
			dst.ProtocolOidc = nil
		} else {
			match++
		}
	} else {
		dst.ProtocolOidc = nil
	}

	// try to unmarshal data into ProtocolSaml
	err = json.Unmarshal(data, &dst.ProtocolSaml)
	if err == nil {
		jsonProtocolSaml, _ := json.Marshal(dst.ProtocolSaml)
		if string(jsonProtocolSaml) == "{}" { // empty struct
			dst.ProtocolSaml = nil
		} else {
			match++
		}
	} else {
		dst.ProtocolSaml = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ProtocolIdVerification = nil
		dst.ProtocolMtls = nil
		dst.ProtocolOAuth = nil
		dst.ProtocolOidc = nil
		dst.ProtocolSaml = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IdentityProviderProtocol)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IdentityProviderProtocol)")
	}
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
