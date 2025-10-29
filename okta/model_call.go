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

// checks if the Call type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Call{}

// Call Attempts to activate a `call` factor with the specified passcode
type Call struct {
	// OTP for the current time window
	PassCode             *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Call Call

// NewCall instantiates a new Call object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCall() *Call {
	this := Call{}
	return &this
}

// NewCallWithDefaults instantiates a new Call object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallWithDefaults() *Call {
	this := Call{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *Call) GetPassCode() string {
	if o == nil || IsNil(o.PassCode) {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetPassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PassCode) {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *Call) HasPassCode() bool {
	if o != nil && !IsNil(o.PassCode) {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *Call) SetPassCode(v string) {
	o.PassCode = &v
}

func (o Call) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Call) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PassCode) {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Call) UnmarshalJSON(data []byte) (err error) {
	varCall := _Call{}

	err = json.Unmarshal(data, &varCall)

	if err != nil {
		return err
	}

	*o = Call(varCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "passCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCall struct {
	value *Call
	isSet bool
}

func (v NullableCall) Get() *Call {
	return v.value
}

func (v *NullableCall) Set(val *Call) {
	v.value = val
	v.isSet = true
}

func (v NullableCall) IsSet() bool {
	return v.isSet
}

func (v *NullableCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCall(val *Call) *NullableCall {
	return &NullableCall{value: val, isSet: true}
}

func (v NullableCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
