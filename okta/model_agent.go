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
	"time"
)

// Agent Agent details
type Agent struct {
	Id *string `json:"id,omitempty"`
	IsHidden *bool `json:"isHidden,omitempty"`
	IsLatestGAedVersion *bool `json:"isLatestGAedVersion,omitempty"`
	LastConnection *time.Time `json:"lastConnection,omitempty"`
	Name *string `json:"name,omitempty"`
	// Operational status of a given agent
	OperationalStatus *string `json:"operationalStatus,omitempty"`
	PoolId *string `json:"poolId,omitempty"`
	// Agent types that are being monitored
	Type *string `json:"type,omitempty"`
	UpdateMessage *string `json:"updateMessage,omitempty"`
	// Status for one agent regarding the status to auto-update that agent
	UpdateStatus *string `json:"updateStatus,omitempty"`
	Version *string `json:"version,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Agent Agent

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent() *Agent {
	this := Agent{}
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Agent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Agent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Agent) SetId(v string) {
	o.Id = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *Agent) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *Agent) HasIsHidden() bool {
	if o != nil && o.IsHidden != nil {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *Agent) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetIsLatestGAedVersion returns the IsLatestGAedVersion field value if set, zero value otherwise.
func (o *Agent) GetIsLatestGAedVersion() bool {
	if o == nil || o.IsLatestGAedVersion == nil {
		var ret bool
		return ret
	}
	return *o.IsLatestGAedVersion
}

// GetIsLatestGAedVersionOk returns a tuple with the IsLatestGAedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetIsLatestGAedVersionOk() (*bool, bool) {
	if o == nil || o.IsLatestGAedVersion == nil {
		return nil, false
	}
	return o.IsLatestGAedVersion, true
}

// HasIsLatestGAedVersion returns a boolean if a field has been set.
func (o *Agent) HasIsLatestGAedVersion() bool {
	if o != nil && o.IsLatestGAedVersion != nil {
		return true
	}

	return false
}

// SetIsLatestGAedVersion gets a reference to the given bool and assigns it to the IsLatestGAedVersion field.
func (o *Agent) SetIsLatestGAedVersion(v bool) {
	o.IsLatestGAedVersion = &v
}

// GetLastConnection returns the LastConnection field value if set, zero value otherwise.
func (o *Agent) GetLastConnection() time.Time {
	if o == nil || o.LastConnection == nil {
		var ret time.Time
		return ret
	}
	return *o.LastConnection
}

// GetLastConnectionOk returns a tuple with the LastConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetLastConnectionOk() (*time.Time, bool) {
	if o == nil || o.LastConnection == nil {
		return nil, false
	}
	return o.LastConnection, true
}

// HasLastConnection returns a boolean if a field has been set.
func (o *Agent) HasLastConnection() bool {
	if o != nil && o.LastConnection != nil {
		return true
	}

	return false
}

// SetLastConnection gets a reference to the given time.Time and assigns it to the LastConnection field.
func (o *Agent) SetLastConnection(v time.Time) {
	o.LastConnection = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Agent) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Agent) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Agent) SetName(v string) {
	o.Name = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *Agent) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *Agent) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *Agent) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *Agent) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *Agent) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *Agent) SetPoolId(v string) {
	o.PoolId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Agent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Agent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Agent) SetType(v string) {
	o.Type = &v
}

// GetUpdateMessage returns the UpdateMessage field value if set, zero value otherwise.
func (o *Agent) GetUpdateMessage() string {
	if o == nil || o.UpdateMessage == nil {
		var ret string
		return ret
	}
	return *o.UpdateMessage
}

// GetUpdateMessageOk returns a tuple with the UpdateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUpdateMessageOk() (*string, bool) {
	if o == nil || o.UpdateMessage == nil {
		return nil, false
	}
	return o.UpdateMessage, true
}

// HasUpdateMessage returns a boolean if a field has been set.
func (o *Agent) HasUpdateMessage() bool {
	if o != nil && o.UpdateMessage != nil {
		return true
	}

	return false
}

// SetUpdateMessage gets a reference to the given string and assigns it to the UpdateMessage field.
func (o *Agent) SetUpdateMessage(v string) {
	o.UpdateMessage = &v
}

// GetUpdateStatus returns the UpdateStatus field value if set, zero value otherwise.
func (o *Agent) GetUpdateStatus() string {
	if o == nil || o.UpdateStatus == nil {
		var ret string
		return ret
	}
	return *o.UpdateStatus
}

// GetUpdateStatusOk returns a tuple with the UpdateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetUpdateStatusOk() (*string, bool) {
	if o == nil || o.UpdateStatus == nil {
		return nil, false
	}
	return o.UpdateStatus, true
}

// HasUpdateStatus returns a boolean if a field has been set.
func (o *Agent) HasUpdateStatus() bool {
	if o != nil && o.UpdateStatus != nil {
		return true
	}

	return false
}

// SetUpdateStatus gets a reference to the given string and assigns it to the UpdateStatus field.
func (o *Agent) SetUpdateStatus(v string) {
	o.UpdateStatus = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Agent) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Agent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Agent) SetVersion(v string) {
	o.Version = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Agent) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Agent) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *Agent) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsHidden != nil {
		toSerialize["isHidden"] = o.IsHidden
	}
	if o.IsLatestGAedVersion != nil {
		toSerialize["isLatestGAedVersion"] = o.IsLatestGAedVersion
	}
	if o.LastConnection != nil {
		toSerialize["lastConnection"] = o.LastConnection
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OperationalStatus != nil {
		toSerialize["operationalStatus"] = o.OperationalStatus
	}
	if o.PoolId != nil {
		toSerialize["poolId"] = o.PoolId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdateMessage != nil {
		toSerialize["updateMessage"] = o.UpdateMessage
	}
	if o.UpdateStatus != nil {
		toSerialize["updateStatus"] = o.UpdateStatus
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Agent) UnmarshalJSON(bytes []byte) (err error) {
	varAgent := _Agent{}

	err = json.Unmarshal(bytes, &varAgent)
	if err == nil {
		*o = Agent(varAgent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "isHidden")
		delete(additionalProperties, "isLatestGAedVersion")
		delete(additionalProperties, "lastConnection")
		delete(additionalProperties, "name")
		delete(additionalProperties, "operationalStatus")
		delete(additionalProperties, "poolId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updateMessage")
		delete(additionalProperties, "updateStatus")
		delete(additionalProperties, "version")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

