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

// AgentJsonSigningKeyRequest - struct for AgentJsonSigningKeyRequest
type AgentJsonSigningKeyRequest struct {
	AgentJsonWebKeyECRequest  *AgentJsonWebKeyECRequest
	AgentJsonWebKeyRsaRequest *AgentJsonWebKeyRsaRequest
}

// AgentJsonWebKeyECRequestAsAgentJsonSigningKeyRequest is a convenience function that returns AgentJsonWebKeyECRequest wrapped in AgentJsonSigningKeyRequest
func AgentJsonWebKeyECRequestAsAgentJsonSigningKeyRequest(v *AgentJsonWebKeyECRequest) AgentJsonSigningKeyRequest {
	return AgentJsonSigningKeyRequest{
		AgentJsonWebKeyECRequest: v,
	}
}

// AgentJsonWebKeyRsaRequestAsAgentJsonSigningKeyRequest is a convenience function that returns AgentJsonWebKeyRsaRequest wrapped in AgentJsonSigningKeyRequest
func AgentJsonWebKeyRsaRequestAsAgentJsonSigningKeyRequest(v *AgentJsonWebKeyRsaRequest) AgentJsonSigningKeyRequest {
	return AgentJsonSigningKeyRequest{
		AgentJsonWebKeyRsaRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AgentJsonSigningKeyRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EC'
	if jsonDict["kty"] == "EC" {
		// try to unmarshal JSON data into AgentJsonWebKeyECRequest
		err = json.Unmarshal(data, &dst.AgentJsonWebKeyECRequest)
		if err == nil {
			return nil // data stored in dst.AgentJsonWebKeyECRequest, return on the first match
		} else {
			dst.AgentJsonWebKeyECRequest = nil
			return fmt.Errorf("failed to unmarshal AgentJsonSigningKeyRequest as AgentJsonWebKeyECRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RSA'
	if jsonDict["kty"] == "RSA" {
		// try to unmarshal JSON data into AgentJsonWebKeyRsaRequest
		err = json.Unmarshal(data, &dst.AgentJsonWebKeyRsaRequest)
		if err == nil {
			return nil // data stored in dst.AgentJsonWebKeyRsaRequest, return on the first match
		} else {
			dst.AgentJsonWebKeyRsaRequest = nil
			return fmt.Errorf("failed to unmarshal AgentJsonSigningKeyRequest as AgentJsonWebKeyRsaRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	if src.AgentJsonWebKeyECRequest != nil {
		return json.Marshal(&src.AgentJsonWebKeyECRequest)
	}

	if src.AgentJsonWebKeyRsaRequest != nil {
		return json.Marshal(&src.AgentJsonWebKeyRsaRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AgentJsonSigningKeyRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AgentJsonWebKeyECRequest != nil {
		return obj.AgentJsonWebKeyECRequest
	}

	if obj.AgentJsonWebKeyRsaRequest != nil {
		return obj.AgentJsonWebKeyRsaRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AgentJsonSigningKeyRequest) GetActualInstanceValue() interface{} {
	if obj.AgentJsonWebKeyECRequest != nil {
		return *obj.AgentJsonWebKeyECRequest
	}

	if obj.AgentJsonWebKeyRsaRequest != nil {
		return *obj.AgentJsonWebKeyRsaRequest
	}

	// all schemas are nil
	return nil
}

type NullableAgentJsonSigningKeyRequest struct {
	value *AgentJsonSigningKeyRequest
	isSet bool
}

func (v NullableAgentJsonSigningKeyRequest) Get() *AgentJsonSigningKeyRequest {
	return v.value
}

func (v *NullableAgentJsonSigningKeyRequest) Set(val *AgentJsonSigningKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonSigningKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonSigningKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonSigningKeyRequest(val *AgentJsonSigningKeyRequest) *NullableAgentJsonSigningKeyRequest {
	return &NullableAgentJsonSigningKeyRequest{value: val, isSet: true}
}

func (v NullableAgentJsonSigningKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonSigningKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
