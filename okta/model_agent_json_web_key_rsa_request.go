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

// checks if the AgentJsonWebKeyRsaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonWebKeyRsaRequest{}

// AgentJsonWebKeyRsaRequest An RSA signing key
type AgentJsonWebKeyRsaRequest struct {
	// The public exponent of the RSA key, represented as a Base64URL-encoded string.  This value is used in combination with the modulus (`n`) to verify signatures and encrypt data.
	E *string `json:"e,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty string `json:"kty"`
	// The modulus of the RSA public key, represented as a Base64URL-encoded string.  This is the primary component of the RSA key and, with the exponent (`e`), is used for cryptographic signature verification and encryption.
	N *string `json:"n,omitempty"`
	// Unique identifier of the JSON Web Key in the AI agent's JSON Web Key Set (JWKS)
	Kid *string `json:"kid,omitempty"`
	// Status of the AI agent JSON Web Key
	Status *string `json:"status,omitempty"`
	// Algorithm that's used in the JSON Web Key
	Alg *string `json:"alg,omitempty"`
	// Acceptable use of the JSON Web Key  You can only use signing keys for AI agents, so the value of `use` is always `sig`.
	Use                  *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentJsonWebKeyRsaRequest AgentJsonWebKeyRsaRequest

// NewAgentJsonWebKeyRsaRequest instantiates a new AgentJsonWebKeyRsaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonWebKeyRsaRequest(kty string) *AgentJsonWebKeyRsaRequest {
	this := AgentJsonWebKeyRsaRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewAgentJsonWebKeyRsaRequestWithDefaults instantiates a new AgentJsonWebKeyRsaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonWebKeyRsaRequestWithDefaults() *AgentJsonWebKeyRsaRequest {
	this := AgentJsonWebKeyRsaRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaRequest) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaRequest) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AgentJsonWebKeyRsaRequest) SetE(v string) {
	o.E = &v
}

// GetKty returns the Kty field value
func (o *AgentJsonWebKeyRsaRequest) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *AgentJsonWebKeyRsaRequest) SetKty(v string) {
	o.Kty = v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaRequest) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaRequest) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *AgentJsonWebKeyRsaRequest) SetN(v string) {
	o.N = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaRequest) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaRequest) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AgentJsonWebKeyRsaRequest) SetKid(v string) {
	o.Kid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentJsonWebKeyRsaRequest) SetStatus(v string) {
	o.Status = &v
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaRequest) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaRequest) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *AgentJsonWebKeyRsaRequest) SetAlg(v string) {
	o.Alg = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaRequest) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaRequest) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaRequest) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AgentJsonWebKeyRsaRequest) SetUse(v string) {
	o.Use = &v
}

func (o AgentJsonWebKeyRsaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonWebKeyRsaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	toSerialize["kty"] = o.Kty
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentJsonWebKeyRsaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kty",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAgentJsonWebKeyRsaRequest := _AgentJsonWebKeyRsaRequest{}

	err = json.Unmarshal(data, &varAgentJsonWebKeyRsaRequest)

	if err != nil {
		return err
	}

	*o = AgentJsonWebKeyRsaRequest(varAgentJsonWebKeyRsaRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "alg")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentJsonWebKeyRsaRequest struct {
	value *AgentJsonWebKeyRsaRequest
	isSet bool
}

func (v NullableAgentJsonWebKeyRsaRequest) Get() *AgentJsonWebKeyRsaRequest {
	return v.value
}

func (v *NullableAgentJsonWebKeyRsaRequest) Set(val *AgentJsonWebKeyRsaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonWebKeyRsaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonWebKeyRsaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonWebKeyRsaRequest(val *AgentJsonWebKeyRsaRequest) *NullableAgentJsonWebKeyRsaRequest {
	return &NullableAgentJsonWebKeyRsaRequest{value: val, isSet: true}
}

func (v NullableAgentJsonWebKeyRsaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonWebKeyRsaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
