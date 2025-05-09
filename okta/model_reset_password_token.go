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

// ResetPasswordToken struct for ResetPasswordToken
type ResetPasswordToken struct {
	ResetPasswordUrl *string `json:"resetPasswordUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResetPasswordToken ResetPasswordToken

// NewResetPasswordToken instantiates a new ResetPasswordToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordToken() *ResetPasswordToken {
	this := ResetPasswordToken{}
	return &this
}

// NewResetPasswordTokenWithDefaults instantiates a new ResetPasswordToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordTokenWithDefaults() *ResetPasswordToken {
	this := ResetPasswordToken{}
	return &this
}

// GetResetPasswordUrl returns the ResetPasswordUrl field value if set, zero value otherwise.
func (o *ResetPasswordToken) GetResetPasswordUrl() string {
	if o == nil || o.ResetPasswordUrl == nil {
		var ret string
		return ret
	}
	return *o.ResetPasswordUrl
}

// GetResetPasswordUrlOk returns a tuple with the ResetPasswordUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordToken) GetResetPasswordUrlOk() (*string, bool) {
	if o == nil || o.ResetPasswordUrl == nil {
		return nil, false
	}
	return o.ResetPasswordUrl, true
}

// HasResetPasswordUrl returns a boolean if a field has been set.
func (o *ResetPasswordToken) HasResetPasswordUrl() bool {
	if o != nil && o.ResetPasswordUrl != nil {
		return true
	}

	return false
}

// SetResetPasswordUrl gets a reference to the given string and assigns it to the ResetPasswordUrl field.
func (o *ResetPasswordToken) SetResetPasswordUrl(v string) {
	o.ResetPasswordUrl = &v
}

func (o ResetPasswordToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResetPasswordUrl != nil {
		toSerialize["resetPasswordUrl"] = o.ResetPasswordUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResetPasswordToken) UnmarshalJSON(bytes []byte) (err error) {
	varResetPasswordToken := _ResetPasswordToken{}

	err = json.Unmarshal(bytes, &varResetPasswordToken)
	if err == nil {
		*o = ResetPasswordToken(varResetPasswordToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resetPasswordUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResetPasswordToken struct {
	value *ResetPasswordToken
	isSet bool
}

func (v NullableResetPasswordToken) Get() *ResetPasswordToken {
	return v.value
}

func (v *NullableResetPasswordToken) Set(val *ResetPasswordToken) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordToken) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordToken(val *ResetPasswordToken) *NullableResetPasswordToken {
	return &NullableResetPasswordToken{value: val, isSet: true}
}

func (v NullableResetPasswordToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

