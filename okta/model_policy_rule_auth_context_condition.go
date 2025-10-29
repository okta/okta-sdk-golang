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
)

// checks if the PolicyRuleAuthContextCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyRuleAuthContextCondition{}

// PolicyRuleAuthContextCondition Specifies an authentication entry point
type PolicyRuleAuthContextCondition struct {
	// Specifies how the user is authenticated
	AuthType             *string `json:"authType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyRuleAuthContextCondition PolicyRuleAuthContextCondition

// NewPolicyRuleAuthContextCondition instantiates a new PolicyRuleAuthContextCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRuleAuthContextCondition() *PolicyRuleAuthContextCondition {
	this := PolicyRuleAuthContextCondition{}
	return &this
}

// NewPolicyRuleAuthContextConditionWithDefaults instantiates a new PolicyRuleAuthContextCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRuleAuthContextConditionWithDefaults() *PolicyRuleAuthContextCondition {
	this := PolicyRuleAuthContextCondition{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *PolicyRuleAuthContextCondition) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRuleAuthContextCondition) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *PolicyRuleAuthContextCondition) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *PolicyRuleAuthContextCondition) SetAuthType(v string) {
	o.AuthType = &v
}

func (o PolicyRuleAuthContextCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRuleAuthContextCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyRuleAuthContextCondition) UnmarshalJSON(data []byte) (err error) {
	varPolicyRuleAuthContextCondition := _PolicyRuleAuthContextCondition{}

	err = json.Unmarshal(data, &varPolicyRuleAuthContextCondition)

	if err != nil {
		return err
	}

	*o = PolicyRuleAuthContextCondition(varPolicyRuleAuthContextCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyRuleAuthContextCondition struct {
	value *PolicyRuleAuthContextCondition
	isSet bool
}

func (v NullablePolicyRuleAuthContextCondition) Get() *PolicyRuleAuthContextCondition {
	return v.value
}

func (v *NullablePolicyRuleAuthContextCondition) Set(val *PolicyRuleAuthContextCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRuleAuthContextCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRuleAuthContextCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRuleAuthContextCondition(val *PolicyRuleAuthContextCondition) *NullablePolicyRuleAuthContextCondition {
	return &NullablePolicyRuleAuthContextCondition{value: val, isSet: true}
}

func (v NullablePolicyRuleAuthContextCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRuleAuthContextCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
