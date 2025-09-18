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

// checks if the Sms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sms{}

// Sms Attempts to activate an `sms` factor with the specified passcode
type Sms struct {
	// OTP for the current time window
	PassCode             *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Sms Sms

// NewSms instantiates a new Sms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSms() *Sms {
	this := Sms{}
	return &this
}

// NewSmsWithDefaults instantiates a new Sms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsWithDefaults() *Sms {
	this := Sms{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *Sms) GetPassCode() string {
	if o == nil || IsNil(o.PassCode) {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sms) GetPassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PassCode) {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *Sms) HasPassCode() bool {
	if o != nil && !IsNil(o.PassCode) {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *Sms) SetPassCode(v string) {
	o.PassCode = &v
}

func (o Sms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PassCode) {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Sms) UnmarshalJSON(data []byte) (err error) {
	varSms := _Sms{}

	err = json.Unmarshal(data, &varSms)

	if err != nil {
		return err
	}

	*o = Sms(varSms)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "passCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSms struct {
	value *Sms
	isSet bool
}

func (v NullableSms) Get() *Sms {
	return v.value
}

func (v *NullableSms) Set(val *Sms) {
	v.value = val
	v.isSet = true
}

func (v NullableSms) IsSet() bool {
	return v.isSet
}

func (v *NullableSms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSms(val *Sms) *NullableSms {
	return &NullableSms{value: val, isSet: true}
}

func (v NullableSms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
