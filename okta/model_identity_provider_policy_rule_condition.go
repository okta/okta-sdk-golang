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

// IdentityProviderPolicyRuleCondition struct for IdentityProviderPolicyRuleCondition
type IdentityProviderPolicyRuleCondition struct {
	IdpIds []string `json:"idpIds,omitempty"`
	Provider *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderPolicyRuleCondition IdentityProviderPolicyRuleCondition

// NewIdentityProviderPolicyRuleCondition instantiates a new IdentityProviderPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderPolicyRuleCondition() *IdentityProviderPolicyRuleCondition {
	this := IdentityProviderPolicyRuleCondition{}
	return &this
}

// NewIdentityProviderPolicyRuleConditionWithDefaults instantiates a new IdentityProviderPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderPolicyRuleConditionWithDefaults() *IdentityProviderPolicyRuleCondition {
	this := IdentityProviderPolicyRuleCondition{}
	return &this
}

// GetIdpIds returns the IdpIds field value if set, zero value otherwise.
func (o *IdentityProviderPolicyRuleCondition) GetIdpIds() []string {
	if o == nil || o.IdpIds == nil {
		var ret []string
		return ret
	}
	return o.IdpIds
}

// GetIdpIdsOk returns a tuple with the IdpIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyRuleCondition) GetIdpIdsOk() ([]string, bool) {
	if o == nil || o.IdpIds == nil {
		return nil, false
	}
	return o.IdpIds, true
}

// HasIdpIds returns a boolean if a field has been set.
func (o *IdentityProviderPolicyRuleCondition) HasIdpIds() bool {
	if o != nil && o.IdpIds != nil {
		return true
	}

	return false
}

// SetIdpIds gets a reference to the given []string and assigns it to the IdpIds field.
func (o *IdentityProviderPolicyRuleCondition) SetIdpIds(v []string) {
	o.IdpIds = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *IdentityProviderPolicyRuleCondition) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicyRuleCondition) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *IdentityProviderPolicyRuleCondition) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *IdentityProviderPolicyRuleCondition) SetProvider(v string) {
	o.Provider = &v
}

func (o IdentityProviderPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdpIds != nil {
		toSerialize["idpIds"] = o.IdpIds
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderPolicyRuleCondition := _IdentityProviderPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varIdentityProviderPolicyRuleCondition)
	if err == nil {
		*o = IdentityProviderPolicyRuleCondition(varIdentityProviderPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "idpIds")
		delete(additionalProperties, "provider")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderPolicyRuleCondition struct {
	value *IdentityProviderPolicyRuleCondition
	isSet bool
}

func (v NullableIdentityProviderPolicyRuleCondition) Get() *IdentityProviderPolicyRuleCondition {
	return v.value
}

func (v *NullableIdentityProviderPolicyRuleCondition) Set(val *IdentityProviderPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderPolicyRuleCondition(val *IdentityProviderPolicyRuleCondition) *NullableIdentityProviderPolicyRuleCondition {
	return &NullableIdentityProviderPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableIdentityProviderPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

