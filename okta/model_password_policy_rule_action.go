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

// PasswordPolicyRuleAction struct for PasswordPolicyRuleAction
type PasswordPolicyRuleAction struct {
	Access *string `json:"access,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRuleAction PasswordPolicyRuleAction

// NewPasswordPolicyRuleAction instantiates a new PasswordPolicyRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRuleAction() *PasswordPolicyRuleAction {
	this := PasswordPolicyRuleAction{}
	return &this
}

// NewPasswordPolicyRuleActionWithDefaults instantiates a new PasswordPolicyRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRuleActionWithDefaults() *PasswordPolicyRuleAction {
	this := PasswordPolicyRuleAction{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *PasswordPolicyRuleAction) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleAction) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *PasswordPolicyRuleAction) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *PasswordPolicyRuleAction) SetAccess(v string) {
	o.Access = &v
}

func (o PasswordPolicyRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyRuleAction := _PasswordPolicyRuleAction{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRuleAction)
	if err == nil {
		*o = PasswordPolicyRuleAction(varPasswordPolicyRuleAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "access")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyRuleAction struct {
	value *PasswordPolicyRuleAction
	isSet bool
}

func (v NullablePasswordPolicyRuleAction) Get() *PasswordPolicyRuleAction {
	return v.value
}

func (v *NullablePasswordPolicyRuleAction) Set(val *PasswordPolicyRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRuleAction(val *PasswordPolicyRuleAction) *NullablePasswordPolicyRuleAction {
	return &NullablePasswordPolicyRuleAction{value: val, isSet: true}
}

func (v NullablePasswordPolicyRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

