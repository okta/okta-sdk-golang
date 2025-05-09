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

// ContinuousAccessPolicyRuleRunWorkflowWorkflow This action runs a workflow
type ContinuousAccessPolicyRuleRunWorkflowWorkflow struct {
	// The `id` of the workflow that runs.
	Id *int32 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleRunWorkflowWorkflow ContinuousAccessPolicyRuleRunWorkflowWorkflow

// NewContinuousAccessPolicyRuleRunWorkflowWorkflow instantiates a new ContinuousAccessPolicyRuleRunWorkflowWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleRunWorkflowWorkflow() *ContinuousAccessPolicyRuleRunWorkflowWorkflow {
	this := ContinuousAccessPolicyRuleRunWorkflowWorkflow{}
	return &this
}

// NewContinuousAccessPolicyRuleRunWorkflowWorkflowWithDefaults instantiates a new ContinuousAccessPolicyRuleRunWorkflowWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleRunWorkflowWorkflowWithDefaults() *ContinuousAccessPolicyRuleRunWorkflowWorkflow {
	this := ContinuousAccessPolicyRuleRunWorkflowWorkflow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleRunWorkflowWorkflow) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleRunWorkflowWorkflow) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleRunWorkflowWorkflow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContinuousAccessPolicyRuleRunWorkflowWorkflow) SetId(v int32) {
	o.Id = &v
}

func (o ContinuousAccessPolicyRuleRunWorkflowWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleRunWorkflowWorkflow) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleRunWorkflowWorkflow := _ContinuousAccessPolicyRuleRunWorkflowWorkflow{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleRunWorkflowWorkflow)
	if err == nil {
		*o = ContinuousAccessPolicyRuleRunWorkflowWorkflow(varContinuousAccessPolicyRuleRunWorkflowWorkflow)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleRunWorkflowWorkflow struct {
	value *ContinuousAccessPolicyRuleRunWorkflowWorkflow
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleRunWorkflowWorkflow) Get() *ContinuousAccessPolicyRuleRunWorkflowWorkflow {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleRunWorkflowWorkflow) Set(val *ContinuousAccessPolicyRuleRunWorkflowWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleRunWorkflowWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleRunWorkflowWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleRunWorkflowWorkflow(val *ContinuousAccessPolicyRuleRunWorkflowWorkflow) *NullableContinuousAccessPolicyRuleRunWorkflowWorkflow {
	return &NullableContinuousAccessPolicyRuleRunWorkflowWorkflow{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleRunWorkflowWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleRunWorkflowWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

