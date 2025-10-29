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

// checks if the AgentJsonWebKeyECRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonWebKeyECRequest{}

// AgentJsonWebKeyECRequest An EC signing key
type AgentJsonWebKeyECRequest struct {
	// The cryptographic curve that's used for the key pair
	Crv *string `json:"crv,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty string `json:"kty"`
	// The public x coordinate for the elliptic curve point
	X *string `json:"x,omitempty"`
	// The public y coordinate for the elliptic curve point
	Y *string `json:"y,omitempty"`
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

type _AgentJsonWebKeyECRequest AgentJsonWebKeyECRequest

// NewAgentJsonWebKeyECRequest instantiates a new AgentJsonWebKeyECRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonWebKeyECRequest(kty string) *AgentJsonWebKeyECRequest {
	this := AgentJsonWebKeyECRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewAgentJsonWebKeyECRequestWithDefaults instantiates a new AgentJsonWebKeyECRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonWebKeyECRequestWithDefaults() *AgentJsonWebKeyECRequest {
	this := AgentJsonWebKeyECRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetCrv() string {
	if o == nil || IsNil(o.Crv) {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetCrvOk() (*string, bool) {
	if o == nil || IsNil(o.Crv) {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasCrv() bool {
	if o != nil && !IsNil(o.Crv) {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *AgentJsonWebKeyECRequest) SetCrv(v string) {
	o.Crv = &v
}

// GetKty returns the Kty field value
func (o *AgentJsonWebKeyECRequest) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *AgentJsonWebKeyECRequest) SetKty(v string) {
	o.Kty = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetX() string {
	if o == nil || IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetXOk() (*string, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *AgentJsonWebKeyECRequest) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetY() string {
	if o == nil || IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetYOk() (*string, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *AgentJsonWebKeyECRequest) SetY(v string) {
	o.Y = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AgentJsonWebKeyECRequest) SetKid(v string) {
	o.Kid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentJsonWebKeyECRequest) SetStatus(v string) {
	o.Status = &v
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *AgentJsonWebKeyECRequest) SetAlg(v string) {
	o.Alg = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECRequest) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECRequest) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECRequest) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AgentJsonWebKeyECRequest) SetUse(v string) {
	o.Use = &v
}

func (o AgentJsonWebKeyECRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonWebKeyECRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Crv) {
		toSerialize["crv"] = o.Crv
	}
	toSerialize["kty"] = o.Kty
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
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

func (o *AgentJsonWebKeyECRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAgentJsonWebKeyECRequest := _AgentJsonWebKeyECRequest{}

	err = json.Unmarshal(data, &varAgentJsonWebKeyECRequest)

	if err != nil {
		return err
	}

	*o = AgentJsonWebKeyECRequest(varAgentJsonWebKeyECRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "crv")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "alg")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentJsonWebKeyECRequest struct {
	value *AgentJsonWebKeyECRequest
	isSet bool
}

func (v NullableAgentJsonWebKeyECRequest) Get() *AgentJsonWebKeyECRequest {
	return v.value
}

func (v *NullableAgentJsonWebKeyECRequest) Set(val *AgentJsonWebKeyECRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonWebKeyECRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonWebKeyECRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonWebKeyECRequest(val *AgentJsonWebKeyECRequest) *NullableAgentJsonWebKeyECRequest {
	return &NullableAgentJsonWebKeyECRequest{value: val, isSet: true}
}

func (v NullableAgentJsonWebKeyECRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonWebKeyECRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
