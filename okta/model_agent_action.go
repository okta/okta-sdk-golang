/*
Okta Admin Management API

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

// checks if the AgentAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentAction{}

// AgentAction Details about the Active Directory or LDAP group membership update
type AgentAction struct {
	// ID of the Active Directory or LDAP group to update
	Id                   string     `json:"id"`
	Parameters           Parameters `json:"parameters"`
	AdditionalProperties map[string]interface{}
}

type _AgentAction AgentAction

// NewAgentAction instantiates a new AgentAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAction(id string, parameters Parameters) *AgentAction {
	this := AgentAction{}
	this.Id = id
	this.Parameters = parameters
	return &this
}

// NewAgentActionWithDefaults instantiates a new AgentAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentActionWithDefaults() *AgentAction {
	this := AgentAction{}
	return &this
}

// GetId returns the Id field value
func (o *AgentAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentAction) SetId(v string) {
	o.Id = v
}

// GetParameters returns the Parameters field value
func (o *AgentAction) GetParameters() Parameters {
	if o == nil {
		var ret Parameters
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *AgentAction) GetParametersOk() (*Parameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *AgentAction) SetParameters(v Parameters) {
	o.Parameters = v
}

func (o AgentAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["parameters"] = o.Parameters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"parameters",
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

	varAgentAction := _AgentAction{}

	err = json.Unmarshal(data, &varAgentAction)

	if err != nil {
		return err
	}

	*o = AgentAction(varAgentAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
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
