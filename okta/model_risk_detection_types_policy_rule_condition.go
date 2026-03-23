/*
Okta Admin Management API

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

// checks if the RiskDetectionTypesPolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskDetectionTypesPolicyRuleCondition{}

// RiskDetectionTypesPolicyRuleCondition <x-lifecycle class=\"oie\"></x-lifecycle> An object that references detected risk events. This object can have an `include` parameter or an `exclude` parameter, but not both.
type RiskDetectionTypesPolicyRuleCondition struct {
	// An array of detected risk events to exclude in the entity policy rule
	Exclude []string `json:"exclude,omitempty"`
	// An array of detected risk events to include in the entity policy rule
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskDetectionTypesPolicyRuleCondition RiskDetectionTypesPolicyRuleCondition

// NewRiskDetectionTypesPolicyRuleCondition instantiates a new RiskDetectionTypesPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskDetectionTypesPolicyRuleCondition() *RiskDetectionTypesPolicyRuleCondition {
	this := RiskDetectionTypesPolicyRuleCondition{}
	return &this
}

// NewRiskDetectionTypesPolicyRuleConditionWithDefaults instantiates a new RiskDetectionTypesPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskDetectionTypesPolicyRuleConditionWithDefaults() *RiskDetectionTypesPolicyRuleCondition {
	this := RiskDetectionTypesPolicyRuleCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *RiskDetectionTypesPolicyRuleCondition) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskDetectionTypesPolicyRuleCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *RiskDetectionTypesPolicyRuleCondition) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *RiskDetectionTypesPolicyRuleCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *RiskDetectionTypesPolicyRuleCondition) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskDetectionTypesPolicyRuleCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *RiskDetectionTypesPolicyRuleCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *RiskDetectionTypesPolicyRuleCondition) SetInclude(v []string) {
	o.Include = v
}

func (o RiskDetectionTypesPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskDetectionTypesPolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskDetectionTypesPolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varRiskDetectionTypesPolicyRuleCondition := _RiskDetectionTypesPolicyRuleCondition{}

	err = json.Unmarshal(data, &varRiskDetectionTypesPolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = RiskDetectionTypesPolicyRuleCondition(varRiskDetectionTypesPolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskDetectionTypesPolicyRuleCondition struct {
	value *RiskDetectionTypesPolicyRuleCondition
	isSet bool
}

func (v NullableRiskDetectionTypesPolicyRuleCondition) Get() *RiskDetectionTypesPolicyRuleCondition {
	return v.value
}

func (v *NullableRiskDetectionTypesPolicyRuleCondition) Set(val *RiskDetectionTypesPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskDetectionTypesPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskDetectionTypesPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskDetectionTypesPolicyRuleCondition(val *RiskDetectionTypesPolicyRuleCondition) *NullableRiskDetectionTypesPolicyRuleCondition {
	return &NullableRiskDetectionTypesPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableRiskDetectionTypesPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskDetectionTypesPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
