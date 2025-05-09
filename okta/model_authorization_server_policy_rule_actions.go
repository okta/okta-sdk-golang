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

// AuthorizationServerPolicyRuleActions struct for AuthorizationServerPolicyRuleActions
type AuthorizationServerPolicyRuleActions struct {
	Token *TokenAuthorizationServerPolicyRuleAction `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleActions AuthorizationServerPolicyRuleActions

// NewAuthorizationServerPolicyRuleActions instantiates a new AuthorizationServerPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleActions() *AuthorizationServerPolicyRuleActions {
	this := AuthorizationServerPolicyRuleActions{}
	return &this
}

// NewAuthorizationServerPolicyRuleActionsWithDefaults instantiates a new AuthorizationServerPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleActionsWithDefaults() *AuthorizationServerPolicyRuleActions {
	this := AuthorizationServerPolicyRuleActions{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActions) GetToken() TokenAuthorizationServerPolicyRuleAction {
	if o == nil || o.Token == nil {
		var ret TokenAuthorizationServerPolicyRuleAction
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActions) GetTokenOk() (*TokenAuthorizationServerPolicyRuleAction, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActions) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given TokenAuthorizationServerPolicyRuleAction and assigns it to the Token field.
func (o *AuthorizationServerPolicyRuleActions) SetToken(v TokenAuthorizationServerPolicyRuleAction) {
	o.Token = &v
}

func (o AuthorizationServerPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRuleActions) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyRuleActions := _AuthorizationServerPolicyRuleActions{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleActions)
	if err == nil {
		*o = AuthorizationServerPolicyRuleActions(varAuthorizationServerPolicyRuleActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyRuleActions struct {
	value *AuthorizationServerPolicyRuleActions
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleActions) Get() *AuthorizationServerPolicyRuleActions {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleActions) Set(val *AuthorizationServerPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleActions(val *AuthorizationServerPolicyRuleActions) *NullableAuthorizationServerPolicyRuleActions {
	return &NullableAuthorizationServerPolicyRuleActions{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

