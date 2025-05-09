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

// GrantTypePolicyRuleCondition Array of grant types that this condition includes. Determines the mechanism that Okta uses to authorize the creation of the tokens.
type GrantTypePolicyRuleCondition struct {
	// Array of grant types that this condition includes.
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantTypePolicyRuleCondition GrantTypePolicyRuleCondition

// NewGrantTypePolicyRuleCondition instantiates a new GrantTypePolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantTypePolicyRuleCondition() *GrantTypePolicyRuleCondition {
	this := GrantTypePolicyRuleCondition{}
	return &this
}

// NewGrantTypePolicyRuleConditionWithDefaults instantiates a new GrantTypePolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantTypePolicyRuleConditionWithDefaults() *GrantTypePolicyRuleCondition {
	this := GrantTypePolicyRuleCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *GrantTypePolicyRuleCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyRuleCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *GrantTypePolicyRuleCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *GrantTypePolicyRuleCondition) SetInclude(v []string) {
	o.Include = v
}

func (o GrantTypePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GrantTypePolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varGrantTypePolicyRuleCondition := _GrantTypePolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varGrantTypePolicyRuleCondition)
	if err == nil {
		*o = GrantTypePolicyRuleCondition(varGrantTypePolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGrantTypePolicyRuleCondition struct {
	value *GrantTypePolicyRuleCondition
	isSet bool
}

func (v NullableGrantTypePolicyRuleCondition) Get() *GrantTypePolicyRuleCondition {
	return v.value
}

func (v *NullableGrantTypePolicyRuleCondition) Set(val *GrantTypePolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantTypePolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantTypePolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantTypePolicyRuleCondition(val *GrantTypePolicyRuleCondition) *NullableGrantTypePolicyRuleCondition {
	return &NullableGrantTypePolicyRuleCondition{value: val, isSet: true}
}

func (v NullableGrantTypePolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantTypePolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

