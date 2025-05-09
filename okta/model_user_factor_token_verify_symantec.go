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

// UserFactorTokenVerifySymantec struct for UserFactorTokenVerifySymantec
type UserFactorTokenVerifySymantec struct {
	// OTP for the next time window
	NextPassCode *int32 `json:"nextPassCode,omitempty"`
	// OTP for the current time window
	PassCode *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenVerifySymantec UserFactorTokenVerifySymantec

// NewUserFactorTokenVerifySymantec instantiates a new UserFactorTokenVerifySymantec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenVerifySymantec() *UserFactorTokenVerifySymantec {
	this := UserFactorTokenVerifySymantec{}
	return &this
}

// NewUserFactorTokenVerifySymantecWithDefaults instantiates a new UserFactorTokenVerifySymantec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenVerifySymantecWithDefaults() *UserFactorTokenVerifySymantec {
	this := UserFactorTokenVerifySymantec{}
	return &this
}

// GetNextPassCode returns the NextPassCode field value if set, zero value otherwise.
func (o *UserFactorTokenVerifySymantec) GetNextPassCode() int32 {
	if o == nil || o.NextPassCode == nil {
		var ret int32
		return ret
	}
	return *o.NextPassCode
}

// GetNextPassCodeOk returns a tuple with the NextPassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenVerifySymantec) GetNextPassCodeOk() (*int32, bool) {
	if o == nil || o.NextPassCode == nil {
		return nil, false
	}
	return o.NextPassCode, true
}

// HasNextPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenVerifySymantec) HasNextPassCode() bool {
	if o != nil && o.NextPassCode != nil {
		return true
	}

	return false
}

// SetNextPassCode gets a reference to the given int32 and assigns it to the NextPassCode field.
func (o *UserFactorTokenVerifySymantec) SetNextPassCode(v int32) {
	o.NextPassCode = &v
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *UserFactorTokenVerifySymantec) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenVerifySymantec) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenVerifySymantec) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *UserFactorTokenVerifySymantec) SetPassCode(v string) {
	o.PassCode = &v
}

func (o UserFactorTokenVerifySymantec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPassCode != nil {
		toSerialize["nextPassCode"] = o.NextPassCode
	}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorTokenVerifySymantec) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorTokenVerifySymantec := _UserFactorTokenVerifySymantec{}

	err = json.Unmarshal(bytes, &varUserFactorTokenVerifySymantec)
	if err == nil {
		*o = UserFactorTokenVerifySymantec(varUserFactorTokenVerifySymantec)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "nextPassCode")
		delete(additionalProperties, "passCode")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorTokenVerifySymantec struct {
	value *UserFactorTokenVerifySymantec
	isSet bool
}

func (v NullableUserFactorTokenVerifySymantec) Get() *UserFactorTokenVerifySymantec {
	return v.value
}

func (v *NullableUserFactorTokenVerifySymantec) Set(val *UserFactorTokenVerifySymantec) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenVerifySymantec) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenVerifySymantec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenVerifySymantec(val *UserFactorTokenVerifySymantec) *NullableUserFactorTokenVerifySymantec {
	return &NullableUserFactorTokenVerifySymantec{value: val, isSet: true}
}

func (v NullableUserFactorTokenVerifySymantec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenVerifySymantec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

