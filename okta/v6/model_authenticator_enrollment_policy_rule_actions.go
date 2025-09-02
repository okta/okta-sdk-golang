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

// AuthenticatorEnrollmentPolicyRuleActions struct for AuthenticatorEnrollmentPolicyRuleActions
type AuthenticatorEnrollmentPolicyRuleActions struct {
	Enroll *AuthenticatorEnrollmentPolicyRuleActionEnroll `json:"enroll,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyRuleActions AuthenticatorEnrollmentPolicyRuleActions

// NewAuthenticatorEnrollmentPolicyRuleActions instantiates a new AuthenticatorEnrollmentPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyRuleActions() *AuthenticatorEnrollmentPolicyRuleActions {
	this := AuthenticatorEnrollmentPolicyRuleActions{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyRuleActionsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyRuleActionsWithDefaults() *AuthenticatorEnrollmentPolicyRuleActions {
	this := AuthenticatorEnrollmentPolicyRuleActions{}
	return &this
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRuleActions) GetEnroll() AuthenticatorEnrollmentPolicyRuleActionEnroll {
	if o == nil || o.Enroll == nil {
		var ret AuthenticatorEnrollmentPolicyRuleActionEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRuleActions) GetEnrollOk() (*AuthenticatorEnrollmentPolicyRuleActionEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRuleActions) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given AuthenticatorEnrollmentPolicyRuleActionEnroll and assigns it to the Enroll field.
func (o *AuthenticatorEnrollmentPolicyRuleActions) SetEnroll(v AuthenticatorEnrollmentPolicyRuleActionEnroll) {
	o.Enroll = &v
}

func (o AuthenticatorEnrollmentPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorEnrollmentPolicyRuleActions) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorEnrollmentPolicyRuleActions := _AuthenticatorEnrollmentPolicyRuleActions{}

	err = json.Unmarshal(bytes, &varAuthenticatorEnrollmentPolicyRuleActions)
	if err == nil {
		*o = AuthenticatorEnrollmentPolicyRuleActions(varAuthenticatorEnrollmentPolicyRuleActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enroll")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyRuleActions struct {
	value *AuthenticatorEnrollmentPolicyRuleActions
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyRuleActions) Get() *AuthenticatorEnrollmentPolicyRuleActions {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleActions) Set(val *AuthenticatorEnrollmentPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyRuleActions(val *AuthenticatorEnrollmentPolicyRuleActions) *NullableAuthenticatorEnrollmentPolicyRuleActions {
	return &NullableAuthenticatorEnrollmentPolicyRuleActions{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

