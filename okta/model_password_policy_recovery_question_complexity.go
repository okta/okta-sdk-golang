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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordPolicyRecoveryQuestionComplexity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoveryQuestionComplexity{}

// PasswordPolicyRecoveryQuestionComplexity struct for PasswordPolicyRecoveryQuestionComplexity
type PasswordPolicyRecoveryQuestionComplexity struct {
	// Minimum length of the password recovery question answer
	MinLength            *int32 `json:"minLength,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryQuestionComplexity PasswordPolicyRecoveryQuestionComplexity

// NewPasswordPolicyRecoveryQuestionComplexity instantiates a new PasswordPolicyRecoveryQuestionComplexity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryQuestionComplexity() *PasswordPolicyRecoveryQuestionComplexity {
	this := PasswordPolicyRecoveryQuestionComplexity{}
	return &this
}

// NewPasswordPolicyRecoveryQuestionComplexityWithDefaults instantiates a new PasswordPolicyRecoveryQuestionComplexity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryQuestionComplexityWithDefaults() *PasswordPolicyRecoveryQuestionComplexity {
	this := PasswordPolicyRecoveryQuestionComplexity{}
	return &this
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryQuestionComplexity) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryQuestionComplexity) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryQuestionComplexity) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *PasswordPolicyRecoveryQuestionComplexity) SetMinLength(v int32) {
	o.MinLength = &v
}

func (o PasswordPolicyRecoveryQuestionComplexity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoveryQuestionComplexity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoveryQuestionComplexity) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoveryQuestionComplexity := _PasswordPolicyRecoveryQuestionComplexity{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoveryQuestionComplexity)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoveryQuestionComplexity(varPasswordPolicyRecoveryQuestionComplexity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minLength")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoveryQuestionComplexity struct {
	value *PasswordPolicyRecoveryQuestionComplexity
	isSet bool
}

func (v NullablePasswordPolicyRecoveryQuestionComplexity) Get() *PasswordPolicyRecoveryQuestionComplexity {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryQuestionComplexity) Set(val *PasswordPolicyRecoveryQuestionComplexity) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryQuestionComplexity) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryQuestionComplexity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryQuestionComplexity(val *PasswordPolicyRecoveryQuestionComplexity) *NullablePasswordPolicyRecoveryQuestionComplexity {
	return &NullablePasswordPolicyRecoveryQuestionComplexity{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryQuestionComplexity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryQuestionComplexity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
