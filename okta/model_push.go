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

// checks if the Push type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Push{}

// Push Sends an asynchronous push notification to the device for approval by the user. You must poll the transaction to determine the state of the verification. See [Retrieve a factor transaction status](./#tag/UserFactor/operation/getFactorTransactionStatus).  Activations have a short lifetime of several minutes and return a `TIMEOUT` if not completed before the timestamp specified in the `expiresAt` param. Use the published activate link to restart the activation process if the activation expires.
type Push struct {
	// Select whether to use a number matching challenge for a `push` factor.  > **Note:** Sending a request with a body is required when you verify a `push` factor with a number matching challenge.
	UseNumberMatchingChallenge *bool `json:"useNumberMatchingChallenge,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _Push Push

// NewPush instantiates a new Push object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPush() *Push {
	this := Push{}
	return &this
}

// NewPushWithDefaults instantiates a new Push object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushWithDefaults() *Push {
	this := Push{}
	return &this
}

// GetUseNumberMatchingChallenge returns the UseNumberMatchingChallenge field value if set, zero value otherwise.
func (o *Push) GetUseNumberMatchingChallenge() bool {
	if o == nil || IsNil(o.UseNumberMatchingChallenge) {
		var ret bool
		return ret
	}
	return *o.UseNumberMatchingChallenge
}

// GetUseNumberMatchingChallengeOk returns a tuple with the UseNumberMatchingChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Push) GetUseNumberMatchingChallengeOk() (*bool, bool) {
	if o == nil || IsNil(o.UseNumberMatchingChallenge) {
		return nil, false
	}
	return o.UseNumberMatchingChallenge, true
}

// HasUseNumberMatchingChallenge returns a boolean if a field has been set.
func (o *Push) HasUseNumberMatchingChallenge() bool {
	if o != nil && !IsNil(o.UseNumberMatchingChallenge) {
		return true
	}

	return false
}

// SetUseNumberMatchingChallenge gets a reference to the given bool and assigns it to the UseNumberMatchingChallenge field.
func (o *Push) SetUseNumberMatchingChallenge(v bool) {
	o.UseNumberMatchingChallenge = &v
}

func (o Push) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Push) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseNumberMatchingChallenge) {
		toSerialize["useNumberMatchingChallenge"] = o.UseNumberMatchingChallenge
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Push) UnmarshalJSON(data []byte) (err error) {
	varPush := _Push{}

	err = json.Unmarshal(data, &varPush)

	if err != nil {
		return err
	}

	*o = Push(varPush)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "useNumberMatchingChallenge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePush struct {
	value *Push
	isSet bool
}

func (v NullablePush) Get() *Push {
	return v.value
}

func (v *NullablePush) Set(val *Push) {
	v.value = val
	v.isSet = true
}

func (v NullablePush) IsSet() bool {
	return v.isSet
}

func (v *NullablePush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePush(val *Push) *NullablePush {
	return &NullablePush{value: val, isSet: true}
}

func (v NullablePush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
