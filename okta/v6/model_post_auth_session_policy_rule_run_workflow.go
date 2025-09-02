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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// PostAuthSessionPolicyRuleRunWorkflow struct for PostAuthSessionPolicyRuleRunWorkflow
type PostAuthSessionPolicyRuleRunWorkflow struct {
	Action *string `json:"action,omitempty"`
	Workflow *EntityRiskPolicyRuleActionRunWorkflowWorkflow `json:"workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionPolicyRuleRunWorkflow PostAuthSessionPolicyRuleRunWorkflow

// NewPostAuthSessionPolicyRuleRunWorkflow instantiates a new PostAuthSessionPolicyRuleRunWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionPolicyRuleRunWorkflow() *PostAuthSessionPolicyRuleRunWorkflow {
	this := PostAuthSessionPolicyRuleRunWorkflow{}
	return &this
}

// NewPostAuthSessionPolicyRuleRunWorkflowWithDefaults instantiates a new PostAuthSessionPolicyRuleRunWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionPolicyRuleRunWorkflowWithDefaults() *PostAuthSessionPolicyRuleRunWorkflow {
	this := PostAuthSessionPolicyRuleRunWorkflow{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRuleRunWorkflow) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRuleRunWorkflow) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRuleRunWorkflow) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PostAuthSessionPolicyRuleRunWorkflow) SetAction(v string) {
	o.Action = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *PostAuthSessionPolicyRuleRunWorkflow) GetWorkflow() EntityRiskPolicyRuleActionRunWorkflowWorkflow {
	if o == nil || o.Workflow == nil {
		var ret EntityRiskPolicyRuleActionRunWorkflowWorkflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionPolicyRuleRunWorkflow) GetWorkflowOk() (*EntityRiskPolicyRuleActionRunWorkflowWorkflow, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *PostAuthSessionPolicyRuleRunWorkflow) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given EntityRiskPolicyRuleActionRunWorkflowWorkflow and assigns it to the Workflow field.
func (o *PostAuthSessionPolicyRuleRunWorkflow) SetWorkflow(v EntityRiskPolicyRuleActionRunWorkflowWorkflow) {
	o.Workflow = &v
}

func (o PostAuthSessionPolicyRuleRunWorkflow) MarshalJSON() ([]byte, error) {
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

func (o *PostAuthSessionPolicyRuleRunWorkflow) UnmarshalJSON(bytes []byte) (err error) {
	varPostAuthSessionPolicyRuleRunWorkflow := _PostAuthSessionPolicyRuleRunWorkflow{}

	err = json.Unmarshal(bytes, &varPostAuthSessionPolicyRuleRunWorkflow)
	if err == nil {
		*o = PostAuthSessionPolicyRuleRunWorkflow(varPostAuthSessionPolicyRuleRunWorkflow)
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

type NullablePostAuthSessionPolicyRuleRunWorkflow struct {
	value *PostAuthSessionPolicyRuleRunWorkflow
	isSet bool
}

func (v NullablePostAuthSessionPolicyRuleRunWorkflow) Get() *PostAuthSessionPolicyRuleRunWorkflow {
	return v.value
}

func (v *NullablePostAuthSessionPolicyRuleRunWorkflow) Set(val *PostAuthSessionPolicyRuleRunWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionPolicyRuleRunWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionPolicyRuleRunWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionPolicyRuleRunWorkflow(val *PostAuthSessionPolicyRuleRunWorkflow) *NullablePostAuthSessionPolicyRuleRunWorkflow {
	return &NullablePostAuthSessionPolicyRuleRunWorkflow{value: val, isSet: true}
}

func (v NullablePostAuthSessionPolicyRuleRunWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionPolicyRuleRunWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

