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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SimulatePolicyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimulatePolicyResult{}

// SimulatePolicyResult The result of the policy evaluation
type SimulatePolicyResult struct {
	Policies             []SimulateResultPoliciesItems `json:"policies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulatePolicyResult SimulatePolicyResult

// NewSimulatePolicyResult instantiates a new SimulatePolicyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulatePolicyResult() *SimulatePolicyResult {
	this := SimulatePolicyResult{}
	return &this
}

// NewSimulatePolicyResultWithDefaults instantiates a new SimulatePolicyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulatePolicyResultWithDefaults() *SimulatePolicyResult {
	this := SimulatePolicyResult{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *SimulatePolicyResult) GetPolicies() []SimulateResultPoliciesItems {
	if o == nil || IsNil(o.Policies) {
		var ret []SimulateResultPoliciesItems
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyResult) GetPoliciesOk() ([]SimulateResultPoliciesItems, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *SimulatePolicyResult) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []SimulateResultPoliciesItems and assigns it to the Policies field.
func (o *SimulatePolicyResult) SetPolicies(v []SimulateResultPoliciesItems) {
	o.Policies = v
}

func (o SimulatePolicyResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimulatePolicyResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimulatePolicyResult) UnmarshalJSON(data []byte) (err error) {
	varSimulatePolicyResult := _SimulatePolicyResult{}

	err = json.Unmarshal(data, &varSimulatePolicyResult)

	if err != nil {
		return err
	}

	*o = SimulatePolicyResult(varSimulatePolicyResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "policies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimulatePolicyResult struct {
	value *SimulatePolicyResult
	isSet bool
}

func (v NullableSimulatePolicyResult) Get() *SimulatePolicyResult {
	return v.value
}

func (v *NullableSimulatePolicyResult) Set(val *SimulatePolicyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatePolicyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatePolicyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatePolicyResult(val *SimulatePolicyResult) *NullableSimulatePolicyResult {
	return &NullableSimulatePolicyResult{value: val, isSet: true}
}

func (v NullableSimulatePolicyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatePolicyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
