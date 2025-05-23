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

// PasswordPolicyRecoveryEmailProperties struct for PasswordPolicyRecoveryEmailProperties
type PasswordPolicyRecoveryEmailProperties struct {
	RecoveryToken *PasswordPolicyRecoveryEmailRecoveryToken `json:"recoveryToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryEmailProperties PasswordPolicyRecoveryEmailProperties

// NewPasswordPolicyRecoveryEmailProperties instantiates a new PasswordPolicyRecoveryEmailProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryEmailProperties() *PasswordPolicyRecoveryEmailProperties {
	this := PasswordPolicyRecoveryEmailProperties{}
	return &this
}

// NewPasswordPolicyRecoveryEmailPropertiesWithDefaults instantiates a new PasswordPolicyRecoveryEmailProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryEmailPropertiesWithDefaults() *PasswordPolicyRecoveryEmailProperties {
	this := PasswordPolicyRecoveryEmailProperties{}
	return &this
}

// GetRecoveryToken returns the RecoveryToken field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryEmailProperties) GetRecoveryToken() PasswordPolicyRecoveryEmailRecoveryToken {
	if o == nil || o.RecoveryToken == nil {
		var ret PasswordPolicyRecoveryEmailRecoveryToken
		return ret
	}
	return *o.RecoveryToken
}

// GetRecoveryTokenOk returns a tuple with the RecoveryToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryEmailProperties) GetRecoveryTokenOk() (*PasswordPolicyRecoveryEmailRecoveryToken, bool) {
	if o == nil || o.RecoveryToken == nil {
		return nil, false
	}
	return o.RecoveryToken, true
}

// HasRecoveryToken returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryEmailProperties) HasRecoveryToken() bool {
	if o != nil && o.RecoveryToken != nil {
		return true
	}

	return false
}

// SetRecoveryToken gets a reference to the given PasswordPolicyRecoveryEmailRecoveryToken and assigns it to the RecoveryToken field.
func (o *PasswordPolicyRecoveryEmailProperties) SetRecoveryToken(v PasswordPolicyRecoveryEmailRecoveryToken) {
	o.RecoveryToken = &v
}

func (o PasswordPolicyRecoveryEmailProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryToken != nil {
		toSerialize["recoveryToken"] = o.RecoveryToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyRecoveryEmailProperties) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyRecoveryEmailProperties := _PasswordPolicyRecoveryEmailProperties{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRecoveryEmailProperties)
	if err == nil {
		*o = PasswordPolicyRecoveryEmailProperties(varPasswordPolicyRecoveryEmailProperties)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "recoveryToken")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyRecoveryEmailProperties struct {
	value *PasswordPolicyRecoveryEmailProperties
	isSet bool
}

func (v NullablePasswordPolicyRecoveryEmailProperties) Get() *PasswordPolicyRecoveryEmailProperties {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryEmailProperties) Set(val *PasswordPolicyRecoveryEmailProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryEmailProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryEmailProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryEmailProperties(val *PasswordPolicyRecoveryEmailProperties) *NullablePasswordPolicyRecoveryEmailProperties {
	return &NullablePasswordPolicyRecoveryEmailProperties{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryEmailProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryEmailProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

