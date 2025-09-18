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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the BehaviorRuleSettingsHistoryBased type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BehaviorRuleSettingsHistoryBased{}

// BehaviorRuleSettingsHistoryBased struct for BehaviorRuleSettingsHistoryBased
type BehaviorRuleSettingsHistoryBased struct {
	MaxEventsUsedForEvaluation   *int32 `json:"maxEventsUsedForEvaluation,omitempty"`
	MinEventsNeededForEvaluation *int32 `json:"minEventsNeededForEvaluation,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _BehaviorRuleSettingsHistoryBased BehaviorRuleSettingsHistoryBased

// NewBehaviorRuleSettingsHistoryBased instantiates a new BehaviorRuleSettingsHistoryBased object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsHistoryBased() *BehaviorRuleSettingsHistoryBased {
	this := BehaviorRuleSettingsHistoryBased{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// NewBehaviorRuleSettingsHistoryBasedWithDefaults instantiates a new BehaviorRuleSettingsHistoryBased object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsHistoryBasedWithDefaults() *BehaviorRuleSettingsHistoryBased {
	this := BehaviorRuleSettingsHistoryBased{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsHistoryBased) GetMaxEventsUsedForEvaluation() int32 {
	if o == nil || IsNil(o.MaxEventsUsedForEvaluation) {
		var ret int32
		return ret
	}
	return *o.MaxEventsUsedForEvaluation
}

// GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsHistoryBased) GetMaxEventsUsedForEvaluationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxEventsUsedForEvaluation) {
		return nil, false
	}
	return o.MaxEventsUsedForEvaluation, true
}

// HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsHistoryBased) HasMaxEventsUsedForEvaluation() bool {
	if o != nil && !IsNil(o.MaxEventsUsedForEvaluation) {
		return true
	}

	return false
}

// SetMaxEventsUsedForEvaluation gets a reference to the given int32 and assigns it to the MaxEventsUsedForEvaluation field.
func (o *BehaviorRuleSettingsHistoryBased) SetMaxEventsUsedForEvaluation(v int32) {
	o.MaxEventsUsedForEvaluation = &v
}

// GetMinEventsNeededForEvaluation returns the MinEventsNeededForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsHistoryBased) GetMinEventsNeededForEvaluation() int32 {
	if o == nil || IsNil(o.MinEventsNeededForEvaluation) {
		var ret int32
		return ret
	}
	return *o.MinEventsNeededForEvaluation
}

// GetMinEventsNeededForEvaluationOk returns a tuple with the MinEventsNeededForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsHistoryBased) GetMinEventsNeededForEvaluationOk() (*int32, bool) {
	if o == nil || IsNil(o.MinEventsNeededForEvaluation) {
		return nil, false
	}
	return o.MinEventsNeededForEvaluation, true
}

// HasMinEventsNeededForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsHistoryBased) HasMinEventsNeededForEvaluation() bool {
	if o != nil && !IsNil(o.MinEventsNeededForEvaluation) {
		return true
	}

	return false
}

// SetMinEventsNeededForEvaluation gets a reference to the given int32 and assigns it to the MinEventsNeededForEvaluation field.
func (o *BehaviorRuleSettingsHistoryBased) SetMinEventsNeededForEvaluation(v int32) {
	o.MinEventsNeededForEvaluation = &v
}

func (o BehaviorRuleSettingsHistoryBased) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BehaviorRuleSettingsHistoryBased) ToMap() (map[string]interface{}, error) {
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

func (o *BehaviorRuleSettingsHistoryBased) UnmarshalJSON(data []byte) (err error) {
	varBehaviorRuleSettingsHistoryBased := _BehaviorRuleSettingsHistoryBased{}

	err = json.Unmarshal(data, &varBehaviorRuleSettingsHistoryBased)

	if err != nil {
		return err
	}

	*o = BehaviorRuleSettingsHistoryBased(varBehaviorRuleSettingsHistoryBased)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxEventsUsedForEvaluation")
		delete(additionalProperties, "minEventsNeededForEvaluation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBehaviorRuleSettingsHistoryBased struct {
	value *BehaviorRuleSettingsHistoryBased
	isSet bool
}

func (v NullableBehaviorRuleSettingsHistoryBased) Get() *BehaviorRuleSettingsHistoryBased {
	return v.value
}

func (v *NullableBehaviorRuleSettingsHistoryBased) Set(val *BehaviorRuleSettingsHistoryBased) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsHistoryBased) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsHistoryBased) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsHistoryBased(val *BehaviorRuleSettingsHistoryBased) *NullableBehaviorRuleSettingsHistoryBased {
	return &NullableBehaviorRuleSettingsHistoryBased{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsHistoryBased) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsHistoryBased) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
