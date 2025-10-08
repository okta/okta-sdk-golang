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

// checks if the IdpDiscoveryPolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpDiscoveryPolicyRuleCondition{}

// IdpDiscoveryPolicyRuleCondition Specifies conditions that must be met during policy evaluation to apply the rule. All policy conditions and conditions for at least one rule must be met to apply the settings specified in the policy and the associated rule.
type IdpDiscoveryPolicyRuleCondition struct {
	App                  *AppAndInstancePolicyRuleCondition `json:"app,omitempty"`
	Network              *PolicyNetworkCondition            `json:"network,omitempty"`
	Platform             *PlatformPolicyRuleCondition       `json:"platform,omitempty"`
	UserIdentifier       *UserIdentifierPolicyRuleCondition `json:"userIdentifier,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpDiscoveryPolicyRuleCondition IdpDiscoveryPolicyRuleCondition

// NewIdpDiscoveryPolicyRuleCondition instantiates a new IdpDiscoveryPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpDiscoveryPolicyRuleCondition() *IdpDiscoveryPolicyRuleCondition {
	this := IdpDiscoveryPolicyRuleCondition{}
	return &this
}

// NewIdpDiscoveryPolicyRuleConditionWithDefaults instantiates a new IdpDiscoveryPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpDiscoveryPolicyRuleConditionWithDefaults() *IdpDiscoveryPolicyRuleCondition {
	this := IdpDiscoveryPolicyRuleCondition{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRuleCondition) GetApp() AppAndInstancePolicyRuleCondition {
	if o == nil || IsNil(o.App) {
		var ret AppAndInstancePolicyRuleCondition
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAndInstancePolicyRuleCondition and assigns it to the App field.
func (o *IdpDiscoveryPolicyRuleCondition) SetApp(v AppAndInstancePolicyRuleCondition) {
	o.App = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRuleCondition) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *IdpDiscoveryPolicyRuleCondition) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRuleCondition) GetPlatform() PlatformPolicyRuleCondition {
	if o == nil || IsNil(o.Platform) {
		var ret PlatformPolicyRuleCondition
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetPlatformOk() (*PlatformPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformPolicyRuleCondition and assigns it to the Platform field.
func (o *IdpDiscoveryPolicyRuleCondition) SetPlatform(v PlatformPolicyRuleCondition) {
	o.Platform = &v
}

// GetUserIdentifier returns the UserIdentifier field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRuleCondition) GetUserIdentifier() UserIdentifierPolicyRuleCondition {
	if o == nil || IsNil(o.UserIdentifier) {
		var ret UserIdentifierPolicyRuleCondition
		return ret
	}
	return *o.UserIdentifier
}

// GetUserIdentifierOk returns a tuple with the UserIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.UserIdentifier) {
		return nil, false
	}
	return o.UserIdentifier, true
}

// HasUserIdentifier returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasUserIdentifier() bool {
	if o != nil && !IsNil(o.UserIdentifier) {
		return true
	}

	return false
}

// SetUserIdentifier gets a reference to the given UserIdentifierPolicyRuleCondition and assigns it to the UserIdentifier field.
func (o *IdpDiscoveryPolicyRuleCondition) SetUserIdentifier(v UserIdentifierPolicyRuleCondition) {
	o.UserIdentifier = &v
}

func (o IdpDiscoveryPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpDiscoveryPolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.UserIdentifier) {
		toSerialize["userIdentifier"] = o.UserIdentifier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdpDiscoveryPolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varIdpDiscoveryPolicyRuleCondition := _IdpDiscoveryPolicyRuleCondition{}

	err = json.Unmarshal(data, &varIdpDiscoveryPolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = IdpDiscoveryPolicyRuleCondition(varIdpDiscoveryPolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "network")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "userIdentifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdpDiscoveryPolicyRuleCondition struct {
	value *IdpDiscoveryPolicyRuleCondition
	isSet bool
}

func (v NullableIdpDiscoveryPolicyRuleCondition) Get() *IdpDiscoveryPolicyRuleCondition {
	return v.value
}

func (v *NullableIdpDiscoveryPolicyRuleCondition) Set(val *IdpDiscoveryPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpDiscoveryPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpDiscoveryPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpDiscoveryPolicyRuleCondition(val *IdpDiscoveryPolicyRuleCondition) *NullableIdpDiscoveryPolicyRuleCondition {
	return &NullableIdpDiscoveryPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableIdpDiscoveryPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpDiscoveryPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
