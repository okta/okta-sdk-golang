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

// checks if the NumberFactorChallengeEmbeddedLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NumberFactorChallengeEmbeddedLinks{}

// NumberFactorChallengeEmbeddedLinks Contains the `challenge` and `correctAnswer` objects for `push` factors that use a number matching challenge
type NumberFactorChallengeEmbeddedLinks struct {
	Challenge            NullableNumberFactorChallengeEmbeddedLinksChallenge `json:"challenge,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NumberFactorChallengeEmbeddedLinks NumberFactorChallengeEmbeddedLinks

// NewNumberFactorChallengeEmbeddedLinks instantiates a new NumberFactorChallengeEmbeddedLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberFactorChallengeEmbeddedLinks() *NumberFactorChallengeEmbeddedLinks {
	this := NumberFactorChallengeEmbeddedLinks{}
	return &this
}

// NewNumberFactorChallengeEmbeddedLinksWithDefaults instantiates a new NumberFactorChallengeEmbeddedLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberFactorChallengeEmbeddedLinksWithDefaults() *NumberFactorChallengeEmbeddedLinks {
	this := NumberFactorChallengeEmbeddedLinks{}
	return &this
}

// GetChallenge returns the Challenge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NumberFactorChallengeEmbeddedLinks) GetChallenge() NumberFactorChallengeEmbeddedLinksChallenge {
	if o == nil || IsNil(o.Challenge.Get()) {
		var ret NumberFactorChallengeEmbeddedLinksChallenge
		return ret
	}
	return *o.Challenge.Get()
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NumberFactorChallengeEmbeddedLinks) GetChallengeOk() (*NumberFactorChallengeEmbeddedLinksChallenge, bool) {
	if o == nil {
		return nil, false
	}
	return o.Challenge.Get(), o.Challenge.IsSet()
}

// HasChallenge returns a boolean if a field has been set.
func (o *NumberFactorChallengeEmbeddedLinks) HasChallenge() bool {
	if o != nil && o.Challenge.IsSet() {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given NullableNumberFactorChallengeEmbeddedLinksChallenge and assigns it to the Challenge field.
func (o *NumberFactorChallengeEmbeddedLinks) SetChallenge(v NumberFactorChallengeEmbeddedLinksChallenge) {
	o.Challenge.Set(&v)
}

// SetChallengeNil sets the value for Challenge to be an explicit nil
func (o *NumberFactorChallengeEmbeddedLinks) SetChallengeNil() {
	o.Challenge.Set(nil)
}

// UnsetChallenge ensures that no value is present for Challenge, not even an explicit nil
func (o *NumberFactorChallengeEmbeddedLinks) UnsetChallenge() {
	o.Challenge.Unset()
}

func (o NumberFactorChallengeEmbeddedLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberFactorChallengeEmbeddedLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Challenge.IsSet() {
		toSerialize["challenge"] = o.Challenge.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NumberFactorChallengeEmbeddedLinks) UnmarshalJSON(data []byte) (err error) {
	varNumberFactorChallengeEmbeddedLinks := _NumberFactorChallengeEmbeddedLinks{}

	err = json.Unmarshal(data, &varNumberFactorChallengeEmbeddedLinks)

	if err != nil {
		return err
	}

	*o = NumberFactorChallengeEmbeddedLinks(varNumberFactorChallengeEmbeddedLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "challenge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumberFactorChallengeEmbeddedLinks struct {
	value *NumberFactorChallengeEmbeddedLinks
	isSet bool
}

func (v NullableNumberFactorChallengeEmbeddedLinks) Get() *NumberFactorChallengeEmbeddedLinks {
	return v.value
}

func (v *NullableNumberFactorChallengeEmbeddedLinks) Set(val *NumberFactorChallengeEmbeddedLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberFactorChallengeEmbeddedLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberFactorChallengeEmbeddedLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberFactorChallengeEmbeddedLinks(val *NumberFactorChallengeEmbeddedLinks) *NullableNumberFactorChallengeEmbeddedLinks {
	return &NullableNumberFactorChallengeEmbeddedLinks{value: val, isSet: true}
}

func (v NullableNumberFactorChallengeEmbeddedLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberFactorChallengeEmbeddedLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
