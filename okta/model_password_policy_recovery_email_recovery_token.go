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

// checks if the PasswordPolicyRecoveryEmailRecoveryToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoveryEmailRecoveryToken{}

// PasswordPolicyRecoveryEmailRecoveryToken struct for PasswordPolicyRecoveryEmailRecoveryToken
type PasswordPolicyRecoveryEmailRecoveryToken struct {
	// Lifetime (in minutes) of the recovery token
	TokenLifetimeMinutes *int32 `json:"tokenLifetimeMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryEmailRecoveryToken PasswordPolicyRecoveryEmailRecoveryToken

// NewPasswordPolicyRecoveryEmailRecoveryToken instantiates a new PasswordPolicyRecoveryEmailRecoveryToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryEmailRecoveryToken() *PasswordPolicyRecoveryEmailRecoveryToken {
	this := PasswordPolicyRecoveryEmailRecoveryToken{}
	var tokenLifetimeMinutes int32 = 10080
	this.TokenLifetimeMinutes = &tokenLifetimeMinutes
	return &this
}

// NewPasswordPolicyRecoveryEmailRecoveryTokenWithDefaults instantiates a new PasswordPolicyRecoveryEmailRecoveryToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryEmailRecoveryTokenWithDefaults() *PasswordPolicyRecoveryEmailRecoveryToken {
	this := PasswordPolicyRecoveryEmailRecoveryToken{}
	var tokenLifetimeMinutes int32 = 10080
	this.TokenLifetimeMinutes = &tokenLifetimeMinutes
	return &this
}

// GetTokenLifetimeMinutes returns the TokenLifetimeMinutes field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryEmailRecoveryToken) GetTokenLifetimeMinutes() int32 {
	if o == nil || IsNil(o.TokenLifetimeMinutes) {
		var ret int32
		return ret
	}
	return *o.TokenLifetimeMinutes
}

// GetTokenLifetimeMinutesOk returns a tuple with the TokenLifetimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryEmailRecoveryToken) GetTokenLifetimeMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenLifetimeMinutes) {
		return nil, false
	}
	return o.TokenLifetimeMinutes, true
}

// HasTokenLifetimeMinutes returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryEmailRecoveryToken) HasTokenLifetimeMinutes() bool {
	if o != nil && !IsNil(o.TokenLifetimeMinutes) {
		return true
	}

	return false
}

// SetTokenLifetimeMinutes gets a reference to the given int32 and assigns it to the TokenLifetimeMinutes field.
func (o *PasswordPolicyRecoveryEmailRecoveryToken) SetTokenLifetimeMinutes(v int32) {
	o.TokenLifetimeMinutes = &v
}

func (o PasswordPolicyRecoveryEmailRecoveryToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoveryEmailRecoveryToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenLifetimeMinutes) {
		toSerialize["tokenLifetimeMinutes"] = o.TokenLifetimeMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoveryEmailRecoveryToken) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoveryEmailRecoveryToken := _PasswordPolicyRecoveryEmailRecoveryToken{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoveryEmailRecoveryToken)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoveryEmailRecoveryToken(varPasswordPolicyRecoveryEmailRecoveryToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tokenLifetimeMinutes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoveryEmailRecoveryToken struct {
	value *PasswordPolicyRecoveryEmailRecoveryToken
	isSet bool
}

func (v NullablePasswordPolicyRecoveryEmailRecoveryToken) Get() *PasswordPolicyRecoveryEmailRecoveryToken {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryEmailRecoveryToken) Set(val *PasswordPolicyRecoveryEmailRecoveryToken) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryEmailRecoveryToken) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryEmailRecoveryToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryEmailRecoveryToken(val *PasswordPolicyRecoveryEmailRecoveryToken) *NullablePasswordPolicyRecoveryEmailRecoveryToken {
	return &NullablePasswordPolicyRecoveryEmailRecoveryToken{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryEmailRecoveryToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryEmailRecoveryToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
