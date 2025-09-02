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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// DeviceSignalCollectionPolicyRuleConditions <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Specifies conditions that must be met during policy evaluation to apply the rule. All policy conditions, as well as conditions for at least one rule must be met to apply the settings specified in the policy and the associated rule.
type DeviceSignalCollectionPolicyRuleConditions struct {
	Network *PolicyNetworkCondition `json:"network,omitempty"`
	Platform *PlatformPolicyRuleCondition `json:"platform,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceSignalCollectionPolicyRuleConditions DeviceSignalCollectionPolicyRuleConditions

// NewDeviceSignalCollectionPolicyRuleConditions instantiates a new DeviceSignalCollectionPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSignalCollectionPolicyRuleConditions() *DeviceSignalCollectionPolicyRuleConditions {
	this := DeviceSignalCollectionPolicyRuleConditions{}
	return &this
}

// NewDeviceSignalCollectionPolicyRuleConditionsWithDefaults instantiates a new DeviceSignalCollectionPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSignalCollectionPolicyRuleConditionsWithDefaults() *DeviceSignalCollectionPolicyRuleConditions {
	this := DeviceSignalCollectionPolicyRuleConditions{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || o.Network == nil {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPolicyRuleConditions) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *DeviceSignalCollectionPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition {
	if o == nil || o.Platform == nil {
		var ret PlatformPolicyRuleCondition
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPolicyRuleConditions) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformPolicyRuleCondition and assigns it to the Platform field.
func (o *DeviceSignalCollectionPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition) {
	o.Platform = &v
}

func (o DeviceSignalCollectionPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceSignalCollectionPolicyRuleConditions) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceSignalCollectionPolicyRuleConditions := _DeviceSignalCollectionPolicyRuleConditions{}

	err = json.Unmarshal(bytes, &varDeviceSignalCollectionPolicyRuleConditions)
	if err == nil {
		*o = DeviceSignalCollectionPolicyRuleConditions(varDeviceSignalCollectionPolicyRuleConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "platform")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceSignalCollectionPolicyRuleConditions struct {
	value *DeviceSignalCollectionPolicyRuleConditions
	isSet bool
}

func (v NullableDeviceSignalCollectionPolicyRuleConditions) Get() *DeviceSignalCollectionPolicyRuleConditions {
	return v.value
}

func (v *NullableDeviceSignalCollectionPolicyRuleConditions) Set(val *DeviceSignalCollectionPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSignalCollectionPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSignalCollectionPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSignalCollectionPolicyRuleConditions(val *DeviceSignalCollectionPolicyRuleConditions) *NullableDeviceSignalCollectionPolicyRuleConditions {
	return &NullableDeviceSignalCollectionPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableDeviceSignalCollectionPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSignalCollectionPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

