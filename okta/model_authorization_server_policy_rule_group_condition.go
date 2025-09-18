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
)

// checks if the AuthorizationServerPolicyRuleGroupCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationServerPolicyRuleGroupCondition{}

// AuthorizationServerPolicyRuleGroupCondition Specifies a set of Groups whose Users are to be included
type AuthorizationServerPolicyRuleGroupCondition struct {
	// Groups to be included
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleGroupCondition AuthorizationServerPolicyRuleGroupCondition

// NewAuthorizationServerPolicyRuleGroupCondition instantiates a new AuthorizationServerPolicyRuleGroupCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleGroupCondition() *AuthorizationServerPolicyRuleGroupCondition {
	this := AuthorizationServerPolicyRuleGroupCondition{}
	return &this
}

// NewAuthorizationServerPolicyRuleGroupConditionWithDefaults instantiates a new AuthorizationServerPolicyRuleGroupCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleGroupConditionWithDefaults() *AuthorizationServerPolicyRuleGroupCondition {
	this := AuthorizationServerPolicyRuleGroupCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleGroupCondition) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleGroupCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleGroupCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *AuthorizationServerPolicyRuleGroupCondition) SetInclude(v []string) {
	o.Include = v
}

func (o AuthorizationServerPolicyRuleGroupCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationServerPolicyRuleGroupCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizationServerPolicyRuleGroupCondition) UnmarshalJSON(data []byte) (err error) {
	varAuthorizationServerPolicyRuleGroupCondition := _AuthorizationServerPolicyRuleGroupCondition{}

	err = json.Unmarshal(data, &varAuthorizationServerPolicyRuleGroupCondition)

	if err != nil {
		return err
	}

	*o = AuthorizationServerPolicyRuleGroupCondition(varAuthorizationServerPolicyRuleGroupCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationServerPolicyRuleGroupCondition struct {
	value *AuthorizationServerPolicyRuleGroupCondition
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleGroupCondition) Get() *AuthorizationServerPolicyRuleGroupCondition {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleGroupCondition) Set(val *AuthorizationServerPolicyRuleGroupCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleGroupCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleGroupCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleGroupCondition(val *AuthorizationServerPolicyRuleGroupCondition) *NullableAuthorizationServerPolicyRuleGroupCondition {
	return &NullableAuthorizationServerPolicyRuleGroupCondition{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleGroupCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleGroupCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
