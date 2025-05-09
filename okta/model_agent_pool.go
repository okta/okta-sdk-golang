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

// AgentPool An AgentPool is a collection of agents that serve a common purpose. An AgentPool has a unique ID within an org, and contains a collection of agents disjoint to every other AgentPool (i.e. no two AgentPools share an Agent).
type AgentPool struct {
	Agents []Agent `json:"agents,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Operational status of a given agent
	OperationalStatus *string `json:"operationalStatus,omitempty"`
	// Agent types that are being monitored
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentPool AgentPool

// NewAgentPool instantiates a new AgentPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPool() *AgentPool {
	this := AgentPool{}
	return &this
}

// NewAgentPoolWithDefaults instantiates a new AgentPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPoolWithDefaults() *AgentPool {
	this := AgentPool{}
	return &this
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *AgentPool) GetAgents() []Agent {
	if o == nil || o.Agents == nil {
		var ret []Agent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetAgentsOk() ([]Agent, bool) {
	if o == nil || o.Agents == nil {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AgentPool) HasAgents() bool {
	if o != nil && o.Agents != nil {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []Agent and assigns it to the Agents field.
func (o *AgentPool) SetAgents(v []Agent) {
	o.Agents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentPool) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentPool) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentPool) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentPool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentPool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentPool) SetName(v string) {
	o.Name = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *AgentPool) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *AgentPool) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *AgentPool) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AgentPool) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AgentPool) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AgentPool) SetType(v string) {
	o.Type = &v
}

func (o AgentPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Agents != nil {
		toSerialize["agents"] = o.Agents
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OperationalStatus != nil {
		toSerialize["operationalStatus"] = o.OperationalStatus
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AgentPool) UnmarshalJSON(bytes []byte) (err error) {
	varAgentPool := _AgentPool{}

	err = json.Unmarshal(bytes, &varAgentPool)
	if err == nil {
		*o = AgentPool(varAgentPool)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "agents")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "operationalStatus")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAgentPool struct {
	value *AgentPool
	isSet bool
}

func (v NullableAgentPool) Get() *AgentPool {
	return v.value
}

func (v *NullableAgentPool) Set(val *AgentPool) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPool) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPool(val *AgentPool) *NullableAgentPool {
	return &NullableAgentPool{value: val, isSet: true}
}

func (v NullableAgentPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

