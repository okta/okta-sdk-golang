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
	"fmt"
)

// checks if the BehaviorRuleSettingsVelocity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BehaviorRuleSettingsVelocity{}

// BehaviorRuleSettingsVelocity struct for BehaviorRuleSettingsVelocity
type BehaviorRuleSettingsVelocity struct {
	VelocityKph          int32 `json:"velocityKph"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleSettingsVelocity BehaviorRuleSettingsVelocity

// NewBehaviorRuleSettingsVelocity instantiates a new BehaviorRuleSettingsVelocity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsVelocity(velocityKph int32) *BehaviorRuleSettingsVelocity {
	this := BehaviorRuleSettingsVelocity{}
	this.VelocityKph = velocityKph
	return &this
}

// NewBehaviorRuleSettingsVelocityWithDefaults instantiates a new BehaviorRuleSettingsVelocity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsVelocityWithDefaults() *BehaviorRuleSettingsVelocity {
	this := BehaviorRuleSettingsVelocity{}
	var velocityKph int32 = 805
	this.VelocityKph = velocityKph
	return &this
}

// GetVelocityKph returns the VelocityKph field value
func (o *BehaviorRuleSettingsVelocity) GetVelocityKph() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VelocityKph
}

// GetVelocityKphOk returns a tuple with the VelocityKph field value
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsVelocity) GetVelocityKphOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VelocityKph, true
}

// SetVelocityKph sets field value
func (o *BehaviorRuleSettingsVelocity) SetVelocityKph(v int32) {
	o.VelocityKph = v
}

func (o BehaviorRuleSettingsVelocity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BehaviorRuleSettingsVelocity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["velocityKph"] = o.VelocityKph

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BehaviorRuleSettingsVelocity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"velocityKph",
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

	varBehaviorRuleSettingsVelocity := _BehaviorRuleSettingsVelocity{}

	err = json.Unmarshal(data, &varBehaviorRuleSettingsVelocity)

	if err != nil {
		return err
	}

	*o = BehaviorRuleSettingsVelocity(varBehaviorRuleSettingsVelocity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "velocityKph")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBehaviorRuleSettingsVelocity struct {
	value *BehaviorRuleSettingsVelocity
	isSet bool
}

func (v NullableBehaviorRuleSettingsVelocity) Get() *BehaviorRuleSettingsVelocity {
	return v.value
}

func (v *NullableBehaviorRuleSettingsVelocity) Set(val *BehaviorRuleSettingsVelocity) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsVelocity) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsVelocity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsVelocity(val *BehaviorRuleSettingsVelocity) *NullableBehaviorRuleSettingsVelocity {
	return &NullableBehaviorRuleSettingsVelocity{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsVelocity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsVelocity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
