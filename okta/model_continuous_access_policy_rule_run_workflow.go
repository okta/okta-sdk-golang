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

// ContinuousAccessPolicyRuleRunWorkflow struct for ContinuousAccessPolicyRuleRunWorkflow
type ContinuousAccessPolicyRuleRunWorkflow struct {
	Action *string `json:"action,omitempty"`
	Workflow *ContinuousAccessPolicyRuleRunWorkflowWorkflow `json:"workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleRunWorkflow ContinuousAccessPolicyRuleRunWorkflow

// NewContinuousAccessPolicyRuleRunWorkflow instantiates a new ContinuousAccessPolicyRuleRunWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleRunWorkflow() *ContinuousAccessPolicyRuleRunWorkflow {
	this := ContinuousAccessPolicyRuleRunWorkflow{}
	return &this
}

// NewContinuousAccessPolicyRuleRunWorkflowWithDefaults instantiates a new ContinuousAccessPolicyRuleRunWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleRunWorkflowWithDefaults() *ContinuousAccessPolicyRuleRunWorkflow {
	this := ContinuousAccessPolicyRuleRunWorkflow{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleRunWorkflow) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleRunWorkflow) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleRunWorkflow) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ContinuousAccessPolicyRuleRunWorkflow) SetAction(v string) {
	o.Action = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleRunWorkflow) GetWorkflow() ContinuousAccessPolicyRuleRunWorkflowWorkflow {
	if o == nil || o.Workflow == nil {
		var ret ContinuousAccessPolicyRuleRunWorkflowWorkflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleRunWorkflow) GetWorkflowOk() (*ContinuousAccessPolicyRuleRunWorkflowWorkflow, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleRunWorkflow) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given ContinuousAccessPolicyRuleRunWorkflowWorkflow and assigns it to the Workflow field.
func (o *ContinuousAccessPolicyRuleRunWorkflow) SetWorkflow(v ContinuousAccessPolicyRuleRunWorkflowWorkflow) {
	o.Workflow = &v
}

func (o ContinuousAccessPolicyRuleRunWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleRunWorkflow) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleRunWorkflow := _ContinuousAccessPolicyRuleRunWorkflow{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleRunWorkflow)
	if err == nil {
		*o = ContinuousAccessPolicyRuleRunWorkflow(varContinuousAccessPolicyRuleRunWorkflow)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "workflow")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleRunWorkflow struct {
	value *ContinuousAccessPolicyRuleRunWorkflow
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleRunWorkflow) Get() *ContinuousAccessPolicyRuleRunWorkflow {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleRunWorkflow) Set(val *ContinuousAccessPolicyRuleRunWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleRunWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleRunWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleRunWorkflow(val *ContinuousAccessPolicyRuleRunWorkflow) *NullableContinuousAccessPolicyRuleRunWorkflow {
	return &NullableContinuousAccessPolicyRuleRunWorkflow{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleRunWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleRunWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

