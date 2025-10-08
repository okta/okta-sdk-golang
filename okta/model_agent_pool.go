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
)

// checks if the AgentPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPool{}

// AgentPool An agent pool is a collection of agents that serve a common purpose. An agent pool has a unique ID within an org, and contains a collection of agents disjoint to every other agent pool, meaning that no two agent pools share an agent.
type AgentPool struct {
	Agents []Agent `json:"agents,omitempty"`
	// Number of agents in the pool that are in a disrupted state
	DisruptedAgents *int32 `json:"disruptedAgents,omitempty"`
	// Agent pool ID
	Id *string `json:"id,omitempty"`
	// Number of agents in the pool that are in an inactive state
	InactiveAgents *int32 `json:"inactiveAgents,omitempty"`
	// Agent pool name
	Name *string `json:"name,omitempty"`
	// Operational status of a given agent
	OperationalStatus *string `json:"operationalStatus,omitempty"`
	// Agent types that are being monitored
	Type                 *string    `json:"type,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
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
	if o == nil || IsNil(o.Agents) {
		var ret []Agent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetAgentsOk() ([]Agent, bool) {
	if o == nil || IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AgentPool) HasAgents() bool {
	if o != nil && !IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []Agent and assigns it to the Agents field.
func (o *AgentPool) SetAgents(v []Agent) {
	o.Agents = v
}

// GetDisruptedAgents returns the DisruptedAgents field value if set, zero value otherwise.
func (o *AgentPool) GetDisruptedAgents() int32 {
	if o == nil || IsNil(o.DisruptedAgents) {
		var ret int32
		return ret
	}
	return *o.DisruptedAgents
}

// GetDisruptedAgentsOk returns a tuple with the DisruptedAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetDisruptedAgentsOk() (*int32, bool) {
	if o == nil || IsNil(o.DisruptedAgents) {
		return nil, false
	}
	return o.DisruptedAgents, true
}

// HasDisruptedAgents returns a boolean if a field has been set.
func (o *AgentPool) HasDisruptedAgents() bool {
	if o != nil && !IsNil(o.DisruptedAgents) {
		return true
	}

	return false
}

// SetDisruptedAgents gets a reference to the given int32 and assigns it to the DisruptedAgents field.
func (o *AgentPool) SetDisruptedAgents(v int32) {
	o.DisruptedAgents = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentPool) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentPool) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentPool) SetId(v string) {
	o.Id = &v
}

// GetInactiveAgents returns the InactiveAgents field value if set, zero value otherwise.
func (o *AgentPool) GetInactiveAgents() int32 {
	if o == nil || IsNil(o.InactiveAgents) {
		var ret int32
		return ret
	}
	return *o.InactiveAgents
}

// GetInactiveAgentsOk returns a tuple with the InactiveAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetInactiveAgentsOk() (*int32, bool) {
	if o == nil || IsNil(o.InactiveAgents) {
		return nil, false
	}
	return o.InactiveAgents, true
}

// HasInactiveAgents returns a boolean if a field has been set.
func (o *AgentPool) HasInactiveAgents() bool {
	if o != nil && !IsNil(o.InactiveAgents) {
		return true
	}

	return false
}

// SetInactiveAgents gets a reference to the given int32 and assigns it to the InactiveAgents field.
func (o *AgentPool) SetInactiveAgents(v int32) {
	o.InactiveAgents = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentPool) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentPool) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.OperationalStatus) {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetOperationalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OperationalStatus) {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *AgentPool) HasOperationalStatus() bool {
	if o != nil && !IsNil(o.OperationalStatus) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AgentPool) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AgentPool) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentPool) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPool) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentPool) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *AgentPool) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o AgentPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !IsNil(o.DisruptedAgents) {
		toSerialize["disruptedAgents"] = o.DisruptedAgents
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InactiveAgents) {
		toSerialize["inactiveAgents"] = o.InactiveAgents
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OperationalStatus) {
		toSerialize["operationalStatus"] = o.OperationalStatus
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentPool) UnmarshalJSON(data []byte) (err error) {
	varAgentPool := _AgentPool{}

	err = json.Unmarshal(data, &varAgentPool)

	if err != nil {
		return err
	}

	*o = AgentPool(varAgentPool)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agents")
		delete(additionalProperties, "disruptedAgents")
		delete(additionalProperties, "id")
		delete(additionalProperties, "inactiveAgents")
		delete(additionalProperties, "name")
		delete(additionalProperties, "operationalStatus")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
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
