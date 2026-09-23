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

// checks if the ParallelScepConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParallelScepConfig{}

// ParallelScepConfig A new SCEP configuration to create against the renewed certificate authority, cloned from an existing one.  For a delegated challenge, Okta overlays `configInfo` on the source configuration's settings rather than replacing them. Omit it entirely to reuse the same Microsoft Entra ID application, including its client secret, which you don't have to re-send. Alternatively, include only the properties you want to change.  `intuneCloudType` is accepted here only when the per-configuration Intune cloud type is enabled for your org. When it isn't, the source configuration's value isn't carried over to the clone either, so the clone falls back to your org's default cloud.
type ParallelScepConfig struct {
	// The challenge type for a configuration created during a renewal. Only `DELEGATED` is supported for now — a static or dynamic clone would need a freshly generated challenge secret, which this operation can't return. Migrate those configurations with `rolloverConfigIds` instead.
	ChallengeType string                           `json:"challengeType"`
	ConfigInfo    *RegistrationAuthorityConfigInfo `json:"configInfo,omitempty"`
	// The name for the new configuration. Give it one that distinguishes it from the source, so the two are easy to tell apart while both are live.
	Name *string `json:"name,omitempty"`
	// The enrollment protocol the configuration serves
	Protocol string `json:"protocol"`
	// The ID of the existing configuration to clone. It stays bound to the certificate it's already on.
	SourceConfigId       string `json:"sourceConfigId"`
	AdditionalProperties map[string]interface{}
}

type _ParallelScepConfig ParallelScepConfig

// NewParallelScepConfig instantiates a new ParallelScepConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParallelScepConfig(challengeType string, protocol string, sourceConfigId string) *ParallelScepConfig {
	this := ParallelScepConfig{}
	this.ChallengeType = challengeType
	this.Protocol = protocol
	this.SourceConfigId = sourceConfigId
	return &this
}

// NewParallelScepConfigWithDefaults instantiates a new ParallelScepConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParallelScepConfigWithDefaults() *ParallelScepConfig {
	this := ParallelScepConfig{}
	return &this
}

// GetChallengeType returns the ChallengeType field value
func (o *ParallelScepConfig) GetChallengeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChallengeType
}

// GetChallengeTypeOk returns a tuple with the ChallengeType field value
// and a boolean to check if the value has been set.
func (o *ParallelScepConfig) GetChallengeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChallengeType, true
}

// SetChallengeType sets field value
func (o *ParallelScepConfig) SetChallengeType(v string) {
	o.ChallengeType = v
}

// GetConfigInfo returns the ConfigInfo field value if set, zero value otherwise.
func (o *ParallelScepConfig) GetConfigInfo() RegistrationAuthorityConfigInfo {
	if o == nil || IsNil(o.ConfigInfo) {
		var ret RegistrationAuthorityConfigInfo
		return ret
	}
	return *o.ConfigInfo
}

// GetConfigInfoOk returns a tuple with the ConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParallelScepConfig) GetConfigInfoOk() (*RegistrationAuthorityConfigInfo, bool) {
	if o == nil || IsNil(o.ConfigInfo) {
		return nil, false
	}
	return o.ConfigInfo, true
}

// HasConfigInfo returns a boolean if a field has been set.
func (o *ParallelScepConfig) HasConfigInfo() bool {
	if o != nil && !IsNil(o.ConfigInfo) {
		return true
	}

	return false
}

// SetConfigInfo gets a reference to the given RegistrationAuthorityConfigInfo and assigns it to the ConfigInfo field.
func (o *ParallelScepConfig) SetConfigInfo(v RegistrationAuthorityConfigInfo) {
	o.ConfigInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParallelScepConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParallelScepConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParallelScepConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParallelScepConfig) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value
func (o *ParallelScepConfig) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *ParallelScepConfig) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *ParallelScepConfig) SetProtocol(v string) {
	o.Protocol = v
}

// GetSourceConfigId returns the SourceConfigId field value
func (o *ParallelScepConfig) GetSourceConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceConfigId
}

// GetSourceConfigIdOk returns a tuple with the SourceConfigId field value
// and a boolean to check if the value has been set.
func (o *ParallelScepConfig) GetSourceConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceConfigId, true
}

// SetSourceConfigId sets field value
func (o *ParallelScepConfig) SetSourceConfigId(v string) {
	o.SourceConfigId = v
}

func (o ParallelScepConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParallelScepConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["challengeType"] = o.ChallengeType
	if !IsNil(o.ConfigInfo) {
		toSerialize["configInfo"] = o.ConfigInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["protocol"] = o.Protocol
	toSerialize["sourceConfigId"] = o.SourceConfigId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ParallelScepConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"challengeType",
		"protocol",
		"sourceConfigId",
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

	varParallelScepConfig := _ParallelScepConfig{}

	err = json.Unmarshal(data, &varParallelScepConfig)

	if err != nil {
		return err
	}

	*o = ParallelScepConfig(varParallelScepConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "challengeType")
		delete(additionalProperties, "configInfo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "sourceConfigId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableParallelScepConfig struct {
	value *ParallelScepConfig
	isSet bool
}

func (v NullableParallelScepConfig) Get() *ParallelScepConfig {
	return v.value
}

func (v *NullableParallelScepConfig) Set(val *ParallelScepConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableParallelScepConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableParallelScepConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParallelScepConfig(val *ParallelScepConfig) *NullableParallelScepConfig {
	return &NullableParallelScepConfig{value: val, isSet: true}
}

func (v NullableParallelScepConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParallelScepConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
