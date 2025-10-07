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

// checks if the EntityRiskPolicyRuleActionRunWorkflowWorkflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRiskPolicyRuleActionRunWorkflowWorkflow{}

// EntityRiskPolicyRuleActionRunWorkflowWorkflow This action runs a workflow
type EntityRiskPolicyRuleActionRunWorkflowWorkflow struct {
	// The `id` of the workflow that runs.
	Id                   *int32 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleActionRunWorkflowWorkflow EntityRiskPolicyRuleActionRunWorkflowWorkflow

// NewEntityRiskPolicyRuleActionRunWorkflowWorkflow instantiates a new EntityRiskPolicyRuleActionRunWorkflowWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleActionRunWorkflowWorkflow() *EntityRiskPolicyRuleActionRunWorkflowWorkflow {
	this := EntityRiskPolicyRuleActionRunWorkflowWorkflow{}
	return &this
}

// NewEntityRiskPolicyRuleActionRunWorkflowWorkflowWithDefaults instantiates a new EntityRiskPolicyRuleActionRunWorkflowWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleActionRunWorkflowWorkflowWithDefaults() *EntityRiskPolicyRuleActionRunWorkflowWorkflow {
	this := EntityRiskPolicyRuleActionRunWorkflowWorkflow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleActionRunWorkflowWorkflow) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleActionRunWorkflowWorkflow) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleActionRunWorkflowWorkflow) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EntityRiskPolicyRuleActionRunWorkflowWorkflow) SetId(v int32) {
	o.Id = &v
}

func (o EntityRiskPolicyRuleActionRunWorkflowWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRiskPolicyRuleActionRunWorkflowWorkflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntityRiskPolicyRuleActionRunWorkflowWorkflow) UnmarshalJSON(data []byte) (err error) {
	varEntityRiskPolicyRuleActionRunWorkflowWorkflow := _EntityRiskPolicyRuleActionRunWorkflowWorkflow{}

	err = json.Unmarshal(data, &varEntityRiskPolicyRuleActionRunWorkflowWorkflow)

	if err != nil {
		return err
	}

	*o = EntityRiskPolicyRuleActionRunWorkflowWorkflow(varEntityRiskPolicyRuleActionRunWorkflowWorkflow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow struct {
	value *EntityRiskPolicyRuleActionRunWorkflowWorkflow
	isSet bool
}

func (v NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow) Get() *EntityRiskPolicyRuleActionRunWorkflowWorkflow {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow) Set(val *EntityRiskPolicyRuleActionRunWorkflowWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleActionRunWorkflowWorkflow(val *EntityRiskPolicyRuleActionRunWorkflowWorkflow) *NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow {
	return &NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleActionRunWorkflowWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
