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

// checks if the UserIdentificationPolicyRuleConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationPolicyRuleConditions{}

// UserIdentificationPolicyRuleConditions <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies conditions that must be met for this rule to apply during policy evaluation. The settings specified in the associated rule are applied when the rule conditions are met.
type UserIdentificationPolicyRuleConditions struct {
	Network              *PolicyNetworkCondition      `json:"network,omitempty"`
	Platform             *PlatformPolicyRuleCondition `json:"platform,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationPolicyRuleConditions UserIdentificationPolicyRuleConditions

// NewUserIdentificationPolicyRuleConditions instantiates a new UserIdentificationPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationPolicyRuleConditions() *UserIdentificationPolicyRuleConditions {
	this := UserIdentificationPolicyRuleConditions{}
	return &this
}

// NewUserIdentificationPolicyRuleConditionsWithDefaults instantiates a new UserIdentificationPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationPolicyRuleConditionsWithDefaults() *UserIdentificationPolicyRuleConditions {
	this := UserIdentificationPolicyRuleConditions{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *UserIdentificationPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *UserIdentificationPolicyRuleConditions) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *UserIdentificationPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *UserIdentificationPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition {
	if o == nil || IsNil(o.Platform) {
		var ret PlatformPolicyRuleCondition
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *UserIdentificationPolicyRuleConditions) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformPolicyRuleCondition and assigns it to the Platform field.
func (o *UserIdentificationPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition) {
	o.Platform = &v
}

func (o UserIdentificationPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationPolicyRuleConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationPolicyRuleConditions) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationPolicyRuleConditions := _UserIdentificationPolicyRuleConditions{}

	err = json.Unmarshal(data, &varUserIdentificationPolicyRuleConditions)

	if err != nil {
		return err
	}

	*o = UserIdentificationPolicyRuleConditions(varUserIdentificationPolicyRuleConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "platform")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationPolicyRuleConditions struct {
	value *UserIdentificationPolicyRuleConditions
	isSet bool
}

func (v NullableUserIdentificationPolicyRuleConditions) Get() *UserIdentificationPolicyRuleConditions {
	return v.value
}

func (v *NullableUserIdentificationPolicyRuleConditions) Set(val *UserIdentificationPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationPolicyRuleConditions(val *UserIdentificationPolicyRuleConditions) *NullableUserIdentificationPolicyRuleConditions {
	return &NullableUserIdentificationPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableUserIdentificationPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
