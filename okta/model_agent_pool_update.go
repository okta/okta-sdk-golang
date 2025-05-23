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

// AgentPoolUpdate Various information about agent auto update configuration
type AgentPoolUpdate struct {
	Agents []Agent `json:"agents,omitempty"`
	// Agent types that are being monitored
	AgentType *string `json:"agentType,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NotifyAdmin *bool `json:"notifyAdmin,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Schedule *AutoUpdateSchedule `json:"schedule,omitempty"`
	SortOrder *int32 `json:"sortOrder,omitempty"`
	// Overall state for the auto-update job from admin perspective
	Status *string `json:"status,omitempty"`
	TargetVersion *string `json:"targetVersion,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentPoolUpdate AgentPoolUpdate

// NewAgentPoolUpdate instantiates a new AgentPoolUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPoolUpdate() *AgentPoolUpdate {
	this := AgentPoolUpdate{}
	return &this
}

// NewAgentPoolUpdateWithDefaults instantiates a new AgentPoolUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPoolUpdateWithDefaults() *AgentPoolUpdate {
	this := AgentPoolUpdate{}
	return &this
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetAgents() []Agent {
	if o == nil || o.Agents == nil {
		var ret []Agent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetAgentsOk() ([]Agent, bool) {
	if o == nil || o.Agents == nil {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasAgents() bool {
	if o != nil && o.Agents != nil {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []Agent and assigns it to the Agents field.
func (o *AgentPoolUpdate) SetAgents(v []Agent) {
	o.Agents = v
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetAgentType() string {
	if o == nil || o.AgentType == nil {
		var ret string
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetAgentTypeOk() (*string, bool) {
	if o == nil || o.AgentType == nil {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasAgentType() bool {
	if o != nil && o.AgentType != nil {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given string and assigns it to the AgentType field.
func (o *AgentPoolUpdate) SetAgentType(v string) {
	o.AgentType = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AgentPoolUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentPoolUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentPoolUpdate) SetName(v string) {
	o.Name = &v
}

// GetNotifyAdmin returns the NotifyAdmin field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetNotifyAdmin() bool {
	if o == nil || o.NotifyAdmin == nil {
		var ret bool
		return ret
	}
	return *o.NotifyAdmin
}

// GetNotifyAdminOk returns a tuple with the NotifyAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetNotifyAdminOk() (*bool, bool) {
	if o == nil || o.NotifyAdmin == nil {
		return nil, false
	}
	return o.NotifyAdmin, true
}

// HasNotifyAdmin returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasNotifyAdmin() bool {
	if o != nil && o.NotifyAdmin != nil {
		return true
	}

	return false
}

// SetNotifyAdmin gets a reference to the given bool and assigns it to the NotifyAdmin field.
func (o *AgentPoolUpdate) SetNotifyAdmin(v bool) {
	o.NotifyAdmin = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AgentPoolUpdate) SetReason(v string) {
	o.Reason = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetSchedule() AutoUpdateSchedule {
	if o == nil || o.Schedule == nil {
		var ret AutoUpdateSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetScheduleOk() (*AutoUpdateSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given AutoUpdateSchedule and assigns it to the Schedule field.
func (o *AgentPoolUpdate) SetSchedule(v AutoUpdateSchedule) {
	o.Schedule = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetSortOrder() int32 {
	if o == nil || o.SortOrder == nil {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetSortOrderOk() (*int32, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AgentPoolUpdate) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentPoolUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetTargetVersion() string {
	if o == nil || o.TargetVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetTargetVersionOk() (*string, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasTargetVersion() bool {
	if o != nil && o.TargetVersion != nil {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *AgentPoolUpdate) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentPoolUpdate) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdate) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentPoolUpdate) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *AgentPoolUpdate) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o AgentPoolUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Agents != nil {
		toSerialize["agents"] = o.Agents
	}
	if o.AgentType != nil {
		toSerialize["agentType"] = o.AgentType
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NotifyAdmin != nil {
		toSerialize["notifyAdmin"] = o.NotifyAdmin
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TargetVersion != nil {
		toSerialize["targetVersion"] = o.TargetVersion
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AgentPoolUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varAgentPoolUpdate := _AgentPoolUpdate{}

	err = json.Unmarshal(bytes, &varAgentPoolUpdate)
	if err == nil {
		*o = AgentPoolUpdate(varAgentPoolUpdate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "agents")
		delete(additionalProperties, "agentType")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notifyAdmin")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "sortOrder")
		delete(additionalProperties, "status")
		delete(additionalProperties, "targetVersion")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAgentPoolUpdate struct {
	value *AgentPoolUpdate
	isSet bool
}

func (v NullableAgentPoolUpdate) Get() *AgentPoolUpdate {
	return v.value
}

func (v *NullableAgentPoolUpdate) Set(val *AgentPoolUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPoolUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPoolUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPoolUpdate(val *AgentPoolUpdate) *NullableAgentPoolUpdate {
	return &NullableAgentPoolUpdate{value: val, isSet: true}
}

func (v NullableAgentPoolUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPoolUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

