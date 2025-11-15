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

// AgentJsonSigningKeyResponse - struct for AgentJsonSigningKeyResponse
type AgentJsonSigningKeyResponse struct {
	AgentJsonWebKeyECResponse  *AgentJsonWebKeyECResponse
	AgentJsonWebKeyRsaResponse *AgentJsonWebKeyRsaResponse
}

// AgentJsonWebKeyECResponseAsAgentJsonSigningKeyResponse is a convenience function that returns AgentJsonWebKeyECResponse wrapped in AgentJsonSigningKeyResponse
func AgentJsonWebKeyECResponseAsAgentJsonSigningKeyResponse(v *AgentJsonWebKeyECResponse) AgentJsonSigningKeyResponse {
	return AgentJsonSigningKeyResponse{
		AgentJsonWebKeyECResponse: v,
	}
}

// AgentJsonWebKeyRsaResponseAsAgentJsonSigningKeyResponse is a convenience function that returns AgentJsonWebKeyRsaResponse wrapped in AgentJsonSigningKeyResponse
func AgentJsonWebKeyRsaResponseAsAgentJsonSigningKeyResponse(v *AgentJsonWebKeyRsaResponse) AgentJsonSigningKeyResponse {
	return AgentJsonSigningKeyResponse{
		AgentJsonWebKeyRsaResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AgentJsonSigningKeyResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EC'
	if jsonDict["kty"] == "EC" {
		// try to unmarshal JSON data into AgentJsonWebKeyECResponse
		err = json.Unmarshal(data, &dst.AgentJsonWebKeyECResponse)
		if err == nil {
			return nil // data stored in dst.AgentJsonWebKeyECResponse, return on the first match
		} else {
			dst.AgentJsonWebKeyECResponse = nil
			return fmt.Errorf("failed to unmarshal AgentJsonSigningKeyResponse as AgentJsonWebKeyECResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RSA'
	if jsonDict["kty"] == "RSA" {
		// try to unmarshal JSON data into AgentJsonWebKeyRsaResponse
		err = json.Unmarshal(data, &dst.AgentJsonWebKeyRsaResponse)
		if err == nil {
			return nil // data stored in dst.AgentJsonWebKeyRsaResponse, return on the first match
		} else {
			dst.AgentJsonWebKeyRsaResponse = nil
			return fmt.Errorf("failed to unmarshal AgentJsonSigningKeyResponse as AgentJsonWebKeyRsaResponse: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	if src.AgentJsonWebKeyECResponse != nil {
		return json.Marshal(&src.AgentJsonWebKeyECResponse)
	}

	if src.AgentJsonWebKeyRsaResponse != nil {
		return json.Marshal(&src.AgentJsonWebKeyRsaResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AgentJsonSigningKeyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AgentJsonWebKeyECResponse != nil {
		return obj.AgentJsonWebKeyECResponse
	}

	if obj.AgentJsonWebKeyRsaResponse != nil {
		return obj.AgentJsonWebKeyRsaResponse
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AgentJsonSigningKeyResponse) GetActualInstanceValue() interface{} {
	if obj.AgentJsonWebKeyECResponse != nil {
		return *obj.AgentJsonWebKeyECResponse
	}

	if obj.AgentJsonWebKeyRsaResponse != nil {
		return *obj.AgentJsonWebKeyRsaResponse
	}

	// all schemas are nil
	return nil
}

type NullableAgentJsonSigningKeyResponse struct {
	value *AgentJsonSigningKeyResponse
	isSet bool
}

func (v NullableAgentJsonSigningKeyResponse) Get() *AgentJsonSigningKeyResponse {
	return v.value
}

func (v *NullableAgentJsonSigningKeyResponse) Set(val *AgentJsonSigningKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonSigningKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonSigningKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonSigningKeyResponse(val *AgentJsonSigningKeyResponse) *NullableAgentJsonSigningKeyResponse {
	return &NullableAgentJsonSigningKeyResponse{value: val, isSet: true}
}

func (v NullableAgentJsonSigningKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonSigningKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
