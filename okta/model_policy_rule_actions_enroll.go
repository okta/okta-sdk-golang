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

// PolicyRuleActionsEnroll struct for PolicyRuleActionsEnroll
type PolicyRuleActionsEnroll struct {
	Self *string `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyRuleActionsEnroll PolicyRuleActionsEnroll

// NewPolicyRuleActionsEnroll instantiates a new PolicyRuleActionsEnroll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRuleActionsEnroll() *PolicyRuleActionsEnroll {
	this := PolicyRuleActionsEnroll{}
	return &this
}

// NewPolicyRuleActionsEnrollWithDefaults instantiates a new PolicyRuleActionsEnroll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRuleActionsEnrollWithDefaults() *PolicyRuleActionsEnroll {
	this := PolicyRuleActionsEnroll{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PolicyRuleActionsEnroll) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRuleActionsEnroll) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PolicyRuleActionsEnroll) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *PolicyRuleActionsEnroll) SetSelf(v string) {
	o.Self = &v
}

func (o PolicyRuleActionsEnroll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyRuleActionsEnroll) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyRuleActionsEnroll := _PolicyRuleActionsEnroll{}

	err = json.Unmarshal(bytes, &varPolicyRuleActionsEnroll)
	if err == nil {
		*o = PolicyRuleActionsEnroll(varPolicyRuleActionsEnroll)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyRuleActionsEnroll struct {
	value *PolicyRuleActionsEnroll
	isSet bool
}

func (v NullablePolicyRuleActionsEnroll) Get() *PolicyRuleActionsEnroll {
	return v.value
}

func (v *NullablePolicyRuleActionsEnroll) Set(val *PolicyRuleActionsEnroll) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRuleActionsEnroll) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRuleActionsEnroll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRuleActionsEnroll(val *PolicyRuleActionsEnroll) *NullablePolicyRuleActionsEnroll {
	return &NullablePolicyRuleActionsEnroll{value: val, isSet: true}
}

func (v NullablePolicyRuleActionsEnroll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRuleActionsEnroll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

