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

// checks if the IdpDiscoveryPlatformPolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpDiscoveryPlatformPolicyRuleCondition{}

// IdpDiscoveryPlatformPolicyRuleCondition Specifies a particular platform or device to match on
type IdpDiscoveryPlatformPolicyRuleCondition struct {
	Exclude              []IdpDiscoveryPlatformConditionEvaluatorPlatform `json:"exclude,omitempty"`
	Include              []IdpDiscoveryPlatformConditionEvaluatorPlatform `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpDiscoveryPlatformPolicyRuleCondition IdpDiscoveryPlatformPolicyRuleCondition

// NewIdpDiscoveryPlatformPolicyRuleCondition instantiates a new IdpDiscoveryPlatformPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpDiscoveryPlatformPolicyRuleCondition() *IdpDiscoveryPlatformPolicyRuleCondition {
	this := IdpDiscoveryPlatformPolicyRuleCondition{}
	return &this
}

// NewIdpDiscoveryPlatformPolicyRuleConditionWithDefaults instantiates a new IdpDiscoveryPlatformPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpDiscoveryPlatformPolicyRuleConditionWithDefaults() *IdpDiscoveryPlatformPolicyRuleCondition {
	this := IdpDiscoveryPlatformPolicyRuleCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) GetExclude() []IdpDiscoveryPlatformConditionEvaluatorPlatform {
	if o == nil || IsNil(o.Exclude) {
		var ret []IdpDiscoveryPlatformConditionEvaluatorPlatform
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) GetExcludeOk() ([]IdpDiscoveryPlatformConditionEvaluatorPlatform, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []IdpDiscoveryPlatformConditionEvaluatorPlatform and assigns it to the Exclude field.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) SetExclude(v []IdpDiscoveryPlatformConditionEvaluatorPlatform) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) GetInclude() []IdpDiscoveryPlatformConditionEvaluatorPlatform {
	if o == nil || IsNil(o.Include) {
		var ret []IdpDiscoveryPlatformConditionEvaluatorPlatform
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) GetIncludeOk() ([]IdpDiscoveryPlatformConditionEvaluatorPlatform, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []IdpDiscoveryPlatformConditionEvaluatorPlatform and assigns it to the Include field.
func (o *IdpDiscoveryPlatformPolicyRuleCondition) SetInclude(v []IdpDiscoveryPlatformConditionEvaluatorPlatform) {
	o.Include = v
}

func (o IdpDiscoveryPlatformPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpDiscoveryPlatformPolicyRuleCondition) ToMap() (map[string]interface{}, error) {
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

func (o *IdpDiscoveryPlatformPolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varIdpDiscoveryPlatformPolicyRuleCondition := _IdpDiscoveryPlatformPolicyRuleCondition{}

	err = json.Unmarshal(data, &varIdpDiscoveryPlatformPolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = IdpDiscoveryPlatformPolicyRuleCondition(varIdpDiscoveryPlatformPolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdpDiscoveryPlatformPolicyRuleCondition struct {
	value *IdpDiscoveryPlatformPolicyRuleCondition
	isSet bool
}

func (v NullableIdpDiscoveryPlatformPolicyRuleCondition) Get() *IdpDiscoveryPlatformPolicyRuleCondition {
	return v.value
}

func (v *NullableIdpDiscoveryPlatformPolicyRuleCondition) Set(val *IdpDiscoveryPlatformPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpDiscoveryPlatformPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpDiscoveryPlatformPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpDiscoveryPlatformPolicyRuleCondition(val *IdpDiscoveryPlatformPolicyRuleCondition) *NullableIdpDiscoveryPlatformPolicyRuleCondition {
	return &NullableIdpDiscoveryPlatformPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableIdpDiscoveryPlatformPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpDiscoveryPlatformPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
