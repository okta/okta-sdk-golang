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

// BehaviorRuleSettingsVelocity struct for BehaviorRuleSettingsVelocity
type BehaviorRuleSettingsVelocity struct {
	VelocityKph int32 `json:"velocityKph"`
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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["velocityKph"] = o.VelocityKph
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BehaviorRuleSettingsVelocity) UnmarshalJSON(bytes []byte) (err error) {
	varBehaviorRuleSettingsVelocity := _BehaviorRuleSettingsVelocity{}

	err = json.Unmarshal(bytes, &varBehaviorRuleSettingsVelocity)
	if err == nil {
		*o = BehaviorRuleSettingsVelocity(varBehaviorRuleSettingsVelocity)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "velocityKph")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

