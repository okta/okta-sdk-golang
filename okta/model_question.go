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

// checks if the Question type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Question{}

// Question Verifies an answer to a `question` factor
type Question struct {
	// Answer to the question
	Answer               *string `json:"answer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Question Question

// NewQuestion instantiates a new Question object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestion() *Question {
	this := Question{}
	return &this
}

// NewQuestionWithDefaults instantiates a new Question object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionWithDefaults() *Question {
	this := Question{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *Question) GetAnswer() string {
	if o == nil || IsNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Question) GetAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *Question) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *Question) SetAnswer(v string) {
	o.Answer = &v
}

func (o Question) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Question) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Question) UnmarshalJSON(data []byte) (err error) {
	varQuestion := _Question{}

	err = json.Unmarshal(data, &varQuestion)

	if err != nil {
		return err
	}

	*o = Question(varQuestion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "answer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuestion struct {
	value *Question
	isSet bool
}

func (v NullableQuestion) Get() *Question {
	return v.value
}

func (v *NullableQuestion) Set(val *Question) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestion(val *Question) *NullableQuestion {
	return &NullableQuestion{value: val, isSet: true}
}

func (v NullableQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
