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

// AppInstancePolicyRuleCondition struct for AppInstancePolicyRuleCondition
type AppInstancePolicyRuleCondition struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppInstancePolicyRuleCondition AppInstancePolicyRuleCondition

// NewAppInstancePolicyRuleCondition instantiates a new AppInstancePolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInstancePolicyRuleCondition() *AppInstancePolicyRuleCondition {
	this := AppInstancePolicyRuleCondition{}
	return &this
}

// NewAppInstancePolicyRuleConditionWithDefaults instantiates a new AppInstancePolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInstancePolicyRuleConditionWithDefaults() *AppInstancePolicyRuleCondition {
	this := AppInstancePolicyRuleCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *AppInstancePolicyRuleCondition) GetExclude() []string {
	if o == nil || o.Exclude == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInstancePolicyRuleCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *AppInstancePolicyRuleCondition) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *AppInstancePolicyRuleCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *AppInstancePolicyRuleCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInstancePolicyRuleCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *AppInstancePolicyRuleCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *AppInstancePolicyRuleCondition) SetInclude(v []string) {
	o.Include = v
}

func (o AppInstancePolicyRuleCondition) MarshalJSON() ([]byte, error) {
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

func (o *AppInstancePolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varAppInstancePolicyRuleCondition := _AppInstancePolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varAppInstancePolicyRuleCondition)
	if err == nil {
		*o = AppInstancePolicyRuleCondition(varAppInstancePolicyRuleCondition)
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

type NullableAppInstancePolicyRuleCondition struct {
	value *AppInstancePolicyRuleCondition
	isSet bool
}

func (v NullableAppInstancePolicyRuleCondition) Get() *AppInstancePolicyRuleCondition {
	return v.value
}

func (v *NullableAppInstancePolicyRuleCondition) Set(val *AppInstancePolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInstancePolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInstancePolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInstancePolicyRuleCondition(val *AppInstancePolicyRuleCondition) *NullableAppInstancePolicyRuleCondition {
	return &NullableAppInstancePolicyRuleCondition{value: val, isSet: true}
}

func (v NullableAppInstancePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInstancePolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

