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

// checks if the TokenHotp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenHotp{}

// TokenHotp Verifies an OTP for a `token:hotp` factor
type TokenHotp struct {
	// OTP for the current time window
	PassCode             *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenHotp TokenHotp

// NewTokenHotp instantiates a new TokenHotp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenHotp() *TokenHotp {
	this := TokenHotp{}
	return &this
}

// NewTokenHotpWithDefaults instantiates a new TokenHotp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenHotpWithDefaults() *TokenHotp {
	this := TokenHotp{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *TokenHotp) GetPassCode() string {
	if o == nil || IsNil(o.PassCode) {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHotp) GetPassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PassCode) {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *TokenHotp) HasPassCode() bool {
	if o != nil && !IsNil(o.PassCode) {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *TokenHotp) SetPassCode(v string) {
	o.PassCode = &v
}

func (o TokenHotp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenHotp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PassCode) {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenHotp) UnmarshalJSON(data []byte) (err error) {
	varTokenHotp := _TokenHotp{}

	err = json.Unmarshal(data, &varTokenHotp)

	if err != nil {
		return err
	}

	*o = TokenHotp(varTokenHotp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "passCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenHotp struct {
	value *TokenHotp
	isSet bool
}

func (v NullableTokenHotp) Get() *TokenHotp {
	return v.value
}

func (v *NullableTokenHotp) Set(val *TokenHotp) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenHotp) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenHotp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenHotp(val *TokenHotp) *NullableTokenHotp {
	return &NullableTokenHotp{value: val, isSet: true}
}

func (v NullableTokenHotp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenHotp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
