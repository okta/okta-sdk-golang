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

// OAuth2ScopesMediationPolicyRuleCondition Array of scopes that the condition includes
type OAuth2ScopesMediationPolicyRuleCondition struct {
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ScopesMediationPolicyRuleCondition OAuth2ScopesMediationPolicyRuleCondition

// NewOAuth2ScopesMediationPolicyRuleCondition instantiates a new OAuth2ScopesMediationPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ScopesMediationPolicyRuleCondition() *OAuth2ScopesMediationPolicyRuleCondition {
	this := OAuth2ScopesMediationPolicyRuleCondition{}
	return &this
}

// NewOAuth2ScopesMediationPolicyRuleConditionWithDefaults instantiates a new OAuth2ScopesMediationPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ScopesMediationPolicyRuleConditionWithDefaults() *OAuth2ScopesMediationPolicyRuleCondition {
	this := OAuth2ScopesMediationPolicyRuleCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *OAuth2ScopesMediationPolicyRuleCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopesMediationPolicyRuleCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *OAuth2ScopesMediationPolicyRuleCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *OAuth2ScopesMediationPolicyRuleCondition) SetInclude(v []string) {
	o.Include = v
}

func (o OAuth2ScopesMediationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ScopesMediationPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ScopesMediationPolicyRuleCondition := _OAuth2ScopesMediationPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varOAuth2ScopesMediationPolicyRuleCondition)
	if err == nil {
		*o = OAuth2ScopesMediationPolicyRuleCondition(varOAuth2ScopesMediationPolicyRuleCondition)
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

type NullableOAuth2ScopesMediationPolicyRuleCondition struct {
	value *OAuth2ScopesMediationPolicyRuleCondition
	isSet bool
}

func (v NullableOAuth2ScopesMediationPolicyRuleCondition) Get() *OAuth2ScopesMediationPolicyRuleCondition {
	return v.value
}

func (v *NullableOAuth2ScopesMediationPolicyRuleCondition) Set(val *OAuth2ScopesMediationPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopesMediationPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopesMediationPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopesMediationPolicyRuleCondition(val *OAuth2ScopesMediationPolicyRuleCondition) *NullableOAuth2ScopesMediationPolicyRuleCondition {
	return &NullableOAuth2ScopesMediationPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableOAuth2ScopesMediationPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopesMediationPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

