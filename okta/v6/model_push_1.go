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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// Push1 Sends an asynchronous push notification to the device for approval by the user. A successful request returns an HTTP 201 response, unlike other factors. You must poll the transaction to determine the state of the verification. See [Retrieve a factor transaction status](./#tag/UserFactor/operation/getFactorTransactionStatus).
type Push1 struct {
	// Select whether to use a number matching challenge for a `push` factor.  > **Note:** Sending a request with a body is required when you verify a `push` factor with a number matching challenge.
	UseNumberMatchingChallenge *bool `json:"useNumberMatchingChallenge,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Push1 Push1

// NewPush1 instantiates a new Push1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPush1() *Push1 {
	this := Push1{}
	return &this
}

// NewPush1WithDefaults instantiates a new Push1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPush1WithDefaults() *Push1 {
	this := Push1{}
	return &this
}

// GetUseNumberMatchingChallenge returns the UseNumberMatchingChallenge field value if set, zero value otherwise.
func (o *Push1) GetUseNumberMatchingChallenge() bool {
	if o == nil || o.UseNumberMatchingChallenge == nil {
		var ret bool
		return ret
	}
	return *o.UseNumberMatchingChallenge
}

// GetUseNumberMatchingChallengeOk returns a tuple with the UseNumberMatchingChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Push1) GetUseNumberMatchingChallengeOk() (*bool, bool) {
	if o == nil || o.UseNumberMatchingChallenge == nil {
		return nil, false
	}
	return o.UseNumberMatchingChallenge, true
}

// HasUseNumberMatchingChallenge returns a boolean if a field has been set.
func (o *Push1) HasUseNumberMatchingChallenge() bool {
	if o != nil && o.UseNumberMatchingChallenge != nil {
		return true
	}

	return false
}

// SetUseNumberMatchingChallenge gets a reference to the given bool and assigns it to the UseNumberMatchingChallenge field.
func (o *Push1) SetUseNumberMatchingChallenge(v bool) {
	o.UseNumberMatchingChallenge = &v
}

func (o Push1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UseNumberMatchingChallenge != nil {
		toSerialize["useNumberMatchingChallenge"] = o.UseNumberMatchingChallenge
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Push1) UnmarshalJSON(bytes []byte) (err error) {
	varPush1 := _Push1{}

	err = json.Unmarshal(bytes, &varPush1)
	if err == nil {
		*o = Push1(varPush1)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "useNumberMatchingChallenge")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePush1 struct {
	value *Push1
	isSet bool
}

func (v NullablePush1) Get() *Push1 {
	return v.value
}

func (v *NullablePush1) Set(val *Push1) {
	v.value = val
	v.isSet = true
}

func (v NullablePush1) IsSet() bool {
	return v.isSet
}

func (v *NullablePush1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePush1(val *Push1) *NullablePush1 {
	return &NullablePush1{value: val, isSet: true}
}

func (v NullablePush1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePush1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

