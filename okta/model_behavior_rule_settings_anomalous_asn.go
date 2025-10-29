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

// checks if the BehaviorRuleSettingsAnomalousASN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BehaviorRuleSettingsAnomalousASN{}

// BehaviorRuleSettingsAnomalousASN struct for BehaviorRuleSettingsAnomalousASN
type BehaviorRuleSettingsAnomalousASN struct {
	MaxEventsUsedForEvaluation   *int32 `json:"maxEventsUsedForEvaluation,omitempty"`
	MinEventsNeededForEvaluation *int32 `json:"minEventsNeededForEvaluation,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _BehaviorRuleSettingsAnomalousASN BehaviorRuleSettingsAnomalousASN

// NewBehaviorRuleSettingsAnomalousASN instantiates a new BehaviorRuleSettingsAnomalousASN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsAnomalousASN() *BehaviorRuleSettingsAnomalousASN {
	this := BehaviorRuleSettingsAnomalousASN{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// NewBehaviorRuleSettingsAnomalousASNWithDefaults instantiates a new BehaviorRuleSettingsAnomalousASN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsAnomalousASNWithDefaults() *BehaviorRuleSettingsAnomalousASN {
	this := BehaviorRuleSettingsAnomalousASN{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousASN) GetMaxEventsUsedForEvaluation() int32 {
	if o == nil || IsNil(o.MaxEventsUsedForEvaluation) {
		var ret int32
		return ret
	}
	return *o.MaxEventsUsedForEvaluation
}

// GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousASN) GetMaxEventsUsedForEvaluationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxEventsUsedForEvaluation) {
		return nil, false
	}
	return o.MaxEventsUsedForEvaluation, true
}

// HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousASN) HasMaxEventsUsedForEvaluation() bool {
	if o != nil && !IsNil(o.MaxEventsUsedForEvaluation) {
		return true
	}

	return false
}

// SetMaxEventsUsedForEvaluation gets a reference to the given int32 and assigns it to the MaxEventsUsedForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousASN) SetMaxEventsUsedForEvaluation(v int32) {
	o.MaxEventsUsedForEvaluation = &v
}

// GetMinEventsNeededForEvaluation returns the MinEventsNeededForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousASN) GetMinEventsNeededForEvaluation() int32 {
	if o == nil || IsNil(o.MinEventsNeededForEvaluation) {
		var ret int32
		return ret
	}
	return *o.MinEventsNeededForEvaluation
}

// GetMinEventsNeededForEvaluationOk returns a tuple with the MinEventsNeededForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousASN) GetMinEventsNeededForEvaluationOk() (*int32, bool) {
	if o == nil || IsNil(o.MinEventsNeededForEvaluation) {
		return nil, false
	}
	return o.MinEventsNeededForEvaluation, true
}

// HasMinEventsNeededForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousASN) HasMinEventsNeededForEvaluation() bool {
	if o != nil && !IsNil(o.MinEventsNeededForEvaluation) {
		return true
	}

	return false
}

// SetMinEventsNeededForEvaluation gets a reference to the given int32 and assigns it to the MinEventsNeededForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousASN) SetMinEventsNeededForEvaluation(v int32) {
	o.MinEventsNeededForEvaluation = &v
}

func (o BehaviorRuleSettingsAnomalousASN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BehaviorRuleSettingsAnomalousASN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxEventsUsedForEvaluation) {
		toSerialize["maxEventsUsedForEvaluation"] = o.MaxEventsUsedForEvaluation
	}
	if !IsNil(o.MinEventsNeededForEvaluation) {
		toSerialize["minEventsNeededForEvaluation"] = o.MinEventsNeededForEvaluation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BehaviorRuleSettingsAnomalousASN) UnmarshalJSON(data []byte) (err error) {
	varBehaviorRuleSettingsAnomalousASN := _BehaviorRuleSettingsAnomalousASN{}

	err = json.Unmarshal(data, &varBehaviorRuleSettingsAnomalousASN)

	if err != nil {
		return err
	}

	*o = BehaviorRuleSettingsAnomalousASN(varBehaviorRuleSettingsAnomalousASN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxEventsUsedForEvaluation")
		delete(additionalProperties, "minEventsNeededForEvaluation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBehaviorRuleSettingsAnomalousASN struct {
	value *BehaviorRuleSettingsAnomalousASN
	isSet bool
}

func (v NullableBehaviorRuleSettingsAnomalousASN) Get() *BehaviorRuleSettingsAnomalousASN {
	return v.value
}

func (v *NullableBehaviorRuleSettingsAnomalousASN) Set(val *BehaviorRuleSettingsAnomalousASN) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsAnomalousASN) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsAnomalousASN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsAnomalousASN(val *BehaviorRuleSettingsAnomalousASN) *NullableBehaviorRuleSettingsAnomalousASN {
	return &NullableBehaviorRuleSettingsAnomalousASN{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsAnomalousASN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsAnomalousASN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
