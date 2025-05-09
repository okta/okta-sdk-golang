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

// LinksQuestions struct for LinksQuestions
type LinksQuestions struct {
	Question *LinksQuestionsQuestion `json:"question,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksQuestions LinksQuestions

// NewLinksQuestions instantiates a new LinksQuestions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksQuestions() *LinksQuestions {
	this := LinksQuestions{}
	return &this
}

// NewLinksQuestionsWithDefaults instantiates a new LinksQuestions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksQuestionsWithDefaults() *LinksQuestions {
	this := LinksQuestions{}
	return &this
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *LinksQuestions) GetQuestion() LinksQuestionsQuestion {
	if o == nil || o.Question == nil {
		var ret LinksQuestionsQuestion
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksQuestions) GetQuestionOk() (*LinksQuestionsQuestion, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *LinksQuestions) HasQuestion() bool {
	if o != nil && o.Question != nil {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given LinksQuestionsQuestion and assigns it to the Question field.
func (o *LinksQuestions) SetQuestion(v LinksQuestionsQuestion) {
	o.Question = &v
}

func (o LinksQuestions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Question != nil {
		toSerialize["question"] = o.Question
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksQuestions) UnmarshalJSON(bytes []byte) (err error) {
	varLinksQuestions := _LinksQuestions{}

	err = json.Unmarshal(bytes, &varLinksQuestions)
	if err == nil {
		*o = LinksQuestions(varLinksQuestions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "question")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksQuestions struct {
	value *LinksQuestions
	isSet bool
}

func (v NullableLinksQuestions) Get() *LinksQuestions {
	return v.value
}

func (v *NullableLinksQuestions) Set(val *LinksQuestions) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksQuestions) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksQuestions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksQuestions(val *LinksQuestions) *NullableLinksQuestions {
	return &NullableLinksQuestions{value: val, isSet: true}
}

func (v NullableLinksQuestions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksQuestions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

