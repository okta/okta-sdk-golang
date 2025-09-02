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

// Sms1 Verifies an OTP sent by an `sms` factor challenge. If you omit `passCode` in the request, a new OTP is sent to the phone.
type Sms1 struct {
	// OTP for the current time window
	PassCode *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Sms1 Sms1

// NewSms1 instantiates a new Sms1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSms1() *Sms1 {
	this := Sms1{}
	return &this
}

// NewSms1WithDefaults instantiates a new Sms1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSms1WithDefaults() *Sms1 {
	this := Sms1{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *Sms1) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sms1) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *Sms1) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *Sms1) SetPassCode(v string) {
	o.PassCode = &v
}

func (o Sms1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Sms1) UnmarshalJSON(bytes []byte) (err error) {
	varSms1 := _Sms1{}

	err = json.Unmarshal(bytes, &varSms1)
	if err == nil {
		*o = Sms1(varSms1)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "passCode")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSms1 struct {
	value *Sms1
	isSet bool
}

func (v NullableSms1) Get() *Sms1 {
	return v.value
}

func (v *NullableSms1) Set(val *Sms1) {
	v.value = val
	v.isSet = true
}

func (v NullableSms1) IsSet() bool {
	return v.isSet
}

func (v *NullableSms1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSms1(val *Sms1) *NullableSms1 {
	return &NullableSms1{value: val, isSet: true}
}

func (v NullableSms1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSms1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

