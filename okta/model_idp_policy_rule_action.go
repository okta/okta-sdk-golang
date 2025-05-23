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

// IdpPolicyRuleAction struct for IdpPolicyRuleAction
type IdpPolicyRuleAction struct {
	Idp *IdpPolicyRuleActionIdp `json:"idp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpPolicyRuleAction IdpPolicyRuleAction

// NewIdpPolicyRuleAction instantiates a new IdpPolicyRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPolicyRuleAction() *IdpPolicyRuleAction {
	this := IdpPolicyRuleAction{}
	return &this
}

// NewIdpPolicyRuleActionWithDefaults instantiates a new IdpPolicyRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPolicyRuleActionWithDefaults() *IdpPolicyRuleAction {
	this := IdpPolicyRuleAction{}
	return &this
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IdpPolicyRuleAction) GetIdp() IdpPolicyRuleActionIdp {
	if o == nil || o.Idp == nil {
		var ret IdpPolicyRuleActionIdp
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleAction) GetIdpOk() (*IdpPolicyRuleActionIdp, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IdpPolicyRuleAction) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IdpPolicyRuleActionIdp and assigns it to the Idp field.
func (o *IdpPolicyRuleAction) SetIdp(v IdpPolicyRuleActionIdp) {
	o.Idp = &v
}

func (o IdpPolicyRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdpPolicyRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	varIdpPolicyRuleAction := _IdpPolicyRuleAction{}

	err = json.Unmarshal(bytes, &varIdpPolicyRuleAction)
	if err == nil {
		*o = IdpPolicyRuleAction(varIdpPolicyRuleAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "idp")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdpPolicyRuleAction struct {
	value *IdpPolicyRuleAction
	isSet bool
}

func (v NullableIdpPolicyRuleAction) Get() *IdpPolicyRuleAction {
	return v.value
}

func (v *NullableIdpPolicyRuleAction) Set(val *IdpPolicyRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPolicyRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPolicyRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPolicyRuleAction(val *IdpPolicyRuleAction) *NullableIdpPolicyRuleAction {
	return &NullableIdpPolicyRuleAction{value: val, isSet: true}
}

func (v NullableIdpPolicyRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPolicyRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

