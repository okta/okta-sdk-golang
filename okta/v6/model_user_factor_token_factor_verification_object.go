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

// UserFactorTokenFactorVerificationObject struct for UserFactorTokenFactorVerificationObject
type UserFactorTokenFactorVerificationObject struct {
	// OTP for the next time window
	NextPassCode *string `json:"nextPassCode,omitempty"`
	// OTP for the current time window
	PassCode *string `json:"passCode,omitempty"`
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
	if o == nil || o.NextPassCode == nil {
		var ret string
		return ret
	}
	return *o.NextPassCode
}

// GetNextPassCodeOk returns a tuple with the NextPassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenFactorVerificationObject) GetNextPassCodeOk() (*string, bool) {
	if o == nil || o.NextPassCode == nil {
		return nil, false
	}
	return o.NextPassCode, true
}

// HasNextPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenFactorVerificationObject) HasNextPassCode() bool {
	if o != nil && o.NextPassCode != nil {
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
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenFactorVerificationObject) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *UserFactorTokenFactorVerificationObject) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *UserFactorTokenFactorVerificationObject) SetPassCode(v string) {
	o.PassCode = &v
}

func (o UserFactorTokenFactorVerificationObject) MarshalJSON() ([]byte, error) {
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

func (o *UserFactorTokenFactorVerificationObject) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorTokenFactorVerificationObject := _UserFactorTokenFactorVerificationObject{}

	err = json.Unmarshal(bytes, &varUserFactorTokenFactorVerificationObject)
	if err == nil {
		*o = UserFactorTokenFactorVerificationObject(varUserFactorTokenFactorVerificationObject)
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

