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

// checks if the TokenPayLoadDataContextAllOfPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPayLoadDataContextAllOfPolicy{}

// TokenPayLoadDataContextAllOfPolicy The authorization server policy used to mint the token
type TokenPayLoadDataContextAllOfPolicy struct {
	// The unique identifier for the policy
	Id                   *string                                 `json:"id,omitempty"`
	Rule                 *TokenPayLoadDataContextAllOfPolicyRule `json:"rule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContextAllOfPolicy TokenPayLoadDataContextAllOfPolicy

// NewTokenPayLoadDataContextAllOfPolicy instantiates a new TokenPayLoadDataContextAllOfPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContextAllOfPolicy() *TokenPayLoadDataContextAllOfPolicy {
	this := TokenPayLoadDataContextAllOfPolicy{}
	return &this
}

// NewTokenPayLoadDataContextAllOfPolicyWithDefaults instantiates a new TokenPayLoadDataContextAllOfPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextAllOfPolicyWithDefaults() *TokenPayLoadDataContextAllOfPolicy {
	this := TokenPayLoadDataContextAllOfPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenPayLoadDataContextAllOfPolicy) SetId(v string) {
	o.Id = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfPolicy) GetRule() TokenPayLoadDataContextAllOfPolicyRule {
	if o == nil || IsNil(o.Rule) {
		var ret TokenPayLoadDataContextAllOfPolicyRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfPolicy) GetRuleOk() (*TokenPayLoadDataContextAllOfPolicyRule, bool) {
	if o == nil || IsNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfPolicy) HasRule() bool {
	if o != nil && !IsNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given TokenPayLoadDataContextAllOfPolicyRule and assigns it to the Rule field.
func (o *TokenPayLoadDataContextAllOfPolicy) SetRule(v TokenPayLoadDataContextAllOfPolicyRule) {
	o.Rule = &v
}

func (o TokenPayLoadDataContextAllOfPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPayLoadDataContextAllOfPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Rule) {
		toSerialize["rule"] = o.Rule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenPayLoadDataContextAllOfPolicy) UnmarshalJSON(data []byte) (err error) {
	varTokenPayLoadDataContextAllOfPolicy := _TokenPayLoadDataContextAllOfPolicy{}

	err = json.Unmarshal(data, &varTokenPayLoadDataContextAllOfPolicy)

	if err != nil {
		return err
	}

	*o = TokenPayLoadDataContextAllOfPolicy(varTokenPayLoadDataContextAllOfPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "rule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenPayLoadDataContextAllOfPolicy struct {
	value *TokenPayLoadDataContextAllOfPolicy
	isSet bool
}

func (v NullableTokenPayLoadDataContextAllOfPolicy) Get() *TokenPayLoadDataContextAllOfPolicy {
	return v.value
}

func (v *NullableTokenPayLoadDataContextAllOfPolicy) Set(val *TokenPayLoadDataContextAllOfPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContextAllOfPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContextAllOfPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContextAllOfPolicy(val *TokenPayLoadDataContextAllOfPolicy) *NullableTokenPayLoadDataContextAllOfPolicy {
	return &NullableTokenPayLoadDataContextAllOfPolicy{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContextAllOfPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContextAllOfPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
