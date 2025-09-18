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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserFactorVerifyResponseWaitingEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorVerifyResponseWaitingEmbedded{}

// UserFactorVerifyResponseWaitingEmbedded struct for UserFactorVerifyResponseWaitingEmbedded
type UserFactorVerifyResponseWaitingEmbedded struct {
	Challenge            NullableNumberFactorChallengeEmbeddedLinksChallenge `json:"challenge,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorVerifyResponseWaitingEmbedded UserFactorVerifyResponseWaitingEmbedded

// NewUserFactorVerifyResponseWaitingEmbedded instantiates a new UserFactorVerifyResponseWaitingEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorVerifyResponseWaitingEmbedded() *UserFactorVerifyResponseWaitingEmbedded {
	this := UserFactorVerifyResponseWaitingEmbedded{}
	return &this
}

// NewUserFactorVerifyResponseWaitingEmbeddedWithDefaults instantiates a new UserFactorVerifyResponseWaitingEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorVerifyResponseWaitingEmbeddedWithDefaults() *UserFactorVerifyResponseWaitingEmbedded {
	this := UserFactorVerifyResponseWaitingEmbedded{}
	return &this
}

// GetChallenge returns the Challenge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorVerifyResponseWaitingEmbedded) GetChallenge() NumberFactorChallengeEmbeddedLinksChallenge {
	if o == nil || IsNil(o.Challenge.Get()) {
		var ret NumberFactorChallengeEmbeddedLinksChallenge
		return ret
	}
	return *o.Challenge.Get()
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorVerifyResponseWaitingEmbedded) GetChallengeOk() (*NumberFactorChallengeEmbeddedLinksChallenge, bool) {
	if o == nil {
		return nil, false
	}
	return o.Challenge.Get(), o.Challenge.IsSet()
}

// HasChallenge returns a boolean if a field has been set.
func (o *UserFactorVerifyResponseWaitingEmbedded) HasChallenge() bool {
	if o != nil && o.Challenge.IsSet() {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given NullableNumberFactorChallengeEmbeddedLinksChallenge and assigns it to the Challenge field.
func (o *UserFactorVerifyResponseWaitingEmbedded) SetChallenge(v NumberFactorChallengeEmbeddedLinksChallenge) {
	o.Challenge.Set(&v)
}

// SetChallengeNil sets the value for Challenge to be an explicit nil
func (o *UserFactorVerifyResponseWaitingEmbedded) SetChallengeNil() {
	o.Challenge.Set(nil)
}

// UnsetChallenge ensures that no value is present for Challenge, not even an explicit nil
func (o *UserFactorVerifyResponseWaitingEmbedded) UnsetChallenge() {
	o.Challenge.Unset()
}

func (o UserFactorVerifyResponseWaitingEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorVerifyResponseWaitingEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Challenge.IsSet() {
		toSerialize["challenge"] = o.Challenge.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorVerifyResponseWaitingEmbedded) UnmarshalJSON(data []byte) (err error) {
	varUserFactorVerifyResponseWaitingEmbedded := _UserFactorVerifyResponseWaitingEmbedded{}

	err = json.Unmarshal(data, &varUserFactorVerifyResponseWaitingEmbedded)

	if err != nil {
		return err
	}

	*o = UserFactorVerifyResponseWaitingEmbedded(varUserFactorVerifyResponseWaitingEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "challenge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorVerifyResponseWaitingEmbedded struct {
	value *UserFactorVerifyResponseWaitingEmbedded
	isSet bool
}

func (v NullableUserFactorVerifyResponseWaitingEmbedded) Get() *UserFactorVerifyResponseWaitingEmbedded {
	return v.value
}

func (v *NullableUserFactorVerifyResponseWaitingEmbedded) Set(val *UserFactorVerifyResponseWaitingEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorVerifyResponseWaitingEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorVerifyResponseWaitingEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorVerifyResponseWaitingEmbedded(val *UserFactorVerifyResponseWaitingEmbedded) *NullableUserFactorVerifyResponseWaitingEmbedded {
	return &NullableUserFactorVerifyResponseWaitingEmbedded{value: val, isSet: true}
}

func (v NullableUserFactorVerifyResponseWaitingEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorVerifyResponseWaitingEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
