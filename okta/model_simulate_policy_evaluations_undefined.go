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

// SimulatePolicyEvaluationsUndefined A list of undefined but not matched policies and rules
type SimulatePolicyEvaluationsUndefined struct {
	Policies []SimulateResultPoliciesItems `json:"policies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulatePolicyEvaluationsUndefined SimulatePolicyEvaluationsUndefined

// NewSimulatePolicyEvaluationsUndefined instantiates a new SimulatePolicyEvaluationsUndefined object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulatePolicyEvaluationsUndefined() *SimulatePolicyEvaluationsUndefined {
	this := SimulatePolicyEvaluationsUndefined{}
	return &this
}

// NewSimulatePolicyEvaluationsUndefinedWithDefaults instantiates a new SimulatePolicyEvaluationsUndefined object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulatePolicyEvaluationsUndefinedWithDefaults() *SimulatePolicyEvaluationsUndefined {
	this := SimulatePolicyEvaluationsUndefined{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *SimulatePolicyEvaluationsUndefined) GetPolicies() []SimulateResultPoliciesItems {
	if o == nil || o.Policies == nil {
		var ret []SimulateResultPoliciesItems
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyEvaluationsUndefined) GetPoliciesOk() ([]SimulateResultPoliciesItems, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *SimulatePolicyEvaluationsUndefined) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []SimulateResultPoliciesItems and assigns it to the Policies field.
func (o *SimulatePolicyEvaluationsUndefined) SetPolicies(v []SimulateResultPoliciesItems) {
	o.Policies = v
}

func (o SimulatePolicyEvaluationsUndefined) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SimulatePolicyEvaluationsUndefined) UnmarshalJSON(bytes []byte) (err error) {
	varSimulatePolicyEvaluationsUndefined := _SimulatePolicyEvaluationsUndefined{}

	err = json.Unmarshal(bytes, &varSimulatePolicyEvaluationsUndefined)
	if err == nil {
		*o = SimulatePolicyEvaluationsUndefined(varSimulatePolicyEvaluationsUndefined)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "policies")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSimulatePolicyEvaluationsUndefined struct {
	value *SimulatePolicyEvaluationsUndefined
	isSet bool
}

func (v NullableSimulatePolicyEvaluationsUndefined) Get() *SimulatePolicyEvaluationsUndefined {
	return v.value
}

func (v *NullableSimulatePolicyEvaluationsUndefined) Set(val *SimulatePolicyEvaluationsUndefined) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatePolicyEvaluationsUndefined) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatePolicyEvaluationsUndefined) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatePolicyEvaluationsUndefined(val *SimulatePolicyEvaluationsUndefined) *NullableSimulatePolicyEvaluationsUndefined {
	return &NullableSimulatePolicyEvaluationsUndefined{value: val, isSet: true}
}

func (v NullableSimulatePolicyEvaluationsUndefined) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatePolicyEvaluationsUndefined) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

