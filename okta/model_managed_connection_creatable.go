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

// ManagedConnectionCreatable - Create a new connection
type ManagedConnectionCreatable struct {
	IdentityAssertionAppInstanceConnectionCreatable *IdentityAssertionAppInstanceConnectionCreatable
	IdentityAssertionCustomASConnectionCreatable    *IdentityAssertionCustomASConnectionCreatable
	STSServiceAccountConnectionCreatable            *STSServiceAccountConnectionCreatable
	STSVaultSecretConnectionCreatable               *STSVaultSecretConnectionCreatable
}

// IdentityAssertionAppInstanceConnectionCreatableAsManagedConnectionCreatable is a convenience function that returns IdentityAssertionAppInstanceConnectionCreatable wrapped in ManagedConnectionCreatable
func IdentityAssertionAppInstanceConnectionCreatableAsManagedConnectionCreatable(v *IdentityAssertionAppInstanceConnectionCreatable) ManagedConnectionCreatable {
	return ManagedConnectionCreatable{
		IdentityAssertionAppInstanceConnectionCreatable: v,
	}
}

// IdentityAssertionCustomASConnectionCreatableAsManagedConnectionCreatable is a convenience function that returns IdentityAssertionCustomASConnectionCreatable wrapped in ManagedConnectionCreatable
func IdentityAssertionCustomASConnectionCreatableAsManagedConnectionCreatable(v *IdentityAssertionCustomASConnectionCreatable) ManagedConnectionCreatable {
	return ManagedConnectionCreatable{
		IdentityAssertionCustomASConnectionCreatable: v,
	}
}

// STSServiceAccountConnectionCreatableAsManagedConnectionCreatable is a convenience function that returns STSServiceAccountConnectionCreatable wrapped in ManagedConnectionCreatable
func STSServiceAccountConnectionCreatableAsManagedConnectionCreatable(v *STSServiceAccountConnectionCreatable) ManagedConnectionCreatable {
	return ManagedConnectionCreatable{
		STSServiceAccountConnectionCreatable: v,
	}
}

// STSVaultSecretConnectionCreatableAsManagedConnectionCreatable is a convenience function that returns STSVaultSecretConnectionCreatable wrapped in ManagedConnectionCreatable
func STSVaultSecretConnectionCreatableAsManagedConnectionCreatable(v *STSVaultSecretConnectionCreatable) ManagedConnectionCreatable {
	return ManagedConnectionCreatable{
		STSVaultSecretConnectionCreatable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ManagedConnectionCreatable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_APP_INSTANCE'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_APP_INSTANCE" {
		// try to unmarshal JSON data into IdentityAssertionAppInstanceConnectionCreatable
		err = json.Unmarshal(data, &dst.IdentityAssertionAppInstanceConnectionCreatable)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionAppInstanceConnectionCreatable, return on the first match
		} else {
			dst.IdentityAssertionAppInstanceConnectionCreatable = nil
			return fmt.Errorf("failed to unmarshal ManagedConnectionCreatable as IdentityAssertionAppInstanceConnectionCreatable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_CUSTOM_AS'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_CUSTOM_AS" {
		// try to unmarshal JSON data into IdentityAssertionCustomASConnectionCreatable
		err = json.Unmarshal(data, &dst.IdentityAssertionCustomASConnectionCreatable)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionCustomASConnectionCreatable, return on the first match
		} else {
			dst.IdentityAssertionCustomASConnectionCreatable = nil
			return fmt.Errorf("failed to unmarshal ManagedConnectionCreatable as IdentityAssertionCustomASConnectionCreatable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_SERVICE_ACCOUNT'
	if jsonDict["connectionType"] == "STS_SERVICE_ACCOUNT" {
		// try to unmarshal JSON data into STSServiceAccountConnectionCreatable
		err = json.Unmarshal(data, &dst.STSServiceAccountConnectionCreatable)
		if err == nil {
			return nil // data stored in dst.STSServiceAccountConnectionCreatable, return on the first match
		} else {
			dst.STSServiceAccountConnectionCreatable = nil
			return fmt.Errorf("failed to unmarshal ManagedConnectionCreatable as STSServiceAccountConnectionCreatable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_VAULT_SECRET'
	if jsonDict["connectionType"] == "STS_VAULT_SECRET" {
		// try to unmarshal JSON data into STSVaultSecretConnectionCreatable
		err = json.Unmarshal(data, &dst.STSVaultSecretConnectionCreatable)
		if err == nil {
			return nil // data stored in dst.STSVaultSecretConnectionCreatable, return on the first match
		} else {
			dst.STSVaultSecretConnectionCreatable = nil
			return fmt.Errorf("failed to unmarshal ManagedConnectionCreatable as STSVaultSecretConnectionCreatable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ManagedConnectionCreatable) MarshalJSON() ([]byte, error) {
	if src.IdentityAssertionAppInstanceConnectionCreatable != nil {
		return json.Marshal(&src.IdentityAssertionAppInstanceConnectionCreatable)
	}

	if src.IdentityAssertionCustomASConnectionCreatable != nil {
		return json.Marshal(&src.IdentityAssertionCustomASConnectionCreatable)
	}

	if src.STSServiceAccountConnectionCreatable != nil {
		return json.Marshal(&src.STSServiceAccountConnectionCreatable)
	}

	if src.STSVaultSecretConnectionCreatable != nil {
		return json.Marshal(&src.STSVaultSecretConnectionCreatable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ManagedConnectionCreatable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IdentityAssertionAppInstanceConnectionCreatable != nil {
		return obj.IdentityAssertionAppInstanceConnectionCreatable
	}

	if obj.IdentityAssertionCustomASConnectionCreatable != nil {
		return obj.IdentityAssertionCustomASConnectionCreatable
	}

	if obj.STSServiceAccountConnectionCreatable != nil {
		return obj.STSServiceAccountConnectionCreatable
	}

	if obj.STSVaultSecretConnectionCreatable != nil {
		return obj.STSVaultSecretConnectionCreatable
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ManagedConnectionCreatable) GetActualInstanceValue() interface{} {
	if obj.IdentityAssertionAppInstanceConnectionCreatable != nil {
		return *obj.IdentityAssertionAppInstanceConnectionCreatable
	}

	if obj.IdentityAssertionCustomASConnectionCreatable != nil {
		return *obj.IdentityAssertionCustomASConnectionCreatable
	}

	if obj.STSServiceAccountConnectionCreatable != nil {
		return *obj.STSServiceAccountConnectionCreatable
	}

	if obj.STSVaultSecretConnectionCreatable != nil {
		return *obj.STSVaultSecretConnectionCreatable
	}

	// all schemas are nil
	return nil
}

type NullableManagedConnectionCreatable struct {
	value *ManagedConnectionCreatable
	isSet bool
}

func (v NullableManagedConnectionCreatable) Get() *ManagedConnectionCreatable {
	return v.value
}

func (v *NullableManagedConnectionCreatable) Set(val *ManagedConnectionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedConnectionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedConnectionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedConnectionCreatable(val *ManagedConnectionCreatable) *NullableManagedConnectionCreatable {
	return &NullableManagedConnectionCreatable{value: val, isSet: true}
}

func (v NullableManagedConnectionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedConnectionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
