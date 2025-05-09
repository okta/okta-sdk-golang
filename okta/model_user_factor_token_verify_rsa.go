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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserFactorTokenVerifyRSA struct for UserFactorTokenVerifyRSA
type UserFactorTokenVerifyRSA struct {
	// OTP for the current time window
	PassCode *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenVerifyRSA UserFactorTokenVerifyRSA

// NewUserFactorTokenVerifyRSA instantiates a new UserFactorTokenVerifyRSA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenVerifyRSA() *UserFactorTokenVerifyRSA {
	this := UserFactorTokenVerifyRSA{}
	return &this
}

// NewUserFactorTokenVerifyRSAWithDefaults instantiates a new UserFactorTokenVerifyRSA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenVerifyRSAWithDefaults() *UserFactorTokenVerifyRSA {
	this := UserFactorTokenVerifyRSA{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *UserFactorTokenVerifyRSA) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenVerifyRSA) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenVerifyRSA) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *UserFactorTokenVerifyRSA) SetPassCode(v string) {
	o.PassCode = &v
}

func (o UserFactorTokenVerifyRSA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorTokenVerifyRSA) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorTokenVerifyRSA := _UserFactorTokenVerifyRSA{}

	err = json.Unmarshal(bytes, &varUserFactorTokenVerifyRSA)
	if err == nil {
		*o = UserFactorTokenVerifyRSA(varUserFactorTokenVerifyRSA)
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

type NullableUserFactorTokenVerifyRSA struct {
	value *UserFactorTokenVerifyRSA
	isSet bool
}

func (v NullableUserFactorTokenVerifyRSA) Get() *UserFactorTokenVerifyRSA {
	return v.value
}

func (v *NullableUserFactorTokenVerifyRSA) Set(val *UserFactorTokenVerifyRSA) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenVerifyRSA) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenVerifyRSA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenVerifyRSA(val *UserFactorTokenVerifyRSA) *NullableUserFactorTokenVerifyRSA {
	return &NullableUserFactorTokenVerifyRSA{value: val, isSet: true}
}

func (v NullableUserFactorTokenVerifyRSA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenVerifyRSA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

