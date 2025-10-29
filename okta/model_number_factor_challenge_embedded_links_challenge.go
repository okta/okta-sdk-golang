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

// checks if the NumberFactorChallengeEmbeddedLinksChallenge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NumberFactorChallengeEmbeddedLinksChallenge{}

// NumberFactorChallengeEmbeddedLinksChallenge Number matching challenge for a `push` factor
type NumberFactorChallengeEmbeddedLinksChallenge struct {
	// The correct answer for a `push` factor that uses a number matching challenge
	CorrectAnswer        *int32 `json:"correctAnswer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NumberFactorChallengeEmbeddedLinksChallenge NumberFactorChallengeEmbeddedLinksChallenge

// NewNumberFactorChallengeEmbeddedLinksChallenge instantiates a new NumberFactorChallengeEmbeddedLinksChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberFactorChallengeEmbeddedLinksChallenge() *NumberFactorChallengeEmbeddedLinksChallenge {
	this := NumberFactorChallengeEmbeddedLinksChallenge{}
	return &this
}

// NewNumberFactorChallengeEmbeddedLinksChallengeWithDefaults instantiates a new NumberFactorChallengeEmbeddedLinksChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberFactorChallengeEmbeddedLinksChallengeWithDefaults() *NumberFactorChallengeEmbeddedLinksChallenge {
	this := NumberFactorChallengeEmbeddedLinksChallenge{}
	return &this
}

// GetCorrectAnswer returns the CorrectAnswer field value if set, zero value otherwise.
func (o *NumberFactorChallengeEmbeddedLinksChallenge) GetCorrectAnswer() int32 {
	if o == nil || IsNil(o.CorrectAnswer) {
		var ret int32
		return ret
	}
	return *o.CorrectAnswer
}

// GetCorrectAnswerOk returns a tuple with the CorrectAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFactorChallengeEmbeddedLinksChallenge) GetCorrectAnswerOk() (*int32, bool) {
	if o == nil || IsNil(o.CorrectAnswer) {
		return nil, false
	}
	return o.CorrectAnswer, true
}

// HasCorrectAnswer returns a boolean if a field has been set.
func (o *NumberFactorChallengeEmbeddedLinksChallenge) HasCorrectAnswer() bool {
	if o != nil && !IsNil(o.CorrectAnswer) {
		return true
	}

	return false
}

// SetCorrectAnswer gets a reference to the given int32 and assigns it to the CorrectAnswer field.
func (o *NumberFactorChallengeEmbeddedLinksChallenge) SetCorrectAnswer(v int32) {
	o.CorrectAnswer = &v
}

func (o NumberFactorChallengeEmbeddedLinksChallenge) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberFactorChallengeEmbeddedLinksChallenge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrectAnswer) {
		toSerialize["correctAnswer"] = o.CorrectAnswer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NumberFactorChallengeEmbeddedLinksChallenge) UnmarshalJSON(data []byte) (err error) {
	varNumberFactorChallengeEmbeddedLinksChallenge := _NumberFactorChallengeEmbeddedLinksChallenge{}

	err = json.Unmarshal(data, &varNumberFactorChallengeEmbeddedLinksChallenge)

	if err != nil {
		return err
	}

	*o = NumberFactorChallengeEmbeddedLinksChallenge(varNumberFactorChallengeEmbeddedLinksChallenge)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "correctAnswer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumberFactorChallengeEmbeddedLinksChallenge struct {
	value *NumberFactorChallengeEmbeddedLinksChallenge
	isSet bool
}

func (v NullableNumberFactorChallengeEmbeddedLinksChallenge) Get() *NumberFactorChallengeEmbeddedLinksChallenge {
	return v.value
}

func (v *NullableNumberFactorChallengeEmbeddedLinksChallenge) Set(val *NumberFactorChallengeEmbeddedLinksChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberFactorChallengeEmbeddedLinksChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberFactorChallengeEmbeddedLinksChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberFactorChallengeEmbeddedLinksChallenge(val *NumberFactorChallengeEmbeddedLinksChallenge) *NullableNumberFactorChallengeEmbeddedLinksChallenge {
	return &NullableNumberFactorChallengeEmbeddedLinksChallenge{value: val, isSet: true}
}

func (v NullableNumberFactorChallengeEmbeddedLinksChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberFactorChallengeEmbeddedLinksChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
