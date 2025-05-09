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

// SimulatePolicyEvaluations struct for SimulatePolicyEvaluations
type SimulatePolicyEvaluations struct {
	Evaluated *SimulatePolicyEvaluationsEvaluated `json:"evaluated,omitempty"`
	// The policy type of the simulate operation
	PolicyType []string `json:"policyType,omitempty"`
	Result *SimulatePolicyResult `json:"result,omitempty"`
	// The result of this entity evaluation
	Status *string `json:"status,omitempty"`
	Undefined *SimulatePolicyEvaluationsUndefined `json:"undefined,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulatePolicyEvaluations SimulatePolicyEvaluations

// NewSimulatePolicyEvaluations instantiates a new SimulatePolicyEvaluations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulatePolicyEvaluations() *SimulatePolicyEvaluations {
	this := SimulatePolicyEvaluations{}
	return &this
}

// NewSimulatePolicyEvaluationsWithDefaults instantiates a new SimulatePolicyEvaluations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulatePolicyEvaluationsWithDefaults() *SimulatePolicyEvaluations {
	this := SimulatePolicyEvaluations{}
	return &this
}

// GetEvaluated returns the Evaluated field value if set, zero value otherwise.
func (o *SimulatePolicyEvaluations) GetEvaluated() SimulatePolicyEvaluationsEvaluated {
	if o == nil || o.Evaluated == nil {
		var ret SimulatePolicyEvaluationsEvaluated
		return ret
	}
	return *o.Evaluated
}

// GetEvaluatedOk returns a tuple with the Evaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyEvaluations) GetEvaluatedOk() (*SimulatePolicyEvaluationsEvaluated, bool) {
	if o == nil || o.Evaluated == nil {
		return nil, false
	}
	return o.Evaluated, true
}

// HasEvaluated returns a boolean if a field has been set.
func (o *SimulatePolicyEvaluations) HasEvaluated() bool {
	if o != nil && o.Evaluated != nil {
		return true
	}

	return false
}

// SetEvaluated gets a reference to the given SimulatePolicyEvaluationsEvaluated and assigns it to the Evaluated field.
func (o *SimulatePolicyEvaluations) SetEvaluated(v SimulatePolicyEvaluationsEvaluated) {
	o.Evaluated = &v
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *SimulatePolicyEvaluations) GetPolicyType() []string {
	if o == nil || o.PolicyType == nil {
		var ret []string
		return ret
	}
	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyEvaluations) GetPolicyTypeOk() ([]string, bool) {
	if o == nil || o.PolicyType == nil {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *SimulatePolicyEvaluations) HasPolicyType() bool {
	if o != nil && o.PolicyType != nil {
		return true
	}

	return false
}

// SetPolicyType gets a reference to the given []string and assigns it to the PolicyType field.
func (o *SimulatePolicyEvaluations) SetPolicyType(v []string) {
	o.PolicyType = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SimulatePolicyEvaluations) GetResult() SimulatePolicyResult {
	if o == nil || o.Result == nil {
		var ret SimulatePolicyResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyEvaluations) GetResultOk() (*SimulatePolicyResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SimulatePolicyEvaluations) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given SimulatePolicyResult and assigns it to the Result field.
func (o *SimulatePolicyEvaluations) SetResult(v SimulatePolicyResult) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SimulatePolicyEvaluations) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyEvaluations) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SimulatePolicyEvaluations) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SimulatePolicyEvaluations) SetStatus(v string) {
	o.Status = &v
}

// GetUndefined returns the Undefined field value if set, zero value otherwise.
func (o *SimulatePolicyEvaluations) GetUndefined() SimulatePolicyEvaluationsUndefined {
	if o == nil || o.Undefined == nil {
		var ret SimulatePolicyEvaluationsUndefined
		return ret
	}
	return *o.Undefined
}

// GetUndefinedOk returns a tuple with the Undefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyEvaluations) GetUndefinedOk() (*SimulatePolicyEvaluationsUndefined, bool) {
	if o == nil || o.Undefined == nil {
		return nil, false
	}
	return o.Undefined, true
}

// HasUndefined returns a boolean if a field has been set.
func (o *SimulatePolicyEvaluations) HasUndefined() bool {
	if o != nil && o.Undefined != nil {
		return true
	}

	return false
}

// SetUndefined gets a reference to the given SimulatePolicyEvaluationsUndefined and assigns it to the Undefined field.
func (o *SimulatePolicyEvaluations) SetUndefined(v SimulatePolicyEvaluationsUndefined) {
	o.Undefined = &v
}

func (o SimulatePolicyEvaluations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Evaluated != nil {
		toSerialize["evaluated"] = o.Evaluated
	}
	if o.PolicyType != nil {
		toSerialize["policyType"] = o.PolicyType
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Undefined != nil {
		toSerialize["undefined"] = o.Undefined
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SimulatePolicyEvaluations) UnmarshalJSON(bytes []byte) (err error) {
	varSimulatePolicyEvaluations := _SimulatePolicyEvaluations{}

	err = json.Unmarshal(bytes, &varSimulatePolicyEvaluations)
	if err == nil {
		*o = SimulatePolicyEvaluations(varSimulatePolicyEvaluations)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "evaluated")
		delete(additionalProperties, "policyType")
		delete(additionalProperties, "result")
		delete(additionalProperties, "status")
		delete(additionalProperties, "undefined")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSimulatePolicyEvaluations struct {
	value *SimulatePolicyEvaluations
	isSet bool
}

func (v NullableSimulatePolicyEvaluations) Get() *SimulatePolicyEvaluations {
	return v.value
}

func (v *NullableSimulatePolicyEvaluations) Set(val *SimulatePolicyEvaluations) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatePolicyEvaluations) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatePolicyEvaluations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatePolicyEvaluations(val *SimulatePolicyEvaluations) *NullableSimulatePolicyEvaluations {
	return &NullableSimulatePolicyEvaluations{value: val, isSet: true}
}

func (v NullableSimulatePolicyEvaluations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatePolicyEvaluations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

