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

// checks if the SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection{}

// SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection Used to define the actions taken as a result of evaluating session violation detection policy
type SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection struct {
	PolicyEvaluation     *SessionViolationDetectionPolicyEvaluation `json:"policyEvaluation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection

// NewSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection instantiates a new SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection() *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection {
	this := SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection{}
	return &this
}

// NewSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetectionWithDefaults instantiates a new SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetectionWithDefaults() *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection {
	this := SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection{}
	return &this
}

// GetPolicyEvaluation returns the PolicyEvaluation field value if set, zero value otherwise.
func (o *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) GetPolicyEvaluation() SessionViolationDetectionPolicyEvaluation {
	if o == nil || IsNil(o.PolicyEvaluation) {
		var ret SessionViolationDetectionPolicyEvaluation
		return ret
	}
	return *o.PolicyEvaluation
}

// GetPolicyEvaluationOk returns a tuple with the PolicyEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) GetPolicyEvaluationOk() (*SessionViolationDetectionPolicyEvaluation, bool) {
	if o == nil || IsNil(o.PolicyEvaluation) {
		return nil, false
	}
	return o.PolicyEvaluation, true
}

// HasPolicyEvaluation returns a boolean if a field has been set.
func (o *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) HasPolicyEvaluation() bool {
	if o != nil && !IsNil(o.PolicyEvaluation) {
		return true
	}

	return false
}

// SetPolicyEvaluation gets a reference to the given SessionViolationDetectionPolicyEvaluation and assigns it to the PolicyEvaluation field.
func (o *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) SetPolicyEvaluation(v SessionViolationDetectionPolicyEvaluation) {
	o.PolicyEvaluation = &v
}

func (o SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyEvaluation) {
		toSerialize["policyEvaluation"] = o.PolicyEvaluation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) UnmarshalJSON(data []byte) (err error) {
	varSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection := _SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection{}

	err = json.Unmarshal(data, &varSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection)

	if err != nil {
		return err
	}

	*o = SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection(varSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "policyEvaluation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection struct {
	value *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection
	isSet bool
}

func (v NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) Get() *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection {
	return v.value
}

func (v *NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) Set(val *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection(val *SessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) *NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection {
	return &NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection{value: val, isSet: true}
}

func (v NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionViolationDetectionPolicyRuleAllOfActionsSessionViolationDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
