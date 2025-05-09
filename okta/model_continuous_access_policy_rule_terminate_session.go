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

// ContinuousAccessPolicyRuleTerminateSession struct for ContinuousAccessPolicyRuleTerminateSession
type ContinuousAccessPolicyRuleTerminateSession struct {
	// The action to take when Continuous Access evaluation detects a failure.
	Action *string `json:"action,omitempty"`
	Slo *ContinuousAccessPolicyRuleTerminateSessionSlo `json:"slo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicyRuleTerminateSession ContinuousAccessPolicyRuleTerminateSession

// NewContinuousAccessPolicyRuleTerminateSession instantiates a new ContinuousAccessPolicyRuleTerminateSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicyRuleTerminateSession() *ContinuousAccessPolicyRuleTerminateSession {
	this := ContinuousAccessPolicyRuleTerminateSession{}
	return &this
}

// NewContinuousAccessPolicyRuleTerminateSessionWithDefaults instantiates a new ContinuousAccessPolicyRuleTerminateSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyRuleTerminateSessionWithDefaults() *ContinuousAccessPolicyRuleTerminateSession {
	this := ContinuousAccessPolicyRuleTerminateSession{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleTerminateSession) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleTerminateSession) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleTerminateSession) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ContinuousAccessPolicyRuleTerminateSession) SetAction(v string) {
	o.Action = &v
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *ContinuousAccessPolicyRuleTerminateSession) GetSlo() ContinuousAccessPolicyRuleTerminateSessionSlo {
	if o == nil || o.Slo == nil {
		var ret ContinuousAccessPolicyRuleTerminateSessionSlo
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessPolicyRuleTerminateSession) GetSloOk() (*ContinuousAccessPolicyRuleTerminateSessionSlo, bool) {
	if o == nil || o.Slo == nil {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *ContinuousAccessPolicyRuleTerminateSession) HasSlo() bool {
	if o != nil && o.Slo != nil {
		return true
	}

	return false
}

// SetSlo gets a reference to the given ContinuousAccessPolicyRuleTerminateSessionSlo and assigns it to the Slo field.
func (o *ContinuousAccessPolicyRuleTerminateSession) SetSlo(v ContinuousAccessPolicyRuleTerminateSessionSlo) {
	o.Slo = &v
}

func (o ContinuousAccessPolicyRuleTerminateSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Slo != nil {
		toSerialize["slo"] = o.Slo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicyRuleTerminateSession) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessPolicyRuleTerminateSession := _ContinuousAccessPolicyRuleTerminateSession{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyRuleTerminateSession)
	if err == nil {
		*o = ContinuousAccessPolicyRuleTerminateSession(varContinuousAccessPolicyRuleTerminateSession)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "slo")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicyRuleTerminateSession struct {
	value *ContinuousAccessPolicyRuleTerminateSession
	isSet bool
}

func (v NullableContinuousAccessPolicyRuleTerminateSession) Get() *ContinuousAccessPolicyRuleTerminateSession {
	return v.value
}

func (v *NullableContinuousAccessPolicyRuleTerminateSession) Set(val *ContinuousAccessPolicyRuleTerminateSession) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicyRuleTerminateSession) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicyRuleTerminateSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicyRuleTerminateSession(val *ContinuousAccessPolicyRuleTerminateSession) *NullableContinuousAccessPolicyRuleTerminateSession {
	return &NullableContinuousAccessPolicyRuleTerminateSession{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicyRuleTerminateSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicyRuleTerminateSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

