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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// TokenPayLoadDataContextAllOfPolicyRule The authorization server policy rule used to mint the token
type TokenPayLoadDataContextAllOfPolicyRule struct {
	// The unique identifier for the policy rule
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContextAllOfPolicyRule TokenPayLoadDataContextAllOfPolicyRule

// NewTokenPayLoadDataContextAllOfPolicyRule instantiates a new TokenPayLoadDataContextAllOfPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContextAllOfPolicyRule() *TokenPayLoadDataContextAllOfPolicyRule {
	this := TokenPayLoadDataContextAllOfPolicyRule{}
	return &this
}

// NewTokenPayLoadDataContextAllOfPolicyRuleWithDefaults instantiates a new TokenPayLoadDataContextAllOfPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextAllOfPolicyRuleWithDefaults() *TokenPayLoadDataContextAllOfPolicyRule {
	this := TokenPayLoadDataContextAllOfPolicyRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfPolicyRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfPolicyRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfPolicyRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenPayLoadDataContextAllOfPolicyRule) SetId(v string) {
	o.Id = &v
}

func (o TokenPayLoadDataContextAllOfPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenPayLoadDataContextAllOfPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	varTokenPayLoadDataContextAllOfPolicyRule := _TokenPayLoadDataContextAllOfPolicyRule{}

	err = json.Unmarshal(bytes, &varTokenPayLoadDataContextAllOfPolicyRule)
	if err == nil {
		*o = TokenPayLoadDataContextAllOfPolicyRule(varTokenPayLoadDataContextAllOfPolicyRule)
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

type NullableTokenPayLoadDataContextAllOfPolicyRule struct {
	value *TokenPayLoadDataContextAllOfPolicyRule
	isSet bool
}

func (v NullableTokenPayLoadDataContextAllOfPolicyRule) Get() *TokenPayLoadDataContextAllOfPolicyRule {
	return v.value
}

func (v *NullableTokenPayLoadDataContextAllOfPolicyRule) Set(val *TokenPayLoadDataContextAllOfPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContextAllOfPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContextAllOfPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContextAllOfPolicyRule(val *TokenPayLoadDataContextAllOfPolicyRule) *NullableTokenPayLoadDataContextAllOfPolicyRule {
	return &NullableTokenPayLoadDataContextAllOfPolicyRule{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContextAllOfPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContextAllOfPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

