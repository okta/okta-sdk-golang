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

// MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints struct for MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints
type MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints struct {
	AaguidGroups []string `json:"aaguidGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints

// NewMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints() *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints {
	this := MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints{}
	return &this
}

// NewMultifactorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults() *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints {
	this := MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints{}
	return &this
}

// GetAaguidGroups returns the AaguidGroups field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) GetAaguidGroups() []string {
	if o == nil || o.AaguidGroups == nil {
		var ret []string
		return ret
	}
	return o.AaguidGroups
}

// GetAaguidGroupsOk returns a tuple with the AaguidGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) GetAaguidGroupsOk() ([]string, bool) {
	if o == nil || o.AaguidGroups == nil {
		return nil, false
	}
	return o.AaguidGroups, true
}

// HasAaguidGroups returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) HasAaguidGroups() bool {
	if o != nil && o.AaguidGroups != nil {
		return true
	}

	return false
}

// SetAaguidGroups gets a reference to the given []string and assigns it to the AaguidGroups field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) SetAaguidGroups(v []string) {
	o.AaguidGroups = v
}

func (o MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AaguidGroups != nil {
		toSerialize["aaguidGroups"] = o.AaguidGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) UnmarshalJSON(bytes []byte) (err error) {
	varMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints := _MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints)
	if err == nil {
		*o = MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints(varMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "aaguidGroups")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints struct {
	value *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) Get() *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) Set(val *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints(val *MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints {
	return &NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

