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

// ManagedConnection - struct for ManagedConnection
type ManagedConnection struct {
	IdentityAssertionAppInstanceConnection *IdentityAssertionAppInstanceConnection
	IdentityAssertionCustomASConnection    *IdentityAssertionCustomASConnection
	STSServiceAccountConnection            *STSServiceAccountConnection
	STSVaultSecretConnection               *STSVaultSecretConnection
}

// IdentityAssertionAppInstanceConnectionAsManagedConnection is a convenience function that returns IdentityAssertionAppInstanceConnection wrapped in ManagedConnection
func IdentityAssertionAppInstanceConnectionAsManagedConnection(v *IdentityAssertionAppInstanceConnection) ManagedConnection {
	return ManagedConnection{
		IdentityAssertionAppInstanceConnection: v,
	}
}

// IdentityAssertionCustomASConnectionAsManagedConnection is a convenience function that returns IdentityAssertionCustomASConnection wrapped in ManagedConnection
func IdentityAssertionCustomASConnectionAsManagedConnection(v *IdentityAssertionCustomASConnection) ManagedConnection {
	return ManagedConnection{
		IdentityAssertionCustomASConnection: v,
	}
}

// STSServiceAccountConnectionAsManagedConnection is a convenience function that returns STSServiceAccountConnection wrapped in ManagedConnection
func STSServiceAccountConnectionAsManagedConnection(v *STSServiceAccountConnection) ManagedConnection {
	return ManagedConnection{
		STSServiceAccountConnection: v,
	}
}

// STSVaultSecretConnectionAsManagedConnection is a convenience function that returns STSVaultSecretConnection wrapped in ManagedConnection
func STSVaultSecretConnectionAsManagedConnection(v *STSVaultSecretConnection) ManagedConnection {
	return ManagedConnection{
		STSVaultSecretConnection: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ManagedConnection) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_APP_INSTANCE'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_APP_INSTANCE" {
		// try to unmarshal JSON data into IdentityAssertionAppInstanceConnection
		err = json.Unmarshal(data, &dst.IdentityAssertionAppInstanceConnection)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionAppInstanceConnection, return on the first match
		} else {
			dst.IdentityAssertionAppInstanceConnection = nil
			return fmt.Errorf("failed to unmarshal ManagedConnection as IdentityAssertionAppInstanceConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_CUSTOM_AS'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_CUSTOM_AS" {
		// try to unmarshal JSON data into IdentityAssertionCustomASConnection
		err = json.Unmarshal(data, &dst.IdentityAssertionCustomASConnection)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionCustomASConnection, return on the first match
		} else {
			dst.IdentityAssertionCustomASConnection = nil
			return fmt.Errorf("failed to unmarshal ManagedConnection as IdentityAssertionCustomASConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_SERVICE_ACCOUNT'
	if jsonDict["connectionType"] == "STS_SERVICE_ACCOUNT" {
		// try to unmarshal JSON data into STSServiceAccountConnection
		err = json.Unmarshal(data, &dst.STSServiceAccountConnection)
		if err == nil {
			return nil // data stored in dst.STSServiceAccountConnection, return on the first match
		} else {
			dst.STSServiceAccountConnection = nil
			return fmt.Errorf("failed to unmarshal ManagedConnection as STSServiceAccountConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_VAULT_SECRET'
	if jsonDict["connectionType"] == "STS_VAULT_SECRET" {
		// try to unmarshal JSON data into STSVaultSecretConnection
		err = json.Unmarshal(data, &dst.STSVaultSecretConnection)
		if err == nil {
			return nil // data stored in dst.STSVaultSecretConnection, return on the first match
		} else {
			dst.STSVaultSecretConnection = nil
			return fmt.Errorf("failed to unmarshal ManagedConnection as STSVaultSecretConnection: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ManagedConnection) MarshalJSON() ([]byte, error) {
	if src.IdentityAssertionAppInstanceConnection != nil {
		return json.Marshal(&src.IdentityAssertionAppInstanceConnection)
	}

	if src.IdentityAssertionCustomASConnection != nil {
		return json.Marshal(&src.IdentityAssertionCustomASConnection)
	}

	if src.STSServiceAccountConnection != nil {
		return json.Marshal(&src.STSServiceAccountConnection)
	}

	if src.STSVaultSecretConnection != nil {
		return json.Marshal(&src.STSVaultSecretConnection)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ManagedConnection) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IdentityAssertionAppInstanceConnection != nil {
		return obj.IdentityAssertionAppInstanceConnection
	}

	if obj.IdentityAssertionCustomASConnection != nil {
		return obj.IdentityAssertionCustomASConnection
	}

	if obj.STSServiceAccountConnection != nil {
		return obj.STSServiceAccountConnection
	}

	if obj.STSVaultSecretConnection != nil {
		return obj.STSVaultSecretConnection
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ManagedConnection) GetActualInstanceValue() interface{} {
	if obj.IdentityAssertionAppInstanceConnection != nil {
		return *obj.IdentityAssertionAppInstanceConnection
	}

	if obj.IdentityAssertionCustomASConnection != nil {
		return *obj.IdentityAssertionCustomASConnection
	}

	if obj.STSServiceAccountConnection != nil {
		return *obj.STSServiceAccountConnection
	}

	if obj.STSVaultSecretConnection != nil {
		return *obj.STSVaultSecretConnection
	}

	// all schemas are nil
	return nil
}

type NullableManagedConnection struct {
	value *ManagedConnection
	isSet bool
}

func (v NullableManagedConnection) Get() *ManagedConnection {
	return v.value
}

func (v *NullableManagedConnection) Set(val *ManagedConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedConnection(val *ManagedConnection) *NullableManagedConnection {
	return &NullableManagedConnection{value: val, isSet: true}
}

func (v NullableManagedConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
