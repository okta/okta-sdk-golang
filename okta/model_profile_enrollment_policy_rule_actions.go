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

// ProfileEnrollmentPolicyRuleActions struct for ProfileEnrollmentPolicyRuleActions
type ProfileEnrollmentPolicyRuleActions struct {
	ProfileEnrollment *ProfileEnrollmentPolicyRuleAction `json:"profileEnrollment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileEnrollmentPolicyRuleActions ProfileEnrollmentPolicyRuleActions

// NewProfileEnrollmentPolicyRuleActions instantiates a new ProfileEnrollmentPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentPolicyRuleActions() *ProfileEnrollmentPolicyRuleActions {
	this := ProfileEnrollmentPolicyRuleActions{}
	return &this
}

// NewProfileEnrollmentPolicyRuleActionsWithDefaults instantiates a new ProfileEnrollmentPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentPolicyRuleActionsWithDefaults() *ProfileEnrollmentPolicyRuleActions {
	this := ProfileEnrollmentPolicyRuleActions{}
	return &this
}

// GetProfileEnrollment returns the ProfileEnrollment field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleActions) GetProfileEnrollment() ProfileEnrollmentPolicyRuleAction {
	if o == nil || o.ProfileEnrollment == nil {
		var ret ProfileEnrollmentPolicyRuleAction
		return ret
	}
	return *o.ProfileEnrollment
}

// GetProfileEnrollmentOk returns a tuple with the ProfileEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleActions) GetProfileEnrollmentOk() (*ProfileEnrollmentPolicyRuleAction, bool) {
	if o == nil || o.ProfileEnrollment == nil {
		return nil, false
	}
	return o.ProfileEnrollment, true
}

// HasProfileEnrollment returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleActions) HasProfileEnrollment() bool {
	if o != nil && o.ProfileEnrollment != nil {
		return true
	}

	return false
}

// SetProfileEnrollment gets a reference to the given ProfileEnrollmentPolicyRuleAction and assigns it to the ProfileEnrollment field.
func (o *ProfileEnrollmentPolicyRuleActions) SetProfileEnrollment(v ProfileEnrollmentPolicyRuleAction) {
	o.ProfileEnrollment = &v
}

func (o ProfileEnrollmentPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProfileEnrollment != nil {
		toSerialize["profileEnrollment"] = o.ProfileEnrollment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileEnrollmentPolicyRuleActions) UnmarshalJSON(bytes []byte) (err error) {
	varProfileEnrollmentPolicyRuleActions := _ProfileEnrollmentPolicyRuleActions{}

	err = json.Unmarshal(bytes, &varProfileEnrollmentPolicyRuleActions)
	if err == nil {
		*o = ProfileEnrollmentPolicyRuleActions(varProfileEnrollmentPolicyRuleActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profileEnrollment")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProfileEnrollmentPolicyRuleActions struct {
	value *ProfileEnrollmentPolicyRuleActions
	isSet bool
}

func (v NullableProfileEnrollmentPolicyRuleActions) Get() *ProfileEnrollmentPolicyRuleActions {
	return v.value
}

func (v *NullableProfileEnrollmentPolicyRuleActions) Set(val *ProfileEnrollmentPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentPolicyRuleActions(val *ProfileEnrollmentPolicyRuleActions) *NullableProfileEnrollmentPolicyRuleActions {
	return &NullableProfileEnrollmentPolicyRuleActions{value: val, isSet: true}
}

func (v NullableProfileEnrollmentPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

