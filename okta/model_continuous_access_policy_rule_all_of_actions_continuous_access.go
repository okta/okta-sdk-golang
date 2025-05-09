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

// ContinuousAccessPolicyRuleAllOfActionsContinuousAccess This object contains a `failureActions` array that defines the specific action to take when Continuous Access evaluation detects a failure.
type ContinuousAccessPolicyRuleAllOfActionsContinuousAccess struct {
	// An array of objects that define the action. It can be empty or contain two `action` value pairs.
	FailureActions []ContinuousAccessFailureActionsObject `json:"failureActions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleAllOfActionsContinuousAccess ContinuousAccessPolicyRuleAllOfActionsContinuousAccess

// NewContinuousAccessPolicyRuleAllOfActionsContinuousAccess instantiates a new ContinuousAccessPolicyRuleAllOfActionsContinuousAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleAllOfActionsContinuousAccess() *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess {
	this := ContinuousAccessPolicyRuleAllOfActionsContinuousAccess{}
	return &this
}

// NewContinuousAccessPolicyRuleAllOfActionsContinuousAccessWithDefaults instantiates a new ContinuousAccessPolicyRuleAllOfActionsContinuousAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleAllOfActionsContinuousAccessWithDefaults() *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess {
	this := ContinuousAccessPolicyRuleAllOfActionsContinuousAccess{}
	return &this
}

// GetFailureActions returns the FailureActions field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) GetFailureActions() []ContinuousAccessFailureActionsObject {
	if o == nil || o.FailureActions == nil {
		var ret []ContinuousAccessFailureActionsObject
		return ret
	}
	return o.FailureActions
}

// GetFailureActionsOk returns a tuple with the FailureActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) GetFailureActionsOk() ([]ContinuousAccessFailureActionsObject, bool) {
	if o == nil || o.FailureActions == nil {
		return nil, false
	}
	return o.FailureActions, true
}

// HasFailureActions returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) HasFailureActions() bool {
	if o != nil && o.FailureActions != nil {
		return true
	}

	return false
}

// SetFailureActions gets a reference to the given []ContinuousAccessFailureActionsObject and assigns it to the FailureActions field.
func (o *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) SetFailureActions(v []ContinuousAccessFailureActionsObject) {
	o.FailureActions = v
}

func (o ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailureActions != nil {
		toSerialize["failureActions"] = o.FailureActions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleAllOfActionsContinuousAccess := _ContinuousAccessPolicyRuleAllOfActionsContinuousAccess{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleAllOfActionsContinuousAccess)
	if err == nil {
		*o = ContinuousAccessPolicyRuleAllOfActionsContinuousAccess(varContinuousAccessPolicyRuleAllOfActionsContinuousAccess)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "failureActions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess struct {
	value *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess) Get() *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess) Set(val *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess(val *ContinuousAccessPolicyRuleAllOfActionsContinuousAccess) *NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess {
	return &NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleAllOfActionsContinuousAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

