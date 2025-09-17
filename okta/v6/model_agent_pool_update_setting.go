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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AgentPoolUpdateSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPoolUpdateSetting{}

// AgentPoolUpdateSetting Setting for auto-update
type AgentPoolUpdateSetting struct {
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

type _AgentPoolUpdateSetting AgentPoolUpdateSetting

// NewAgentPoolUpdateSetting instantiates a new AgentPoolUpdateSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPoolUpdateSetting() *AgentPoolUpdateSetting {
	this := AgentPoolUpdateSetting{}
	return &this
}

// NewAgentPoolUpdateSettingWithDefaults instantiates a new AgentPoolUpdateSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPoolUpdateSettingWithDefaults() *AgentPoolUpdateSetting {
	this := AgentPoolUpdateSetting{}
	return &this
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetAgentType() string {
	if o == nil || IsNil(o.AgentType) {
		var ret string
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetAgentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AgentType) {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasAgentType() bool {
	if o != nil && !IsNil(o.AgentType) {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given string and assigns it to the AgentType field.
func (o *AgentPoolUpdateSetting) SetAgentType(v string) {
	o.AgentType = &v
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.ContinueOnError) {
		return nil, false
	}
	return o.ContinueOnError, true
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasContinueOnError() bool {
	if o != nil && !IsNil(o.ContinueOnError) {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given bool and assigns it to the ContinueOnError field.
func (o *AgentPoolUpdateSetting) SetContinueOnError(v bool) {
	o.ContinueOnError = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetLatestVersion() string {
	if o == nil || IsNil(o.LatestVersion) {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetLatestVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *AgentPoolUpdateSetting) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetMinimalSupportedVersion returns the MinimalSupportedVersion field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetMinimalSupportedVersion() string {
	if o == nil || IsNil(o.MinimalSupportedVersion) {
		var ret string
		return ret
	}
	return *o.MinimalSupportedVersion
}

// GetMinimalSupportedVersionOk returns a tuple with the MinimalSupportedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetMinimalSupportedVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimalSupportedVersion) {
		return nil, false
	}
	return o.MinimalSupportedVersion, true
}

// HasMinimalSupportedVersion returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasMinimalSupportedVersion() bool {
	if o != nil && !IsNil(o.MinimalSupportedVersion) {
		return true
	}

	return false
}

// SetMinimalSupportedVersion gets a reference to the given string and assigns it to the MinimalSupportedVersion field.
func (o *AgentPoolUpdateSetting) SetMinimalSupportedVersion(v string) {
	o.MinimalSupportedVersion = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetPoolId() string {
	if o == nil || IsNil(o.PoolId) {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.PoolId) {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasPoolId() bool {
	if o != nil && !IsNil(o.PoolId) {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *AgentPoolUpdateSetting) SetPoolId(v string) {
	o.PoolId = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetPoolName() string {
	if o == nil || IsNil(o.PoolName) {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetPoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.PoolName) {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasPoolName() bool {
	if o != nil && !IsNil(o.PoolName) {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *AgentPoolUpdateSetting) SetPoolName(v string) {
	o.PoolName = &v
}

// GetReleaseChannel returns the ReleaseChannel field value if set, zero value otherwise.
func (o *AgentPoolUpdateSetting) GetReleaseChannel() string {
	if o == nil || IsNil(o.ReleaseChannel) {
		var ret string
		return ret
	}
	return *o.ReleaseChannel
}

// GetReleaseChannelOk returns a tuple with the ReleaseChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPoolUpdateSetting) GetReleaseChannelOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseChannel) {
		return nil, false
	}
	return o.ReleaseChannel, true
}

// HasReleaseChannel returns a boolean if a field has been set.
func (o *AgentPoolUpdateSetting) HasReleaseChannel() bool {
	if o != nil && !IsNil(o.ReleaseChannel) {
		return true
	}

	return false
}

// SetReleaseChannel gets a reference to the given string and assigns it to the ReleaseChannel field.
func (o *AgentPoolUpdateSetting) SetReleaseChannel(v string) {
	o.ReleaseChannel = &v
}

func (o AgentPoolUpdateSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPoolUpdateSetting) ToMap() (map[string]interface{}, error) {
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

func (o *AgentPoolUpdateSetting) UnmarshalJSON(data []byte) (err error) {
	varAgentPoolUpdateSetting := _AgentPoolUpdateSetting{}

	err = json.Unmarshal(data, &varAgentPoolUpdateSetting)

	if err != nil {
		return err
	}

	*o = AgentPoolUpdateSetting(varAgentPoolUpdateSetting)

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

type NullableAgentPoolUpdateSetting struct {
	value *AgentPoolUpdateSetting
	isSet bool
}

func (v NullableAgentPoolUpdateSetting) Get() *AgentPoolUpdateSetting {
	return v.value
}

func (v *NullableAgentPoolUpdateSetting) Set(val *AgentPoolUpdateSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPoolUpdateSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPoolUpdateSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPoolUpdateSetting(val *AgentPoolUpdateSetting) *NullableAgentPoolUpdateSetting {
	return &NullableAgentPoolUpdateSetting{value: val, isSet: true}
}

func (v NullableAgentPoolUpdateSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPoolUpdateSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
