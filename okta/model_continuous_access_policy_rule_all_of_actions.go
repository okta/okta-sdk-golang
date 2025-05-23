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

// ContinuousAccessPolicyRuleAllOfActions The action to take in response to a failure of the reevaluated global session policy or authentication polices.
type ContinuousAccessPolicyRuleAllOfActions struct {
	ContinuousAccess *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess `json:"continuousAccess,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleAllOfActions ContinuousAccessPolicyRuleAllOfActions

// NewContinuousAccessPolicyRuleAllOfActions instantiates a new ContinuousAccessPolicyRuleAllOfActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleAllOfActions() *ContinuousAccessPolicyRuleAllOfActions {
	this := ContinuousAccessPolicyRuleAllOfActions{}
	return &this
}

// NewContinuousAccessPolicyRuleAllOfActionsWithDefaults instantiates a new ContinuousAccessPolicyRuleAllOfActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleAllOfActionsWithDefaults() *ContinuousAccessPolicyRuleAllOfActions {
	this := ContinuousAccessPolicyRuleAllOfActions{}
	return &this
}

// GetContinuousAccess returns the ContinuousAccess field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleAllOfActions) GetContinuousAccess() ContinuousAccessPolicyRuleAllOfActionsContinuousAccess {
	if o == nil || o.ContinuousAccess == nil {
		var ret ContinuousAccessPolicyRuleAllOfActionsContinuousAccess
		return ret
	}
	return *o.ContinuousAccess
}

// GetContinuousAccessOk returns a tuple with the ContinuousAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleAllOfActions) GetContinuousAccessOk() (*ContinuousAccessPolicyRuleAllOfActionsContinuousAccess, bool) {
	if o == nil || o.ContinuousAccess == nil {
		return nil, false
	}
	return o.ContinuousAccess, true
}

// HasContinuousAccess returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleAllOfActions) HasContinuousAccess() bool {
	if o != nil && o.ContinuousAccess != nil {
		return true
	}

	return false
}

// SetContinuousAccess gets a reference to the given ContinuousAccessPolicyRuleAllOfActionsContinuousAccess and assigns it to the ContinuousAccess field.
func (o *ContinuousAccessPolicyRuleAllOfActions) SetContinuousAccess(v ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) {
	o.ContinuousAccess = &v
}

func (o ContinuousAccessPolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinuousAccess != nil {
		toSerialize["continuousAccess"] = o.ContinuousAccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleAllOfActions) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleAllOfActions := _ContinuousAccessPolicyRuleAllOfActions{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleAllOfActions)
	if err == nil {
		*o = ContinuousAccessPolicyRuleAllOfActions(varContinuousAccessPolicyRuleAllOfActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "continuousAccess")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleAllOfActions struct {
	value *ContinuousAccessPolicyRuleAllOfActions
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleAllOfActions) Get() *ContinuousAccessPolicyRuleAllOfActions {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleAllOfActions) Set(val *ContinuousAccessPolicyRuleAllOfActions) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleAllOfActions) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleAllOfActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleAllOfActions(val *ContinuousAccessPolicyRuleAllOfActions) *NullableContinuousAccessPolicyRuleAllOfActions {
	return &NullableContinuousAccessPolicyRuleAllOfActions{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleAllOfActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

