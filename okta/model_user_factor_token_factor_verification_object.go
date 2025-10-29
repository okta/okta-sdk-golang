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

// checks if the UserFactorTokenFactorVerificationObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorTokenFactorVerificationObject{}

// UserFactorTokenFactorVerificationObject struct for UserFactorTokenFactorVerificationObject
type UserFactorTokenFactorVerificationObject struct {
	// OTP for the next time window
	NextPassCode *string `json:"nextPassCode,omitempty"`
	// OTP for the current time window
	PassCode             *string `json:"passCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenFactorVerificationObject UserFactorTokenFactorVerificationObject

// NewUserFactorTokenFactorVerificationObject instantiates a new UserFactorTokenFactorVerificationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenFactorVerificationObject() *UserFactorTokenFactorVerificationObject {
	this := UserFactorTokenFactorVerificationObject{}
	return &this
}

// NewUserFactorTokenFactorVerificationObjectWithDefaults instantiates a new UserFactorTokenFactorVerificationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenFactorVerificationObjectWithDefaults() *UserFactorTokenFactorVerificationObject {
	this := UserFactorTokenFactorVerificationObject{}
	return &this
}

// GetNextPassCode returns the NextPassCode field value if set, zero value otherwise.
func (o *UserFactorTokenFactorVerificationObject) GetNextPassCode() string {
	if o == nil || IsNil(o.NextPassCode) {
		var ret string
		return ret
	}
	return *o.NextPassCode
}

// GetNextPassCodeOk returns a tuple with the NextPassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenFactorVerificationObject) GetNextPassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.NextPassCode) {
		return nil, false
	}
	return o.NextPassCode, true
}

// HasNextPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenFactorVerificationObject) HasNextPassCode() bool {
	if o != nil && !IsNil(o.NextPassCode) {
		return true
	}

	return false
}

// SetNextPassCode gets a reference to the given string and assigns it to the NextPassCode field.
func (o *UserFactorTokenFactorVerificationObject) SetNextPassCode(v string) {
	o.NextPassCode = &v
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *UserFactorTokenFactorVerificationObject) GetPassCode() string {
	if o == nil || IsNil(o.PassCode) {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenFactorVerificationObject) GetPassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PassCode) {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenFactorVerificationObject) HasPassCode() bool {
	if o != nil && !IsNil(o.PassCode) {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *UserFactorTokenFactorVerificationObject) SetPassCode(v string) {
	o.PassCode = &v
}

func (o UserFactorTokenFactorVerificationObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorTokenFactorVerificationObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPassCode) {
		toSerialize["nextPassCode"] = o.NextPassCode
	}
	if !IsNil(o.PassCode) {
		toSerialize["passCode"] = o.PassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorTokenFactorVerificationObject) UnmarshalJSON(data []byte) (err error) {
	varUserFactorTokenFactorVerificationObject := _UserFactorTokenFactorVerificationObject{}

	err = json.Unmarshal(data, &varUserFactorTokenFactorVerificationObject)

	if err != nil {
		return err
	}

	*o = UserFactorTokenFactorVerificationObject(varUserFactorTokenFactorVerificationObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nextPassCode")
		delete(additionalProperties, "passCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorTokenFactorVerificationObject struct {
	value *UserFactorTokenFactorVerificationObject
	isSet bool
}

func (v NullableUserFactorTokenFactorVerificationObject) Get() *UserFactorTokenFactorVerificationObject {
	return v.value
}

func (v *NullableUserFactorTokenFactorVerificationObject) Set(val *UserFactorTokenFactorVerificationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenFactorVerificationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenFactorVerificationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenFactorVerificationObject(val *UserFactorTokenFactorVerificationObject) *NullableUserFactorTokenFactorVerificationObject {
	return &NullableUserFactorTokenFactorVerificationObject{value: val, isSet: true}
}

func (v NullableUserFactorTokenFactorVerificationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenFactorVerificationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
