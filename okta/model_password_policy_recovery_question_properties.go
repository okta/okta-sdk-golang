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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordPolicyRecoveryQuestionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoveryQuestionProperties{}

// PasswordPolicyRecoveryQuestionProperties struct for PasswordPolicyRecoveryQuestionProperties
type PasswordPolicyRecoveryQuestionProperties struct {
	Complexity           *PasswordPolicyRecoveryQuestionComplexity `json:"complexity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryQuestionProperties PasswordPolicyRecoveryQuestionProperties

// NewPasswordPolicyRecoveryQuestionProperties instantiates a new PasswordPolicyRecoveryQuestionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryQuestionProperties() *PasswordPolicyRecoveryQuestionProperties {
	this := PasswordPolicyRecoveryQuestionProperties{}
	return &this
}

// NewPasswordPolicyRecoveryQuestionPropertiesWithDefaults instantiates a new PasswordPolicyRecoveryQuestionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryQuestionPropertiesWithDefaults() *PasswordPolicyRecoveryQuestionProperties {
	this := PasswordPolicyRecoveryQuestionProperties{}
	return &this
}

// GetComplexity returns the Complexity field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryQuestionProperties) GetComplexity() PasswordPolicyRecoveryQuestionComplexity {
	if o == nil || IsNil(o.Complexity) {
		var ret PasswordPolicyRecoveryQuestionComplexity
		return ret
	}
	return *o.Complexity
}

// GetComplexityOk returns a tuple with the Complexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryQuestionProperties) GetComplexityOk() (*PasswordPolicyRecoveryQuestionComplexity, bool) {
	if o == nil || IsNil(o.Complexity) {
		return nil, false
	}
	return o.Complexity, true
}

// HasComplexity returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryQuestionProperties) HasComplexity() bool {
	if o != nil && !IsNil(o.Complexity) {
		return true
	}

	return false
}

// SetComplexity gets a reference to the given PasswordPolicyRecoveryQuestionComplexity and assigns it to the Complexity field.
func (o *PasswordPolicyRecoveryQuestionProperties) SetComplexity(v PasswordPolicyRecoveryQuestionComplexity) {
	o.Complexity = &v
}

func (o PasswordPolicyRecoveryQuestionProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoveryQuestionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Complexity) {
		toSerialize["complexity"] = o.Complexity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoveryQuestionProperties) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoveryQuestionProperties := _PasswordPolicyRecoveryQuestionProperties{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoveryQuestionProperties)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoveryQuestionProperties(varPasswordPolicyRecoveryQuestionProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "complexity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoveryQuestionProperties struct {
	value *PasswordPolicyRecoveryQuestionProperties
	isSet bool
}

func (v NullablePasswordPolicyRecoveryQuestionProperties) Get() *PasswordPolicyRecoveryQuestionProperties {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryQuestionProperties) Set(val *PasswordPolicyRecoveryQuestionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryQuestionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryQuestionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryQuestionProperties(val *PasswordPolicyRecoveryQuestionProperties) *NullablePasswordPolicyRecoveryQuestionProperties {
	return &NullablePasswordPolicyRecoveryQuestionProperties{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryQuestionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryQuestionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
