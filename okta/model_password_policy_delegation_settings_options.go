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

// PasswordPolicyDelegationSettingsOptions struct for PasswordPolicyDelegationSettingsOptions
type PasswordPolicyDelegationSettingsOptions struct {
	SkipUnlock *bool `json:"skipUnlock,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyDelegationSettingsOptions PasswordPolicyDelegationSettingsOptions

// NewPasswordPolicyDelegationSettingsOptions instantiates a new PasswordPolicyDelegationSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyDelegationSettingsOptions() *PasswordPolicyDelegationSettingsOptions {
	this := PasswordPolicyDelegationSettingsOptions{}
	return &this
}

// NewPasswordPolicyDelegationSettingsOptionsWithDefaults instantiates a new PasswordPolicyDelegationSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyDelegationSettingsOptionsWithDefaults() *PasswordPolicyDelegationSettingsOptions {
	this := PasswordPolicyDelegationSettingsOptions{}
	return &this
}

// GetSkipUnlock returns the SkipUnlock field value if set, zero value otherwise.
func (o *PasswordPolicyDelegationSettingsOptions) GetSkipUnlock() bool {
	if o == nil || o.SkipUnlock == nil {
		var ret bool
		return ret
	}
	return *o.SkipUnlock
}

// GetSkipUnlockOk returns a tuple with the SkipUnlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyDelegationSettingsOptions) GetSkipUnlockOk() (*bool, bool) {
	if o == nil || o.SkipUnlock == nil {
		return nil, false
	}
	return o.SkipUnlock, true
}

// HasSkipUnlock returns a boolean if a field has been set.
func (o *PasswordPolicyDelegationSettingsOptions) HasSkipUnlock() bool {
	if o != nil && o.SkipUnlock != nil {
		return true
	}

	return false
}

// SetSkipUnlock gets a reference to the given bool and assigns it to the SkipUnlock field.
func (o *PasswordPolicyDelegationSettingsOptions) SetSkipUnlock(v bool) {
	o.SkipUnlock = &v
}

func (o PasswordPolicyDelegationSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkipUnlock != nil {
		toSerialize["skipUnlock"] = o.SkipUnlock
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyDelegationSettingsOptions) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyDelegationSettingsOptions := _PasswordPolicyDelegationSettingsOptions{}

	err = json.Unmarshal(bytes, &varPasswordPolicyDelegationSettingsOptions)
	if err == nil {
		*o = PasswordPolicyDelegationSettingsOptions(varPasswordPolicyDelegationSettingsOptions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "skipUnlock")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyDelegationSettingsOptions struct {
	value *PasswordPolicyDelegationSettingsOptions
	isSet bool
}

func (v NullablePasswordPolicyDelegationSettingsOptions) Get() *PasswordPolicyDelegationSettingsOptions {
	return v.value
}

func (v *NullablePasswordPolicyDelegationSettingsOptions) Set(val *PasswordPolicyDelegationSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyDelegationSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyDelegationSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyDelegationSettingsOptions(val *PasswordPolicyDelegationSettingsOptions) *NullablePasswordPolicyDelegationSettingsOptions {
	return &NullablePasswordPolicyDelegationSettingsOptions{value: val, isSet: true}
}

func (v NullablePasswordPolicyDelegationSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyDelegationSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

