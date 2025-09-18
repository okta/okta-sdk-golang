/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

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

// checks if the ProfileEnrollmentPolicyRuleActivationRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileEnrollmentPolicyRuleActivationRequirement{}

// ProfileEnrollmentPolicyRuleActivationRequirement Contains a single Boolean property that indicates whether `emailVerification` should occur (`true`) or not (`false`, default)
type ProfileEnrollmentPolicyRuleActivationRequirement struct {
	EmailVerification    *bool `json:"emailVerification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileEnrollmentPolicyRuleActivationRequirement ProfileEnrollmentPolicyRuleActivationRequirement

// NewProfileEnrollmentPolicyRuleActivationRequirement instantiates a new ProfileEnrollmentPolicyRuleActivationRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentPolicyRuleActivationRequirement() *ProfileEnrollmentPolicyRuleActivationRequirement {
	this := ProfileEnrollmentPolicyRuleActivationRequirement{}
	return &this
}

// NewProfileEnrollmentPolicyRuleActivationRequirementWithDefaults instantiates a new ProfileEnrollmentPolicyRuleActivationRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentPolicyRuleActivationRequirementWithDefaults() *ProfileEnrollmentPolicyRuleActivationRequirement {
	this := ProfileEnrollmentPolicyRuleActivationRequirement{}
	return &this
}

// GetEmailVerification returns the EmailVerification field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleActivationRequirement) GetEmailVerification() bool {
	if o == nil || IsNil(o.EmailVerification) {
		var ret bool
		return ret
	}
	return *o.EmailVerification
}

// GetEmailVerificationOk returns a tuple with the EmailVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleActivationRequirement) GetEmailVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerification) {
		return nil, false
	}
	return o.EmailVerification, true
}

// HasEmailVerification returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleActivationRequirement) HasEmailVerification() bool {
	if o != nil && !IsNil(o.EmailVerification) {
		return true
	}

	return false
}

// SetEmailVerification gets a reference to the given bool and assigns it to the EmailVerification field.
func (o *ProfileEnrollmentPolicyRuleActivationRequirement) SetEmailVerification(v bool) {
	o.EmailVerification = &v
}

func (o ProfileEnrollmentPolicyRuleActivationRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileEnrollmentPolicyRuleActivationRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailVerification) {
		toSerialize["emailVerification"] = o.EmailVerification
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileEnrollmentPolicyRuleActivationRequirement) UnmarshalJSON(data []byte) (err error) {
	varProfileEnrollmentPolicyRuleActivationRequirement := _ProfileEnrollmentPolicyRuleActivationRequirement{}

	err = json.Unmarshal(data, &varProfileEnrollmentPolicyRuleActivationRequirement)

	if err != nil {
		return err
	}

	*o = ProfileEnrollmentPolicyRuleActivationRequirement(varProfileEnrollmentPolicyRuleActivationRequirement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emailVerification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileEnrollmentPolicyRuleActivationRequirement struct {
	value *ProfileEnrollmentPolicyRuleActivationRequirement
	isSet bool
}

func (v NullableProfileEnrollmentPolicyRuleActivationRequirement) Get() *ProfileEnrollmentPolicyRuleActivationRequirement {
	return v.value
}

func (v *NullableProfileEnrollmentPolicyRuleActivationRequirement) Set(val *ProfileEnrollmentPolicyRuleActivationRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentPolicyRuleActivationRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentPolicyRuleActivationRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentPolicyRuleActivationRequirement(val *ProfileEnrollmentPolicyRuleActivationRequirement) *NullableProfileEnrollmentPolicyRuleActivationRequirement {
	return &NullableProfileEnrollmentPolicyRuleActivationRequirement{value: val, isSet: true}
}

func (v NullableProfileEnrollmentPolicyRuleActivationRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentPolicyRuleActivationRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
