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

// PlatformPolicyRuleCondition struct for PlatformPolicyRuleCondition
type PlatformPolicyRuleCondition struct {
	Exclude []PlatformConditionEvaluatorPlatform `json:"exclude,omitempty"`
	Include []PlatformConditionEvaluatorPlatform `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformPolicyRuleCondition PlatformPolicyRuleCondition

// NewPlatformPolicyRuleCondition instantiates a new PlatformPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformPolicyRuleCondition() *PlatformPolicyRuleCondition {
	this := PlatformPolicyRuleCondition{}
	return &this
}

// NewPlatformPolicyRuleConditionWithDefaults instantiates a new PlatformPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformPolicyRuleConditionWithDefaults() *PlatformPolicyRuleCondition {
	this := PlatformPolicyRuleCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *PlatformPolicyRuleCondition) GetExclude() []PlatformConditionEvaluatorPlatform {
	if o == nil || o.Exclude == nil {
		var ret []PlatformConditionEvaluatorPlatform
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPolicyRuleCondition) GetExcludeOk() ([]PlatformConditionEvaluatorPlatform, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *PlatformPolicyRuleCondition) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []PlatformConditionEvaluatorPlatform and assigns it to the Exclude field.
func (o *PlatformPolicyRuleCondition) SetExclude(v []PlatformConditionEvaluatorPlatform) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *PlatformPolicyRuleCondition) GetInclude() []PlatformConditionEvaluatorPlatform {
	if o == nil || o.Include == nil {
		var ret []PlatformConditionEvaluatorPlatform
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPolicyRuleCondition) GetIncludeOk() ([]PlatformConditionEvaluatorPlatform, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *PlatformPolicyRuleCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []PlatformConditionEvaluatorPlatform and assigns it to the Include field.
func (o *PlatformPolicyRuleCondition) SetInclude(v []PlatformConditionEvaluatorPlatform) {
	o.Include = v
}

func (o PlatformPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformPolicyRuleCondition := _PlatformPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varPlatformPolicyRuleCondition)
	if err == nil {
		*o = PlatformPolicyRuleCondition(varPlatformPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePlatformPolicyRuleCondition struct {
	value *PlatformPolicyRuleCondition
	isSet bool
}

func (v NullablePlatformPolicyRuleCondition) Get() *PlatformPolicyRuleCondition {
	return v.value
}

func (v *NullablePlatformPolicyRuleCondition) Set(val *PlatformPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformPolicyRuleCondition(val *PlatformPolicyRuleCondition) *NullablePlatformPolicyRuleCondition {
	return &NullablePlatformPolicyRuleCondition{value: val, isSet: true}
}

func (v NullablePlatformPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

