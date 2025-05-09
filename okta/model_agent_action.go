/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// AgentAction Details about the AD Group membership update
type AgentAction struct {
	// ID of the AD group to update
	Id *string `json:"id,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentAction AgentAction

// NewAgentAction instantiates a new AgentAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAction() *AgentAction {
	this := AgentAction{}
	return &this
}

// NewAgentActionWithDefaults instantiates a new AgentAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentActionWithDefaults() *AgentAction {
	this := AgentAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentAction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentAction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentAction) SetId(v string) {
	o.Id = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AgentAction) GetParameters() Parameters {
	if o == nil || o.Parameters == nil {
		var ret Parameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAction) GetParametersOk() (*Parameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AgentAction) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given Parameters and assigns it to the Parameters field.
func (o *AgentAction) SetParameters(v Parameters) {
	o.Parameters = &v
}

func (o AgentAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AgentAction) UnmarshalJSON(bytes []byte) (err error) {
	varAgentAction := _AgentAction{}

	err = json.Unmarshal(bytes, &varAgentAction)
	if err == nil {
		*o = AgentAction(varAgentAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAgentAction struct {
	value *AgentAction
	isSet bool
}

func (v NullableAgentAction) Get() *AgentAction {
	return v.value
}

func (v *NullableAgentAction) Set(val *AgentAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentAction(val *AgentAction) *NullableAgentAction {
	return &NullableAgentAction{value: val, isSet: true}
}

func (v NullableAgentAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

