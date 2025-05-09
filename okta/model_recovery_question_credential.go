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

// RecoveryQuestionCredential Specifies a secret question and answer that's validated (case insensitive) when a user forgets their password or unlocks their account. The answer property is write-only.
type RecoveryQuestionCredential struct {
	// The recovery question answer
	Answer *string `json:"answer,omitempty"`
	// The recovery question
	Question *string `json:"question,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryQuestionCredential RecoveryQuestionCredential

// NewRecoveryQuestionCredential instantiates a new RecoveryQuestionCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryQuestionCredential() *RecoveryQuestionCredential {
	this := RecoveryQuestionCredential{}
	return &this
}

// NewRecoveryQuestionCredentialWithDefaults instantiates a new RecoveryQuestionCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryQuestionCredentialWithDefaults() *RecoveryQuestionCredential {
	this := RecoveryQuestionCredential{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *RecoveryQuestionCredential) GetAnswer() string {
	if o == nil || o.Answer == nil {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryQuestionCredential) GetAnswerOk() (*string, bool) {
	if o == nil || o.Answer == nil {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *RecoveryQuestionCredential) HasAnswer() bool {
	if o != nil && o.Answer != nil {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *RecoveryQuestionCredential) SetAnswer(v string) {
	o.Answer = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *RecoveryQuestionCredential) GetQuestion() string {
	if o == nil || o.Question == nil {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryQuestionCredential) GetQuestionOk() (*string, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *RecoveryQuestionCredential) HasQuestion() bool {
	if o != nil && o.Question != nil {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *RecoveryQuestionCredential) SetQuestion(v string) {
	o.Question = &v
}

func (o RecoveryQuestionCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Answer != nil {
		toSerialize["answer"] = o.Answer
	}
	if o.Question != nil {
		toSerialize["question"] = o.Question
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryQuestionCredential) UnmarshalJSON(bytes []byte) (err error) {
	varRecoveryQuestionCredential := _RecoveryQuestionCredential{}

	err = json.Unmarshal(bytes, &varRecoveryQuestionCredential)
	if err == nil {
		*o = RecoveryQuestionCredential(varRecoveryQuestionCredential)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "answer")
		delete(additionalProperties, "question")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRecoveryQuestionCredential struct {
	value *RecoveryQuestionCredential
	isSet bool
}

func (v NullableRecoveryQuestionCredential) Get() *RecoveryQuestionCredential {
	return v.value
}

func (v *NullableRecoveryQuestionCredential) Set(val *RecoveryQuestionCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryQuestionCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryQuestionCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryQuestionCredential(val *RecoveryQuestionCredential) *NullableRecoveryQuestionCredential {
	return &NullableRecoveryQuestionCredential{value: val, isSet: true}
}

func (v NullableRecoveryQuestionCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryQuestionCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

