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

// checks if the SessionViolationDetectionPolicyRuleAllOfConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionViolationDetectionPolicyRuleAllOfConditions{}

// SessionViolationDetectionPolicyRuleAllOfConditions Specifies policy evaluation conditions required to apply the rule. All policy conditions and conditions for at least one rule must be met to apply the settings specified in the policy and the associated rule.
type SessionViolationDetectionPolicyRuleAllOfConditions struct {
	Network              *PolicyNetworkCondition       `json:"network,omitempty"`
	RiskScore            *RiskScorePolicyRuleCondition `json:"riskScore,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionViolationDetectionPolicyRuleAllOfConditions SessionViolationDetectionPolicyRuleAllOfConditions

// NewSessionViolationDetectionPolicyRuleAllOfConditions instantiates a new SessionViolationDetectionPolicyRuleAllOfConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionViolationDetectionPolicyRuleAllOfConditions() *SessionViolationDetectionPolicyRuleAllOfConditions {
	this := SessionViolationDetectionPolicyRuleAllOfConditions{}
	return &this
}

// NewSessionViolationDetectionPolicyRuleAllOfConditionsWithDefaults instantiates a new SessionViolationDetectionPolicyRuleAllOfConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionViolationDetectionPolicyRuleAllOfConditionsWithDefaults() *SessionViolationDetectionPolicyRuleAllOfConditions {
	this := SessionViolationDetectionPolicyRuleAllOfConditions{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetRiskScore() RiskScorePolicyRuleCondition {
	if o == nil || IsNil(o.RiskScore) {
		var ret RiskScorePolicyRuleCondition
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool) {
	if o == nil || IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) HasRiskScore() bool {
	if o != nil && !IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given RiskScorePolicyRuleCondition and assigns it to the RiskScore field.
func (o *SessionViolationDetectionPolicyRuleAllOfConditions) SetRiskScore(v RiskScorePolicyRuleCondition) {
	o.RiskScore = &v
}

func (o SessionViolationDetectionPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionViolationDetectionPolicyRuleAllOfConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionViolationDetectionPolicyRuleAllOfConditions) UnmarshalJSON(data []byte) (err error) {
	varSessionViolationDetectionPolicyRuleAllOfConditions := _SessionViolationDetectionPolicyRuleAllOfConditions{}

	err = json.Unmarshal(data, &varSessionViolationDetectionPolicyRuleAllOfConditions)

	if err != nil {
		return err
	}

	*o = SessionViolationDetectionPolicyRuleAllOfConditions(varSessionViolationDetectionPolicyRuleAllOfConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "riskScore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionViolationDetectionPolicyRuleAllOfConditions struct {
	value *SessionViolationDetectionPolicyRuleAllOfConditions
	isSet bool
}

func (v NullableSessionViolationDetectionPolicyRuleAllOfConditions) Get() *SessionViolationDetectionPolicyRuleAllOfConditions {
	return v.value
}

func (v *NullableSessionViolationDetectionPolicyRuleAllOfConditions) Set(val *SessionViolationDetectionPolicyRuleAllOfConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionViolationDetectionPolicyRuleAllOfConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionViolationDetectionPolicyRuleAllOfConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionViolationDetectionPolicyRuleAllOfConditions(val *SessionViolationDetectionPolicyRuleAllOfConditions) *NullableSessionViolationDetectionPolicyRuleAllOfConditions {
	return &NullableSessionViolationDetectionPolicyRuleAllOfConditions{value: val, isSet: true}
}

func (v NullableSessionViolationDetectionPolicyRuleAllOfConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionViolationDetectionPolicyRuleAllOfConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
