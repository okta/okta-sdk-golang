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

// checks if the PasswordPolicyRecoveryQuestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoveryQuestion{}

// PasswordPolicyRecoveryQuestion struct for PasswordPolicyRecoveryQuestion
type PasswordPolicyRecoveryQuestion struct {
	Properties *PasswordPolicyRecoveryQuestionProperties `json:"properties,omitempty"`
	// Whether or not the factor is active
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryQuestion PasswordPolicyRecoveryQuestion

// NewPasswordPolicyRecoveryQuestion instantiates a new PasswordPolicyRecoveryQuestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryQuestion() *PasswordPolicyRecoveryQuestion {
	this := PasswordPolicyRecoveryQuestion{}
	return &this
}

// NewPasswordPolicyRecoveryQuestionWithDefaults instantiates a new PasswordPolicyRecoveryQuestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryQuestionWithDefaults() *PasswordPolicyRecoveryQuestion {
	this := PasswordPolicyRecoveryQuestion{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryQuestion) GetProperties() PasswordPolicyRecoveryQuestionProperties {
	if o == nil || IsNil(o.Properties) {
		var ret PasswordPolicyRecoveryQuestionProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryQuestion) GetPropertiesOk() (*PasswordPolicyRecoveryQuestionProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryQuestion) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given PasswordPolicyRecoveryQuestionProperties and assigns it to the Properties field.
func (o *PasswordPolicyRecoveryQuestion) SetProperties(v PasswordPolicyRecoveryQuestionProperties) {
	o.Properties = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryQuestion) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryQuestion) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryQuestion) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PasswordPolicyRecoveryQuestion) SetStatus(v string) {
	o.Status = &v
}

func (o PasswordPolicyRecoveryQuestion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoveryQuestion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoveryQuestion) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoveryQuestion := _PasswordPolicyRecoveryQuestion{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoveryQuestion)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoveryQuestion(varPasswordPolicyRecoveryQuestion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "properties")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoveryQuestion struct {
	value *PasswordPolicyRecoveryQuestion
	isSet bool
}

func (v NullablePasswordPolicyRecoveryQuestion) Get() *PasswordPolicyRecoveryQuestion {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryQuestion) Set(val *PasswordPolicyRecoveryQuestion) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryQuestion(val *PasswordPolicyRecoveryQuestion) *NullablePasswordPolicyRecoveryQuestion {
	return &NullablePasswordPolicyRecoveryQuestion{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
