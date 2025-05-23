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

// UserFactorSecurityQuestionProfile struct for UserFactorSecurityQuestionProfile
type UserFactorSecurityQuestionProfile struct {
	// Answer to the question
	Answer *string `json:"answer,omitempty"`
	// Unique key for the question
	Question *string `json:"question,omitempty"`
	// Human-readable text displayed to the user
	QuestionText *string `json:"questionText,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSecurityQuestionProfile UserFactorSecurityQuestionProfile

// NewUserFactorSecurityQuestionProfile instantiates a new UserFactorSecurityQuestionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSecurityQuestionProfile() *UserFactorSecurityQuestionProfile {
	this := UserFactorSecurityQuestionProfile{}
	return &this
}

// NewUserFactorSecurityQuestionProfileWithDefaults instantiates a new UserFactorSecurityQuestionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSecurityQuestionProfileWithDefaults() *UserFactorSecurityQuestionProfile {
	this := UserFactorSecurityQuestionProfile{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *UserFactorSecurityQuestionProfile) GetAnswer() string {
	if o == nil || o.Answer == nil {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSecurityQuestionProfile) GetAnswerOk() (*string, bool) {
	if o == nil || o.Answer == nil {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *UserFactorSecurityQuestionProfile) HasAnswer() bool {
	if o != nil && o.Answer != nil {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *UserFactorSecurityQuestionProfile) SetAnswer(v string) {
	o.Answer = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *UserFactorSecurityQuestionProfile) GetQuestion() string {
	if o == nil || o.Question == nil {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSecurityQuestionProfile) GetQuestionOk() (*string, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *UserFactorSecurityQuestionProfile) HasQuestion() bool {
	if o != nil && o.Question != nil {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *UserFactorSecurityQuestionProfile) SetQuestion(v string) {
	o.Question = &v
}

// GetQuestionText returns the QuestionText field value if set, zero value otherwise.
func (o *UserFactorSecurityQuestionProfile) GetQuestionText() string {
	if o == nil || o.QuestionText == nil {
		var ret string
		return ret
	}
	return *o.QuestionText
}

// GetQuestionTextOk returns a tuple with the QuestionText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSecurityQuestionProfile) GetQuestionTextOk() (*string, bool) {
	if o == nil || o.QuestionText == nil {
		return nil, false
	}
	return o.QuestionText, true
}

// HasQuestionText returns a boolean if a field has been set.
func (o *UserFactorSecurityQuestionProfile) HasQuestionText() bool {
	if o != nil && o.QuestionText != nil {
		return true
	}

	return false
}

// SetQuestionText gets a reference to the given string and assigns it to the QuestionText field.
func (o *UserFactorSecurityQuestionProfile) SetQuestionText(v string) {
	o.QuestionText = &v
}

func (o UserFactorSecurityQuestionProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Answer != nil {
		toSerialize["answer"] = o.Answer
	}
	if o.Question != nil {
		toSerialize["question"] = o.Question
	}
	if o.QuestionText != nil {
		toSerialize["questionText"] = o.QuestionText
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorSecurityQuestionProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorSecurityQuestionProfile := _UserFactorSecurityQuestionProfile{}

	err = json.Unmarshal(bytes, &varUserFactorSecurityQuestionProfile)
	if err == nil {
		*o = UserFactorSecurityQuestionProfile(varUserFactorSecurityQuestionProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "answer")
		delete(additionalProperties, "question")
		delete(additionalProperties, "questionText")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorSecurityQuestionProfile struct {
	value *UserFactorSecurityQuestionProfile
	isSet bool
}

func (v NullableUserFactorSecurityQuestionProfile) Get() *UserFactorSecurityQuestionProfile {
	return v.value
}

func (v *NullableUserFactorSecurityQuestionProfile) Set(val *UserFactorSecurityQuestionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSecurityQuestionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSecurityQuestionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSecurityQuestionProfile(val *UserFactorSecurityQuestionProfile) *NullableUserFactorSecurityQuestionProfile {
	return &NullableUserFactorSecurityQuestionProfile{value: val, isSet: true}
}

func (v NullableUserFactorSecurityQuestionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSecurityQuestionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

