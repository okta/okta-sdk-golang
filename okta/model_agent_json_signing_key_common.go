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
)

// checks if the AgentJsonSigningKeyCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonSigningKeyCommon{}

// AgentJsonSigningKeyCommon struct for AgentJsonSigningKeyCommon
type AgentJsonSigningKeyCommon struct {
	// Algorithm that's used in the JSON Web Key
	Alg *string `json:"alg,omitempty"`
	// Acceptable use of the JSON Web Key  You can only use signing keys for AI agents, so the value of `use` is always `sig`.
	Use                  *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentJsonSigningKeyCommon AgentJsonSigningKeyCommon

// NewAgentJsonSigningKeyCommon instantiates a new AgentJsonSigningKeyCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonSigningKeyCommon() *AgentJsonSigningKeyCommon {
	this := AgentJsonSigningKeyCommon{}
	return &this
}

// NewAgentJsonSigningKeyCommonWithDefaults instantiates a new AgentJsonSigningKeyCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonSigningKeyCommonWithDefaults() *AgentJsonSigningKeyCommon {
	this := AgentJsonSigningKeyCommon{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *AgentJsonSigningKeyCommon) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonSigningKeyCommon) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *AgentJsonSigningKeyCommon) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *AgentJsonSigningKeyCommon) SetAlg(v string) {
	o.Alg = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AgentJsonSigningKeyCommon) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonSigningKeyCommon) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AgentJsonSigningKeyCommon) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AgentJsonSigningKeyCommon) SetUse(v string) {
	o.Use = &v
}

func (o AgentJsonSigningKeyCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonSigningKeyCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *AgentJsonSigningKeyCommon) UnmarshalJSON(data []byte) (err error) {
	varAgentJsonSigningKeyCommon := _AgentJsonSigningKeyCommon{}

	err = json.Unmarshal(data, &varAgentJsonSigningKeyCommon)

	if err != nil {
		return err
	}

	*o = AgentJsonSigningKeyCommon(varAgentJsonSigningKeyCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentJsonSigningKeyCommon struct {
	value *AgentJsonSigningKeyCommon
	isSet bool
}

func (v NullableAgentJsonSigningKeyCommon) Get() *AgentJsonSigningKeyCommon {
	return v.value
}

func (v *NullableAgentJsonSigningKeyCommon) Set(val *AgentJsonSigningKeyCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonSigningKeyCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonSigningKeyCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonSigningKeyCommon(val *AgentJsonSigningKeyCommon) *NullableAgentJsonSigningKeyCommon {
	return &NullableAgentJsonSigningKeyCommon{value: val, isSet: true}
}

func (v NullableAgentJsonSigningKeyCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonSigningKeyCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
