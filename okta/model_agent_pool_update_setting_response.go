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
)

// checks if the AgentPoolUpdateSettingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPoolUpdateSettingResponse{}

// AgentPoolUpdateSettingResponse Settings for auto-update
type AgentPoolUpdateSettingResponse struct {
	// Agent types that are being monitored
	AgentType *string `json:"agentType,omitempty"`
	// Continues the update even if some agents fail to update
	ContinueOnError *bool `json:"continueOnError,omitempty"`
	// Latest version of the agent
	LatestVersion *string `json:"latestVersion,omitempty"`
	// Minimal version of the agent
	MinimalSupportedVersion *string `json:"minimalSupportedVersion,omitempty"`
	// ID of the agent pool that the settings apply to
	PoolId *string `json:"poolId,omitempty"`
	// Pool name
	PoolName *string `json:"poolName,omitempty"`
	// Release channel for auto-update
	ReleaseChannel       *string `json:"releaseChannel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentPoolUpdateSettingResponse AgentPoolUpdateSettingResponse

// NewAgentPoolUpdateSettingResponse instantiates a new AgentPoolUpdateSettingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPoolUpdateSettingResponse() *AgentPoolUpdateSettingResponse {
	this := AgentPoolUpdateSettingResponse{}
	return &this
}

// NewAgentPoolUpdateSettingResponseWithDefaults instantiates a new AgentPoolUpdateSettingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPoolUpdateSettingResponseWithDefaults() *AgentPoolUpdateSettingResponse {
	this := AgentPoolUpdateSettingResponse{}
	return &this
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetAgentType() string {
	if o == nil || IsNil(o.AgentType) {
		var ret string
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetAgentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AgentType) {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasAgentType() bool {
	if o != nil && !IsNil(o.AgentType) {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given string and assigns it to the AgentType field.
func (o *AgentPoolUpdateSettingResponse) SetAgentType(v string) {
	o.AgentType = &v
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.ContinueOnError) {
		return nil, false
	}
	return o.ContinueOnError, true
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasContinueOnError() bool {
	if o != nil && !IsNil(o.ContinueOnError) {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given bool and assigns it to the ContinueOnError field.
func (o *AgentPoolUpdateSettingResponse) SetContinueOnError(v bool) {
	o.ContinueOnError = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetLatestVersion() string {
	if o == nil || IsNil(o.LatestVersion) {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetLatestVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *AgentPoolUpdateSettingResponse) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetMinimalSupportedVersion returns the MinimalSupportedVersion field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetMinimalSupportedVersion() string {
	if o == nil || IsNil(o.MinimalSupportedVersion) {
		var ret string
		return ret
	}
	return *o.MinimalSupportedVersion
}

// GetMinimalSupportedVersionOk returns a tuple with the MinimalSupportedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetMinimalSupportedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimalSupportedVersion) {
		return nil, false
	}
	return o.MinimalSupportedVersion, true
}

// HasMinimalSupportedVersion returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasMinimalSupportedVersion() bool {
	if o != nil && !IsNil(o.MinimalSupportedVersion) {
		return true
	}

	return false
}

// SetMinimalSupportedVersion gets a reference to the given string and assigns it to the MinimalSupportedVersion field.
func (o *AgentPoolUpdateSettingResponse) SetMinimalSupportedVersion(v string) {
	o.MinimalSupportedVersion = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetPoolId() string {
	if o == nil || IsNil(o.PoolId) {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.PoolId) {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasPoolId() bool {
	if o != nil && !IsNil(o.PoolId) {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *AgentPoolUpdateSettingResponse) SetPoolId(v string) {
	o.PoolId = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetPoolName() string {
	if o == nil || IsNil(o.PoolName) {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetPoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.PoolName) {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasPoolName() bool {
	if o != nil && !IsNil(o.PoolName) {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *AgentPoolUpdateSettingResponse) SetPoolName(v string) {
	o.PoolName = &v
}

// GetReleaseChannel returns the ReleaseChannel field value if set, zero value otherwise.
func (o *AgentPoolUpdateSettingResponse) GetReleaseChannel() string {
	if o == nil || IsNil(o.ReleaseChannel) {
		var ret string
		return ret
	}
	return *o.ReleaseChannel
}

// GetReleaseChannelOk returns a tuple with the ReleaseChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSettingResponse) GetReleaseChannelOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseChannel) {
		return nil, false
	}
	return o.ReleaseChannel, true
}

// HasReleaseChannel returns a boolean if a field has been set.
func (o *AgentPoolUpdateSettingResponse) HasReleaseChannel() bool {
	if o != nil && !IsNil(o.ReleaseChannel) {
		return true
	}

	return false
}

// SetReleaseChannel gets a reference to the given string and assigns it to the ReleaseChannel field.
func (o *AgentPoolUpdateSettingResponse) SetReleaseChannel(v string) {
	o.ReleaseChannel = &v
}

func (o AgentPoolUpdateSettingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPoolUpdateSettingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentType) {
		toSerialize["agentType"] = o.AgentType
	}
	if !IsNil(o.ContinueOnError) {
		toSerialize["continueOnError"] = o.ContinueOnError
	}
	if !IsNil(o.LatestVersion) {
		toSerialize["latestVersion"] = o.LatestVersion
	}
	if !IsNil(o.MinimalSupportedVersion) {
		toSerialize["minimalSupportedVersion"] = o.MinimalSupportedVersion
	}
	if !IsNil(o.PoolId) {
		toSerialize["poolId"] = o.PoolId
	}
	if !IsNil(o.PoolName) {
		toSerialize["poolName"] = o.PoolName
	}
	if !IsNil(o.ReleaseChannel) {
		toSerialize["releaseChannel"] = o.ReleaseChannel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentPoolUpdateSettingResponse) UnmarshalJSON(data []byte) (err error) {
	varAgentPoolUpdateSettingResponse := _AgentPoolUpdateSettingResponse{}

	err = json.Unmarshal(data, &varAgentPoolUpdateSettingResponse)

	if err != nil {
		return err
	}

	*o = AgentPoolUpdateSettingResponse(varAgentPoolUpdateSettingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agentType")
		delete(additionalProperties, "continueOnError")
		delete(additionalProperties, "latestVersion")
		delete(additionalProperties, "minimalSupportedVersion")
		delete(additionalProperties, "poolId")
		delete(additionalProperties, "poolName")
		delete(additionalProperties, "releaseChannel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentPoolUpdateSettingResponse struct {
	value *AgentPoolUpdateSettingResponse
	isSet bool
}

func (v NullableAgentPoolUpdateSettingResponse) Get() *AgentPoolUpdateSettingResponse {
	return v.value
}

func (v *NullableAgentPoolUpdateSettingResponse) Set(val *AgentPoolUpdateSettingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPoolUpdateSettingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPoolUpdateSettingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPoolUpdateSettingResponse(val *AgentPoolUpdateSettingResponse) *NullableAgentPoolUpdateSettingResponse {
	return &NullableAgentPoolUpdateSettingResponse{value: val, isSet: true}
}

func (v NullableAgentPoolUpdateSettingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPoolUpdateSettingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
