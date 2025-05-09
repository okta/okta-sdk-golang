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

// TokenAuthorizationServerPolicyRuleActionInlineHook struct for TokenAuthorizationServerPolicyRuleActionInlineHook
type TokenAuthorizationServerPolicyRuleActionInlineHook struct {
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenAuthorizationServerPolicyRuleActionInlineHook TokenAuthorizationServerPolicyRuleActionInlineHook

// NewTokenAuthorizationServerPolicyRuleActionInlineHook instantiates a new TokenAuthorizationServerPolicyRuleActionInlineHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAuthorizationServerPolicyRuleActionInlineHook() *TokenAuthorizationServerPolicyRuleActionInlineHook {
	this := TokenAuthorizationServerPolicyRuleActionInlineHook{}
	return &this
}

// NewTokenAuthorizationServerPolicyRuleActionInlineHookWithDefaults instantiates a new TokenAuthorizationServerPolicyRuleActionInlineHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAuthorizationServerPolicyRuleActionInlineHookWithDefaults() *TokenAuthorizationServerPolicyRuleActionInlineHook {
	this := TokenAuthorizationServerPolicyRuleActionInlineHook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenAuthorizationServerPolicyRuleActionInlineHook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationServerPolicyRuleActionInlineHook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenAuthorizationServerPolicyRuleActionInlineHook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenAuthorizationServerPolicyRuleActionInlineHook) SetId(v string) {
	o.Id = &v
}

func (o TokenAuthorizationServerPolicyRuleActionInlineHook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenAuthorizationServerPolicyRuleActionInlineHook) UnmarshalJSON(bytes []byte) (err error) {
	varTokenAuthorizationServerPolicyRuleActionInlineHook := _TokenAuthorizationServerPolicyRuleActionInlineHook{}

	err = json.Unmarshal(bytes, &varTokenAuthorizationServerPolicyRuleActionInlineHook)
	if err == nil {
		*o = TokenAuthorizationServerPolicyRuleActionInlineHook(varTokenAuthorizationServerPolicyRuleActionInlineHook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenAuthorizationServerPolicyRuleActionInlineHook struct {
	value *TokenAuthorizationServerPolicyRuleActionInlineHook
	isSet bool
}

func (v NullableTokenAuthorizationServerPolicyRuleActionInlineHook) Get() *TokenAuthorizationServerPolicyRuleActionInlineHook {
	return v.value
}

func (v *NullableTokenAuthorizationServerPolicyRuleActionInlineHook) Set(val *TokenAuthorizationServerPolicyRuleActionInlineHook) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAuthorizationServerPolicyRuleActionInlineHook) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAuthorizationServerPolicyRuleActionInlineHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAuthorizationServerPolicyRuleActionInlineHook(val *TokenAuthorizationServerPolicyRuleActionInlineHook) *NullableTokenAuthorizationServerPolicyRuleActionInlineHook {
	return &NullableTokenAuthorizationServerPolicyRuleActionInlineHook{value: val, isSet: true}
}

func (v NullableTokenAuthorizationServerPolicyRuleActionInlineHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAuthorizationServerPolicyRuleActionInlineHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

