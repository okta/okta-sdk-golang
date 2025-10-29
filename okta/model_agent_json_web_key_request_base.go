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

// checks if the AgentJsonWebKeyRequestBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonWebKeyRequestBase{}

// AgentJsonWebKeyRequestBase struct for AgentJsonWebKeyRequestBase
type AgentJsonWebKeyRequestBase struct {
	// Unique identifier of the JSON Web Key in the AI agent's JSON Web Key Set (JWKS)
	Kid *string `json:"kid,omitempty"`
	// Status of the AI agent JSON Web Key
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentJsonWebKeyRequestBase AgentJsonWebKeyRequestBase

// NewAgentJsonWebKeyRequestBase instantiates a new AgentJsonWebKeyRequestBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonWebKeyRequestBase() *AgentJsonWebKeyRequestBase {
	this := AgentJsonWebKeyRequestBase{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewAgentJsonWebKeyRequestBaseWithDefaults instantiates a new AgentJsonWebKeyRequestBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonWebKeyRequestBaseWithDefaults() *AgentJsonWebKeyRequestBase {
	this := AgentJsonWebKeyRequestBase{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRequestBase) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRequestBase) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRequestBase) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AgentJsonWebKeyRequestBase) SetKid(v string) {
	o.Kid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRequestBase) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRequestBase) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRequestBase) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentJsonWebKeyRequestBase) SetStatus(v string) {
	o.Status = &v
}

func (o AgentJsonWebKeyRequestBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonWebKeyRequestBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentJsonWebKeyRequestBase) UnmarshalJSON(data []byte) (err error) {
	varAgentJsonWebKeyRequestBase := _AgentJsonWebKeyRequestBase{}

	err = json.Unmarshal(data, &varAgentJsonWebKeyRequestBase)

	if err != nil {
		return err
	}

	*o = AgentJsonWebKeyRequestBase(varAgentJsonWebKeyRequestBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentJsonWebKeyRequestBase struct {
	value *AgentJsonWebKeyRequestBase
	isSet bool
}

func (v NullableAgentJsonWebKeyRequestBase) Get() *AgentJsonWebKeyRequestBase {
	return v.value
}

func (v *NullableAgentJsonWebKeyRequestBase) Set(val *AgentJsonWebKeyRequestBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonWebKeyRequestBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonWebKeyRequestBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonWebKeyRequestBase(val *AgentJsonWebKeyRequestBase) *NullableAgentJsonWebKeyRequestBase {
	return &NullableAgentJsonWebKeyRequestBase{value: val, isSet: true}
}

func (v NullableAgentJsonWebKeyRequestBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonWebKeyRequestBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
