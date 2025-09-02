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

// Email1 Verifies an OTP sent by an `email` factor challenge. If you omit `passCode` in the request, a new OTP is sent to the phone.
type Email1 struct {
	// OTP for the current time window
	PassCode *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Email1 Email1

// NewEmail1 instantiates a new Email1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmail1() *Email1 {
	this := Email1{}
	return &this
}

// NewEmail1WithDefaults instantiates a new Email1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmail1WithDefaults() *Email1 {
	this := Email1{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *Email1) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email1) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *Email1) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *Email1) SetPassCode(v string) {
	o.PassCode = &v
}

func (o Email1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Email1) UnmarshalJSON(bytes []byte) (err error) {
	varEmail1 := _Email1{}

	err = json.Unmarshal(bytes, &varEmail1)
	if err == nil {
		*o = Email1(varEmail1)
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

type NullableEmail1 struct {
	value *Email1
	isSet bool
}

func (v NullableEmail1) Get() *Email1 {
	return v.value
}

func (v *NullableEmail1) Set(val *Email1) {
	v.value = val
	v.isSet = true
}

func (v NullableEmail1) IsSet() bool {
	return v.isSet
}

func (v *NullableEmail1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmail1(val *Email1) *NullableEmail1 {
	return &NullableEmail1{value: val, isSet: true}
}

func (v NullableEmail1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmail1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

