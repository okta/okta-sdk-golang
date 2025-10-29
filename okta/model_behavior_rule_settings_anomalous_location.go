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
	"fmt"
)

// checks if the BehaviorRuleSettingsAnomalousLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BehaviorRuleSettingsAnomalousLocation{}

// BehaviorRuleSettingsAnomalousLocation struct for BehaviorRuleSettingsAnomalousLocation
type BehaviorRuleSettingsAnomalousLocation struct {
	MaxEventsUsedForEvaluation   *int32 `json:"maxEventsUsedForEvaluation,omitempty"`
	MinEventsNeededForEvaluation *int32 `json:"minEventsNeededForEvaluation,omitempty"`
	Granularity                  string `json:"granularity"`
	// Required when `granularity` is `LAT_LONG`. Radius from the provided coordinates in kilometers.
	RadiusKilometers     *int32 `json:"radiusKilometers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleSettingsAnomalousLocation BehaviorRuleSettingsAnomalousLocation

// NewBehaviorRuleSettingsAnomalousLocation instantiates a new BehaviorRuleSettingsAnomalousLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsAnomalousLocation(granularity string) *BehaviorRuleSettingsAnomalousLocation {
	this := BehaviorRuleSettingsAnomalousLocation{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	this.Granularity = granularity
	return &this
}

// NewBehaviorRuleSettingsAnomalousLocationWithDefaults instantiates a new BehaviorRuleSettingsAnomalousLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsAnomalousLocationWithDefaults() *BehaviorRuleSettingsAnomalousLocation {
	this := BehaviorRuleSettingsAnomalousLocation{}
	var maxEventsUsedForEvaluation int32 = 20
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	var minEventsNeededForEvaluation int32 = 0
	this.MinEventsNeededForEvaluation = &minEventsNeededForEvaluation
	return &this
}

// GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousLocation) GetMaxEventsUsedForEvaluation() int32 {
	if o == nil || IsNil(o.MaxEventsUsedForEvaluation) {
		var ret int32
		return ret
	}
	return *o.MaxEventsUsedForEvaluation
}

// GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) GetMaxEventsUsedForEvaluationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxEventsUsedForEvaluation) {
		return nil, false
	}
	return o.MaxEventsUsedForEvaluation, true
}

// HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) HasMaxEventsUsedForEvaluation() bool {
	if o != nil && !IsNil(o.MaxEventsUsedForEvaluation) {
		return true
	}

	return false
}

// SetMaxEventsUsedForEvaluation gets a reference to the given int32 and assigns it to the MaxEventsUsedForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousLocation) SetMaxEventsUsedForEvaluation(v int32) {
	o.MaxEventsUsedForEvaluation = &v
}

// GetMinEventsNeededForEvaluation returns the MinEventsNeededForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousLocation) GetMinEventsNeededForEvaluation() int32 {
	if o == nil || IsNil(o.MinEventsNeededForEvaluation) {
		var ret int32
		return ret
	}
	return *o.MinEventsNeededForEvaluation
}

// GetMinEventsNeededForEvaluationOk returns a tuple with the MinEventsNeededForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) GetMinEventsNeededForEvaluationOk() (*int32, bool) {
	if o == nil || IsNil(o.MinEventsNeededForEvaluation) {
		return nil, false
	}
	return o.MinEventsNeededForEvaluation, true
}

// HasMinEventsNeededForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) HasMinEventsNeededForEvaluation() bool {
	if o != nil && !IsNil(o.MinEventsNeededForEvaluation) {
		return true
	}

	return false
}

// SetMinEventsNeededForEvaluation gets a reference to the given int32 and assigns it to the MinEventsNeededForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousLocation) SetMinEventsNeededForEvaluation(v int32) {
	o.MinEventsNeededForEvaluation = &v
}

// GetGranularity returns the Granularity field value
func (o *BehaviorRuleSettingsAnomalousLocation) GetGranularity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) GetGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *BehaviorRuleSettingsAnomalousLocation) SetGranularity(v string) {
	o.Granularity = v
}

// GetRadiusKilometers returns the RadiusKilometers field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousLocation) GetRadiusKilometers() int32 {
	if o == nil || IsNil(o.RadiusKilometers) {
		var ret int32
		return ret
	}
	return *o.RadiusKilometers
}

// GetRadiusKilometersOk returns a tuple with the RadiusKilometers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) GetRadiusKilometersOk() (*int32, bool) {
	if o == nil || IsNil(o.RadiusKilometers) {
		return nil, false
	}
	return o.RadiusKilometers, true
}

// HasRadiusKilometers returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousLocation) HasRadiusKilometers() bool {
	if o != nil && !IsNil(o.RadiusKilometers) {
		return true
	}

	return false
}

// SetRadiusKilometers gets a reference to the given int32 and assigns it to the RadiusKilometers field.
func (o *BehaviorRuleSettingsAnomalousLocation) SetRadiusKilometers(v int32) {
	o.RadiusKilometers = &v
}

func (o BehaviorRuleSettingsAnomalousLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BehaviorRuleSettingsAnomalousLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxEventsUsedForEvaluation) {
		toSerialize["maxEventsUsedForEvaluation"] = o.MaxEventsUsedForEvaluation
	}
	if !IsNil(o.MinEventsNeededForEvaluation) {
		toSerialize["minEventsNeededForEvaluation"] = o.MinEventsNeededForEvaluation
	}
	toSerialize["granularity"] = o.Granularity
	if !IsNil(o.RadiusKilometers) {
		toSerialize["radiusKilometers"] = o.RadiusKilometers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BehaviorRuleSettingsAnomalousLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"granularity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBehaviorRuleSettingsAnomalousLocation := _BehaviorRuleSettingsAnomalousLocation{}

	err = json.Unmarshal(data, &varBehaviorRuleSettingsAnomalousLocation)

	if err != nil {
		return err
	}

	*o = BehaviorRuleSettingsAnomalousLocation(varBehaviorRuleSettingsAnomalousLocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxEventsUsedForEvaluation")
		delete(additionalProperties, "minEventsNeededForEvaluation")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "radiusKilometers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBehaviorRuleSettingsAnomalousLocation struct {
	value *BehaviorRuleSettingsAnomalousLocation
	isSet bool
}

func (v NullableBehaviorRuleSettingsAnomalousLocation) Get() *BehaviorRuleSettingsAnomalousLocation {
	return v.value
}

func (v *NullableBehaviorRuleSettingsAnomalousLocation) Set(val *BehaviorRuleSettingsAnomalousLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsAnomalousLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsAnomalousLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsAnomalousLocation(val *BehaviorRuleSettingsAnomalousLocation) *NullableBehaviorRuleSettingsAnomalousLocation {
	return &NullableBehaviorRuleSettingsAnomalousLocation{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsAnomalousLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsAnomalousLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
