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

// TokenSoftwareTotp1 Verifies an OTP for a `token:software:totp` factor
type TokenSoftwareTotp1 struct {
	// OTP for the current time window
	PassCode *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenSoftwareTotp1 TokenSoftwareTotp1

// NewTokenSoftwareTotp1 instantiates a new TokenSoftwareTotp1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenSoftwareTotp1() *TokenSoftwareTotp1 {
	this := TokenSoftwareTotp1{}
	return &this
}

// NewTokenSoftwareTotp1WithDefaults instantiates a new TokenSoftwareTotp1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenSoftwareTotp1WithDefaults() *TokenSoftwareTotp1 {
	this := TokenSoftwareTotp1{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *TokenSoftwareTotp1) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSoftwareTotp1) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *TokenSoftwareTotp1) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *TokenSoftwareTotp1) SetPassCode(v string) {
	o.PassCode = &v
}

func (o TokenSoftwareTotp1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenSoftwareTotp1) UnmarshalJSON(bytes []byte) (err error) {
	varTokenSoftwareTotp1 := _TokenSoftwareTotp1{}

	err = json.Unmarshal(bytes, &varTokenSoftwareTotp1)
	if err == nil {
		*o = TokenSoftwareTotp1(varTokenSoftwareTotp1)
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

type NullableTokenSoftwareTotp1 struct {
	value *TokenSoftwareTotp1
	isSet bool
}

func (v NullableTokenSoftwareTotp1) Get() *TokenSoftwareTotp1 {
	return v.value
}

func (v *NullableTokenSoftwareTotp1) Set(val *TokenSoftwareTotp1) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenSoftwareTotp1) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenSoftwareTotp1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenSoftwareTotp1(val *TokenSoftwareTotp1) *NullableTokenSoftwareTotp1 {
	return &NullableTokenSoftwareTotp1{value: val, isSet: true}
}

func (v NullableTokenSoftwareTotp1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenSoftwareTotp1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

