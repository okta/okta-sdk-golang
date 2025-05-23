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

// AuthorizationServerPolicyRuleConditions struct for AuthorizationServerPolicyRuleConditions
type AuthorizationServerPolicyRuleConditions struct {
	GrantTypes *GrantTypePolicyRuleCondition `json:"grantTypes,omitempty"`
	People *AuthorizationServerPolicyPeopleCondition `json:"people,omitempty"`
	Scopes *OAuth2ScopesMediationPolicyRuleCondition `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleConditions AuthorizationServerPolicyRuleConditions

// NewAuthorizationServerPolicyRuleConditions instantiates a new AuthorizationServerPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleConditions() *AuthorizationServerPolicyRuleConditions {
	this := AuthorizationServerPolicyRuleConditions{}
	return &this
}

// NewAuthorizationServerPolicyRuleConditionsWithDefaults instantiates a new AuthorizationServerPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleConditionsWithDefaults() *AuthorizationServerPolicyRuleConditions {
	this := AuthorizationServerPolicyRuleConditions{}
	return &this
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleConditions) GetGrantTypes() GrantTypePolicyRuleCondition {
	if o == nil || o.GrantTypes == nil {
		var ret GrantTypePolicyRuleCondition
		return ret
	}
	return *o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleConditions) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given GrantTypePolicyRuleCondition and assigns it to the GrantTypes field.
func (o *AuthorizationServerPolicyRuleConditions) SetGrantTypes(v GrantTypePolicyRuleCondition) {
	o.GrantTypes = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleConditions) GetPeople() AuthorizationServerPolicyPeopleCondition {
	if o == nil || o.People == nil {
		var ret AuthorizationServerPolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleConditions) GetPeopleOk() (*AuthorizationServerPolicyPeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given AuthorizationServerPolicyPeopleCondition and assigns it to the People field.
func (o *AuthorizationServerPolicyRuleConditions) SetPeople(v AuthorizationServerPolicyPeopleCondition) {
	o.People = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition {
	if o == nil || o.Scopes == nil {
		var ret OAuth2ScopesMediationPolicyRuleCondition
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleConditions) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given OAuth2ScopesMediationPolicyRuleCondition and assigns it to the Scopes field.
func (o *AuthorizationServerPolicyRuleConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition) {
	o.Scopes = &v
}

func (o AuthorizationServerPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantTypes != nil {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRuleConditions) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyRuleConditions := _AuthorizationServerPolicyRuleConditions{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleConditions)
	if err == nil {
		*o = AuthorizationServerPolicyRuleConditions(varAuthorizationServerPolicyRuleConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "grantTypes")
		delete(additionalProperties, "people")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyRuleConditions struct {
	value *AuthorizationServerPolicyRuleConditions
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleConditions) Get() *AuthorizationServerPolicyRuleConditions {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleConditions) Set(val *AuthorizationServerPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleConditions(val *AuthorizationServerPolicyRuleConditions) *NullableAuthorizationServerPolicyRuleConditions {
	return &NullableAuthorizationServerPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

