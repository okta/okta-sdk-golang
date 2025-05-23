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

// BehaviorRuleSettingsAnomalousIP struct for BehaviorRuleSettingsAnomalousIP
type BehaviorRuleSettingsAnomalousIP struct {
	MaxEventsUsedForEvaluation *int32 `json:"maxEventsUsedForEvaluation,omitempty"`
	MinEventsNeededForEvaluation *int32 `json:"minEventsNeededForEvaluation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleSettingsAnomalousIP BehaviorRuleSettingsAnomalousIP

// NewBehaviorRuleSettingsAnomalousIP instantiates a new BehaviorRuleSettingsAnomalousIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsAnomalousIP() *BehaviorRuleSettingsAnomalousIP {
	this := BehaviorRuleSettingsAnomalousIP{}
	var maxEventsUsedForEvaluation int32 = 50
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// NewBehaviorRuleSettingsAnomalousIPWithDefaults instantiates a new BehaviorRuleSettingsAnomalousIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsAnomalousIPWithDefaults() *BehaviorRuleSettingsAnomalousIP {
	this := BehaviorRuleSettingsAnomalousIP{}
	var maxEventsUsedForEvaluation int32 = 50
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousIP) GetMaxEventsUsedForEvaluation() int32 {
	if o == nil || o.MaxEventsUsedForEvaluation == nil {
		var ret int32
		return ret
	}
	return *o.MaxEventsUsedForEvaluation
}

// GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousIP) GetMaxEventsUsedForEvaluationOk() (*int32, bool) {
	if o == nil || o.MaxEventsUsedForEvaluation == nil {
		return nil, false
	}
	return o.MaxEventsUsedForEvaluation, true
}

// HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousIP) HasMaxEventsUsedForEvaluation() bool {
	if o != nil && o.MaxEventsUsedForEvaluation != nil {
		return true
	}

	return false
}

// SetMaxEventsUsedForEvaluation gets a reference to the given int32 and assigns it to the MaxEventsUsedForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousIP) SetMaxEventsUsedForEvaluation(v int32) {
	o.MaxEventsUsedForEvaluation = &v
}

// GetMinEventsNeededForEvaluation returns the MinEventsNeededForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousIP) GetMinEventsNeededForEvaluation() int32 {
	if o == nil || o.MinEventsNeededForEvaluation == nil {
		var ret int32
		return ret
	}
	return *o.MinEventsNeededForEvaluation
}

// GetMinEventsNeededForEvaluationOk returns a tuple with the MinEventsNeededForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousIP) GetMinEventsNeededForEvaluationOk() (*int32, bool) {
	if o == nil || o.MinEventsNeededForEvaluation == nil {
		return nil, false
	}
	return o.MinEventsNeededForEvaluation, true
}

// HasMinEventsNeededForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousIP) HasMinEventsNeededForEvaluation() bool {
	if o != nil && o.MinEventsNeededForEvaluation != nil {
		return true
	}

	return false
}

// SetMinEventsNeededForEvaluation gets a reference to the given int32 and assigns it to the MinEventsNeededForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousIP) SetMinEventsNeededForEvaluation(v int32) {
	o.MinEventsNeededForEvaluation = &v
}

func (o BehaviorRuleSettingsAnomalousIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxEventsUsedForEvaluation != nil {
		toSerialize["maxEventsUsedForEvaluation"] = o.MaxEventsUsedForEvaluation
	}
	if o.MinEventsNeededForEvaluation != nil {
		toSerialize["minEventsNeededForEvaluation"] = o.MinEventsNeededForEvaluation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BehaviorRuleSettingsAnomalousIP) UnmarshalJSON(bytes []byte) (err error) {
	varBehaviorRuleSettingsAnomalousIP := _BehaviorRuleSettingsAnomalousIP{}

	err = json.Unmarshal(bytes, &varBehaviorRuleSettingsAnomalousIP)
	if err == nil {
		*o = BehaviorRuleSettingsAnomalousIP(varBehaviorRuleSettingsAnomalousIP)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "maxEventsUsedForEvaluation")
		delete(additionalProperties, "minEventsNeededForEvaluation")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBehaviorRuleSettingsAnomalousIP struct {
	value *BehaviorRuleSettingsAnomalousIP
	isSet bool
}

func (v NullableBehaviorRuleSettingsAnomalousIP) Get() *BehaviorRuleSettingsAnomalousIP {
	return v.value
}

func (v *NullableBehaviorRuleSettingsAnomalousIP) Set(val *BehaviorRuleSettingsAnomalousIP) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsAnomalousIP) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsAnomalousIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsAnomalousIP(val *BehaviorRuleSettingsAnomalousIP) *NullableBehaviorRuleSettingsAnomalousIP {
	return &NullableBehaviorRuleSettingsAnomalousIP{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsAnomalousIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsAnomalousIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

