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

// AuthorizationServerPolicyRuleUserCondition Specifies a set of Users to be included
type AuthorizationServerPolicyRuleUserCondition struct {
	// Users to be included
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleUserCondition AuthorizationServerPolicyRuleUserCondition

// NewAuthorizationServerPolicyRuleUserCondition instantiates a new AuthorizationServerPolicyRuleUserCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleUserCondition() *AuthorizationServerPolicyRuleUserCondition {
	this := AuthorizationServerPolicyRuleUserCondition{}
	return &this
}

// NewAuthorizationServerPolicyRuleUserConditionWithDefaults instantiates a new AuthorizationServerPolicyRuleUserCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleUserConditionWithDefaults() *AuthorizationServerPolicyRuleUserCondition {
	this := AuthorizationServerPolicyRuleUserCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleUserCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleUserCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleUserCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *AuthorizationServerPolicyRuleUserCondition) SetInclude(v []string) {
	o.Include = v
}

func (o AuthorizationServerPolicyRuleUserCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRuleUserCondition) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyRuleUserCondition := _AuthorizationServerPolicyRuleUserCondition{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleUserCondition)
	if err == nil {
		*o = AuthorizationServerPolicyRuleUserCondition(varAuthorizationServerPolicyRuleUserCondition)
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

type NullableAuthorizationServerPolicyRuleUserCondition struct {
	value *AuthorizationServerPolicyRuleUserCondition
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleUserCondition) Get() *AuthorizationServerPolicyRuleUserCondition {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleUserCondition) Set(val *AuthorizationServerPolicyRuleUserCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleUserCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleUserCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleUserCondition(val *AuthorizationServerPolicyRuleUserCondition) *NullableAuthorizationServerPolicyRuleUserCondition {
	return &NullableAuthorizationServerPolicyRuleUserCondition{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleUserCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleUserCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

