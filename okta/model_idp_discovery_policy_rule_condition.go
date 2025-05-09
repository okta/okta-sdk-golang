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

// IdpDiscoveryPolicyRuleCondition struct for IdpDiscoveryPolicyRuleCondition
type IdpDiscoveryPolicyRuleCondition struct {
	App *AppAndInstancePolicyRuleCondition `json:"app,omitempty"`
	Network *PolicyNetworkCondition `json:"network,omitempty"`
	UserIdentifier *UserIdentifierPolicyRuleCondition `json:"userIdentifier,omitempty"`
	Platform *PlatformPolicyRuleCondition `json:"platform,omitempty"`
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
	if o == nil || o.App == nil {
		var ret AppAndInstancePolicyRuleCondition
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasApp() bool {
	if o != nil && o.App != nil {
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
	if o == nil || o.Network == nil {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *IdpDiscoveryPolicyRuleCondition) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetUserIdentifier returns the UserIdentifier field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRuleCondition) GetUserIdentifier() UserIdentifierPolicyRuleCondition {
	if o == nil || o.UserIdentifier == nil {
		var ret UserIdentifierPolicyRuleCondition
		return ret
	}
	return *o.UserIdentifier
}

// GetUserIdentifierOk returns a tuple with the UserIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool) {
	if o == nil || o.UserIdentifier == nil {
		return nil, false
	}
	return o.UserIdentifier, true
}

// HasUserIdentifier returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasUserIdentifier() bool {
	if o != nil && o.UserIdentifier != nil {
		return true
	}

	return false
}

// SetUserIdentifier gets a reference to the given UserIdentifierPolicyRuleCondition and assigns it to the UserIdentifier field.
func (o *IdpDiscoveryPolicyRuleCondition) SetUserIdentifier(v UserIdentifierPolicyRuleCondition) {
	o.UserIdentifier = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *IdpDiscoveryPolicyRuleCondition) GetPlatform() PlatformPolicyRuleCondition {
	if o == nil || o.Platform == nil {
		var ret PlatformPolicyRuleCondition
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPolicyRuleCondition) GetPlatformOk() (*PlatformPolicyRuleCondition, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicyRuleCondition) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformPolicyRuleCondition and assigns it to the Platform field.
func (o *IdpDiscoveryPolicyRuleCondition) SetPlatform(v PlatformPolicyRuleCondition) {
	o.Platform = &v
}

func (o IdpDiscoveryPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.UserIdentifier != nil {
		toSerialize["userIdentifier"] = o.UserIdentifier
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdpDiscoveryPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varIdpDiscoveryPolicyRuleCondition := _IdpDiscoveryPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varIdpDiscoveryPolicyRuleCondition)
	if err == nil {
		*o = IdpDiscoveryPolicyRuleCondition(varIdpDiscoveryPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "network")
		delete(additionalProperties, "userIdentifier")
		delete(additionalProperties, "platform")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

