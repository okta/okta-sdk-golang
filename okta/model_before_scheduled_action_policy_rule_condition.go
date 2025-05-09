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

// BeforeScheduledActionPolicyRuleCondition struct for BeforeScheduledActionPolicyRuleCondition
type BeforeScheduledActionPolicyRuleCondition struct {
	Duration *Duration `json:"duration,omitempty"`
	LifecycleAction *ScheduledUserLifecycleAction `json:"lifecycleAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BeforeScheduledActionPolicyRuleCondition BeforeScheduledActionPolicyRuleCondition

// NewBeforeScheduledActionPolicyRuleCondition instantiates a new BeforeScheduledActionPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeforeScheduledActionPolicyRuleCondition() *BeforeScheduledActionPolicyRuleCondition {
	this := BeforeScheduledActionPolicyRuleCondition{}
	return &this
}

// NewBeforeScheduledActionPolicyRuleConditionWithDefaults instantiates a new BeforeScheduledActionPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeforeScheduledActionPolicyRuleConditionWithDefaults() *BeforeScheduledActionPolicyRuleCondition {
	this := BeforeScheduledActionPolicyRuleCondition{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BeforeScheduledActionPolicyRuleCondition) GetDuration() Duration {
	if o == nil || o.Duration == nil {
		var ret Duration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeforeScheduledActionPolicyRuleCondition) GetDurationOk() (*Duration, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BeforeScheduledActionPolicyRuleCondition) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given Duration and assigns it to the Duration field.
func (o *BeforeScheduledActionPolicyRuleCondition) SetDuration(v Duration) {
	o.Duration = &v
}

// GetLifecycleAction returns the LifecycleAction field value if set, zero value otherwise.
func (o *BeforeScheduledActionPolicyRuleCondition) GetLifecycleAction() ScheduledUserLifecycleAction {
	if o == nil || o.LifecycleAction == nil {
		var ret ScheduledUserLifecycleAction
		return ret
	}
	return *o.LifecycleAction
}

// GetLifecycleActionOk returns a tuple with the LifecycleAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeforeScheduledActionPolicyRuleCondition) GetLifecycleActionOk() (*ScheduledUserLifecycleAction, bool) {
	if o == nil || o.LifecycleAction == nil {
		return nil, false
	}
	return o.LifecycleAction, true
}

// HasLifecycleAction returns a boolean if a field has been set.
func (o *BeforeScheduledActionPolicyRuleCondition) HasLifecycleAction() bool {
	if o != nil && o.LifecycleAction != nil {
		return true
	}

	return false
}

// SetLifecycleAction gets a reference to the given ScheduledUserLifecycleAction and assigns it to the LifecycleAction field.
func (o *BeforeScheduledActionPolicyRuleCondition) SetLifecycleAction(v ScheduledUserLifecycleAction) {
	o.LifecycleAction = &v
}

func (o BeforeScheduledActionPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.LifecycleAction != nil {
		toSerialize["lifecycleAction"] = o.LifecycleAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BeforeScheduledActionPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varBeforeScheduledActionPolicyRuleCondition := _BeforeScheduledActionPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varBeforeScheduledActionPolicyRuleCondition)
	if err == nil {
		*o = BeforeScheduledActionPolicyRuleCondition(varBeforeScheduledActionPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "lifecycleAction")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBeforeScheduledActionPolicyRuleCondition struct {
	value *BeforeScheduledActionPolicyRuleCondition
	isSet bool
}

func (v NullableBeforeScheduledActionPolicyRuleCondition) Get() *BeforeScheduledActionPolicyRuleCondition {
	return v.value
}

func (v *NullableBeforeScheduledActionPolicyRuleCondition) Set(val *BeforeScheduledActionPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableBeforeScheduledActionPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableBeforeScheduledActionPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeforeScheduledActionPolicyRuleCondition(val *BeforeScheduledActionPolicyRuleCondition) *NullableBeforeScheduledActionPolicyRuleCondition {
	return &NullableBeforeScheduledActionPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableBeforeScheduledActionPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeforeScheduledActionPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

