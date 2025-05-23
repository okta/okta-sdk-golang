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

// BehaviorRuleSettingsAnomalousDevice struct for BehaviorRuleSettingsAnomalousDevice
type BehaviorRuleSettingsAnomalousDevice struct {
	MaxEventsUsedForEvaluation *int32 `json:"maxEventsUsedForEvaluation,omitempty"`
	MinEventsNeededForEvaluation *int32 `json:"minEventsNeededForEvaluation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleSettingsAnomalousDevice BehaviorRuleSettingsAnomalousDevice

// NewBehaviorRuleSettingsAnomalousDevice instantiates a new BehaviorRuleSettingsAnomalousDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsAnomalousDevice() *BehaviorRuleSettingsAnomalousDevice {
	this := BehaviorRuleSettingsAnomalousDevice{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// NewBehaviorRuleSettingsAnomalousDeviceWithDefaults instantiates a new BehaviorRuleSettingsAnomalousDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsAnomalousDeviceWithDefaults() *BehaviorRuleSettingsAnomalousDevice {
	this := BehaviorRuleSettingsAnomalousDevice{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousDevice) GetMaxEventsUsedForEvaluation() int32 {
	if o == nil || o.MaxEventsUsedForEvaluation == nil {
		var ret int32
		return ret
	}
	return *o.MaxEventsUsedForEvaluation
}

// GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousDevice) GetMaxEventsUsedForEvaluationOk() (*int32, bool) {
	if o == nil || o.MaxEventsUsedForEvaluation == nil {
		return nil, false
	}
	return o.MaxEventsUsedForEvaluation, true
}

// HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousDevice) HasMaxEventsUsedForEvaluation() bool {
	if o != nil && o.MaxEventsUsedForEvaluation != nil {
		return true
	}

	return false
}

// SetMaxEventsUsedForEvaluation gets a reference to the given int32 and assigns it to the MaxEventsUsedForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousDevice) SetMaxEventsUsedForEvaluation(v int32) {
	o.MaxEventsUsedForEvaluation = &v
}

// GetMinEventsNeededForEvaluation returns the MinEventsNeededForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousDevice) GetMinEventsNeededForEvaluation() int32 {
	if o == nil || o.MinEventsNeededForEvaluation == nil {
		var ret int32
		return ret
	}
	return *o.MinEventsNeededForEvaluation
}

// GetMinEventsNeededForEvaluationOk returns a tuple with the MinEventsNeededForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousDevice) GetMinEventsNeededForEvaluationOk() (*int32, bool) {
	if o == nil || o.MinEventsNeededForEvaluation == nil {
		return nil, false
	}
	return o.MinEventsNeededForEvaluation, true
}

// HasMinEventsNeededForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousDevice) HasMinEventsNeededForEvaluation() bool {
	if o != nil && o.MinEventsNeededForEvaluation != nil {
		return true
	}

	return false
}

// SetMinEventsNeededForEvaluation gets a reference to the given int32 and assigns it to the MinEventsNeededForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousDevice) SetMinEventsNeededForEvaluation(v int32) {
	o.MinEventsNeededForEvaluation = &v
}

func (o BehaviorRuleSettingsAnomalousDevice) MarshalJSON() ([]byte, error) {
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

func (o *BehaviorRuleSettingsAnomalousDevice) UnmarshalJSON(bytes []byte) (err error) {
	varBehaviorRuleSettingsAnomalousDevice := _BehaviorRuleSettingsAnomalousDevice{}

	err = json.Unmarshal(bytes, &varBehaviorRuleSettingsAnomalousDevice)
	if err == nil {
		*o = BehaviorRuleSettingsAnomalousDevice(varBehaviorRuleSettingsAnomalousDevice)
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

type NullableBehaviorRuleSettingsAnomalousDevice struct {
	value *BehaviorRuleSettingsAnomalousDevice
	isSet bool
}

func (v NullableBehaviorRuleSettingsAnomalousDevice) Get() *BehaviorRuleSettingsAnomalousDevice {
	return v.value
}

func (v *NullableBehaviorRuleSettingsAnomalousDevice) Set(val *BehaviorRuleSettingsAnomalousDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsAnomalousDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsAnomalousDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsAnomalousDevice(val *BehaviorRuleSettingsAnomalousDevice) *NullableBehaviorRuleSettingsAnomalousDevice {
	return &NullableBehaviorRuleSettingsAnomalousDevice{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsAnomalousDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsAnomalousDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

